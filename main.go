package main

import (
    "fmt"
)

type Course int

const (
    Basic Course = iota
    Premium
)

func (c Course) String() string {
    switch c {
    case Basic:
        return "Basic"
    case Premium:
        return "Premium"
    default:
        return ""
    }
}

// Application は通信キャリアの契約申し込みをイメージしたものです。
type Application struct {
    Course Course                // ベーシックプランとプレミアプランがある
    SubscribeSupportService bool // サポートサービスのオプション
    SubscribeMovieService bool   // 動画サービスのオプション
    SubscribeBackupService bool  // データバックアップサービスのオプション
}

func main() {
    // Functional Option Pattern (FOP)
    fopApp := NewApplicationWithFOP(Premium,
        WithBackupService(true),
        WithSupport(true),
        WithMovie(false),
    )

    // Builder Pattern (BP)
    bpApp := NewApplicationWithBP(Premium).
        WithBackupService(true).
        WithSupport(true).
        WithMovie(false).
        Build()

    fmt.Printf("%+v\n", fopApp) // &{Course:Premium SubscribeSupportService:true SubscribeMovieService:false SubscribeBackupService:true}
    fmt.Printf("%+v\n", bpApp)  // &{Course:Premium SubscribeSupportService:true SubscribeMovieService:false SubscribeBackupService:true}
}
