FROM golang:alpine

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

RUN mkdir /app
WORKDIR /app

COPY . .
RUN go mod download
RUN go mod tidy

# Build the Go app
#RUN go build -o main .
RUN go run .

# Expose port 9000
EXPOSE 9000
# Run the executable
# CMD [ "/main" ]

#Setup hot-reload for dev environment
RUN go install -mod=mod github.com/githubnemo/CompileDaemon
RUN go install -mod=mod golang.org/x/tools/gopls@latest
# ENTRYPOINT CompileDaemon --build="go build -a -installsuffix cgo -o main ." --command=./main
#ENTRYPOINT CompileDaemon --command="go run ." --command=./main
ENTRYPOINT ["/go/bin/CompileDaemon", ...]