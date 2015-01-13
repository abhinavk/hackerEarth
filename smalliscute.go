package main
import "fmt"

func main() {
	var n,a,b int

	var divisibles []int
	fmt.Scanf("%d %d %d",&n,&a,&b)
	for i:=0;i<n;i++ {
		if i%a==0 {
			divisibles = append(divisibles,i)
		} else if i%b==0 {
			divisibles = append(divisibles,i)
		}
	}
	var sum uint64
	for i:=0;i<len(divisibles);i++ {
		sum += uint64(divisibles[i])
	}
	fmt.Println(sum)
}
