package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/k0mpilator/bolidprnt-emulator/internal/config"
	"github.com/rs/zerolog/log"
	"github.com/tarm/serial"
	"golang.org/x/text/encoding/charmap"
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
	}

	for {
		c := rand.Intn(10)
		d := rand.Intn(10)

		t := time.Now().Format("02.01 15:04:05")
		qf := conf.Qualifiers[rand.Intn(len(conf.Qualifiers))]
		dt := conf.Detector[rand.Intn(len(conf.Detector))]
		b := []byte(fmt.Sprintf("|%s|%s|  %v    |   %v|%s|№ ПАРОЛЯ: 7     |\r\n", t, qf, c, d, dt))

		encoder := charmap.CodePage866.NewEncoder()
		eb, err := encoder.Bytes(b)
		if err != nil {
			log.Error().Err(err).Msg("")
			continue
		}

		_, err = s.Write(eb)
		if err != nil {
			log.Error().Err(err).Msg("")
		}

		/*detector := chardet.NewTextDetector()
		result, err := detector.DetectBest([]byte(eb))
		if err != nil {
			log.Error().Err(err).Msg("")
		}
		fmt.Println(result)*/

		fmt.Println(eb)

		fmt.Printf("|%s|%s|  %v    |   %v|%s|№ ПАРОЛЯ: 7     |\r\n", t, qf, c, d, dt)
		time.Sleep(time.Second * 5)
	}
}
