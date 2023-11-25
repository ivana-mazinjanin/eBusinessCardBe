# eBusinessCardBe 

# Golang REST API 


## Development Setup
First of all create a new **.env** file in the root of the project directory. see **example.env** file for all required environment variables.

After setting up environment variables, run the following command to start development and database servers.

```bash
docker compose up
```
> before running this command **docker** and **docker compose** must be installed.

## Technology
- Language (golang)


### Libraries
- Router (gorilla/mux)
- Server (net/http)
- Live Reload (cosmtrek/air)


## API Documentation

### Authentication
> **POST** ``/auth/login``

Login with username/email and password.

##### Body

```json
{
    "id": "abc123",
    "password": "abc123",
}
```

#### Output

```json
{
    "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
    "user": {
        "id": "1",
        "name": "ABC 123",
        "username": "abc123",
        "email": "admin@abc123.io",
    }
}
```

> **POST** ``/auth/signup``

Create a new user in the database.

##### Body

```json
{
    "name": "ABC 123",
    "username": "abc123",
    "email": "admin@abc123.io",
    "password": "abc123",
}
```

#### Output

```json
{
    "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
    "user": {
        "id": "1",
        "name": "ABC 123",
        "username": "abc123",
        "email": "admin@abc123.io",
    }
}
```

### Data Manipulation

All endpoints are protected, must send valid **jwt** as ``Authorization`` header with each request.

> **GET** &nbsp; ``/food/all``

Get All Food Items

#### Output

```json
[
    {
        "id": "1",
        "name": "Apples",
        "quantity": 100,
        "selling_price": "100 USD",
    },
    {
        "id": "2",
        "name": "Mangos",
        "quantity": 97,
        "selling_price": "120 USD",
    }
]
```

> **GET** &nbsp; ``/food/<name>``

Get single Food Item by its name. name should be lowercase (e.g /food/apples)

#### Output

```json
{
    "id": "1",
    "name": "Apples",
    "quantity": 100,
    "selling_price": "100 USD",
}
```

> **POST** &nbsp; ``/food``

Add a new food item to the database.

##### Body

```json
{
    "name": "Oranges",
    "quantity": 44,
    "selling_price": "80 USD",
}
```

#### Output

```json
{
    "id": "1",
    "name": "Oranges",
    "quantity": 44,
    "selling_price": "80 USD",
}
```
> **DELETE** &nbsp; ``/food/<id>``

Delete one Food Item from the database.

#### Output

```json
{
    "id": "1",
    "name": "Oranges",
    "quantity": 44,
    "selling_price": "80 USD",
}
```