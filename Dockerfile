FROM golang:1.19
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod tidy
COPY . ./
RUN go build -o /server
CMD [ "/server" ]