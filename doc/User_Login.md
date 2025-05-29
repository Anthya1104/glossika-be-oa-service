# APIs

## User Login API

**POST** `/api/v1/auth/login`

**Request Body**

```
{
    "email":"test2@gmail.com", // required, string
    "password":"testpswD@" // required, string
}
```

**Response**

`200` OK

```
{
    "version": "1.0.0",
    "error": "",
    "data": {
        "token": "auth-token"
    }
}
```
