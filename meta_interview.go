package main

import (
	"fmt"
	"strings"
)
func main() {
	mag := "Meta Platforms, Inc. to Change Ticker Symbol to 'META' on June 9 \nMENLO PARK, Calif., May 31, 2022 /PRNewswire/ -- Meta Platforms, Inc. (Nasdaq: FB) today announced that its Class A common stock will begin trading on NASDAQ under the ticker symbol 'META' prior to market open on June 9, 2022. This will replace the company's current ticker symbol 'FB', which has been used since its initial public offering in 2012. The new ticker symbol aligns with the company's rebranding from Facebook to Meta, announced on October 28, 2021."
	letters := make(map[byte]int8)
	mag = strings.ToLower(mag)
	mag = strings.ReplaceAll(mag, ".", "")
	mag = strings.ReplaceAll(mag, ".", "")
	for _, letter := range []byte(mag) {
		letters[letter]++
	}
	fmt.Println(letters)
	//get user input 
	var user string = ""
	fmt.Scan(&user)
	user = strings.ToLower(user)
	okay := true
	for _,letter := range []byte(user) {
		if letters[letter] - 1 != 0 {
			letters[letter]--
		}else {
			okay = false
		}
	}
	fmt.Println(okay)
}


func allow(val string) string{
	var newValue strings.Builder
	allow := string[]{"A", "B", C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z.}

}