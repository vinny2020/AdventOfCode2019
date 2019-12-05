package main

import "math"
import "fmt"

func FuelCalc(mass float64) int {
	//To find the fuel required for a module, take its mass, divide by three, round down, and subtract 2.
	var fuel float64 = math.Trunc(mass / 3)
	return int(fuel - 2)
}


func main() {

	type module struct {
		mass float64
		fuel int
	}

	var mod1 = module{12, 0}
	mod1.fuel = FuelCalc(mod1.mass)
	var mod2 = module{14, 0}
	mod2.fuel = FuelCalc(mod2.mass)
	var mod3 = module{1969, 0}
	mod3.fuel = FuelCalc(mod3.mass)
	var mod4 = module{100756, 0}
	mod4.fuel = FuelCalc(mod4.mass)

	var mods [] module
	mods = append(mods, mod1)
	mods = append(mods, mod2)
	mods = append(mods, mod3)
	mods = append(mods, mod4)

	var totalModsFuel = 0

	for i := 0; i < len(mods); i++ {
		totalModsFuel += int(mods[i].mass)
		fmt.Println("the raw mass is" ,  int(mods[i].mass))
		fmt.Println("the fuel is" ,  int(mods[i].fuel))

	}

	fmt.Println( "The total fuel is ", totalModsFuel)

}
