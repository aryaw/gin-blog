FROM golang:alpine

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

RUN mkdir /app
WORKDIR /app

COPY . .
COPY .env .

RUN go get -d -v ./...

RUN go install -v ./...

# Build the Go app
#RUN go build -o /build
# Expose port 8080
#EXPOSE 8080
# Run the executable
#CMD [ "/build" ]

#Setup hot-reload for dev environment
RUN go get github.com/githubnemo/CompileDaemon
RUN go get -v golang.org/x/tools/gopls

ENTRYPOINT CompileDaemon --build="go build -a -installsuffix cgo -o main ." --command=./main