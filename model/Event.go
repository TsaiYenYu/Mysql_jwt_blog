package model

import (
	"time"

	"github.com/wejectchen/ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Event struct {
	ID         uint      `gorm:"primary_key;auto_increment"  json:"id"`
	Name       string    `gorm:"type:varchar(20);not null" json:"name"`
	Host       string    `gorm:"type:varchar(20);not null" json:"host"`
	When       string    `gorm:"type:varchar(100);not null" json:"when"`
	Who        string    `gorm:"type:varchar(100);not null" json:"who"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

// 查询分类是否存在
func CheckEvent(name string) (code int) {
	var event Event
	db.Select("id").Where("name = ?", name).First(&event)
	//沒有找到是0
	if event.ID > 0 {
		return errmsg.ERROR_CATENAME_USED //2001
	}
	return errmsg.SUCCSE
}

// CreateCate 新增分类
func CreateEvent(data *Event) int {
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
func GetEventInfo(id int) (Event, int) {
	var event Event
	db.Where("id = ?", id).First(&event)
	return event, errmsg.SUCCSE
}

// GetCate 查询分类列表
func GetEvent(pageSize int, pageNum int) ([]Event, int64) {
	var event []Event
	var total int64
	err = db.Find(&event).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	db.Model(&event).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return event, total
}

// EditCate 编辑分类信息
func EditEvent(id int, data *Event) int {
	var event Event
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	maps["updated_at"] = now
	err = db.Model(&event).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// DeleteCate 删除分类
func DeleteEvent(id int) int {
	var event Event
	err = db.Where("id = ? ", id).Delete(&event).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
