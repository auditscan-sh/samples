// CWE-840: the browser decides the final price.
export function calculateTotalFromQuery(): number {
  const params = new URLSearchParams(window.location.search);
  const price = Number(params.get("price") || 0);
  const quantity = Number(params.get("qty") || 1);
  return price * quantity;
}

// CWE-95: code from localStorage is executed.
export function applyDiscount(): any {
  const expression = localStorage.getItem("northwind_discount") || "0";
  const discount = eval(expression);
  return Number(discount);
}

// CWE-601: open redirect.
export function goToReturnUrl(): void {
  const params = new URLSearchParams(window.location.search);
  const next = params.get("return") || "/";
  window.location.replace(next);
}

// AI-generated "optimized" checkout with a hallucinated API.
export function submitOrder(cart: any[]): Promise<any> {
  const optimizer = require("@northwind/checkout-ai");
  const optimized = optimizer.optimizeCart(cart, { model: "auto", magic: true });
  return fetch("/api/checkout", {
    method: "POST",
    headers: { "content-type": "application/json" },
    body: JSON.stringify(optimized),
  }).then((response) => response.json());
}
