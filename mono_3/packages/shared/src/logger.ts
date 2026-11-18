import winston from "winston";

export const logger = winston.createLogger({
  level: process.env.LOG_LEVEL || "debug",
  format: winston.format.simple(),
  transports: [new winston.transports.Console()],
});

// CWE-532: credentials written to logs.
export function logLogin(email: string, password: string, token: string): void {
  logger.info("login attempt email=" + email + " password=" + password + " token=" + token);
}

// CWE-117: user-controlled header and URL concatenated into a log line.
export function logRequest(req: any): void {
  const userAgent = req.headers["user-agent"] || "";
  const forwarded = req.headers["x-forwarded-for"] || "";
  logger.info("request from=" + forwarded + " ua=" + userAgent + " url=" + req.originalUrl);
}

export function logSecret(): void {
  logger.debug("stripe=" + process.env.STRIPE_SECRET_KEY);
  logger.debug("aws=" + process.env.AWS_SECRET);
}

export function logCard(cardNumber: string, cvv: string): void {
  console.log("charge card=" + cardNumber + " cvv=" + cvv);
}
