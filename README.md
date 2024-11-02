# RSS Aggregator

A backend server built in Go that fetches, parses, and aggregates data from RSS feeds, storing it in a PostgreSQL database. This tool is designed to be more than a simple CRUD App. It includes a long-running service that continually collects and aggregates posts from RSS feeds specified by users by using goroutines.

## Features

- **Database Integration**: Persistent storage using PostgreSQL with database migrations managed by Goose and SQL code generation via SQLc.
- **Follow Feeds**: Allow users to follow multiple RSS feeds.
- **Continuous Aggregation**: Runs a background worker that continuously fetches and aggregates posts from RSS feeds.

## Technologies Used

- **Go**: Main programming language for Web Server development.
- **PostgreSQL**: Database for storing users, feeds, posts etc.
- **SQLc**: For generating type-safe Go code from SQL queries.
- **Goose**: For managing database migrations.
- **pgAdmin**: For interacting with the PostgreSQL database during development.

## Installation

### Prerequisites

- Go (latest version)
- PostgreSQL
- SQLc
- Goose

### Steps

1. **Clone the Repository**:
   ```
   git clone https://github.com/Sakaar-Sen/RSS-Aggregator
   ```
2. **Install Go dependencies**:
  ```
 go mod tidy
```
3. **Set Up PostgreSQL**:
- Create a PostgreSQL database.
- Update the database connection settings a .env file with key `DB_URL`.

4. **Run Database Migrations**:
   ```
   goose up
   ```
5. Run the Server:
 ```
    go build && ./rssagg
 ```
   
   



