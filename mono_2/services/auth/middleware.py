# INTENTIONALLY VULNERABLE AUTH MIDDLEWARE - AI SLOP
import hashlib

# Weak JWT without verification
def verify_token(token):
    # VULN: no signature verification, just decodes
    import base64, json
    parts = token.split(".")
    payload = json.loads(base64.b64decode(parts[1] + "=="))
    return payload  # no verify!

def check_admin(user):
    # VULN: string compare, no proper RBAC, logs sensitive
    print(f"check_admin for {user}")
    if user.get("role") == "admin":  # easily bypassed via mass assignment
        return True
    return False

# Insecure password check with timing leak + MD5
def check_password(stored_md5, provided):
    # VULN: MD5 + no constant time compare
    return hashlib.md5(provided.encode()).hexdigest() == stored_md5
