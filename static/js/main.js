// Основной JavaScript файл для сайта каталога фильмов

document.addEventListener('DOMContentLoaded', function() {
  // Загрузка данных о фильмах
  fetch(`/static/data/movies.json?t=${Date.now()}`)
      .then(response => response.json())
      .then(data => {
        // Отображение фильмов по категориям
        displayMoviesByCategory(data);
        // Инициализация всплывающих подсказок
        initializeTooltips();
        // Инициализация анимаций при прокрутке
        initializeScrollAnimations();
        // Проверка якоря в URL и скроллинг до категории
        const hash = window.location.hash.substring(1); // Убираем # из #comedy
        if (hash) {
          scrollToCategory(hash);
        }
        // Обработка кликов по ссылкам в хедере
        initializeHeaderLinks();

        // Инициализация модального окна
        const modal = document.getElementById("movieModal");
        if (!modal) {
          createMovieModal();
        } else {
          initializeMovieModal(modal);
        }
      })
      .catch(error => console.error('Ошибка загрузки данных о фильмах:', error));
});

// Функция для обработки кликов по ссылкам в хедере
function initializeHeaderLinks() {
  const navLinks = document.querySelectorAll('header nav ul li a');
  navLinks.forEach(link => {
    link.addEventListener('click', function(e) {
      const href = this.getAttribute('href');
      const [path, hash] = href.split('#'); // Разделяем путь и якорь
      const categoryKey = hash;

      // Если мы уже на странице /movies
      if (window.location.pathname === '/movies' && categoryKey) {
        e.preventDefault(); // Предотвращаем переход по ссылке
        scrollToCategory(categoryKey); // Скроллим к категории
      }
      // Если на другой странице, переход произойдёт как обычно
    });
  });
}

// Функция для скроллинга до категории
function scrollToCategory(categoryKey) {
  const categorySection = document.querySelector(`.category-section[data-category="${categoryKey}"]`);
  if (categorySection) {
    categorySection.scrollIntoView({ behavior: 'smooth', block: 'center' });
  } else {
    console.log(`Категория ${categoryKey} не найдена`);
  }
}

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
    categorySection.dataset.category = categoryKey; // Добавляем атрибут для поиска

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
    prevButton.innerHTML = '<';
    prevButton.addEventListener('click', () => {
      movieSlider.scrollBy({ left: -600, behavior: 'smooth' });
    });

    const nextButton = document.createElement('div');
    nextButton.className = 'slider-nav slider-nav-next';
    nextButton.innerHTML = '>';
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

  // Создаем оверлей для затемнения при наведении
  const overlay = document.createElement('div');
  overlay.className = 'movie-overlay';

  posterContainer.appendChild(posterImg);
  posterContainer.appendChild(overlay);

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

  movieItem.appendChild(movieCard);

  // Эффекты при наведении
  movieCard.addEventListener('mouseenter', function() {
    // Притемнение постера и увеличение
    overlay.classList.add('active');
    posterImg.style.transform = 'scale(1.1)';
    posterImg.style.transition = 'transform 0.3s ease';

    // Отображение информации о фильме
    infoContainer.classList.add('visible');
  });

  movieCard.addEventListener('mouseleave', function() {
    // Удаление притемнения
    overlay.classList.remove('active');
    posterImg.style.transform = 'scale(1)';

    // Скрытие информации
    infoContainer.classList.remove('visible');
  });

  // Открытие модального окна при клике
  movieCard.addEventListener('click', function(e) {
    e.stopPropagation(); // Предотвращаем всплытие события

    const modal = document.getElementById("movieModal");

    // Заполняем модальное окно информацией о фильме
    document.getElementById("modalTitle").textContent = movie.title;
    document.getElementById("modalYear").textContent = `Год выпуска: ${movie.year}`;
    document.getElementById("modalDescription").textContent = movie.fullDescription || movie.description || 'Подробное описание отсутствует';

    const modalImage = document.getElementById("modalImage");
    if (modalImage) {
      modalImage.src = movie.imagePath || '/static/images/placeholder.jpg';
      modalImage.alt = `Постер фильма "${movie.title}"`;
    }

    // Добавляем ссылку, если она есть
    const modalBody = modal.querySelector(".modal-body");
    let linkElement = modal.querySelector(".modal-link"); // Проверяем, есть ли уже ссылка
    if (movie.link) {
      if (!linkElement) {
        // Если элемента ссылки ещё нет, создаём его
        linkElement = document.createElement("a");
        linkElement.className = "modal-link";
        linkElement.target = "_blank"; // Открывать в новой вкладке
        linkElement.rel = "noopener noreferrer"; // Безопасность
        modalBody.appendChild(linkElement);
      }
      linkElement.href = movie.link;
      linkElement.textContent = "Смотреть на Кинопоиске"; // Или другой текст
    } else if (linkElement) {
      // Если ссылки нет в данных, удаляем элемент, если он был
      linkElement.remove();
    }

    // Позиционируем модальное окно рядом с курсором
    const modalWidth = modal.offsetWidth;
    const modalHeight = modal.offsetHeight;
    const viewportWidth = window.innerWidth;
    const viewportHeight = window.innerHeight;

    let left = e.clientX + 10; // Смещение вправо от курсора на 10px
    let top = e.clientY + window.scrollY + 10; // Смещение вниз от курсора на 10px с учетом прокрутки

    // Корректируем позицию, чтобы окно не выходило за пределы экрана
    if (left + modalWidth > viewportWidth - 10) {
      left = e.clientX - modalWidth - 10; // Показываем слева от курсора, если не помещается справа
    }
    if (top + modalHeight > window.scrollY + viewportHeight - 10) {
      top = e.clientY + window.scrollY - modalHeight - 10; // Показываем выше курсора, если не помещается снизу
    }
    if (left < 10) {
      left = 10; // Не даём выйти за левую границу
    }
    if (top < window.scrollY + 10) {
      top = window.scrollY + 10; // Не даём выйти за верхнюю границу видимой области
    }

    modal.style.left = `${left}px`;
    modal.style.top = `${top}px`;

    // Показываем модальное окно
    modal.classList.add("visible");
  });

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
      if (!tooltip) return;

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

