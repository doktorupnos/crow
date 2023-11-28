# API Documentation

### GET /healthz

* Description: Health-Check or Readiness endpoint.
* Status Code: 200  when the service is available. TODO: Enhance to return a 503 when not.

## Admin Endpoints

### POST /admin/panic

* Description: Make the server panic.
* Status Code: 500
* Goal: The server should automatically recover from panics.

### GET /admin/error

* Description: Test the JSON reporting.
* Status Code: 500
* Response Body:
```
{
    "error": "Internal Server Error"
}
```

### POST /admin/sleep

* Description: Sleep for a minute before responding.
* Status Code: 200 if not cancelled.
* Goal: Handle cancellation with timeouts & graceful shutdown.

### POST /admin/jwt

* Description: Validates a JWT.
* Status Code:
* 200 - Valid
* 401 - Not Valid
* Goal: Test the implementatio of JWTs.
