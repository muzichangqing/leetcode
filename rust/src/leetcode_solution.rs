use std::cmp;

pub struct Solution {}

impl Solution {
   
    // 20230619 1262. 可被三整除的最大和
    pub fn max_sum_div_three(nums: Vec<i32>) -> i32 {
        let mut f = [0, -0x3f3f3f3f, -0x3f3f3f3f];
        for num in nums {
            let mut g = [0; 3];
            for i in 0..3 {
                g[((i+num) % 3) as usize] = cmp::max(f[((i+num)%3) as usize], f[i as usize] + num)
            }
            f = g;
        }
        f[0]
    }

    // 2496. 数组中字符串的最大值
    pub fn maximum_value(strs: Vec<String>) -> i32 {
        let mut max = 0;

        for str in &strs {
            let num: i32 = match str.parse() {
                Ok(num) => num,
                Err(_) => {
                    str.chars().count() as i32
                },
            };
            max = cmp::max(max, num)
        }
        max
    }
    
    pub fn check_overlap(radius: i32, x_center: i32, y_center: i32, x1: i32, y1: i32, x2: i32, y2: i32) -> bool {
        let mut dist = 0;
        if x_center < x1 || x_center > x2 {
            dist += cmp::min((x1 - x_center).pow(2), (x2 - x_center).pow(2));
        }
        if y_center < y1 || y_center > y2 {
            dist += cmp::min((y1-y_center).pow(2), (y2-y_center).pow(2));
        }
        return dist <= radius.pow(2);
    }
}