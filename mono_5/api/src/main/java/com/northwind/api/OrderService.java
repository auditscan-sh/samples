package com.northwind.api;

import com.northwind.common.LegacyXml;
import java.io.ByteArrayInputStream;
import java.util.List;
import java.util.Map;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.data.mongodb.core.query.Criteria;
import org.springframework.data.mongodb.core.query.Query;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;
import org.w3c.dom.Document;

@Service
public class OrderService {

    @Autowired
    private OrderRepository orderRepository;

    @Autowired
    private MongoTemplate mongoTemplate;

    private final RestTemplate restTemplate = new RestTemplate();

    // CWE-89 chain: controller -> service -> repository.
    public List<Map<String, Object>> findByCustomer(String customer) {
        return orderRepository.findByCustomer(customer);
    }

    public List<Map<String, Object>> sorted(String column, String direction) {
        return orderRepository.sorted(column, direction);
    }

    // CWE-918: server-side fetch of a caller-controlled URL.
    public String fetchReceipt(String url) {
        return restTemplate.getForObject(url, String.class);
    }

    // CWE-943: raw Mongo criteria; `$where` comes from the caller.
    public List<Map> mongoSearch(String expression) {
        Query query = new Query(Criteria.where("$where").is(expression));
        return mongoTemplate.find(query, Map.class, "orders");
    }

    // CWE-611: delegates to the unsafe XML helper.
    public Document parseXml(String body) throws Exception {
        return LegacyXml.parseUnsafe(new ByteArrayInputStream(body.getBytes("UTF-8")));
    }

    // CWE-643 chain: controller -> service -> XPath string.
    public String userPath(String name) {
        return LegacyXml.userPath(name);
    }

    // CWE-840: total is computed from a client-supplied unit price.
    public long total(long clientPrice, long quantity) {
        return clientPrice * quantity;
    }

    public void updateStatus(String orderId, String status) {
        orderRepository.updateStatus(orderId, status);
    }
}
