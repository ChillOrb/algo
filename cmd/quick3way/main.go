/*
Sorts a sequence of strings from standard input using quick sort 3 way.


% more tiny.txt

S O R T E X A M P L E

% go run cmd/quick3way/main.go < tiny.txt

A E E L M O P R S T X                 [ one string per line ]



% more words3.txt

bed bug dad yes zoo ... all bad yet


% go run cmd/quick3way/main.go < words3.txt

all bad bed bug dad ... yes yet zoo    [ one string per line ]
*/

package main

import (
	"fmt"

	"github.com/shellfly/algo"
	"github.com/shellfly/algo/stdin"
)

func main() {
	items := stdin.ReadAllStrings()
	fmt.Println("items: ", items)
	algo.QuickSort3way(algo.StringSlice(items))
	if !algo.IsSorted(algo.StringSlice(items)) {
		fmt.Println("Not Sorted")
		fmt.Println(items)
	} else {
		fmt.Println("sorted items: ", items)
	}
}
