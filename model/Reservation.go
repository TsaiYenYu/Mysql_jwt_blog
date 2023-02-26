package model

import (
	"time"

	"github.com/wejectchen/ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Reservation struct {
	ID           uint      `gorm:"primary_key;auto_increment"  json:"id"`
	Date         string    `gorm:"type:varchar(45);not null" json:"date"`
	Name         string    `gorm:"type:varchar(45);not null" json:"name"`
	Charger      string    `gorm:"type:varchar(45);not null" json:"charger"`
	Event        string    `gorm:"type:varchar(45);not null" json:"event"`
	Time         string    `gorm:"type:varchar(45);not null" json:"time"`
	Invited_list string    `gorm:"type:varchar(45);not null" json:"invited_list"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
}

// CheckCategory 查询分类是否存在
func CheckReservation(name string) (code int) {
	var reservation Reservation
	db.Select("id").Where("name = ?", name).First(&reservation)
	//沒有找到是0
	if reservation.ID > 0 {
		return errmsg.ERROR_CATENAME_USED //2001
	}
	return errmsg.SUCCSE
}

// CreateCate 新增分类
func CreateReservation(data *Reservation) int {
	dataCreated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	data.Created_at = dataCreated_at
	dataUpdated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	data.Updated_at = dataUpdated_at
	result := db.Create(&data)
	if result.Error != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

// GetCateInfo 查询单个分类信息
func GetReservationInfo(id int) (Reservation, int) {
	var reservation Reservation
	db.Where("id = ?", id).First(&reservation)
	return reservation, errmsg.SUCCSE
}

// GetCate 查询分类列表
func GetReservation(pageSize int, pageNum int) ([]Reservation, int64) {
	var reservation []Reservation
	var total int64
	err = db.Find(&reservation).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	db.Model(&reservation).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return reservation, total
}

// EditCate 编辑分类信息
func EditReservation(id int, data *Reservation) int {
	var reservation Reservation
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	maps["charger"] = data.Charger
	maps["event"] = data.Event
	maps["time"] = data.Time
	maps["invited_list"] = data.Invited_list
	now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	maps["updated_at"] = now
	err = db.Model(&reservation).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// DeleteCate 删除分类
func DeleteReservation(id int) int {
	var reservation Reservation
	err = db.Where("id = ? ", id).Delete(&reservation).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
