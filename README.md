# Blogist Backend API

REST API ready to use.

## How to clone

```bash 
git clone git@github.com:ekosalvinus/blogist-be-v1.git
```

## How to Run

```bash
go run main.go
```

## Endpoints

| Method | URL | Auth |
|--------|-----|------|
| GET | http://localhost:8080/blog | Not required |
| GET | http://localhost:8080/articles | Basic Auth |

## Example Request

### Request `/articles` with curl

**Using username and password:**
```bash
curl -u admin:admin123 http://localhost:8080/articles
```

**Or using manual header:**
```bash
curl -H "Authorization: Basic YWRtaW46YWRtaW4xMjM=" http://localhost:8080/articles
```

## Authentication

Middleware uses HTTP Basic Auth — if credentials are incorrect or missing, it will return **401 Unauthorized**.