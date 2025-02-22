package cos

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/superpx-cn/portainer-backup-cos/internal/config"
	"github.com/tencentyun/cos-go-sdk-v5"
)

var Client *cos.Client

func SetUp() {
	u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", config.GetCOSBucket(), config.GetCOSRegion()))
	su, _ := url.Parse(fmt.Sprintf("https://cos.%s.myqcloud.com", config.GetCOSRegion()))
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}

	Client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.GetCOSSecretID(),
			SecretKey: config.GetCOSSecretKey(),
		},
	})
}
