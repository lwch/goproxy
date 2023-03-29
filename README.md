# goproxy

自建go mod代理服务

## 一键启动

docker方式启动：

```shell
$ docker run -d -p 8080:8080 lwch/goproxy
```

docker-compose方式启动：

```yaml
version: "3"

services:
  goproxy:
    restart: always
    image: lwch/goproxy
    ports:
      - 8080:8080
    volumes:
      - type: volume
        source: goproxy-data
        target: /tmp

volumes:
  goproxy-data:
```

## 使用

将GOPROXY环境变量设置为该代理的请求地址即可，如：

```shell
GOPROXY=http://example:8080,direct
```

## 文件存储

目前仅支持本地文件存储，将来会添加minio方式的文件存储，配置文件中的默认缓存时长为24小时