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

- 201 - Created successfully
- 400 - Bad Request - Failed to decode request body or Body is not valid
- 401 - Unauthorized - JWT has expired

Checks:

- 0 < Body Length <= 128

## GET /posts

- Description: Get all Posts.
- Requires: JWT.
- Pagination:
- Query Parameters:
  - page: Paginated list. The first page has index 0
  - limit: Limit the entries for a page. If not set a sane default is provided.
  - u: The user name for the owner of the post.
    Example: GET /posts?u=zoumas&page=2&limit=10
    Get the third page of 10 at most posts by user zoumas.

Example Response Body: `http://localhost:8000/posts?u=zoumas&page=0`

````json
[
    {
        "user_name": "zoumas",
        "body": "What?",
        "id": "4646178c-5f70-4cc5-9394-77fecb6f3319",
        "created_at": 1703803617,
        "updated_at": 1703803617,
        "user_id": "d65c089f-d1cf-41e2-8523-78322143a1fe",
        "likes": 0,
        "liked_by_user": false,
        "self": true
    },
    {
        "user_name": "zoumas",
        "body": "Comments?\n\n\n",
        "id": "770351cd-3f62-4456-a3d0-365c6be087d5",
        "created_at": 1703773627,
        "updated_at": 1703773627,
        "user_id": "d65c089f-d1cf-41e2-8523-78322143a1fe",
        "likes": 1,
        "liked_by_user": true,
        "self": true
    },
    {
        "user_name": "zoumas",
        "body": "No?\n",
        "id": "3716db8a-8d47-4e7f-9a5f-1f65ac17b7f9",
        "created_at": 1703773606,
        "updated_at": 1703773606,
        "user_id": "d65c089f-d1cf-41e2-8523-78322143a1fe",
        "likes": 1,
        "liked_by_user": true,
        "self": true
    },
    {
        "user_name": "zoumas",
        "body": "10",
        "id": "e3815421-9e48-4395-9cae-81e8fa00bd9e",
        "created_at": 1703588260,
        "updated_at": 1703588260,
        "user_id": "d65c089f-d1cf-41e2-8523-78322143a1fe",
        "likes": 1,
        "liked_by_user": true,
        "self": true
    },
    {
        "user_name": "zoumas",
        "body": "9",
        "id": "9605c291-3943-4ca0-8ed4-029a832a0e9f",
        "created_at": 1703588257,
        "updated_at": 1703588257,
        "user_id": "d65c089f-d1cf-41e2-8523-78322143a1fe",
        "likes": 1,
        "liked_by_user": true,
        "self": true
    }
]
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
````

Status Code:

- 200 - Updated successfully
- 400 - Failed to parse Post's ID or request's body is not valid.

Checks:

- 0 < Body Length <= 128

## DELETE /posts/{id}

Requires: JWT.

Input: Post's ID via URL parameter.

Example Request: `DELETE http://localhost:8000/posts/2ab14a7f-ecb0-40fd-96ba-99fddddc53cb`

Status Code:

- 200 - Delete successfully
- 400 - Failed to parse Post's ID or Post doesn't belong to User.
- 401 - Unauthorized
