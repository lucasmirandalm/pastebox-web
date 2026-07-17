# Pastebox Web

<p align="center">
  <strong>A clean paste management web application built with Go.</strong>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-Backend-00ADD8?style=for-the-badge&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/chi-Router-111111?style=for-the-badge" alt="chi router">
  <img src="https://img.shields.io/badge/PostgreSQL-Database-4169E1?style=for-the-badge&logo=postgresql" alt="PostgreSQL">
  <img src="https://img.shields.io/badge/Tailwind%20CSS-UI-38B2AC?style=for-the-badge&logo=tailwindcss" alt="Tailwind CSS">
  <img src="https://img.shields.io/badge/Docker-Environment-2496ED?style=for-the-badge&logo=docker" alt="Docker">
</p>

---

## About

**Pastebox Web** is a server-side web application for creating, organizing, editing, searching, and favoriting text pastes.

The project is built with Go, PostgreSQL, server-rendered HTML templates, Tailwind CSS, and a layered backend architecture.

This project is being developed as a learning-focused backend application with attention to:

- clean project organization;
- HTTP routing;
- HTML rendering;
- PostgreSQL integration;
- migrations;
- authentication;
- authorization;
- ownership rules;
- form handling;
- validation;
- sessions;
- testing;
- production-minded structure.

---

## Tech Stack

- **Go** — main backend language
- **chi** — HTTP router
- **html/template** — server-side HTML rendering
- **PostgreSQL** — relational database
- **golang-migrate** — database migrations
- **Docker Compose** — local development database
- **Tailwind CSS CDN** — interface styling during development

---

## Planned Features

- Paste listing
- Paste creation
- Paste editing
- Paste deletion
- Favorite and unfavorite actions
- Text search
- Favorite filtering
- Pagination
- User registration
- User login and logout
- Session-based authentication
- Per-user paste ownership
- Protected private dashboard
- Permission-based paste access
- Public read-only paste links
- Public link regeneration
- Automated tests

---

## Design System

Pastebox Web uses a clean, minimal, black-and-white visual style.

| Token           | Color              | Usage                              |
| --------------- | ------------------ | ---------------------------------- |
| Background      | `#ffffff`          | Main page background               |
| Surface         | neutral gray tones | Cards, inputs, and secondary areas |
| Primary         | `#000000`          | Buttons, active states, emphasis   |
| Text            | neutral black      | Main text                          |
| Muted text      | neutral gray       | Secondary text                     |

The interface uses a centered `900px` container to keep the layout readable and consistent.

Typography:

- **Sora** for headings, logo, and important UI labels
- **Inter** for body text and general interface content

---

## Architecture

The application follows a layered structure:

```txt
Browser
  ↓
chi router
  ↓
HTTP handlers
  ↓
services
  ↓
repositories
  ↓
PostgreSQL
```

The goal is to keep responsibilities separated:

- **handlers** deal with HTTP requests, forms, redirects, and responses;
- **services** handle business rules;
- **repositories** handle database access and SQL;
- **templates** render HTML pages;
- **config** centralizes environment-based configuration;
- **database** manages the PostgreSQL connection.

---

## Project Structure

```txt
pastebox-web/
├── cmd/
│   └── web/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── database/
│   │   └── database.go
│   ├── paste/
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repository.go
│   │   ├── model.go
│   │   └── errors.go
│   └── render/
│       └── render.go
├── migrations/
├── ui/
│   └── templates/
│       ├── layouts/
│       │   └── base.html
│       └── pages/
│           ├── home.html
│           ├── new.html
│           └── edit.html
├── docker-compose.yml
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

---

## Requirements

- Go
- Docker
- Docker Compose
- PostgreSQL client tools
- golang-migrate

---

## Installation

Clone the repository:

```bash
git clone git@github.com:lucasmirandalm/pastebox-web.git
```

Enter the project folder:

```bash
cd pastebox-web
```

Install Go dependencies:

```bash
go mod download
```

Run the application:

```bash
go run ./cmd/web
```

Open Pastebox Web:

```txt
http://localhost:8080
```

---

## Environment Variables

The application will use a `.env` file during development.

Example:

```env
PORT=8080
DATABASE_URL=postgres://pastebox_user:pastebox_password@127.0.0.1:5435/pastebox_web?sslmode=disable
```

Variables:

| Name           | Description                  |
| -------------- | ---------------------------- |
| `PORT`         | HTTP server port             |
| `DATABASE_URL` | PostgreSQL connection string |

---

## Database

During development, PostgreSQL will run through Docker Compose.

Planned local database setup:

```bash
docker compose up -d
```

Run migrations:

```bash
migrate -path migrations -database "$DATABASE_URL" up
```

Rollback migrations:

```bash
migrate -path migrations -database "$DATABASE_URL" down
```

---

## Main Routes

Planned routes:

```txt
GET  /                         paste library
GET  /pastes/new               new paste form
POST /pastes                   create paste
GET  /pastes/{id}/edit         edit paste form
POST /pastes/{id}              update paste
POST /pastes/{id}/favorite     toggle favorite
POST /pastes/{id}/delete       delete paste
GET  /login                    login form
POST /login                    login user
POST /logout                   logout user
GET  /register                 registration form
POST /register                 create account
GET  /p/{public_id}            public read-only paste
```

---

## Learning Goals

This project is being built to practice professional backend development with Go, including:

- HTTP routing with chi
- server-side rendering with `html/template`
- PostgreSQL integration
- SQL migrations
- form handling
- validation
- session-based authentication
- authorization and ownership rules
- layered architecture
- dependency injection
- error handling
- testing handlers, services, and repositories
- building a real web application from start to finish

---

## Project Status

Pastebox Web is currently **in development**.

The project is being built step by step, with a focus on understanding each part of the backend architecture before adding the next layer.

---

## License

This project is open source and available under the MIT License.

---

<p align="center">
  Built with Go, PostgreSQL, Tailwind CSS, Docker, and a lot of practice.
</p>
