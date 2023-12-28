# Profile

## GET /profile

- Description: View a user's profile.
- Requires: JWT.
- Optional Query Parameter:
  - u : a user's name
    `GET /profile?u=zoumas`
    If the u query parameter is not set, the jwt owner's profile is returned.

Response Body:
User 'zoumas' wants to view the profile of User 'doukas' who they follow.

```json
{
  "name": "doukas",
  "follower_count": 1,
  "following_count": 0,
  "id": "1fad8a3a-7c5b-4f54-a6ee-c8fd0d917adb",
  "self": false,
  "following": true
}
```
