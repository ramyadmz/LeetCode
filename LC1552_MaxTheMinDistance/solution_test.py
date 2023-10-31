import unittest

from solution import Solution

class TestSolution(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()
    def test_maxDistance(self):
        self.assertEqual(self.solution.maxDistance([1, 2, 3, 4, 7], 3), 2)
    

if __name__=="__main__":
    unittest.main()