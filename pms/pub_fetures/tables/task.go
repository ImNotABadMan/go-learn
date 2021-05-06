package tables

import (
	"time"
)

type Publication_task struct {
	Id           int
	Site_id      int
	Admin_id     int
	State        int
	Pulled_at    time.Time `gorm:"default:null"`
	Pull_state   int
	Deleted_at   time.Time `gorm:"default:null"`
	Delete_state int
	Created_at   time.Time            `gorm:"default:null"`
	Updated_at   time.Time            `gorm:"default:null"`
	PuProduct    Publication_products `gorm:"foreignKey:Publication_id;references:Id"`
}

type Publication_products struct {
	Publication_id   int
	ProductID        int    `gorm:"column:productID"`
	Third_product_id string `gorm:"column:productID"`
	Name             string
	Product_code     string
	Currency_id      int
	Enabled          int
	Is_taxable       int
	Is_logistics     int
	Is_group         int
	Created_at       time.Time                 `gorm:"default:null"`
	Updated_at       time.Time                 `gorm:"default:null"`
	PubProductSub    []Publication_product_sub `gorm:"foreignKey:Publication_id;references:Publication_id"`
}

type Publication_product_sub struct {
	Id                    int
	Publication_id        int
	Co_productID          int    `gorm:"column:co_productID"`
	Co_product_code       string `gorm:"column:co_product_code"`
	Co_third_product_id   string `gorm:"column:co_third_product_id"`
	Co_third_inventory_id string `gorm:"column:co_third_inventory_id"`
	Price                 float64
	Original_price        float64
	In_stock              int
	Out_stock_strategy    int
	Created_at            time.Time `gorm:"default:null"`
	Updated_at            time.Time `gorm:"default:null"`
}

type Publication_product_category struct {
	Id               int
	Publication_id   int
	Site_category_id int
	Remarks          string
	Created_at       time.Time `gorm:"default:null"`
	Updated_at       time.Time `gorm:"default:null"`
	Third_collect_id string
}

type Site_category struct {
	Id                int
	Site_id           int
	Third_category_id string
	Title             string
	Type              string
	Published         int
	Local_parent      int
	Remarks           string
	Created_at        time.Time `gorm:"default:null"`
	Updated_at        time.Time `gorm:"default:null"`
}

type Site_config struct {
	Id           int
	Name         string
	Url          string
	Description  string
	Platform     string
	Api_id       string
	Api_key      string
	Api_password string
	Secret_key   string
	Access_token string
	Location_ids string
	AdminID      int
	Created_at   time.Time `gorm:"default:null"`
	Updated_at   time.Time `gorm:"default:null"`
}
