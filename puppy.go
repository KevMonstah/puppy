package Puppy

import (
	"log"

	"github.com/KevMonstah/Dog"
)

func Bark() string {
	log.Println("Puppy Bark")
	return "Puppy.Bark"
}

func Barks() string {
	log.Println("Puppy Bark Bark Bark")
	return "Puppy.Barks a lot!!!"
}

func Sleeps(s string) string {
	log.Println("Puppy sleeps")
	Dog.Sleeps("Cinnamon")
	return "My Puppy sleeps"
}

func Runs(s string) string {
	log.Println(s + " the dog runs")
	Dog.Runs("Runes")
	return "My Puppy Runs"
}

func Chills(s string) string {
	log.Println(s + "the dog is chilling")
	Dog.Chills("Nunes")
	return "My Puppy Chills"
}

func From11() {
	log.Println("from version 1.1.0")
}

func From12() {
	log.Println("from version 1.2.0")
}

func From13() {
	log.Println("from version 1.3.0")
}
