// 2623. 记忆函数
/**
 * @param {Function} fn
 */
function memoize(fn) {
    let cache = new Map();
    return function(...args) {
        args_key = args.toString();
        if (!cache.has(args_key)) {
            cache.set(args_key, fn(...args))
        }
        return cache.get(args_key);
    }
}

let callCount = 0;
const memoizeFn = memoize(function(a, b) {
    callCount += 1;
    return a + b;
});
console.log(memoizeFn(2, 3));
console.log(memoizeFn(2, 3));
console.log(callCount);
