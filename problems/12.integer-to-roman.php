<?php

class Solution {

    /**
     * @param Integer $num
     * @return String
     */
    function intToRoman($num) {
        $map = [
            1000 => 'M',
            900 => 'CM',
            500 => 'D',
            400 => 'CD',
            100 => 'C',
            90 => 'XC',
            50 => 'L',
            40 => 'XL',
            10 => 'X',
            9 => 'IX',
            5 => 'V',
            4 => 'IV',
            1 => 'I',
        ];

        $list = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];

        $index = 0;
        $ans = '';
        while($num != 0) {
            $cursor = $list[$index];
            if ($num < $cursor) {
                $index++;
                continue;
            }
            $ans .= $map[$cursor];
            $num -= $cursor;
        }
        return $ans;
    }
}

$solution = new Solution();
echo $solution->intToRoman(1994) . PHP_EOL;