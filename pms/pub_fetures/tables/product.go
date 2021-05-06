package tables

import (
	"time"
)

type Products struct {
	//gorm.Model
	ProductID          uint `gorm:"primaryKey;column:productID"`
	CategoryID         int  `gorm:"column:categoryID"`
	Name               string
	Name_cn            string
	Product_code       string
	Product_code_type  int
	Group_parent_sku   string
	Isgroup            int
	Group_parent       int
	Platform_id        int
	In_stock           int
	Enabled            int
	Sort_order         int
	Default_picture    int
	Viewed_times       int
	Onhold             int
	Items_sold         int
	Weight             float64
	Cubic              float64
	Just_imported      int
	Added_admin        string
	Modified_admin     int
	In_stock_photo_id  int
	In_stock_photo_url string
	Date_added         time.Time
	Date_modified      time.Time
	Formal_deleted_at  time.Time
	Deleted_at         time.Time
	Additional         Product_additional `gorm:"foreignKey:ProductID;references:ProductID"`
	Prices             []Product_price    `gorm:"foreignKey:ProductID;references:ProductID"`
}

type Product_additional struct {
	//gorm.Model
	ProductID             int `gorm:"primaryKey;column:productID"`
	Original_product_code string
	Supplier_name         string
	Supplier_type         string
	Supplier_url          string
	Purchase_agent        int
	Listing_agent         int
	Review_agent          int
	Lead_time             int
	Seasonal              int
	Months                string
	Cost                  float64
	Item_location         string
	Discontinued          int
	Remark                string
	Sale_remark           string
	Purchase_remark       string
	Vparcel               float64
	Cgroup                string
	Cost_rmb              float64
	Sizetable             string
	Property_string       string
	Group_info            string
	Purchase_price        float64
	Out_stock_strategy    int
	//Product Product `gorm:"references:productID"`
}

type Product_price struct {
	Id             int `gorm:"primaryKey"`
	ProductID      int `gorm:"column:productID"`
	Price          float64
	Original_price float64
	CurrencyID     int
	Created_at     time.Time
	Updated_at     time.Time
}
