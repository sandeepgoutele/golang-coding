package modules

func GetMinString(str, pat string) string {
	hashP := make(map[byte]int)
	for idx := 0; idx < len(pat); idx++ {
		hashP[pat[idx]]++
	}

	hashS := make(map[byte]int)
	start, count, minLen, resultStr := 0, 0, len(str), ""
	for idx := 0; idx < len(str); idx++ {
		hashS[str[idx]]++
		if _, ok := hashP[str[idx]]; ok || hashP[str[idx]] >= hashS[str[idx]] {
			count++
		}

		if count == len(pat) {
			for {
				_, ok := hashP[str[start]]
				if !ok || hashS[str[start]] > hashP[str[start]] {
					if hashS[str[start]] > hashP[str[start]] {
						hashS[str[start]]--
					}
					start++
				} else {
					break
				}
			}
			if idx-start+1 < minLen {
				minLen = idx - start + 1
				resultStr = str[start : idx+1]
			}
		}
	}
	return resultStr
}
