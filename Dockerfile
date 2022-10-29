FROM golang:1.19.2
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build /app/cmd/app/main.go 
CMD ["/app/main"]