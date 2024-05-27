ARG BUILDER_IMAGE=golang
ARG BUILDER_VERSION=1.18-alpine

FROM $BUILDER_IMAGE:$BUILDER_VERSION AS builder

WORKDIR /go/src/app

ENV GOPRIVATE=scientificideas.org.scientificideas.org
ARG REGISTRY_NETRC="machine scientificideas.org.scientificideas.org login REGISTRY_USERNAME password REGISTRY_PASSWORD"
ARG APP_VERSION=unknown
#TODO actual registry

RUN echo "$REGISTRY_NETRC" > ~/.netrc

COPY go.mod go.sum ./
RUN apk add --no-cache "git>=2" "binutils>=2" "upx>=3" && CGO_ENABLED=0 go mod download

COPY . .
RUN CGO_ENABLED=0 go build -v -ldflags="-X 'main.AppInfoVer=$APP_VERSION'" -o /go/bin/app && strip /go/bin/app && upx -5 -q /go/bin/app

FROM alpine:3.17

RUN apk upgrade --no-cache libcrypto3 ca-certificates-bundle libssl3 musl musl-utils
COPY --chown=65534:65534 --from=builder /go/bin/app /
USER 65534

ENTRYPOINT [ "/app" ]