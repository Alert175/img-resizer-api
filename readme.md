# API Для работы с Bitrix24 АН Doma

## Запуск в prod версии:
- `docker-compose up --build -d`

#

## Запуск в dev версии:

- Изменить конфигурацию .env для СУБД
- В файле docker-compose.yaml закоментировать сервис API и оставить конфиг для СУБД
- Выполнить команду - `docker-compose up --build`
- Выполнить компанды:
`go mod vendor`, `swag init --parseDependency --parseInternal -g main.go`