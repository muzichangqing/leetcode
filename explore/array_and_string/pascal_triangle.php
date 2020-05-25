<?php

class Solution {

    /**
     * @param Integer $numRows
     * @return Integer[][]
     */
    function generate($numRows) {
        $ans = [];
        if ($numRows < 1) {
            return $ans;
        }
        for ($i = 1; $i <= $numRows; $i++) {
            $row = [];
            $row[0] = 1;
            $lastRow = $ans[$i - 2] ?? [];
            for ($j = 1; $j < $i - 1; $j++) {
                $row[$j] = $lastRow[$j - 1] + $lastRow[$j];
            }
            $row[$i-1] = 1;
            $ans[] = $row;
        }
        return $ans;
    }
}

$solution = new Solution();
$numRows = 6;
print_r($solution->generate($numRows));