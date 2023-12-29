# Environment Variables and Configuration

NOTE: Environment Variable names are not final.

The server expects the following environment variables to be set:

- `ADDR`: `hostname:port`. The address for the server to bind to and listen for incoming requests.
- `CORS_ORIGIN`: `http://*`. The single origin allowed by the server. This is to be the URL of the front-end.
- `DSN`: `postgres://postgres:postgres@localhost:5432/crow`. The Data Source Name for the Postgres database.
- `JWT_SECRET`: the secret key used to sign JWTs. Use `openssl rand -base64 64` to generate one for youself.
- `JWT_LIFETIME`: a parsable strings by Go's `time.ParseDuration` function. Serves as the lifetime of a JWT.
- DEFAULT_POSTS_PAGE_SIZE: default page size for the posts.
- DEFAULT_FOLLOWS_PAGE_SIZE: default page size for the following and followers lists.

## Local Development

When running the server, a `-local` flag can be used to depend on a `.env` file.
The `.env` file is loaded using `godotenv`. BEWARE: `godotenv` does not override environment variables.
