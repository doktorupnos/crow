# API DOCUMENTATION - Users

## POST /users

Description: Creates a new User.
Request Body:
```json
{
    "name": "zoumas",
    "password": "ilias"
}
```
Status Code: 
* 201 - created 
* 400 - anything else
Checks:
* 0 < Name Length < 20
* 0 < Password Length < 72
* Valid Characters: a-z, A-Z, 0-9, `_` (the underscore). Applied for both the name and the password.

## GET /users

Description: Retrieve all Users.
Request Body: None
Status Code: 
* 200 - retrieved 
* 500 - failed to retrieve from database
Response Body:
```json
[
    {
        "id": "abf71427-2d2c-436a-bffd-8b697e86cf34",
        "created_at": "2023-11-28T00:16:40.213105+02:00",
        "updated_at": "2023-11-28T00:16:40.213105+02:00",
        "name": "zoumas"
    },
    {
        "id": "9ad8a9e4-bf8c-456e-a984-4106978d7c6a",
        "created_at": "2023-11-28T00:44:36.370385+02:00",
        "updated_at": "2023-11-28T00:44:36.370385+02:00",
        "name": "doukas"
    },
    {
        "id": "298bde74-3d86-420b-92f4-a2abb2a91a8d",
        "created_at": "2023-11-28T00:44:47.492365+02:00",
        "updated_at": "2023-11-28T00:44:47.492365+02:00",
        "name": "papazoglou"
    }
]
```

## GET /users/{name}

Description: Get User by Name
Request Body: None, the name is given from the URL parameter.

Example: 

`GET /users/zoumas`

Response Body:
```json
{
    "id": "abf71427-2d2c-436a-bffd-8b697e86cf34",
    "created_at": "2023-11-28T00:16:40.213105+02:00",
    "updated_at": "2023-11-28T00:16:40.213105+02:00",
    "name": "zoumas"
}
```

## DELETE /users

With Authorization: Basic

Request Body: None
Response Body: None
Status Code: 
* 200 - deleted 
* 401 - unauthorized
* 500 - Database Error

## PUT /users

With JWT

Request Body:
```json
{
    "name": "zoumas",
    "password": "ilias"
}
```
Response Body: The updated user resource
Status Code:
* 200 - Updated
* 401 - Unauthorized
* 500 - Database Error
