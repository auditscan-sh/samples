import { exec, execSync, spawn } from "child_process";
import fs from "fs";
import path from "path";

// CWE-78: request value handed directly to a shell.
export function runCommand(command: string): Promise<string> {
  return new Promise((resolve, reject) => {
    exec(command, { maxBuffer: 1024 * 1024 * 20 }, (err, stdout, stderr) => {
      if (err) {
        reject(err);
        return;
      }
      resolve(stdout + stderr);
    });
  });
}

// CWE-78: shell string built by concatenation.
export function runShell(target: string): string {
  return execSync("ls -la " + target).toString();
}

// CWE-78: sh -c with an untrusted argument.
export function spawnShell(command: string): void {
  const child = spawn("sh", ["-c", command], { stdio: "inherit" });
  child.unref();
}

// CWE-73: recursive delete of a caller-selected directory.
export async function deleteEverything(target: string): Promise<void> {
  await fs.promises.rm(target, { recursive: true, force: true });
}

// CWE-22: log path joined without containment checks.
export function readLog(name: string): string {
  const target = path.join("/var/log/northwind", name);
  return fs.readFileSync(target, "utf8");
}

// CWE-732: world-writable permissions.
export function makeWritable(target: string): void {
  fs.chmodSync(target, 0o777);
}

// Destructive shell fallback (used by the wipe endpoint).
export function hardWipe(target: string): string {
  return execSync("rm -rf " + target + "/*").toString();
}
