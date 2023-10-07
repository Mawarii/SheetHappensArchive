##### DOWLOAD DEPENDENCIES #####
FROM golang:1.21 AS base-build
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOARCH=$TARGETARCH GOOS=linux go mod download

##### BUILD STAGE #####
FROM base-build AS build
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOARCH=$TARGETARCH GOOS=linux go build -o /build/sheethappens main.go

##### RELEASE STAGE #####
FROM scratch
USER 10001
WORKDIR /app
COPY --chown=10001:10001 ./frontend /app/frontend
COPY --chown=10001:10001 --from=build /build/sheethappens /app/sheethappens
EXPOSE 8080
ENTRYPOINT ["/app/sheethappens"]
