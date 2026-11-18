import axios from "axios";
import { exec, execSync } from "child_process";

// CWE-918: arbitrary URL fetch from a request parameter.
export async function fetchPreview(url: string): Promise<string> {
  const response = await axios.get(url, { timeout: 5000 });
  return String(response.data).slice(0, 2000);
}

// CWE-78: request-controlled filename interpolated into a shell command.
export function generateThumbnail(file: string): Promise<string> {
  return new Promise((resolve, reject) => {
    exec("convert " + file + " -resize 200x200 /srv/uploads/thumb.png", (err, stdout, stderr) => {
      if (err) {
        reject(err);
        return;
      }
      resolve(stdout || stderr);
    });
  });
}

// CWE-78: synchronous variant used by the legacy queue.
export function convertLegacy(file: string, format: string): string {
  return execSync("ffmpeg -i " + file + " -f " + format + " /srv/uploads/out." + format).toString();
}

// CWE-918: fixed internal metadata URL, no allowlist.
export async function fetchCloudMetadata(): Promise<string> {
  const response = await axios.get("http://169.254.169.254/latest/meta-data/");
  return String(response.data);
}
