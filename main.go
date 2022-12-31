package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{"1", "Read Killing Commendatore", false},
	{"2", "Learn Go", false},
	{"3", "Play Elden Ring", false},
}

func getTodos(cnxt *gin.Context) {

	cnxt.IndentedJSON(http.StatusOK, todos)
}

func addTodo(cnxt *gin.Context) {

	var newTodo todo

	if err := cnxt.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	cnxt.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoById(id string) (*todo, error) {

	for i, t := range todos {

		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func getTodo(cnxt *gin.Context) {

	id := cnxt.Param("id")

	t, err := getTodoById(id)

	if err != nil {

		cnxt.IndentedJSON(http.StatusNotFound, gin.H{"messgae": "Todo not found"})
		return
	}

	cnxt.IndentedJSON(http.StatusOK, t)
}

func toggleTodoStatus(cnxt *gin.Context) {

	id := cnxt.Param("id")

	t, err := getTodoById(id)

	if err != nil {

		cnxt.IndentedJSON(http.StatusNotFound, gin.H{"messgae": "Todo not found"})
		return
	}

	t.Completed = !t.Completed

	cnxt.IndentedJSON(http.StatusOK, t)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.Run("localhost:8080")
}
