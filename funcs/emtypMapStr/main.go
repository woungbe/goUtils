package main

import "fmt"

type Identifications struct {
	Identifications []ChainalysisData `json:"identifications"`
}

type ChainalysisData struct {
	Category    string `json:"category" db:"category"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Url         string `json:"url" db:"url"`
}

func main() {
	emtypmapstr()
}

func emtypmapstr() {

	var args Identifications
	fmt.Printf("%+v", args)

	if len(args.Identifications) == 0 {
		fmt.Println("정상이구요")
	} else {
		fmt.Println("be정상이구요")
	}

}
