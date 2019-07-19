#include <stdio.h>
#include <stdlib.h>

double findMedian(int *nums, int numsSize) {
    double median;
    if (numsSize % 2 == 0) {
        median = (nums[numsSize / 2 - 1] + nums[numsSize / 2]) / 2.0;
    } else {
        median = nums[numsSize / 2];
    }
    return median;
}

double findMedianSortedArrays(int* nums1, int nums1Size, int* nums2, int nums2Size){
    if (nums1Size == 0) {
        return findMedian(nums2, nums2Size);
    }
    if (nums2Size == 0) {
        return findMedian(nums1, nums1Size);
    }

    int i, j, m, n;
    int *A, *B;
    if (nums1Size > nums2Size) {
        A = nums2;
        B = nums1;
        m = nums2Size;
        n = nums1Size;
    } else {
        A = nums1;
        B = nums2;
        m = nums1Size;
        n = nums2Size;
    }

    for (i=0; i<=m; i++) {
        j = (m + n) / 2 - i;
        int condition1 = 0;
        int condition2 = 0;

        // TODO: 这里有溢出
        // A[i-1] < B[j]
        if (j > n - 1 || i == 0) {
            condition1 = 1;
        } else if (A[i-1] <= B[j]) {
            condition1 = 1;
        }
        // B[j-1] < A[i]
        if (i > m - 1 || j == 0) {
            condition2 = 1;
        } else if ( A[i] >= B[j-1]) {
            condition2 = 1;
        }

        if (condition1 && condition2) {
            // is it!
            if (i + j > n - i + m - j) {
                if (i == 0) {
                    return B[j-1];
                }
                if (j == 0) {
                    return A[i-1];
                }
                return A[i-1] > B[j-1] ? A[i-1] : B[j - 1];
            }
            if (i + j < n - i + m - j) {
                if (i == m) {
                    return B[j];
                }
                if (j == n) {
                    return A[i];
                }
                return A[i] < B[j] ? A[i] : B[j];
            }
            int left;
            int right;
            if (i == 0) {
                left = B[j-1];
            } else if (j == 0) {
                left = A[i-1];
            } else {
                left = A[i-1] > B[j-1] ? A[i-1] : B[j-1];
            }
            if (i==m) {
                right = B[j];
            } else if (j==n) {
                right = A[i];
            } else {
                right = A[i] < B[j] ? A[i] : B[j];
            }
            return (left + right) / 2.0;
        }

    }

    return 0;
}



int main() {
    int num1Size, num2Size;
    num1Size = 1;
    num2Size = 1;
    int *nums1 = (int *)malloc(num1Size * sizeof(int));
    int *nums2 = (int *)malloc(num2Size * sizeof(int));

    nums1[0] = 1;

    nums2[0] = 1;

    double median = findMedianSortedArrays(nums1, num1Size, nums2, num2Size);

    printf("Result: %f\n", median);

    free(nums1);
    free(nums2);

    return 0;
}