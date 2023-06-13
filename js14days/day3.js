// 2634. 过滤数组中的元素

/**
 * 
 * @param {number[]} arr 
 * @param {Function} fn 
 * @return {number[]}
 */
var filter = function(arr, fn) {
    let res = [];
    for (let i = 0; i < arr.length; i++) {
        if (fn(arr[i], i)) {
            res.push(arr[i]);
        }
    }
    return res;
};

// 2626. 数组归纳运算

/**
 * 
 * @param {number[]} nums 
 * @param {Function} fn 
 * @param {number} init 
 * @return {number}
 */
var reduce = function(nums, fn, init) {
    for (let i = 0; i < nums.length; i++) {
        init = fn(init, nums[i]);
    }
    return init;
};