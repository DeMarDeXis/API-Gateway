# API Gateways's Routes

HOST = ```http://localhost:8083```

Services:
- Users = -etc-
- Auth = https://github.com/DeMarDeXis/Auth
- Lessons = https://github.com/DeMarDeXis/lessons

1. **POST** > sign_up = ```/auth/sign-up```\
_JSON-Body:_
```
{
    "email": "FredieFreeman@gmail.com",
    "name": "Fredie",
    "surname": "Freeman",
    "login": "MVP",
    "password": "DodgersLA"
}
```
_JSON-Response:_
```
{
    "id": 3
}
```
2. **POST** > sign_in = ```/auth/sign-in```\
_JSON-Body:_\
```
{
    "login": "MVP",
    "password": "DodgersLA"
}
```
_JSON-Response:_
```
{
    "token": "7cb0d735cf31caa5e219473831ecb4e98504a0cc7901f54f6e7024798b5ec164"
}
```
3. **POST** > create_course = ```/courses/create```\
   _JSON-Body:_
```
{
    "name": "myname",
    "desc": "SupportYankees"
}
```

_JSON-Response:_
```
"status": "ok"
```
4. **GET** > get_courses = ```/courses/```\
   _JSON-Response:_
```
[
    {
        "course_id": 1,
        "name": "IT",
        "desc": "Its IT course",
        "created_at": "2024-10-16T08:26:15.965443Z",
        "updated_at": "2024-10-16T08:26:15.965443Z",
        "owner_id": 1
    },
    {
        "course_id": 2,
        "name": "ADMINISTRATION",
        "desc": "Its ADMINISTRATION course",
        "created_at": "2024-10-16T08:29:25.181718Z",
        "updated_at": "2024-10-16T08:29:25.181718Z",
        "owner_id": 1
    },
    {
        "course_id": 6,
        "name": "Python",
        "desc": "SupportDodgers",
        "created_at": "2024-11-05T06:39:22.052693Z",
        "updated_at": "2024-11-05T06:39:22.052693Z",
        "owner_id": 7
    },
    {
        "course_id": 7,
        "name": "Hornets",
        "desc": "LaMelo Ball MVP",
        "created_at": "2024-11-11T07:32:27.297443Z",
        "updated_at": "2024-11-11T07:32:27.297443Z",
        "owner_id": 9
    }
]
```

5. **GET** > get_course = ```/courses/7```\
   _JSON-Response:_
```
{
    "course_id": 7,
    "name": "Hornets",
    "desc": "LaMelo Ball MVP",
    "created_at": "2024-11-11T07:32:27.297443Z",
    "updated_at": "2024-11-11T07:32:27.297443Z",
    "owner_id": 9
}
```

6. **PUT** > update_course = ```/courses/7```\
   _JSON-Body:_
```
{
  //"name": "Hornets",
  "desc": "Miller Brandon MVP"
}
```
_JSON-Response:_
```
"status": "ok"
```

7. **DELETE** > delete_course = ```/courses/7```\
   _JSON-Response:_
```
"status": "ok"
```
