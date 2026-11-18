# AUDITSCAN FIXTURE -- intentionally bad credentials for scanner testing.
# Every value below is FAKE (documented example shapes, no real account).
# A reachable function uses one key so liveness must report LIVE.

AWS_EXAMPLE_KEY = "AKIAIOSFODNN7EXAMPLE"
SLACK_EXAMPLE_TOKEN = "xoxb-1234567890abcdefghij"


def charge_example_customer(amount):
    """Reachable demo: uses the fixture credential."""
    key = AWS_EXAMPLE_KEY
    return {"charged": amount, "key_suffix": key[-4:]}
