# URL Shortener API

## Описание
Этот проект представляет собой REST API для сокращения URL-адресов. Пользователь может отправить длинный URL и получить короткий алиас, по которому можно перейти на исходный адрес.

## Стек технологий
- Golang
- PostgreSQL
- Chi Router
- slog (Structured Logging)
- GitHub Actions (для CI/CD)

## Установка и запуск

### Локальный запуск
1. Установите зависимости:
   ```sh
   go mod tidy
   ```
2. Заполните файл `config/local.yaml` с параметрами для базы данных и сервера.
   ```sh
    env: "local"
    database:
       host: "localhost"
       port: 5432
       user: "postgres"
       password: "your_password"
       dbname: "your_db_name"
       sslmode: "disable"
    
    http_server:
       address: "localhost:8082"
       timeout: 4s
       idle_timeout: 60s
       user: "your_username"
       password: "your_password"
   ```
4. Запустите сервер:
   ```sh
   go run cmd/url-shortener/main.go
   ```

### Docker (опционально)
1. Соберите и запустите контейнеры:
   ```sh
   docker-compose up --build
   ```

## API эндпоинты

### Сохранение URL
- **POST** `/url`
- **Headers:**
  ```json
  {
    "Authorization": "Basic base64(user:password)"
  }
  ```
- **Body:**
  ```json
  {
    "url": "https://example.com",
    "alias": "myalias"
  }
  ```
- **Response (200 OK):**
  ```json
  {
    "message": "URL saved successfully",
    "alias": "myalias"
  }
  ```

### Перенаправление
- **GET** `/{alias}`
- **Response (302 Found):** Перенаправление на исходный URL.
- **Response (404 Not Found):** Если alias не найден.

## Тестирование
Запуск всех тестов:
```sh
go test ./...
```

## CI/CD (GitHub Actions)
Проект автоматически деплоится через GitHub Actions. При каждом push в `main` происходят:
- Сборка и тестирование проекта
- Деплой на сервер

## Автор
Veylin2175
