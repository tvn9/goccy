// Sleeping Barber Challenge
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// veriables for time delay use for simulate the shop capacity, waiting time
// arrival rate, hair cut duration, and time the shop open...
var (
	numberOfBarber  = 0
	seatingCapacity = 20
	arrivalRate     = 100
	cutDuration     = time.Second * time.Duration(rand.Intn(3))
	timeOpen        = time.Second * time.Duration(rand.Intn(10))
)

func main() {
	// print welcome message
	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("---------------------------")

	// create channels if we need any
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	// create the barbershop
	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: numberOfBarber,
		BarberDoneChan:  doneChan,
		ClientsChan:     clientChan,
		Open:            true,
	}

	color.Green("The shop is open for the day!")

	// add barbers
	shop.addBarber("Thanh")
	shop.addBarber("Mike")
	shop.addBarber("Kevin")
	shop.addBarber("Kim")
	shop.addBarber("Tim")

	// start the barbershop as a goroutine
	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShop()
		closed <- true
	}()

	// add clients
	i := 1

	go func() {
		for {
			// get a random number
			ranNum := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(ranNum)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	// block until the barbershop is closed
	// time.Sleep(3 * time.Second)
	<-closed

}
