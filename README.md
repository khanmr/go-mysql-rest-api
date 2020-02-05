# Golang REST API with MySQL Database

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
GET api/books
GET api/books/{id}
POST api/books
PUT api/books/{id}
DELETE api/books/{id}

*Use Postman to send POST, PUT, DELETE requests
```
