package main

import (
	"errors"
	"fmt"
	"strings"
)

func commandInspect(state *config, opts ...string) error {
	if len(opts) == 1 {
		fmt.Print("Need a pokemon name with inspect command")
		return errors.New("")
	}

	if x, ok := state.Pokedex[opts[1]]; ok {
		//		fmt.Printf("Pokemon %v is in Pokedex", x.Name)
		//		return errors.New("")
		height := x.Height
		weight := x.Weight
		var hp, attack, defense, specialattack, specialdefense, speed int
		types := make([]string, 0)
		for _, j := range x.Stats {
			if j.Stat.Name == "hp" {
				hp = j.BaseStat
			} else if j.Stat.Name == "attack" {
				attack = j.BaseStat
			} else if j.Stat.Name == "defense" {
				defense = j.BaseStat
			} else if j.Stat.Name == "special-attack" {
				specialattack = j.BaseStat
			} else if j.Stat.Name == "special-defense" {
				specialdefense = j.BaseStat
			} else if j.Stat.Name == "speed" {
				speed = j.BaseStat
			} else {
				continue
			}
		}
		for _, k := range x.Types {
			types = append(types, " - "+k.Type.Name)
		}

		fmt.Printf("Name: %v\n", x.Name)
		fmt.Printf("Height: %v\n", height)
		fmt.Printf("Weight: %v\n", weight)
		fmt.Printf("Stats:\n")
		fmt.Printf(" -hp: %v\n", hp)
		fmt.Printf(" -attack: %v\n", attack)
		fmt.Printf(" -defense: %v\n", defense)
		fmt.Printf(" -special-attack: %v\n", specialattack)
		fmt.Printf(" -special-defense: %v\n", specialdefense)
		fmt.Printf(" -speed: %v\n", speed)

		results := strings.Join(types, "\n")
		fmt.Print("Types:\n")
		fmt.Print(results)
		return errors.New("")

	} else {
		fmt.Printf("you have not caught that pokemon")
		return errors.New("")
	}

}
