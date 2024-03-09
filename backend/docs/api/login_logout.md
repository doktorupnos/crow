# Login & Logout

## POST /login

* Requires: Authorization Basic.
* Returns: a JWT as a cookie under the name "token".

### Errors

* Missing Authorization Basic header
```json
{
  "error": "missing Authorization Basic header"
}
```

* User does not exist

```json
{
  "error": "record not found"
}
```

* Passwords don't match
```json
{
  "error": "wrong password"
}
```

## POST /logout

Requires: JWT.
Returns: renewed JWT in the same manner as Login. The created token has already expired, effectively invalidating the token that is set in the cookies.
