#!/bin/bash

# 컨테이너 이름
CONTAINER_NAME="gorizon"

# 호스트 디렉토리와 컨테이너의 마운트 경로
HOST_DIR=~/dev/go/gorizon
CONTAINER_DIR=/gorizon
PORT=8080

# 이미지 이름
IMAGE_NAME="golang:1.19"

# 컨테이너 실행
docker run --name $CONTAINER_NAME -it -d -p $PORT:$PORT -v $HOST_DIR:$CONTAINER_DIR $IMAGE_NAME /bin/bash