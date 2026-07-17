# Pastebox

<p align="center">
  <strong>A clean paste-sharing web app built with pure Django.</strong>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Django-Pure%20Django-092E20?style=for-the-badge&logo=django" alt="Django">
  <img src="https://img.shields.io/badge/Tailwind%20CSS-UI-38B2AC?style=for-the-badge&logo=tailwindcss" alt="Tailwind CSS">
  <img src="https://img.shields.io/badge/PostgreSQL-Database-4169E1?style=for-the-badge&logo=postgresql" alt="PostgreSQL">
  <img src="https://img.shields.io/badge/Docker-Environment-2496ED?style=for-the-badge&logo=docker" alt="Docker">
  <img src="https://img.shields.io/badge/Tests-12%20Passing-brightgreen?style=for-the-badge" alt="12 passing tests">
</p>

---

<table align="center">
  <tr>
    <td align="center">
      <img src="github_assets/pasteboxdesktop.gif" alt="Pastebox desktop demo" width="520">
      <br>
      <strong>Desktop</strong>
    </td>
    <td align="center">
      <img src="github_assets/pasteboxmobile.gif" alt="Pastebox mobile demo" width="220">
      <br>
      <strong>Mobile</strong>
    </td>
  </tr>
</table>

---

## About

**Pastebox** is a paste-sharing web application built with pure Django, Tailwind CSS, PostgreSQL, and Docker.

Users can create private accounts, manage their own text pastes, search and filter their workspace, mark pastes as favorites, and share read-only public links.

Pastebox was developed as a learning-focused Django project with attention to project organization, authentication, permissions, business rules, responsive design, and automated tests.

---

## Features

* Pure Django web application
* PostgreSQL database
* Docker-based development environment
* Tailwind CSS interface
* Custom visual design system
* Responsive desktop and mobile layouts
* User registration
* User login and logout
* Session-based authentication
* Protected private paste dashboard
* Per-user paste ownership
* Paste creation
* Paste listing
* Paste editing
* Paste deletion
* Favorite and unfavorite actions
* Text search
* Favorite filtering
* Pagination
* Public shareable paste links
* Public read-only paste pages
* Copy-to-clipboard public links
* Public slug regeneration after paste updates
* Old public link invalidation after paste updates
* Permission-based paste access
* Django messages for user feedback
* Automatically dismissed feedback messages
* Django Admin integration
* Automated model and view tests

---

## Design System

Pastebox uses a clean and minimal visual style.

| Token           | Color     | Usage                                      |
| --------------- | --------- | ------------------------------------------ |
| Main background | `#ffffff` | Page background                            |
| Surface         | `#eeeeee` | Cards, inputs, and secondary areas         |
| Accent          | `#ffbb11` | Hover states, focus states, and highlights |
| Text            | `#2f2f2f` | Main text                                  |
| Muted text      | `#6b6b6b` | Secondary text                             |

The interface uses a centered `900px` container to keep the layout readable and consistent.

---

## Project Structure

```txt
pastebox/
├── github_assets/
│   ├── pasteboxdesktop.gif
│   └── pasteboxmobile.gif
├── accounts/
│   ├── migrations/
│   ├── admin.py
│   ├── apps.py
│   ├── models.py
│   ├── tests.py
│   ├── urls.py
│   └── views.py
├── config/
│   ├── settings.py
│   ├── urls.py
│   ├── asgi.py
│   └── wsgi.py
├── pastes/
│   ├── migrations/
│   ├── admin.py
│   ├── apps.py
│   ├── forms.py
│   ├── models.py
│   ├── tests.py
│   ├── urls.py
│   └── views.py
├── static/
│   ├── src/
│   │   └── input.css
│   └── css/
│       └── output.css
├── templates/
│   ├── base.html
│   ├── partials/
│   │   ├── messages.html
│   │   └── navbar.html
│   ├── accounts/
│   │   ├── login.html
│   │   └── register.html
│   └── pastes/
│       ├── paste_form.html
│       ├── paste_list.html
│       ├── paste_public_detail.html
│       └── partials/
│           ├── pagination.html
│           ├── paste_card.html
│           └── paste_filter_form.html
├── compose.yml
├── Dockerfile
├── manage.py
├── package.json
├── package-lock.json
├── requirements.txt
├── .env.example
├── .gitattributes
├── .gitignore
└── README.md
```

---

## Requirements

* Docker
* Docker Compose

The application runs inside Docker containers, so Python, PostgreSQL, and Node.js do not need to be installed directly on the host machine.

