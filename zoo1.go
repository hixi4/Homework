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

// Функція для створення зоопарку зі звірами
func createZoo() Zoo {
	return Zoo{
		Animals: []Animal{
			{"Лев", 1},
			{"Тигр", 2},
			{"Мавпа", 3},
			{"Змія", 4},
			{"Слон", 5},
		},
	}
}

// Функція для створення кліток
func createCages() []*Cage {
	return []*Cage{
		{ID: 1},
		{ID: 2},
		{ID: 3},
		{ID: 4},
		{ID: 5},
	}
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

// Функція для виведення інформації про клітки і їхніх мешканців
func printCagesInfo(cages []*Cage) {
	for _, cage := range cages {
		fmt.Printf("Клітка %d містить звірів: ", cage.ID)
		for _, animal := range cage.Animals {
			fmt.Printf("%s ", animal.Name)
		}
		fmt.Println()
	}
}

func main() {
	// Створюємо зоопарк
	zoo := createZoo()

	// Створюємо клітки
	cages := createCages()

	// Ситуація, коли звірі втекли з кліток
	fmt.Println("Звірі утекли з зоопарку!")

	// Повернення звірів в клітки
	zoo.ReturnToCage(cages)

	// Виводимо інформацію про клітки і їхніх мешканців
	printCagesInfo(cages)
}
