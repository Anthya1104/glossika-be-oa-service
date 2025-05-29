## User Registration API

**POST** `/api/v1/users`

- 400 if invalid email pattern
- 400 if invalid password pattern
- 424 if send email failed
- 500 if bcrypt hash password failed
- 500 if DB create user faild（except duplicate entry）
- 500 if JWT token generation failed

## User Login API

**POST** `/api/v1/auth/login`

- 200 if user login success
- 400 if user login with missing req body or invalid JSON
- 400 if user not found by email
- 401 if password mismatch
- 403 if user not activated
- 500 if JWT token generation failed

## User Activation API

**GET** `/api/v1/users/verify?token=abcdefg123456`

- 200 if user activation succeeds with a valid token
- 400 if the token is missing in the query
- 400 if the token is invalid or expired
- 400 if the token type is not "email_verify"
- 500 if updating the user activation status in the database fails

## Get Recommendation API

**POST** `/api/v1/recommendations`

- 200 if user gets recommendations successfully (with cache hit)
- 200 if user gets recommendations successfully (with cache miss, DB fetch)
- 400 if request body is missing or invalid (JSON binding error)
- 500 if fetching user recommendations from DB fails
- 500 if fetching products from DB fails
- 200 if user has no recommendations (empty list returned)
