// GraphQL surface with introspection enabled and no query-cost controls.
export function buildSchema(): any {
  return {
    introspection: true,
    playground: true,
    execute(query: string): any {
      // CWE-94: request query concatenated into an executable expression.
      const fn = new Function("context", "return " + query);
      return fn({});
    },
    depthLimit: 0,
    complexityLimit: 0,
  };
}

// Safe false-positive bait: no request input.
export function resolvePublicProduct(): any {
  return { id: "demo", name: "Sample product" };
}
