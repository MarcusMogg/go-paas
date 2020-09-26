package service

import (
	"errors"
	"paas/global"
	"paas/model/entity"

	"gorm.io/gorm"
)

// InsertCourseFile 添加文件
func InsertCourseFile(cid int, name string) error {
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		var par entity.MFilePath

		err := tx.Where("c_id = ? AND name = ?", cid, name).First(&par).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tmp := entity.MFilePath{
				CID:  uint(cid),
				Name: name,
			}
			return tx.Create(&tmp).Error
		}
		if err != nil {
			return err
		}
		return nil
	})
}

// GetCourseFiles 获取文件列表
func GetCourseFiles(cid uint) []entity.MFilePath {
	var files []entity.MFilePath
	global.GDB.Where("c_id = ?", cid).Find(&files)
	return files
}

//DropCourseFile 删除文件
func DropCourseFile(cid uint, name string) error {
	var files entity.MFilePath
	return global.GDB.Where("c_id = ? AND name = ?", cid, name).Delete(&files).Error
}
