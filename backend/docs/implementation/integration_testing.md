# Integration Testing

Initially integration testing was done by connecting to a local Postgres database.
This approach is cumbersome. A new developer needs to setup the database according to a special configuration in order to just run tests.

Then, we used `docker compose`. However this came with its own problems.
Port collisions, stale data, lack of automation, race conditions. These problems results in flaky tests.
