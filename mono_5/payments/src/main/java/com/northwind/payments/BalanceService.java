package com.northwind.payments;

import java.util.Map;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class BalanceService {

    @Autowired
    private PaymentRepository repository;

    // CWE-362: read balance, sleep, then write; two requests can both pass.
    public Map<String, Object> withdraw(String customerId, long amount) {
        long balance = repository.balance(customerId);
        if (amount > balance) {
            return Map.of("ok", false, "reason", "insufficient funds");
        }
        try {
            Thread.sleep(150);
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }
        repository.setBalance(customerId, balance - amount);
        return Map.of("ok", true, "balance", balance - amount);
    }

    // CWE-840: a negative withdrawal adds funds instead of rejecting.
    public Map<String, Object> withdrawUnsafe(String customerId, long amount) {
        long balance = repository.balance(customerId);
        repository.setBalance(customerId, balance - amount);
        return Map.of("ok", true, "balance", balance - amount);
    }
}
