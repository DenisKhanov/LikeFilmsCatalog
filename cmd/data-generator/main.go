// Структура для хранения данных о фильмах
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Movie представляет информацию о фильме
type Movie struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Year        int    `json:"year"`
	Category    string `json:"category"`
	Description string `json:"description"`
	ImagePath   string `json:"imagePath"`
}

// Категории фильмов
var categories = map[string]string{
	"drama":      "Драма",
	"comedy":     "Комедия",
	"fantasy":    "Фантастика и фэнтези",
	"thriller":   "Триллер и детектив",
	"biography":  "Биографический",
	"historical": "Исторический и военный",
	"melodrama":  "Мелодрама",
}

// Функция для создания структуры данных фильмов из списка пользователя
func main() {
	// Создаем каталог для изображений, если его нет
	imagesDir := filepath.Join("static", "images", "movies")
	if err := os.MkdirAll(imagesDir, 0755); err != nil {
		fmt.Printf("Ошибка при создании каталога для изображений: %v\n", err)
		return
	}

	// Создаем карту для хранения фильмов по категориям
	moviesByCategory := make(map[string][]Movie)

	// Драма
	dramaMovies := []Movie{
		{ID: "arthur-king", Title: "Артур, ты король", Year: 2024, Category: "drama", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "white-bird", Title: "Белая птица: Новое чудо", Year: 2023, Category: "drama", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "eternal-sunshine", Title: "Вечное сияние чистого разума", Year: 2004, Category: "drama", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "beautiful-mind", Title: "Игры разума", Year: 2001, Category: "drama", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "centaur", Title: "Кентавр", Year: 2023, Category: "drama", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "million-dollar-baby", Title: "Малышка на миллион", Year: 2004, Category: "drama", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "all-quiet", Title: "На западном фронте без перемен", Year: 2022, Category: "drama", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "caddo-lake", Title: "Озеро Каддо", Year: 2024, Category: "drama", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "one-life", Title: "Одна жизнь", Year: 2023, Category: "drama", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "first-day", Title: "Первый день моей жизни", Year: 2023, Category: "drama", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "seven-years-tibet", Title: "Семь лет в Тибете", Year: 1997, Category: "drama", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "intertwined-fates", Title: "Сплетение судеб", Year: 2023, Category: "drama", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "killers-flower-moon", Title: "Убийцы цветочной луны", Year: 2023, Category: "drama", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "good-will-hunting", Title: "Умница Уилл Хантинг", Year: 1997, Category: "drama", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "hitman-last-job", Title: "Хитмен. Последнее дело", Year: 2023, Category: "drama", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "dogman", Title: "Догмен", Year: 2023, Category: "drama", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "count-monte-cristo", Title: "Граф Монте-Кристо", Year: 2024, Category: "drama", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "gladiator", Title: "Гладиатор", Year: 2000, Category: "drama", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "awakening", Title: "Пробуждение", Year: 1990, Category: "drama", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "pianist", Title: "Пианист", Year: 2002, Category: "drama", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "erin-brockovich", Title: "Эрин Брокович", Year: 2000, Category: "drama", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "no-family", Title: "Без семьи", Year: 2018, Category: "drama", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "cast-away", Title: "Изгой", Year: 2000, Category: "drama", ImagePath: "/static/images/movies/placeholder.jpg"},
	}
	moviesByCategory["drama"] = dramaMovies

	// Комедия
	comedyMovies := []Movie{
		{ID: "chef-battle", Title: "Битва шефов", Year: 2023, Category: "comedy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "jerry-marge", Title: "Джерри и Мардж играют по-крупному", Year: 2022, Category: "comedy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "kitchen-stars", Title: "Кухня со звездами", Year: 2023, Category: "comedy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "mr-blake", Title: "Мистер Блейк к вашим услугам", Year: 2023, Category: "comedy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "terrible-neighbor", Title: "Ужасный сосед", Year: 2023, Category: "comedy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "change-up", Title: "Хочу как ты", Year: 2011, Category: "comedy", ImagePath: "/static/images/movies/placeholder.jpg"},
	}
	moviesByCategory["comedy"] = comedyMovies

	// Фантастика и фэнтези
	fantasyMovies := []Movie{
		{ID: "avatar-2", Title: "Аватар 2", Year: 2022, Category: "fantasy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "edge-of-tomorrow", Title: "Грань будущего", Year: 2014, Category: "fantasy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "deja-vu", Title: "Дежавю", Year: 2006, Category: "fantasy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "dune", Title: "Дюна", Year: 2021, Category: "fantasy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "dune-2", Title: "Дюна 2", Year: 2024, Category: "fantasy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "interstellar", Title: "Интерстеллар", Year: 2014, Category: "fantasy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "source-code", Title: "Исходный код", Year: 2011, Category: "fantasy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "inception", Title: "Начало", Year: 2010, Category: "fantasy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "oblivion", Title: "Обливион", Year: 2013, Category: "fantasy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "passengers", Title: "Пассажиры", Year: 2016, Category: "fantasy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "sunshine", Title: "Пекло", Year: 2007, Category: "fantasy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "dungeons-dragons", Title: "Подземелья и драконы: честь среди воров", Year: 2023, Category: "fantasy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "paradise", Title: "Рай земной", Year: 2023, Category: "fantasy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "dreamland", Title: "Страна снов", Year: 2022, Category: "fantasy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "dark-reflections", Title: "Темные отражения", Year: 2018, Category: "fantasy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "real-steel", Title: "Живая сталь", Year: 2011, Category: "fantasy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "tenet", Title: "Довод", Year: 2020, Category: "fantasy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "martian", Title: "Марсианин", Year: 2015, Category: "fantasy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "jacket", Title: "Пиджак", Year: 2004, Category: "fantasy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "pandorum", Title: "Пандорум", Year: 2009, Category: "fantasy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "contact", Title: "Контакт", Year: 1997, Category: "fantasy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "moon", Title: "Луна 2112", Year: 2009, Category: "fantasy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "abyss", Title: "Бездна", Year: 1989, Category: "fantasy", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "gattaca", Title: "Гаттака", Year: 1997, Category: "fantasy", ImagePath: "/static/images/movies/placeholder.jpg"},
	}
	moviesByCategory["fantasy"] = fantasyMovies

	// Триллер и детектив
	thrillerMovies := []Movie{
		{ID: "hypnotic", Title: "Гипнотик", Year: 2023, Category: "thriller", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "two-three-demon", Title: "Два, три, демон приди", Year: 2022, Category: "thriller", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "catch-me", Title: "Поймай меня, если сможешь", Year: 2002, Category: "thriller", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "stalker", Title: "Сталкер", Year: 2023, Category: "thriller", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "stop-word", Title: "Стоп слово", Year: 2023, Category: "thriller", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "gods-creation", Title: "Творение Господне", Year: 2017, Category: "thriller", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "trap", Title: "Ловушка", Year: 2024, Category: "thriller", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "leave-the-world-behind", Title: "Весь мир позади", Year: 2023, Category: "thriller", ImagePath: "/static/images/movies/placeholder.jpg"},
	}
	moviesByCategory["thriller"] = thrillerMovies

	// Биографический
	biographyMovies := []Movie{
		{ID: "gran-turismo", Title: "Gran Turismo", Year: 2023, Category: "biography", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "how-to-hack-exam", Title: "Как взломать экзамен", Year: 2024, Category: "biography", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "snow-brotherhood", Title: "Снежное братство", Year: 2023, Category: "biography", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "tetris", Title: "Тетрис", Year: 2023, Category: "biography", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "green-book", Title: "Зеленая книга", Year: 2018, Category: "biography", ImagePath: "/static/images/movies/placeholder.jpg"},
	}
	moviesByCategory["biography"] = biographyMovies

	// Исторический и военный
	historicalMovies := []Movie{
		{ID: "no-answer", Title: "Без ответа", Year: 2023, Category: "historical", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "bandit", Title: "Бандит", Year: 2022, Category: "historical", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "left-behind", Title: "Оставленные", Year: 2023, Category: "historical", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "fires", Title: "Пожары", Year: 2010, Category: "historical", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "300", Title: "300 спартанцев", Year: 2006, Category: "historical", ImagePath: "/static/images/movies/placeholder.jpg"},
		{ID: "all-quiet-western", Title: "На Западном фронте без перемен", Year: 2022, Category: "historical", ImagePath: "/static/images/movies/placeholder.jpg"},
	}
	moviesByCategory["historical"] = historicalMovies

	// Мелодрама
	melodramaMovies := []Movie{
		{ID: "serendipity", Title: "Интуиция", Year: 2001, Category: "melodrama", ImagePath: "/static/images/movies/placeholder.jpg"},
	}
	moviesByCategory["melodrama"] = melodramaMovies

	// Сохраняем данные в JSON файл
	jsonData, err := json.MarshalIndent(moviesByCategory, "", "  ")
	if err != nil {
		fmt.Printf("Ошибка при маршалинге JSON: %v\n", err)
		return
	}

	// Создаем каталог для данных, если его нет
	dataDir := filepath.Join("static", "data")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		fmt.Printf("Ошибка при создании каталога для данных: %v\n", err)
		return
	}

	// Записываем JSON в файл
	jsonFilePath := filepath.Join(dataDir, "movies.json")
	if err := ioutil.WriteFile(jsonFilePath, jsonData, 0644); err != nil {
		fmt.Printf("Ошибка при записи JSON файла: %v\n", err)
		return
	}

	// Создаем заглушку для изображений
	placeholderContent := []byte(`<svg width="300" height="450" xmlns="http://www.w3.org/2000/svg">
		<rect width="300" height="450" fill="#2d3748"/>
		<text x="150" y="225" font-family="Arial" font-size="24" fill="#a0aec0" text-anchor="middle">Изображение фильма</text>
	</svg>`)

	placeholderPath := filepath.Join(imagesDir, "placeholder.jpg")
	if err := ioutil.WriteFile(placeholderPath, placeholderContent, 0644); err != nil {
		fmt.Printf("Ошибка при создании заглушки для изображений: %v\n", err)
		return
	}

	fmt.Println("Структура данных фильмов успешно создана и сохранена в", jsonFilePath)
}
