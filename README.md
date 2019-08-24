# drone-baidu-sitemap

## 简介

用于更新 baidu site 的 drone 插件

## 原理

baidu 要求的 sitemap.xml 与标准的 sitemap.xml 不同。通过解析普通的 sitemap.xml，调用 baidu 的「主动提交」接口更新连接

![](https://img.henryhou.com/img/20190824223417.png)

## Docker 调用方法

```bash
docker run --rm \
           -e PLUGIN_SITE=host \   # site 例如 henryhou.com
           -e PLUGIN_TOKEN=token \ # baidu token
           -e PLUGIN_PATH=./public/sitemap.xml \
           -v $(pwd):$(pwd) \
           -w $(pwd) henryhou009/drone-baidu-sitemap
```

## Drone 配置

```yaml
kind: pipeline
name: default

steps:
  - name: baidu-sitemap
    image: henryhou009/drone-baidu-sitemap:linux-amd64
    settings:
      site: henryhou.com
      token: 
        from_secret: baidu_sitemap_token
      path: ./public/sitemap.xml
```