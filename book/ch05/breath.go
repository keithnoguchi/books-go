package ch05

func BreathFirst(f func(string) []string, workList []string) {
	seen := make(map[string]bool)

	for len(workList) > 0 {
		items := workList
		workList = nil
		for _, item := range items {
			if seen[item] {
				continue
			}
			seen[item] = true
			workList = append(workList, f(item)...)
		}
	}
}
