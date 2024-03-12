package storage

import "rent-car/models"


type IStorage interface {
	CloseDB()
	Car() ICarStorage
	Customer() ICustomerStorage
}

type ICarStorage interface {
Create(models.Car)(string,error)
GetByID(id string)(models.Car,error)
GetAll(request models.GetAllCarsRequest) (models.GetAllCarsResponse, error)
Update(models.Car) (string, error)
Delete(string) error
}

type ICustomerStorage interface {
Create(models.Customer)(string,error)
GetByID(id string)(models.Customer,error)
GetAll(request models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error)
Update(models.Customer) (string, error)
Delete(string) error
}