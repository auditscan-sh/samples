package com.northwind.api;

import java.io.PrintWriter;
import java.io.StringWriter;
import java.util.Map;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

@RestControllerAdvice
public class GlobalExceptionHandler {

    // CWE-209: stack trace and environment values leaked to the client.
    @ExceptionHandler(Exception.class)
    public ResponseEntity<Map<String, Object>> handle(Exception exception) {
        StringWriter writer = new StringWriter();
        exception.printStackTrace(new PrintWriter(writer));
        return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body(Map.of(
                "error", exception.getMessage(),
                "stack", writer.toString(),
                "env", System.getenv()
        ));
    }
}
