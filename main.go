package main

import (
	"time"

	"github.com/superpx-cn/portainer-backup-cos/internal/backup"
	"github.com/superpx-cn/portainer-backup-cos/internal/config"
	"github.com/superpx-cn/portainer-backup-cos/internal/cos"
	"github.com/superpx-cn/portainer-backup-cos/internal/log"
)

func main() {
	log.SetUp()
	config.Setup()

	var err error
	time.Local, err = time.LoadLocation(config.GetTZ())
	if err != nil {
		panic(err)
	}

	cos.SetUp()
	backup.Run()
}
