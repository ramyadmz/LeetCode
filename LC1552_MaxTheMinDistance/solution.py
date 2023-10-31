from typing import List
class Solution:
    def maxDistance(self, position: List[int], m: int) -> int:
        position.sort()
        def possible(dist):
            prev=position[0]
            count=1
            for p in position[1:]:
                if p - prev >= dist:
                    prev=p
                    count+=1

            return count >= m
        
        low,high,ans=1,position[-1]-position[0],-1
        while low <= high:
            mid=low+(high-low)//2
            if possible(mid):
                ans = mid
                low = mid+1
            else:
                high = mid -1
        return ans