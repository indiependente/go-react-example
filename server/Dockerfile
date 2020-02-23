# Builder image
FROM golang:1.14-rc-alpine as builder
WORKDIR $GOPATH/src/github.com/indiependente/go-react-example
COPY go.mod .
COPY go.sum .
ENV GOPRIVATE=gitb.com/indiependente
RUN go mod download && go get github.com/GeertJohan/go.rice && go get github.com/GeertJohan/go.rice/rice
ADD ./client/dist ./client/dist
RUN rice embed-go
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /svc ./server/
RUN rm rice-box.go

# Deployable image
FROM gcr.io/distroless/static
WORKDIR /app
COPY --from=builder /svc /app/
EXPOSE 8000
ENTRYPOINT ["/app/svc"]
