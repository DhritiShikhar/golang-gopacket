// This program prints all devices on your machine

// Author: Dhriti Shikhar

package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

func main() {
	// enumerate all network interfaces on current machine
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	// Interface structure has name, description and address fields
	for _, d := range devices {
		fmt.Println("\n", d.Name)
		fmt.Println("\n", d.Description)

		for _, a := range d.Addresses {
			fmt.Println("IP::::::", a.IP)
			fmt.Println("Netmask::::::", a.Netmask)
		}
		fmt.Println("-----------------------")
	}
}
