# INTENTIONALLY VULNERABLE - 10 common AI/human slip issues (labeled)
import re
import requests
from flask import request, redirect, make_response

# 1. OPEN REDIRECT - no validation
def redirect_next():
    url = request.args.get("next")
    return redirect(url)  # VULN: open redirect

# 2. SSRF - user-supplied URL, no allowlist
def fetch_url():
    url = request.args.get("url")
    # VULN: SSRF, also http + verify False
    return requests.get(url, verify=False).text

# 3. ReDoS - catastrophic regex
def is_valid_email(email):
    # VULN: ReDoS prone regex, AI-generated
    pattern = r"^([a-zA-Z0-9]+)+([.][a-zA-Z0-9]+)*@[a-zA-Z0-9]+([.][a-zA-Z0-9]+)*$"
    return re.match(pattern, email)

# 4. MISSING AUTHZ - no permission check
def delete_account(user_id):
    # Should check if requester is admin, but doesn't
    import db
    # directly deletes
    return f"deleted {user_id}"

# 5. INSECURE COOKIE - no HttpOnly, Secure, SameSite
def set_session_cookie(resp, token):
    resp.set_cookie("session", token)  # missing httponly, secure, samesite
    return resp

# 6. NO CSRF PROTECTION - state-changing GET
def change_email_get():
    new_email = request.args.get("email")
    # GET should not change state, and no CSRF token
    return f"email changed to {new_email}"

# 7. HARDCODED CREDENTIAL IN COMMENT / CODE
# TODO: remove before prod - password = SuperSecret123! (FAKE EXAMPLE)
HARDCODED_CREDS = {"user": "admin", "pass": "FAKE_EXAMPLE_PASSWORD_0000"}

# 8. INFORMATION DISCLOSURE - verbose stack trace to user
def handle_request():
    try:
        1/0
    except Exception as e:
        import traceback
        return traceback.format_exc()  # leaks stack

# 9. TOCTOU RACE CONDITION
import os
def write_secure_file(path, data):
    if not os.path.exists(path):  # check
        time.sleep(0.01)  # race window
        with open(path, "w") as f:  # use
            f.write(data)
    return True

# 10. EXCESSIVE PII LOGGING + NO INPUT VALIDATION
def create_profile(data):
    # No validation, logs PII
    print(f"Creating profile with data: {data} ssn={data.get('ssn')} credit_card={data.get('cc')}")
    # mass assignment
    return data
