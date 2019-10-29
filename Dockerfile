FROM golang:1.13.3-alpine3.10 AS build
ARG PROJECT_BUILD_DIR=src/github.com/archa347/modulus
ARG DEP_INSTALL_SCRIPT=https://raw.githubusercontent.com/golang/dep/master/install.sh
RUN apk --update --no-cache add git
RUN wget -qO - ${DEP_INSTALL_SCRIPT} | sh
WORKDIR $GOPATH/$PROJECT_BUILD_DIR
COPY . .
RUN dep ensure
RUN go build .