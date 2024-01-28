### Employee Management API
#### 1. Get Employee by Email {Admin, GA, Employee}

Request :

- Method : `GET`
- Endpoint : `/employees/email/{email}`
- Params : `email`
- Header :
    - Authorization : Bearer Token
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
        "id": "string",
        "name": "string",
        "email": "example@mail.com",
        "division": "string",
        "position": "string",
        "role": "ADMIN/GA/EMPLOYEE",
        "contact": "string"
    }
}
```

#### 2. Get Employees {Admin}

Request :

- Method : `GET`
- Endpoint : `/employees`
- Query : `page` optional
- Query : `size` optional
- Header :
    - Authorization : Bearer Token
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
            "name": "string",
            "email": "example@mail.com",
            "division": "string",
            "position": "string",
            "role": "ADMIN/GA/EMPLOYEE",
            "contact": "string"
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


#### 3. Get Deleted Employees {Admin}

Request :

- Method : `GET`
- Endpoint : `/employees/deleted`
- Query : `page` optional
- Query : `size` optional
- Header :
    - Authorization : Bearer Token
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
            "name": "string",
            "email": "example@mail.com",
            "division": "string",
            "position": "string",
            "role": "ADMIN/GA/EMPLOYEE",
            "contact": "string"
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

#### 4. Get Employee by ID {Admin, GA, Employee}

Request :

- Method : `GET`
- Endpoint : `/employees/{id}`
- Params : `id` 
- Header :
    - Authorization : Bearer Token
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
        "id": "string",
        "name": "string",
        "email": "example@mail.com",
        "division": "string",
        "position": "string",
        "role": "ADMIN/GA/EMPLOYEE",
        "contact": "string"
    }
}
```

#### 5. Get Employee by ID {Admin}

Request :

- Method : `DELETE`
- Endpoint : `/employees/{id}`
- Params : `id`
- Header :
    - Authorization : Bearer Token
    - Content-Type : application/json
    - Accept : application/json

Response :
``` json
{
    "status": {
        "code": 200,
        "message": "success"
    },
    "data": "success"
}
```

#### 6. Create Employee {Admin}

Request :

- Method : `POST`
- Endpoint : `/employees`
- Header :
    - Authorization : Bearer Token
    - Content-Type : application/json
    - Accept : application/json

Request :

```json
{
	"name" : "string", 
	"email": "example@mail.com",
	"password":"string",
	"division":"string",
	"position":"string",
	"role":"ADMIN/GA/EMPLOYEE",
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
        "role": "ADMIN/GA/EMPLOYEE",
        "contact": "string"
    }
}
```


#### 7. Update Employee {Admin}

Request :

- Method : `PATCH`
- Endpoint : `/employees/{id}`
- Params : `id`
- Header :
    - Authorization : Bearer Token
    - Content-Type : application/json
    - Accept : application/json

Request :

```json
{
	"name" : "string", 
	"email": "example@mail.com",
	"password":"string",
	"division":"string",
	"position":"string",
	"role":"ADMIN/GA/EMPLOYEE",
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
        "role": "ADMIN,GA,EMPLOYEE",
        "contact": "string"
    }
}
```
