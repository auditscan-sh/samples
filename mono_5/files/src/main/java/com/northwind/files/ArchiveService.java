package com.northwind.files;

import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.util.zip.ZipEntry;
import java.util.zip.ZipInputStream;
import org.springframework.stereotype.Service;

@Service
public class ArchiveService {

    // CWE-22 (Zip Slip): entry names are joined without containment checks.
    public void extract(String archivePath, String destination) throws Exception {
        byte[] buffer = new byte[8192];
        try (ZipInputStream zip = new ZipInputStream(new FileInputStream(archivePath))) {
            ZipEntry entry;
            while ((entry = zip.getNextEntry()) != null) {
                File target = new File(destination, entry.getName());
                if (entry.isDirectory()) {
                    target.mkdirs();
                    continue;
                }
                new File(target.getParent()).mkdirs();
                try (FileOutputStream out = new FileOutputStream(target)) {
                    int length;
                    while ((length = zip.read(buffer)) > 0) {
                        out.write(buffer, 0, length);
                    }
                }
            }
        }
    }

    // CWE-22: single entry extraction with no check.
    public void extractEntry(String archivePath, String entryName, String destination) throws Exception {
        try (ZipInputStream zip = new ZipInputStream(new FileInputStream(archivePath))) {
            ZipEntry entry;
            while ((entry = zip.getNextEntry()) != null) {
                if (entry.getName().equals(entryName)) {
                    File target = new File(destination, entry.getName());
                    try (FileOutputStream out = new FileOutputStream(target)) {
                        out.write(zip.readAllBytes());
                    }
                }
            }
        }
    }

    // Intentional resource leak (no close).
    public int countEntries(String archivePath) throws Exception {
        ZipInputStream zip = new ZipInputStream(new FileInputStream(archivePath));
        int count = 0;
        while (zip.getNextEntry() != null) {
            count++;
        }
        return count;
    }
}
