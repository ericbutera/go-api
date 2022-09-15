FROM golang:1.19-alpine AS build

WORKDIR /src/

COPY . ./

RUN go mod download
RUN CGO_ENABLED=0 go build -o /bin/app

FROM scratch
LABEL org.opencontainers.image.description go-api demo application
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app", "server"]
