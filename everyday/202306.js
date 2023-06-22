// 2481. 分割圆的最少切割次数
/**
 * 
 * @param {number} n
 * @return {number} 
 */
var numberOfCuts = function(n) {
    if (n === 1) {
        return 0;
    }
    if (n % 2 === 0) {
        return n / 2;
    }
    return n;
};
/**
 * 1254. 统计封闭岛屿的数目
 * @param {number[][]} grid
 * @return {number} 
 */
var closedIsland = function(grid) {
    let height = grid.length;
    let width = grid[0].length;
    let dfs = function(i, j) {
        if (i === -1 || i === height || j === -1 || j === width) {
            return false;
        }
        if (grid[i][j] !== 0) {
            return true;
        }
        grid[i][j] = 2;
        // 避免短路，一个一个执行，而不是 bfs(i-1,j) && dfs(i+1, j) && ...
        let up = dfs(i-1, j);
        let down = dfs(i+1, j);
        let left = dfs(i, j-1);
        let right = dfs(i, j+1);
        return up && down && left && right;
    };
    let count = 0;
    for (let i = 0; i < height; i++) {
        for (let j = 0; j < width; j++) {
            if (grid[i][j] !== 0) {
                continue;
            }
            if (dfs(i, j)) {
                count++;
            }
        }
    }
    return count;
};

// 1595. 连通两组点的最小成本
var connectTwoGroups = function(cost) {
    const size1 = cost.length;
    const size2 = cost[0].length;
    const m = 1 << size2;
    const dp = Array.from(Array(size1 + 1), () => new Array(m).fill(Number.MAX_SAFE_INTEGER / 2));

    dp[0][0] = 0;

    for (let i = 1; i <= size1; i++) {
        for (let s = 0; s < m; s++) {
        for (let k = 0; k < size2; k++) {
            if ((s & (1 << k)) === 0) {
            continue;
            }
            dp[i][s] = Math.min(dp[i][s], dp[i][s ^ (1 << k)] + cost[i - 1][k]);
            dp[i][s] = Math.min(dp[i][s], dp[i - 1][s] + cost[i - 1][k]);
            dp[i][s] = Math.min(dp[i][s], dp[i - 1][s ^ (1 << k)] + cost[i - 1][k]);
        }
        }
    }

    return dp[size1][m - 1];
};

// 面试题 16.19. 水域大小
const pondSizes = (land) => {
    const m = land.length;
    const n = land[0].length;
    const resList = [];
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (land[i][j] === 0) {
                resList.push(dfs(land, i, j));
            }
        }
    }
    const res = resList.sort((a, b) => a - b);
    return res;
};

const dfs = (land, x, y) => {
    const m = land.length;
    const n = land[0].length;
    if (x < 0 || x >= m || y < 0 || y >= n || land[x][y] !== 0) {
        return 0;
    }
    land[x][y] = -1;
    let res = 1;
    for (let dx = -1; dx <= 1; dx++) {
        for (let dy = -1; dy <= 1; dy++) {
            if (dx === 0 && dy === 0) {
                continue;
            }
            res += dfs(land, x + dx, y + dy);
        }
    }
    return res;
};