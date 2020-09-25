package service

import (
	"errors"
	"paas/global"
	"paas/model/entity"

	"gorm.io/gorm"
)

// InsertStudent 添加学生
func InsertStudent(cid uint, name string, status uint) (uint, error) {
	var uid uint
	err := global.GDB.Transaction(func(tx *gorm.DB) error {
		tx.Model(&entity.MUser{}).Select("id").Where("user_name = ?", name).Scan(&uid)
		if uid == 0 {
			return errors.New("查无此人")
		}
		cs := entity.CourseStudents{
			CourseID:  cid,
			StudentID: uid,
		}
		return tx.Create(&cs).Error
	})
	return uid, err
}

// GetStudents 获取学生列表
func GetStudents(cid uint) []entity.MUser {
	var users []entity.MUser
	global.GDB.Table("m_users").Joins("JOIN course_students ON m_users.id = course_students.student_id").
		Where("course_students.course_id = ?", cid).Find(&users)
	return users
}

// DeleteStudent 删除学生
func DeleteStudent(uid, cid uint) error {
	var id uint
	global.GDB.Model(&entity.Course{}).Select("teacher_id").Where("id = ?", cid).Scan(&id)
	if id == uid {
		return errors.New("不能删除老师")
	}
	global.GDB.Where("student_id = ? AND course_id = ?", uid, cid).Delete(&entity.CourseStudents{})
	return nil
}
