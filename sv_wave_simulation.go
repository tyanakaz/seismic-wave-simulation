package main

import (
	"fmt"
	"math"
)

// Grid size
const N = 100

// Time steps
const M = 1000

// The wave speed
const C = 0.2

// Absorbing layer thickness
const AbsorbingLayer = 10

// Absorption coefficient
const Absorption = 0.1

func updateBoundary(u [][]float64) {
	for i := 0; i < AbsorbingLayer; i++ {
		factor := math.Exp(-Absorption * float64(AbsorbingLayer-i))
		for j := 0; j < N; j++ {
			u[i][j] *= factor
			u[N-1-i][j] *= factor
			u[j][i] *= factor
			u[j][N-1-i] *= factor
		}
	}
}

func solveWaveEquation() [][]float64 {
	u := make([][]float64, N)
	v := make([][]float64, N)
	for i := range u {
		u[i] = make([]float64, N)
		v[i] = make([]float64, N)
	}

	for i := N/2 - 10; i < N/2+10; i++ {
		for j := N/2 - 10; j < N/2+10; j++ {
			u[i][j] = math.Exp(-float64((i-N/2)*(i-N/2)+(j-N/2)*(j-N/2)) / 100.0)
		}
	}

	uNext := make([][]float64, N)
	for i := range uNext {
		uNext[i] = make([]float64, N)
	}

	for t := 0; t < M; t++ {
		for i := 1; i < N-1; i++ {
			for j := 1; j < N-1; j++ {
				uNext[i][j] = 2*u[i][j] - u[i][j] + C*C*(u[i+1][j] + u[i-1][j] + u[i][j+1] + u[i][j-1] - 4*u[i][j])
			}
		}
		updateBoundary(uNext)
		u, uNext = uNext, u
	}

	return u
}

func main() {
	results := solveWaveEquation()
	for _, row := range results {
		for _, value := range row {
			fmt.Printf("%f ", value)
		}
		fmt.Println()
	}
}
