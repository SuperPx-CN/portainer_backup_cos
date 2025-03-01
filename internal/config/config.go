package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	BackupInterval time.Duration
	BackupLimit    int
	COSBucket      string
	COSRegion      string
	COSSecretID    string
	COSSecretKey   string
	Folder         string
	PortainerToken string
	PortainerURL   string
	TZ             string
}

var cfg Config

func Setup() {

	cfg = Config{
		BackupInterval: parseBackupInterval(),
		BackupLimit:    getIntEnv("BACKUP_LIMIT", 4321),
		COSBucket:      getStringEnv("COS_BUCKET", "bucket-name"),
		COSRegion:      getStringEnv("COS_REGION", "ap-guangzhou"),
		COSSecretID:    getStringEnv("COS_SECRET_ID", ""),
		COSSecretKey:   getStringEnv("COS_SECRET_KEY", ""),
		Folder:         getStringEnv("FOLDER", "/"),
		PortainerToken: getStringEnv("PORTAINER_TOKEN", ""),
		PortainerURL:   getStringEnv("PORTAINER_URL", "http://127.0.0.1:9000"),
		TZ:             getStringEnv("TZ", "UTC"),
	}
}

func getStringEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}

func getIntEnv(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	result, _ := strconv.Atoi(value)
	return result
}

func parseBackupInterval() time.Duration {
	str := getStringEnv("BACKUP_INTERVAL", "10m")
	duration, err := time.ParseDuration(str)
	if err != nil {
		duration = time.Minute * 10
	}

	return duration
}

func GetBackupInterval() time.Duration {
	return cfg.BackupInterval
}
func GetBackupLimit() int {
	return cfg.BackupLimit
}
func GetCOSBucket() string {
	return cfg.COSBucket
}
func GetCOSRegion() string {
	return cfg.COSRegion
}
func GetCOSSecretID() string {
	return cfg.COSSecretID
}
func GetCOSSecretKey() string {
	return cfg.COSSecretKey
}
func GetFolder() string {
	return cfg.Folder
}
func GetPortainerToken() string {
	return cfg.PortainerToken
}
func GetPortainerURL() string {
	return cfg.PortainerURL
}
func GetTZ() string {
	return cfg.TZ
}
