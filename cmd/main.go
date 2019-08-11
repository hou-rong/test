package main

import (
	"../pkg/controller/impl"
	"context"
	"os"
	"strings"
)

func main() {
	ctx := context.Background()

	site := os.Getenv("PLUGIN_SITE")
	token := os.Getenv("PLUGIN_TOKEN")
	path := os.Getenv("PLUGIN_PATH")

	SiteMapController := impl.DefaultSiteMapController
	urls := SiteMapController.GetSiteMapUrls(ctx, path)
	SiteMapController.SendData(ctx, site, token, strings.Join(urls, "\n"))
}
