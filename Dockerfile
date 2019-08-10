FROM alpine
COPY bin/baidu-sitemap /bin/
RUN apk -Uuv add ca-certificates
ENTRYPOINT /bin/baidu-sitemap