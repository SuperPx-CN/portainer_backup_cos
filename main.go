package main

import (
	"github.com/superpx-cn/portainer-backup-cos/internal/backup"
	"github.com/superpx-cn/portainer-backup-cos/internal/config"
	"github.com/superpx-cn/portainer-backup-cos/internal/cos"
	"log"
	"time"
)

func main() {
	config.Load()

	log.Printf("TZ %v", config.GetTZ())
	var err error
	time.Local, err = time.LoadLocation(config.GetTZ())
	if err != nil {
		panic(err)
	}

	cos.SetUp()
	backup.Run()
}
