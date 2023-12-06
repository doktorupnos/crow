# Post Likes

## POST /post_likes

Description: Like a post.
Requires: JWT.
Request Body:
```json
{
    "post_id": "9275a3f8-eb56-4b73-bad0-7e5f23e6329b"
}
```
Status Code:
* 201 - Created
* 400 - Bad Request
* 401 - Unauthorized

## DELETE /post_likes

Description: Unlike a post.
Requires: JWT.
Request Body:
```json
{
    "post_id": "9275a3f8-eb56-4b73-bad0-7e5f23e6329b"
}
```
Status Code:
* 200 - Unliked successfully
* 400 - Bad Request
* 401 - Unauthorized
