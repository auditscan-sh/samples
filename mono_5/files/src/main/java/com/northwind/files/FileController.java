package com.northwind.files;

import java.util.Map;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.multipart.MultipartFile;

@RestController
@RequestMapping("/files")
public class FileController {

    @Autowired
    private StorageService storageService;

    @Autowired
    private ArchiveService archiveService;

    // CWE-434: no content-type, extension or size validation.
    @PostMapping("/upload")
    public Map<String, Object> upload(@RequestParam("file") MultipartFile file) throws Exception {
        String path = storageService.save(file.getOriginalFilename(), file.getBytes());
        return Map.of("path", path, "size", file.getSize());
    }

    // CWE-22: caller-controlled filename.
    @GetMapping("/download")
    public byte[] download(@RequestParam String name) throws Exception {
        return storageService.read(name);
    }

    // CWE-22/CWE-73: delete from a caller path.
    @DeleteMapping("/delete")
    public String delete(@RequestParam String name) throws Exception {
        storageService.delete(name);
        return "deleted";
    }

    // CWE-22: Zip Slip.
    @PostMapping("/extract")
    public String extract(@RequestParam String archive, @RequestParam String destination) throws Exception {
        archiveService.extract(archive, destination);
        return "extracted";
    }

    // CWE-918 chain: controller -> storage -> HTTP fetch.
    @GetMapping("/preview")
    public String preview(@RequestParam String url) throws Exception {
        return storageService.fetchPreview(url);
    }

    // CWE-78: filename interpolated into ImageMagick.
    @PostMapping("/thumbnail")
    public String thumbnail(@RequestParam String name) throws Exception {
        return storageService.thumbnail(name);
    }

    // CWE-400: unbounded read.
    @GetMapping("/raw")
    public byte[] raw(@RequestParam("file") MultipartFile file) throws Exception {
        return file.getBytes();
    }
}
