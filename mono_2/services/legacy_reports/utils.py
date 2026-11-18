# Hidden deep file with API key - not in root or env
# INTENTIONALLY VULNERABLE FIXTURE - FAKE SECRET
import requests
import os

# Fake API key - should be flagged as committed secret
GITHUB_TOKEN_LEGACY = "ghp_FAKE_EXAMPLE_TOKEN_0000000000000000000000"
INTERNAL_API_KEY = "FAKE-API-KEY-legacy-reports-0000-not-real"

def sync_external():
    # unverified non-https call + file delete
    data = requests.get("http://legacy-reports.example.com/sync", verify=False).text
    # deletes local file based on external input
    os.remove("/tmp/legacy_cache.json")

def fetch_insecure():
    # Another http call, no timeout, no rate limit
    return requests.post("http://example.com/api/collect", data={"key": INTERNAL_API_KEY}, verify=False)
