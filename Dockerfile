FROM golang:1.17 as Build

LABEL maintainer="xylonx"

WORKDIR /build
COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o main -ldflags="-w -s" main.go

FROM xylonx/cn-ubuntu:latest as Prod

ARG PROJECT_NAME=icc-core

WORKDIR /opt/${PROJECT_NAME}
COPY --from=Build /build/main ./${PROJECT_NAME}

RUN echo "./${PROJECT_NAME} -c=config.yaml" >>start.sh && \
    chmod 755 -R /opt/${PROJECT_NAME}

CMD ./start.sh