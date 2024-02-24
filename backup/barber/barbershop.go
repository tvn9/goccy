package main

import (
	"fmt"
	"time"
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
		fmt.Printf("%s check for clients in waiting room.\n", barber)

		for {
			// if there are no clients, the barber goes to Sleep.
			if len(s.ClientsChan) == 0 {
				fmt.Printf("There is no client, so %s take a nap.\n", barber)
			}

			client, ok := <-s.ClientsChan
			if ok {
				if isSleeping {
					fmt.Printf("%s wakes %s up.\n", client, barber)
					isSleeping = false
				}
				// cut hair
				s.cutHair(barber, client)
			} else {
				// shop is closed, send the barber home and close the goroutine
				s.sendBarberHome(barber)
				return
			}
		}

	}()
}

func (s *BarberShop) cutHair(barber, client string) {
	fmt.Printf("%s is cutting %s's hair.\n", barber, client)
	time.Sleep(s.HairCutDuration)
	fmt.Printf("%s is finish cutting %s's hair.\n", barber, client)
}

func (s *BarberShop) sendBarberHome(barber string) {
	fmt.Printf("%s is going home.", barber)
	s.BarberDoneChan <- true
}

func (s *BarberShop) closeShopForDay() {
	fmt.Println("Closing shop for the day.")

	close(s.ClientsChan)
	s.Open = false

	for a := 1; a <= s.NumberOfBarbers; a++ {
		<-s.BarberDoneChan
	}

	close(s.BarberDoneChan)

	fmt.Println("_________________________")
	fmt.Println("The barbershop is closed!")
}

func (s *BarberShop) addClient(client string) {
	// Print out a message
	fmt.Printf("*** %s arrives!\n", client)

	if s.Open {
		select {
		case s.ClientsChan <- client:
			fmt.Printf("%s takes a seat in the waiting room.\n", client)
		default:
			fmt.Printf("The waiting room is full, so %s leaves.\n", client)
		}
	} else {
		fmt.Printf("The shop is already closed, so %s leaves!\n", client)
	}
}
