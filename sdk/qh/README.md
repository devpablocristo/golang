# QH MVP

Welcome to the **QH MVP**! This project aims to create a minimal viable product (MVP) for a social network focused on events, enabling users and companies to interact, create, participate in events, form groups, and provide ratings and feedback.

## Table of Contents

- [Project Overview](#project-overview)
- [Architecture](#architecture)
- [Microservices](#microservices)
  - [Authentication Service](#authentication-service)
  - [Authorization Service](#authorization-service)
  - [User Service](#user-service)
  - [Event Service](#event-service)
  - [Group Service](#group-service)
  - [Rating Service](#rating-service)
  - [Notification Service](#notification-service)
- [Technologies](#technologies)
- [Getting Started](#getting-started)
- [Future Enhancements](#future-enhancements)
- [Contributing](#contributing)
- [License](#license)

## Project Overview

The **Event Social Network MVP** is designed to facilitate interactions between individual users and companies around events. Key functionalities include:

- **User Authentication and Authorization**: Secure login and role-based access control.
- **User and Company Profiles**: Manage personal and business profiles.
- **Event Management**: Create, modify, and view events.
- **Group Formation**: Allow users to form groups to attend events together.
- **Ratings and Reviews**: Users can rate and comment on events and companies.
- **Notifications**: Keep users informed about relevant updates and activities.

## Architecture

The project follows a **Microservices Architecture**, where each functionality is encapsulated within its own service. This approach ensures scalability, maintainability, and the ability to develop and deploy services independently.

![Architecture Diagram](architecture-diagram.png)

*Note: Replace the placeholder with an actual architecture diagram as needed.*

## Microservices

### Authentication Service

**Name:** `Authe`

**Function:**  
Handles user registration, login, verification of credentials, and the generation and validation of authentication tokens (JWT, OAuth).

**Key Features:**
- User registration and login for both individuals and companies.
- Token-based authentication for secure access.
- Integration with third-party authentication providers (e.g., Google, Facebook).

**Database:** PostgreSQL or MongoDB

### Authorization Service

**Name:** `Autho`

**Function:**  
Manages user roles and permissions, implementing role-based access control (RBAC) to ensure users have appropriate access levels.

**Key Features:**
- Define and assign roles (e.g., admin, user).
- Manage permissions for various actions within the platform.
- Validate user privileges for accessing specific resources or functionalities.

**Database:** Redis or similar caching system

### User Service

**Name:** `Users`

**Function:**  
Manages user profiles for both individuals and companies, including creation, modification, deletion, and retrieval of user data.

**Key Features:**
- Separate handling of individual and company profiles.
- Manage user-specific attributes (e.g., personal details, company information).
- Support for future scalability and complexity enhancements.

**Database:** PostgreSQL or MongoDB

### Event Service

**Name:** `Events`

**Function:**  
Facilitates the creation, modification, and visualization of events, including managing posts related to these events.

**Key Features:**
- Create and manage event details (date, location, category).
- Allow companies to publish and promote events.
- Integrate post management for event-related discussions and announcements.

**Database:** PostgreSQL or Event Store

### Group Service

**Name:** `Group`

**Function:**  
Enables users to form groups for attending events together, fostering social interaction and collaboration.

**Key Features:**
- Create and manage user groups.
- Associate groups with specific events.
- Define roles within groups (e.g., group admin, member).

**Database:** PostgreSQL or MongoDB

### Rating Service

**Name:** `Rating`

**Function:**  
Provides a system for users to rate and comment on events and companies, enhancing transparency and trust within the platform.

**Key Features:**
- Users can leave ratings (e.g., star ratings) and comments.
- Aggregate ratings for events and companies.
- Moderate and manage inappropriate feedback.

**Database:** MongoDB

### Notification Service

**Name:** `Notification`

**Function:**  
Handles the delivery of notifications to users about important activities, such as new events, group invitations, and received ratings.

**Key Features:**
- Send notifications via email, push, or other channels.
- Notify users about relevant updates and interactions.
- Manage notification preferences and delivery schedules.

**Database:** Redis or RabbitMQ

## Technologies

- **Programming Language:** Go (Golang)
- **Frameworks:** Gin-Gonic for API development
- **Databases:** PostgreSQL, MongoDB, Redis, Cassandra (for messaging)
- **Messaging & Queues:** RabbitMQ, Kafka (optional)
- **Authentication:** JWT, OAuth
- **Containerization:** Docker, Docker Compose
- **Orchestration:** Kubernetes (optional for scaling)
- **Monitoring & Logging:** Prometheus, Grafana, ELK Stack
- **Version Control:** Git
- **CI/CD:** GitHub Actions, Jenkins (optional)
- **Other Tools:** Viper for configuration management, Go-Micro for microservices framework

## Getting Started

### Prerequisites

- **Go** (version 1.22 or later)
- **Docker** and **Docker Compose**
- **Git**
- **Node.js** and **npm** (if frontend is included)
- **Database Systems**: PostgreSQL, MongoDB, Redis

### Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/yourusername/event-social-network-mvp.git
   cd event-social-network-mvp
   ```

2. **Set Up Environment Variables**

   Create a `.env` file in the root directory and configure the necessary environment variables for each microservice.

3. **Run Microservices with Docker Compose**

   ```bash
   docker-compose up -d
   ```

4. **Access the Services**

   Each microservice will be accessible on its configured port. Refer to the `docker-compose.yml` file for port mappings.

### Development

1. **Navigate to a Microservice Directory**

   ```bash
   cd services/authentication
   ```

2. **Run Locally**

   ```bash
   go run main.go
   ```

3. **Testing**

   Write and run tests for each microservice to ensure functionality.

   ```bash
   go test ./...
   ```

### API Documentation

Each microservice should have its own API documentation, possibly using tools like Swagger or Postman. Ensure that APIs are well-documented for easy integration and maintenance.

## Future Enhancements

While the MVP focuses on essential functionalities, future iterations can include:

- **Separate User Services**: Distinguish between individual and company user services for better scalability and maintainability.
- **Chat Service**: Real-time messaging between users and within groups.
- **Payment Service**: Handle payments for paid events or premium features.
- **Analytics Service**: Collect and analyze user interaction data for insights.
- **Media Service**: Manage multimedia content uploads and storage.
- **Search Service**: Implement advanced search capabilities using ElasticSearch.

## Contributing

Contributions are welcome! Please follow these steps:

1. **Fork the Repository**
2. **Create a Feature Branch**

   ```bash
   git checkout -b feature/new-feature
   ```

3. **Commit Your Changes**

   ```bash
   git commit -m "Add new feature"
   ```

4. **Push to the Branch**

   ```bash
   git push origin feature/new-feature
   ```

5. **Open a Pull Request**

## License

This project is licensed under the [MIT License](LICENSE).
