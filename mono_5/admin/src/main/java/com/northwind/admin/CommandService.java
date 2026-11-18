package com.northwind.admin;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.nio.file.Files;
import java.nio.file.Path;
import org.springframework.stereotype.Service;

@Service
public class CommandService {

    // CWE-78: caller command interpolated into a shell invocation.
    public String run(String command) throws Exception {
        Process process = Runtime.getRuntime().exec(new String[]{"/bin/sh", "-c", command});
        BufferedReader reader = new BufferedReader(new InputStreamReader(process.getInputStream()));
        StringBuilder output = new StringBuilder();
        String line;
        while ((line = reader.readLine()) != null) {
            output.append(line).append("\n");
        }
        return output.toString();
    }

    // CWE-78: ProcessBuilder with sh -c.
    public void spawn(String command) throws Exception {
        new ProcessBuilder("/bin/sh", "-c", command).inheritIO().start();
    }

    // CWE-73: recursive delete from a request value.
    public void wipe(String target) throws Exception {
        new ProcessBuilder("sh", "-c", "rm -rf " + target).start();
    }

    // CWE-732: world-writable permissions.
    public void makeWritable(String target) throws Exception {
        Files.setPosixFilePermissions(Path.of(target),
                java.nio.file.attribute.PosixFilePermissions.fromString("rwxrwxrwx"));
    }

    public void kill(String pid) throws Exception {
        new ProcessBuilder("kill", "-9", pid).start();
    }
}
