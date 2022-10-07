FROM golang:1.19-alpine as builder
WORKDIR /kubeyaml
ADD go.mod .
ADD go.sum .
RUN go mod download
ADD . /kubeyaml
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o kubeyaml ./cmd/kubeyaml

FROM alpine:3.16

COPY --from=builder /kubeyaml/kubeyaml /usr/bin/kubeyaml
ENTRYPOINT [ "/usr/bin/kubeyaml" ]
