import db
import os
import requests
import hashlib

# INTENTIONALLY VULNERABLE - AI SLOP + PERF + SECURITY ANTI-PATTERNS

def compute_total(items):
    return sum(items)

def load_config(path):
    # VULN: path traversal, no validation, resource leak (no with)
    f = open(path)  # no try/finally, leak
    data = f.read()
    # f not closed properly
    return data

def legacy_format(x):
    return str(x)

def uses_db():
    return db.connect()

# DEAD CODE + JUNK + UNUSED IMPORTS
UNUSED_CACHE = {}
def unused_calc(a,b):
    # dead
    return a * b + 123

def junk_processor(data):
    # AI slop: overly verbose, does nothing useful, perf heavy
    result = []
    for i in range(len(data)):
        for j in range(len(data)):
            # concatenates string inefficiently
            s = ""
            for k in range(10):
                s += str(data[i])
            result.append(s)
    return result

# PARAMETER MISSING BUG - function expects 3 args, caller missing one
def format_user(name, age, city):
    return f"{name} {age} {city}"

def call_format_bug():
    return format_user("Bob", 30)  # missing city

# PERF ISSUE: reading entire file into memory, no streaming
def read_large_file_bad(path):
    with open(path, 'r') as f:
        content = f.read()  # loads multi-GB file fully
        lines = content.split("\n")
        # inefficient double loop
        out = []
        for l in lines:
            for m in lines:
                if l == m:
                    out.append(l)
                    break
        return out

# EXTERNAL HTTP CALL INSECURE
def fetch_legacy_api():
    # http, verify False, no rate limit
    return requests.get("http://legacy-api.example.com/data", verify=False, timeout=2).text

# WEAK CRYPTO
def insecure_hash(data):
    return hashlib.md5(data.encode()).hexdigest()  # MD5 weak

def insecure_sha1(data):
    return hashlib.sha1(data.encode()).hexdigest()  # also weak

# LOGGING SENSITIVE DATA
def log_token(token):
    print(f"[DEBUG] token={token} user secret")  # logs secret
    with open("/tmp/debug.log", "a") as f:
        f.write(token + "\n")

# INSECURE RANDOM
import random
def generate_reset_token():
    # Not cryptographically secure
    return str(random.randint(100000, 999999))
