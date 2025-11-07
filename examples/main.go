package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	service := NewRickAndMortyService()
	totalPages := 5

	ctx := context.Background()

	// Canal para errores
	errors := make(chan error, totalPages)

	// Mapa para guardar resultados por página
	var resultsMap sync.Map

	var wg sync.WaitGroup

	for page := 1; page <= totalPages; page++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			chars, err := service.GetCharacters(ctx, p)
			if err != nil {
				errors <- fmt.Errorf("página %d: %v", p, err)
				return
			}
			// Guardar resultados en el mapa de forma concurrente
			resultsMap.Store(p, chars)
		}(page)
	}

	// Esperar a que todas las goroutines terminen
	wg.Wait()
	close(errors)

	// Imprimir errores si los hubo
	for err := range errors {
		fmt.Println("Error:", err)
	}

	// Imprimir resultados en orden de página
	resultsMap.Range(func(key, value any) bool {
		page := key.(int)
		chars := value.([]Character)

		fmt.Printf("Página %d:\n", page)
		for _, c := range chars {
			fmt.Printf("  [%s] %s (%s)\n", c.ID, c.Name, c.Status)
		}
		fmt.Println("---------------")

		return true // seguir iterando
	})
}
