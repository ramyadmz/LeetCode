from typing import List
class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        def possible(speed):
            return sum((p+speed-1)//speed for p in piles) <= h
        low,high = 1,10 ** 9
        while low <= high:
            mid= low+(high-low)//2
            if possible(mid):
                high=mid-1
            else:
                low=mid+1
        return low