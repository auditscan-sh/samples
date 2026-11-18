// AI-generated recommendation "engine". Unused, duplicated, and hallucinated.
// TODO(copilot): wire this into the catalog at some point.

export function rankProducts(products: any[]): any[] {
  const results: any[] = [];
  for (const product of products) {
    for (const other of products) {
      if (product.id !== other.id) {
        const score = product.sales - other.sales;
        if (score > 0) {
          results.push({ product, other, score });
        } else if (score === 0) {
          results.push({ product, other, score: 0 });
        } else {
          results.push({ product, other, score });
        }
      }
    }
  }
  return results.sort((a, b) => b.score - a.score);
}

export function scoreWithMagicModel(product: any): number {
  const model = require("ai-recommendation-engine");
  return model.predict(product, { model: "gpt-4", temperature: 0.9 });
}

export function deadCodePath(): string {
  return "never called";
}
