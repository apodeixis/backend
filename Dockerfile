FROM golang1.19:alpine3.18 as buildbase

WORKDIR /go/src/github.com/apodeixis/backend

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/backend main.go

FROM alpine:3.18

COPY --from=buildbase /usr/local/bin/backend /usr/local/bin/backend
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["backend"]
