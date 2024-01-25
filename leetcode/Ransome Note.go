func canConstruct(ransomNote string, magazine string) bool {
	rmap := make(map[string]int)
	mmap := make(map[string]int)

	for i := 0; i < len(ransomNote); i++ {
		if _, ok := rmap[string(ransomNote[i])]; ok == true {
			rmap[string(ransomNote[i])] += 1
		} else {
			rmap[string(ransomNote[i])] = 1
		}
	}
	for i := 0; i < len(magazine); i++ {
		if _, ok := mmap[string(magazine[i])]; ok == true {
			mmap[string(magazine[i])] += 1
		} else {
			mmap[string(magazine[i])] = 1
		}
	}
	for key, value := range rmap {
		if mmap[key] < value {
			return false
		}
	}
	return true
}