FROM golang:1.19 AS build

ENV GOPATH /go
WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/otp-cli main.go

RUN strip /go/bin/otp-cli
RUN test -e /go/bin/otp-cli

FROM alpine:latest

LABEL org.opencontainers.image.source=https://github.com/chyroc/otp-cli
LABEL org.opencontainers.image.description="Generate OTP Code Tool."
LABEL org.opencontainers.image.licenses="Apache-2.0"

COPY --from=build /go/bin/otp-cli /bin/otp-cli

CMD /bin/otp-cli