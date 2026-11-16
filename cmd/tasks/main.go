package main

import (
	"fmt"

	"github.com/viktorbeznosov/job4j_go_lang_base/internal/base"
)

func main() {
    memorySender := &base.MemorySender{}
    memoryNotificationService, err := base.NewNotificationService(memorySender)
    if (err != nil) {
        fmt.Printf("не удалось создать сервис: %v\n", err)
        return
    }
    memoryNotificationService.Notify("Somebody", "Hello")
    fmt.Println(memorySender.Messages)

    countSender := &base.CountSender{}
    countNotificationService, err := base.NewNotificationService(countSender)
    if (err != nil) {
        fmt.Printf("не удалось создать сервис: %v\n", err)
        return
    }
    countNotificationService.Notify("Somebody", "Hello")
    countNotificationService.Notify("Somebody", "Hello")
    fmt.Println(countSender.Calls)
}


