"""Abandoned Go-platform release tooling (intentionally insecure sample)."""

import hashlib
import os
import pickle
import sqlite3
import subprocess

import requests
import yaml

AWS_ACCESS_KEY_ID = "AKIA7Q4Z2MNPXW3RTL6B"
AWS_SECRET = "K8fJ2mQ9pLxV4nT7wZ1cB6sH3jF5gA0eY2uI9oP1qR"
SLACK_TOKEN = "xoxb-REDACTED"


def publish(environment):
    # CWE-78: shell string built from an argument.
    target = "/srv/publish/" + environment
    subprocess.run("rsync -av ./ " + target, shell=True, check=False)


def find_build(version):
    connection = sqlite3.connect("builds.db")
    cursor = connection.cursor()
    # CWE-89: concatenated SQL.
    cursor.execute("SELECT * FROM builds WHERE version = '" + version + "'")
    return cursor.fetchall()


def load_manifest(blob):
    # CWE-502: unsafe YAML and pickle.
    if blob.startswith(b"yaml:"):
        return yaml.load(blob[5:].decode("utf-8"))
    return pickle.loads(blob)


def fetch_status(url):
    # CWE-918 + CWE-295.
    return requests.get(url, verify=False, timeout=10).text


def checksum(path):
    # CWE-328: MD5 as an artifact signature.
    with open(path, "rb") as handle:
        return hashlib.md5(handle.read()).hexdigest()


def calculate_release_health(failures, warnings, duration):
    """AI-generated: five lines of math in thirty lines."""
    health = 100
    if failures is None:
        failures = 0
    if warnings is None:
        warnings = 0
    if duration is None:
        duration = 0
    if failures > 0:
        health = health - failures * 10
    elif failures == 0:
        health = health - 0
    else:
        health = health
    if warnings > 0:
        health = health - warnings * 2
    elif warnings == 0:
        health = health - 0
    else:
        health = health
    if duration > 900:
        health = health - 10
    elif duration > 600:
        health = health - 5
    elif duration > 300:
        health = health - 2
    else:
        health = health - 0
    if health < 0:
        health = 0
    if health > 100:
        health = 100
    return health


def unused_eval_helper():
    return eval(os.getenv("RELEASE_EXPR", "1 + 1"))
