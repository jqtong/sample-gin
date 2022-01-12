package card

import (
	"nucarf.com/store_service/api/conf"
)

// card
type Card struct {
	ID        int64  `json:"id" gorm:"column:id"`
	Name      string `json:"name" column:"name"`
	No        string `json:"no" column:"no"`
	CreatedAt int    `json:"created_at" column:"created_at"`
}

// TableName
func (Card) TableName() string {
	return "sp_card"
}

func (card *Card) List(page, size int) (cards []Card, err error) {
	offset := (page - 1) * size
	err = conf.Mysql.Where("status=?", 1).Offset(offset).Limit(size).Order("id desc").Find(&cards).Error
	return
}
