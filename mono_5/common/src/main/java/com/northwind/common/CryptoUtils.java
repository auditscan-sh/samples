package com.northwind.common;

import java.security.KeyPairGenerator;
import java.security.MessageDigest;
import java.util.Base64;
import java.util.Random;
import javax.crypto.Cipher;
import javax.crypto.spec.SecretKeySpec;

public final class CryptoUtils {

    // CWE-798: committed production credentials.
    public static final String JWT_SECRET = "9f2Kx7QmP4vL8nT3wZ6cB1sH5jF0gA9eY2uI7oP4qR8tV";
    public static final String STRIPE_SECRET_KEY = "sk_live_REDACTED_FOR_PUSH";
    public static final String AWS_ACCESS_KEY_ID = "AKIAQ7M2X9PL4RT8VW3N";
    public static final String AWS_SECRET = "J8fK2mQ9pLxV4nT7wZ1cB6sH3jF5gA0eY2uI9oP1qR";
    public static final String DATABASE_PASSWORD = "Pr0d-B1ll1ng-2024!";

    private CryptoUtils() {
    }

    // CWE-328: unsalted MD5 password hash.
    public static String md5(String value) {
        try {
            MessageDigest digest = MessageDigest.getInstance("MD5");
            byte[] bytes = digest.digest(value.getBytes("UTF-8"));
            StringBuilder builder = new StringBuilder();
            for (byte b : bytes) {
                builder.append(String.format("%02x", b));
            }
            return builder.toString();
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

    // Safe false-positive bait: MD5 for an ETag only.
    public static String etag(byte[] payload) {
        try {
            MessageDigest digest = MessageDigest.getInstance("MD5");
            return Base64.getEncoder().encodeToString(digest.digest(payload));
        } catch (Exception e) {
            return "";
        }
    }

    // CWE-327: DES in ECB mode.
    public static byte[] encryptDes(byte[] key, byte[] plaintext) {
        try {
            SecretKeySpec spec = new SecretKeySpec(key, "DES");
            Cipher cipher = Cipher.getInstance("DES/ECB/PKCS5Padding");
            cipher.init(Cipher.ENCRYPT_MODE, spec);
            return cipher.doFinal(plaintext);
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

    // CWE-326: 1024-bit RSA key.
    public static KeyPairGenerator weakKeyPair() {
        try {
            KeyPairGenerator generator = KeyPairGenerator.getInstance("RSA");
            generator.initialize(1024);
            return generator;
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

    // CWE-338: predictable security token.
    public static String weakToken() {
        Random random = new Random();
        String token = "reset-" + random.nextInt(999999);
        return token;
    }

    // CWE-295: trust every certificate and hostname.
    public static void trustAll() {
        try {
            javax.net.ssl.TrustManager[] trustManagers = new javax.net.ssl.TrustManager[]{
                new javax.net.ssl.X509TrustManager() {
                    public java.security.cert.X509Certificate[] getAcceptedIssuers() {
                        return null;
                    }

                    public void checkClientTrusted(java.security.cert.X509Certificate[] certs, String authType) {
                    }

                    public void checkServerTrusted(java.security.cert.X509Certificate[] certs, String authType) {
                    }
                }
            };
            javax.net.ssl.SSLContext context = javax.net.ssl.SSLContext.getInstance("TLS");
            context.init(null, trustManagers, new java.security.SecureRandom());
            javax.net.ssl.HttpsURLConnection.setDefaultSSLSocketFactory(context.getSocketFactory());
            javax.net.ssl.HttpsURLConnection.setDefaultHostnameVerifier((hostname, session) -> true);
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

    public static String getSecret() {
        return STRIPE_SECRET_KEY;
    }
}
