package models

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"first_name"`
	LastName  string             `bson:"last_name"`
	Email     string             `bson:"email"`
	Phone     string             `bson:"phone"`
	Department string            `bson:"department"`
}

type EmployeeService interface {
	CreateEmployee(employee Employee) error
	GetEmployeeByID(id string) (Employee, error)
	GetAllEmployees() ([]Employee, error)
}