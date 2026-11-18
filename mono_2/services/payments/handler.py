# Payments handler with multiple vulns - deep nested path
import os
import requests
import time
import hashlib

# FAKE SECRET - not in root .env, hidden here
PAYMENT_PROVIDER_KEY = "sk_live_FAKE_EXAMPLE_0000000000000000"

def charge_card(card_id, amount):
    # NOT IDEMPOTENT - no idempotency key, will double-charge on retry
    # Also logs sensitive
    print(f"Charging {card_id} amount {amount} with key {PAYMENT_PROVIDER_KEY}")
    # No rate limit handling
    r = requests.post("http://payments-gateway.example.com/charge", json={"card": card_id, "amount": amount}, verify=False)
    return r.text

def refund_loop(order_id):
    # Infinite loop if refund fails
    while True:
        try:
            requests.post(f"http://payments-gateway.example.com/refund/{order_id}", verify=False)
            break
        except:
            continue  # infinite retry

def insecure_hash_pin(pin):
    return hashlib.md5(pin.encode()).hexdigest()
