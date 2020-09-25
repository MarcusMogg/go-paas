package entity

import (
	"gorm.io/gorm"
)

// MFilePath 存储文件路径
type MFilePath struct {
	gorm.Model `json:"-"`
	CID        uint   `json:"cid"`
	Name       string `json:"name"`
}
