package initialize

import (
	"paas/global"
	"paas/model/entity"
)

// DBTables 迁移 schema
func DBTables() {
	global.GDB.AutoMigrate(&entity.MUser{})
	global.GDB.AutoMigrate(&entity.Course{})
	global.GDB.AutoMigrate(&entity.CourseStudents{})
	global.GDB.AutoMigrate(&entity.MFilePath{})
	global.GDB.AutoMigrate(&entity.DockerImage{})
	global.GDB.AutoMigrate(&entity.ContainerUser{})
}
