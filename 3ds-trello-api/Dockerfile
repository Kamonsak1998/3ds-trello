FROM golang:1.13

ARG APP_ENV
ENV APP_ENV $APP_ENV

WORKDIR /go/src/app/

COPY . .

RUN GO111MODULE=on go mod tidy && GO111MODULE=on go get
RUN go build -o main .

CMD ["./main"]

