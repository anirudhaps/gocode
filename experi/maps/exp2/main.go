package main

import (
	"fmt"
)

type vertex struct {
	x int
	y int
}

//var coords = map[string]vertex{
//	"longitude": vertex{
//		23, 39,
//	},
//	"latitude": vertex{
//		34, 91,
//	},
//}

func main() {
	// coords := map[string]vertex{
	// 	"longitude": vertex{
	// 		23, 91,
	// 	},
	// 	"latitude": vertex{
	// 		78, 65,
	// 	},
	// }
	// coords := map[string]vertex{
	// 	"longitude": {
	// 		24, 67,
	// 	},
	// 	"latitude": {
	// 		45, 89,
	// 	},
	// }
	// coords := map[string]vertex{
	// 	"longitude": {23, 91},
	// 	"latitude":  {34, 56},
	// }
	var coords = map[string]vertex{
		"longitude": {23, 45},
		"latitude":  {47, 89},
	}
	for key, val := range coords {
		fmt.Println(key, val)
	}
	fmt.Println("Keys and values:")
	fmt.Println("longitude.x:", coords["longitude"].x)
	fmt.Println("latitude.x:", coords["latitude"].x)
	var elem, ok = coords["longitude"]
	fmt.Println("elem:", elem, "ok:", ok)
	if ok {
		fmt.Println("Sum of coords:", elem.x+elem.y)
	}
	coords["base"] = vertex{0, 0}
	fmt.Println(coords)
	delete(coords, "latitude")
	fmt.Println(coords)
}
