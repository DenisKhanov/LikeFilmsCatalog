// Файл для реализации функциональности всплывающих описаний при наведении
document.addEventListener('DOMContentLoaded', function() {
  // Инициализация всплывающих подсказок при наведении
  initializeHoverDescriptions();
});

// Функция для инициализации всплывающих описаний при наведении
function initializeHoverDescriptions() {
  // Обработка событий наведения на карточки фильмов
  document.addEventListener('mouseover', function(event) {
    // Находим ближайшую карточку фильма
    const movieCard = event.target.closest('.movie-card');
    if (!movieCard) return;
    
    // Показываем всплывающее описание
    const tooltip = movieCard.querySelector('.movie-tooltip');
    if (tooltip) {
      // Позиционируем подсказку относительно курсора
      positionTooltip(tooltip, event);
      
      // Делаем подсказку видимой
      tooltip.style.opacity = '1';
      tooltip.style.visibility = 'visible';
    }
  }, true);
  
  // Обработка событий ухода мыши с карточек фильмов
  document.addEventListener('mouseout', function(event) {
    // Находим ближайшую карточку фильма
    const movieCard = event.target.closest('.movie-card');
    if (!movieCard) return;
    
    // Если мышь уходит с карточки фильма, скрываем подсказку
    if (!movieCard.contains(event.relatedTarget)) {
      const tooltip = movieCard.querySelector('.movie-tooltip');
      if (tooltip) {
        tooltip.style.opacity = '0';
        tooltip.style.visibility = 'hidden';
      }
    }
  }, true);
  
  // Обработка движения мыши для динамического позиционирования подсказки
  document.addEventListener('mousemove', function(event) {
    const movieCard = event.target.closest('.movie-card');
    if (!movieCard) return;
    
    const tooltip = movieCard.querySelector('.movie-tooltip');
    if (tooltip && tooltip.style.visibility === 'visible') {
      positionTooltip(tooltip, event);
    }
  }, true);
}

// Функция для позиционирования всплывающей подсказки
function positionTooltip(tooltip, event) {
  // Получаем размеры и позицию карточки
  const card = event.target.closest('.movie-card');
  const cardRect = card.getBoundingClientRect();
  
  // Получаем размеры подсказки
  const tooltipRect = tooltip.getBoundingClientRect();
  
  // Определяем позицию подсказки
  let top, left;
  
  // Позиционируем подсказку над карточкой
  top = cardRect.top - tooltipRect.height - 10;
  left = cardRect.left + (cardRect.width / 2) - (tooltipRect.width / 2);
  
  // Проверяем, не выходит ли подсказка за пределы экрана
  if (top < 0) {
    // Если подсказка выходит за верхний край экрана, показываем её под карточкой
    top = cardRect.bottom + 10;
  }
  
  if (left < 0) {
    left = 0;
  } else if (left + tooltipRect.width > window.innerWidth) {
    left = window.innerWidth - tooltipRect.width;
  }
  
  // Применяем позицию к подсказке
  tooltip.style.top = `${top}px`;
  tooltip.style.left = `${left}px`;
}

// Функция для создания эффекта размытия фона при наведении
function createBlurEffect() {
  document.querySelectorAll('.movie-card').forEach(card => {
    card.addEventListener('mouseenter', function() {
      // Добавляем класс для размытия всех остальных карточек
      document.querySelectorAll('.movie-card').forEach(otherCard => {
        if (otherCard !== card) {
          otherCard.classList.add('blur-effect');
        }
      });
    });
    
    card.addEventListener('mouseleave', function() {
      // Убираем размытие со всех карточек
      document.querySelectorAll('.movie-card').forEach(otherCard => {
        otherCard.classList.remove('blur-effect');
      });
    });
  });
}

// Функция для добавления эффекта параллакса при наведении
function addParallaxEffect() {
  document.querySelectorAll('.movie-poster').forEach(poster => {
    poster.addEventListener('mousemove', function(e) {
      const card = e.currentTarget.closest('.movie-card');
      const rect = card.getBoundingClientRect();
      
      // Вычисляем положение курсора относительно карточки
      const x = e.clientX - rect.left;
      const y = e.clientY - rect.top;
      
      // Вычисляем смещение в процентах
      const xPercent = (x / rect.width) * 100;
      const yPercent = (y / rect.height) * 100;
      
      // Применяем эффект параллакса к изображению
      const img = poster.querySelector('img');
      if (img) {
        img.style.transform = `translate(${(xPercent - 50) * 0.05}px, ${(yPercent - 50) * 0.05}px) scale(1.1)`;
      }
    });
    
    poster.addEventListener('mouseleave', function(e) {
      const img = e.currentTarget.querySelector('img');
      if (img) {
        img.style.transform = 'scale(1)';
      }
    });
  });
}

// Инициализация дополнительных эффектов при загрузке страницы
document.addEventListener('DOMContentLoaded', function() {
  // Добавляем эффект размытия
  createBlurEffect();
  
  // Добавляем эффект параллакса
  addParallaxEffect();
});