---

## Installation

Clone the repository:

```bash
git clone git@github.com:lucasmirandalm/pastebox.git
```

Enter the project folder:

```bash
cd pastebox
```

Create the local environment file:

```bash
cp .env.example .env
```

Build the containers:

```bash
docker compose build
```

Start the application:

```bash
docker compose up
```

Open Pastebox:

```txt
http://localhost:8000
```

---

## Environment Variables

The `.env.example` file contains the variables required by the application.

```env
POSTGRES_DB=pastebox
POSTGRES_USER=pastebox_user
POSTGRES_PASSWORD=pastebox_password
POSTGRES_HOST=db
POSTGRES_PORT=5432
POSTGRES_EXTERNAL_PORT=5433
```

`POSTGRES_PORT` is used for communication between the Django and PostgreSQL containers.

`POSTGRES_EXTERNAL_PORT` exposes PostgreSQL to the host machine and can be changed when port `5432` is already in use.

---

## Database Setup

Run the migrations:

```bash
docker compose run --rm web python manage.py migrate
```

Create a superuser:

```bash
docker compose run --rm web python manage.py createsuperuser
```

Access the Django Admin:

```txt
http://localhost:8000/admin/
```

---

## Tailwind CSS

Pastebox uses Tailwind CSS through a dedicated Docker service.

The Tailwind source file is located at:

```txt
static/src/input.css
```

The generated CSS file is located at:

```txt
static/css/output.css
```

During development, the Tailwind container runs in watch mode and recompiles the stylesheet when templates or source styles change.

---

## Usage

Start all application services:

```bash
docker compose up
```

Useful URLs:

```txt
http://localhost:8000/
http://localhost:8000/register/
http://localhost:8000/login/
http://localhost:8000/admin/
```

The authenticated dashboard includes:

* text search
* favorite filtering
* paste listing
* pagination
* paste creation
* paste editing
* paste deletion
* favorite and unfavorite actions
* public link sharing

---

## Authentication and Permissions

Pastebox uses Django's built-in authentication and session system.

Private views require authentication, and all private paste queries are restricted by the current user:

```python
Paste.objects.filter(owner=request.user)
```

Update, favorite, and delete operations also verify ownership before changing a paste.

This prevents users from viewing or modifying another user's private pastes.

---

## Public Links

Each paste receives a unique public UUID.

Public links can be opened without authentication and display a read-only version of the paste.

When a paste is edited and saved, Pastebox generates a new public slug. As a result:

* the previous public link becomes invalid;
* the newest public link becomes active;
* favorite and unfavorite actions do not change the public link.

---

## Running Tests

Pastebox includes **12 essential automated tests** covering its main models, permissions, business rules, and views.

Run the complete test suite:

```bash
docker compose run --rm web python manage.py test
```

Run only the tests from the `pastes` app:

```bash
docker compose run --rm web python manage.py test pastes
```

Run the tests with detailed output:

```bash
docker compose run --rm web python manage.py test pastes -v 2
```

If Docker requires elevated permissions on your machine:

```bash
sudo docker compose run --rm web python manage.py test
```

### Test coverage

The test suite verifies that:

* the string representation of a paste returns its title;
* a public slug is generated automatically;
* users only see their own private pastes;
* authenticated users can create pastes;
* unauthenticated users cannot create pastes;
* paste owners can update their own pastes;
* users cannot update another user's paste;
* updating a paste generates a new public slug;
* favoriting changes only the favorite status;
* paste owners can delete their own pastes;
* public paste pages are accessible without login;
* text search and favorite filtering work together.

The Django test runner creates a temporary test database and destroys it after the suite finishes. Tests do not modify the development database.

---

## Why I Built This

I built Pastebox as a Django learning project to practice:

* pure Django development
* PostgreSQL integration
* Docker-based environments
* Tailwind CSS styling
* authentication and authorization
* session-based login
* ownership and permissions
* Django models and forms
* class organization
* reusable template partials
* Django messages
* URL routing
* public UUID links
* business rules
* responsive interface design
* automated model and view testing
* debugging real application errors

---

## Project Status

Pastebox is complete as a learning project.

The implemented scope includes the main application workflow, authentication, paste management, public sharing, responsive design, permissions, and essential automated tests.

Future Django concepts and additional features will be explored in separate projects.

---

## License

This project is open source and available under the MIT License.

---

<p align="center">
  Built with Django, Tailwind CSS, PostgreSQL, Docker, and a lot of practice.
