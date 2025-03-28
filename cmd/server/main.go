package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

// MovieWithDescription представляет информацию о фильме
type MovieWithDescription struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	Year            int    `json:"year"`
	Category        string `json:"category"`
	Description     string `json:"description"`
	ImagePath       string `json:"imagePath"`
	FullDescription string `json:"fullDescription"`
}

// Категории фильмов
var categoryNames = map[string]string{
	"drama":      "Драма",
	"comedy":     "Комедия",
	"fantasy":    "Фантастика и фэнтези",
	"thriller":   "Триллер и детектив",
	"biography":  "Биографический",
	"historical": "Исторический и военный",
	"melodrama":  "Мелодрама",
}

// Шаблоны
var templates *template.Template

func main() {
	// Загружаем шаблоны
	var err error
	templates, err = template.ParseGlob("templates/*/*.html")

	//templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Ошибка при загрузке шаблонов: %v", err)
	}

	// Статические файлы
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Маршруты
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/movies", handleMovies)
	http.HandleFunc("/category/", handleCategory)

	// API маршруты
	http.HandleFunc("/api/movies", handleAPIMovies)
	http.HandleFunc("/api/movies/", handleAPIMoviesByCategory)
	http.HandleFunc("/api/movie/", handleAPIMovie)

	// Запуск сервера
	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Обработчик главной страницы
func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	data := map[string]interface{}{
		"title": "Каталог фильмов",
	}

	err := templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Обработчик страницы со всеми фильмами
func handleMovies(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"title": "Все фильмы",
	}

	err := templates.ExecuteTemplate(w, "movies.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Обработчик страницы категории
func handleCategory(w http.ResponseWriter, r *http.Request) {
	category := strings.TrimPrefix(r.URL.Path, "/category/")
	categoryTitle, ok := categoryNames[category]
	if !ok {
		categoryTitle = category
	}

	data := map[string]interface{}{
		"title":    categoryTitle,
		"category": category,
	}

	err := templates.ExecuteTemplate(w, "movies.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Обработчик API для получения всех фильмов
func handleAPIMovies(w http.ResponseWriter, r *http.Request) {
	// Загружаем данные о фильмах из JSON файла
	jsonFilePath := filepath.Join("static", "data", "movies.json")
	jsonData, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		http.Error(w, "Не удалось загрузить данные о фильмах", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Отправляем JSON данные
	w.Write(jsonData)
}

// Обработчик API для получения фильмов по категории
func handleAPIMoviesByCategory(w http.ResponseWriter, r *http.Request) {
	category := strings.TrimPrefix(r.URL.Path, "/api/movies/")

	// Загружаем данные о фильмах из JSON файла
	jsonFilePath := filepath.Join("static", "data", "movies.json")
	jsonData, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		http.Error(w, "Не удалось загрузить данные о фильмах", http.StatusInternalServerError)
		return
	}

	// Распаковываем JSON в карту фильмов
	var moviesByCategory map[string][]MovieWithDescription
	err = json.Unmarshal(jsonData, &moviesByCategory)
	if err != nil {
		http.Error(w, "Ошибка при обработке данных о фильмах", http.StatusInternalServerError)
		return
	}

	// Получаем фильмы по категории
	movies, ok := moviesByCategory[category]
	if !ok {
		http.Error(w, "Категория не найдена", http.StatusNotFound)
		return
	}

	// Сериализуем фильмы в JSON
	moviesJSON, err := json.Marshal(movies)
	if err != nil {
		http.Error(w, "Ошибка при сериализации данных", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Отправляем JSON данные
	w.Write(moviesJSON)
}

// Обработчик API для получения информации о конкретном фильме
func handleAPIMovie(w http.ResponseWriter, r *http.Request) {
	movieID := strings.TrimPrefix(r.URL.Path, "/api/movie/")

	// Загружаем данные о фильмах из JSON файла
	jsonFilePath := filepath.Join("static", "data", "movies.json")
	jsonData, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		http.Error(w, "Не удалось загрузить данные о фильмах", http.StatusInternalServerError)
		return
	}

	// Распаковываем JSON в карту фильмов
	var moviesByCategory map[string][]MovieWithDescription
	err = json.Unmarshal(jsonData, &moviesByCategory)
	if err != nil {
		http.Error(w, "Ошибка при обработке данных о фильмах", http.StatusInternalServerError)
		return
	}

	// Ищем фильм по ID во всех категориях
	for _, movies := range moviesByCategory {
		for _, movie := range movies {
			if movie.ID == movieID {
				// Сериализуем фильм в JSON
				movieJSON, err := json.Marshal(movie)
				if err != nil {
					http.Error(w, "Ошибка при сериализации данных", http.StatusInternalServerError)
					return
				}

				// Устанавливаем заголовок Content-Type
				w.Header().Set("Content-Type", "application/json")

				// Отправляем JSON данные
				w.Write(movieJSON)
				return
			}
		}
	}

	http.Error(w, "Фильм не найден", http.StatusNotFound)
}
