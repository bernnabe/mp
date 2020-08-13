# Start from golang base image
FROM golang:stretch as build

# Set the current working directory inside the container 
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download 
COPY . .

# Build the Go app
RUN go build -o main .

EXPOSE 8080

### Put the binary onto Heroku image
FROM heroku/heroku:18
COPY --from=build /app/main .
CMD ["./main"]
