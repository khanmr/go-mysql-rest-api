# Golang REST API with MySQL Database

## Summary

This is a rest api written in golang for a books database. Users can create, read, update and delete books using the endpoints.

## Quick Start

```
Install mux router:

go get -u github.com/gorilla/mux
```

```
Create Database and insert data:

CREATE DATABASE gorestapi;

USE gorestapi;

CREATE TABLE books (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    isbn VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL
);

INSERT INTO books (isbn, title, author) VALUES (
    '9781451639612',
    'The 7 Habits of Highly Effective People: Powerful Lessons in Personal Change',
    'Stephen Covey'
);

INSERT INTO books (isbn, title, author) VALUES (
    '014312417X',
    'Mastery',
    'Robert Green'
);
```

```
Update username and password: "<username>:<password>@/gorestapi"
```

```
go run main.go
```

### Endpoints

```
GET localhost:3000/books
GET localhost:3000/books/{id}
POST localhost:3000/books
PUT localhost:3000/books/{id}
DELETE localhost:3000/books/{id}

*Use Postman to send POST, PUT, DELETE requests
```
