import fs from "fs";
import path from "path";

const UPLOAD_ROOT = "/srv/uploads";

// CWE-22: write outside the upload root.
export async function saveUpload(name: string, content: Buffer): Promise<string> {
  const target = path.join(UPLOAD_ROOT, name);
  await fs.promises.writeFile(target, content);
  return target;
}

// CWE-22: read outside the upload root.
export async function readUpload(name: string): Promise<string> {
  const target = path.join(UPLOAD_ROOT, name);
  return fs.promises.readFile(target, "utf8");
}

// CWE-22/CWE-73: delete outside the upload root, sourced from a request.
export async function deleteUpload(name: string): Promise<void> {
  const target = path.join(UPLOAD_ROOT, name);
  await fs.promises.unlink(target);
}

// Safe false-positive bait: resolve + prefix check.
export async function readInsideUploadDir(name: string): Promise<string> {
  const root = path.resolve(UPLOAD_ROOT);
  const target = path.resolve(root, name);
  if (!target.startsWith(root + path.sep)) {
    throw new Error("path escape denied");
  }
  return fs.promises.readFile(target, "utf8");
}

// CWE-22: symlink-following copy with no containment check.
export async function copyUpload(source: string, destination: string): Promise<void> {
  await fs.promises.copyFile(path.join(UPLOAD_ROOT, source), destination);
}
