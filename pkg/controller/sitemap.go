package controller

import (
	"context"
)

type SiteMapController interface {
	GetSiteMapUrls(ctx context.Context, path string) (urls []string)
	SendData(ctx context.Context, site, token, data string)
}
