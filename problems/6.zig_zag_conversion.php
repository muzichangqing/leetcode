<?php

class Solution {

    /**
     * @param String $s
     * @param Integer $numRows
     * @return String
     */
    function convert($s, $numRows) {
        if ($numRows < 2) {
            return $s;
        }
        $arr = str_split($s);
        $rows = [];
        $currentLine = 0;
        $dire = 1;
        foreach($arr as  $c) {
            $rows[$currentLine][] = $c;
            if ($currentLine == $numRows - 1) {
                $dire = -1;
            } elseif ($currentLine == 0 && $dire == -1) {
                $dire = 1;
            }
            $currentLine = $dire > 0 ? $currentLine + 1 : $currentLine - 1;
        }
        $rs = "";
        foreach($rows as $row) {
            $rs .= join($row);
        }
        return $rs;
    }
}