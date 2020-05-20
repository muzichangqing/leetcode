

int maxArea(int* height, int heightSize){
    int i = 0, j = heightSize - 1, max = 0;
    while(i != j) {
        int area = (j - i) * (height[i] < height[j] ? height[i] : height[j]);
        if (area > max) max = area;
        if (height[i] < height[j]) i++;
        else j--;
    }
    return max;
}