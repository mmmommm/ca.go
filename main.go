package main

import (
    "fmt"
)

type Kind string

const (
    Mountain Kind = "Mountain"
    Road Kind = "Road"
)

func (k Kind) String() string {
    switch k {
    case Mountain:
        return "Mountain"
    case Road:
        return "Road"
    default:
        return "Not exist such kind of bicycle"
    }
}

type Bicycle struct {
		Kind Kind // マウンテン、ロード
		BodyType string
		TotalGear int
}

func main() {
    // Functional Option Pattern (FOP)
    fopApp := NewBicycleWithFOP(Road,
        WithBodyType("aluminum"),
        WithTotalGear(4, 8),
    )

    // Builder Pattern (BP)
    bpApp := NewBicycleWithBP(Mountain).
        WithBodyType("Carbon").
        WithTotalGear(3, 9).
        Build()

    fmt.Printf("%+v\n", fopApp) // &{Course:Premium SubscribeSupportService:true SubscribeMovieService:false SubscribeBackupService:true}
    fmt.Printf("%+v\n", bpApp)  // &{Course:Premium SubscribeSupportService:true SubscribeMovieService:false SubscribeBackupService:true}
}
