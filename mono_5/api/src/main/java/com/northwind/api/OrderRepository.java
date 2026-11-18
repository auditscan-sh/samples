package com.northwind.api;

import java.util.List;
import java.util.Map;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.RowMapper;
import org.springframework.stereotype.Repository;

@Repository
public class OrderRepository {

    @Autowired
    private JdbcTemplate jdbcTemplate;

    // CWE-89 sink reached from OrderController -> OrderService.
    public List<Map<String, Object>> findByCustomer(String customer) {
        return jdbcTemplate.query(
                "SELECT id, customer, total, status FROM orders WHERE customer = '" + customer + "'",
                (RowMapper<Map<String, Object>>) (rs, rowNum) -> Map.of(
                        "id", rs.getString("id"),
                        "customer", rs.getString("customer"),
                        "total", rs.getBigDecimal("total"),
                        "status", rs.getString("status")
                ));
    }

    // Safe false-positive bait: bound parameter.
    public List<Map<String, Object>> findByIdParameterized(String id) {
        return jdbcTemplate.query(
                "SELECT id, customer, total, status FROM orders WHERE id = ?",
                (RowMapper<Map<String, Object>>) (rs, rowNum) -> Map.of(
                        "id", rs.getString("id"),
                        "customer", rs.getString("customer"),
                        "total", rs.getBigDecimal("total"),
                        "status", rs.getString("status")
                ),
                id
        );
    }

    // CWE-89: stored nickname concatenated later.
    public List<Map<String, Object>> findByNickname(String nickname) {
        return jdbcTemplate.queryForList(
                "SELECT id, total FROM orders WHERE nickname = '" + nickname + "'");
    }

    // CWE-89: ORDER BY injection.
    public List<Map<String, Object>> sorted(String column, String direction) {
        return jdbcTemplate.queryForList(
                "SELECT id, total FROM orders ORDER BY " + column + " " + direction);
    }

    // CWE-89: update with concatenation.
    public void updateStatus(String orderId, String status) {
        jdbcTemplate.execute("UPDATE orders SET status = '" + status + "' WHERE id = '" + orderId + "'");
    }
}
