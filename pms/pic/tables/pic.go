package tables

import (
	"gorm.io/gorm"
	"strconv"
)

type Picture struct {
	PhotoID        int64 `gorm:"primaryKey;column:photoID"`
	ProductID      int64 `gorm:"column:productID"`
	Filename       string
	Thumbnail      string
	Enlarged       string
	Smallthumbnail string
	PhotoOrder     int8
	Newthumbnail   string
	Image_HD       string `gorm:"column:image_HD"`
	Mobile_image   string
	Origin_image   string
}

func GetAll(db *gorm.DB, limit int64, productID int64) []Picture {
	var pics []Picture
	if productID == 0 {
		productID = 56414
	}
	db.Raw("select * from ss_product_pictures where productID < ? order by productID desc limit "+strconv.FormatInt(limit, 10), productID).Scan(&pics)

	return pics
}

func CountPic(db *gorm.DB, productID int64) int64 {
	var count int64
	db.Raw("select count(*) from ss_product_pictures where productID < ? ", productID).Scan(&count)

	return count
}
