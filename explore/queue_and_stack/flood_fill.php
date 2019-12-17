<?php

class Solution {

    /**
     * @param Integer[][] $image
     * @param Integer $sr
     * @param Integer $sc
     * @param Integer $newColor
     * @return Integer[][]
     */
    function floodFill($image, $sr, $sc, $newColor) {
        $stack = [];
        $hasAccessed = [];
        $stack[] = [$sr, $sc];
        $hasAccessed[$sr][$sc] = 1;
        $color = $image[$sr][$sc];
        
        while (!empty($stack)) {
            list($x, $y) = array_pop($stack);
            $image[$x][$y] = $newColor;
            //up
            $ux = $x - 1;
            if (isset($image[$ux][$y]) && $image[$ux][$y] == $color && !isset($hasAccessed[$ux][$y])) {
                $hasAccessed[$ux][$y] = 1;
                $stack[] = [$ux, $y];
            }
            //down
            $dx = $x + 1;
            if (isset($image[$dx][$y]) && $image[$dx][$y] == $color && !isset($hasAccessed[$dx][$y])) {
                $hasAccessed[$dx][$y] = 1;
                $stack[] = [$dx, $y];
            }
            //left
            $ly = $y - 1;
            if (isset($image[$x][$ly]) && $image[$x][$ly] == $color && !isset($hasAccessed[$x][$ly])) {
                $hasAccessed[$x][$ly] = 1;
                $stack[] = [$x, $ly];
            }
            //right
            $ry = $y + 1;
            if (isset($image[$x][$ry]) && $image[$x][$ry] == $color && !isset($hasAccessed[$x][$ry])) {
                $hasAccessed[$x][$ly] = 1;
                $stack[] = [$x, $ry];
            }
        }

        return $image;
    }

    function test() {
        $image = [[1,1,1],[1,1,0],[1,0,1]];
        $newImage = $this->floodFill($image, 1, 1, 2);
        print_r($newImage);
    }
}