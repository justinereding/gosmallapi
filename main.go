package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ToDo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []ToDo{
	{ID: "1", Item: "Jeffy", Completed: false},
	{ID: "2", Item: "Jeffffy", Completed: false},
	{ID: "3", Item: "Jeffffffy", Completed: false},
}

func getToDos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func createTodo(context *gin.Context) {
	var newTodo ToDo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")

	ToDo, err := getToDoByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "To do not found"})
	}

	context.IndentedJSON(http.StatusOK, ToDo)
}

func getToDoByID(id string) (*ToDo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("To Do not found")
}

func main() {
	router := gin.Default()
	router.GET("/todos", getToDos)
	router.GET("/todos/:id", getTodo)
	router.POST("/todos", createTodo)
	router.Run("localhost:8080")
}
