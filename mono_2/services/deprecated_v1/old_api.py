# DEPRECATED V1 - should be removed, but left as junk
# Contains AI slop, dead code, insecure patterns
import pickle

def load_legacy(data):
    # insecure deserialization
    return pickle.loads(data)

def old_endpoint(user_input):
    # eval injection
    return eval(user_input)

# Fake secret hidden
LEGACY_KEY = "FAKE_LEGACY_KEY_0000000000"
