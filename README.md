#vk-mini-app-backend

###API

### User routes:
**`POST`**
+ */user/register*
```json5
{
    "id": 123123,
    "name": "Name",
    "lastname": "Lastname",
    "city": "",
    "country": "",
    "sex": 2,
    "timezone": 10,
    "photo_100": "test100.com",
    "photo_200": "test200.com",
    "photo_max_orig": "testorig.com"
}
```

**`GET`** 
- */user* (with cookie `user_id`)

**`GET`** 
- /user/all

### Team routes
**`POST`**
+ */team/create*
```json5
{
    "id": 123123,
    "title": "Liquid",
    "tag": "Liquid",
    "photo_100": "test.jpg",
    "photo_200": "test.jpg",
    "photo_max_orig": "test.jpg",
    "rating": 500,
    "match_id": 228
}
```

**`GET`** 
- */team* (with cookie `team_id`)

**`GET`** 
- /team/all