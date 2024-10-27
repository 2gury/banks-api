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

---

Ручка получения списка банков
localhost:8080/v1/banks/get

Request:
```
{
    "language": "kz"
}
```

Response:
```
{
    "banks": [
        {
            "id": "1",
            "external_id": "4612",
            "external_legacy_id": "6726",
            "name": "JetCar KZ - CPS",
            "logo": "https://my.s3-cdn.com/offers/thumbs/5/103198.png",
            "url": "https://jetcar.kz/cash/credit",
            "createdAt": "2024-10-18T18:25:07.264364Z",
            "updatedAt": "2024-10-18T18:25:07.264364Z",
            "period_from": "3",
            "period_to": "84",
            "amount_from": "300000",
            "amount_to": "15000000",
            "rate_from": 3.2,
            "rate_to": 3.7,
            "review_time": "1",
            "registration": [
                "Гражданство Казахстана"
            ],
            "credit_purpose": [
                "Потребительские нужды"
            ],
            "documents": [
                "ИИН",
                "Мобильный телефон",
                "VIN",
                "Гос. номер автомобиля",
                "Номер техпаспорта"
            ],
            "obtain_method": [
                "Банковская карта VISA",
                "Банковская карта MasterCard"
            ],
            "description": ""
        },
        {
            "id": "2",
            "external_id": "4586",
            "external_legacy_id": "6700",
            "name": "KMF KZ - CPS",
            "logo": "https://my.s3-cdn.com/offers/thumbs/8/103166.png",
            "url": "https://online.kmf.kz/app/pre-partner/?partner=kmf",
            "createdAt": "2024-10-18T18:25:07.281991Z",
            "updatedAt": "2024-10-18T18:25:07.281991Z",
            "period_from": "3",
            "period_to": "36",
            "amount_from": "0",
            "amount_to": "3000000",
            "rate_from": 27,
            "rate_to": 0,
            "review_time": "1",
            "registration": [
                "Гражданство Казахстана"
            ],
            "credit_purpose": [
                "Потребительские нужды"
            ],
            "documents": [
                "Выписка со счёта",
                "Заработная плата",
                "ИИН",
                "Мобильный телефон"
            ],
            "obtain_method": [
                "Онлайн выдача"
            ],
            "description": ""
        },
        {
            "id": "3",
            "external_id": "4565",
            "external_legacy_id": "6679",
            "name": "Swiss Capital KZ - CPS",
            "logo": "https://my.s3-cdn.com/offers/thumbs/6/103148.png",
            "url": "https://swisscapital.kz/",
            "createdAt": "2024-10-18T18:25:07.297493Z",
            "updatedAt": "2024-10-18T18:25:07.297493Z",
            "period_from": "0",
            "period_to": "60",
            "amount_from": "0",
            "amount_to": "20000000",
            "rate_from": 0.1,
            "rate_to": 38,
            "review_time": "0",
            "registration": [
                "Гражданство Казахстана"
            ],
            "credit_purpose": [
                "Финансирование для бизнеса и потребительских целей"
            ],
            "documents": [
                "Удостоверение личности или паспорт",
                "Свидетельство о регистрации ТС (техпаспорт)",
                "Нотариальное согласие от супруга(и) (если в браке)"
            ],
            "obtain_method": [
                "Оформление в филиале",
                "Выдача на банковскую карту"
            ],
            "description": ""
        }
    ]
}
```

---

Ручка создание нового банка
POST localhost:8080/v1/banks/update

Request:
```
{
    "bank": {
        "external_id": "4612",
        "external_legacy_id": "6726",
        "name": "JetCarz KZ - CPS",
        "logo": "https://my.s3-cdn.com/offers/thumbs/5/103198.png",
        "url": "https://jetcar.kz/cash/credit",
        "createdAt": "2024-10-18T18:25:07.264364Z",
        "updatedAt": "2024-10-18T18:25:07.264364Z",
        "period_from": "3",
        "period_to": "84",
        "amount_from": "300000",
        "amount_to": "15000000",
        "rate_from": 3.2,
        "rate_to": 3.7,
        "review_time": "1",
        "registration": [
            "Гражданство Казахстана"
        ],
        "credit_purpose": [
            "Бащкадры"
        ],
        "documents": [
            "ИИН",
            "Мобильный телефон",
            "VIN",
            "Гос. номер автомобиля",
            "Номер техпаспорта"
        ],
        "obtain_method": [
            "Банковская карта VISA",
            "Банковская карта MasterCard"
        ],
        "description": ""
    }
}
```

Response:
```
{}
```

---

Ручка изменения существующего банка
POST localhost:8080/v1/banks/update

