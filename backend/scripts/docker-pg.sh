#!/bin/bash

docker run -d \
  --name crow_pg \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=crow \
  -p 5432:5432 \
  postgres
