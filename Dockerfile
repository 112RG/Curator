FROM golang:1.21-alpine as build

WORKDIR /go/src/curator
RUN apk update && apk add automake build-base
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG GIT_COMMIT=unspecified
ARG BUILD_DATE=unspecified

LABEL git_commit=$GIT_COMMIT
LABEL build_date=$BUILD_DATE

ENV GIT_COMMIT=$GIT_COMMIT
ENV BUILD_DATE=$BUILD_DATE

RUN go build -o curator .

FROM alpine
WORKDIR /app
COPY --from=flyio/litefs:0.5 /usr/local/bin/litefs /usr/local/bin/litefs
COPY --from=build /go/src/curator/docker-config/etc/litefs.static-lease.yml /tmp
COPY --from=build /go/src/curator/docker-config/litefs.yml /etc
COPY --from=build /go/src/curator/curator /app/curator
COPY --from=build /go/src/curator/views /app/views
COPY --from=build /go/src/curator/favicon.ico /app/favicon.ico
COPY --from=build /go/src/curator/static /app/static


RUN apk add ca-certificates fuse3 sqlite

ENV GIN_MODE=release
ENTRYPOINT litefs mount