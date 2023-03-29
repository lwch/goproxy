FROM golang AS build

COPY go.mod go.sum main.go /build/
COPY internal /build/internal/

WORKDIR /build

RUN go build

FROM golang

COPY conf/goproxy.yaml /etc/goproxy.yaml
COPY --from=build /build/goproxy /usr/bin/goproxy

ENTRYPOINT [ "/usr/bin/goproxy", "-conf", "/etc/goproxy.yaml" ]