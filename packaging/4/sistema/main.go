package main

import (
	"github.com/bruno-holanda15/go_expert_fc/packaging/4/math"
	"github.com/google/uuid"
)

func main() {
	m := math.Math{A: 1, B: 2}
	println(m.Add())
	println(uuid.New().String())
}