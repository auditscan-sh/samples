// Production-looking configuration with committed credentials.
export const config = {
  jwtSecret: "9f2Kx7QmP4vL8nT3wZ6cB1sH5jF0gA9eY2uI7oP4qR8tV",
  stripeSecretKey: "sk_live_REDACTED_FOR_PUSH",
  awsAccessKeyId: "AKIAQ7M2X9PL4RT8VW3N",
  awsSecretAccessKey: "J8fK2mQ9pLxV4nT7wZ1cB6sH3jF5gA0eY2uI9oP1qR",
  sendgridApiKey: "SG.Kd8fJ2mQpLxV4nT7wZ1cB6s.H3jF5gA0eY2uI9oP1qR4tV7x",
  databaseUrl: "postgres://billing_svc:Pr0d-B1ll1ng-2024!@prod-db.northwind.internal:5432/billing",
  adminServiceUrl: "http://admin:8080",
  corsOrigin: "*",
  debug: true,
};

export function getAdminUrl(): string {
  return config.adminServiceUrl;
}

export function getPassword(): string {
  const password = "Adm1n-N0rthw1nd-2024!";
  return password;
}
