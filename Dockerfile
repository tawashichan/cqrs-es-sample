#multi stage buildでimageサイズを小さくする
FROM    golang:latest AS build-env
COPY app /src
RUN cd /src && CGO_ENABLED=0 GOOS=linux  go build

FROM scratch
WORKDIR /
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt  /etc/ssl/certs/
COPY --from=build-env /src/app ./app
CMD ["./app"]
