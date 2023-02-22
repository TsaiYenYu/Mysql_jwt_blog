package model

import (
	"time"

	"github.com/wejectchen/ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Location struct {
	ID         uint      `gorm:"primary_key;auto_increment"  json:"id"`
	Name       string    `gorm:"type:varchar(20);not null" json:"name"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

// CheckCategory 查询分类是否存在
func CheckLocation(name string) (code int) {
	var location Location
	db.Select("id").Where("name = ?", name).First(&location)
	//沒有找到是0
	if location.ID > 0 {
		return errmsg.ERROR_CATENAME_USED //2001
	}
	return errmsg.SUCCSE
}

// CreateCate 新增分类
func CreateLocation(data *Location) int {

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
func GetLocationInfo(id int) (Location, int) {
	var location Location
	db.Where("id = ?", id).First(&location)
	return location, errmsg.SUCCSE
}

// GetCate 查询分类列表
func GetLocation(pageSize int, pageNum int) ([]Location, int64) {
	var location []Location
	var total int64
	err = db.Find(&location).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	db.Model(&location).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return location, total
}

// EditCate 编辑分类信息
func EditLocation(id int, data *Location) int {
	var location Location
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	maps["updated_at"] = now
	err = db.Model(&location).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// DeleteCate 删除分类
func DeleteLocation(id int) int {
	var location Location
	err = db.Where("id = ? ", id).Delete(&location).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
