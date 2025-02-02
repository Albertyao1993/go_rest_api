# Go REST API
Demo project for RESTful API implementation using Go

## Core Features

### User Management
- User registration and authentication
- Secure password hashing
- JWT token-based authentication
- Login/Logout functionality

### Event Management
- Event registration
- Event cancellation
- Event listing and details
- User-event association

### Database Structure
```sql
-- Users Table
CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Events Table
CREATE TABLE events (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    event_date DATETIME NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- User-Event Registration Table
CREATE TABLE event_registrations (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    event_id BIGINT NOT NULL,
    registration_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (event_id) REFERENCES events(id),
    UNIQUE KEY unique_registration (user_id, event_id)
);
```

### Authentication Middleware
- JWT token validation
- Protected route access control
- Token expiration handling
- Role-based authorization (optional)

### API Endpoints
- `POST /api/register` - User registration
- `POST /api/login` - User login
- `GET /api/events` - List all events
- `POST /api/events` - Create new event
- `POST /api/events/:id/register` - Register for event
- `DELETE /api/events/:id/register` - Cancel event registration
- `GET /api/user/events` - List user's registered events

### Implementation Details
- RESTful API design principles
- CRUD operations
- SQL database integration
- Secure authentication flow
- Error handling and validation