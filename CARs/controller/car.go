package controller

import (
	"fmt"
	"rent_car/models"

	"strconv"

	"time"
)

func (c *Controller) CreateCar() {
	car := getCarInfo()
	var count = 0
	var a string = strconv.Itoa(car.Year)
	for i := 0; i < len(a); i++ {
		count++
	}
	if car.Year <= 0 || car.Year > time.Now().Year()+1 || count > 4 {
		fmt.Println("year intput is not correct")
		return
	}
	id, err := c.Store.Car.Create(car)
	if err != nil {
		fmt.Println("error while creating car, err: ", err)
		return
	}
	fmt.Printf("Car created successfully with ID: %v\n", id)
}

func (c *Controller) UpdateCar() {
	car := getCarforUpdating()
	var count = 0
	var a string = strconv.Itoa(car.Year)
	for i := 0; i < len(a); i++ {
		count++
	}
	if car.Year <= 0 || car.Year > time.Now().Year()+1 || count > 4 {
		fmt.Println("year intput is not correct")
		return
	}
	err := c.Store.Car.Update(car)
	if err != nil {
		fmt.Println("error while updating car, err: ", err)
		return
	}
	fmt.Printf("Car updated successfully :")
}

func getCarforUpdating() models.Car {
	var car models.Car
	fmt.Println(`Enter new car data 
	Name
	Year
	Brand
	Model
	HoursePower
	Colour
	EngineCap
	ID`)
	fmt.Scan(&car.Name, &car.Year, &car.Brand, &car.Model, &car.HoursePower, &car.Colour, &car.EngineCap, &car.Id)

	return car
}

func (c *Controller) Delete() {
	carId := getIdToDelete()
	err := c.Store.Car.Delete(carId)
	if err != nil {
		fmt.Println("error while deleting car, err: ", err)
	}
	fmt.Println("successfully deleted")

}

func getCarInfo() models.Car {
	car := models.Car{}
	fmt.Println(`Enter the car datas 
	Name
	Year
	Brand
	Model
	HoursePower
	Colour
	EngineCap`)
	fmt.Scan(&car.Name, &car.Year, &car.Brand, &car.Model, &car.HoursePower, &car.Colour, &car.EngineCap)

	return car
}

func getIdToDelete() string {
	var car models.Car
	fmt.Println("Enter id you want to delete: ")
	fmt.Scan(&car.Id)
	return "Deleted"
}
