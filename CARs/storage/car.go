package storage

import (
	"database/sql"
	"fmt"
	"rent_car/models"

	"github.com/google/uuid"
)

type carRepo struct {
	db *sql.DB
}

func NewCar(db *sql.DB) carRepo {
	return carRepo{
		db: db,
	}
}
func (c *carRepo) Create(car models.Car) (string, error) {

	id := uuid.New()

	query := ` INSERT INTO cars (
		id,
		name,
		brand,
		year,
		model,
		hourse_power,
		colour,
		engine_cap)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8) 
	`

	res, err := c.db.Exec(query,
		id.String(),
		car.Name, car.Brand, car.Year,
		car.Model, car.HoursePower,
		car.Colour, car.EngineCap)

	if err != nil {
		return "", err
	}

	fmt.Printf("%+v\n", res)

	return id.String(), nil
}

func (i *carRepo) Update(car models.Car) error {
	_, err := i.db.Exec(
		`UPDATE cars 
		SET name=$1,
			brand=$2,
			year=$3,
			model=$4,
			hoursepower=$5,
			colour=$6,
			engine_cap=$7,
			updated_at=NOW()
			WHERE id=$8`, car.Name, car.Brand, car.Year, car.Model, car.HoursePower, car.Colour, car.EngineCap, car.Id)
	if err != nil {
		fmt.Println("error while updating car err: ", err)
		return err
	}

	return nil
}

func (i *carRepo) Delete(id string) error {
	_, err := i.db.Exec(
		`DELETE FROM cars
			WHERE id=$1`, id)
	if err != nil {
		fmt.Println("error while deleting car info err: ", err)
		return err
	}

	return nil
}

func (i *carRepo) GetAll() ([]models.Car, error) {
	cars := []models.Car{}
	rows, err := i.db.Query(`select 
	id,
	name,
	code,
	created_at from cars WHERE deleted_at is null`)
	if err != nil {
		fmt.Println("error while getting all cars err: ", err)
		return nil, err
	}

	for rows.Next() {
		c := models.Car{}
		if err = rows.Scan(&c.Id, &c.Name, &c.Brand, &c.Colour, &c.Model, &c.EngineCap, &c.HoursePower, &c.Year, &c.CreatedAt); err != nil {
			fmt.Println("error while scanning cars err: ", err)
			return nil, err
		}
		cars = append(cars, c)
	}

	return cars, nil
}
