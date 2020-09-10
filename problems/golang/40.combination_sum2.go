package leetcode

var combinations2 [][]int

func combinationSum2(candidates []int, target int) [][]int {
	combinations2 = [][]int{}
	freqs := make(map[int]int) // key 数字  value 该数字在candidates中出现的次数
	for _, element := range candidates {
		freqs[element]++
	}
	freqCandidates := make([][]int, len(freqs))
	freqCandidatesIndex := 0
	for key, value := range freqs {
		freqCandidates[freqCandidatesIndex] = []int{key, value}
		freqCandidatesIndex++
	}
	backtraceForCombinationSum2(freqCandidates, make([]int, 0), target, 0)
	return combinations2
}

func backtraceForCombinationSum2(freqCandidates [][]int, path []int, rest int, pos int) {

	for i := pos; i < len(freqCandidates); i++ {
		num := freqCandidates[i][0]
		maxFreq := freqCandidates[i][1]
		for freq := 1; freq <= maxFreq; freq++ {
			if num*freq > rest {
				break // 都是正整数，没必要往下乘
			} else if num*freq == rest {
				combination := make([]int, len(path)+freq)
				for index, value := range path {
					combination[index] = value
				}
				for index := len(path); index < len(path)+freq; index++ {
					combination[index] = num
				}
				combinations2 = append(combinations2, combination)
				break
			} else {
				for count := 0; count < freq; count++ {
					path = append(path, num)
				}
				backtraceForCombinationSum2(freqCandidates, path, rest-freq*num, i+1)
				path = path[:len(path)-freq]
			}
		}

	}

}
