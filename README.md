# TaskAPI

TaskAPI — RESTful API для управления статусами задач на Go с использованием чистой архитектуры и Docker.

Формулировку задания можно найти в папке `doc`.
```
├── doc/                   # Задание 
The_Task.jpeg 
```
## 📦 Архитектура проекта

```
├── cmd/                    # Точка входа
├── internal/
│   ├── entities/           # Сущности
│   ├── logger/             # Логирование
│   ├── scr/tasks/          # CRUD и обработчики задач
│   ├── services/           # Сервисы
│   └── taskWatcher.go                 # Прочие внутренние пакеты
├── logs/                   # Логи приложения
├── scripts/                # Скрипты запуска
├── Dockerfile              # Docker-образ
├── docker-compose.yml      # Docker Compose
├── go.mod, go.sum          # Зависимости Go
└── README.md
```

## ⚙️ Требования

- Go 1.23+
- Docker и Docker Compose

## 🚀 Быстрый старт
Необходимо находиться в корне   проекта 

### Через Docker

```bash
docker-compose up --build
```

### Через скрипт

```bash
./scripts/run.sh
```

## 🌐 API Эндпоинты

#### id — строго целое численное значение

| Метод | Путь         | Описание                |
|-------|--------------|-------------------------|
| GET   | /tasks       | Получить все задачи     |
| GET   | /tasks/{id}  | Получить задачу по id   |
| POST  | /tasks/{id}       | Создать новую задачу по id    |
| DELETE| /tasks/{id}  | Удалить задачу по id          |

### Примеры запросов

```http
curl -X GET http://localhost:8080/tasks

curl -X GET http://localhost:8080/tasks/{1}

curl -X POST http://localhost:8080/tasks/{1}

curl -X DELETE http://localhost:8080/tasks/{1}
```

### Пример ответа на Get запрос
```json
{
    "TaskInfo":{"id":"1"},
    "WorkStatus":"STARTED",
    "CreatedAt":"2025-06-24T11:05:12Z","Duration":"5.830005803s",
    "Completed":false
    }
```

## ⏱️ Сервис Watcher

Сервис **Watcher** отвечает за отслеживание и автоматическое обновление статусов задач в памяти приложения (их длительности). Он работает в фоне, периодически проверяя длительность выполнения задач, чтобы поддерживать актуальность данных без участия пользователя.

```
internal/
└── services/
    └── taskCacheWatcher.go
```

## ⚙️ Переменные окружения

- `TASKAPI_LISTEN_PORT` — порт для запуска API (по умолчанию :8080)
- `SYNC_TASKSTATUS_INTERVAL_SEC` — интервал синхронизации времени выполнения задачи через (по умолчанию 3s). 

```
internal/
└── utils/
    └── envs.go
```
## 📝 Логирование

- Логи пишутся в файл `logs/app.log` и выводятся в консоль.
- Поддерживаются уровни: TRACE, DEBUG, INFO, WARN, ERROR.

## 🐳 Работа с Docker
Возможно команды не будут работать без прав суперпользователя: `sudo`
- Сборка и запуск: `docker-compose up --build`
- Остановка: `docker-compose down`
- Очистка неиспользуемых образов: `docker image prune`

## 👤 Автор

- [Pavel-K16 @KPAudz -- tg](https://github.com/Pavel-K16)
