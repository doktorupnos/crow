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
        "id": "2ab14a7f-ecb0-40fd-96ba-99fddddc53cb",
        "created_at": "2023-11-29T00:57:25.38232+02:00",
        "updated_at": "2023-11-29T00:57:25.38232+02:00",
        "body": "First post!",
        "user_id": "abf71427-2d2c-436a-bffd-8b697e86cf34",
        "user_name": "ilias"
    },
    {
        "id": "1d770217-4b94-4b75-ba95-3ddc1199dd66",
        "created_at": "2023-11-29T01:06:30.365035+02:00",
        "updated_at": "2023-11-29T01:06:30.365035+02:00",
        "body": "🐹",
        "user_id": "25a789c2-21a0-48a9-96ec-f869beb50760",
        "user_name": "stefanos"
    }
]
```

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

