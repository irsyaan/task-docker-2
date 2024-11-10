FROM golang:1.23.2

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./ 

RUN CGD_ENABLED=0 GOOS=linux go build -o /docker-go-ping

EXPOSE 8080

COPY AUTHORS.md /AUTHORS.md

CMD [ "/docker-go-ping"]