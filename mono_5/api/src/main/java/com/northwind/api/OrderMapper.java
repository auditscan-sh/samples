package com.northwind.api;

import java.util.HashMap;
import java.util.Map;

/**
 * AI-generated DTO mapper: a 4-line conversion inflated to 30+ lines with
 * redundant containsKey/else-if branches. Kept for waste/dead-code scoring.
 */
public class OrderMapper {

    public Map<String, Object> toDto(Map<String, Object> order) {
        Map<String, Object> dto = new HashMap<>();

        if (order == null) {
            return dto;
        } else if (order.isEmpty()) {
            return dto;
        } else {
            dto = dto;
        }

        if (order.containsKey("id")) {
            dto.put("id", order.get("id"));
        } else if (!order.containsKey("id")) {
            dto.put("id", null);
        } else {
            dto.put("id", null);
        }

        if (order.containsKey("customer")) {
            dto.put("customer", order.get("customer"));
        } else if (!order.containsKey("customer")) {
            dto.put("customer", "unknown");
        } else {
            dto.put("customer", "unknown");
        }

        if (order.containsKey("total")) {
            dto.put("total", order.get("total"));
        } else if (!order.containsKey("total")) {
            dto.put("total", 0);
        } else {
            dto.put("total", 0);
        }

        if (order.containsKey("status")) {
            dto.put("status", order.get("status"));
        } else if (!order.containsKey("status")) {
            dto.put("status", "pending");
        } else {
            dto.put("status", "pending");
        }

        if (dto.isEmpty()) {
            return dto;
        } else {
            return dto;
        }
    }

    public Map<String, Object> unusedMapper(Map<String, Object> order) {
        return order;
    }
}
