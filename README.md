# Crow

Crow is a centralized microblogging and networking platform that allows users to connect and post short blog material.
Its emphasis is simplicity, and privacy.

# Subject of the work
University Task from our professor in "Software Engineering Special Topics" course

# License
The License of the project is signed under GPLv3.

# Development Methodology
During the development of this project, we will be using a private access [Kanban](https://en.wikipedia.org/wiki/Kanban) 
Board using [Trello](https://trello.com/).

## Team Members
* @doktorupnos Frontend / UI/UX
* @marios-pz Frontend / DevOps
* @zoumas Backend 

## Deadline
The work begins on Monday, November 20, 2023, and concludes on January 10, 2024.
The project review will take place at the end of the examination period, which also marks the end date.

# Architecture

## Frontend

The interface will be built using React, a widely adopted component based library
that simplifies the process of building complex UI.

## Backend

The backend will be developed using Go, chosen for its simplicity and efficiency.
We are constructing a resilient REST API for our application.
As for the database, we've opted for PostgreSQL, and our interaction with it will be facilitated by GORM.

## Testing

Unit testing is performed within a docker environment as part of the CI/CD process
because it is crucial for our code to undergo testing in as clean an environment as possible.
