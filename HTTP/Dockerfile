FROM golang:1.21

WORKDIR /app

COPY go.mod .



RUN go mod download && go mod verify

COPY . ./

RUN go build -o ./out/estudandohttp .
EXPOSE 5001
ENTRYPOINT ["./out/estudandohttp"]