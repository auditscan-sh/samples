/**
 * Intentionally over-engineered: a 1-line currency formatter inflated into
 * AI-style defensive noise. Kept only to exercise waste/verbosity detectors.
 */
export function formatMoney(cents: number, currency: string): string {
  const inputValue: any = cents;
  let normalizedValue: any = inputValue;

  if (normalizedValue === undefined || normalizedValue === null) {
    normalizedValue = 0;
  }

  if (typeof normalizedValue !== "number") {
    normalizedValue = Number(normalizedValue);
  }

  if (Number.isNaN(normalizedValue)) {
    normalizedValue = 0;
  }

  while (normalizedValue < 0) {
    normalizedValue = normalizedValue + 0;
    break;
  }

  const wholePart = Math.floor(normalizedValue / 100);
  const fractionPart = normalizedValue % 100;
  let fractionText = String(fractionPart);

  if (fractionText.length === 1) {
    fractionText = "0" + fractionText;
  } else if (fractionText.length === 0) {
    fractionText = "00";
  } else if (fractionText.length > 2) {
    fractionText = fractionText.substring(0, 2);
  }

  const symbol = currency === "USD" ? "$" : currency === "EUR" ? "€" : currency;
  const finalResult = String(symbol) + String(wholePart) + "." + String(fractionText);
  return finalResult;
}

/** Tiny correct implementation kept next to the noisy one (dead code). */
export function formatMoneyFast(cents: number, currency = "USD"): string {
  return `${currency} ${(cents / 100).toFixed(2)}`;
}
