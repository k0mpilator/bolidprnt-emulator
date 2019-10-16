package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/k0mpilator/bolidprnt-emulator/internal/config"
	"github.com/tarm/serial"
)

func main() {

	// Read config yaml
	conf := config.NewConfig("config.yml")

	// Serial port settings
	config := &serial.Config{
		Name:        conf.PortName,
		Baud:        conf.Baud,
		ReadTimeout: conf.ReadTimeout,
		Size:        conf.Size,
	}

	s, err := serial.OpenPort(config)
	if err != nil {
		fmt.Printf("Error open serial port %v", err)
		log.Fatalf("can't open serial port %v", err)
	}

	for {
		d := rand.Intn(10)
		t := time.Now().Format("01.02 15:04:05")
		e := conf.Qualifiers[rand.Intn(len(conf.Qualifiers))]
		//sec := section[rand.Intn(len(section))]
		b := []byte(fmt.Sprintf("|%s|%s|  %v    |   %v|0|№ ПАРОЛЯ: 7     |\r\n", t, e, d, d))
		_, err := s.Write(b)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("|%s|%s|  %v    |   %v|0|№ ПАРОЛЯ: 7     |\r\n", t, e, d, d)
		time.Sleep(time.Second * 1)
	}
}
