### Guest API
#### LOGIN {Admin, GA, Employee}

Request :

- Method : `POST`
- Endpoint : `/auth/login`
- Header :
    - Content-Type : application/json
    - Accept : application/json
- Body :

```json
{
  "email": "string",
  "password": "string"
}
```

Response :
``` json
{
    "status": {
        "code": 200,
        "message": "success"
    },
    "data": {
        "token": "jwt-token"
    }
}
```

#### Register {Admin}

Request :

- Method : `POST`
- Endpoint : `/auth/register`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
  "name": "string",
  "email":"example@mail.com",
  "password":"string",
  "division":"string",
  "position":"string",
  "role":"string",
  "contact":"string"
}
```

Response :
``` json
{
    "status": {
        "code": 201,
        "message": "success"
    },
    "data": {
        "id": "string",
        "name": "string",
        "email": "example@mail.com",
        "division": "string",
        "position": "string",
        "role": "string",
        "contact": "string"
    }
}
```
