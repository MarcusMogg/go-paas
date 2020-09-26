package request

// DockerPullReq 从docker hub下载镜像
type DockerPullReq struct {
	CourseIDReq
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

// ContainerCreateReq 创建镜像时所需的
type ContainerCreateReq struct {
	CourseIDReq
	Image string   `json:"image"`
	Env   []string `json:"env"`
	Port  int      `json:"port"`
	Num   int      `json:"num"`
}

//ContainerAllocReq 将容器分配给用户
type ContainerAllocReq struct {
	CourseIDReq
	ContainerID uint `json:"containerid"` // 容器ID
	UID         uint `json:"uid"`         // 指定的用户ID
}

//ContainerDelReq 容器删除
type ContainerDelReq struct {
	CourseIDReq
	ContainerID uint `json:"containerid"` // 容器ID
}
