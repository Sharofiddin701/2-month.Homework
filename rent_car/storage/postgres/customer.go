package storage

import (
	"database/sql"
	"fmt"
	"rent-car/models"
	"rent-car/pkg"

	"github.com/google/uuid"
)

type customerRepo struct{
	db *sql.DB
}

func NewCustomer(db *sql.DB) customerRepo {
	return customerRepo{
		db: db,
	}
}

func (c *customerRepo) Create(customer models.Customer) (string,error) {
id := uuid.New()

query := `insert into customers(id,first_name,last_name,gmail,phone,is_blocked) values($1,$2,$3,$4,$5,$6)`

_,err := c.db.Exec(query,id.String(),customer.FirstName,customer.LastName,customer.Gmail,customer.Phone,customer.Is_Blocked)
if err != nil{
	return "error:",err
}
return id.String(),nil
}

func (c *customerRepo) Update(customer models.Customer) (string,error) {
	query:=`update customers set 
	first_name=$1,
	last_name=$2,
	gmail=$3,
	phone=$4,
	is_blocked=$5,
	updated_at=CURRENT_TIMESTAMP
	WHERE id = $6 AND deleted_at=0
	`
	_,err := c.db.Exec(query,
		customer.FirstName,
		customer.LastName,
		customer.Gmail,
		customer.Phone,
		customer.Is_Blocked,
		customer.Id)
	if err != nil {
		return "",err
	}
	return customer.Id,nil
}

func (c *customerRepo) GetAll(req models.GetAllCustomersRequest) (models.GetAllCustomersResponse,error) {
	var (
		resp = models.GetAllCustomersResponse{}
		filter = ""
	)
     offset :=(req.Page -1) * req.Limit
	if req.Search !=""{
		filter +=fmt.Sprintf(`and first_name ILIKE '%%%v%%'`,req.Search)
	}
	
     filter += fmt.Sprintf("OFFSET %v LIMIT %v",offset,req.Limit)
	 fmt.Println("filter:",filter)
	rows,err :=c.db.Query(`select count(id) over(),
	id,
	first_name,
	last_name,	
	gmail,
	phone,
	is_blocked,
	created_at::date,
	updated_at FROM customers WHERE deleted_at = 0`+ filter + ``)
	if err != nil {
		return resp,err
	}
	for rows.Next(){
		var(
			customer = models.Customer{}
			updateAt sql.NullString
		)
		if err := rows.Scan(
			&resp.Count,
			&customer.Id,
			&customer.FirstName,
			&customer.LastName,
			&customer.Gmail,
			&customer.Phone,
			&customer.Is_Blocked,
			&customer.CreatedAt,
			&updateAt);err != nil {
				return resp,err
			}
			customer.UpdatedAt = pkg.NullStringToString(updateAt)
			resp.Customers = append(resp.Customers, customer)
	}
	return resp,nil
}

func (c *customerRepo) GetByID(id string) (models.Customer,error) {
	customer:=models.Customer{}

	if err := c.db.QueryRow(`select id,first_name,last_name,gmail,phone,is_blocked from customers where id = $1`,id).Scan(
		&customer.Id,
		&customer.FirstName,
		&customer.LastName,
		&customer.Gmail,
		&customer.Phone,
		&customer.Is_Blocked,);err != nil{
			return models.Customer{},err
		}
		return customer,nil
}

func (c *customerRepo) Delete(id string) error{
	queary := `delete from customers where id = $1`
	_,err :=c.db.Exec(queary,id)
	if err != nil {
		return err
	}
	return nil
}