</p>
# Pastebox

<p align="center">
  <strong>A clean paste-sharing web app built with pure Django.</strong>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Django-Pure%20Django-092E20?style=for-the-badge&logo=django" alt="Django">
  <img src="https://img.shields.io/badge/Tailwind%20CSS-UI-38B2AC?style=for-the-badge&logo=tailwindcss" alt="Tailwind CSS">
  <img src="https://img.shields.io/badge/PostgreSQL-Database-4169E1?style=for-the-badge&logo=postgresql" alt="PostgreSQL">
  <img src="https://img.shields.io/badge/Docker-Environment-2496ED?style=for-the-badge&logo=docker" alt="Docker">
  <img src="https://img.shields.io/badge/Tests-12%20Passing-brightgreen?style=for-the-badge" alt="12 passing tests">
</p>

---

<table align="center">
  <tr>
    <td align="center">
      <img src="github_assets/pasteboxdesktop.gif" alt="Pastebox desktop demo" width="520">
      <br>
      <strong>Desktop</strong>
    </td>
    <td align="center">
      <img src="github_assets/pasteboxmobile.gif" alt="Pastebox mobile demo" width="220">
      <br>
      <strong>Mobile</strong>
    </td>
  </tr>
</table>

---

## About

**Pastebox** is a paste-sharing web application built with pure Django, Tailwind CSS, PostgreSQL, and Docker.

Users can create private accounts, manage their own text pastes, search and filter their workspace, mark pastes as favorites, and share read-only public links.

Pastebox was developed as a learning-focused Django project with attention to project organization, authentication, permissions, business rules, responsive design, and automated tests.

---

## Features

* Pure Django web application
* PostgreSQL database
* Docker-based development environment
* Tailwind CSS interface
* Custom visual design system
* Responsive desktop and mobile layouts
* User registration
* User login and logout
* Session-based authentication
* Protected private paste dashboard
* Per-user paste ownership
* Paste creation
* Paste listing
* Paste editing
* Paste deletion
* Favorite and unfavorite actions
* Text search
* Favorite filtering
* Pagination
* Public shareable paste links
* Public read-only paste pages
* Copy-to-clipboard public links
* Public slug regeneration after paste updates
* Old public link invalidation after paste updates
* Permission-based paste access
* Django messages for user feedback
* Automatically dismissed feedback messages
* Django Admin integration
* Automated model and view tests

---

## Design System

Pastebox uses a clean and minimal visual style.

| Token           | Color     | Usage                                      |
| --------------- | --------- | ------------------------------------------ |
| Main background | `#ffffff` | Page background                            |
| Surface         | `#eeeeee` | Cards, inputs, and secondary areas         |
| Accent          | `#ffbb11` | Hover states, focus states, and highlights |
| Text            | `#2f2f2f` | Main text                                  |
| Muted text      | `#6b6b6b` | Secondary text                             |

The interface uses a centered `900px` container to keep the layout readable and consistent.

---

## Project Structure

```txt
pastebox/
├── github_assets/
│   ├── pasteboxdesktop.gif
│   └── pasteboxmobile.gif
├── accounts/
│   ├── migrations/
│   ├── admin.py
│   ├── apps.py
│   ├── models.py
│   ├── tests.py
│   ├── urls.py
│   └── views.py
├── config/
│   ├── settings.py
│   ├── urls.py
│   ├── asgi.py
│   └── wsgi.py
├── pastes/
│   ├── migrations/
│   ├── admin.py
│   ├── apps.py
│   ├── forms.py
│   ├── models.py
│   ├── tests.py
│   ├── urls.py
│   └── views.py
├── static/
│   ├── src/
│   │   └── input.css
│   └── css/
│       └── output.css
├── templates/
│   ├── base.html
│   ├── partials/
│   │   ├── messages.html
│   │   └── navbar.html
│   ├── accounts/
│   │   ├── login.html
│   │   └── register.html
│   └── pastes/
│       ├── paste_form.html
│       ├── paste_list.html
│       ├── paste_public_detail.html
│       └── partials/
│           ├── pagination.html
│           ├── paste_card.html
│           └── paste_filter_form.html
├── compose.yml
├── Dockerfile
├── manage.py
├── package.json
├── package-lock.json
├── requirements.txt
├── .env.example
├── .gitattributes
├── .gitignore
└── README.md
```

---

## Requirements

* Docker
* Docker Compose

The application runs inside Docker containers, so Python, PostgreSQL, and Node.js do not need to be installed directly on the host machine.

