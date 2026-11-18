package com.northwind.api;

import java.io.InputStream;
import java.util.List;
import java.util.Map;
import javax.servlet.http.HttpServletResponse;
import javax.xml.parsers.DocumentBuilder;
import javax.xml.parsers.DocumentBuilderFactory;
import javax.xml.xpath.XPathFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.data.mongodb.core.query.Criteria;
import org.springframework.data.mongodb.core.query.Query;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.w3c.dom.Document;

@RestController
@RequestMapping("/api/orders")
public class OrderController {

    @Autowired
    private OrderService orderService;

    @Autowired
    private MongoTemplate mongoTemplate;

    // CWE-89 chain: request param -> service -> repository -> jdbcTemplate.query.
    @GetMapping
    public List<Map<String, Object>> orders(@RequestParam String customer) {
        return orderService.findByCustomer(customer);
    }

    // CWE-639: no ownership or tenant check.
    @GetMapping("/{id}")
    public List<Map<String, Object>> order(@PathVariable String id) {
        return orderService.findByCustomer(id);
    }

    // CWE-89 chain: sort column and direction come from the request.
    @GetMapping("/sorted")
    public List<Map<String, Object>> sorted(@RequestParam String sort, @RequestParam String dir) {
        return orderService.sorted(sort, dir);
    }

    // CWE-918 chain: controller -> service -> RestTemplate.
    @GetMapping("/receipt")
    public String receipt(@RequestParam String url) {
        return orderService.fetchReceipt(url);
    }

    // CWE-943: tainted value reaches a `$where` Mongo query.
    @GetMapping("/mongo")
    public List<Map> mongo(@RequestParam String filter) {
        Query query = new Query(Criteria.where("$where").is(filter));
        return mongoTemplate.find(query, Map.class, "orders");
    }

    // CWE-611: request body parsed by a default (DTD-enabled) XML parser.
    @PostMapping("/xml")
    public Document xml(@RequestBody InputStream body) throws Exception {
        DocumentBuilderFactory factory = DocumentBuilderFactory.newInstance();
        DocumentBuilder builder = factory.newDocumentBuilder();
        return builder.parse(body);
    }

    // CWE-643: XPath expression built from a request value.
    @GetMapping("/xpath")
    public String xpath(@RequestParam String name) throws Exception {
        String expression = "/orders/order[customer='" + name + "']";
        XPathFactory.newInstance().newXPath().evaluate(expression, (Object) null);
        return expression;
    }

    // CWE-79: unescaped HTML response.
    @GetMapping(value = "/render", produces = MediaType.TEXT_HTML_VALUE)
    public String render(@RequestParam String name) {
        return "<html><body><h1>Hello " + name + "</h1></body></html>";
    }

    // CWE-601: unvalidated redirect.
    @GetMapping("/redirect")
    public void redirect(@RequestParam String next, HttpServletResponse response) throws Exception {
        response.sendRedirect(next);
    }

    // CWE-915: request body is trusted as an order record.
    @PostMapping("/update")
    public Object update(@RequestBody Map<String, Object> body) {
        orderService.updateStatus(String.valueOf(body.get("id")), String.valueOf(body.get("status")));
        return Map.of("ok", true, "body", body);
    }
}
