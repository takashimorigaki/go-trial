FROM golang:latest

# working directory
ENV APP_ROOT /go-trial/api
WORKDIR $APP_ROOT

EXPOSE 3000