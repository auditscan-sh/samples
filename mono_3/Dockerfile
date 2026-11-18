FROM node:latest

# Intentionally bad: root user, latest base, secrets baked into image layers.
USER root
WORKDIR /app

ENV NODE_ENV=production \
    JWT_SECRET=9f2Kx7QmP4vL8nT3wZ6cB1sH5jF0gA9eY2uI7oP4qR8tV \
    STRIPE_SECRET_KEY=sk_live_REDACTED_FOR_PUSH \
    DATABASE_URL=postgres://billing_svc:Pr0d-B1ll1ng-2024!@prod-db.northwind.internal:5432/billing

RUN apt-get update && apt-get install -y curl netcat-openbsd vim && rm -rf /var/lib/apt/lists/*
RUN npm install -g npm@latest pm2

COPY . .
RUN npm install --production

EXPOSE 22 3000 9229
CMD ["node", "packages/gateway/dist/server.js"]
