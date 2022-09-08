package main

import (
	//"encoding/csv"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

// todos represents data about a record Todos.
type todos struct {
	ID   int    `csv:"id"`
	Name string `csv:"name"`
	Done bool   `csv:"true"`
}

var todo = []todos{
	{ID: 1, Name: "Dilbag", Done: true},
	{ID: 2, Name: "Eric", Done: true},
	{ID: 3, Name: "Deepa", Done: true},
}

// gettodos responds with the list of all list as JSON.
func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todo)

}

// posttodo adds an todos from JSON received in the request body.
func Posttodo(c *gin.Context) {
	var newtodo todos

	// Call BindJSON to bind the received JSON to
	// newTodo.
	if err := c.BindJSON(&newtodo); err != nil {
		return
	}
	//todo = strconv.ParseInt()

	// Add the new todo to the slice.
	todo = append(todo, newtodo)
	c.IndentedJSON(http.StatusCreated, newtodo)

}

func PatchtodosByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of todos, looking for
	// an todos whose ID value matches the parameter.
	for _, a := range todo {
		if strconv.Itoa(a.ID) == id {
			a.Done = !a.Done
			id_int, _ := strconv.Atoi(id)

			todo[id_int-1].Done = !todo[id_int-1].Done
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todos not found"})
}

func DeletetodosByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of todos, looking for
	// an todos whose ID value matches the parameter.
	for _, a := range todo {
		if strconv.Itoa(a.ID) == id {
			a.Done = !a.Done
			id_int, _ := strconv.Atoi(id)
			todo = append(todo[:id_int-1], todo[id_int:]...)

			//	todo[id_int-1].Done = !todo[id_int-1].Done
			c.IndentedJSON(http.StatusOK, todo)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todos not found"})
}

func main() {

	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", Posttodo)
	router.PATCH("/todos/:id", PatchtodosByID)
	router.DELETE("/todos/:id", DeletetodosByID)

	router.Run("localhost:5000")
}
