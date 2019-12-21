<?

class Solution {

    /**
     * @param Integer[][] $matrix
     * @return Integer[][]
     */
    function updateMatrix($matrix) {
        $rowLen = count($matrix);
        if ($rowLen === 0) {
            return $matrix;
        }
        $colLen = count($matrix[0]);
        // 初始化
        $initValue = PHP_INT_MAX - $rowLen * $colLen;
        foreach($matrix as $i => $rows) {
            foreach($rows as $j => $value) {
                if ($value !== 0) {
                    $matrix[$i][$j] = $initValue;
                }
            }
        }
        // 左上到右下
        for ($i = 0; $i < $rowLen; $i++) {
            for ($j = 0; $j < $colLen; $j++) {
                $current = $matrix[$i][$j];
                // right
                if ($i < $rowLen - 1) {
                    if ($matrix[$i+1][$j] !== 0 && $matrix[$i+1][$j] > $current + 1) {
                        $matrix[$i+1][$j] = $current + 1;
                    }
                }
                // down
                if ($j < $colLen - 1) {
                    if ($matrix[$i][$j+1] !== 0 && $matrix[$i][$j+1] > $current + 1) {
                        $matrix[$i][$j+1] = $current + 1;
                    }
                }
            }
        }
        // 右下到左上
        for ($i = $rowLen - 1; $i >= 0; $i--) {
            for($j = $colLen - 1; $j >= 0; $j--) {
                $current = $matrix[$i][$j];
                // left
                if ($i > 0) {
                    if ($matrix[$i-1][$j] !== 0 && $matrix[$i-1][$j] > $current + 1) {
                        $matrix[$i-1][$j] = $current + 1;
                    }
                }
                // up
                if ($j > 0) {
                    if ($matrix[$i][$j-1] !== 0 && $matrix[$i][$j-1] > $current + 1) {
                        $matrix[$i][$j-1] = $current + 1;
                    }
                }
            }
        }
        return $matrix;
    }
}