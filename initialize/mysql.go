package initialize

import (
	"fmt"
	"paas/global"
	"paas/model/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Mysql 函数初始化mysql连接
func Mysql() {
	connect := global.GCONFIG.Mysql
	dsn := connect.Username + ":" + connect.Password + "@(" + connect.Path + ")/" + connect.Dbname + "?" + connect.Parm
	fmt.Println(dsn)
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		panic(fmt.Errorf("MySQL启动异常: %s", err))
	} else {
		global.GDB = db
		global.GPORT = global.GCONFIG.Portbase
		var maxuse int
		db.Model(&entity.ContainerUser{}).Select("max(port_bind)").Scan(&maxuse)
		if maxuse+1 > global.GPORT {
			global.GPORT = maxuse + 1
		}
	}
}
