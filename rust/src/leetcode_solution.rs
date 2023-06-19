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
}