package everyday

import (
	"sort"
	"strings"
)

func removeSubfolders(folders []string) []string {
	sort.Slice(folders, func(i, j int) bool {
		return folders[i] < folders[j]
	})
	prefix := folders[0] + "/"
	ans := []string{folders[0]}
	for i := 1; i < len(folders); i++ {
		if !strings.HasPrefix(folders[i], prefix) {
			prefix = folders[i] + "/"
			ans = append(ans, folders[i])
		}
	}
	return ans
}