Request:
```
{
    "bank": {
        "id": "1",
        "external_id": "4612",
        "external_legacy_id": "6726",
        "name": "JetCarz KZ - CPS",
        "logo": "https://my.s3-cdn.com/offers/thumbs/5/103198.png",
        "url": "https://jetcar.kz/cash/credit",
        "createdAt": "2024-10-18T18:25:07.264364Z",
        "updatedAt": "2024-10-18T18:25:07.264364Z",
        "period_from": "3",
        "period_to": "84",
        "amount_from": "300000",
        "amount_to": "15000000",
        "rate_from": 3.2,
        "rate_to": 3.7,
        "review_time": "1",
        "registration": [
            "Гражданство Казахстана"
        ],
        "credit_purpose": [
            "Бащкадры"
        ],
        "documents": [
            "ИИН",
            "Мобильный телефон",
            "VIN",
            "Гос. номер автомобиля",
            "Номер техпаспорта"
        ],
        "obtain_method": [
            "Банковская карта VISA",
            "Банковская карта MasterCard"
        ],
        "description": ""
    }
}
```

Response:
```
{}
```

---

Ручка удаления банка по id
POST localhost:8080/v1/banks/delete

Request:
```
{
    "id": 23
}
```

Response:
```
{}
```

---


Ручка получения возможных банковский предложений
POST localhost:8080/v1/banks/get-possible

Request:
```
{}
```

Response:
```
{
    "banks": [
        {
            "external_id": "476",
            "external_legacy_id": "873",
            "name": "BANNN - CPS",
            "logo": "https://my.s3-cdn.com/offers/thumbs/0/103226.png"
        },
        {
            "external_id": "848",
            "external_legacy_id": "274",
            "name": "TEST BSNK - CPS",
            "logo": "https://my.s3-cdn.com/offers/thumbs/5/103198.png"
        }
    ]
}
```

---

Ручка запроса инфы по банку (тяжелая ручка)
POST localhost:8080/v1/banks/request-bank-information

Request:
```
{
    "external_id": 4658
}
```

Response:
```
{
    "bank": {
        "id": "0",
        "external_id": "4658",
        "external_legacy_id": "6772",
        "name": "Ecommoney KZ - CPS",
        "logo": "https://my.s3-cdn.com/offers/thumbs/0/103226.png",
        "url": "https://ecommoney.kz/",
        "createdAt": "2024-10-19T09:57:18.769071Z",
        "updatedAt": "2024-10-19T09:57:18.769071Z",
        "period_from": "30",
        "period_to": "365",
        "amount_from": "100000",
        "amount_to": "500000",
        "rate_from": 38.3,
        "rate_to": 46,
        "review_time": "0",
        "registration": [],
        "credit_purpose": [],
        "documents": [],
        "obtain_method": [
            "Микрокредит можно оформить, не выходя из дома или работы, прямо на номер счета (IBAN) карты или номер текущего счета (IBAN) в банке;",
            "прозрачные условия получения кредита;",
            "заем быстро, без залога, без длительных процедур оформления и проверок."
        ],
        "description": "Финансовая платформа, предоставляющая онлайн-кредиты для клиентов в Казахстане. Процесс оформления микрокредита прост и интуитивно понятен, а средства поступают на счёт быстро и без лишних задержек."
    }
}
```
---

Ручка получения списка отзывов
POST localhost:8080/v1/reviews/get-list

Request:
```
{}
```

Response:
```
{
    "reviews": [
        {
            "id": "1",
            "content": "Понравился процесс оформления кредита",
            "is_approved": false,
            "user_email": "test@mail.ru",
            "user_phone": "+893455474543",
            "rating": "4",
            "bank_id": "12",
            "user_name": "Бауржан",
            "bank": "JetCarz KZ - CPS",
            "date": "2024-10-22T16:13:33.676555Z"
        },
        {
            "id": "12",
            "content": "Все понравилось",
            "is_approved": true,
            "user_email": "testing@mail.ru",
            "user_phone": "+892325474543",
            "rating": "1",
            "bank_id": "23",
            "user_name": "Бауржан",
            "bank": "JetCarz KZ - CPS",
            "date": "2024-10-22T16:13:33.676555Z"
        }
    ]
}
```

---

Ручка создания отзыва
POST localhost:8080/v1/reviews/create

Request:
```
{
    "content": "Все понравилось",
    "user_email": "testing@mail.ru",
    "user_phone": "+892325474543",
    "rating": "1",
    "bank_id": "23",
    "user_name": "Бауржан"
}
```

Response:
```
{}
```

---

Ручка аппрува/деклайна отзыва
POST localhost:8080/v1/reviews/update

Request:
```
{
    "id": "12",
    "is_approved": true
}
```

Response:
```
{}
```

---

Ручка включения/выключения автомодерации
POST localhost:8080/v1/automoderation/update

Request:
```
{
    "automoderation_enable": true
}
```

Response:
```
{}
```

---


