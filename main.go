package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

type Movie struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Year   int     `json:"year"`
	Rating float64 `json:"rating"`
	Genre  string  `json:"genre"`
}

func main() {
	var err error
	db, err = sql.Open("postgres", "host=localhost user=almat dbname=movies sslmode=disable")
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/movies", func(c *gin.Context) {
		rows, err := db.Query("SELECT id, title, genre, year, rating FROM movies")
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()
		var movies []Movie
		for rows.Next() {
			var m Movie
			rows.Scan(&m.ID, &m.Title, &m.Genre, &m.Year, &m.Rating)
			movies = append(movies, m)

		}
		c.JSON(200, movies)
	})
	r.GET("/movies/top", func(c *gin.Context) {
		rows, _ := db.Query("SELECT id, title, genre, year, rating FROM movies ORDER BY rating DESC LIMIT 5")
		defer rows.Close()
		var movies []Movie
		for rows.Next() {
			var m Movie
			rows.Scan(&m.ID, &m.Title, &m.Genre, &m.Year, &m.Rating)
			movies = append(movies, m)
		}
		c.JSON(200, movies)
	})

	r.POST("/movies", func(c *gin.Context) {
		var m Movie
		c.ShouldBindJSON(&m)
		err := db.QueryRow("INSERT INTO movies(title, genre, year, rating) VALUES($1, $2, $3, $4) RETURNING id",
			m.Title, m.Genre, m.Year, m.Rating).Scan(&m.ID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, m)
	})

	r.DELETE("/movies/:title", func(c *gin.Context) {
		title := c.Param("title")
		db.Exec("DELETE FROM movies WHERE title = $1", title)
		c.JSON(200, gin.H{"message": "Movie deleted"})
	})

	r.PUT("/movies/:id", func(c *gin.Context) {
		id := c.Param("id")
		var m Movie
		c.ShouldBindJSON(&m)
		db.Exec("UPDATE movies SET title = $1, genre = $2, year = $3, rating = $4 WHERE id = $5",
			m.Title, m.Genre, m.Year, m.Rating, id)
		c.JSON(200, m)
	})

	r.Run(":8080")
}
