FROM golang:1.12.5-alpine3.9 AS build

ENV CGO_ENABLED=0

WORKDIR /go/src/github.com/stackrox/default-authz-plugin
COPY . .
RUN go build -o /default-authz-plugin .

FROM scratch
COPY --from=build /default-authz-plugin /

ENTRYPOINT ["/default-authz-plugin"]
