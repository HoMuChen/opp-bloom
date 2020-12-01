package main

import (
        "fmt"

        "github.com/HoMuChen/opp-bloom"
)

func main() {
        set := oppbloom.New(1000000)

        set.Add([]byte(`b41401af-0383-4016-af77-d1cc96ad77f8`))
        set.Add([]byte(`5b679c1b-2ed9-4ea2-9071-c93e0a73fed1`))

        fmt.Println(set.Contain([]byte(`b41401af-0383-4016-af77-d1cc96ad77f8`)))//true
        fmt.Println(set.Contain([]byte(`Hi! This is from opp-bloom`)))          //false
}
