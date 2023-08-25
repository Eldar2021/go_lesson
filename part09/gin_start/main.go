package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Age      int    `json:"age"`
}

func deleteAtIndex(slice []Person, index int) []Person {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	list := []Person{}
	r := gin.Default()

	/*
	* curl localhost:8080/persons
	 */
	r.GET("/persons", func(ctx *gin.Context) {
		ctx.IndentedJSON(200, list)
	})

	/*
		curl http://localhost:8080/add \
		--include \
		--header "Content-Type: application/json" \
		--request "POST" \
		--data '{"id": 1,"name": "Eldi","lastname": "Almazbek","age": 22}'
	*/
	r.POST("/add", func(ctx *gin.Context) {
		var newPerson Person
		if err := ctx.BindJSON(&newPerson); err != nil {
			log.Fatal(err)
			return
		}
		list = append(list, newPerson)
		ctx.IndentedJSON(http.StatusCreated, newPerson)
	})

	/*
	* curl http://localhost:8080/persons/1
	 */
	r.GET("/persons/:id", func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		for _, p := range list {
			if id == p.Id {
				ctx.IndentedJSON(http.StatusOK, p)
				return
			}
		}
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person not found"})
	})

	/*
	* curl -i -X DELETE http://localhost:8080/delete/1
	 */
	r.DELETE("/delete/:id", func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		for _, p := range list {
			if id == p.Id {
				list = deleteAtIndex(list, p.Id)
				ctx.IndentedJSON(http.StatusOK, p)
				return
			}
		}
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person not found"})
	})

	r.Run()
}
