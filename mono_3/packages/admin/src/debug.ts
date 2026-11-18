import fs from "fs";
import v8 from "v8";
import inspector from "inspector";

// CWE-200: every environment variable (including keys) is returned.
export function dumpEnvironment(): Record<string, string | undefined> {
  return { ...process.env };
}

// CWE-73: heap snapshot path is caller-controlled.
export function writeHeap(target: string): void {
  const snapshot = v8.writeHeapSnapshot(target);
  console.log("wrote", snapshot);
}

// CWE-215: debugger opened by an HTTP route.
export function openDebugger(): void {
  inspector.open(9229, "0.0.0.0", true);
}

export function readSecretFile(): string {
  return fs.readFileSync("/proc/self/environ", "utf8");
}
