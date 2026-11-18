import sys
from util import compute_total
from service import increment


def process(items):
    total = compute_total(items)
    return total


def unused_helper(x):
    return x * 2


def run_console():
    cmd = input("> ")
    result = eval(cmd)
    return result


def main():
    data = [1, 2, 3]
    print(process(data))
    increment()


if __name__ == "__main__":
    main()
