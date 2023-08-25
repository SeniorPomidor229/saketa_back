package services

import (
	"context"
	"saketa/configs"
	"saketa/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeServiceMongo struct {
	db *configs.DB
}

func NewEmployeeService(db *configs.DB) *EmployeeServiceMongo {
	return &EmployeeServiceMongo{db: db}
}

func (es *EmployeeServiceMongo) CreateEmployee(employee models.Employee) error {
	_, err := es.db.Collection.InsertOne(context.Background(), employee)
	return err
}

func (es *EmployeeServiceMongo) GetEmployeeByID(id string) (models.Employee, error) {
	var employee models.Employee
	objID, _ := primitive.ObjectIDFromHex(id)
	err := es.db.Collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&employee)
	return employee, err
}

func (es *EmployeeServiceMongo) GetAllEmployees() ([]models.Employee, error) {
	var employees []models.Employee
	cursor, err := es.db.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var employee models.Employee
		err := cursor.Decode(&employee)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, nil
}