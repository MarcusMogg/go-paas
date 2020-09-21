package service

import (
	"errors"
	"paas/global"
	"paas/model/entity"

	"gorm.io/gorm"
)

// InsertCourse 插入数据
func InsertCourse(c *entity.Course, uid uint) error {
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		if err := global.GDB.Create(c).Error; err != nil {
			return err
		}
		cs := entity.CourseStudents{
			CourseID:  c.ID,
			StudentID: uid,
		}
		return tx.Create(&cs).Error
	})
}

// IsCourseTeacher 是否是课程的创建教师id
func IsCourseTeacher(cid, uid uint, tx *gorm.DB) bool {
	var ct entity.Course
	result := tx.Where("id = ? AND teacher_id = ?", cid, uid).First(&ct)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

// InCourse 检查id是否属于课程
func InCourse(cid, uid uint, tx *gorm.DB) bool {
	result := tx.Where("course_id = ? AND student_id = ? AND status = ?", cid, uid, 1).First(&entity.CourseStudents{})
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

// GetCourseByID 通过课程id获取课程信息
func GetCourseByID(id uint) *entity.Course {
	var c entity.Course
	global.GDB.Where("id = ?", id).First(&c)
	return &c
}

// GetMyCourses 通过id获取课程列表
func GetMyCourses(id uint) []entity.Course {
	var c []entity.Course
	global.GDB.Table("courses").Joins("JOIN courses on course_students.course_id = courses.id").
		Where("course_students.student_id = ?", id).Find(&c)
	return c
}

// GetCourses 获取课程列表
func GetCourses(pagenum, pagesize int, keyword string) ([]entity.Course, int64) {
	var c []entity.Course = make([]entity.Course, 0, pagesize)
	var total int64
	offset := (pagenum - 1) * pagesize

	global.GDB.Model(&entity.Course{}).Where("name LIKE ?", "%"+keyword+"%").Count(&total).Offset(offset).Limit(pagesize).Find(&c)
	return c, total
}

// UpdateCourse 修改课程信息
func UpdateCourse(c *entity.Course, user *entity.MUser) error {
	return global.GDB.Save(c).Error
}

// CourseExist 通过课程id判断课程是否存在
func CourseExist(cid uint) error {
	var c entity.Course
	return global.GDB.First(&c, cid).Error
}

// DropCourse 删除课程
func DropCourse(id uint, uid uint) error {
	return global.GDB.Transaction(func(tx *gorm.DB) error {
		var c entity.Course
		c.ID = id

		if err := tx.Delete(&c).Error; err != nil {
			return err
		}
		if err := tx.Where("course_id = ?", id).Delete(&entity.CourseStudents{}).Error; err != nil {
			return err
		}
		return nil
	})
}
