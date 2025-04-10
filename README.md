# TODOCLI

Простое CLI-приложение для управления повседневными задачами на языке Go.

Позволяет добавлять, просматривать, отмечать выполненными и сохранять задачи в JSON-файл.

## Установка

1. Убедитесь, что у вас установлен Go 1.20 или выше.
2. Клонируйте репозиторий:

```bash
git clone https://github.com/mkbdlnv/todocli.git
cd todocli
```

## Запуск 

Из корневой папки выполните
```bash
go run cmd/main.go
```


## Доступные команды

Программа предлагает выбрать команду из меню:

1. **Create task** – добавить новую задачу  
2. **See upcoming tasks** – посмотреть список предстоящих задач  
3. **See finished tasks** - посмортерть выполненных задач
4. **Mark as done** - отметить задачу как выполненная
0. **Stop program** - завершить программу

## Структура проекта

- `cmd/main.go` — точка входа
- `internal/app/` — логика приложения и меню
- `internal/models/` — структура `Task`
- `internal/storage/` — интерфейс и реализация хранения задач
- `tasks.json` — файл, в который сохраняются задачи

## Тесты
Для запуска тестов можете использовать комманды:
```bash
go test ./... -v #для запуска всех тестов
go test ./... -coverprofile=coverage.out #для создания файла отчета и покрытием
go tool cover -html=coverage.out # для генерации html отчета
```


## License

MIT
