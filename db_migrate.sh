#!/bin/bash
# Install the migrate binary by following the installation guide https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
migrate -database postgres://bucketeer:bucketeer_pass@localhost:5432/bucketeer_db?sslmode=disable -path db/migrations up