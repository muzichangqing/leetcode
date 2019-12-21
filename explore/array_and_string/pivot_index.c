int pivotIndex(int* nums, int numsSize){
    if (numsSize == 0) {
        return -1;
    }
    int sumR = 0;
    for (int i = 0; i < numsSize; i++) {
        sumR += nums[i];
    }
    int sumL = 0;
    for (int i = 0; i < numsSize; i++) {
        if (i > 0) {
            sumL += nums[i-1];
            sumR -= nums[i-1];
        }
        if (sumL == sumR - nums[i]) {
            return i;
        }
    }
    return -1;
}