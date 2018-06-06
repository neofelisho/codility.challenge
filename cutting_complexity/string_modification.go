package cuttingcomplexity

// StringModification is used to calculate the minimun changes of the string,
// to satisfy that the Longest interval of letters 'M' = K.
// The string S is consist of 'L' and 'M', and we can change any "M" to "L" or any "L" to "M".
// The objective of this challenge is to calculate the minimum number of changes
// we have to make in order to achieve the desired longest M-interval length K
func StringModification(S string, K int) int {
	if len(S) < K {
		return -1
	}

	isOk := false
	answer := 0
	counter := 0
	for i := 0; i < len(S); i++ {
		if S[i] == 'M' {
			counter++
		} else {
			counter = 0
		}
		if counter == K {
			isOk = true
		}
		if counter > K {
			S = S[:i] + "L" + S[i+1:]
			answer++
			counter = 0
		}
	}

	if isOk {
		return answer
	}

	counter = 0
	for i := 0; i < K; i++ {
		if S[i] == 'L' {
			counter++
		}
	}
	if S[K] == 'M' {
		answer = counter + 1
	} else {
		answer = counter
	}

	newCounter := 0
	for i := K; i < len(S); i++ {
		if S[i-K] == 'L' {
			counter--
		}
		if S[i] == 'L' {
			counter++
		}
		newCounter = 0
		if i-K != 0 && S[i-K-1] == 'M' {
			newCounter++
		}
		if i < len(S)-1 && S[i+1] == 'M' {
			newCounter++
		}
		if (counter + newCounter) < answer {
			answer = counter + newCounter
		}
	}

	return answer
}
