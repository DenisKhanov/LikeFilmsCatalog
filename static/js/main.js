// Основной JavaScript файл для сайта каталога фильмов

document.addEventListener('DOMContentLoaded', function() {
  // Загрузка данных о фильмах
  fetch('/static/data/movies.json')
    .then(response => response.json())
    .then(data => {
      // Отображение фильмов по категориям
      displayMoviesByCategory(data);
      // Инициализация всплывающих подсказок
      initializeTooltips();
      // Инициализация анимаций при прокрутке
      initializeScrollAnimations();
    })
    .catch(error => console.error('Ошибка загрузки данных о фильмах:', error));

  // Анимация хедера при прокрутке
  initializeHeaderAnimation();
});

// Функция для отображения фильмов по категориям
function displayMoviesByCategory(moviesByCategory) {
  const mainContainer = document.getElementById('movies-container');
  if (!mainContainer) return;

  // Очищаем контейнер
  mainContainer.innerHTML = '';

  // Отображаем фильмы по категориям
  Object.keys(moviesByCategory).forEach((categoryKey, index) => {
    const categoryName = getCategoryDisplayName(categoryKey);
    const movies = moviesByCategory[categoryKey];
    
    // Создаем секцию для категории
    const categorySection = document.createElement('section');
    categorySection.className = 'category-section mb-12';
    categorySection.style.animationDelay = `${0.1 * index}s`;
    
    // Добавляем заголовок категории
    const categoryTitle = document.createElement('h2');
    categoryTitle.className = 'category-title text-2xl font-bold mb-6';
    categoryTitle.textContent = categoryName;
    categorySection.appendChild(categoryTitle);
    
    // Создаем слайдер для фильмов
    const sliderContainer = document.createElement('div');
    sliderContainer.className = 'slider-container relative';
    
    const movieSlider = document.createElement('div');
    movieSlider.className = 'movie-slider';
    
    // Добавляем фильмы в слайдер
    movies.forEach(movie => {
      const movieCard = createMovieCard(movie);
      movieSlider.appendChild(movieCard);
    });
    
    // Добавляем кнопки навигации слайдера
    const prevButton = document.createElement('div');
    prevButton.className = 'slider-nav slider-nav-prev';
    prevButton.innerHTML = '&lt;';
    prevButton.addEventListener('click', () => {
      movieSlider.scrollBy({ left: -600, behavior: 'smooth' });
    });
    
    const nextButton = document.createElement('div');
    nextButton.className = 'slider-nav slider-nav-next';
    nextButton.innerHTML = '&gt;';
    nextButton.addEventListener('click', () => {
      movieSlider.scrollBy({ left: 600, behavior: 'smooth' });
    });
    
    sliderContainer.appendChild(movieSlider);
    sliderContainer.appendChild(prevButton);
    sliderContainer.appendChild(nextButton);
    
    categorySection.appendChild(sliderContainer);
    mainContainer.appendChild(categorySection);
  });
}

// Функция для создания карточки фильма
function createMovieCard(movie) {
  const movieItem = document.createElement('div');
  movieItem.className = 'movie-slider-item';
  
  const movieCard = document.createElement('div');
  movieCard.className = 'movie-card';
  movieCard.dataset.movieId = movie.id;
  
  // Постер фильма
  const posterContainer = document.createElement('div');
  posterContainer.className = 'movie-poster';
  
  const posterImg = document.createElement('img');
  posterImg.src = movie.imagePath || '/static/images/movies/placeholder.jpg';
  posterImg.alt = `Постер фильма "${movie.title}"`;
  posterImg.loading = 'lazy';
  
  posterContainer.appendChild(posterImg);
  
  // Информация о фильме
  const infoContainer = document.createElement('div');
  infoContainer.className = 'movie-info';
  
  const title = document.createElement('h3');
  title.className = 'movie-title';
  title.textContent = movie.title;
  
  const year = document.createElement('div');
  year.className = 'movie-year';
  year.textContent = movie.year;
  
  const description = document.createElement('div');
  description.className = 'movie-description';
  description.textContent = movie.description || 'Описание отсутствует';
  
  infoContainer.appendChild(title);
  infoContainer.appendChild(year);
  infoContainer.appendChild(description);
  
  // Всплывающая подсказка с полным описанием
  const tooltip = document.createElement('div');
  tooltip.className = 'movie-tooltip';
  
  const tooltipTitle = document.createElement('h4');
  tooltipTitle.className = 'tooltip-title';
  tooltipTitle.textContent = movie.title;
  
  const tooltipYear = document.createElement('div');
  tooltipYear.className = 'tooltip-year';
  tooltipYear.textContent = movie.year;
  
  const tooltipDescription = document.createElement('div');
  tooltipDescription.className = 'tooltip-description';
  tooltipDescription.textContent = movie.fullDescription || movie.description || 'Подробное описание отсутствует';
  
  tooltip.appendChild(tooltipTitle);
  tooltip.appendChild(tooltipYear);
  tooltip.appendChild(tooltipDescription);
  
  // Собираем карточку
  movieCard.appendChild(posterContainer);
  movieCard.appendChild(infoContainer);
  movieCard.appendChild(tooltip);
  
  movieItem.appendChild(movieCard);
  
  return movieItem;
}

// Функция для получения отображаемого имени категории
function getCategoryDisplayName(categoryKey) {
  const categoryNames = {
    'drama': 'Драма',
    'comedy': 'Комедия',
    'fantasy': 'Фантастика и фэнтези',
    'thriller': 'Триллер и детектив',
    'biography': 'Биографический',
    'historical': 'Исторический и военный',
    'melodrama': 'Мелодрама'
  };
  
  return categoryNames[categoryKey] || categoryKey;
}

// Инициализация всплывающих подсказок
function initializeTooltips() {
  // Позиционирование всплывающих подсказок
  document.querySelectorAll('.movie-card').forEach(card => {
    card.addEventListener('mouseenter', function() {
      const tooltip = this.querySelector('.movie-tooltip');
      
      // Получаем размеры и позицию карточки
      const cardRect = this.getBoundingClientRect();
      
      // Позиционируем подсказку над карточкой
      tooltip.style.bottom = '100%';
      tooltip.style.left = '50%';
      tooltip.style.transform = 'translateX(-50%)';
      
      // Проверяем, не выходит ли подсказка за пределы экрана
      const tooltipRect = tooltip.getBoundingClientRect();
      
      if (tooltipRect.left < 0) {
        tooltip.style.left = '0';
        tooltip.style.transform = 'none';
      }
      
      if (tooltipRect.right > window.innerWidth) {
        tooltip.style.left = 'auto';
        tooltip.style.right = '0';
        tooltip.style.transform = 'none';
      }
      
      if (tooltipRect.top < 0) {
        tooltip.style.bottom = 'auto';
        tooltip.style.top = '100%';
      }
    });
  });
}

// Инициализация анимаций при прокрутке
function initializeScrollAnimations() {
  // Анимация появления элементов при прокрутке
  const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.classList.add('visible');
        observer.unobserve(entry.target);
      }
    });
  }, { threshold: 0.1 });
  
  document.querySelectorAll('.category-section').forEach(section => {
    observer.observe(section);
  });
}

// Анимация хедера при прокрутке
function initializeHeaderAnimation() {
  const header = document.querySelector('.header');
  if (!header) return;
  
  window.addEventListener('scroll', () => {
    if (window.scrollY > 50) {
      header.classList.add('scrolled');
    } else {
      header.classList.remove('scrolled');
    }
  });
}
