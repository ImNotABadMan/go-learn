package tables

import "time"

type Task struct {
	Id           int
	Site_id      int
	Admin_id     int
	State        int
	Pulled_at    time.Time
	Pull_state   int
	Deleted_at   time.Time
	Delete_state int
	Created_at   time.Time
	Updated_at   time.Time
}

type TaskProduct struct {
	Publication_id   int
	ProductID        int
	Third_product_id string
	Name             string
	Product_code     string
	Currency_id      int
	Enabled          int
	Is_taxable       int
	Is_logistics     int
	Is_group         int
	Created_at       time.Time
	Updated_at       time.Time
}
