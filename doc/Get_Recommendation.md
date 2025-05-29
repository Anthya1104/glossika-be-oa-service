# APIs

## User Login API

**POST** `/api/v1/recommendations`

**Reqeust Header**

```
Content-Type: application/json
Authorization: "auth-token" // get auth token from POST /api/v1/auth/login
```

**Request Body**

```
{
    "page":1, // required, int
    "pageSize":50 //required, int
}
```

**Response**

`200` OK

```
{
    "version": "1.0.0",
    "error": "",
    "data": {
        "total": 3,
        "recommendation_list": [
            {
                "product_id": 1,
                "product_name": "National",
                "description": "I hand shake.",
                "price": 456.93
            },
            {
                "product_id": 2,
                "product_name": "Same",
                "description": "Fish keep structure.",
                "price": 862.97
            },
            {
                "product_id": 4,
                "product_name": "Issue",
                "description": "Side care couple kid.",
                "price": 582.68
            }
            ]
    }
}

```
