package tables

import (
	"time"
)

type Products struct {
	//gorm.Model
	ProductID       int `gorm:"primaryKey;column:productID"`
	CategoryID      int `gorm:"column:categoryID"`
	Name            string
	NameCn          string
	ProductCode     string
	ProductCodeType int
	GroupParentSku  string
	Isgroup         int
	GroupParent     int
	PlatformID      int
	InStock         int
	Enabled         int
	SortOrder       int
	DefaultPicture  int
	ViewedTimes     int
	Onhold          int
	ItemsSold       int
	Weight          float64
	Cubic           float64
	JustImported    int
	AddedAdmin      string
	ModifiedAdmin   int
	InStockPhotoID  int
	InStockPhotoUrl string
	DateAdded       time.Time
	DateModified    time.Time
	FormalDeletedAt time.Time
	DeletedAt       time.Time
	Additional      ProductAdditional `gorm:"foreignKey:ProductID;references:ProductID"`
	Prices          []ProductPrice    `gorm:"foreignKey:ProductID;references:ProductID"`
}

type ProductAdditional struct {
	//gorm.Model
	ProductID           int `gorm:"primaryKey;column:productID"`
	OriginalProductCode string
	SupplierName        string
	SupplierType        string
	SupplierUrl         string
	PurchaseAgent       int
	ListingAgent        int
	DeveloperAgent      int
	LeadTime            int
	Seasonal            int
	Months              string
	Cost                float64
	ItemLocation        string
	Discontinued        int
	Remark              string
	SaleRemark          string
	PurchaseRemark      string
	Vparcel             float64
	Cgroup              string
	CostRmb             float64
	Sizetable           string
	PropertyString      string
	GroupInfo           string
	PurchasePrice       float64
	OutStockStrategy    int
	//Product Product `gorm:"references:productID"`
}

type ProductPrice struct {
	ID            int `gorm:"primaryKey"`
	ProductID     int `gorm:"column:productID"`
	Price         float64
	OriginalPrice float64
	CurrencyID    int `gorm:"column:currencyID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
