package controller

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	cake_model "rest_api/models"

	"github.com/gin-gonic/gin"
)

const layoutDateTime = "2006-01-02 15:04:05"

type Repo struct {
	DB *sql.DB
}

func (db *Repo) GetCake(c *gin.Context) {
	selDB, err := db.DB.Query("SELECT * FROM cakes ORDER By rating DESC, title ASC")
	if err != nil {
		panic(err.Error())
	}
	cake := cake_model.CAKE{}
	cakes := []cake_model.CAKE{}
	for selDB.Next() {
		var (
			id                        int32
			title, description, image string
			rating                    float64
			created_at, updated_at    string
		)
		err = selDB.Scan(&id, &title, &description, &rating, &image, &created_at, &updated_at)
		if err != nil {
			panic(err.Error())
		}
		cake.ID = id
		cake.Title = title
		cake.Description = description
		cake.Rating = rating
		cake.Image = image
		cake.CreatedAt, err = time.Parse(layoutDateTime, created_at)
		if err != nil {
			log.Fatal(err)
		}
		cake.UpdatedAt, err = time.Parse(layoutDateTime, created_at)
		if err != nil {
			log.Fatal(err)
		}
		cakes = append(cakes, cake)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "get data cakes",
		"data":    cakes,
	})
	defer selDB.Close()
}

func (db *Repo) GetCakeById(c *gin.Context) {
	id := c.Param("id")
	selDB, err := db.DB.Query("SELECT * FROM cakes WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}
	cake := cake_model.CAKE{}
	cakes := []cake_model.CAKE{}
	for selDB.Next() {
		var (
			id                        int32
			title, description, image string
			rating                    float64
			created_at, updated_at    string
		)
		err = selDB.Scan(&id, &title, &description, &rating, &image, &created_at, &updated_at)
		if err != nil {
			panic(err.Error())
		}
		cake.ID = id
		cake.Title = title
		cake.Description = description
		cake.Rating = rating
		cake.Image = image
		cake.CreatedAt, err = time.Parse(layoutDateTime, created_at)
		if err != nil {
			log.Fatal(err)
		}
		cake.UpdatedAt, err = time.Parse(layoutDateTime, created_at)
		if err != nil {
			log.Fatal(err)
		}
		cakes = append(cakes, cake)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "get data cake by ID",
		"data":    cakes,
	})
	defer selDB.Close()
}

func (db *Repo) AddCake(c *gin.Context) {
	cake := cake_model.CAKE{}
	if err := c.ShouldBindJSON(&cake); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insert, err := db.DB.Query("INSERT INTO cakes (title,description,rating,image,created_at,updated_at) VALUES (?,?,?,?,?,?)",
		cake.Title, cake.Description, cake.Rating, cake.Image, time.Now().Format(layoutDateTime), time.Now().Format(layoutDateTime))
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	selDB, err := db.DB.Query("SELECT * FROM cakes")
	if err != nil {
		panic(err.Error())
	}
	cakes := []cake_model.CAKE{}
	for selDB.Next() {
		var (
			id                        int32
			title, description, image string
			rating                    float64
			created_at, updated_at    string
		)
		err = selDB.Scan(&id, &title, &description, &rating, &image, &created_at, &updated_at)
		if err != nil {
			panic(err.Error())
		}
		cake.ID = id
		cake.Title = title
		cake.Description = description
		cake.Rating = rating
		cake.Image = image
		cake.CreatedAt, err = time.Parse(layoutDateTime, created_at)
		if err != nil {
			log.Fatal(err)
		}
		cake.UpdatedAt, err = time.Parse(layoutDateTime, created_at)
		if err != nil {
			log.Fatal(err)
		}
		cakes = append(cakes, cake)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "add data cake",
		"data":    cakes,
	})
	defer selDB.Close()
}

func (db *Repo) UpdateCake(c *gin.Context) {
	id := c.Param("id")
	cake := cake_model.CAKE{}
	if err := c.ShouldBindJSON(&cake); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update, err := db.DB.Query("UPDATE cakes SET title=?, description=?, rating=?, image=?, updated_at=? WHERE id=?",
		cake.Title, cake.Description, cake.Rating, cake.Image, time.Now().Format(layoutDateTime), id)
	if err != nil {
		panic(err.Error())
	}
	defer update.Close()

	selDB, err := db.DB.Query("SELECT * FROM cakes WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}
	cakes := []cake_model.CAKE{}
	for selDB.Next() {
		var (
			id                        int32
			title, description, image string
			rating                    float64
			created_at, updated_at    string
		)
		err = selDB.Scan(&id, &title, &description, &rating, &image, &created_at, &updated_at)
		if err != nil {
			panic(err.Error())
		}
		cake.ID = id
		cake.Title = title
		cake.Description = description
		cake.Rating = rating
		cake.Image = image
		cake.CreatedAt, err = time.Parse(layoutDateTime, created_at)
		if err != nil {
			log.Fatal(err)
		}
		cake.UpdatedAt, err = time.Parse(layoutDateTime, created_at)
		if err != nil {
			log.Fatal(err)
		}
		cakes = append(cakes, cake)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "update data cake",
		"data":    cakes,
	})

	defer selDB.Close()
}

func (db *Repo) DeleteCakeById(c *gin.Context) {
	id := c.Param("id")
	delete, err := db.DB.Query("DELETE From cakes WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}
	defer delete.Close()
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "delete id:" + id + " is success",
	})

}
