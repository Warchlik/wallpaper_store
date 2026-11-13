FROM golang:1.27-alpine AS build
WORKDIR /src

COPY pkg ./pkg
COPY services/gateway ./services/gateway

RUN go work init ./pkg ./services/gateway

WORKDIR /src/services/gateway
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /bin/gateway ./cmd/server

FROM gcr.io/distroless/static-debian12:nonroot AS runtime
COPY --from=build /bin/gateway /gateway
ENTRYPOINT ["/gateway"]
