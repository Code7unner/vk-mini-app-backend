#vk-mini-app-backend

[![Docs Status](https://img.shields.io/badge/Docs-Status-brightgreen.svg?style=flat)](https://github.com/Code7unner/vk-mini-app-backend/blob/main/README.md)

###API

<!-- TOC depthFrom:1 depthTo:2 withLinks:1 updateOnSave:1 orderedList:0 -->
- [User](#user-routes)
- [Team](#team-routes)
- [Steam](#steam-routes)
- [Match](#match-routes)
<!-- /TOC -->

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
- */team/:id* (with cookie `user_id`)

**`GET`** 
- /team/all

### Steam routes
**`GET`**
+ */steam/login*

**`GET`**
+ */steam/user*

### Match routes
**`POST`**
+ */match/create*
```json5
{
    "team_left_id": 123123,
	"team_right_id": 322322,
	"time_created": "22-04-2020",
	"time_started": "22-04-2020",
	"team_left_ready": true,
	"team_right_ready": false
}
```
**`GET`**
+ */match/:id*
