package controller

import (
	"context"
)

// SiteMapController 百度 SiteMap 控制器
type SiteMapController interface {
	GetSiteMapUrls(ctx context.Context, path string) (urls []string)
	SendData(ctx context.Context, site, token, data string)
}
