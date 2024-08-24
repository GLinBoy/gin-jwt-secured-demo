# GIN JWT secured

## cURL requests

### Tags

- Get all tags

  ```bash
  curl --location 'localhost:8888/api/tag'
  ```

- Get a tag by ID

  ```bash
  curl --location 'localhost:8888/api/tag/1'
  ```

- Create a tag

  ```bash
  curl --location 'localhost:8888/api/tag' \
  --header 'Content-Type: application/json' \
  --data '{
      "name": "test"
    }'
  ```

- Update a tag

  ```bash
  curl --location --request PATCH 'localhost:8888/api/tag/1' \
  --header 'Content-Type: application/json' \
  --data '{
      "id": 1,
      "name": "updated tag"
      }'
  ```

- Delete a tag

  ```bash
  curl --location --request DELETE 'localhost:8888/api/tag/1'
  ```

### Users

- Signup

  ```bash
  curl --location 'localhost:8888/api/auth/signup' \
  --header 'Content-Type: application/json' \
  --data-raw '{
      "fullname": "John Doe",
      "email": "john.doe@email.com",
      "password": "123456"
      }'
  ```

- Signin

  ```bash
  curl --location 'localhost:8888/api/auth/signin' \
  --header 'Content-Type: application/json' \
  --data-raw '{
      "email": "john.doe@email.com",
      "password": "123456"
      }'
  ```
