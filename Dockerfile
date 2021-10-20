FROM golang:1.16-buster AS builder
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download -x

COPY . ./
ARG SKAFFOLD_GO_GCFLAGS
RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -ldflags="-X 'main.release=`git describe --tags --abbrev=0`'" -o /bin/go-app cmd/server/*.go

FROM gcr.io/distroless/base
ENV GOTRACEBACK=single
WORKDIR /app

COPY --from=builder /bin/go-app ./
EXPOSE 8080

CMD ["./go-app"]
