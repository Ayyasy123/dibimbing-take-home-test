# Event Management System

This project is an Event Management System built using Go (Golang) with the Gin framework. It provides APIs for managing users, events, and tickets. The system allows users to register, login, create events, purchase tickets, and generate reports. It also includes role-based access control (RBAC) to ensure that only authorized users can perform certain actions.

## Technologies Used
- **Go (Golang)**: The primary programming language used for the backend.
- **Gin**: A web framework for building APIs in Go.
- **GORM**: An ORM (Object-Relational Mapping) library for interacting with the database.
- **JWT (JSON Web Tokens)**: Used for authentication and authorization.
- **Bcrypt**: Used for password hashing.
- **MySQL**: The database used for storing application data.

## API Documentation

### User Endpoints
- **POST /register**: Register a new user.
- **POST /login**: Login and receive a JWT token.
- **POST /register/admin**: Register a new admin.
- **GET /users**: Retrieve all users (admin only).
- **GET /users/:id**: Retrieve a user by ID.
- **PUT /users/:id**: Update a user.
- **DELETE /users/:id**: Delete a user (admin only).
- **GET /users/report**: Generate a user report (admin only).

### Event Endpoints
- **POST /events**: Create a new event (admin only).
- **GET /events**: Retrieve all events.
- **GET /events/:id**: Retrieve an event by ID.
- **PUT /events/:id**: Update an event (admin only).
- **DELETE /events/:id**: Delete an event (admin only).
- **PATCH /events/:id**: Cancel an event (admin only).
- **GET /events/search**: Search for events.
- **GET /events/report**: Generate an event report (admin only).

### Ticket Endpoints
- **POST /tickets**: Purchase a ticket.
- **GET /tickets**: Retrieve all tickets (admin only).
- **GET /tickets/:id**: Retrieve a ticket by ID.
- **PUT /tickets/:id**: Update a ticket (admin only).
- **DELETE /tickets/:id**: Delete a ticket (admin only).
- **GET /tickets/user**: Retrieve all tickets for the logged-in user.
- **PATCH /tickets/:id**: Cancel a ticket.
- **GET /tickets/report**: Generate a ticket sales report (admin only).
- **GET /tickets/report/event**: Retrieve tickets sold per event (admin only).
