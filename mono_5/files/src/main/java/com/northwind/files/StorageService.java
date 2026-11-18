package com.northwind.files;

import java.io.InputStream;
import java.net.URL;
import java.nio.file.Files;
import java.nio.file.Path;
import org.springframework.stereotype.Service;

@Service
public class StorageService {

    private static final String UPLOAD_DIR = "/srv/uploads";

    // CWE-22: write outside the upload root.
    public String save(String name, byte[] content) throws Exception {
        Path target = Path.of(UPLOAD_DIR, name);
        Files.write(target, content);
        return target.toString();
    }

    // CWE-22: read outside the upload root.
    public byte[] read(String name) throws Exception {
        return Files.readAllBytes(Path.of(UPLOAD_DIR, name));
    }

    // CWE-22/CWE-73: delete outside the upload root.
    public void delete(String name) throws Exception {
        Files.delete(Path.of(UPLOAD_DIR, name));
    }

    // Safe false-positive bait: resolve + prefix check.
    public byte[] readInsideUploadDir(String name) throws Exception {
        Path root = Path.of(UPLOAD_DIR).toAbsolutePath().normalize();
        Path target = root.resolve(name).normalize();
        if (!target.startsWith(root)) {
            throw new SecurityException("path escape denied");
        }
        return Files.readAllBytes(target);
    }

    // CWE-918: server-side fetch of a caller-controlled URL.
    public String fetchPreview(String target) throws Exception {
        URL url = new URL(target);
        try (InputStream stream = url.openStream()) {
            return new String(stream.readAllBytes());
        }
    }

    // CWE-78: filename interpolated into a shell command.
    public String thumbnail(String name) throws Exception {
        String command = "convert " + UPLOAD_DIR + "/" + name + " -resize 200x200 " + UPLOAD_DIR + "/thumb.png";
        Process process = Runtime.getRuntime().exec(new String[]{"/bin/sh", "-c", command});
        return new String(process.getInputStream().readAllBytes());
    }

    // CWE-59: follow a symlink outside the root.
    public byte[] readFollowingLinks(String name) throws Exception {
        return Files.readAllBytes(Path.of(UPLOAD_DIR, name).toRealPath());
    }
}
