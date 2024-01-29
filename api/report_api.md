### Download Report API
#### 1. Download Report {Admin, GA}

Request :

- Method : `GET`
- Endpoint : `/Facilities/id/{id}`
- Query : `start` optional
- Query : `end` optional
- Default check for last month transaction if query empty
- Header :
    - Authorization : Bearer Token
    - Content-Type : application/json
    - Accept : application/json

Response :
excel File
