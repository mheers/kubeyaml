FROM golang:1.19-alpine
WORKDIR /kubeyaml
ADD go.mod .
ADD go.sum .
RUN go mod download
ADD . /kubeyaml
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o kubeyaml ./cmd/kubeyaml

FROM alpine:3.17
ARG USER=default
ENV HOME /home/$USER

USER $USER
WORKDIR $HOME

COPY --from=0 /kubeyaml/kubeyaml ./kubeyaml
ENTRYPOINT [ "./kubeyaml" ]
