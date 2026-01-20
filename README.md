# Решение тестового задания
## Запуск
git clone https://github.com/tearannosaur/Test-Task

Далее необходимо перейти в папку решения с помощью команды cd Test-Task

После перехода, нужно ввести команду в терминале docker-compose up --build 


**Эндпоинт POST /numbers**

Для проверки можно использовать как POSTMAN так и команды в терминале

Пример команды: Invoke-RestMethod -Uri http://localhost:8080/numbers `  -Method POST `  -Headers @{ "Content-Type" = "application/json" } `  -Body '{"number":3}'

Запрос
```json
{
  "number": 3
}
```

Ответ
```json

"success": [3]

```
Ошибки:
400 incorrect data
500 internal server error
