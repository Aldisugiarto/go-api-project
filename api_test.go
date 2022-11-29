// Golang REST API unit testing program
package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"rest_api/controller"
	"strconv"
	"testing"

	"rest_api/config"

	"encoding/json"
	cake_model "rest_api/models"

	"github.com/gin-gonic/gin"
)

func connDB() controller.Repo {
	db := config.InitDB(&gin.Context{})
	c := controller.Repo{
		DB: db,
	}
	return c
}
func TestAddCake(t *testing.T) {
	c := connDB()
	r := gin.Default()
	r.POST("/cakes", c.AddCake)

	cake := cake_model.CAKE{
		Title:       "Banana cheesecake",
		Description: "A cheesecake made of lemon",
		Rating:      10,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
	}
	jsonValue, _ := json.Marshal(cake)
	req, err := http.NewRequest("POST", "/cakes", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println(w.Body)

	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

}
func TestGetCake(t *testing.T) {
	gin.SetMode(gin.TestMode)

	c := connDB()
	r := gin.Default()
	r.GET("/cakes", c.GetCake)
	req, err := http.NewRequest(http.MethodGet, "/cakes", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	fmt.Println(w.Body)

	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
func TestGetCakeById(t *testing.T) {
	c := connDB()
	r := gin.Default()
	r.GET("/cakes/:id", c.GetCakeById)

	cakeId := `1`
	req, err := http.NewRequest(http.MethodGet, "/cakes/"+cakeId, nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println(w.Body)

	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
func TestUpdateCake(t *testing.T) {
	c := connDB()
	r := gin.Default()
	r.PATCH("/cakes/:id", c.UpdateCake)

	cake := cake_model.CAKE{
		ID:          1,
		Title:       "Banana cheesecake",
		Description: "A cheesecake made of lemon",
		Rating:      10,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
	}
	jsonValue, _ := json.Marshal(cake)
	req, err := http.NewRequest("PATCH", "/cakes/"+strconv.Itoa(int(cake.ID)), bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println(w.Body)

	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestDeleteCakeById(t *testing.T) {
	c := connDB()
	r := gin.Default()
	r.DELETE("/cakes/:id", c.DeleteCakeById)

	cakeId := `1`
	req, err := http.NewRequest(http.MethodDelete, "/cakes/"+cakeId, nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println(w.Body)

	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
