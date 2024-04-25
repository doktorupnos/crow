<h1 align="center">Crow</h1>

Crow is a centralized microblogging and networking platform that allows users to connect and post short blog material.
Its emphasis is simplicity, and privacy.

The project was developed as part of our university course "[Special Topics in Software Engineering](https://ice.uniwa.gr/en/education-1/undergraduate/courses/special-topics-in-software-engineering/)" and is licensed under GPLv3.

## ✏️Development Methodology
During the development of this project, we will be using a private access [Kanban](https://en.wikipedia.org/wiki/Kanban) 
Board using [Trello](https://trello.com/).

![image](https://github.com/doktorupnos/wip-chat/assets/30930688/aa7fe0d2-fd69-407e-a94c-65f87049da84)

## 🧑‍💻Team Members

- [zoumas](https://github.com/zoumas) | Lead Backend <br/>
- [doktorupnos](https://github.com/doktorupnos) | Lead Frontend - UI/UX <br/>
- [marios-pz](https://github.com/marios-pz) | Frontend - DevOps <br/>

## 📐Architecture

Decoupled RESTful web service and a SPA front end. 

## 🧑‍🔧Backend

The backend is being developed in Go, chosen for its simplicity and efficiency.

### 💽Database Layer & ORM

The communication with the Postgres Database is done via ORM. [GORM](https://gorm.io) is Go's most popular ORM and it was used solery for that reason. If we had a choice, we would go with a Database First approach using [goose](https://pressly.github.io/goose/) as a migration tools and a SQL compiler like [sqlc](https://sqlc.dev/). Although as it turned out, GORM's automigration came in handy for integration tests with testcontainters.

An example:

Under the `user` package, the model `user.User` is defined. An interface `user.Repo` declares all the methods a type must implement in order to be used as a `user.Repo`. Then in the database layer a type named `GormUserRepo` implements `user.Repo`. `GormUserRepo` encapsulates the database communication but does is not aware of  any business logic. 
Finally, the business logic is implemented by a type named `user.Service`. 
It is worth noting that `user.Service` doesn't depend on a specific data source, it could be a database, a message broker or a KINESIS stream. Or a in memory structure making it easy to test via mocking.

### 🧬Dependency Injection

Go offers great support for Dependency Injection with its flexible interfaces.

To implement the restriction of the assignment that Dependency Injection must be used, the Repository Pattern was deployed.

The Repository pattern allows for the decoupling of the data from the business logic and the controllers. It also aids the 3 Layer restriction.

### 🔐Auth (Authentication/Authorization)

Authentication middleware makes up for DRY code.

Basic password authentication was used in the `POST /api/login` and `DELETE /api/users` endpoints. 
While JWT access tokens are used across the majority of the API to handle safe access to resources.

### 🧪Integration Testing

Initially integration testing was done by connecting to a local Postgres database. 
This approach is cumbersome. 
A new developer needs to setup the database according to a special configuration in order to just run tests.

Then, we used docker compose. 
However this came with its own problems. 
Port collisions, stale data, lack of automation, race conditions. These problems results in flaky tests.

We concluded to use [testcontainters](https://testcontainers.com/).
We can achieve indempotent behavior while testing what matters against real services.

## 🧑‍🎨Frontend

The web-app interface is built using the [Next.js](https://nextjs.org/) React framework. No external CSS frameworks are used in the final UI design.

Most of the icons used are taken directly from [Bootstrap Icons](https://icons.getbootstrap.com/).
