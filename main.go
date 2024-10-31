package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	TAMANHO_LARGURA = 1280
	TAMANHO_ALTURA  = 720
)

func CriarLista(n int) []int32 {
	lista := make([]int32, n)
	for i := range lista {
		lista[i] = int32(i) + 1
	}
	for i := range lista {
		temp := lista[i]
		rand := rand.IntN(n)
		lista[i] = lista[rand]
		lista[rand] = temp
	}
	return lista
}

func Swap(a, b *int32) {
	if *a != *b {
		(*a) ^= (*b)
		(*b) ^= (*a)
		(*a) ^= (*b)
	}
}

func Sorted(l []int32) bool {
	for i := range len(l) - 1 {
		if l[i] > l[i+1] {
			return false
		}
	}
	return true
}

func BubbleSort(l []int32) {
	for i := range len(l) {
		for j := range len(l) - i - 1 {
			if l[j] > l[j+1] {
				Swap(&l[j], &l[j+1])
			}
			time.Sleep(time.Millisecond)
		}
	}
}

func GnomeSort(l []int32) {
	pos := 0
	tamanho := len(l)
	for pos < tamanho {
		time.Sleep(time.Millisecond)
		if pos == 0 || l[pos] >= l[pos-1] {
			pos++
		} else {
			Swap(&l[pos], &l[pos-1])
			pos--
		}
	}
}

func MergeSort(l []int32) {
	if len(l) <= 1 {
		return
	}
	MergeSort(l[:len(l)/2])
	MergeSort(l[len(l)/2:])
	aux := make([]int32, len(l))
	ie, id, i := 0, len(l)/2, 0
	for ie < len(l)/2 || id < len(l) {
		time.Sleep(time.Millisecond * 10)
		if ie < len(l)/2 && (id >= len(l) || l[ie] <= l[id]) {
			aux[i] = l[ie]
			ie++
		} else {
			aux[i] = l[id]
			id++
		}
		i++
	}
	for i := range l {
		l[i] = aux[i]
		time.Sleep(time.Millisecond * 10)
	}
}

func main() {

	lista := CriarLista(100)

	fmt.Println(lista)

	go func() {
		time.Sleep(time.Second)
		fmt.Println("comecou sort")
		MergeSort(lista)
		fmt.Println("terminou sort")
	}()

	rl.SetTargetFPS(60)
	rl.InitWindow(TAMANHO_LARGURA, TAMANHO_ALTURA, "asd")
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		for k, v := range lista {
			rl.DrawRectangle(int32(k*(TAMANHO_LARGURA/len(lista))), TAMANHO_ALTURA-(v*int32((TAMANHO_ALTURA/len(lista)))), int32(TAMANHO_LARGURA/len(lista)), v*(int32(TAMANHO_ALTURA/len(lista))), rl.White)
		}

		rl.DrawFPS(10, 10)
		rl.EndDrawing()

	}
	rl.CloseWindow()
}
