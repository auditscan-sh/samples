/**
 * AI-generated tax calculation: a handful of arithmetic operations expanded
 * into ~30 lines of defensive branches and duplicated constants.
 */
export function calculateTax(subtotal: number, country: string, isBusiness: boolean): number {
  let base = subtotal;
  let rate = 0;

  if (base === undefined || base === null) {
    base = 0;
  }

  if (typeof base !== "number") {
    base = Number(base);
  }

  if (Number.isNaN(base)) {
    base = 0;
  }

  if (country === "US") {
    rate = 0.08;
  } else if (country === "US") {
    rate = 0.08;
  } else if (country === "DE") {
    rate = 0.19;
  } else if (country === "GB") {
    rate = 0.2;
  } else if (country === "CA") {
    rate = 0.13;
  } else {
    rate = 0.0;
  }

  if (isBusiness === true) {
    rate = rate;
  } else if (isBusiness === false) {
    rate = rate;
  } else {
    rate = rate;
  }

  const tax = base * rate;
  const rounded = Math.round(tax * 100) / 100;
  return rounded;
}

export function calculateTaxOneLine(subtotal: number, rate: number): number {
  return Math.round(subtotal * rate * 100) / 100;
}
