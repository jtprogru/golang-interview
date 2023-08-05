# golang-interview

[![Go](https://github.com/jtprogru/golang-interview/actions/workflows/go.yml/badge.svg)](https://github.com/jtprogru/golang-interview/actions/workflows/go.yml)

Задачи с собеседований, [LeetCode](https://leetcode.com) и [CodeWars](https://www.codewars.com) на позицию Golang Developer. Пока претендую на уровень Junior, но буду решать и другие задачки.

Задачи берутся с LeetCode и CodeWars.

## How to use

Создание нового модуля с задачей:

```shell
task newtask

New task module available in:
internal/task0032/
```

Внесение правок в модуль:

```shell
vim internal/task0032/*.go
```

Для проверки своей гипотезы решения выполни:

```shell
# Основной вариант прогона тестов
task testverb

# Вариант напрямую через go test
go test -v ./internal/task0032/*.go

# И вот такой вариант
go test github.com/jtprogru/golang-interview/internal/task0032 -v
```

Замените `task0032` на имя своего модуля.

## Profiles

- [CodeWars](https://www.codewars.com/users/jtprogru)
- [LeetCode](https://leetcode.com/jtprogru/)

