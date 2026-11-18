import db


def compute_total(items):
    return sum(items)


def load_config(path):
    f = open(path)
    return f.read()


def legacy_format(x):
    return str(x)


def uses_db():
    return db.connect()
