mod leetcode_solution;

pub fn add(left: usize, right: usize) -> usize {
    left + right
}

pub fn strs_strings(strs: Vec<&str>) -> Vec<String> {
    let mut res = Vec::new();
    for str in strs {
        res.push(String::from(str))
    }
    return res
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

    #[test]
    fn maximum_value() {
        let strs = vec!["alic3", "bob", "0000003", "4", "00000"];
        let correct = 5;
        let result = Solution::maximum_value(strs_strings(strs));
        assert_eq!(correct, result);
    }
}

