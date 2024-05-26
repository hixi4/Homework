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

// Визначаємо структуру кліток
type Cage struct {
	ID      int
	Animals []Animal
}

// Функція для повернення звірів в їх клітки
func (z *Zoo) ReturnToCage(cages []*Cage) {
	for _, animal := range z.Animals {
		for _, cage := range cages {
			if cage.ID == animal.Cage {
				cage.Animals = append(cage.Animals, animal)
				fmt.Printf("%s повернення в клітку %d\n", animal.Name, animal.Cage)
			}
		}
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
	// Створюю клітки
	cages := []*Cage{
		{ID: 1},
		{ID: 2},
		{ID: 3},
		{ID: 4},
		{ID: 5},
	}

	// Ситуація, коли звірі втекли з кліток
	fmt.Println("Звірі утекли з зоопарку!")

	// Функція для повернення звірів в клітки
	zoo.ReturnToCage(cages)

	// Виводимо інформацію про клітки і їхніх мешканців
	for _, cage := range cages {
		fmt.Printf("Клітка %d містить звірів: ", cage.ID)
		for _, animal := range cage.Animals {
			fmt.Printf("%s ", animal.Name)
		}
		fmt.Println()
	}
}
