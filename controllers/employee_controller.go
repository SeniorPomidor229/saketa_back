package controllers

import(
	"saketa/services"
	"saketa/models"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	*services.EmployeeServiceMongo
}

func NewApp(employeeService *services.EmployeeServiceMongo) *App {
	return &App{EmployeeServiceMongo: employeeService}
}

func (app *App) SetupRoutes() *fiber.App {
	router := fiber.New()

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Employee Directory!")
	})

	router.Get("/employees", app.getAllEmployees)
	router.Get("/employees/:id", app.getEmployeeByID)
	router.Post("/employees", app.createEmployee)

	return router
}

func (app *App) getAllEmployees(c *fiber.Ctx) error {
	employees, err := app.EmployeeServiceMongo.GetAllEmployees()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(employees)
}

func (app *App) getEmployeeByID(c *fiber.Ctx) error {
	employeeID := c.Params("id")
	employee, err := app.EmployeeServiceMongo.GetEmployeeByID(employeeID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Employee not found"})
	}
	return c.JSON(employee)
}

func (app *App) createEmployee(c *fiber.Ctx) error {
	var employee models.Employee
	if err := c.BodyParser(&employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := app.EmployeeServiceMongo.CreateEmployee(employee); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create employee"})
	}

	return c.JSON(fiber.Map{"message": "Employee created successfully"})
}