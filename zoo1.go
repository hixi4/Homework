package main

import (
	"fmt"
)

// Визначаємо структуру для звірів
type Animal struct {
	Name string
	Cage int
}

// Визначаємо структуру для зоопарка
type Zoo struct {
	Animals []Animal
}

// Функція для повернення звірів в їх клітки
func (z *Zoo) ReturnToCage() {
	for _, animal := range z.Animals {
		fmt.Printf("%s повернення в клітку %d\n", animal.Name, animal.Cage)
	}
}

func main() {
	// Створюю звірів
	zoo := Zoo{
		Animals: []Animal{
			{"Лев", 1},
			{"Тигр", 2},
			{"Мавпа", 3},
			{"Змія", 4},
			{"Слон", 5},
		},
	}

	// Ситуация, коли звірі втекли з кліток
	fmt.Println("Звірі утекли з зоопарку!")

	// Функція для повернення звірів в клітки
	zoo.ReturnToCage()
}
