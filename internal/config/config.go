package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	TZ             string
	BackupInterval time.Duration
	BackupLimit    int
	COSBucket      string
	COSRegion      string
	COSSecretID    string
	COSSecretKey   string
	PortainerURL   string
	PortainerToken string
}

var cfg Config

func Load() {
	// init cfg
	cfg = Config{
		TZ:             getEnv("TZ", "UTC"),
		COSSecretID:    os.Getenv("COS_SECRET_ID"),
		COSSecretKey:   os.Getenv("COS_SECRET_KEY"),
		PortainerURL:   getEnv("PORTAINER_URL", "http://127.0.0.1:9000"),
		PortainerToken: os.Getenv("PORTAINER_TOKEN"),
	}
	loadBackupInterval()
	loadBackupLimit()
	loadCOSBucket()
	loadCOSRegion()
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}

func GetTZ() string {
	return cfg.TZ
}
func GetBackupLimit() int {
	return cfg.BackupLimit
}

func GetBackupInterval() time.Duration {
	return cfg.BackupInterval
}

func GetPortainerURL() string {
	return cfg.PortainerURL
}
func GetPortainerToken() string {
	return cfg.PortainerToken
}

func GetCOSBucket() string {
	return cfg.COSBucket
}

func GetCOSRegion() string {
	return cfg.COSRegion
}

func GetCOSSecretId() string {
	return cfg.COSSecretID
}
func GetCOSSecretKey() string {
	return cfg.COSSecretKey
}

func loadCOSRegion() {
	cfg.COSRegion = os.Getenv("COS_REGION")
}
func loadBackupLimit() {
	cfg.BackupLimit, _ = strconv.Atoi(os.Getenv("BACKUP_LIMIT"))
}
func loadCOSBucket() {
	cfg.COSBucket = os.Getenv("COS_BUCKET")
}
func loadBackupInterval() {
	str := getEnv("BACKUP_INTERVAL", "10m")

	duration, err := time.ParseDuration(str)
	if err != nil {
		duration = time.Minute * 10
		log.Fatalf("解析 BACKUP_INTERVAL 失败: %v,使用默认值 %v", err, duration)
	}

	cfg.BackupInterval = duration
}
