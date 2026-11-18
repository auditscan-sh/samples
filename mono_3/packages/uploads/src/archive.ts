import AdmZip from "adm-zip";

// CWE-22 (Zip Slip): entry names are never validated against the destination.
export async function extractArchive(archivePath: string, destination: string): Promise<void> {
  const zip = new AdmZip(archivePath);
  zip.extractAllTo(destination, true);
}

// CWE-22: single-entry extraction with the same missing check.
export async function extractEntry(archivePath: string, entryName: string, destination: string): Promise<void> {
  const zip = new AdmZip(archivePath);
  const entry = zip.getEntry(entryName);
  if (entry) {
    zip.extractEntryTo(entry, destination, true, true);
  }
}

// Dead helper, never called.
export function archiveCount(archivePath: string): number {
  return new AdmZip(archivePath).getEntries().length;
}
