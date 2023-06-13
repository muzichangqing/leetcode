// 2629. 复合函数
/**
 * @param {Function[]} functions
 * @return {Function}
 */
var compose = function(functions) {
	return function(x) {
        for (let i = functions.length-1; i >= 0; i--) {
            x = functions[i](x);
        }
        return x;
    }
};

// 2666. 只允许一次函数调用
/**
 * @param {Function} fn
 * @return {Function}
 */
var once = function(fn) {
    let called = false;
    return function(...args){
        if (called) {
            return undefined;
        }
        called = true;
        return fn(...args);
    }
};