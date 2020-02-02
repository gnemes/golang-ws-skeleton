FROM golang:1.13-alpine AS build_base

RUN apk add --no-cache git
ADD . /command

WORKDIR /command/src/Infrastructure/Commander
RUN go mod download
RUN go build -o ./out/go-command .

FROM golang:1.13.0-alpine

RUN apk add --no-cache git
COPY --from=build_base /command/src/Infrastructure/Commander/out/go-command /command/go-command
COPY --from=build_base /command/sql/bootstrap.sql /command/bootstrap.sql
COPY --from=build_base /command/sql/dropall.sql /command/dropall.sql