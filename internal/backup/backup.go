package backup

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/superpx-cn/portainer-backup-cos/internal/config"
	"github.com/superpx-cn/portainer-backup-cos/internal/cos"
	"github.com/superpx-cn/portainer-backup-cos/internal/portainer"
	TencentCOS "github.com/tencentyun/cos-go-sdk-v5"
)

func Run() {
	for {
		localPath := portainer.Backup()

		if localPath != "" {
			remotePath := storeBackup(localPath)

			log.Printf("备份成功 %s", remotePath)
			_ = os.Remove(localPath)
		}
		cleanBackups()

		time.Sleep(config.GetBackupInterval())
	}
}

func cleanBackups() {
	if config.GetBackupLimit() <= 0 {
		return
	}

	opt := &TencentCOS.BucketGetOptions{Prefix: "portainer-backup_"}

	res, _, err := cos.Client.Bucket.Get(context.Background(), opt)
	if err != nil {
		log.Println("获取备份列表失败：", err)
		return
	}

	var keys []string
	for _, content := range res.Contents {
		keys = append(keys, content.Key)
	}

	if len(keys) <= config.GetBackupLimit() {
		return
	}

	sort.Strings(keys)
	for _, key := range keys[:len(keys)-config.GetBackupLimit()] {
		_, err := cos.Client.Object.Delete(context.Background(), key)
		if err != nil {
			log.Println("删除旧备份失败：", err)
		}
	}

}

func storeBackup(filePath string) string {
	key := filepath.Join(config.GetFolder(), filepath.Base(filePath))

	_, _, err := cos.Client.Object.Upload(context.Background(), key, filePath, nil)
	if err != nil {
		panic(err)
	}

	return key
}
