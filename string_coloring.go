package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
)

type rgb struct {
	r int
	g int
	b int
}

func print_colored(text string) {
	freq := 0.1
	position := 0
	const dfreq = 0.01

	for _, c := range text {
		if c == '\n' {
			freq += dfreq
			position = 0
		}
		colors, err := rainbow(freq, position)
		position += 1
		if err != nil {
			log.Fatal(err)
		}
		esc_seq, err := make_color_esc_seq(colors)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v%c", esc_seq, c)
	}
	fmt.Printf("\n")
}

func rainbow(freq float64, i int) (rgb, error) {
	var colors rgb
	colors.r = int(math.Sin(freq*float64(i)+0)*127 + 128)
	colors.g = int(math.Sin(freq*float64(i)+2*math.Pi/3)*127 + 128)
	colors.b = int(math.Sin(freq*float64(i)+4*math.Pi/3)*127 + 128)

	// These errors should be mathematically impossible, but just in case
	// there was an error in my math.
	if colors.r < 0 || colors.g < 0 || colors.b < 0 {
		err := errors.New("Negative RGB Value found")
		return colors, err
	}

	if colors.r > 256 || colors.g > 256 || colors.b > 256 {
		err := errors.New("RGB Value exceeds 256")
		return colors, err
	}
	return colors, nil
}

func make_color_esc_seq(colors rgb) (string, error) {
	// Validate the colors struct
	n_errors := 0
	if colors.r < 0 || colors.r > 256 {
		n_errors += 1
	}
	if colors.g < 0 || colors.g > 256 {
		n_errors += 1
	}
	if colors.b < 0 || colors.b > 256 {
		n_errors += 1
	}
	if n_errors > 0 {
		err := errors.New("Invalid rgb struct provided")
		return "", err
	}

	seq := "\x1b[38;2;"
	seq += strconv.Itoa(colors.r) + ";"
	seq += strconv.Itoa(colors.g) + ";"
	seq += strconv.Itoa(colors.b) + "m"
	return seq, nil
}
