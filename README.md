# RSOI_Microservices

## Что это?
Выполнение лабораторных работ по Распределённым системам обработки информации.
1. не в этом репозитории
2. Написать приложение на микросервисной архитектуре, архитектура описана в лабораторной. 
  * Должен присутствоват агрегационный сервис, через который должны проходить все запросы к системе.
  * Должно быть два запроса, требующих изменения данных на нескольких сервисах для их выполнения.
  * Должен быть запрос, требующий сбора данных с нескольких сервисов.
  * Всё это без авторизации (вкостылена базовая авторизация с агрегационного сервиса)
  * Юнит тестирование
3. Реализовать распределённый транзакции. 
  * Откат при ошибке одного из серсисов.
  * Реализовать отложенное выполнение при падении какого-либо сервиса.
4. Фронтенд
5. Авторизация
  * Реализоват codegrant
  * Реализовать авторизацию с ~~бледжеком и шлюхами~~ access и refresh токенами
  * Спрятать всё это за gateway, пока обычная реализация RFC6749
6. Статистика
  * Сообщения посылаются на сервер статистики асинхронно
  * Реализуется гарантированная доставка и обработка сообщений (с таймаутами ретраями и прочим)
  * Доступ к статистике давать только админам.

## Известные проблемы
Я знаю ряд проблем у этого проекта, но пока не дошли руки исправить.
1. Косяки с правописанием
2. Немного копипасты. (Немного повоевал с ней, но чуть-чуть осталось)
3. Архитектурные косяки (типа auth за gateway и валидации на auth). Некоторые из них объясняются постановкой задачи. (задачи приведены не полностью)
4. Безопасность. Всё в конфигах и makefile, знаю, что это не безопасно и т.д. делается не для продакшана, но жизнь упрощает.
5. Конфиги грузятся немного криво, переделывать начал, но не закончил.
6. Тесты, сделаны в рамках второй лабораторной работы и забыты

## Что в проекте
В этой ветке небольшие скрипты для сборки и запуска проекта. 
```
.
├── logs
├── pid
├── security
│   ├── auth
│   ├── event
│   ├── registration
│   ├── stats
│   └── user
└── src
    └── github.com
        └── wolf1996
            ├── auth
            ├── events
            ├── frontend
            ├── gateway
            ├── master
            ├── registration
            ├── stats
            └── user

```
структура проекта примерно такова, в корень дерева ставится `$GOPATH`. В папки грузятся ветки в соответствии с названиями.
Далее всё достаточно просто. 
Собрать всё 
```
make build_all
```
Запустить всё
```
make start_all
```
Остановить всё
```
make stop_all
```
Все логи пишутся в папку logs. В security хранятся сертификаты для ssl. 

Все зависимости у сервисов менджерятся dep-ом.