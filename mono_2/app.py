import sys
import os
import json
import hashlib
import pickle
import requests
import time
from util import compute_total
from service import increment

# INTENTIONALLY VULNERABLE FIXTURE - FOR SCANNER TESTING ONLY
# This file contains AI-slop patterns, dead code, perf anti-patterns, etc.

# --- DEAD CODE / JUNK (unused, never called) ---
def unused_helper(x):
    return x * 2

def dead_function_a(a, b, c):  # param c never used, junk logic
    temp = []
    for i in range(100):
        temp.append(a * b)
    return temp[0]

def dead_function_b():
    x = "this is dead code"
    y = {"unused": 123}
    return None

UNUSED_GLOBAL = {"foo": "bar", "junk": [1,2,3]*100}
ANOTHER_UNUSED = lambda x: x

# AI SLOP: overly verbose, redundant comments, unnecessary wrapper
def process(items):
    """
    This function processes items by computing total.
    It takes items as input and returns total.
    It is very important to compute total correctly.
    """
    # Check if items is not None
    if items is not None:
        # Compute total using helper
        total = compute_total(items)
        # Return the total
        return total
    else:
        # If items is None, return 0
        return 0

# PARAMETER MISSING / MISMATCH - caller will miss required arg
def create_user(name, email, role="user"):
    return {"name": name, "email": email, "role": role}

def call_create_user_bug():
    # Missing required 'email' param - will raise TypeError but AI slop ignores
    return create_user("alice")  # BUG: missing email

# PERFORMANCE BIG ISSUE: O(n^2) nested loop, doing same work repeatedly, no pagination
def get_all_users_slow(users):
    result = []
    for u in users:
        for v in users:  # O(n^2) - should be O(n)
            if u["id"] == v["id"]:
                result.append(u)
    # Also loads everything into memory without limit
    return result

def compute_total_slow(items):
    # Recomputes sum in inefficient way, N+1 style
    total = 0
    for i in range(len(items)):
        for j in range(len(items)):
            total += items[i]  # adds n times too many
    return total

# FUNCTION THAT DELETES FILES - PATH TRAVERSAL VULN (scanner should flag)
def delete_user_file(user_input):
    # VULN: no sanitization, allows ../../etc/passwd style
    # FIXTURE: only operates under /tmp/vuln_demo, but pattern is intentionally bad
    base = "/tmp/vuln_demo"
    path = os.path.join(base, user_input)
    # Intentionally vulnerable: no check if path escapes base
    if os.path.exists(path):
        os.remove(path)  # DANGEROUS PATTERN - flagged by scanner
        return f"deleted {path}"
    return "not found"

def delete_logs_unsafe():
    # Mass delete without confirmation - bad pattern
    import shutil
    shutil.rmtree("/tmp/logs", ignore_errors=True)  # overly broad

# EXTERNAL CALL TO UNVERIFIED/NON-HTTPS URL - insecure
def fetch_external_data(user_id):
    # VULN: http not https, verify=False, no timeout handling properly
    url = f"http://example.com/api/user/{user_id}"  # should be https
    resp = requests.get(url, verify=False, timeout=1)  # verify=False is bad
    return resp.text

def send_telemetry(data):
    # Calls unverified external URL, leaks data
    requests.post("http://insecure-telemetry.example.com/collect", data=data, verify=False)

# INFINITE LOOP ON SOME ISSUE
def retry_forever(task):
    # No backoff, no max retries, infinite loop if task always fails
    while True:
        try:
            result = task()
            if result == "success":
                break
        except Exception:
            pass  # swallows exception and loops forever
        # No sleep limit, no break condition

def poll_infinite():
    while True:  # never exits
        time.sleep(0.1)
        # does nothing, burns CPU

# NO RATE LIMIT HANDLING - sensitive place
def call_external_api_no_rate_limit(items):
    # Will hammer API without rate limit, no 429 handling
    results = []
    for item in items:
        r = requests.get(f"http://example.com/api/data/{item}", verify=False)
        results.append(r.json())  # no check for 429, no backoff
    return results

# SENSITIVE PLACE THAT'S NOT IDEMPOTENT WHICH SHOULD BE BUT NOT
# Payment processing should be idempotent but isn't
payment_processed = set()
def process_payment(order_id, amount):
    # VULN: not idempotent - double charge if retried, no idempotency key
    # Also no transaction, race condition
    global payment_processed
    # Should check if already processed, but logic is flawed
    print(f"Processing payment {order_id} for {amount}")  # logs sensitive
    increment()  # side effect not idempotent
    # Simulate charging
    time.sleep(0.01)
    payment_processed.add(order_id)
    return {"status": "charged", "order_id": order_id, "amount": amount}

def checkout_non_idempotent(cart_id):
    # Called twice will charge twice
    return process_payment(cart_id, 100)

# MORE AI SLOP - verbose, redundant error handling
def run_console():
    try:
        cmd = input("> ")
        result = eval(cmd)  # VULN: code injection - scanner should flag
        print(f"Result is {result}")
        return result
    except Exception as e:
        # AI slop: catches all exceptions and just passes
        pass
    finally:
        # useless finally
        pass

# WEAK CRYPTO + LOGGING SENSITIVE DATA + VERBOSE ERROR
def hash_password_insecure(password):
    # VULN: MD5 for password - weak crypto
    return hashlib.md5(password.encode()).hexdigest()

def login(username, password):
    hashed = hash_password_insecure(password)
    print(f"DEBUG: login attempt {username}:{password} hash {hashed}")  # logs sensitive
    try:
        # fake check
        return hashed == "abc"
    except Exception as e:
        # verbose error leak
        return f"Error: {str(e)} {repr(e)} stack: {e.__traceback__}"

# INSECURE DESERIALIZATION
def load_session(cookie):
    # VULN: pickle.loads on user input
    return pickle.loads(cookie)

# XSS - unsanitized output
def render_profile(name):
    # VULN: reflected XSS
    return f"<div>Hello {name}</div>"  # no escaping

# CORS WILDCARD + MISSING AUTH
def api_get_user(request):
    # No auth check!
    headers = {"Access-Control-Allow-Origin": "*"}  # overly permissive
    user = {"id": 1, "name": "admin"}
    return json.dumps(user), headers

def main():
    data = [1, 2, 3]
    print(process(data))
    increment()
    # dead call
    call_create_user_bug()
    # perf issue demo
    get_all_users_slow([{"id": i} for i in range(100)])
    # non-idempotent demo
    process_payment("order123", 99.99)


if __name__ == "__main__":
    main()
