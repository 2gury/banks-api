# banks-api

Ручка получение перевода из гугла
POST localhost:8080/v1/translate/request-translation-text
Request:
```
{
    "language": "kz",
    "text": "привет!"
}
```

Response:
```
{
    "language": "kz",
    "text": "Сәлем!"
}
```
---

Ручка получения перевода из базы
POST localhost:8080/v1/translate/get-translation-text
Request:
```
{
    "language": "kz",
    "text": "Прове"
}
```

Response:
```
{
    "translation": {
        "id": "3",
        "lexeme": "Прове",
        "translated_lexeme": "Бащкадры",
        "source_language": "ru",
        "target_language": "kk"
    }
}
```
---

Ручка создания нового перевода
POST localhost:8080/v1/translate/create
Request:
```
{
    "translation": {
        "lexeme": "Провеаааа",
        "translated_lexeme": "Бащкадры",
        "source_language": "ru",
        "target_language": "kk"
    }
}
```

Response:
```
{
    "translation": {
        "id": "5",
        "lexeme": "Провеаааа",
        "translated_lexeme": "Бащкадры",
        "source_language": "ru",
        "target_language": "kk"
    }
}
```

---

Ручка изменения текущего перевода
POST localhost:8080/v1/translate/create
Request:
```
{
    "translation": {
        "id": "5",
        "lexeme": "Провеааа",
        "translated_lexeme": "Бащкадры",
        "source_language": "ru",
        "target_language": "kk"
    }
}
```

Response:
```
{
    "translation": {
        "id": "5",
        "lexeme": "Провеааа",
        "translated_lexeme": "Бащкадры",
        "source_language": "ru",
        "target_language": "kk"
    }
}
```

---

Ручка получения переводов
POST localhost:8080/v1/translate/get-list
Request:
```
{}
```

Response:
```
{
    "translations": [
        {
            "id": "1",
            "lexeme": "Проверddddка",
            "translated_lexeme": "Бащкадры",
            "source_language": "ru",
            "target_language": "kk"
        },
        {
            "id": "4",
            "lexeme": "Проверddddка",
            "translated_lexeme": "Бащкадры",
            "source_language": "ru",
            "target_language": "kk"
        },
        {
            "id": "3",
            "lexeme": "Прове",
            "translated_lexeme": "Бащкадры",
            "source_language": "ru",
            "target_language": "kk"
        },
        {
            "id": "2",
            "lexeme": "Потребительские нужды",
            "translated_lexeme": "Бащкадры",
            "source_language": "ru",
            "target_language": "kk"
        },
        {
            "id": "5",
            "lexeme": "Провеааа",
            "translated_lexeme": "Бащкадры",
            "source_language": "ru",
            "target_language": "kk"
        }
    ]
}
```


