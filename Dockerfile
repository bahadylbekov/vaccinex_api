FROM golang:1.12.7-alpine3.10
RUN mkdir /app
COPY . /app
WORKDIR /app

RUN apk add git
RUN apk add --update make
RUN go mod download

RUN make build

CMD ["./vaccinex_api"]