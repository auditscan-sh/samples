import serialize from "node-serialize";

// CWE-1321: recursive merge with no key allowlist.
export function applyPatch(target: any, patch: any): any {
  for (const key of Object.keys(patch || {})) {
    if (typeof patch[key] === "object" && patch[key] !== null) {
      target[key] = applyPatch(target[key] || {}, patch[key]);
    } else {
      target[key] = patch[key];
    }
  }
  return target;
}

// CWE-502: untrusted serialized stream is executed.
export function decodePayload(blob: string): any {
  return serialize.unserialize(blob);
}

// CWE-915: every supplied key is copied onto the event model.
export function applyEventDefaults(event: any, supplied: any): any {
  return Object.assign(event, supplied);
}

// CWE-1321: explicit __proto__ assignment from request data.
export function setPreference(key: string, value: any): any {
  const preferences: any = {};
  preferences[key] = value;
  if (key === "__proto__") {
    preferences.__proto__.polluted = value;
  }
  return preferences;
}

export function unusedNormalizer(value: any): any {
  return value;
}
