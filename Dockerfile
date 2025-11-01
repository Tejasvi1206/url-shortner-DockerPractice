FROM golang:1.25-alpine


WORKDIR /app

COPY . .


RUN go build -o url-shortener


EXPOSE 3000


CMD ["./url-shortener"]
