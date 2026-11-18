COUNTER = 0
import time
import requests
import os

# INTENTIONALLY VULNERABLE SERVICE LAYER

def increment():
    global COUNTER
    COUNTER += 1
    return COUNTER

# NOT IDEMPOTENT - should be idempotent for retry safety
def create_order(order_id, items):
    # If called twice, creates duplicate order, no idempotency key check
    print(f"Creating order {order_id}")  # no check if exists
    with open(f"/tmp/order_{order_id}.json", "w") as f:
        import json
        json.dump({"items": items}, f)
    return True

# NO RATE LIMIT, INFINITE RETRY
def call_payment_gateway(payload):
    while True:  # infinite loop if gateway down
        try:
            r = requests.post("http://payments.example.com/pay", json=payload, verify=False)
            if r.status_code == 429:
                # BUG: ignores rate limit, just retries immediately without backoff
                continue
            return r.json()
        except Exception:
            time.sleep(0.1)
            continue  # infinite retry, no max attempts

# FILE DELETE + PATH TRAVERSAL
def cleanup_upload(filename):
    # VULN: path traversal
    path = os.path.join("/tmp/uploads", filename)
    os.remove(path)  # no validation

# MASS ASSIGNMENT VULN
def update_user(user, data):
    # VULN: blindly sets all fields from user input, allows is_admin overwrite
    for k, v in data.items():
        setattr(user, k, v)
    return user

# DEAD CODE
def dead_service_helper():
    x = [1,2,3]
    y = {k: v for k, v in enumerate(x)}
    return None

def another_dead(a,b,c,d):
    # params c,d unused, junk
    return a+b
