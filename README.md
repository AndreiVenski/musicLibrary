# Реализация онлайн библиотеки песен

## Запуск проекта
После скачивания откройте файл .env и введите значение для MUSICSERVICE_APIURL (строка подключения к api описанном свагером во втором пункте тз), после используйте команду :
`docker compose up`

## Структура проекта:

    api:
    └── docs - swagger
    cmd:
    └── main.go - точка входа в программу , обьявление самых основных структур
    config:
    └── config.go - чтение файла .env и запись переменных окружения в Config{}
    docker:
    └── Dockerfile
    internal:
    ├── libraryService:
    │   ├── delivery - реализация handlers и роутеров
    │   ├── repository - реализация libRepo для работы с db
    │   ├── infrastucture - реализация клиента http и подключение к стороннему апи
    │   └── usecase - реализация libUseCase (бизнес правила приложения)
    ├── models - модели используемые в internal (некоторые используются только для документации swagger)
    └── server - настройка а также запуск сервера и инициализация всех частей libraryService
    migrations - миграция базы данных перед запуском сервера, запускается в cmd/main.go
    pkg:
    ├── db - реализация подключения к базе данных
    ├── httpErrors - юзерские ошибки
    ├── logger - базовая реализация логгера
    └── utils - содержит работу с контекстом запроса, валидацию структур и базовую работу с ошибками
    .env - файл конфигурации , нужно заполнить строку для стороннего апи
    docker-compose.yml - запуск контейнера с postgres и самим сервером

## Выполнение условий тз :
- Реализованы 7 rest методов :
  - Get library information (получение песен из базы данных с фильтрацией по данным полученным от пользователя)
  - Add a new song (добавление песни в базу данных и получение дополнительных данных от сторонего API)
  - Delete a song (удаление песни по названию группы и песни)
  - Delete a song by ID (удаление песни по id песни)
  - Get music text information (получение определенного куплета из песни)
  - Update music information (изменение деталей песни по названию группы и песни)
  - Update music information by ID (изменение всех деталей включая назваение группы и песни по ID)
- Миграция базы данных при старте сервера (файл миграции находится в папке migrations)
- Код покрыт логами 
- Сгенирирован swagger и доступен по пути (host:port)/swagger/

Для удобства тестирования при первой миграции базы данных добавляются две записи песен :
- ('testgroup1', 'testsong', '2000', 'test verdes\n\n test verdes2 \n\n test verdes3', 'http://example.com/'),
- ('testgroup2', 'testsong', '2000', 'some text', 'http://example.com/')

Контакты:
- venskiandrei32@gmail.com
- telegram: @ban_ka