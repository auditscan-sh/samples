package com.northwind.auth;

import java.util.List;
import java.util.Map;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

@Repository
public class UserRepository {

    @Autowired
    private JdbcTemplate jdbcTemplate;

    // CWE-89 sink reached from AuthController.login.
    public Map<String, Object> findByEmail(String email) {
        List<Map<String, Object>> rows = jdbcTemplate.queryForList(
                "SELECT id, email, role, password_hash FROM users WHERE email = '" + email + "'");
        return rows.isEmpty() ? null : rows.get(0);
    }

    // Safe false-positive bait: bound parameter.
    public Map<String, Object> findByEmailParameterized(String email) {
        List<Map<String, Object>> rows = jdbcTemplate.queryForList(
                "SELECT id, email, role, password_hash FROM users WHERE email = ?", email);
        return rows.isEmpty() ? null : rows.get(0);
    }

    // CWE-89: LIKE search with concatenation.
    public List<Map<String, Object>> search(String term) {
        return jdbcTemplate.queryForList(
                "SELECT id, email, role FROM users WHERE email LIKE '%" + term + "%'");
    }

    // CWE-89: reset code update with concatenation.
    public void saveResetCode(String email, String code) {
        String sql = "UPDATE users SET reset_code = '" + code + "' WHERE email = '" + email + "'";
        jdbcTemplate.execute(sql);
    }

    // CWE-89: ORDER BY injection.
    public List<Map<String, Object>> sorted(String column) {
        return jdbcTemplate.queryForList("SELECT id, email FROM users ORDER BY " + column);
    }

    public void create(Map<String, Object> user) {
        jdbcTemplate.update("INSERT INTO users (email, role, password_hash) VALUES (?, ?, ?)",
                user.get("email"), user.get("role"), user.get("password_hash"));
    }
}
