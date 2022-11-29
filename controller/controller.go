package controller

import (
	"database/sql"
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
	row, err := db.DB.Query("SELECT * FROM privyTest.cakes ORDER By rating DESC, title ASC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Errors",
			"message": err.Error(),
		})
		return
	}
	defer row.Close()
	if row.Next() {
		selDB, err := db.DB.Query("SELECT * FROM privyTest.cakes ORDER By rating DESC, title ASC")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Errors",
				"message": err.Error(),
			})
			return
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
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  "Errors",
					"message": err.Error(),
				})
				return
			}
			cake.ID = id
			cake.Title = title
			cake.Description = description
			cake.Rating = rating
			cake.Image = image
			cake.CreatedAt, err = time.Parse(layoutDateTime, created_at)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  "Errors",
					"message": err.Error(),
				})
				return
			}
			cake.UpdatedAt, err = time.Parse(layoutDateTime, created_at)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  "Errors",
					"message": err.Error(),
				})
				return
			}
			cakes = append(cakes, cake)
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "get data cakes",
			"data":    cakes,
		})
		defer selDB.Close()
		defer row.Close()
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Errors",
			"message": "There is no data in database",
		})
		defer row.Close()
		return
	}

}

func (db *Repo) GetCakeById(c *gin.Context) {
	id := c.Param("id")
	row, err := db.DB.Query("SELECT * FROM privyTest.cakes WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Errors",
			"message": err.Error(),
		})
		return
	}
	if row.Next() {
		selDB, err := db.DB.Query("SELECT * FROM privyTest.cakes WHERE id=?", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Errors",
				"message": err.Error(),
			})
			return
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
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  "Errors",
					"message": err.Error(),
				})
				return
			}
			cake.ID = id
			cake.Title = title
			cake.Description = description
			cake.Rating = rating
			cake.Image = image
			cake.CreatedAt, err = time.Parse(layoutDateTime, created_at)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  "Errors",
					"message": err.Error(),
				})
				return
			}
			cake.UpdatedAt, err = time.Parse(layoutDateTime, created_at)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  "Errors",
					"message": err.Error(),
				})
				return
			}
			cakes = append(cakes, cake)
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "get data cake by ID",
			"data":    cakes,
		})
		defer selDB.Close()
		defer row.Close()
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Errors",
			"message": "There is no data for ID: " + id,
		})
		defer row.Close()
		return
	}

}

func (db *Repo) AddCake(c *gin.Context) {
	cake := cake_model.CAKE{}
	if err := c.ShouldBindJSON(&cake); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Errors",
			"message": err.Error(),
		})
		return
	}

	insert, err := db.DB.Query("INSERT INTO privyTest.cakes (title,description,rating,image,created_at,updated_at) VALUES (?,?,?,?,?,?)",
		cake.Title, cake.Description, cake.Rating, cake.Image, time.Now().Format(layoutDateTime), time.Now().Format(layoutDateTime))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Errors",
			"message": err.Error(),
		})
		return
	}
	defer insert.Close()

	selDB, err := db.DB.Query("SELECT * FROM privyTest.cakes")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Errors",
			"message": err.Error(),
		})
		return
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
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Errors",
				"message": err.Error(),
			})
			return
		}
		cake.ID = id
		cake.Title = title
		cake.Description = description
		cake.Rating = rating
		cake.Image = image
		cake.CreatedAt, err = time.Parse(layoutDateTime, created_at)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Errors",
				"message": err.Error(),
			})
			return
		}
		cake.UpdatedAt, err = time.Parse(layoutDateTime, created_at)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Errors",
				"message": err.Error(),
			})
			return
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
	row, err := db.DB.Query("SELECT * FROM privyTest.cakes WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Errors",
			"message": err.Error(),
		})
		return
	}
	if row.Next() {
		cake := cake_model.CAKE{}
		if err := c.ShouldBindJSON(&cake); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "Errors",
				"message": err.Error(),
			})
			return
		}

		update, err := db.DB.Query("UPDATE privyTest.cakes SET title=?, description=?, rating=?, image=?, updated_at=? WHERE id=?",
			cake.Title, cake.Description, cake.Rating, cake.Image, time.Now().Format(layoutDateTime), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Errors",
				"message": err.Error(),
			})
			return
		}
		defer update.Close()

		selDB, err := db.DB.Query("SELECT * FROM privyTest.cakes WHERE id=?", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Errors",
				"message": err.Error(),
			})
			return
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
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  "Errors",
					"message": err.Error(),
				})
				return
			}
			cake.ID = id
			cake.Title = title
			cake.Description = description
			cake.Rating = rating
			cake.Image = image
			cake.CreatedAt, err = time.Parse(layoutDateTime, created_at)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  "Errors",
					"message": err.Error(),
				})
				return
			}
			cake.UpdatedAt, err = time.Parse(layoutDateTime, created_at)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  "Errors",
					"message": err.Error(),
				})
				return
			}
			cakes = append(cakes, cake)
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "update data cake",
			"data":    cakes,
		})

		defer selDB.Close()
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Errors",
			"message": "There is no data for ID: " + id,
		})
		defer row.Close()
		return
	}

}

func (db *Repo) DeleteCakeById(c *gin.Context) {
	id := c.Param("id")
	row, err := db.DB.Query("SELECT * FROM privyTest.cakes WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Errors",
			"message": err.Error(),
		})
		return
	}
	if row.Next() {
		delete, err := db.DB.Query("DELETE From privyTest.cakes WHERE id=?", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Errors",
				"message": err.Error(),
			})
			return
		}
		defer delete.Close()
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "delete id:" + id + " is success",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Errors",
			"message": "There is no data for ID: " + id,
		})
		defer row.Close()
		return
	}
}
