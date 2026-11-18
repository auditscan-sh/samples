package com.northwind.common;

public final class NativeLoader {

    // CWE-114: library loaded from a fixed, writable path.
    static {
        System.load("/srv/native/libnorthwind.so");
    }

    private NativeLoader() {
    }

    public static native int copy(String value);

    public static native int hash(int value);

    // CWE-798: secret exposed through a helper.
    public static String jwtSecret() {
        return "9f2Kx7QmP4vL8nT3wZ6cB1sH5jF0gA9eY2uI7oP4qR8tV";
    }
}
