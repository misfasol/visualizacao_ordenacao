package main

import (
	"fmt"
	"math"
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
	return lista
}

func RandomizarLista(lista []int32) {
	for i := range lista {
		temp := lista[i]
		rand := rand.IntN(len(lista))
		lista[i] = lista[rand]
		lista[rand] = temp
	}
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

func BubbleSort(l []int32, tempo time.Duration) {
	for i := range len(l) {
		for j := range len(l) - i - 1 {
			if l[j] > l[j+1] {
				Swap(&l[j], &l[j+1])
			}
			time.Sleep(tempo)
		}
	}
}

func GnomeSort(l []int32, tempo time.Duration) {
	pos := 0
	tamanho := len(l)
	for pos < tamanho {
		time.Sleep(tempo)
		if pos == 0 || l[pos] >= l[pos-1] {
			pos++
		} else {
			Swap(&l[pos], &l[pos-1])
			pos--
		}
	}
}

func MergeSort(l []int32, tempo time.Duration) {
	if len(l) <= 1 {
		return
	}
	MergeSort(l[:len(l)/2], tempo)
	MergeSort(l[len(l)/2:], tempo)
	aux := make([]int32, len(l))
	ie, id, i := 0, len(l)/2, 0
	for ie < len(l)/2 || id < len(l) {
		time.Sleep(tempo)
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
		time.Sleep(tempo)
	}
}

func main() {

	lista := CriarLista(100)
	RandomizarLista(lista)

	fmt.Println(lista)

	// go func() {
	// 	time.Sleep(time.Second)
	// 	fmt.Println("comecou sort")
	// 	MergeSort(lista, time.Nanosecond)
	// 	fmt.Println("terminou sort")
	// }()

	rl.SetTargetFPS(60)
	rl.InitWindow(TAMANHO_LARGURA, TAMANHO_ALTURA, "asd")
	var posx, posy, wid, hei int32 = 0, 0, 0, 0
	contador := 0
	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(rl.KeyR) {
			RandomizarLista(lista)
			go func() {
				time.Sleep(time.Millisecond * 100)
				fmt.Println("comecou sort")
				MergeSort(lista, time.Nanosecond*10000000)
				fmt.Println("terminou sort")
			}()
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		for k, v := range lista {
			posx = int32(float32(k) * (float32(TAMANHO_LARGURA) / float32(len(lista))))
			posy = int32(float32(TAMANHO_ALTURA) - (float32(v) * (TAMANHO_ALTURA / float32(len(lista)))))
			wid = int32(math.Ceil(float64(TAMANHO_LARGURA) / float64(len(lista))))
			hei = int32(float32(v) * (float32(TAMANHO_ALTURA) / float32(len(lista))))
			contador++
			rl.DrawRectangle(posx, posy, wid, hei, rl.White)
		}
		contador = 0

		rl.DrawFPS(10, 10)
		rl.EndDrawing()

	}
	rl.CloseWindow()
}
