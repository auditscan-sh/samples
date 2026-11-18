# ABANDONED ETL - dead code, junk, unused service example
# INTENTIONALLY VULNERABLE FIXTURE
import requests

# Never called, dead code
def etl_job():
    # fetches all without pagination - perf issue
    data = requests.get("http://example.com/etl/all", verify=False).json()
    # O(n^2) processing
    for a in data:
        for b in data:
            if a == b:
                print(a)

def unused_transform(x):
    return x * 2

# Infinite loop
def run_forever():
    while True:
        etl_job()
