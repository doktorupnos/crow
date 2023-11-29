# Implementation Details

## Architecture

Decoupled RESTful web service.

## Database Layer - ORM

The communication with the Postgres Database is done via ORM. GORM is Go's most popular ORM and it was used solely for that reason.

Note: If we had a choice, we would go the route of a migration tool like `goose` + a SQL compiler like `sqlc`.

## Dependency Injection

Go offers great support for Dependency Injection with its flexible interfaces.

To implement the restriction of the assignment that Dependency Injection **must** be used, the Repository Pattern was deployed.

The Repository pattern allows for the decouple of the data from the business logic and the controllers. It also aids the 3 Layer restriction.

An example:

For the `User` type, an interface named `UserRepo` references all the methods a type must implement in order to be used as a `UserRepo`. 
Then in the database layer a type named `GormUserRepo` implements `UserRepo` via the  GORM API. `GormUserRepo` encapsulates the database communication but does not have any business logic. 
Finally, the business logic is implemented in the application layer by a type named `UserService`.
It is worth noting that `UserService` doesn't depend on a specific data source, making it easy to test via mocking. Though if that's worth doing is a separate issue altogether. The important part is that we are not explicitly depending on a data source.

## Authorization / Authentication

Authentication middleware was put in place in order have DRY code. 

Basic password authentication was used in the login and user deletion endpoints. While JWT access tokens are used across the majority of the API to handle safe access to resources.

## Testing

Unit tests will be implemented where necessary.

E2E test are a WIP (should be performed at the end of every sprint, or in a TDD fashion with each deployment).
