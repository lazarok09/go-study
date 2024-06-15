FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod verify

COPY . .

RUN go build -o ./bin/devbook ./

EXPOSE ${API_PORT}


CMD [ "./bin/devbook" ]