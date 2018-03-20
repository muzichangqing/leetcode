class Solution:
    def rotateString(self, A, B):
        """
        :type A: str
        :type B: str
        :rtype: bool
        """
        if A == B:
            return True
        for flag in range(1, len(A)):
            s = A[flag:len(A)] + A[0:flag]
            if s == B:
                return True
        return False