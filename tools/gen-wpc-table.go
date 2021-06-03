package main

import (
	"fmt"
	"log"
)

type devmap struct {
	id  string
	num int
}

var flipperCoils = []string{
	"FLLH",
	"FLLM",
	"FLRH",
	"FLRM",
	"FULH",
	"FULM",
	"FURH",
	"FURM",
}

func main() {
	log.SetFlags(0)

	// outFile := os.Args[1]
	// f, err := os.Create(outFile)
	// if err != nil {
	// 	log.Fatalf("unable to open file: %v", err)
	// }
	// os.Stdout = f

	devices := make([]devmap, 0, 0)

	fmt.Println("package wpc")
	fmt.Println()

	fmt.Println("// Coils and flashers")
	fmt.Println("const (")
	for i := 1; i <= 44; i++ {
		var num int
		if i <= 28 {
			num = (i - 1) + 30
		}
		if i >= 29 && i <= 36 {
			num = (i - 29) + 32
		}
		if i >= 37 {
			num = (i - 37) + 144
		}
		id := fmt.Sprintf("C%02d", i)
		devices = append(devices, devmap{id: id, num: num})
		fmt.Printf("%v = %d\n", id, num)
	}
	fmt.Println(")")
	fmt.Println()

	fmt.Println("// General illumination")
	fmt.Println("const (")
	for i := 1; i <= 5; i++ {
		num := (i - 1) + 72
		id := fmt.Sprintf("G%02d", i)
		devices = append(devices, devmap{id: id, num: num})
		fmt.Printf("%v = %d\n", id, num)
	}
	fmt.Println(")")
	fmt.Println()

	fmt.Println("// Lamps")
	fmt.Println("const (")
	num := 80
	for c := 1; c <= 8; c++ {
		for r := 1; r <= 8; r++ {
			id := fmt.Sprintf("L%0d%d", c, r)
			devices = append(devices, devmap{id: id, num: num})
			fmt.Printf("%v = %d\n", id, num)
			num++
		}
	}
	fmt.Println(")")
	fmt.Println()

	fmt.Println("// Switch matrix")
	fmt.Println("const (")
	num = 32
	for c := 1; c <= 8; c++ {
		for r := 1; r <= 8; r++ {
			id := fmt.Sprintf("S%0d%d", c, r)
			devices = append(devices, devmap{id: id, num: num})
			fmt.Printf("%v = %d\n", id, num)
			num++
		}
		num += 8
	}
	fmt.Println(")")
	fmt.Println()

	fmt.Println("// Dedicated switches")
	fmt.Println("const (")
	for i := 1; i <= 8; i++ {
		num := (i - 1) + 8
		id := fmt.Sprintf("SD%d", i)
		devices = append(devices, devmap{id: id, num: num})
		fmt.Printf("%v = %d\n", id, num)
	}
	fmt.Println(")")
	fmt.Println()

	fmt.Println("// Flipper switches")
	fmt.Println("const (")
	for i := 1; i <= 8; i++ {
		num := (i - 1) + 0
		id := fmt.Sprintf("SF%d", i)
		devices = append(devices, devmap{id: id, num: num})
		fmt.Printf("%v = %d\n", id, num)
	}
	fmt.Println(")")
	fmt.Println()

	fmt.Println("// Flipper coils")
	fmt.Println("const (")
	for i := 1; i <= 8; i++ {
		num := (i - 1) + 35
		id := flipperCoils[i-1]
		devices = append(devices, devmap{id: id, num: num})
		fmt.Printf("%v = %d\n", id, num)
	}
	fmt.Println(")")
	fmt.Println()

	fmt.Println("var Devices = map[string]uint8 {")
	for _, d := range devices {
		fmt.Printf("\"%v\": %d,\n", d.id, d.num)
	}
	fmt.Println("}")
	fmt.Println()

	fmt.Println("var SwitchNames = map[uint8]string {")
	for _, d := range devices {
		if d.id[0] == 'S' {
			fmt.Printf("%d: \"%s\",\n", d.num, d.id)
		}
	}
	fmt.Println("}")
	fmt.Println()
}
