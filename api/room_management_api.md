### Employee Management API

#### 1. Get Rooms {Admin, Employee, GA}

Request :

- Method : `GET`
- Endpoint : `/room`
- Query : `page` optional
- Query : `size` optional
- Header :
    - Content-Type : application/json
    - Accept : application/json

Response :
``` json
{
    "status": {
        "code": 200,
        "message": "success"
    },
    "data": [
        {
            "id": "string",
            "code_room": "string",
            "room_type": "example@mail.com",
            "capacity": "string",
            "facilities": "string",
        }
    ],
    "paging": {
        "page": 1, (default value)
        "rowsPerPage": 5, (default value)
        "totalPages": 1,
        "totalRows": 1
    }
}
```


#### 2. Get By Id Room {Admin, Employee, GA}

Request :

- Method : `GET`
- Endpoint : `/room/:id`
- Header :
    - Content-Type : application/json
    - Accept : application/json

Response :
``` json
{
    "status": {
        "code": 200,
        "message": "success"
    },
    "data": {
        {
            "id": "string",
            "code_room": "string",
            "room_type": "example@mail.com",
            "capacity": "string",
            "facilities": "string",
        }
    }
}

#### 3. Create Room {Admin}

Request :

- Method : `POST`
- Endpoint : `/room/create`
- Header :
    - Content-Type : application/json
    - Accept : application/json

Request :

```json
{
    "code_room": "string",
    "room_type": "string",
    "facilities": "string",
    "capacity": int,
}

Response :
``` json
{
    "status": {
        "code": 201,
        "message": "created successfully"
    },
    "data": {
        "id": "string",
        "code_room": "string",
        "room_type": "string",
        "facilities": "string",
        "capacity": int,
    }
}
```


#### 4. Update Room {Admin}

Request :

- Method : `PUT`
- Endpoint : `/room/:id`
- Params : `id`
- Header :
    - Content-Type : application/json
    - Accept : application/json

Request :

```json
{
    "code_room": "string",
    "room_type": "string",
    "facilities": "string",
    "capacity": int,
}
```

Response :
``` json
{
    "status": {
        "code": 201,
        "message": "update successfully"
    },
    "data": {
        "id": "string",
        "code_room": "string",
        "room_type": "string",
        "facilities": "string",
        "capacity": int,
    }
}
