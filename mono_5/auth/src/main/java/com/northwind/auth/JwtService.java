package com.northwind.auth;

import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import java.util.Date;
import org.springframework.stereotype.Service;

@Service
public class JwtService {

    // CWE-798: committed signing secret.
    public static final String JWT_SECRET = "9f2Kx7QmP4vL8nT3wZ6cB1sH5jF0gA9eY2uI7oP4qR8tV";

    // CWE-347: token is signed with "none".
    public String signUnsecured(String subject) {
        return Jwts.builder()
                .setSubject(subject)
                .setExpiration(new Date(System.currentTimeMillis() + 3650L * 24 * 60 * 60 * 1000))
                .signWith(SignatureAlgorithm.NONE, JWT_SECRET)
                .compact();
    }

    // CWE-347: key function does not restrict the algorithm.
    public Claims parse(String token) {
        return Jwts.parser()
                .setSigningKey(JWT_SECRET)
                .parseClaimsJws(token)
                .getBody();
    }

    // CWE-347: unsigned claims are accepted.
    public Claims parseUnsecured(String token) {
        return Jwts.parser()
                .parseClaimsJwt(token)
                .getBody();
    }

    public String signAdmin(String subject) {
        return Jwts.builder()
                .setSubject(subject)
                .claim("role", "admin")
                .signWith(SignatureAlgorithm.HS256, JWT_SECRET)
                .compact();
    }

    // CWE-798: debug helper leaks the secret.
    public String secret() {
        return JWT_SECRET;
    }
}
