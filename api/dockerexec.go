package api

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"paas/global"
	"paas/service"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/docker/docker/api/types"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Terminal websocket 连接容器
func Terminal(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("container"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	container := service.ContainerID(uint(id))
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	hr, err := exec(container, c.Query("workdir"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer hr.Close()
	defer func() {
		hr.Conn.Write([]byte("exit\r"))
	}()

	go func() {
		wsWriterCopy(hr.Conn, conn)
	}()
	wsReaderCopy(conn, hr.Conn)
}

func exec(container string, workdir string) (hr types.HijackedResponse, err error) {
	ctx := context.Background()
	ir, err := global.GDOCKER.ContainerExecCreate(ctx, container, types.ExecConfig{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		WorkingDir:   workdir,
		Cmd:          []string{"/bin/bash"},
		Tty:          true,
	})
	if err != nil {
		return
	}

	hr, err = global.GDOCKER.ContainerExecAttach(ctx, ir.ID, types.ExecStartCheck{Detach: false, Tty: true})
	if err != nil {
		return
	}
	return
}

func wsWriterCopy(reader io.Reader, writer *websocket.Conn) {
	buf := make([]byte, 8192)
	for {
		nr, err := reader.Read(buf)
		if nr > 0 {
			err := writer.WriteMessage(websocket.BinaryMessage, buf[0:nr])
			if err != nil {
				return
			}
		}
		if err != nil {
			return
		}
	}
}

func wsReaderCopy(reader *websocket.Conn, writer io.Writer) {
	for {
		messageType, p, err := reader.ReadMessage()
		if err != nil {
			return
		}
		if messageType == websocket.TextMessage {
			writer.Write(p)
		}
	}
}
