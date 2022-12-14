# Phonebook GoLang REST API

This is a simple REST API written in GoLang for learning purposes.


| Methods | Endpoints           | Action                             |
| --------- | :-------------------- | ------------------------------------ |
| GET     | /api/v1/persons     | Show details of all contacts       |
| GET     | /api/v1/persons/:id | Show details of a specific contact |
| POST    | /api/v1/persons     | Add a new person to the phonebook  |
| PATCH   | /api/v1/persons/:id | Edit details of a contact          |
| DELETE  | /api/v1/persons/:id | Delete contact from phonebook      |
| GET     | /api/v1/info        | Show application information       |
| GET     | /api/v1/healthz     | Uptime check of the server         |

## Deployment to Google Cloud Run

Prerequisites:

* A project on Google Cloud

### Imperatively with gcloud

### Declaratively with Terrafom
