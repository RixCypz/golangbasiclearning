package main

import (
	"net/http"
	"os"
	"strconv"

	"github/rixcypz/gin/model"

	"github.com/gin-gonic/gin"
)

var students = []model.Student{
	{Id: 1, Firstname: "Harry", Lastname: "Potter", Age: 20},
	{Id: 2, Firstname: "Tomas", Lastname: "Sheldon", Age: 18},
	{Id: 3, Firstname: "Virgil", Lastname: "Van Lor", Age: 19},
}

func main() {
	r := gin.New()

	r.GET("/student", listStudentsHandler)
	r.POST("/student", createStudentsHandler)
	r.DELETE("/student/:id", deleteStudentHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	r.Run(":" + port)
}

func listStudentsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, students)
}

func createStudentsHandler(c *gin.Context) {
	var student model.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	students = append(students, student)
	c.JSON(http.StatusCreated, student)
}

func deleteStudentHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	for i, a := range students {
		if a.Id == id {
			students = append(students[:i], students[i+1:]...)
			break
		}
	}
	c.Status(http.StatusNoContent)
}
