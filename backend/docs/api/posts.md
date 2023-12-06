# Posts

## POST /posts

Description: Creates a new Post.

Requires: JWT.

A post is associated (belongs) to the user that creates it (inferred from the JWT).

Request Body:
```json
{
    "body": "First post!"
}
```

Status Code:
* 201 - Created successfully
* 400 - Bad Request - Failed to decode request body or Body is not valid
* 401 - Unauthorized - JWT has expired

Checks:
* 0 < Body Length <= 128

## GET /posts

Description: Get all Posts.
Requires: JWT.

Example Response Body:
```json
[
    {
        "id": "0b72bd22-c3a1-40ed-a48d-9ea745ced037",
        "created_at": 1701867268,
        "updated_at": 1701867268,
        "body": "7",
        "user_id": "8b42f8bd-d4b9-44c0-ae44-3c77346d2137",
        "user_name": "test",
        "likes": 0,
        "liked_by_user": false
    },
    {
        "id": "2a4e7619-725d-4c8c-94ac-8dea7c202d04",
        "created_at": 1701867265,
        "updated_at": 1701867265,
        "body": "6",
        "user_id": "8b42f8bd-d4b9-44c0-ae44-3c77346d2137",
        "user_name": "test",
        "likes": 1,
        "liked_by_user": true
    }
]```

Status Code:
* 200 - Retrieved successfully
* 404 - Database Error

## PUT /posts/{id}

Requires: JWT.

Request Body:
```json
{
    "body": "First post!"
}
```

Status Code:
* 200 - Updated successfully
* 400 - Failed to parse Post's ID or request's body is not valid.

Checks:
* 0 < Body Length <= 128

## DELETE /posts/{id}

Requires: JWT.

Input: Post's ID via URL parameter.

Example Request: `DELETE http://localhost:8000/posts/2ab14a7f-ecb0-40fd-96ba-99fddddc53cb`

Status Code:
* 200 - Delete successfully
* 400 - Failed to parse Post's ID or Post doesn't belong to User.
* 401 - Unauthorized

