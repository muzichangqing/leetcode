// 2620.计数器
/**
 * 
 * @param {number} n 
 * @returns {Function} counter
 */
var createCounter = function(n) {
    return function() {
        return n++;
    };
};
// const counter = createCounter(10);
// console.log(counter());
// console.log(counter());
// console.log(counter());

// 2667.创建 Hello World 函数
/**
 * 
 * @returns {Function}
 */
var createHelloWorld = function() {
    return function(...args) {
        return "Hello World";
    };
};
const f = createHelloWorld();
console.log(f());