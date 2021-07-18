#!/usr/bin/env bash
#
docker run --name comments-api-db -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
