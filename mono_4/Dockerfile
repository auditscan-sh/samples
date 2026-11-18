# Intentionally bad container build.
FROM golang:latest AS build
WORKDIR /src
COPY . .
RUN go work sync && CGO_ENABLED=0 go build -o /out/gateway ./services/gateway && \
    CGO_ENABLED=0 go build -o /out/auth ./services/auth && \
    CGO_ENABLED=0 go build -o /out/orders ./services/orders && \
    CGO_ENABLED=0 go build -o /out/payments ./services/payments && \
    CGO_ENABLED=0 go build -o /out/files ./services/files && \
    CGO_ENABLED=0 go build -o /out/admin ./services/admin && \
    CGO_ENABLED=0 go build -o /out/worker ./services/worker

FROM alpine:latest
USER root
RUN apk add --no-cache curl bash netcat-openbsd sudo
ENV JWT_SECRET=4Qm9pLxV2nT7wZ1cB6sH3jF5gA0eY8uI9oP2qR4tV7x
ENV DATABASE_URL=postgres://billing_svc:Pr0d-B1ll1ng-2024!@prod-db.northwind.internal:5432/billing
COPY --from=build /out/gateway /usr/local/bin/gateway
COPY --from=build /out/auth /usr/local/bin/auth
COPY --from=build /out/orders /usr/local/bin/orders
COPY --from=build /out/payments /usr/local/bin/payments
COPY --from=build /out/files /usr/local/bin/files
COPY --from=build /out/admin /usr/local/bin/admin
COPY --from=build /out/worker /usr/local/bin/worker
COPY --from=build /out/worker /usr/local/bin/billing-worker
COPY --from=build /out/worker /usr/local/bin/legacy-sync
EXPOSE 8080 22 6060
CMD ["gateway"]
