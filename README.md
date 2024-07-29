# Neosurge Portfolio Management Service

## Description
A simplified portfolio management service that manages user portfolios, tracks asset values, and provides basic analytics.

## Setup and Run
1. Clone the repository:
    ```sh
    git clone https://github.com/AVtheking/user_portfolio_management.git
    cd user_portfolio_management
    ```

2. Create a `.env` file with the following variables:
    ```
     PORT=
    DB_USER=""
    DB_PASS=""
    DB_HOST=""
    DB_PORT=
    DB_NAME=""
    JWT_ACCESS_SECRET=""
    JWT_REFRESH_SECRET=""
    ```

3. Run the application:
    ```sh
    go run main.go
    ```

## API Endpoints

### Auth

- **Sign Up**
  - **URL**: `/api/v1/auth/signup`
  - **Method**: `POST`
  - **Description**: Registers a new user.
  - **Request Body**:
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```

- **Login**
  - **URL**: `/api/v1/auth/login`
  - **Method**: `POST`
  - **Description**: Authenticates a user and returns a JWT token.
  - **Request Body**:
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```

### User

- **Get User Details**
  - **URL**: `/api/v1/user/`
  - **Method**: `GET`
  - **Description**: Retrieves the authenticated user's details.
  - **Headers**:
    - `Authorization: Bearer <JWT_TOKEN>`

### Portfolio

- **Create Portfolio**
  - **URL**: `/api/v1/portfolio/`
  - **Method**: `POST`
  - **Description**: Creates a new portfolio.
  - **Headers**:
    - `Authorization: Bearer <JWT_TOKEN>`
  - **Request Body**:
    ```json
    {
      "name": "string"
    }
    ```

- **Update Portfolio**
  - **URL**: `/api/v1/portfolio/:id`
  - **Method**: `PUT`
  - **Description**: Updates an existing portfolio.
  - **Headers**:
    - `Authorization: Bearer <JWT_TOKEN>`
  - **Request Body**:
    ```json
    {
      "name": "string"
    }
    ```

- **Get Portfolio**
  - **URL**: `/api/v1/portfolio/:id`
  - **Method**: `GET`
  - **Description**: Retrieves a portfolio by ID.
  - **Headers**:
    - `Authorization: Bearer <JWT_TOKEN>`

- **Delete Portfolio**
  - **URL**: `/api/v1/portfolio/:id`
  - **Method**: `DELETE`
  - **Description**: Deletes a portfolio by ID.
  - **Headers**:
    - `Authorization: Bearer <JWT_TOKEN>`

### Asset

- **Add Asset**
  - **URL**: `/api/v1/asset/:portfolioId`
  - **Method**: `POST`
  - **Description**: Adds a new asset to a portfolio.
  - **Headers**:
    - `Authorization: Bearer <JWT_TOKEN>`
  - **Request Body**:
    ```json
    {
      "name": "string",
      "value": "number"
    }
    ```

- **Update Asset**
  - **URL**: `/api/v1/asset/:assetId`
  - **Method**: `PUT`
  - **Description**: Updates an existing asset.
  - **Headers**:
    - `Authorization: Bearer <JWT_TOKEN>`
  - **Request Body**:
    ```json
    {
      "name": "string",
      "value": "number"
    }
    ```

- **Delete Asset**
  - **URL**: `/api/v1/asset/:assetId`
  - **Method**: `DELETE`
  - **Description**: Deletes an asset by ID.
  - **Headers**:
    - `Authorization: Bearer <JWT_TOKEN>`

- **Get Asset**
  - **URL**: `/api/v1/asset/:assetId`
  - **Method**: `GET`
  - **Description**: Retrieves an asset by ID.
  - **Headers**:
    - `Authorization: Bearer <JWT_TOKEN>`

- **Get All Assets in a Portfolio**
  - **URL**: `/api/v1/asset/portfolio/:portfolioId`
  - **Method**: `GET`
  - **Description**: Retrieves all assets in a specific portfolio.
  - **Headers**:
    - `Authorization: Bearer <JWT_TOKEN>`

### Analytics

- **Get Total Portfolio Value**
  - **URL**: `/api/v1/analytics/totalValue/:portfolioId`
  - **Method**: `GET`
  - **Description**: Calculates the total value of a portfolio.
  - **Headers**:
    - `Authorization: Bearer <JWT_TOKEN>`

- **Get Average Return**
  - **URL**: `/api/v1/analytics/averageReturn/:portfolioId`
  - **Method**: `GET`
  - **Description**: Calculates the average return of a portfolio.
  - **Headers**:
    - `Authorization: Bearer <JWT_TOKEN>`

## Report
### Design Choices
- **Frameworks and Libraries**: Gin for HTTP routing, Gorm for ORM, and JWT-go for JWT handling were chosen for their simplicity and effectiveness in building web services.
- **Code Structure**: The project is structured to separate concerns between handlers, models, middleware, and database initialization.

### Assumptions and Future Extensions
- **Assumptions**: User authentication is performed using JWT, and assets are manually tracked by users.
- **Future Extensions**: Real-time asset tracking, more advanced analytics, and user roles and permissions can be added to extend the functionality.
