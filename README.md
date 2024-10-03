
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
[
  {
    "title": "The Great Gatsby",
    "author": "F. Scott Fitzgerald",
    "genre": "fiction"
  },
  {
    "title": "To Kill a Mockingbird",
    "author": "Harper Lee",
    "genre": "fiction"
  }
]
```

---

#### 2. Schedule Book Pickup

**Endpoint:** `/schedule`

**Method:** `POST`

**Description:** Schedules a book pickup.

**Request Body:**
| Name       | Type    | Description                         |
|------------|---------|-------------------------------------|
| book_id    | string  | The ID of the book to be picked up  |
| date       | string  | The date for the pickup (YYYY-MM-DD) |

**Example Request:**
```bash
POST http://localhost:8080/schedule
```

**Example Body:**
```json
{
  "book_id": "12345",
  "date": "2024-10-05"
}
```

**Example Response:**
```json
{
  "message": "Pickup scheduled successfully."
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
[
  {
    "book_id": "12345",
    "pickup_date": "2024-10-05"
  },
  {
    "book_id": "67890",
    "pickup_date": "2024-10-10"
  }
]
```

---