FROM golang:1.16.5-alpine3.14

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
RUN apk del automake build-base
ENV GIN_MODE=release
EXPOSE 9999
CMD ["./curator"]