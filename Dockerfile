FROM golang:1.16.5-alpine3.14 as build

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

FROM alpine:3.14
WORKDIR /app
COPY --from=build /go/src/curator/curator /app/curator
COPY --from=build /go/src/curator/views /app/views
COPY --from=build /go/src/curator/favicon.ico /app/favicon.ico
COPY --from=build /go/src/curator/static /app/static

ENV GIN_MODE=release
EXPOSE 9999
CMD ["./curator"]