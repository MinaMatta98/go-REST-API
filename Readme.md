# Basic Go API

> [!WARNING]
> This repo is still under construction, but is open to public viewing

This is a basic api written in [GoLang](https://go.dev/) for the simple minded that uses the [Repository Design Pattern](https://www.umlboard.com/design-patterns/repository.html)


## Features
Features include:

- API routing (you don't say)
- Database management
- ORM via [GORM](https://gorm.io/) (coming soon...)

## Running the server

```bash
go run ./cmd/main.go
```

Make sure that the database is setup and the corrosponding Database name is saved in your `.env` file with a variable name of `DATABASE_URL` and password as `DB_PASSWORD`.

## Requirements
The following are the base requirements:
- [mariadb](https://mariadb.org/download/) 
- [GoLang](https://go.dev/) version >= 1.20 (Or atleast that's what i'm certain it works on)

> [!NOTE]
> Happy Coding!
