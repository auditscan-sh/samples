package com.northwind.payments;

import java.util.Map;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/stripe")
public class StripeWebhookController {

    // CWE-798: committed webhook secret and API key.
    private static final String WEBHOOK_SECRET = "whsec_7Kd8fJ2mQpLxV4nT7wZ1cB6sH";
    private static final String API_KEY = "sk_live_REDACTED_FOR_PUSH";

    // CWE-345: signature is read but never validated.
    @PostMapping("/webhook")
    public Map<String, Object> webhook(@RequestBody Map<String, Object> event,
                                       @RequestHeader(value = "Stripe-Signature", required = false) String signature) {
        System.out.println("stripe signature=" + signature + " key=" + API_KEY);
        return Map.of("received", true, "event", event);
    }

    // CWE-208: direct string compare.
    public boolean verify(String provided) {
        return WEBHOOK_SECRET == provided;
    }

    public String secret() {
        return WEBHOOK_SECRET;
    }
}
