package main

import (
	"net/http"
	"os"
	"strconv"

	"github/rixcypz/gin/model"

	"github.com/gin-gonic/gin"
)

// initial student data
var students = []model.Student{
	{Id: 1, Firstname: "Harry", Lastname: "Potter", Age: 20},
	{Id: 2, Firstname: "Tomas", Lastname: "Sheldon", Age: 18},
	{Id: 3, Firstname: "Virgil", Lastname: "Van Lor", Age: 19},
}

func main() {
	r := gin.New()

	//call services
	r.GET("/student", listStudentsHandler)
	r.POST("/student", createStudentsHandler)
	r.DELETE("/student/:id", deleteStudentHandler)

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	r.Run(":" + port)
}

func listStudentsHandler(c *gin.Context) {
	// Return the list of students as JSON response
	c.JSON(http.StatusOK, students)
}

func createStudentsHandler(c *gin.Context) {
	// Bind JSON request body to Student struct
	var student model.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		// Return a bad request response if binding fails
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Append the new student to the list
	students = append(students, student)
	// Return the created student as JSON response
	c.JSON(http.StatusCreated, student)
}

func deleteStudentHandler(c *gin.Context) {
	// Get student ID from path parameter
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// Return a bad request response if ID conversion fails
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	// Iterate through students to find and remove the specified student
	for i, a := range students {
		if a.Id == id {
			students = append(students[:i], students[i+1:]...)
			break
		}
	}
	// Return a no content response after successful deletion
	c.Status(http.StatusNoContent)
}
