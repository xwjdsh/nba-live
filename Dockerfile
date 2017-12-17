FROM golang:1.9 as builder
WORKDIR /go/src/github.com/xwjdsh/nba-live
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o nba-live

FROM alpine:latest  
LABEL maintainer="iwendellsun@gmail.com"
RUN apk --no-cache add ca-certificates terminus-font
WORKDIR /go/src/github.com/xwjdsh/nba-live
COPY --from=builder /go/src/github.com/xwjdsh/nba-live/nba-live .
COPY --from=builder /go/src/github.com/xwjdsh/nba-live/vendor/github.com/common-nighthawk/go-figure/fonts/slant.flf ./vendor/github.com/common-nighthawk/go-figure/fonts/slant.flf
CMD ["./nba-live"]
