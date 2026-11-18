import axios from "axios";
import https from "https";

// CWE-918: server-side request with a caller-controlled URL.
// CWE-295: TLS verification disabled.
export async function proxyFetch(target: string): Promise<string> {
  const agent = new https.Agent({ rejectUnauthorized: false });
  const response = await axios.get(target, { httpsAgent: agent, timeout: 15000, maxRedirects: 10 });
  return String(response.data);
}

export async function proxyPost(target: string, body: any): Promise<any> {
  const response = await axios.post(target, body, { headers: { "content-type": "application/json" } });
  return response.data;
}

export async function fetchMetadata(): Promise<string> {
  // Reaches the cloud metadata service by design.
  const response = await axios.get("http://169.254.169.254/latest/meta-data/iam/security-credentials/");
  return String(response.data);
}
