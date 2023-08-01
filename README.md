# GoAPIUserTransactions

This project on start initialises PostgreSQL database and Redis, reads .env file to establish chosen connection and creates 2 users with random positive balace and starts API server.
Данный проект при старте создает контейнеры с PostgreSQL и Redis, читает .env файл для установки выбранного соединения, в миграции создает 2х пользователей  со случайным положительным балансом и стартует API сервер.

# To start the app local using following commands:
1) to build the app use "docker compose build"
2) to start the app use "docker compose up"
3) to finish the app use "docker compose down"

# Для запуска локально используйте следующие команды:
1) для сборки образов "docker compose build"
2) для запуска приложения "docker compose up"
3) для завершения "docker compose down"

# Request examples/Примеры запросов

1) To get users with balance/для получения пользователей и их баланса

curl --location 'http://localhost:8000/users'

2) To make a transaction between users 

curl --location 'http://localhost:8000/transfer' \
--form 'UserFrom="1"' \
--form 'UserTo="2"' \
--form 'Amount="10"'
