package com.northwind.payments;

import java.util.List;
import java.util.Map;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

@RestController
@RequestMapping("/payments")
public class PaymentController {

    @Autowired
    private RefundService refundService;

    @Autowired
    private BalanceService balanceService;

    private final RestTemplate restTemplate = new RestTemplate();

    // CWE-840: amount and currency are fully client-controlled.
    @PostMapping("/charge")
    public Map<String, Object> charge(@RequestParam long amount, @RequestParam String token) {
        return refundService.charge(amount, token);
    }

    // CWE-306 + CWE-840: unauthenticated refund with a trusted amount.
    @PostMapping("/refund")
    public Map<String, Object> refund(@RequestBody Map<String, Object> body) {
        String refundId = String.valueOf(body.get("refundId"));
        long amount = Long.parseLong(String.valueOf(body.get("amount")));
        return refundService.issueRefund(refundId, amount);
    }

    // CWE-639: any caller can read any refund.
    @GetMapping("/refund")
    public Map<String, Object> getRefund(@RequestParam String id) {
        return refundService.getRefund(id);
    }

    // CWE-362: concurrent withdrawal can double-spend.
    @PostMapping("/withdraw")
    public Map<String, Object> withdraw(@RequestParam String customerId, @RequestParam long amount) {
        return balanceService.withdraw(customerId, amount);
    }

    // CWE-345: signature header is accepted but never verified.
    @PostMapping("/webhook")
    public Map<String, Object> webhook(@RequestBody String body,
                                       @RequestHeader(value = "Stripe-Signature", required = false) String signature) {
        return Map.of("received", true, "signature", String.valueOf(signature), "body", body);
    }

    // CWE-918 chain: controller -> service -> RestTemplate.
    @GetMapping("/receipt")
    public String receipt(@RequestParam String url) {
        return refundService.fetchReceipt(url);
    }

    // CWE-89 chain: controller -> service -> repository.
    @GetMapping("/search")
    public List<Map<String, Object>> search(@RequestParam String customer) {
        return refundService.findByCustomer(customer);
    }

    // CWE-79: refund note rendered as HTML.
    @GetMapping(value = "/note", produces = MediaType.TEXT_HTML_VALUE)
    public String note(@RequestParam String note) {
        return "<div class='note'>" + note + "</div>";
    }

    // CWE-644: receipt link built from the Host header.
    @GetMapping("/receipt-link")
    public String receiptLink(@RequestHeader("Host") String host, @RequestParam String id) {
        return "https://" + host + "/payments/refund?id=" + id;
    }
}
