package com.northwind.admin;

import java.util.Map;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class DebugController {

    // CWE-200: environment dump.
    @GetMapping("/debug/env")
    public Map<String, String> env() {
        return System.getenv();
    }

    // CWE-200: system properties and secrets.
    @GetMapping("/debug/properties")
    public Map<Object, Object> properties() {
        return System.getProperties();
    }

    // CWE-73: heap dump path from a request.
    @GetMapping("/debug/heap")
    public String heap(@RequestParam String path) throws Exception {
        com.sun.management.HotSpotDiagnosticMXBean bean =
                java.lang.management.ManagementFactory.getPlatformMXBean(
                        com.sun.management.HotSpotDiagnosticMXBean.class);
        bean.dumpHeap(path, true);
        return path;
    }

    // CWE-79: debug page renders caller HTML.
    @GetMapping(value = "/debug/render", produces = MediaType.TEXT_HTML_VALUE)
    public String render(@RequestParam String value) {
        return "<pre>" + value + "</pre>";
    }

    // CWE-798: secret helper.
    @GetMapping("/debug/secret")
    public String secret() {
        return "9f2Kx7QmP4vL8nT3wZ6cB1sH5jF0gA9eY2uI7oP4qR8tV";
    }
}
