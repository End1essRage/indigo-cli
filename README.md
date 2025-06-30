### Быстрый старт

#### Требования ПО:
1) go v1.24.x
2) docker

#### Установка и создание проекта:
1) `go install github.com/end1essrage/indigo-cli@latest`
2) `indigo-cli` в папке с проектом

#### Запуск сервисов
```bash
# Linux/macOS/Windows WSL
BOT_TOKEN="<ваш_токен>" docker-compose up -d --build

# Windows CMD
cmd /C "set BOT_TOKEN=<ваш_токен> && docker-compose up -d --build"

# Windows PowerShell
$env:BOT_TOKEN="<ваш_токен>"; docker-compose up -d --build
```
#### Запуск сервисов
`docker-compose down`

#### Полезные ссылки:
1) "Дока"



