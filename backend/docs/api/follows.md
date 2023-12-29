# Follows

## POST /follow

- Description: Follow a user.
- Requires: JWT.
- Request Body:

```json
{
  "user_id"= uuid
}

```

The user_id is the id the user wants to follow.

- Response Status Code:
  - 200: Successfully followed
  - 400: Error

## POST /unfollow

- Description: Unfollow a user.
- Requires: JWT.
- Request Body:

```json
{
  "user_id"= uuid
}

```

- Response Status Code:
  - 200: Successfully unfollowed
  - 400: Error

## GET /following

- Description: Get a paginated list of following users.
- Requires: JWT.
- Query Parameters:
  - u : a user's name
    `GET /following?u=zoumas`
    If the u query parameter is not set, the jwt owner's following list is returned.
  - page: a page is a paginated list of results.
    GET /following?page=2
    If the page query parameter is not set, it is inferred as page=0. 0 is the index for the first page of results.
  - limit: page size.
    GET /following?page=0&limit=10
    Limit the page to 10 entries.
    If the limit query parameter is not set, a sane default is provided.

## GET /followers

- Description: Get a paginated list of follower users.
- Requires: JWT.
- Query Parameters:
  - u : a user's name
    `GET /followers?u=zoumas`
    If the u query parameter is not set, the jwt owner's followers list is returned.
  - page: similar to /following.
  - limit: similar to /following.

## GET /following_count

- Description: Get the total count of following users.
- Requires: JWT.
- Optional Query Parameter:
  - u : a user's name
    `GET /following_count?u=zoumas`
    If the u query parameter is not set, the jwt owner's following count is returned.
- Returns: An integer.

## GET /followers_count

- Description: Get the total count of follower users.
- Requires: JWT.
- Optional Query Parameter:
  - u : a user's name
    `GET /followers_count?u=zoumas`
- Returns: An integer.
  If the u query parameter is not set, the jwt owner's follower count is returned.
