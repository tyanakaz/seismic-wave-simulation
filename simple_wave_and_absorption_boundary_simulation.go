package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
)

// Grid size
const N = 100

// The wave speed
const C = 0.2

// Time steps
const M = 1000

// The absorption factor
const Absorption = 0.5

func solveWaveEquation() [][][]float64 {
	// The wave function at time t and t-1
	u := make([][]float64, N)
	uPrev := make([][]float64, N)
	for i := range u {
		u[i] = make([]float64, N)
		uPrev[i] = make([]float64, N)
	}

	// Initial condition: a pulse in the middle
	for i := N/2 - 10; i < N/2+10; i++ {
		for j := N/2 - 10; j < N/2+10; j++ {
			u[i][j] = math.Exp(-float64((i-N/2)*(i-N/2)+(j-N/2)*(j-N/2)) / 100.0)
		}
	}

	// The wave function at the next time step
	uNext := make([][]float64, N)
	for i := range uNext {
		uNext[i] = make([]float64, N)
	}

	results := make([][][]float64, M)
	for i := range results {
		results[i] = make([][]float64, N)
		for j := range results[i] {
			results[i][j] = make([]float64, N)
		}
	}

	// Time-stepping loop
	for t := 0; t < M; t++ {
		// Spatial loop
		for i := 1; i < N-1; i++ {
			for j := 1; j < N-1; j++ {
				// Finite difference scheme
				uNext[i][j] = 2*u[i][j] - uPrev[i][j] + C*C*(u[i-1][j]+u[i+1][j]+u[i][j-1]+u[i][j+1]-4*u[i][j])
			}
		}

		// Apply absorbing boundary conditions
		for i := 0; i < N; i++ {
			uNext[0][i] = Absorption * (uNext[1][i] + uPrev[1][i]) / 2
			uNext[N-1][i] = Absorption * (uNext[N-2][i] + uPrev[N-2][i]) / 2
			uNext[i][0] = Absorption * (uNext[i][1] + uPrev[i][1]) / 2
			uNext[i][N-1] = Absorption * (uNext[i][N-2] + uPrev[i][N-2]) / 2
		}

		// Rotate arrays for the next time step
		for i := range u {
			uPrev[i], u[i], uNext[i] = u[i], uNext[i], uPrev[i]
		}

		// Copy the current state to results
		for i := range results[t] {
			copy(results[t][i], u[i])
		}
	}

	return results
}

func main() {
	// Run the simulation
	results := solveWaveEquation()

	for t := range results {
		// Open a CSV file for writing
		file, err := os.Create(fmt.Sprintf("csv/frame_%04d.csv", t))
		if err != nil {
			panic(err)
		}

		writer := csv.NewWriter(file)

		for i := range results[t] {
			row := make([]string, N)
			for j := range results[t][i] {
				row[j] = strconv.FormatFloat(results[t][i][j], 'f', -1, 64)
			}

			if err := writer.Write(row); err != nil {
				panic(err)
			}
		}

		writer.Flush()
		file.Close()
	}
}
