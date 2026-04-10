# Movies API

REST API for managing movies, built with Go, Gin, and PostgreSQL.

## Tech Stack
- Go
- Gin
- PostgreSQL

##Getting started

1. Create database:

CREATE DATABASE movies;
\c movies
CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    genre TEXT,
    year INT,
    rating DECIMAL(3,1) DEFAULT 0
);

2. Run the server:

go build main.go && ./main

##Endpoints

GET    /movies         - Get all movies
GET    /movies/top     - Get top 5 by rating
POST   /movies         - Add a movie
PUT    /movies/:id     - Update a movie
DELETE /movies/:title  - Delete a movie

##Example

curl -X POST http://localhost:8080/movies \
-H "Content-Type: application/json" \
-d '{"title":"Inception","genre":"Sci-Fi","year":2010,"rating":8.8}'
