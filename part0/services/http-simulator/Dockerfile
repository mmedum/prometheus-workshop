FROM golang:1.13-alpine as builder

LABEL maintainer="Mark Medum Bundgaard <mmedum@gmail.com>"

# To fix go get and build with cgo
RUN apk add --no-cache --virtual .build-deps \
    bash \
    gcc \
    git \
    musl-dev

RUN mkdir build
COPY . /build
WORKDIR /build

RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o httpsimulator .
RUN adduser -S -D -H -h /build httpsimulator
USER httpsimulator

FROM scratch
COPY --from=builder /build/httpsimulator /app/
WORKDIR /app
EXPOSE 8080
CMD ["./httpsimulator"]
