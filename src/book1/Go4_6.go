package main

const Pi = 3.14159

func main() {
	//显示类型定义
	const b1 string = "abc"
	//隐式类型定义
	const b2 = "abc"

	const c1 = 2 / 3
	//Const initializer 'getNumber()' is not a constant
	//const c2 = getNumber()

	const Ln2 = 0.128796391273987129837 /
		12389172398
	// this is a precise reciprocal
	const Log2E = 1 / Ln2
	// float constant
	const Billion = 1e9
	const hardEight = (1 << 100) >> 97
	println(hardEight)

	const beef, two, c = "meat", 2, "veg"
	const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6
	const (
		Monday2, Tuesday2, Wednesday2 = 1, 2, 3
		Thursday2, Friday2, Saturday2 = 4, 5, 6
	)

	println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
	//enum
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)

	const (
		a1 = iota
		a2 = iota
		a3 = iota
	)
	const (
		a11 = iota
		a12
		a13
	)
	println(a1, a2, a3)
	println(a11, a12, a13)

	type Color int
	const (
		RED    Color = iota // 0
		ORANGE              // 1
		YELLOW              // 2
		GREEN               // 3
		BLUE                // 4
		INDIGO              // 5
		VIOLET              // 6
	)
	println(RED, ORANGE, YELLOW, GREEN)

}
func getNumber() int {
	return 1
}
