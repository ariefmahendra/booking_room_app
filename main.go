package main

import "booking-room/delivery"

func main() {
	delivery.NewServer().Run()
}