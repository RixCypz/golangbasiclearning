package main

import (
	"github/rixcypz/fibertest/model"
	"net/http"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Initial student data
var students = []model.Student{
	{Id: 1, Firstname: "Harry", Lastname: "Potter", Age: 20},
	{Id: 2, Firstname: "Tomas", Lastname: "Sheldon", Age: 18},
	{Id: 3, Firstname: "Virgil", Lastname: "Van Lor", Age: 19},
}

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// Define a route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Routes
	app.Get("/student", listStudentsHandler)
	app.Post("/student", createStudentsHandler)
	app.Delete("/student/:id", deleteStudentHandler)

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	// Start Fiber server
	app.Listen(":" + port)
}

func listStudentsHandler(c *fiber.Ctx) error {
	// Return the list of students as JSON response
	return c.JSON(students)
}

func createStudentsHandler(c *fiber.Ctx) error {
	// Bind JSON request body to Student struct
	var student model.Student
	if err := c.BodyParser(&student); err != nil {
		// Return a bad request response if binding fails
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Append the new student to the list
	students = append(students, student)
	// Return the created student as JSON response
	return c.Status(http.StatusCreated).JSON(student)
}

func deleteStudentHandler(c *fiber.Ctx) error {
	// Get student ID from path parameter
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// Return a bad request response if ID conversion fails
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id"})
	}
	// Iterate through students to find and remove the specified student
	for i, a := range students {
		if a.Id == id {
			students = append(students[:i], students[i+1:]...)
			break
		}
	}
	// Return a no content response after successful deletion
	return c.SendStatus(http.StatusNoContent)
}
