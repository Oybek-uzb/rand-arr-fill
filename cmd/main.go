package cmd

import (
	"log"
	"rand-arr-fill/internal"
)

const (
	rowsCount    = 5
	columnsCount = 5
	randomLimit  = 100
)

func main() {
	randomizer := internal.NewRandomizer()
	randomIntegers, err := randomizer.GetUniqueRandomIntegers(internal.GetUniqueRandomIntegersInput{
		SizeX:       rowsCount,
		SizeY:       columnsCount,
		RandomLimit: randomLimit,
	})
	if err != nil {
		log.Fatal(err)
	}

	array, err := internal.NewIntArray(rowsCount, columnsCount)
	if err != nil {
		log.Fatal(err)
	}

	if err := array.Fill(randomIntegers); err != nil {
		log.Fatal(err)
	}

	if err := internal.Print(array); err != nil {
		log.Fatal(err)
	}
}
