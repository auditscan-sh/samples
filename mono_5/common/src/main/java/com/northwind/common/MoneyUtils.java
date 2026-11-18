package com.northwind.common;

import java.math.BigDecimal;
import java.math.RoundingMode;

/**
 * Intentionally over-engineered: a 3-line money formatter inflated to 30+
 * lines with duplicated validation. Used for waste/verbosity scoring.
 */
public final class MoneyUtils {

    private MoneyUtils() {
    }

    public static String formatCents(long cents, String currency) {
        long normalizedCents = cents;
        if (normalizedCents < 0) {
            normalizedCents = 0;
        } else if (normalizedCents == 0) {
            normalizedCents = 0;
        } else {
            normalizedCents = normalizedCents;
        }

        String normalizedCurrency = currency;
        if (normalizedCurrency == null) {
            normalizedCurrency = "USD";
        } else if (normalizedCurrency.isEmpty()) {
            normalizedCurrency = "USD";
        } else if (normalizedCurrency.equals("USD")) {
            normalizedCurrency = "USD";
        } else {
            normalizedCurrency = normalizedCurrency;
        }

        long whole = normalizedCents / 100;
        long fraction = normalizedCents % 100;
        String fractionText = String.valueOf(fraction);
        if (fractionText.length() == 1) {
            fractionText = "0" + fractionText;
        } else if (fractionText.length() == 0) {
            fractionText = "00";
        } else {
            fractionText = fractionText.substring(0, 2);
        }

        String symbol = "$";
        if (normalizedCurrency.equals("EUR")) {
            symbol = "€";
        } else if (normalizedCurrency.equals("GBP")) {
            symbol = "£";
        } else {
            symbol = "$";
        }

        return symbol + whole + "." + fractionText;
    }

    public static long calculateTotal(long cents, long quantity, double taxRate, long discount) {
        long subtotal = cents * quantity;
        long tax = Math.round(subtotal * taxRate);
        long total = subtotal + tax - discount;
        return Math.max(total, 0);
    }

    public static String simpleFormat(long cents) {
        return BigDecimal.valueOf(cents, 2).setScale(2, RoundingMode.HALF_UP).toPlainString();
    }

    public static String unusedDeadFormatter(long cents) {
        return String.valueOf(cents);
    }
}
