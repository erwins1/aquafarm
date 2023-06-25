# Aquafarm Repository

## Prerequisites
  1. Go v1.19
  2. Mockery (for unit test) detail: https://github.com/vektra/mockery

## Initial Setup
  1. Name this project folder as aquafarm
  2. Put this project to $GOPATH/src/
  3. Copy `.env.example` file to `.env` and update the connection
  4. Create bin folder
  5. run `go mod vendor`

## Binary List
  1. app (HTTP)

## How To Run
  1. Run docker `docker-compose up -d` to create the datastore
  2. Run using command `make run-[binary]`

## Run Test
  1. Run using command `make test`
