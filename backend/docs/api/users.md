# Users

## POST /users

Description: Creates a new User.

Request Body:
```json
{
    "name": "zoumas",
    "password": "ilias"
}
```

Conditions:
* 0 < Name Length < 20
* 0 < Password Length < 72
* Valid Characters: a-z, A-Z, 0-9, `_`. Applied for both the name and the password.

* Status Code: 
* 201 - User created successfully 
* 400 - Anything else

## GET /users

Description: Retrieve all Users.

Response Body Example:
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

Status Code: 
* 200 - Retrieved successfully
* 500 - Database Error

## GET /users/{name}

Description: Get User by Name.

Input: The user's name via URL Parameter.

Example Request: 

`GET /users/zoumas`

Example Response Body:
```json
{
    "id": "abf71427-2d2c-436a-bffd-8b697e86cf34",
    "created_at": "2023-11-28T00:16:40.213105+02:00",
    "updated_at": "2023-11-28T00:16:40.213105+02:00",
    "name": "zoumas"
}
```

Status Code:
* 400 - Missing URL parameter
* 404 - Database Error
* 200 - Retrieve successfully

## PUT /users

Description: Updates a User's name and password.

Requires: JWT

Note: The token is set via Cookies and automatically sent to the server. So the client doesn't need to worry about manually setting Authorization headers.

Existing User: `zoumas:ilias`

Example Request Body:
```json
{
    "name": "ilias",
    "password": "zoumas"
}
```

Example Response Body:
```json
{
    "id": "abf71427-2d2c-436a-bffd-8b697e86cf34",
    "created_at": "2023-11-28T00:16:40.213105+02:00",
    "updated_at": "2023-11-28T23:50:15.402891+02:00",
    "name": "ilias"
}
```

Status Code:
* 200 - Updated successfully
* 400 - Couldn't decode json request body
* 401 - Unauthorized
* 500 - Database Error or violation of the rules set for the Name and Password

## DELETE /users

Description: Deletes a User.

Requires: Authorization Basic.

* Header: Authorization
* Value: Basic username:password

Status Code: 
* 200 - Deleted successfully
* 401 - Unauthorized
* 500 - Database Error
* 
