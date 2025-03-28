package models

// Movie представляет информацию о фильме
type Movie struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Year        int    `json:"year"`
	Category    string `json:"category"`
	Description string `json:"description"`
	ImagePath   string `json:"imagePath"`
}

// GetCategories возвращает список всех категорий фильмов
func GetCategories() []string {
	return []string{
		"Драма",
		"Комедия",
		"Фантастика и фэнтези",
		"Триллер и детектив",
		"Биографический",
		"Исторический и военный",
		"Мелодрама",
	}
}

// MovieStore хранит все фильмы
var MovieStore = make(map[string][]Movie)
