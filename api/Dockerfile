FROM golang:1.15.2-alpine3.12 as builder
COPY . /go/src
WORKDIR /go/src
RUN go build -o /app main.go

FROM alpine:3.12
COPY --from=builder /app /bin/
EXPOSE 80
CMD ["/bin/app"]
