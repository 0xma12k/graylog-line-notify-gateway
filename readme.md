# Graylog Webhook with Line Notify

Graylog HTTP notification with Line Notify

## Installtions

### Precompiled binaries
Precompiled binaries for released versions are available in the  [release page.](https://github.com/0xma12k/graylog-line-notify-gateway/releases)

1. Download binaries to local

```
$ wget https://github.com/0xma12k/graylog-line-notify-gateway/releases/download/v0.0.1/graylog-line-notify-gateway_0.0.1_linux_amd64.tar.gz
```

2. Extract binaries from tar file 

```
$ tar -xvf graylog-line-notify-gateway_0.0.1_linux_amd64.tar.gz
```

3. Move binaries to `/usr/local/bin`

```
$ mv graylog-line-notify-gateway /usr/local/bin/
```


### Docker images

Docker images are available on [Docker Hub](https://hub.docker.com/r/0xma12k/graylog-line-notify-gateway)

```
$ docker run --rm -d -p 3000:3000 -v $(pwd)/config.yml:/app/config.yml  0xma12k/graylog-line-notify-gateway:0.0.1 --config config.yml
```

## Configurations

Example of configuration file

```yml
server_port: 3000
log_level: debug

default_template: default

templates:
  - name: default
    path: templates/default.templ
```

Running gateway with the following command.

```
$ graylog-line-notify-gateway --config config.yml
```

## Usages

Gateway is a http server is listenning for Graylog HTTP notification at path `/line`, for example

```
localhost:3000/line?token={line_notify_token}&template={template_name}
```

Gateway use go [text/template](https://pkg.go.dev/text/template) package for generate message and parse [Graylog's object](https://github.com/0xma12k/graylog-line-notify-gateway/blob/main/internal/entity/graylog.go) to.