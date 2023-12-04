package main

import (
	"github.com/fatih/color"
	"time"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%q check the room for clients.", barber)

		for {
			// if there are no clients, the barber goes to sleep
			if len(shop.ClientsChan) == 0 {
				color.Yellow("There is nothing to do, so %q takes a nap.", barber)
				isSleeping = true
			}

			client, shopOpen := <-shop.ClientsChan

			if shopOpen {
				if isSleeping {
					color.Yellow("%q wakes %q up.", client, barber)
					isSleeping = false
				}
				// cut hair
				shop.cutHair(barber, client)
			} else {
				// shop is closed, send the barber home and close the goroutine
				shop.sendBarberHome(barber)
				return
			}
		}
	}()
}

func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("%q is cutting %q's hair.", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("%q is finishing cutting %q's hair.", barber, client)
}

func (shop *BarberShop) sendBarberHome(barber string) {
	color.Cyan("%q is going home.", barber)
	shop.BarbersDoneChan <- true
}

func (shop *BarberShop) shopNowClosed() {
	color.Cyan("Closing shop for the day.")

	close(shop.ClientsChan)
	shop.Open = false

	for a := 1; a < shop.NumberOfBarbers; a++ {
		<-shop.BarbersDoneChan
	}

	close(shop.BarbersDoneChan)

	color.Green("-----------------------------------------")
	color.Green("The barbershop is now closed for the day.")
}

func (shop *BarberShop) addClient(client string) {
	// print out a message
	color.Green("*** %q arrives!", client)

	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			color.Yellow("%q takes a sear in the waiting room.", client)
		default:
			color.Red("The waiting room is full, so %q leaves.", client)
		}
	} else {
		color.Red("The shop is already closed, so %q leaves!", client)
	}
}
