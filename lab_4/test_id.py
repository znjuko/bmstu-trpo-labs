import unittest
from quailified_id_regexp import get_qualified_id


class TestQualifiedID(unittest.TestCase):
    def test_1(self):
        res = get_qualified_id('')
        self.assertEqual(res, None)

    def test_2(self):
        res = get_qualified_id('attribute')
        self.assertEqual(res.group(), 'attribute')

    def test_3(self):
        res = get_qualified_id('attribute_1.attribute_2')
        self.assertEqual(res.group(), 'attribute_1.attribute_2')

    def test_4(self):
        res = get_qualified_id('attribute_1.attribute_2.attribute_3')
        self.assertEqual(res.group(), 'attribute_1.attribute_2.attribute_3')

    def test_5(self):
        res = get_qualified_id('attribute_1.attribute_2.attribute_3.x')
        self.assertEqual(res.group(), 'attribute_1.attribute_2.attribute_3.x')

    def test_6(self):
        res = get_qualified_id('x')
        self.assertEqual(res.group(), 'x')

    def test_7(self):
        res = get_qualified_id('_')
        self.assertEqual(res.group(), '_')

    def test_8(self):
        res = get_qualified_id('x._')
        self.assertEqual(res.group(), 'x._')

    def test_9(self):
        res = get_qualified_id('3x')
        self.assertEqual(res, None)

    def test_8(self):
        res = get_qualified_id('y.1')
        self.assertEqual(res, None)

    def test_9(self):
        res = get_qualified_id('.z')
        self.assertEqual(res, None)


if __name__ == '__main__':
    unittest.main()