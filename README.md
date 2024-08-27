# Loan Tracker API

## Overview

This Loan Tracker API is a backend project developed using Golang with the Gin framework. The API provides functionalities for both users and admins, allowing users to apply for loans and manage their accounts, while admins can manage user accounts and loan details. The project is built following clean architecture principles to ensure a modular, maintainable, and scalable codebase.

## Table of Contents

- [Overview](#overview)
- [Project Objectives](#project-objectives)
- [Project Structure](#project-structure)
- [Functional Requirements](#functional-requirements)
  - [User Management](#user-management)
  - [Admin Functionalities](#admin-functionalities)
- [Endpoints](#endpoints)
  - [User Management Endpoints](#user-management-endpoints)
  - [Admin Functionalities Endpoints](#admin-functionalities-endpoints)
- [Authentication and Authorization](#authentication-and-authorization)
- [Running the Project](#running-the-project)
- [Postman Documentation](#postman-documentation)

## Project Objectives

The main objectives of this project are to:

1. Develop a RESTful API using Golang and the Gin framework.
2. Implement user and admin functionalities.
3. Apply clean architecture principles.
4. Ensure secure and efficient handling of data.
5. Document the API using Postman.

## Project Structure

The project follows the clean architecture principles, separating the codebase into different layers such as `domain`, `usecase`, `repository`, and `delivery`. This structure ensures that the business logic is decoupled from external frameworks, making the code more maintainable and testable.

```plaintext
├── cmd/                   # Entry point of the application
├── domain/                # Entities and interfaces (business logic)
├── infrastructure/        # External services, database connection
├── repository/            # Data access logic
├── usecase/               # Application logic (use cases)
├── delivery/              # HTTP handlers (controllers)
├── config/                # Configuration files
├── main.go                # Main application file
└── README.md              # Project documentation
```

## Functional Requirements

### User Management

1.  **User Registration**

    - Endpoint: POST /users/register
    - Description: Register a new user with email, password, and profile details.
    - Response: Success or error message.

2.  **Email Verification**

    - Endpoint: GET /users/verify-email
    - Description: Verify the user's email address using a token sent to their email.
    - Parameters:

      - token: Verification token sent via email
      - email: User's email address

    - Flow of Events:

      1.  User registers on the platform.
      2.  System sends a verification email with a unique token.
      3.  User clicks on the verification link.
      4.  User is directed to the email verification endpoint.
      5.  The system verifies the token and activates the user's account.

    - Response: Success or error message.

3.  **User Login**

    - Endpoint: POST /users/login
    - Description: Authenticate user and provide access and refresh tokens.
    - Flow of Events:

      1.  User submits credentials.
      2.  System validates credentials.
      3.  System checks whether the user's account is verified.

    - Response: Access and refresh tokens or error message.

4.  **Token Refresh**

    - Endpoint: POST /users/token/refresh
    - Description: Refresh access token using refresh token.
    - Response: New access token or error message.

5.  **User Profile**

    - Endpoint: GET /users/profile
    - Description: Retrieve authenticated user profile.
    - Response: User profile data.

6.  **Password Reset Request**

    - Endpoint: POST /users/password-reset
    - Description: Send password reset link to user's email.
    - Response: Success or error message.

7.  **Password Update After Reset**

    - Endpoint: POST /users/password-update
    - Description: Update the user's password using the token received in the password reset email.
    - Flow of Events:

      1.  User receives the password reset link via email.
      2.  User clicks the link and is directed to a password reset page.
      3.  User submits the new password along with the token.
      4.  The system verifies the token and updates the password.

    - Response: Success or error message.

### Admin Functionalities

1.  **View All Users**

    - Endpoint: GET /admin/users
    - Description: Retrieve a list of all users.
    - Response: List of users.

2.  **Delete User Account**

    - Endpoint: DELETE /admin/users/{id}
    - Description: Delete a specific user account.
    - Response: Success or error message.

## Endpoints

### User Management Endpoints

- **POST** /users/register: Register a new user.
- **GET** /users/verify-email: Verify user's email.
- **POST** /users/login: Login a user.
- **POST** /users/token/refresh: Refresh access token.
- **GET** /users/profile: Retrieve user profile.
- **POST** /users/password-reset: Request password reset.
- **POST** /users/password-update: Update password after reset.

### Admin Functionalities Endpoints

- **GET** /admin/users: Retrieve all users.
- **DELETE** /admin/users/{id}: Delete a user account.

## Authentication and Authorization

The API uses JWT (JSON Web Tokens) for authenticating and authorizing users. Access tokens and refresh tokens are used to secure endpoints and manage user sessions.

- **Access Token**: Short-lived token used to access protected resources.
- **Refresh Token**: Long-lived token used to obtain a new access token.

### 2.2 Loan Management

#### Apply for Loan

- **Endpoint:** `POST /loans`
- **Description:** Submit a loan application.
- **Response:** Loan application status.

#### View Loan Status

- **Endpoint:** `GET /loans/{id}`
- **Description:** Retrieve the status of a specific loan of the current user.
- **Response:** Loan status details.

#### View All Loans (Admin)

- **Endpoint:** `GET /admin/loans`
- **Description:** Retrieve all loan applications.
- **Response:** List of loan applications.
- **Parameters:**
  - `status`: `pending` | `approved` | `rejected` (optional, default: all)
  - `order`: `asc` | `desc` (optional, default: `asc` for pending, `desc` for reviewed)

## Postman Documentation

A Postman collection has been provided for testing the API endpoints. You can import the collection into Postman and use it to interact with the API.

- Postman Link [Postman Documentation](https://documenter.getpostman.com/view/34226868/2sAXjGduTq)

## License

This project is licensed under the MIT License. See the LICENSE file for more details.

Feel free to contribute to this project by opening issues or submitting pull requests.
