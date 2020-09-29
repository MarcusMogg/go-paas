package global

import (
	"paas/config"

	"github.com/docker/docker/client"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	// GCONFIG 全局配置内容
	GCONFIG config.Config
	// GVP 读取配置
	GVP *viper.Viper
	// GDB 数据库连接
	GDB *gorm.DB
	//GDOCKER docker连接
	GDOCKER *client.Client
	// GPORT 未使用的端口
	GPORT int
)

// TimeTemplateDay 时间转换模板，到天
const TimeTemplateDay = "2006-01-02"

// TimeTemplateSec 时间转换模板，到秒
const TimeTemplateSec = "2006-01-02 15:04:05"
