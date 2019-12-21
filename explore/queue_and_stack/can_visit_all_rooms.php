<?php

class Solution {

    /**
     * @param Integer[][] $rooms
     * @return Boolean
     */
    function canVisitAllRooms($rooms) {
        $keys = [0];
        $openRooms = [];
        while(!empty($keys)) {
            $key = array_pop($keys);
            if (!isset($openRooms[$key])) {
                // 开门拿钥匙
                $openRooms[$key] = 1;
                $keys = array_merge($keys, $rooms[$key]);
            }
        }
        return count($rooms) == count($openRooms);
    }
}