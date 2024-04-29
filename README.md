# Project wishlist

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## Environment

Create a .env file in the root of the project with the following content

```bash
PORT=8080
HOST=localhost
DATABASE_URL=file:.wishlist/wishlist.db

ADMIN_USER=admin
ADMIN_PASSWORD=admin
```

## Database

```bash
GOOSE_MIGRATION_DIR=schema GOOSE_DRIVER=sqlite3 GOOSE_DBSTRING=.wishlist/wishlist.db goose up
```

## Prerequisites

- https://github.com/cosmtrek/air
- https://github.com/pressly/goose
- https://docs.sqlc.dev/
- https://github.com/pvolok/mprocs

```bash
go install github.com/cosmtrek/air@latest
go install github.com/pressly/goose/v3/cmd/goose@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install github.com/a-h/templ/cmd/templ@latest
```
