### Facilities Management API
#### 1. Get Facility Detail by ID {Admin, GA, Employee}

Request :

- Method : `GET`
- Endpoint : `/Facilities/id/{id}`
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
        "codeName": "string",
        "facilitiesType": "string",
        "status": "string",
        "createdAt": "string",
        "updatedAt": "string"
    }
}
```

#### 2. Get Facility Detail by Code Name {Admin, GA, Employee}

Request :

- Method : `GET`
- Endpoint : `/Facilities/name/{name}`
- Params : `name`
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
        "codeName": "string",
        "facilitiesType": "string",
        "status": "string",
        "position": "string",
        "createdAt": "string",
        "updatedAt": "string"
    }
}
```

#### 3. Get All Facilities {Admin}

Request :

- Method : `GET`
- Endpoint : `/facilities`
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
            "codeName": "string",
            "FacilitiesType": "string",
            "status": "string"
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

#### 4. Get All Facilities by Type {Admin}

Request :

- Method : `GET`
- Endpoint : `/facilities/type/{type}`
- Query : `FacilitiesType`
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
            "codeName": "string",
            "FacilitiesType": "string",
            "status": "string"
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

#### 5. Get All Facilities by Status {Admin}

Request :

- Method : `GET`
- Endpoint : `/facilities/status/{status}`
- Query : `status`
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
            "codeName": "string",
            "FacilitiesType": "string",
            "status": "string"
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

#### 6. Get Deleted Facilities {Admin}

Request :

- Method : `GET`
- Endpoint : `/facilities/deleted`
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
            "codeName": "string",
            "FacilitiesType": "string",
            "status": "string",
            "createdAt": "string",
            "updatedAt": "string",
            "deletedAt": "string"
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

#### 7. Create Facility {Admin}

Request :

- Method : `POST`
- Endpoint : `/facilities`
- Header :
    - Authorization : Bearer Token
    - Content-Type : application/json
    - Accept : application/json

Request :

```json
{
	"codeName" : "string", 
	"FacilitiesType": "string",
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
        "codeName": "string",
        "FacilitiesType": "string",
        "status": "AVAILABLE", (Default Value)
        "createdAt": "string"
    }
}
```


#### 7. Update Facility {Admin}

Request :

- Method : `PUT`
- Endpoint : `/facilities/{id}`
- Params : `id`
- Header :
    - Authorization : Bearer Token
    - Content-Type : application/json
    - Accept : application/json

Request :

```json
{
	"codeName" : "string", 
	"FacilitiiesType": "string",
	"status":"string"
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
        "id" : "string", 
        "codeName" : "string", 
        "FacilitiiesType": "string",
        "status":"string",
        "updatedAt": "string"
    }
}
```
#### 8. Delete Facility by Id {Admin}

Request :

- Method : `DELETE`
- Endpoint : `/facilities/{id}`
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
    "data": { "id has been deleted"}
}
```
#### 7. Update Facility {Admin}

Request :

- Method : `PUT`
- Endpoint : `/facilities/{id}`
- Params : `id`
- Header :
    - Authorization : Bearer Token
    - Content-Type : application/json
    - Accept : application/json

Request :

```json
{
	"codeName" : "string", 
	"FacilitiiesType": "string",
	"status":"string"
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
        "id" : "string", 
        "codeName" : "string", 
        "FacilitiiesType": "string",
        "status":"string",
        "updatedAt": "string"
    }
}
```
#### 8. Delete Facility by Code Name {Admin}

Request :

- Method : `DELETE`
- Endpoint : `/facilities/{name}`
- Params : `name`
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
    "data": { "codeName has been deleted"}
}
```
