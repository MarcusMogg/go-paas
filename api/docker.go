package api

import (
	"context"
	"fmt"
	"io"
	"os"
	"paas/global"
	"paas/model/entity"
	"paas/model/request"
	"paas/model/response"
	"paas/service"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
)

// DockerPull 从docker hub拉去镜像
func DockerPull(c *gin.Context) {
	var p request.DockerPullReq
	if err := c.BindJSON(&p); err == nil {
		ctx := context.Background()
		if p.Tag == "" {
			p.Tag = "latest"
		}
		out, err := global.GDOCKER.ImagePull(ctx, fmt.Sprintf("%s:%s", p.Name, p.Tag), types.ImagePullOptions{})
		if err != nil {
			response.FailWithMessage("镜像不存在", c)
		}
		defer out.Close()

		io.Copy(os.Stdout, out)
		response.Ok(c)
	} else {
		response.FailValidate(c)
	}
}

// GetDockerImages 获取镜像列表
func GetDockerImages(c *gin.Context) {
	var id request.CourseIDReq
	if err := c.BindJSON(&id); err == nil {
		response.OkWithData(service.GetImages(id.ID), c)
	} else {
		response.FailValidate(c)
	}
}

// DelDockerImage 删除镜像
func DelDockerImage(c *gin.Context) {
	var id request.DelStudentReq
	if err := c.BindJSON(&id); err == nil {
		service.DelImages(&id)
		response.Ok(c)
	} else {
		response.FailValidate(c)
	}
}

// CreateContainer 批量创建容器
func CreateContainer(c *gin.Context) {
	var cc request.ContainerCreateReq
	if err := c.BindJSON(&cc); err == nil {
		res := 0
		for cc.Num > 0 {
			id, err := service.ContainerCreate(&cc)
			if err == nil {
				res++
				service.ContainerRename(id, fmt.Sprintf("%d-%s-%d", cc.CourseIDReq.ID, id[0:6], time.Now().Unix()))
				service.ContainerStart(id)
			} else {
				fmt.Println(err.Error())
			}
			cc.Num--
		}
		response.OkWithData(res, c)
	} else {
		response.FailValidate(c)
	}
}

// GetContainerList 获取容器列表
func GetContainerList(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	var id request.CourseIDReq
	if err := c.BindJSON(&id); err == nil {
		res := service.GetContainerList(user.ID, id.ID)
		response.OkWithData(res, c)
	} else {
		response.FailValidate(c)
	}
}

// GetContainer 获取容器详情
func GetContainer(c *gin.Context) {
	var id request.GetByID
	if err := c.BindJSON(&id); err == nil {
		res, err := service.ContainerStats(id.ID)
		if err != nil {
			response.Fail(c)
		} else {
			response.OkWithData(res, c)
		}
	} else {
		response.FailValidate(c)
	}
}

// AllocContainer 分配容器
func AllocContainer(c *gin.Context) {
	var id request.ContainerAllocReq
	if err := c.BindJSON(&id); err == nil {
		service.ContainerAlloc(id.ContainerID, id.UID)
		response.Ok(c)
	} else {
		response.FailValidate(c)
	}
}

// DropAllContainer 删除容器
func DropAllContainer(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	user := claim.(*entity.MUser)
	var id request.CourseIDReq
	if err := c.BindJSON(&id); err == nil {
		res := service.GetContainerList(user.ID, id.ID)
		for _, i := range res {
			fmt.Println(i.ID)
			err := service.ContainerRemove(i.ID)
			if err != nil {
				response.FailWithMessage(err.Error(), c)
				return
			}
		}
		response.Ok(c)
	} else {
		response.FailValidate(c)
	}
}

// DropContainer 删除容器
func DropContainer(c *gin.Context) {
	var id request.ContainerDelReq
	if err := c.BindJSON(&id); err == nil {
		err := service.ContainerRemove(id.ContainerID)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.Ok(c)
	} else {
		response.FailValidate(c)
	}
}
