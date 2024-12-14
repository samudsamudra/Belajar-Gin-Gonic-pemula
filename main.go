package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var toDoList = map[int]string{}
var currentID = 1 // ID unik yang akan bertambah otomatis

func main() {
	router := gin.Default()

	// Endpoint: Ambil semua tugas
	router.GET("/todos", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"tasks": toDoList})
	})

	// Endpoint: Ambil tugas berdasarkan ID
	router.GET("/todos/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		task, exists := toDoList[id]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"id": id, "task": task})
	})

	// Endpoint: Tambahkan tugas baru
	router.POST("/todos", func(c *gin.Context) {
		var newTask struct {
			Task string `json:"task"`
		}

		if err := c.BindJSON(&newTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		toDoList[currentID] = newTask.Task
		c.JSON(http.StatusCreated, gin.H{"id": currentID, "message": "Task added"})
		currentID++
	})

	// Endpoint: Perbarui tugas berdasarkan ID
	router.PUT("/todos/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		var updatedTask struct {
			Task string `json:"task"`
		}
		if err := c.BindJSON(&updatedTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, exists := toDoList[id]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		toDoList[id] = updatedTask.Task
		c.JSON(http.StatusOK, gin.H{"message": "Task updated", "id": id})
	})

	// Endpoint: Hapus tugas berdasarkan ID
	router.DELETE("/todos/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		_, exists := toDoList[id]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		delete(toDoList, id)
		c.JSON(http.StatusOK, gin.H{"message": "Task deleted", "id": id})
	})

	router.Run(":8080")
}
