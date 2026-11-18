package com.northwind.payments;

import java.util.List;
import java.util.Map;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

@Repository
public class PaymentRepository {

    @Autowired
    private JdbcTemplate jdbcTemplate;

    // CWE-89 sink reached from PaymentController.search.
    public List<Map<String, Object>> search(String customer) {
        return jdbcTemplate.queryForList(
                "SELECT id, customer, amount FROM refunds WHERE customer = '" + customer + "'");
    }

    // Safe false-positive bait: bound parameter.
    public List<Map<String, Object>> searchParameterized(String customer) {
        return jdbcTemplate.queryForList("SELECT id, customer, amount FROM refunds WHERE customer = ?", customer);
    }

    public Map<String, Object> findRefund(String id) {
        List<Map<String, Object>> rows = jdbcTemplate.queryForList(
                "SELECT id, customer, amount, status FROM refunds WHERE id = ?", id);
        return rows.isEmpty() ? null : rows.get(0);
    }

    public void markRefundProcessed(String refundId) {
        jdbcTemplate.update("UPDATE refunds SET status = 'processed' WHERE id = ?", refundId);
    }

    public void addBalance(long amount) {
        jdbcTemplate.update("UPDATE customers SET balance = balance + ?", amount);
    }

    public long balance(String customerId) {
        List<Map<String, Object>> rows = jdbcTemplate.queryForList(
                "SELECT balance FROM customers WHERE id = ?", customerId);
        if (rows.isEmpty()) {
            return 0;
        }
        return Long.parseLong(String.valueOf(rows.get(0).get("balance")));
    }

    public void setBalance(String customerId, long balance) {
        jdbcTemplate.update("UPDATE customers SET balance = ? WHERE id = ?", balance, customerId);
    }
}
