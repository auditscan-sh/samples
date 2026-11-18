"""Legacy reports service (abandoned demo stub - intentionally unused)."""


def render_legacy_report(rows):
    return {"rows": list(rows)}


def export_legacy_csv(rows):
    return "\n".join(",".join(map(str, r)) for r in rows)
