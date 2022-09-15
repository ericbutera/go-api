FROM golang:1.19-alpine AS build

LABEL org.opencontainers.image.description go-api demo application

WORKDIR /src/

COPY . ./

RUN go mod download
RUN CGO_ENABLED=0 go build -o /bin/app

FROM scratch
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app", "server"]
