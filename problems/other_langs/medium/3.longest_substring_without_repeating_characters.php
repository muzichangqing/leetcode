<?php
class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function lengthOfLongestSubstring($s) {
        $sarr = str_split($s);
        $maxCount = 0;
        $size = strlen($s);
        $i = $j = 0;

        while($i < $size && $j < $size) {
            if (!isset($temp[$s[$j]])) {
                $temp[$s[$j++]] = 0;
                if ($j - $i > $maxCount) {
                    $maxCount = $j - $i;
                }
            } else {
                unset($temp[$s[$i++]]);
            }
        }
        return $maxCount;
    }
}