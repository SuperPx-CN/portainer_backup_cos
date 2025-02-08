package portainer

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/superpx-cn/portainer-backup-cos/internal/config"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

func Backup() string {
	url := fmt.Sprintf("%s/api/backup", config.GetPortainerURL())

	body := map[string]string{
		"password": "",
	}

	bodyJson, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyJson))
	if err != nil {
		log.Println("创建请求失败:", err)
		return ""
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", config.GetPortainerToken())

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("请求失败：", err)
		return ""
	}

	if resp.StatusCode != 200 {
		log.Println("请求失败：", resp)
		return ""
	}

	defer resp.Body.Close()

	if resp.Header.Get("Content-Type") != "application/x-gzip" {
		log.Printf("启动的 Content-Type 是 application/x-gzip，返回的确却是 %v", resp.Header.Get("Content-Type"))
		return ""
	}

	gzipReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		fmt.Println("Error creating gzip reader:", err)
		return ""
	}

	defer gzipReader.Close()

	tempDir := os.TempDir() // 获取系统临时目录

	contentDisposition := resp.Header.Get("Content-Disposition")
	re := regexp.MustCompile(`filename=([^\s;]+)`)
	matches := re.FindStringSubmatch(contentDisposition)
	var filename string
	if len(matches) > 1 {
		filename = matches[1]
	}

	tempFile, err := os.Create(filepath.Join(tempDir, filename))

	if err != nil {
		log.Println("Error creating temp tempFile:", err)
		return ""
	}

	defer tempFile.Close()

	_, err = io.Copy(tempFile, gzipReader)
	if err != nil {
		fmt.Println("Error writing to temp file:", err)
		return ""
	}

	return filepath.Join(tempDir, filename)
}
