package main

import "fmt"

// func main() {
// 	x := make(map[string]int)
// 	x["key"] = 10
// 	fmt.Println(x)	
// }

// func main() {
//     elements := make(map[string]string)
//     elements["H"] = "Hydrogen"
//     elements["He"] = "Helium"
//     elements["Li"] = "Lithium"
//     elements["Be"] = "Beryllium"
//     elements["B"] = "Boron"
//     elements["C"] = "Carbon"
//     elements["N"] = "Nitrogen"
//     elements["O"] = "Oxygen"
//     elements["F"] = "Fluorine"
//     elements["Ne"] = "Neon"
//     if name, ok := elements["F"]; ok {
//     	fmt.Println(name, ok)
//     }
// }

func main() {
	elements := map[string]string{
    "H": "Hydrogen",
    "He": "Helium",
    "Li": "Lithium",
    "Be": "Beryllium",
    "B": "Boron",
    "C": "Carbon",
    "N": "Nitrogen",
    "O": "Oxygen",
    "F": "Fluorine",
    "Ne": "Neon",
}
	if name, ok := elements["F"]; ok {
     	fmt.Println(name, ok)
    }
}


