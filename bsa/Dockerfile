# builder
FROM golang:alpine AS build-env
LABEL maintainer="Yannick Foeillet <bzhtux@gmail.com>"

RUN apk --no-cache add build-base git mercurial gcc curl
RUN mkdir -p /go/src/github.com/bzhtux/mgg
ADD . /go/src/github.com/bzhtux/mgg
RUN cd /go/src/github.com/bzhtux/mgg && go get ./... && go build -o mgg-api cmd/main.go


# final image
FROM alpine
LABEL maintainer="Yannick Foeillet <bzhtux@gmail.com>"

RUN apk --no-cache add curl jq
WORKDIR /app
COPY --from=build-env /go/src/github.com/bzhtux/mgg/mgg-api /app/
ENTRYPOINT ./mgg-api
