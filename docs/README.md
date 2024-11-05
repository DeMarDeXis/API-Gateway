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