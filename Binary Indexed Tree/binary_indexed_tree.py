class NumArray(object):

    def __init__(self, nums):
        """
        :type nums: List[int]
        """
        self.nums = nums
        self.n = len(nums)
        self.BIT = [0] * (self.n + 1)

        self.constructBIT()

    # Time: nlog(n) | Space: O(n)
    def constructBIT(self):
        for i in range(self.n):
            BITIndex = i + 1
            while BITIndex <= self.n:
                self.BIT[BITIndex] += self.nums[i]
                BITIndex += (BITIndex & (-BITIndex))        

    # Time: log(n) | Space: O(1)
    def update(self, index, val):
        """
        :type index: int
        :type val: int
        :rtype: None
        """
        delta = val - self.nums[index]
        self.nums[index] = val
        BITIndex = index + 1
        while BITIndex <= self.n:
            self.BIT[BITIndex] += delta
            BITIndex += (BITIndex & (-BITIndex))

    # Time: log(n) | Space: O(1)
    def sumRange(self, left, right):
        """
        :type left: int
        :type right: int
        :rtype: int
        """
        return self.getRangeSum(right) - self.getRangeSum(left-1)
    
    def getRangeSum(self, index):
        _sum = 0
        BITIndex = index+1
        while BITIndex  > 0:
            _sum += self.BIT[BITIndex ]
            BITIndex  -= (BITIndex  & (-BITIndex ))
        return _sum

if __name__ == "__main__":
    # index 0   1  2  3   4   5   6  7  8  9  10  11  12  13 14 15
    nums = [2, -1, 8, 9, 22, -10, 1, 3, 6, 7, 11, -6, 12, -2, 1, 4]
    numArr = NumArray(nums)
    print("sumRange(5, 13)", numArr.sumRange(5, 13))
    print("Before update")
    print("nums", numArr.nums)
    print("BIT array", numArr.BIT)
    numArr.update(5, -16)
    print("After update")
    print("nums", numArr.nums)
    print("BIT array", numArr.BIT)

    """
    """