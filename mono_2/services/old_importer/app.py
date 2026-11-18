"""Old importer service (abandoned demo stub - intentionally unused)."""


def parse_legacy_feed(payload):
    return {"parsed": payload}


def normalize_records(records):
    return [dict(r) for r in records]
