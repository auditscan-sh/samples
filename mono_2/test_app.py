import unittest

from util import compute_total
from service import increment
from app import process


class TestApp(unittest.TestCase):
    def test_compute(self):
        self.assertEqual(compute_total([1, 2, 3]), 6)

    def test_increment(self):
        self.assertEqual(increment(), 1)

    def test_process(self):
        self.assertEqual(process([1, 2, 3]), 6)


if __name__ == "__main__":
    unittest.main()
