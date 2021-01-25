package main

import "fmt"

/// database
// List Customer And Product Without Sale

//To solve this, we will follow these steps −
//
//Take a map rem to store remainders. Set ans := 0
//for all elements i in time −
//	if i is divisible by 0 and 0 in rem, then ans := ans + rem[0]
//	o otherwise when 60 – (i mod 60) in rem, then ans := ans + rem[60 – (i mod 60)]
//	if i mod 60 in rem, then rem[i mod 60] := rem[i mod 60] + 1
//	otherwise rem[i mod 60] := 1
//return the ans

func playlist(songs []int32) int64 {
	// Write your code here
	var numOfSongs int64
	mapOfSongs := make(map[int32]int64)
	for _, song := range songs {
		if song % 60 == 0 && mapOfSongs[0] != 0 {
			numOfSongs += mapOfSongs[0]
		} else if mapOfSongs[60-(song%60)] != 0 {
			numOfSongs += mapOfSongs[60-(song%60)]
		}
		if mapOfSongs[song%60] != 0 {
			mapOfSongs[song%60] += 1
		} else {
			mapOfSongs[song%60] = 1
		}
	}
	return numOfSongs
}

func main() {
	arr := []int32{5, 30, 20, 150, 100, 40}
	fmt.Println(playlist(arr))
}
