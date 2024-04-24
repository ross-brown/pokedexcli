package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config) error {
	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range res.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = res.Next
	cfg.prevLocationAreaURL = res.Previous
	return nil
}

func callbackMapb(cfg *config) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}
	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range res.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = res.Next
	cfg.prevLocationAreaURL = res.Previous
	return nil
}