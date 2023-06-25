package main

import (
	"fmt"
	"time"

	"learn-go/channel"
	"learn-go/defers"
	"learn-go/interfaces"
	"learn-go/pointer"
)

func main() {
	fmt.Println("Hello Java_Rin!")
}

func interfaceUsage() {
	r := interfaces.Rectangle{Width: 5, Height: 10}
	c := interfaces.Circle{Radius: 15}

	interfaces.Measure(r)
	interfaces.Measure(c)
}

func deferUsage() {
	fmt.Println("Welcome to calculator")
	defer fmt.Println("End")
	defers.Add(20, 10)
}

func pointerUsage() {
	var i int = 1
	fmt.Println("i Value: ", i)
	fmt.Println("i Address: ", &i)
	pointer.ChangeValue(&i)
	fmt.Println("Changed i Value: ", i)
	fmt.Println("i Address: ", &i)
}

func goRoutineUsage() {
	go channel.F("Hello Java_Rin")
	go channel.F("Message 2")
	time.Sleep(5 * time.Second)
}

func channelUsage() {
	ch := make(chan string)
	go channel.Process1(ch, "Data1")
	fmt.Println(<- ch)
	time.Sleep(5 * time.Second)
}