
all: baidu-sitemap

baidu-sitemap:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/baidu-sitemap cmd/main.go


