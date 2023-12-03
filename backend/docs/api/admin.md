# Administration/Misc endpoints

A collection of endpoints that are used for testing/demonstration purposes.

## POST /admin/jwt

* Description: Validates a JWT.
* Motivation: Test the server JWT Authentication middleware.
* Status Code:
    * 200 - Authorized
    * 401 - Unauthorized

## POST /admin/panic

* Description: Cause the server to panic.
* Motivation: The server should make use of middleware to automatically recover from run-time erorrs and print a stack trace.
* Status Code: Always 500.

## POST /admin/sleep

* Description: Sleep for 1 minute before responding.
* Motivation: Handling cancellation via timeouts & graceful shutdown.
* Status Code:
    * 200 - The server has slept for a whole minute and responded.
    * 500 - The socket hang up.

* Note: The default duration for graceful shutdown is configured to be 30 seconds. So, a request to sleep immediately followed by a signal to shutdown the server will result in the socket hanging up.

## GET /admin/error

* Description: Always respond with 500 status code and an error message.
* Motivation: Test the server's JSON respondWith functions.
* Response Body:
```
{
    "error": "Internal Server Error"
}
```

