# Event RSVP System – Event & Guest Response Management

A full-stack web application built to simplify event management and guest response tracking.  
The Event RSVP System allows organizers to create events, invite guests, and track RSVP responses (Yes / No / Maybe) through a clean and structured workflow.

---

## Project Overview

The Event RSVP System is designed for scenarios such as parties, meetings, weddings, and corporate events where organizers need an easy way to manage invitations and monitor guest attendance.

The application follows a layered backend architecture and is deployed using containerized services to ensure scalability, maintainability, and clean separation of concerns.

---

## Key Features

### Event Management
- Create and manage events
- Define event details such as title and date
- View all created events

### Guest Management
- Add and manage guest details
- Maintain unique guest records
- Associate guests with events

### RSVP Tracking
- Guests can respond with **Yes / No **
- Track RSVP status per event
- Maintain RSVP history

### Architecture & Design
- Clean layered architecture (Handler → Service → Repository)
- DTO-based request validation
- RESTful API design
- Clear separation of frontend and backend

---

## Tech Stack

### Backend
- Go (Golang)
- Gin (HTTP framework)
- GORM (ORM)
- PostgreSQL

### Frontend
- React
- Axios

### DevOps & Deployment
- Docker
- Docker Compose
- Nginx (Reverse Proxy)

---

## Getting Started

### Prerequisites
- Docker
- Docker Compose
- PostgreSQL (for local, non-docker usage)

---

## How to Run

### 1.Clone the Project
```bash
git clone https://github.com/SajadCT/rsvp-system.git
cd rsvp-system
```

## 2.Environment Configuration

The application uses environment variables for database configuration.

### Create `.env` File

Create a `.env` file in the backend root directory and add the following:

```.env
DB_PORT=5432
DB_USER=rsvp
DB_PASSWORD=your_password
DB_NAME=rsvp_db
DB_SSLMODE=disable
```
## Running the Project with Docker

This project is fully containerized and can be started using **Docker Compose**.

```bash
docker compose up --build -d
```

