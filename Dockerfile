FROM golang:alpine

RUN apk update && apk add git

ENV APP_PATH /go/src/app
RUN mkdir -p $APP_PATH
WORKDIR $APP_PATH
COPY . $APP_PATH
RUN chown -R daemon:daemon $APP_PATH

RUN go-wrapper download
RUN go-wrapper install
USER daemon
# this will ideally be built by the `go-wrapper` ;)
ENTRYPOINT ["go-wrapper", "run"]
