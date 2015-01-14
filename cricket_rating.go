package main

import "fmt"

func max(a,b int32) int32 {
	if a>b {
		return a
	} else {
		return b
	}
}

func main() {
	var i,n int32
	fmt.Scanf("%d",&n)
	var arr [100007]int
	for i=0;i<n;i++ {
		fmt.Scanf("%d",&arr[i])
	}
	var maxsum, runningsum int32
	maxsum = 0
	runningsum = 0
	for i=0;i<n;i++ {
		runningsum = max(0,runningsum + int32(arr[i]))
		maxsum = max(maxsum, runningsum)
	}
	fmt.Println(maxsum)
}
