package main

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
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
	Link            string `json:"link"` // Добавлено для поддержки ссылок
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
	templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Ошибка при загрузке шаблонов: %v", err)
	}

	// Настройка Gin
	gin.SetMode(gin.ReleaseMode) // Используйте gin.DebugMode для отладки
	router := gin.Default()

	// Статические файлы
	router.Static("/static", "./static")

	// Маршруты
	router.GET("/", handleIndex)
	router.GET("/movies", handleMovies)
	router.GET("/category/:category", handleCategory)

	// API маршруты
	router.GET("/api/movies", handleAPIMovies)
	router.GET("/api/movies/:category", handleAPIMoviesByCategory)
	router.GET("/api/movie/:id", handleAPIMovie)

	// Настройка HTTP-сервера
	filmsServer := &http.Server{
		Addr:    ":8080", // Можно изменить на "localhost:8080", если нужен только локальный доступ
		Handler: router,
	}

	// Запуск сервера в горутине
	go func() {
		log.Printf("Сервер запущен на http://localhost%s", filmsServer.Addr)
		if err := filmsServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка запуска HTTP-сервера: %v", err)
		}
	}()

	// Ожидание сигнала завершения
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-signalChan
	log.Printf("Получен сигнал завершения: %v", sig)

	// Graceful shutdown с таймаутом 5 секунд
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := filmsServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("Ошибка при остановке HTTP-сервера: %v", err)
	} else {
		log.Println("HTTP-сервер успешно остановлен")
	}

	log.Println("Сервер завершил работу")
}

// Обработчик главной страницы
func handleIndex(c *gin.Context) {
	if c.Request.URL.Path != "/" {
		c.Status(http.StatusNotFound)
		return
	}

	data := map[string]interface{}{
		"title": "Каталог фильмов",
		"page":  "index",
	}

	err := templates.ExecuteTemplate(c.Writer, "base.html", data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка рендеринга шаблона: %v", err)
	}
}

// Обработчик страницы со всеми фильмами
func handleMovies(c *gin.Context) {
	data := map[string]interface{}{
		"title": "Все фильмы",
		"page":  "movies",
	}

	err := templates.ExecuteTemplate(c.Writer, "base.html", data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка рендеринга шаблона: %v", err)
	}
}

// Обработчик страницы категории
func handleCategory(c *gin.Context) {
	category := c.Param("category")
	categoryTitle, ok := categoryNames[category]
	if !ok {
		categoryTitle = category
	}

	data := map[string]interface{}{
		"title":    categoryTitle,
		"category": category,
		"pages":    "category", // Исправлено на "page" для консистентности
	}

	err := templates.ExecuteTemplate(c.Writer, "base.html", data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка рендеринга шаблона: %v", err)
	}
}

// Обработчик API для получения всех фильмов
func handleAPIMovies(c *gin.Context) {
	jsonFilePath := filepath.Join("static", "data", "movies.json")
	jsonData, err := os.ReadFile(jsonFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось загрузить данные о фильмах"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.Writer.Write(jsonData)
}

// Обработчик API для получения фильмов по категории
func handleAPIMoviesByCategory(c *gin.Context) {
	category := c.Param("category")

	jsonFilePath := filepath.Join("static", "data", "movies.json")
	jsonData, err := os.ReadFile(jsonFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось загрузить данные о фильмах"})
		return
	}

	var moviesByCategory map[string][]MovieWithDescription
	err = json.Unmarshal(jsonData, &moviesByCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обработке данных о фильмах"})
		return
	}

	movies, ok := moviesByCategory[category]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Категория не найдена"})
		return
	}

	c.JSON(http.StatusOK, movies)
}

// Обработчик API для получения информации о конкретном фильме
func handleAPIMovie(c *gin.Context) {
	movieID := c.Param("id")

	jsonFilePath := filepath.Join("static", "data", "movies.json")
	jsonData, err := os.ReadFile(jsonFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось загрузить данные о фильмах"})
		return
	}

	var moviesByCategory map[string][]MovieWithDescription
	err = json.Unmarshal(jsonData, &moviesByCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обработке данных о фильмах"})
		return
	}

	for _, movies := range moviesByCategory {
		for _, movie := range movies {
			if movie.ID == movieID {
				c.JSON(http.StatusOK, movie)
				return
			}
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Фильм не найден"})
}
