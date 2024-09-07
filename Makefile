# Установи переменные окружения
DB_URL=postgres://postgres:postgres@localhost:5432/yourdb?sslmode=disable
MIGRATIONS_DIR=./migrations

# Команда для создания новой миграции
create-migration:
	@docker run --rm -v $(PWD)/migrations:/migrations migrate/migrate create -ext sql -dir /migrations -seq $(name)

# Команда для запуска миграций
migrate-up:
	@docker run --rm -v $(PWD)/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database $(DB_URL) up

# Команда для отката миграций
migrate-down:
	@docker run --rm -v $(PWD)/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database $(DB_URL) down -all

# Команда для сброса базы данных
migrate-drop:
	@docker run --rm -v $(PWD)/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database $(DB_URL) drop

# Команда для принудительной установки версии миграции
migrate-force:
	@docker run --rm -v $(PWD)/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database $(DB_URL) force $(version)

# Команда для генерации моков с помощью Mockery через Docker
generate-mocks:
	@docker run --rm -v $(PWD):/src -w /src vektra/mockery --config .mockery.yaml

# Команда для сборки приложения
build:
	@go build -o bin/app ./cmd

# Команда для запуска приложения
run:
	@go run ./cmd

# Команда для локальной разработки (выполняет миграции и запускает сервер)
dev: migrate-up run

# Команда для запуска докер-контейнеров (если нужно использовать Docker)
docker-up:
	@docker-compose up -d

# Команда для остановки докер-контейнеров
docker-down:
	@docker-compose down

# Команда для обновления зависимостей
deps:
	@go mod tidy

# Команда для форматирования кода
fmt:
	@gofmt -s -w .

# Очистка скомпилированных файлов
clean:
	@rm -rf bin
