# Flow

```mermaid
sequenceDiagram
    participant Client
    participant BE
    participant DB

    Client->>BE: POST /api/v1/users (register user)
    BE->>DB: Insert user data (is_activated = false)
    DB-->>BE: User created
    BE-->>Client: 200 OK + message (in demo, with activation link)

    Client->>Client: Click activation link (email)
    Client->>BE: GET /api/v1/users/verify?token=abcdefg123456
    BE->>DB: Update user.is_activated = true
    DB-->>BE: Update success
    BE-->>Client: 200 OK + success message
```

# APIs

## User Registration API

**POST** `/api/v1/users`

**Request Body**

```
{
    "email":"test92@gmail.com", //required, string
    "password":"testpswD@" //required, string
}
```

**Response**

`200` OK

```
{
    "version": "1.0.0",
    "error": "",
    "data": "" // would get activation link in this demo
}
```

## Get Activation API

**GET** `/api/v1/users/verify?token=abcdefg123456`

**Request Header**

```
Content-Type: application/json
```

**Response**
`200` OK

```
{
    "version": "1.0.0",
    "error": "",
    "data": ""
}
```
