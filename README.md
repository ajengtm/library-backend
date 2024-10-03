
# Library Backend API

This is the backend service for managing library book data and pickup schedules. The service provides APIs to get books by genre and manage book pickup schedules.

## Getting Started

### Prerequisites
- Go version 1.16 or above
- A working library-backend project structure with handlers

### Installation

Clone the repository:
```bash
git clone https://github.com/yourusername/library-backend.git
```

Navigate to the project directory:
```bash
cd library-backend
```

Install the dependencies:
```bash
go mod tidy
```

### Running the Application

To run the application, use the following command:
```bash
go run main.go
```

## API Documentation

### Base URL
```
http://localhost:8080
```

### Endpoints

#### 1. Get Books by Genre

**Endpoint:** `/books`

**Method:** `GET`

**Description:** Retrieves a list of books filtered by genre.

**Query Parameters:**
| Name   | Type   | Description      |
|--------|--------|------------------|
| genre  | string | Genre of the books (e.g., `fiction`, `non-fiction`) |

**Example Request:**
```bash
GET http://localhost:8080/books?genre=fiction
```

**Example Response:**
```json
{
    "data": [
        {
            "title": "Pride and Prejudice",
            "author": "[]",
            "edition_number": "4035"
        },
        {
            "title": "Treasure Island",
            "author": "[]",
            "edition_number": "1984"
        }
    ],
    "message": "Books retrieved successfully",
    "status": "success"
}
```

---

#### 2. Schedule Book Pickup

**Endpoint:** `/schedule`

**Method:** `POST`

**Description:** Schedules a book pickup.

**Request Body:**
| Name       | Type    | Description                          |
|------------|---------|--------------------------------------|
| user       | string  | The name of the user scheduling the pickup |
| book_title | string  | The title of the book to be picked up |
| schedule   | string  | The date and time for the pickup in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ) |

**Example Request:**
```bash
POST http://localhost:8080/schedule
```

**Example Body:**
```json
{
    "user": "Ajeng",
    "book_title": "Women in Love",
    "schedule": "2024-10-02T10:00:00Z"
}
```

**Example Response:**
```json
{
    "message": "Book pick-up scheduled successfully",
    "status": "success"
}
```

---

#### 3. Get All Pickup Schedules

**Endpoint:** `/schedules`

**Method:** `GET`

**Description:** Retrieves all scheduled book pickups.

**Example Request:**
```bash
GET http://localhost:8080/schedules
```

**Example Response:**
```json
{
    "data": [
        {
            "book_title": "Women in Love",
            "pick_up_time": ""
        }
    ],
    "message": "Pick-up schedules retrieved successfully",
    "status": "success"
}
```

---