package main

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/movies", func(c *gin.Context) {
		rows, _ := db.Query("SELECT id, title, genre, year, rating FROM movies")
		defer rows.Close()
		var movies []Movie
		for rows.Next() {
			var m Movie
			rows.Scan(&m.ID, &m.Title, &m.Genre, &m.Year, &m.Rating)
			movies = append(movies, m)
		}
		c.JSON(200, movies)
	})
	return r
}

func TestGetMovies(t *testing.T) {
	var err error
	db, err = sql.Open("postgres", "host=localhost user=almat dbname=movies sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/movies", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Expected 200 but got %d", w.Code)
	}
}
