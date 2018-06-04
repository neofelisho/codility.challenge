package zinc2018

import "strconv"

// Solution is used to find that: Kate was given a birthday gift of three theater tickets.
// Now she is browsing the theater program for the next N days.
// On the program, performances are named by integers. Every day, one performance is staged.
// Kate wants to choose three days (not necessarily consecutive) to go to the theater.
// In how many ways can she use her tickets? Two ways are different if the sequences of watched performances are different.
// Kate likes the theater, so she may watch one performance more than once.
// For example, if N = 4 and theater program looks as following: [1, 2, 1, 1],
// Kate has four possibilities to choose the dates: [1, 2, 1, 1], [1, 2, 1, 1], [1, 2, 1, 1], and [1, 2, 1, 1],
// but they create only three different sequences: (1, 2, 1), (1, 1, 1) and (2, 1, 1).
//The correct answer for this example is 3. Notice that the order of performances matters, so the first and the last sequences are considered different.
func TheaterTickets(A []int) int {
	temp := []string{}
	counter := 0
	for i := 0; i < len(A)-2; i++ {
		for j := len(A) - 1; j >= i+2; j-- {
			current := strconv.Itoa(A[i]) + "," + strconv.Itoa(A[j])
			if contains(temp, current) {
				continue
			}
			temp = append(temp, current)
			between := []int{}
			for k := 1; k < j-i; k++ {
				between = appendIfMissing(between, A[i+k])
			}
			counter += len(between)
		}
	}
	return counter
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func appendIfMissing(slice []int, i int) []int {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}
