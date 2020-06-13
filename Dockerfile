# build image
FROM golang:latest AS build-env
ENV GOPATH=/go \
    GOBIN=$GOPATH/bin \
    CGO_ENABLED=0
WORKDIR /app
COPY ./ ./
RUN go get && go build -o serve.out

# execute image
FROM gcr.io/distroless/base
WORKDIR /app

COPY --from=build-env /app/serve.out .
ENV TZ=Asia/Tokyo

WORKDIR /share/misc
WORKDIR /app

EXPOSE 8080
CMD ["./serve.out"]
