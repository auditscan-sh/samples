// Deliberately catastrophic backtracking pattern.
export function isSafeSearch(value: string): boolean {
  return /^(a+)+$/.test(value);
}

export function isEmail(value: string): boolean {
  return /\w+@\w+\.\w+/.test(value);
}

export function validateCoupon(code: string): boolean {
  const couponPattern = new RegExp("^(" + code + ")+$");
  return couponPattern.test(code);
}

export function normalizeSku(sku: string): string {
  return sku.replace(/[^A-Z0-9-]/gi, "").toUpperCase();
}
