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

func BogoSort(l []int32, tempo time.Duration) {
	for !Sorted(l) {
		time.Sleep(tempo)
		RandomizarLista(l)
	}
}

func BubbleSort(l []int32, tempo time.Duration) {
	tam := len(l)
	for i := range tam {
		for j := range tam - i - 1 {
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
	tam := len(l)
	meio := tam / 2
	MergeSort(l[:meio], tempo)
	MergeSort(l[meio:], tempo)
	aux := make([]int32, tam)
	ie, id, i := 0, meio, 0
	for ie < meio || id < tam {
		time.Sleep(tempo)
		if ie < meio && (id >= tam || l[ie] <= l[id]) {
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

func QuickSort(l []int32, tempo time.Duration) {
	tam := len(l)
	if tam < 2 {
		return
	}
	pivo := tam - 1
	ie, id := 0, pivo-1
	for ie <= id {
		for ie <= id && l[ie] < l[pivo] {
			ie++
			time.Sleep(tempo)
		}
		for ie <= id && l[id] > l[pivo] {
			id--
			time.Sleep(tempo)
		}
		if ie < id {
			Swap(&l[ie], &l[id])
		}
	}
	Swap(&l[ie], &l[pivo])
	QuickSort(l[:ie], tempo)
	QuickSort(l[ie+1:], tempo)
}

func main() {

	lista := CriarLista(10000)
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
	opcao := 1
	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(rl.KeyOne) {
			opcao = 1
		} else if rl.IsKeyPressed(rl.KeyTwo) {
			opcao = 2
		} else if rl.IsKeyPressed(rl.KeyThree) {
			opcao = 3
		} else if rl.IsKeyPressed(rl.KeyFour) {
			opcao = 4
		} else if rl.IsKeyPressed(rl.KeyFive) {
			opcao = 5
		}

		if rl.IsKeyPressed(rl.KeySpace) {
			RandomizarLista(lista)
			go func() {
				time.Sleep(time.Millisecond * 100)
				fmt.Println("comecou sort")
				switch opcao {
				case 1:
					BubbleSort(lista, time.Nanosecond*1000000)
				case 2:
					GnomeSort(lista, time.Nanosecond*1000000)
				case 3:
					MergeSort(lista, time.Nanosecond*1000000)
				case 4:
					QuickSort(lista, time.Nanosecond*500)
				case 5:
					BogoSort(lista, time.Nanosecond*5000000)
				}
				fmt.Println("terminou sort")
			}()
		}

		rl.BeginDrawing()

		switch opcao {
		case 1:
			rl.DrawText("bubble", 10, 30, 20, rl.White)
		case 2:
			rl.DrawText("gnome", 10, 30, 20, rl.White)
		case 3:
			rl.DrawText("merge", 10, 30, 20, rl.White)
		case 4:
			rl.DrawText("quick", 10, 30, 20, rl.White)
		case 5:
			rl.DrawText("bogo", 10, 30, 20, rl.White)
		}

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
