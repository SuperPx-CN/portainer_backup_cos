package cos

import (
	"fmt"
	"github.com/superpx-cn/portainer-backup-cos/internal/config"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

var Client *cos.Client

func SetUp() {
	u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", config.GetCOSBucket(), config.GetCOSRegion()))
	su, _ := url.Parse(fmt.Sprintf("https://cos.%s.myqcloud.com", config.GetCOSRegion()))
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}

	Client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.GetCOSSecretId(),
			SecretKey: config.GetCOSSecretKey(),
		},
	})
}
