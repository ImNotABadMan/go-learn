package tables

import (
	"time"
)

// 蛇形命名，c_id => cID
type PublicationTask struct {
	ID          int
	SiteID      int
	AdminID     int
	State       int
	PulledAt    time.Time `gorm:"default:null"`
	PullState   int
	DeletedAt   time.Time `gorm:"default:null"`
	DeleteState int
	CreatedAt   time.Time           `gorm:"default:null;autoCreateTime"`
	UpdatedAt   time.Time           `gorm:"default:null;autoUpdateTime"`
	PuProduct   PublicationProducts `gorm:"foreignKey:PublicationID;references:ID"`
}

type PublicationProducts struct {
	PublicationID  int
	ProductID      int `gorm:"column:productID"`
	ThirdProductID string
	Name           string
	ProductCode    string
	CurrencyID     int
	Enabled        int `gorm:"default:0"`
	IsTaxable      int
	IsLogistics    int
	IsGroup        int
	CreatedAt      time.Time               `gorm:"default:null"`
	UpdatedAt      time.Time               `gorm:"default:null"`
	PubProductSub  []PublicationProductSub `gorm:"foreignKey:PublicationID;references:PublicationID"`
}

type PublicationProductSub struct {
	ID                 int
	PublicationID      int
	CoProductID        int    `gorm:"column:co_productID"`
	CoProductCode      string `gorm:"column:co_product_code"`
	CoThirdProductId   string `gorm:"column:co_third_product_id"`
	CoThirdInventoryId string `gorm:"column:co_third_inventory_id"`
	Price              float64
	OriginalPrice      float64
	InStock            int
	OutStockStrategy   int
	CreatedAt          time.Time `gorm:"default:null"`
	UpdatedAt          time.Time `gorm:"default:null"`
}

type PublicationProductCategory struct {
	Id             int
	PublicationID  int
	SiteCategoryID int
	Remarks        string
	CreatedAt      time.Time `gorm:"default:null"`
	UpdatedAt      time.Time `gorm:"default:null"`
	ThirdCollectID string
}

type SiteCategory struct {
	Id              int
	SiteID          int
	ThirdCategoryID string
	Title           string
	Type            string
	Published       int
	LocalParent     int
	Remarks         string
	CreatedAt       time.Time `gorm:"default:null"`
	UpdatedAt       time.Time `gorm:"default:null"`
}

type SiteConfig struct {
	Id          int
	Name        string
	Url         string
	Description string
	Platform    string
	ApiID       string
	ApiKey      string
	ApiPassword string
	SecretKey   string
	AccessToken string
	LocationIds string
	AdminID     int
	CreatedAt   time.Time `gorm:"default:null"`
	UpdatedAt   time.Time `gorm:"default:null"`
}
