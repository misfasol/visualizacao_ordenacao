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

}

func Sorted(l []int32) bool {
	for i := range len(l) - 1 {
		if l[i] > l[i+1] {
			return false
		}
	}
	return true
}

func Bubble(l []int32) {
	time.Sleep(time.Second)
	for i := range len(l) {
		for j := range len(l) - i - 1 {
			if l[j] > l[j+1] {
				temp := l[j]
				l[j] = l[j+1]
				l[j+1] = temp
			}
			time.Sleep(time.Millisecond)
		}
	}
}

func main() {

	lista := CriarLista(100)

	fmt.Println(lista)
	go Bubble(lista)

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
