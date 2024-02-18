package main

import (
	"fmt"
	"math/rand"
	"time"
)

// variables
var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {

	// send our random number generator

	// print welcome message
	fmt.Println("The Sleeping Barber Problem")
	fmt.Println("___________________________")

	// create channels if we need any
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	// create data structure for barbershop
	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarberDoneChan:  doneChan,
		Open:            true,
	}

	fmt.Println("The shop is open for the day!")

	// add barbers
	shop.addBarber("Frank")
	shop.addBarber("Kim")
	shop.addBarber("Jason")
	shop.addBarber("Lin")
	shop.addBarber("Tao")
	shop.addBarber("Mai")
	shop.addBarber("Toan")

	// start the barbershop as a goroutine
	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	// add clients
	i := 1

	go func() {
		for {
			// get a random number
			randNum := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randNum)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	// block until the barber shop is closed
	<-closed
}
