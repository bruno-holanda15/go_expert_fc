package main

type MyNumber int

// constraint
type Number interface {
	~int | ~float64
}

// usando generics
func Soma[T Number] (lista map[string]T) T {
	var soma T

	for _,n := range lista {
		soma += n
	}

	return soma
}

func main() {
	m := map[string]int{"brunin":300,"pat":300,"apt":10}
	f := map[string]float64{"brunin":300.0,"pat":300.0,"apt":12.0}
	m2 := map[string]MyNumber{"brunin":300,"pat":300,"apt":10}

	println(Soma(m))
	println(Soma(f))
	println(Soma(m2))
}

