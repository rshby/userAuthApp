
## UserAuthAPP
### How to run this app
**1. Clone Repository** <br>
Clone this repository into your local computer. Make sure that Docker is already installed on your computer.

**2. Open In IDE** <br>
Open it on your favorite IDE, you can use Goland, VSCode. etc..

**3. Run With Docker Compose** <br>
Open terminal/cmd and **navigate to this root project folder**. Type: <br>
```
docker-compose up -d
```
**4. App and Database will ready to use** <br>
This app will use port `5005` on your local computer. <br>
The database name is `userauth_db`, and it will expose port 5432 since it uses PostgreSQL. Inside the database, two tables have already been created : `accounts` and `users`

<br>

### List Of Endpoint
| Endpoint  | Method |
| ------------- |:-------------:|
| http://localhost:5005/api/v1/register|POST|
| http://localhost:5005/api/v1/user| POST|
| http://localhost:5005/api/v1/login| POST|
| http://localhost:5005/api/v1/users| GET| <br>
* **Register Account** <br>

Type this endpoint to postman with **POST** method. <br>
```
http://localhost:5005/api/v1/register
```
Request Body :
```
{
  "username" : "john@gmail.com",
  "password" : "123456789012"
}
```
Response Body :
```
{
    "status_code": 200,
    "status": "ok",
    "message": "success create account",
    "data": {
        "id": 2,
        "username": "john@gmail.com",
        "password": "$2a$10$J3laTGDqqK0z8Srn5q3RueaG3awFzvzMBpyHD6E0REuRa.CcZscI."
    }
}
```

* **Create User** <br>

Type this endpoint to postman with **POST** method.
```
http://localhost:5005/api/v1/user
```
Request Body : <br>
The `account_id` is obtained from the property id when you register an account.
```
{
    "full_name" : "John Doe",
    "address" : "Jakarta Selatan",
    "account_id" : 2
}
```
Response Body :
```
{
    "status_code": 200,
    "status": "ok",
    "message": "success create new user",
    "data": {
        "id": 2,
        "full_name": "John Doe",
        "address": "Jakarta Selatan",
        "account": {
            "id": 2,
            "username": "john@gmail.com",
            "password": "$2a$10$J3laTGDqqK0z8Srn5q3RueaG3awFzvzMBpyHD6E0REuRa.CcZscI."
        }
    }
}
```

* **Login** <br>

Type this endpoint in Postman with **POST** Method.
```
http://localhost:5005/api/v1/login
```
Request Body :
```
{
    "username" : "john@gmail.com",
    "password" : "123456789012"
}
```
Response Body :
```
{
    "status_code": 200,
    "status": "ok",
    "message": "success login",
    "data": {
        "login_at": "2023-10-04 15:46:51",
        "token": "eyJhbGcxxx........."
    }
}
```

* **Get Users** <br>

Type this endpoint in Postman with **GET** Method.
```
http://localhost:5005/api/v1/users
```
Please insert the token you obtained during login into the Auth Bearer Token <br>

Response Body :
```
{
    "status_code": 200,
    "status": "ok",
    "message": "success get data",
    "data": {
        "account": {
            "id": 2,
            "username": "john@gmail.com"
        },
        "user": {
            "id": 1,
            "full_name": "John Doe",
            "address": "Jakarta Selatan"
        }
    }
}
```