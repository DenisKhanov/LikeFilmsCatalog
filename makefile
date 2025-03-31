# Переменные
SERVER_DIR = cmd/server
SERVER_BIN = server
GO = go
GOFLAGS = -v

# Цель по умолчанию
.PHONY: all
all: build

# Сборка обоих приложений
.PHONY: build
build: deps build-server


# Сборка сервера
.PHONY: build-server
build-server:
	$(GO) build $(GOFLAGS) -o $(SERVER_BIN) ./$(SERVER_DIR)

# Запуск сервера и бота
.PHONY: run
run: build
	@echo "Starting server..."
	@./$(SERVER_BIN) > server.log 2>&1 || (echo "Server failed to start. Check server.log for errors."; exit 1)

# Остановка запущенных процессов
.PHONY: stop
stop:
	-pkill -f $(SERVER_BIN)

# Очистка бинарных файлов
.PHONY: clean
clean:
	rm -f $(SERVER_BIN)
	rm -f server.log

# Установка зависимостей
.PHONY: deps
deps:
	$(GO) mod tidy

# Форматирование кода
.PHONY: fmt
fmt:
	$(GO) fmt ./...

# Тестирование
.PHONY: test
test:
	$(GO) test ./... -v