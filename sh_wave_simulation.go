package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
)

const N = 1000 // Grid size
const C = 0.2 // Wave speed
const M = 1000 // Time steps
const dt = 0.01
const dx = 0.01

func solveWaveEquation() [][][]float64 {
	// Initialize wave function at time t and t-1
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
				uNext[i][j] = 2*u[i][j] - uPrev[i][j] + C*C*dt*dt*(u[i-1][j]+u[i+1][j]+u[i][j-1]+u[i][j+1]-4*u[i][j])/(dx*dx)
			}
		}

		// Absorbing boundary condition
		for i := 0; i < N; i++ {
			uNext[i][0] = 0.5 * (u[i][0] + uPrev[i][0])
			uNext[i][N-1] = 0.5 * (u[i][N-1] + uPrev[i][N-1])
			uNext[0][i] = 0.5 * (u[0][i] + uPrev[0][i])
			uNext[N-1][i] = 0.5 * (u[N-1][i] + uPrev[N-1][i])
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
