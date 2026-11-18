"""Deprecated v1 API service (demo stub - intentionally unused)."""


def handle_v1(request):
    return {"version": 1, "echo": request}


def migrate_v1_to_v2(request):
    return {"version": 2, "echo": request}
