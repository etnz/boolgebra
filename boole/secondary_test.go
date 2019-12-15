package boole

import "fmt"

func ExampleQuantified() {

	A := ID("A")
	B := ID("B")
	C := ID("C")
	fmt.Println(Exactly(2, A, B, C))
	fmt.Println(AtLeast(2, A, B, C))
	fmt.Println(AtMost(2, A, B, C))

	//Output:
	// "A" & "B" & !"C" | !"A" & "B" & "C" | "A" & !"B" & "C"
	// "A" & "B" | "B" & "C" | "A" & "C"
	// !"C" | !"A" | !"B"
}
