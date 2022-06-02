import unittest
from ip_regexp import get_ip


class TestGetIP(unittest.TestCase):
    def test_1(self):
        res = get_ip('')
        self.assertEqual(res, None)

    def test_2(self):
        res = get_ip('0.0.0.0')
        self.assertEqual(res.group(), '0.0.0.0')

    def test_3(self):
        res = get_ip('0.00.000.0000')
        self.assertEqual(res, None)

    def test_4(self):
        res = get_ip('127.255.93.34')
        self.assertEqual(res.group(), '127.255.93.34')

    def test_5(self):
        res = get_ip('17.25.093.034')
        self.assertEqual(res.group(), '17.25.093.034')

    def test_6(self):
        res = get_ip('1.1.1.1')
        self.assertEqual(res.group(), '1.1.1.1')

    def test_7(self):
        res = get_ip('3.0.4.1')
        self.assertEqual(res.group(), '3.0.4.1')

    def test_8(self):
        res = get_ip('255.255.255.255')
        self.assertEqual(res.group(), '255.255.255.255')

    def test_9(self):
        res = get_ip('2.100.76.199')
        self.assertEqual(res.group(), '2.100.76.199')

    def test_10(self):
        res = get_ip(' 27.27.27.27')
        self.assertEqual(res.group(), '27.27.27.27')

    def test_11(self):
        res = get_ip('24.1.1.1 ')
        self.assertEqual(res.group(), '24.1.1.1')

    def test_12(self):
        res = get_ip('24.1.1.1 word')
        self.assertEqual(res.group(), '24.1.1.1')

    def test_13(self):
        res = get_ip('word 24.1.1.1')
        self.assertEqual(res.group(), '24.1.1.1')

    def test_14(self):
        res = get_ip('250.1.1.1v')
        self.assertEqual(res, None)

    def test_15(self):
        res = get_ip('1.256.1.1')
        self.assertEqual(res, None)

    def test_16(self):
        res = get_ip('1.1.1.256')
        self.assertEqual(res, None)

    def test_17(self):
        res = get_ip('1.1.1.')
        self.assertEqual(res, None)

    def test_18(self):
        res = get_ip('1.1.256')
        self.assertEqual(res, None)

    def test_19(self):
        res = get_ip('1.1.-1.1')
        self.assertEqual(res, None)

    def test_20(self):
        res = get_ip('1.1.1.a')
        self.assertEqual(res, None)

    def test_21(self):
        res = get_ip('1.1.a.1')
        self.assertEqual(res, None)

    def test_22(self):
        res = get_ip('1.1a.1.1')
        self.assertEqual(res, None)


if __name__ == '__main__':
    unittest.main()