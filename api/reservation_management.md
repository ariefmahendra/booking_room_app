### Reservation Management API

#### 1. Get All Reservation

Request :

- Method : `GET`
- Endpoint : `/api/v1/reservation/`
- Query : `page` optional
- Query : `size` optional
- Header :
  - Authorization : Bearer Token
  - Content-Type : application/json
  - Accept : application/json

Response :

```json
{
  "status": {
    "code": 200,
    "message": "success"
  },
  "data": [
    {
      "id": "string",
      "employee_id": "string",
      "employee_name": "string",
      "room_code": "string",
      "booked_start_date": "2024-01-25T09:00:00Z",
      "booked_end_date": "2024-01-27T11:00:00Z",
      "note": "string",
      "approval_status": "ACCEPT/PENDNG/DECLINE",
      "apprv_note": "string",
      "facilities": [
        {
          "id": "string",
          "code": "string",
          "type": "string"
        }
      ]
    },
  ],
  "paging": {
    "page": int,
    "rowsPerPage": int,
    "totalPages": int,
    "totalRows": int
  }
}
```

#### 2. Get Reservation By Id

Request :

- Method : `GET`
- Endpoint : ` /api/v1/reservation/get/:id`
- Params : `id`
- Header :
  - Authorization : Bearer Token
  - Content-Type : application/json
  - Accept : application/json

Response :

```json
{
  "status": {
    "code": 200,
    "message": "success"
  },
  "data": {
    "id": "string",
    "employee_id": "string",
    "employee_name": "string",
    "room_code": "string",
    "booked_start_date": "2024-01-25T09:00:00Z",
    "booked_end_date": "2024-01-27T11:00:00Z",
    "note": "string",
    "approval_status": "ACCEPT/PENDNG/DECLINE",
    "apprv_note": "string",
    "facilities": [
      {
        "id": "string",
        "code": "string",
        "type": "string"
      }
    ]
  }
}
```

#### 3. Get Reservation By Employee Id

Request :

- Method : `GET`
- Endpoint : `/api/v1/reservation/employee/:id`
- Params : `id`
- Query : `page` optional
- Query : `size` optional
- Header :
  - Authorization : Bearer Token
  - Content-Type : application/json
  - Accept : application/json

Response :

```json
{
  "status": {
    "code": 200,
    "message": "success"
  },
  "data": [
    {
      "id": "string",
      "employee_id": "string",
      "employee_name": "string",
      "room_code": "string",
      "booked_start_date": "2024-01-25T09:00:00Z",
      "booked_end_date": "2024-01-27T11:00:00Z",
      "note": "string",
      "approval_status": "ACCEPT/PENDNG/DECLINE",
      "apprv_note": "string",
      "facilities": [
        {
          "id": "string",
          "code": "string",
          "type": "string"
        }
      ]
    },
  ],
  "paging": {
    "page": int,
    "rowsPerPage": int,
    "totalPages": int,
    "totalRows": int
  }
}
```

#### 4. Create Reservation

Request :

- Method : `POST`
- Endpoint : `/api/v1/reservation/ `
- Header :
  - Authorization : Bearer Token
  - Content-Type : application/json
  - Accept : application/json

```json
{
  "email": "budi@mail.com",
  "room_code": "R001",
  "booked_start_date": "2024-02-01T09:00:00Z",
  "booked_end_date": "2024-02-02T11:00:00Z",
  "note": "Tesss",
  "facilities": [
    {
      "code": "SCR1"
    },
    {
      "code": "PRJ3"
    }
  ]
}
```

Response :

```json
{
  "status": {
    "code": 201,
    "message": "success"
  },
  "data": {
    "id": "24c2d37a-9475-4687-a6dc-167ce7c16f83",
    "employee_id": "89037fe3-53c0-44c3-aa74-644e381ab621",
    "employee_name": "Budi",
    "room_code": "R001",
    "booked_start_date": "2024-02-01T09:00:00Z",
    "booked_end_date": "2024-02-02T11:00:00Z",
    "note": "Tesss",
    "approval_status": "PENDING",
    "apprv_note": "",
    "facilities": [
      {
        "id": "db3dc28e-b888-47f8-850a-8228e42ea379",
        "code": "SCR1",
        "type": "screen"
      },
      {
        "id": "24bcf012-34bf-45d1-a986-117586eb362e",
        "code": "PRJ3",
        "type": "projector"
      }
    ]
  }
}
```

#### 5. Get Approval Pending List

Request :

- Method : `GET`
- Endpoint : `/api/v1/reservation/approval`
- Query : `page` optional
- Query : `size` optional
- Header :
  - Authorization : Bearer Token
  - Content-Type : application/json
  - Accept : application/json

Response :

```json
{
  "status": {
    "code": 200,
    "message": "success"
  },
  "data": [
    {
      "id": "string",
      "employee_id": "string",
      "employee_name": "string",
      "room_code": "string",
      "booked_start_date": "2024-01-25T09:00:00Z",
      "booked_end_date": "2024-01-27T11:00:00Z",
      "note": "string",
      "approval_status": "ACCEPT/PENDNG/DECLINE",
      "apprv_note": "string",
      "facilities": [
        {
          "id": "string",
          "code": "string",
          "type": "string"
        }
      ]
    },
  ],
  "paging": {
    "page": int,
    "rowsPerPage": int,
    "totalPages": int,
    "totalRows": int
  }
}
```

#### 6. Put Approval Status

Request :

- Method : `PUT`
- Endpoint : `/api/v1/reservation/approval`
- Header :
  - Authorization : Bearer Token
  - Content-Type : application/json
  - Accept : application/json

```json
{
  "id": "string",
  "approval_status": "ACCEPT/DECLINE",
  "apprv_note": "string"
}
```

Response :

```json
{
  "status": {
    "code": 200,
    "message": "success"
  },
  "data": [
    {
      "id": "string",
      "employee_id": "string",
      "employee_name": "string",
      "room_code": "string",
      "booked_start_date": "2024-01-25T09:00:00Z",
      "booked_end_date": "2024-01-27T11:00:00Z",
      "note": "string",
      "approval_status": "ACCEPT/DECLINE",
      "apprv_note": "string",
      "facilities": [
        {
          "id": "string",
          "code": "string",
          "type": "string"
        }
      ]
    },
  ],
  "paging": {
    "page": int,
    "rowsPerPage": int,
    "totalPages": int,
    "totalRows": int
  }
}
```

#### 7. Delete Reservation

Request :

- Method : `DELETE`
- Endpoint : `/api/v1/reservation/:id`
- Params : `id`
- Header :

  - Authorization : Bearer Token
  - Content-Type : application/json
  - Accept : application/json

  Response :

  ```json
  {
    "status": {
      "code": 200,
      "message": "success"
    },
    "data": "Reservation Deleted"
  }
  ```

#### 8. Get Available Rooms

Request :

- Method : `GET`
- Endpoint : `/api/v1/reservation/available`
- Header :

  - Authorization : Bearer Token
  - Content-Type : application/json
  - Accept : application/json

- BodyJSON :

```json
{
  "start_date": "2024-01-25 09:00:00",
  "end_date": "2024-01-27 11:00:00"
}
```

Response :

```json
{
"status": {
  "code": 200,
  "message": "success"
},
"data": [
  {
    "id": "string",
    "code_room": "string",
    "room_type": "string",
    "facilities": "string",
    "capacity": int
  },
  {
    "id": "string",
    "code_room": "string",
    "room_type": "string",
    "facilities": "string",
    "capacity": int
  }
]
}
```
