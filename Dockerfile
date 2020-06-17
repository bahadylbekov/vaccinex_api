FROM golang:1.12.7-alpine3.10
RUN mkdir /app
COPY . /app
WORKDIR /app

RUN go mod download

RUN make build

CMD ["/app/vaccinex_api"]