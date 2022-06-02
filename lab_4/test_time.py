import unittest
from time_regexp import get_time


class TestGetTime(unittest.TestCase):
    def test_1(self):
        res = get_time('')
        self.assertEqual(res, None)

    def test_2(self):
        res = get_time('somedatahere')
        self.assertEqual(res, None)

    def test_3(self):
        res = get_time('Time today is 10:30 PM')
        self.assertEqual(res.group(), '10:30 PM')

    def test_4(self):
        res = get_time('Time today is 10:1000 PM')
        self.assertEqual(res, None)

    def test_5(self):
        res = get_time('Time today is 00:00 PM')
        self.assertEqual(res, None)

    def test_6(self):
        res = get_time('Time today is 00:00 VM')
        self.assertEqual(res, None)

    def test_7(self):
        res = get_time('Time today is 00:00')
        self.assertEqual(res, None)

    def test_8(self):
        res = get_time('Time today is 10:30 AM')
        self.assertEqual(res.group(), '10:30 AM')

