# Library Management System

### Конфигурация базы данных

Для работы с PostgreSQL вам потребуется создать базу данных и настроить параметры подключения. Ниже приведена инструкция:

1. **Создание базы данных**: Вы можете создать базу данных с помощью следующей команды SQL:
    ```sql
    CREATE DATABASE library;
    ```
2. **Настройка подключения**: Укажите параметры подключения к вашей базе данных в файле `db.go`. Пример строки подключения:
    ```go
    dsn := "host=localhost user=youruser dbname=library password=yourpassword port=5432 sslmode=disable TimeZone=Asia/Shanghai"
    ```
    Измените параметры `youruser`, `yourpassword`, и другие, соответственно вашим настройкам.

3. **Миграция таблиц**: GORM автоматически создаст необходимые таблицы при первом запуске приложения благодаря функции `AutoMigrate`. Все структуры данных будут автоматически синхронизированы с вашей базой данных.

### Зависимости

Этот проект зависит от следующих библиотек:
- **GORM**: ORM библиотека для работы с базой данных.
- **Gorilla Mux**: Фреймворк для маршрутизации в Go.

Для установки необходимых библиотек выполните следующие команды:
```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/gorilla/mux


### Основные функции

- **Добавление новых книг**: Укажите название, автора, описание и жанр книги для добавления в библиотеку.
- **Просмотр списка книг**: Получите список всех книг, доступных в библиотеке, с названием и автором.
- **Поиск книг**: Поиск книг по названию или автору.
- **Удаление книг**: Удалите книгу из библиотеки по её уникальному идентификатору.
- **Работа с жанрами**: Добавление новых жанров и получение списка всех доступных жанров.

## Начало работы

### Предварительные требования

Убедитесь, что у вас установлен Go версии 1.16 или выше. Это можно проверить, выполнив команду:

```sh
go version
```

### Установка

Склонируйте репозиторий проекта:

```sh
git clone https://github.com/<ваш_логин>/library-management-system.git
cd library-management-system
```
Затем загрузите и установите необходимые зависимости:
```sh
go mod tidy
```
### Запуск проекта

Для запуска сервера выполните следующую команду в корневом каталоге проекта:

```sh
go run .
```

Сервер запустится на `localhost:8080`.



## Использование

### Добавление книги

```sh
curl -X POST -H "Content-Type: application/json" -d '{"title":"Название книги", "author":"Автор книги", "description":"Описание книги", "genre":"Жанр книги"}' http://localhost:8080/book
```

### Получение списка всех книг

```sh
curl http://localhost:8080/books
```

### Получение информации о книге по ID

```sh
curl http://localhost:8080/book/{id}
```

### Удаление книги по ID

```sh
curl -X DELETE http://localhost:8080/book/{id}
```

### Получение списка жанров

```sh
curl http://localhost:8080/genres
```

### Получение списка книг по жанру

```sh
curl http://localhost:8080/books/genre/{жанр}
```

### Поиск книг по названию или автору

Для выполнения поиска книг по ключевому слову или фразе в названии или имени автора, используйте следующий запрос:

```sh
curl "http://localhost:8080/books/search?q=ваш_поисковый_запрос"
```

Замените `ваш_поисковый_запрос` на интересующее вас ключевое слово или фразу:

```sh
curl "http://localhost:8080/books/search?q=python для чайников"
```
