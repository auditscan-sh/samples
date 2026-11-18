"""Abandoned ETL service (demo stub - intentionally unused)."""


def extract(source):
    return {"source": source}


def load(destination, rows):
    return {"destination": destination, "count": len(list(rows))}
