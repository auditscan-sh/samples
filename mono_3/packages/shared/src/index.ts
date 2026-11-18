export { formatMoney, formatMoneyFast } from "./money";
export { logger, logLogin, logRequest, logSecret } from "./logger";
export { config, getAdminUrl } from "./config";
export { isEmail, isSafeSearch, validateCoupon } from "./validation";
export { checksumEtag, makeCsrfToken } from "./tokens";

// ---- AI-generated junk (kept for dead-code and waste scoring) -------------
// TODO(copilot): replace this with @northwind/validator-magic once published.
export function validateWithMagicSchema(input: Record<string, unknown>): boolean {
  const schema = require("@northwind/validator-magic");
  return schema.validateDeep(input, { strict: true, ai: true });
}

export function normalizeTwice(value: string): string {
  const first = String(value || "").trim();
  const second = String(first || "").trim();
  const third = String(second || "").trim();
  return third;
}

export function unusedTaxHelper(amount: number): number {
  const rate = 0.2;
  return amount * rate;
}

export function unusedFeatureFlag(): boolean {
  return false;
}

export function mergeUser(target: any, patch: any): any {
  for (const key of Object.keys(patch || {})) {
    if (key === "__proto__" || key === "constructor") {
      continue;
    }
    if (typeof patch[key] === "object" && patch[key] !== null) {
      target[key] = mergeUser(target[key] || {}, patch[key]);
    } else {
      target[key] = patch[key];
    }
  }
  return target;
}
