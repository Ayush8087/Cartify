# Cartify - E-Commerce Microservices System

Cartify is an e-commerce system built using a microservices architecture. Each service is responsible for a specific functionality, such as user management, product catalog, order processing, payment handling, and notifications. The services are implemented in Go and use the Gorilla Mux router for HTTP routing.

## Features

- **User Management**: Register, login, and manage user profiles.
- **Product Catalog**: Manage products, including details and pricing.
- **Order Management**: Create, update, and track orders.
- **Payment Processing**: Handle payments and refunds.
- **Notifications**: Send email/SMS notifications to users.

## Technologies Used

- **Go**: Programming language used for building the microservices.
- **Gorilla Mux**: Web framework for handling HTTP requests and routing.
- **JSON**: Used for data serialization and communication between services.
- **HTTP**: Communication protocol for inter-service communication.

## Installation and Setup

### Prerequisites
- Go (version 1.16 or above)
- Git

### Steps to Run the Application
1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/cartify.git
    cd cartify
    ```

2. Navigate to each service directory and run the service:
    ```bash
    cd user-service
    go run main.go
    ```

    Repeat the above step for `product-service`, `order-service`, `payment-service`, and `notification-service`.

3. The services will be available at the following ports:
    - **User Service**: `http://localhost:8081`
    - **Product Service**: `http://localhost:8082`
    - **Order Service**: `http://localhost:8083`
    - **Payment Service**: `http://localhost:8084`
    - **Notification Service**: `http://localhost:8085`

## API Endpoints and Sample Responses

### User Service

#### Register a User
**POST** `/register`

**Request Body**:
```json
{
  "id": "1",
  "name": "Test User",
  "email": "test@example.com",
  "password": "password123"
}
```

**Response**:
```json
{
  "id": "1",
  "name": "Test User",
  "email": "test@example.com",
  "password": "password123"
}
```

**POST** `/login`

**Request Body**:
```json
{
  "email": "test@example.com",
  "password": "password123"
}
```

**Response**:
```json
{
  "id": "1",
  "name": "Test User",
  "email": "test@example.com",
  "password": "password123"
}
```

### Product Service

#### Get All Products
**GET** `/products`

**Response**:
```json
[
  {
    "id": "1",
    "name": "Laptop",
    "price": "1000"
  },
  {
    "id": "2",
    "name": "Smartphone",
    "price": "500"
  }
]
```

### Order Service

#### Create an Order
**POST** `/orders`

**Request Body**:
```json
{
  "id": "1",
  "userId": "1",
  "productId": "101",
  "quantity": 2
}
```

**Response**:
```json
{
  "id": "1",
  "userId": "1",
  "productId": "101",
  "quantity": 2
}
```

#### List All Orders
**GET** `/orders`

**Response**:
```json
[
  {
    "id": "1",
    "userId": "1",
    "productId": "101",
    "quantity": 2
  }
]
```

### Payment Service

#### Create a Payment
**POST** `/payments`

**Request Body**:
```json
{
  "id": "1",
  "userId": "1",
  "amount": 2000
}
```

**Response**:
```json
{
  "id": "1",
  "userId": "1",
  "amount": 2000
}
```

#### List All Payments
**GET** `/payments`

**Response**:
```json
[
  {
    "id": "1",
    "userId": "1",
    "amount": 2000
  }
]
```

### Notification Service

#### Get All Notifications
**GET** `/notifications`

**Response**:
```json
[
  {
    "id": "1",
    "userId": "123",
    "message": "Your order has been shipped!"
  }
]
```
