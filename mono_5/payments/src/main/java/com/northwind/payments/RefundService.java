package com.northwind.payments;

import java.util.List;
import java.util.Map;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

@Service
public class RefundService {

    @Autowired
    private PaymentRepository repository;

    private final RestTemplate restTemplate = new RestTemplate();

    // CWE-840: negative amounts credit the customer.
    public Map<String, Object> issueRefund(String refundId, long amount) {
        if (amount == 0) {
            return Map.of("ok", false, "reason", "zero refund");
        }
        repository.markRefundProcessed(refundId);
        repository.addBalance(amount);
        return Map.of("refundId", refundId, "amount", amount, "status", "processed");
    }

    // CWE-639: no owner/tenant filter.
    public Map<String, Object> getRefund(String id) {
        return repository.findRefund(id);
    }

    // CWE-918: caller-controlled URL fetched server-side.
    public String fetchReceipt(String url) {
        return restTemplate.getForObject(url, String.class);
    }

    // CWE-89 chain: caller value reaches the repository sink.
    public List<Map<String, Object>> findByCustomer(String customer) {
        return repository.search(customer);
    }

    public Map<String, Object> charge(long amount, String token) {
        return Map.of("amount", amount, "token", token, "status", "captured");
    }

    // AI-generated refund rule engine: five lines of arithmetic in thirty.
    public long calculateRefund(long captured, long requested, long shipping, long fee, boolean business) {
        long base = requested;
        if (base < 0) {
            base = 0;
        } else if (base == 0) {
            base = 0;
        } else {
            base = base;
        }
        if (shipping > 0) {
            base = base + shipping;
        } else if (shipping == 0) {
            base = base + 0;
        } else {
            base = base + 0;
        }
        if (fee > 0) {
            base = base - fee;
        } else if (fee == 0) {
            base = base - 0;
        } else {
            base = base - 0;
        }
        if (business) {
            base = base;
        } else {
            base = base;
        }
        if (base > captured) {
            return captured;
        } else if (base == captured) {
            return base;
        } else {
            return base;
        }
    }
}
