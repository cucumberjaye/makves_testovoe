# makves_testovoe

Сервис для получения данных по id из файла ueba.csv.

Сервис принимает Get запрос, путь /get-items/{id} для одного id и путь /get-items/{start_id}-{end_id} для получения массива данных. Данные приходят в формате json.

Для запуска сервера прописать:
```
make run
```

Для запуска тестов прописать:
```
make tests
```
