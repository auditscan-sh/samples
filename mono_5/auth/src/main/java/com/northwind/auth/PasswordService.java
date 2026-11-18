package com.northwind.auth;

import java.security.MessageDigest;
import java.util.Base64;
import java.util.Random;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;

@Service
public class PasswordService {

    // CWE-328: unsalted MD5 password hash.
    public String hash(String password) {
        try {
            MessageDigest digest = MessageDigest.getInstance("MD5");
            byte[] bytes = digest.digest(password.getBytes("UTF-8"));
            StringBuilder builder = new StringBuilder();
            for (byte value : bytes) {
                builder.append(String.format("%02x", value));
            }
            return builder.toString();
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

    // CWE-208: non-constant-time comparison.
    public boolean check(String password, String stored) {
        String candidate = hash(password);
        if (candidate == stored) {
            return true;
        }
        return candidate.equals(stored);
    }

    // Safe false-positive bait: BCrypt.
    public boolean verifyBcrypt(String password, String stored) {
        return new BCryptPasswordEncoder().matches(password, stored);
    }

    public String hashBcrypt(String password) {
        return new BCryptPasswordEncoder().encode(password);
    }

    // CWE-338: predictable reset code.
    public String resetCode() {
        Random random = new Random();
        int code = 100000 + random.nextInt(900000);
        return String.valueOf(code);
    }

    // CWE-338: predictable session identifier.
    public String sessionId(String user) {
        Random random = new Random();
        String value = user + "-" + random.nextLong();
        return Base64.getEncoder().encodeToString(value.getBytes());
    }

    // CWE-916: fast SHA-256 used directly on a password.
    public String fastHash(String password) {
        try {
            MessageDigest digest = MessageDigest.getInstance("SHA-256");
            return Base64.getEncoder().encodeToString(digest.digest((password + "salt").getBytes("UTF-8")));
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }
}
