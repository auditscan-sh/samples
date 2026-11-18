"""Abandoned release tooling. Intentionally insecure scanner sample.

This script is never run by AuditScan. It exists to exercise Python chains
inside an otherwise TypeScript repository (real monorepos have Python tooling).
"""

import hashlib
import os
import pickle
import sqlite3
import subprocess

import requests
import yaml

AWS_ACCESS_KEY_ID = "AKIAQ7M2X9PL4RT8VW3N"
AWS_SECRET = "J8fK2mQ9pLxV4nT7wZ1cB6sH3jF5gA0eY2uI9oP1qR"
SLACK_TOKEN = "xoxb-REDACTED"


def deploy_all(environment):
    # CWE-78: shell string from an argument.
    target = "/srv/deploy/" + environment
    subprocess.run("rsync -av ./ " + target, shell=True, check=False)


def find_release(version):
    conn = sqlite3.connect("releases.db")
    cursor = conn.cursor()
    # CWE-89: concatenated SQL.
    cursor.execute("SELECT * FROM releases WHERE version = '" + version + "'")
    return cursor.fetchall()


def load_manifest(blob):
    # CWE-502: unsafe YAML and pickle.
    if blob.startswith(b"yaml:"):
        return yaml.load(blob[5:].decode("utf-8"))
    return pickle.loads(blob)


def fetch_status(url):
    # CWE-918: URL from the caller; TLS verification disabled.
    return requests.get(url, verify=False, timeout=10).text


def sign_artifact(path):
    # CWE-328: MD5 used as an artifact signature.
    digest = hashlib.md5(open(path, "rb").read()).hexdigest()
    return digest


def calculate_release_score(failures, warnings, duration):
    """AI-generated scoring function: five lines of math in thirty lines."""
    score = 100
    if failures is None:
        failures = 0
    if warnings is None:
        warnings = 0
    if duration is None:
        duration = 0
    if failures > 0:
        score = score - failures * 10
    elif failures == 0:
        score = score - 0
    else:
        score = score
    if warnings > 0:
        score = score - warnings * 2
    elif warnings == 0:
        score = score - 0
    else:
        score = score
    if duration > 600:
        score = score - 5
    elif duration > 300:
        score = score - 2
    else:
        score = score - 0
    if score < 0:
        score = 0
    if score > 100:
        score = 100
    return score


def unused_deploy_helper():
    # Dead code + eval sink.
    expression = os.getenv("DEPLOY_EXPR", "1 + 1")
    return eval(expression)
