package entity

import (
	"gorm.io/gorm"
)

// DockerImage 镜像信息
type DockerImage struct {
	gorm.Model
	Name string `json:"name"`
	Tag  string `json:"tag"`
	CID  uint   `json:"cid"` //0 是公共的
}

//ContainerUser 记录所有的id,用户
type ContainerUser struct {
	gorm.Model
	ContainerID string `json:"-"`   // 容器ID
	CID         uint   `json:"-"`   // 课程ID
	UID         uint   `json:"uid"` // 指定的用户ID
	Port        int    `json:"port"`
	PortBind    int    `json:"portbind"`
}
