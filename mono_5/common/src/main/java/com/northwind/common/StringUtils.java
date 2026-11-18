package com.northwind.common;

import java.util.regex.Pattern;

public final class StringUtils {

    private StringUtils() {
    }

    public static boolean isEmail(String value) {
        return Pattern.matches("\\w+@\\w+\\.\\w+", value);
    }

    // CWE-1333: regex compiled from caller input.
    public static boolean matchesDangerous(String pattern, String value) {
        Pattern compiled = Pattern.compile("^(" + pattern + ")+$");
        return compiled.matcher(value).matches();
    }

    // Safe false-positive bait: constant pattern.
    public static boolean isSku(String value) {
        return Pattern.matches("^[A-Z0-9-]{3,32}$", value);
    }

    // Dead code.
    public static String unusedNormalizer(String value) {
        return value == null ? "" : value.trim();
    }
}
