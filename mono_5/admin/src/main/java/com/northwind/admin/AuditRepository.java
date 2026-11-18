package com.northwind.admin;

import java.util.List;
import java.util.Map;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

@Repository
public class AuditRepository {

    @Autowired
    private JdbcTemplate jdbcTemplate;

    // CWE-89 sink reached from AdminController.audit.
    public List<Map<String, Object>> findByActor(String actor) {
        return jdbcTemplate.queryForList(
                "SELECT id, actor, action FROM audit_log WHERE actor = '" + actor + "'");
    }

    // CWE-89: ORDER BY injection.
    public List<Map<String, Object>> sorted(String column) {
        return jdbcTemplate.queryForList("SELECT id, actor, action FROM audit_log ORDER BY " + column);
    }
}
