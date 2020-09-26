package service

import (
	"context"
	"encoding/json"
	"fmt"
	"paas/global"
	"paas/model/entity"
	"paas/model/request"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/gin-gonic/gin"
)

// InsertDockerImage 插入镜像
func InsertDockerImage(p *request.DockerPullReq) {
	d := entity.DockerImage{
		CID:  p.ID,
		Name: p.Name,
		Tag:  p.Tag,
	}
	global.GDB.Create(&d)
}

// GetImages 获取镜像列表
func GetImages(cid uint) []entity.DockerImage {
	var res []entity.DockerImage
	global.GDB.Where("c_id = 0 OR c_id = ?", cid).Find(&res)
	return res
}

//DelImages 删除镜像
func DelImages(p *request.DelStudentReq) {
	global.GDB.Where("id = ? AND c_id = ?", p.GetByID.ID, p.CourseIDReq.ID).Delete(&entity.DockerImage{})
}

// ContainerCreate 创建容器
func ContainerCreate(cc *request.ContainerCreateReq) (string, error) {
	ctx := context.Background()
	var netconfig *container.HostConfig = nil
	var useport int = 0
	if cc.Port != 0 {
		useport = global.GPORT
		netconfig = &container.HostConfig{
			PortBindings: nat.PortMap{
				nat.Port(fmt.Sprintf("%d/tcp", cc.Port)): []nat.PortBinding{
					nat.PortBinding{HostPort: fmt.Sprintf("%d/tcp", useport)}},
			},
		}
		global.GPORT++
	}
	resp, err := global.GDOCKER.ContainerCreate(ctx, &container.Config{
		Image: cc.Image,
		Env:   cc.Env,
	}, netconfig, nil, "")
	if err != nil {
		return "", err
	}

	return resp.ID, global.GDB.Create(&entity.ContainerUser{
		ContainerID: resp.ID,
		CID:         cc.CourseIDReq.ID,
		Port:        cc.Port,
		PortBind:    useport,
	}).Error
}

// ContainerStart 容器启动
func ContainerStart(id string) error {
	ctx := context.Background()
	return global.GDOCKER.ContainerStart(ctx, id, types.ContainerStartOptions{})
}

// ContainerRename 容器重命名
func ContainerRename(id, name string) error {
	ctx := context.Background()
	return global.GDOCKER.ContainerRename(ctx, id, name)
}

// ContainerID 获取容器的dockerid
func ContainerID(id uint) string {
	c := entity.ContainerUser{}
	global.GDB.Where("id = ?", id).First(&c)
	return c.ContainerID
}

// ContainerRemove 容器删除
func ContainerRemove(id uint) error {
	ctx := context.Background()
	cid := ContainerID(id)
	//fmt.Println(cid)
	err := global.GDOCKER.ContainerRemove(ctx, cid, types.ContainerRemoveOptions{
		Force: true,
	})
	if err != nil {
		return err
	}
	return global.GDB.Where("id = ?", id).Delete(&entity.ContainerUser{}).Error
}

// ContainerAlloc 将容器分配给某个用户
func ContainerAlloc(id, uid uint) {
	global.GDB.Model(&entity.ContainerUser{}).Where("id = ?", id).Update("u_id", uid)
}

//ContainerStats 获取容器详情
func ContainerStats(id uint) (gin.H, error) {
	ctx := context.Background()
	cid := ContainerID(id)
	out, err := global.GDOCKER.ContainerStats(ctx, cid, false)
	if err != nil {
		return nil, err
	}
	d := json.NewDecoder(out.Body)
	stats := gin.H{}
	d.Decode(&stats)

	// CPU使用率
	systemCPUUsage := (stats["cpu_stats"].(map[string]interface{}))["system_cpu_usage"].(float64)
	cpuUsage := ((stats["cpu_stats"].(map[string]interface{}))["cpu_usage"].(map[string]interface{}))["total_usage"].(float64)
	cpuUtilization := cpuUsage / systemCPUUsage
	memUsage := (stats["memory_stats"].(map[string]interface{}))["usage"].(float64)
	memLimit := (stats["memory_stats"].(map[string]interface{}))["limit"].(float64)

	return gin.H{
		"cpu":      cpuUtilization,
		"memuse":   memUsage,
		"memlimit": memLimit,
	}, nil
}

// GetContainerList 获取课程容器列表
func GetContainerList(uid, cid uint) []entity.ContainerUser {
	var res []entity.ContainerUser
	if IsCourseTeacher(cid, uid, global.GDB) {
		global.GDB.Where("c_id = ?", cid).Find(&res)
	} else {
		global.GDB.Where("c_id = ? and uid = ?", cid, uid).Find(&res)
	}
	return res
}
