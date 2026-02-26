Быстрая установка
1. Создать .custom-gcl.yml
version: v2.10.1

plugins:
  - module: github.com/aeglukhov/loglinter
    path: .
2. Собрать кастомный golangci-lint
golangci-lint custom

Бинарник появится, например:

./custom-gcl
3. Подключить плагин в .golangci.yml
version: "2"

linters:
  default: none
  enable:
    - mylinter

  settings:
    custom:
      mylinter:
        type: module
4. Запуск
./custom-gcl run