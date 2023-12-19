package Controllers

import (
	"net/http"

	// "../Models"
	"example.go/Models"
	"github.com/gin-gonic/gin"
)

// list all todos

func GetTodos(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	var todo []Models.Todo
	err := Models.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// create todo
func CreateATodo(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	var todo Models.Todo
	c.BindJSON(&todo)
	err := Models.CreateAToDo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetATodo(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	id := c.Params.ByName("id")
	var todo Models.Todo
	err := Models.GetATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateATodo(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	var todo Models.Todo
	id := c.Params.ByName("id")
	err := Models.GetATodo(&todo, id)
	if err != nil {
		c.JSON(http.StatusNotFound, todo)
	}

	c.BindJSON(&todo)
	err = Models.UpdateATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	var todo Models.Todo
	id := c.Params.ByName("id")
	err := Models.DeleteATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
