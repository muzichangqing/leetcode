mod leetcode_solution;

pub fn add(left: usize, right: usize) -> usize {
    left + right
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::leetcode_solution::Solution;

    #[test]
    fn it_works() {
        let result = add(2, 2);
        assert_eq!(result, 4);
    }

    #[test]
    fn max_sum_div_three() {
        let nums = vec![3, 6, 5, 1, 8];
        let correct = 18;
        let result = Solution::max_sum_div_three(nums);
        assert_eq!(result, correct);
    }
}

