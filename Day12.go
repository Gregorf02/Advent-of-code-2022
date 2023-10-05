package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
)

type Point struct {
	X, Y      int
	Elevation int
	Steps     int
}

func main() {
	// Apri il file di input
	file, err := os.Open("day12.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Leggi il file riga per riga
	var heightmap []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		heightmap = append(heightmap, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Trova la posizione di partenza (S) e la destinazione (E)
	var start, destination Point
	for y, row := range heightmap {
		for x, char := range row {
			if char == 'S' {
				start = Point{X: x, Y: y, Elevation: 0, Steps: 0}
			} else if char == 'E' {
				destination = Point{X: x, Y: y, Elevation: 25, Steps: 0} // Supponendo che 'z' abbia un valore ASCII di 122
			}
		}
	}

	// Applica la ricerca in ampiezza (BFS) per trovare il minor numero di passi
	steps := bfs(heightmap, start, destination)
	fmt.Println("Numero minimo di passi richiesti:", steps)
}

func bfs(heightmap []string, start, destination Point) int {
  neighbors := getValidNeighbors(heightmap, current, destination)
	visited := make(map[Point]bool)
	queue := list.New()

	queue.PushBack(start)
	visited[start] = true

	for queue.Len() > 0 {
		current := queue.Remove(queue.Front()).(Point)

		neighbors := getValidNeighbors(heightmap, current)
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				queue.PushBack(neighbor)
				visited[neighbor] = true

				if neighbor.X == destination.X && neighbor.Y == destination.Y {
					return neighbor.Steps
				}
			}
		}
	}

	return -1 // Se la destinazione non è raggiungibile
}

func getValidNeighbors(heightmap []string, p, destination Point) []Point {
	directions := []Point{
		{X: p.X - 1, Y: p.Y, Elevation: 0, Steps: 0}, // Sinistra
		{X: p.X + 1, Y: p.Y, Elevation: 0, Steps: 0}, // Destra
		{X: p.X, Y: p.Y - 1, Elevation: 0, Steps: 0}, // Su
		{X: p.X, Y: p.Y + 1, Elevation: 0, Steps: 0}, // Giù
	}

	var neighbors []Point
	for _, dir := range directions {
		if isValidMove(heightmap, p, dir, destination) {
			dir.Elevation = int(heightmap[dir.Y][dir.X]) - 97
			dir.Steps = p.Steps + 1
			neighbors = append(neighbors, dir)
		}
	}

	return neighbors
}

func isValidMove(heightmap []string, p, q Point) bool {
	if q.X >= 0 && q.X < len(heightmap[0]) && q.Y >= 0 && q.Y < len(heightmap) {
		currentElevation := int(heightmap[p.Y][p.X]) - 97
		destinationElevation := int(heightmap[q.Y][q.X]) - 97
		return destinationElevation <= currentElevation+1
	}
	return false
}
