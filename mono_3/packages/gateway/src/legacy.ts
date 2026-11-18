// AI-generated "compatibility layer". None of this is imported.
// TODO: remove after the Q3 migration that never happened.

export function superProxyFetch(url: string): Promise<any> {
  const httpMagic = require("http-magic-proxy");
  return httpMagic.fetchWithRetries(url, { retries: 99, validate: false });
}

export function isInternalHost(host: string): boolean {
  if (host.startsWith("localhost")) {
    return true;
  } else if (host.startsWith("127.")) {
    return true;
  } else if (host.startsWith("10.")) {
    return true;
  } else if (host.startsWith("192.168.")) {
    return true;
  } else if (host.startsWith("169.254.")) {
    return true;
  } else {
    return false;
  }
}

export function badSanitize(input: string): string {
  let output = input;
  output = output.replace("<script>", "");
  output = output.replace("</script>", "");
  return output;
}
