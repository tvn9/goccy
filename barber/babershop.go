package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarberDoneChan  chan bool
	ClientsChan     chan string
	Open            bool
}

func (s *BarberShop) addBarber(barber string) {
	s.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s Check for clients in waiting room.", barber)

		for {
			// if there are no client, the barber goes to sleep
			if len(s.ClientsChan) == 0 {
				color.Yellow("There is no client, so %s take a nap.", barber)
				isSleeping = true
			}

			client, shopOpen := <-s.ClientsChan
			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up.", client, barber)
					isSleeping = false
				}
				// cut hair
				s.cutHair(barber, client)
			} else {
				// shop is closed, so barber goes home
				s.sendBarberHome(barber)
				return
			}
		}
	}()
}

func (s *BarberShop) cutHair(barber string, client string) {
	color.Green("%s is cutting %s's hair", barber, client)
	time.Sleep(s.HairCutDuration)
	color.Green("%s is finished cutting %s's hair.", barber, client)
}

func (s *BarberShop) sendBarberHome(barber string) {
	color.Cyan("%s is going home.", barber)
	s.BarberDoneChan <- true
}

func (s *BarberShop) closeShop() {
	color.Cyan("Closing shop for the day.")

	close(s.ClientsChan)
	s.Open = false

	for i := 1; i <= s.NumberOfBarbers; i++ {
		<-s.BarberDoneChan
	}
	close(s.BarberDoneChan)

	color.Green("The Barbershop is now closed!")
}

func (s *BarberShop) addClient(client string) {
	// print out a message
	color.Green("*** %s arrives!", client)
	if s.Open {
		select {
		case s.ClientsChan <- client:
			color.Yellow("%s takes a seat in the waiting room.", client)
		default:
			color.Red("The waiting room is full, so %s leaves.", client)
		}
	} else {
		color.Red("The shop is already closed, so %s leaves!", client)
	}
}
