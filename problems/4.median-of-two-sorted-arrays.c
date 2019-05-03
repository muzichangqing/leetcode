

double findMedianSortedArrays(int* nums1, int nums1Size, int* nums2, int nums2Size){
    
}


/*
i + j = n - i + m - j
i + j = (n + m) / 2
j = (n + m) / 2 - i
if n > m maybe j < 0
so n < m

condition:
n <= m
0 <= i <= n
j = (n + m) / 2 - i

if (i > 0 && j < m) A[i-1] < B[j]
if (j > 0 && i < n) B[j-1] < A[i]
*/
