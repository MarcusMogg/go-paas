package initialize

import (
	"paas/global"

	"github.com/docker/docker/client"
)

const version = "1.38"

// Docker 初始化docker client
func Docker() {
	cli, err := client.NewClientWithOpts(client.WithHost(global.GCONFIG.DockerHost), client.WithVersion(version))
	if err != nil {
		panic(err)
	}
	global.GDOCKER = cli
}
