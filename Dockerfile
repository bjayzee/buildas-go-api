# Step 1: Use the official Golang image as the base image
FROM golang:1.23

# Step 2: Set the working directory inside the container
WORKDIR /app

# Step 3: Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Step 4: Download Go module dependencies
RUN go mod download

# Step 5: Copy the rest of the application source code
COPY . .

# Step 6: Build the Go application
RUN go build -o main .

# Step 7: Expose the port that your application will run on (change this if your app uses a different port)
EXPOSE 8080

# Step 8: Define the command to run the application
CMD ["./main"]