---

## Installation

Clone the repository:

```bash
git clone git@github.com:lucasmirandalm/pastebox.git
```

Enter the project folder:

```bash
cd pastebox
```

Create the local environment file:

```bash
cp .env.example .env
```

Build the containers:

```bash
docker compose build
```

Start the application:

```bash
docker compose up
```

Open Pastebox:

```txt
http://localhost:8000
```

---

## Environment Variables

The `.env.example` file contains the variables required by the application.

```env
POSTGRES_DB=pastebox
POSTGRES_USER=pastebox_user
POSTGRES_PASSWORD=pastebox_password
POSTGRES_HOST=db
POSTGRES_PORT=5432
POSTGRES_EXTERNAL_PORT=5433
```

`POSTGRES_PORT` is used for communication between the Django and PostgreSQL containers.

`POSTGRES_EXTERNAL_PORT` exposes PostgreSQL to the host machine and can be changed when port `5432` is already in use.

---

## Database Setup

Run the migrations:

```bash
docker compose run --rm web python manage.py migrate
```

Create a superuser:

```bash
docker compose run --rm web python manage.py createsuperuser
```

Access the Django Admin:

```txt
http://localhost:8000/admin/
```

---

## Tailwind CSS

Pastebox uses Tailwind CSS through a dedicated Docker service.

The Tailwind source file is located at:

```txt
static/src/input.css
```

The generated CSS file is located at:

```txt
static/css/output.css
```

During development, the Tailwind container runs in watch mode and recompiles the stylesheet when templates or source styles change.

---

## Usage

Start all application services:

```bash
docker compose up
```

Useful URLs:

```txt
http://localhost:8000/
http://localhost:8000/register/
http://localhost:8000/login/
http://localhost:8000/admin/
```

The authenticated dashboard includes:

* text search
* favorite filtering
* paste listing
* pagination
* paste creation
* paste editing
* paste deletion
* favorite and unfavorite actions
* public link sharing

---

## Authentication and Permissions

Pastebox uses Django's built-in authentication and session system.

Private views require authentication, and all private paste queries are restricted by the current user:

```python
Paste.objects.filter(owner=request.user)
```

Update, favorite, and delete operations also verify ownership before changing a paste.

This prevents users from viewing or modifying another user's private pastes.

---

## Public Links

Each paste receives a unique public UUID.

Public links can be opened without authentication and display a read-only version of the paste.

When a paste is edited and saved, Pastebox generates a new public slug. As a result:

* the previous public link becomes invalid;
* the newest public link becomes active;
* favorite and unfavorite actions do not change the public link.

---

## Running Tests

Pastebox includes **12 essential automated tests** covering its main models, permissions, business rules, and views.

Run the complete test suite:

```bash
docker compose run --rm web python manage.py test
```

Run only the tests from the `pastes` app:

```bash
docker compose run --rm web python manage.py test pastes
```

Run the tests with detailed output:

```bash
docker compose run --rm web python manage.py test pastes -v 2
```

If Docker requires elevated permissions on your machine:

```bash
sudo docker compose run --rm web python manage.py test
```

### Test coverage

The test suite verifies that:

* the string representation of a paste returns its title;
* a public slug is generated automatically;
* users only see their own private pastes;
* authenticated users can create pastes;
* unauthenticated users cannot create pastes;
* paste owners can update their own pastes;
* users cannot update another user's paste;
* updating a paste generates a new public slug;
* favoriting changes only the favorite status;
* paste owners can delete their own pastes;
* public paste pages are accessible without login;
* text search and favorite filtering work together.

The Django test runner creates a temporary test database and destroys it after the suite finishes. Tests do not modify the development database.

---

## Why I Built This

I built Pastebox as a Django learning project to practice:

* pure Django development
* PostgreSQL integration
* Docker-based environments
* Tailwind CSS styling
* authentication and authorization
* session-based login
* ownership and permissions
* Django models and forms
* class organization
* reusable template partials
* Django messages
* URL routing
* public UUID links
* business rules
* responsive interface design
* automated model and view testing
* debugging real application errors

---

## Project Status

Pastebox is complete as a learning project.

The implemented scope includes the main application workflow, authentication, paste management, public sharing, responsive design, permissions, and essential automated tests.

Future Django concepts and additional features will be explored in separate projects.

---

## License

This project is open source and available under the MIT License.

---

<p align="center">
  Built with Django, Tailwind CSS, PostgreSQL, Docker, and a lot of practice.
</p>

