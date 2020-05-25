<?php

class Solution {

    /**
     * @param Integer[][] $matrix
     * @return Integer[]
     */
    function spiralOrder($matrix) {
        $ans = [];

        $rows = count($matrix);
        if ($rows < 1) {
            return $ans;
        }
        $cols = count($matrix[0]);
        if ($cols < 1) {
            return $ans;
        }

        $i = $j = 0;
        $min_i = $min_j = 0;
        $max_i = $rows - 1;
        $max_j = $cols - 1;

        $dire_i = 0;
        $dire_j = 1;

        $remains = $rows * $cols;
        while ($remains > 0) {
            $ans[] = $matrix[$i][$j];
            $remains -= 1;
            if ($j == $max_j && $dire_j > 0) {
                $dire_j = 0;
                $dire_i = 1;
                $min_i += 1;
            } elseif ($j == $min_j && $dire_j < 0) {
                $dire_j = 0;
                $dire_i = -1;
                $max_i -= 1;
            } elseif ($i == $max_i && $dire_i > 0) {
                $dire_i = 0;
                $dire_j = -1;
                $max_j -= 1;
            } elseif($i == $min_i && $dire_i < 0) {
                $dire_i = 0;
                $dire_j = 1;
                $min_j += 1;
            }
            $i += $dire_i;
            $j += $dire_j;
        }

        return $ans;
    }
}

$solution = new Solution();
$matrix = [
    [1, 2, 3, 4],
    [5, 6, 7, 8],
    [9,10,11,12]
];
$ans = $solution->spiralOrder($matrix);
echo implode(",", $ans) . "\n";

