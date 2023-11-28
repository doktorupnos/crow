# API DOCUMENTATION - Login & Logout

## POST /login

Requires Authorization: Basic
Returns a JWT as a cookie under the name: "token"

## POST /logout

Requires: JWT
Returns:  New JWT as a cookie in the same manner as Login. 
The created token has already expired, effectively invalidating the token that is set in the cookies.