// Функция для создания модального окна, если его нет на странице
function createMovieModal() {
  // Создаем структуру модального окна
  const modal = document.createElement("div");
  modal.id = "movieModal";
  modal.className = "movie-modal";

  const modalContent = document.createElement("div");
  modalContent.className = "modal-content";

  const closeBtn = document.createElement("span");
  closeBtn.className = "close";
  closeBtn.innerHTML = "×";

  const modalHeader = document.createElement("div");
  modalHeader.className = "modal-header";

  const modalTitle = document.createElement("h2");
  modalTitle.id = "modalTitle";

  const modalImageContainer = document.createElement("div");
  modalImageContainer.className = "modal-image-container";

  const modalImage = document.createElement("img");
  modalImage.id = "modalImage";
  modalImage.alt = "Постер фильма";

  const modalBody = document.createElement("div");
  modalBody.className = "modal-body";

  const modalYear = document.createElement("p");
  modalYear.id = "modalYear";
  modalYear.className = "modal-year";

  const modalDescription = document.createElement("p");
  modalDescription.id = "modalDescription";
  modalDescription.className = "modal-description";

  // Собираем модальное окно
  modalHeader.appendChild(modalTitle);
  modalImageContainer.appendChild(modalImage);
  modalBody.appendChild(modalYear);
  modalBody.appendChild(modalDescription);

  modalContent.appendChild(closeBtn);
  modalContent.appendChild(modalHeader);
  modalContent.appendChild(modalImageContainer);
  modalContent.appendChild(modalBody);

  modal.appendChild(modalContent);

  // Добавляем модальное окно в документ
  document.body.appendChild(modal);

  // Инициализируем обработчики событий сразу после создания
  initializeMovieModal(modal);
}

// Инициализация обработчиков событий для модального окна
function initializeMovieModal(modal) {
  const closeButton = modal.querySelector(".close");

  // Проверяем, что кнопка найдена
  if (!closeButton) {
    console.error("Кнопка закрытия (.close) не найдена в модальном окне!");
    return;
  }

  // Удаляем существующий обработчик, если он есть, чтобы избежать дублирования
  closeButton.removeEventListener("click", closeModalHandler);
  closeButton.addEventListener("click", closeModalHandler);

  // Обработчик для закрытия окна
  function closeModalHandler(e) {
    e.stopPropagation(); // Предотвращаем всплытие события
    console.log("Клик по кнопке закрытия сработал!"); // Отладочный вывод
    modal.classList.remove("visible");
  }

  // Закрытие при клике вне модального окна
  window.removeEventListener("click", outsideClickHandler);
  window.addEventListener("click", outsideClickHandler);
  function outsideClickHandler(event) {
    if (event.target === modal) {
      console.log("Клик вне модального окна сработал!");
      modal.classList.remove("visible");
    }
  }

  // Закрытие по клавише Escape
  document.removeEventListener("keydown", escapeKeyHandler);
  document.addEventListener("keydown", escapeKeyHandler);
  function escapeKeyHandler(event) {
    if (event.key === "Escape" && modal.classList.contains("visible")) {
      console.log("Escape нажата, закрываем модальное окно!");
      modal.classList.remove("visible");
    }
  }
}