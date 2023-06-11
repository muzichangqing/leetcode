// 2267. 计数器II

/**
 * 
 * @param {integer} init
 * return { increment: Function, decrement: Function, reset Function} 
 */
var createCounter = function(init) {
    let num = init
    return {
        increment: () => ++num,
        decrement: () => --num,
        reset: () => { num = init; return num; } ,
    }
};

// const counter = createCounter(5);
// console.log(counter.increment());
// console.log(counter.reset());
// console.log(counter.decrement());

// 2635. 转换数组中的每个函数

/**
 * 
 * @param {number[]} arr 
 * @param {Function} fn
 * @return {number[]} 
 */
var map = function(arr, fn) {
    let returned_arr = [];
    for (let i = 0; i < arr.length; i++) {
        returned_arr.push(fn(arr[i], i));
    }
    return returned_arr;
};