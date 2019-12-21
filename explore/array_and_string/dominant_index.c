int dominantIndex(const int* nums, int numsSize){
    int max;
    int index;
    max = 0;
    index = -1;
    for(int i = 0; i < numsSize; i++) {
        if (nums[i] > max) {
            if (nums[i] >= 2 * max) {
                index = i;
            } else {
                index = -1;
            }
            max = nums[i];
        } else if (max < nums[i] * 2) {
            index = -1;
        }
    }
    return index;
}