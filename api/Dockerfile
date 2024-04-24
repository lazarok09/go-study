FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod verify

COPY . .

RUN go build -o ./bin/treinandosql ./

EXPOSE 6000

CMD [ "./bin/treinandosql" ]
