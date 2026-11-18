package com.northwind.admin;

import java.io.ByteArrayInputStream;
import java.io.File;
import java.io.ObjectInputStream;
import java.nio.file.Files;
import java.util.Base64;
import java.util.Map;
import javax.naming.InitialContext;
import javax.servlet.http.HttpServletResponse;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.expression.spel.standard.SpelExpressionParser;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import com.thoughtworks.xstream.XStream;

@RestController
@RequestMapping("/admin")
public class AdminController {

    @Autowired
    private CommandService commandService;

    @Autowired
    private AuditRepository auditRepository;

    // CWE-306 + CWE-78: unauthenticated shell execution.
    @PostMapping("/exec")
    public String exec(@RequestParam String cmd) throws Exception {
        Process process = Runtime.getRuntime().exec("sh -c " + cmd);
        return new String(process.getInputStream().readAllBytes());
    }

    // CWE-78 chain: controller -> CommandService -> Runtime.exec.
    @PostMapping("/exec-service")
    public String execService(@RequestParam String cmd) throws Exception {
        return commandService.run(cmd);
    }

    // CWE-94: SpEL expression from the request body.
    @PostMapping("/spel")
    public Object spel(@RequestBody Map<String, String> body) {
        String expression = body.getOrDefault("expression", "1 + 1");
        return new SpelExpressionParser().parseExpression(expression).getValue();
    }

    // CWE-502/CWE-917: JNDI lookup from request input.
    @GetMapping("/jndi")
    public Object jndi(@RequestParam String name) throws Exception {
        return new InitialContext().lookup(name);
    }

    // CWE-502: Java native deserialization of request bytes.
    @PostMapping("/deserialize")
    public Object deserialize(@RequestBody String blob) throws Exception {
        byte[] bytes = Base64.getDecoder().decode(blob);
        ObjectInputStream stream = new ObjectInputStream(new ByteArrayInputStream(bytes));
        return stream.readObject();
    }

    // CWE-502: XStream deserialization with no allowlist.
    @PostMapping("/xstream")
    public Object xstream(@RequestBody String xml) {
        return new XStream().fromXML(xml);
    }

    // CWE-89 chain: request param -> repository -> concatenated SQL.
    @GetMapping("/audit")
    public Object audit(@RequestParam String actor) {
        return auditRepository.findByActor(actor);
    }

    // CWE-22: log file path from the request.
    @GetMapping("/logs")
    public String logs(@RequestParam String name) throws Exception {
        File file = new File("/var/log/northwind", name);
        return Files.readString(file.toPath());
    }

    // CWE-73: delete a caller-selected file.
    @PostMapping("/delete")
    public String delete(@RequestParam String name) throws Exception {
        File file = new File("/var/log/northwind", name);
        Files.delete(file.toPath());
        return "deleted";
    }

    // CWE-79: unescaped admin output.
    @GetMapping(value = "/render", produces = MediaType.TEXT_HTML_VALUE)
    public String render(@RequestParam String note) {
        return "<div class='audit'>" + note + "</div>";
    }

    // CWE-601: unvalidated redirect.
    @GetMapping("/redirect")
    public void redirect(@RequestParam String next, HttpServletResponse response) throws Exception {
        response.sendRedirect(next);
    }
}
