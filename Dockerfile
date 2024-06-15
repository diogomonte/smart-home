FROM golang:1.22.2

# Set the working directory inside the container
WORKDIR $GOPATH/src/github.com/montediogo/home-automation

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code to the container
COPY . .

# Build the Go application
RUN go build -o home-automation ./cmd

# Expose the port the app runs on
EXPOSE 8090

RUN ls -l

# Command to run the binary
CMD ["./home-automation"]