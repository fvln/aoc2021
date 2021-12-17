package day17

import "fmt"

type Trench struct {
	minX int
	maxX int
	minY int
	maxY int
}

func max(a int, b int) int {
	if a > b { return a }
	return b
}

func probeTouchesTrench(velX, velY int, trench Trench) bool {
	for posX, posY := 0,0; posX <= trench.maxX && posY >= trench.minY; {
		posX += velX
		posY += velY

		//fmt.Println(posX, posY)

		if trench.minX <= posX && posX <= trench.maxX {
			if trench.minY <= posY && posY <= trench.maxY {
				return true
			}
		}

		velX = max(0, velX - 1)
		velY -= 1

	}

	return false
}

func probeTouchesTrenchVertically(velY int, trench Trench) bool {
	for posY := 0; posY >= trench.minY; {
		posY += velY

		if trench.minY <= posY && posY <= trench.maxY {
			return true
		}

		velY -= 1
	}

	return false
}

func getMaxY(velY int) int {
	posY := 0

	for ; velY >=0; velY -= 1 {
		posY += velY
	}

	return posY
}

func getMaxXVelocity(trench Trench) int {

	for velX := 0; ; velX++ {

		posX := 0
		for tmpVelX := velX ; tmpVelX > 0; tmpVelX-- {
			posX += tmpVelX
		}

		if posX > trench.maxX {
			// We've gone too far!
			return velX
		}
	}
}

func step1(trench Trench) {
	for y:=0; y < 123; y++ {
		if probeTouchesTrenchVertically(y, trench) {
			fmt.Println("Step1 : Candidate Y velocity:", y, getMaxY(y))
		}
	}
}

func step2(trench Trench) {

	nbPossibleVelocities := 0

	for y:=trench.minY; y < 123; y++ {
		// Approximation grossière FTW
		for x:=0; x <= trench.maxX; x++ {
			if probeTouchesTrench(x, y, trench) {
				//fmt.Println("Candidate velocity:", x, y)
				nbPossibleVelocities++
			}

		}
	}

	fmt.Println("Step2:", nbPossibleVelocities)
}

func Solve() {
	//testTrench := Trench{
	//	minX: 20,
	//	maxX: 30,
	//	minY: -10,
	//	maxY: -5,
	//}
	
	prodTrench := Trench{
		minX: 124,
		maxX: 174,
		minY: -123,
		maxY: -86,
	}

	step1(prodTrench)
	step2(prodTrench)
}