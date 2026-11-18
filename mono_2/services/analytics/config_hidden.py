# HIDDEN CONFIG NOT IN ROOT OR .env - scanner should find this
# INTENTIONALLY VULNERABLE FIXTURE - FAKE KEYS ONLY
STRIPE_API_KEY = "sk_test_FAKE_EXAMPLE_0000000000000000"
AWS_SECRET_KEY = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
OPENAI_KEY = "sk-FAKE-EXAMPLE-1234567890-not-real-0000"

# Also insecure http endpoint
TELEMETRY_URL = "http://insecure-collector.example.com/v1/track"
