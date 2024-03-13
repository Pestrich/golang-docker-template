# Golang Docker Template

Шаблон Golang приложения с использованием Docker Compose. Приложение собирается в два этапа, поэтому конечный образ имеет небольшой размер. Но из-за этого не работает `docker compose watch` для автоматической пересборки. Также включает в себя базу данных PostgreSQL

## Локальная разработка

&#x2705; Выполнить клонирование проекта\
&#x2705; Создать `.env` файл на основе [.env.example](./.env.example)

Затем выполнить сборку контейнеров

```bash
docker compose -f docker-compose.yml up -d --build
```

## Общий рабочий процесс

&#x2705; Пишем код\
&#x2705; Пишем тесты\
&#x2705; Проходим линтер\
&#x2705; Обновляем [openapi документацию](./openapi.json) (при необходимости)

## Команды

### Сборка контейнеров

```bash
docker compose -f docker-compose.yml up -d --build
```

### Запуск контейнеров

```bash
docker compose -f docker-compose.yml up -d
```
