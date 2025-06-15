package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func shuffle(data []string) {
	n := len(data)
	rand.Seed(time.Now().Unix())
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

var t int

func init() {
	flag.IntVar(&t, "t", 1, "制限時間(分)")
	flag.Parse()
}

func main() {
	var (
		ch_rcv = input(os.Stdin)
		tm     = time.After(time.Duration(t) * time.Minute)
		words  = []string{"raccoon", "dog", "cat", "elephant", "tiger", "excellentswimmer is long language"}
		score  = 0
	)

	fmt.Println()
	shuffle(words)
	fmt.Println("タイピングゲームを始めます。制限時間は", t, "分。1語1点、", len(words), "点満点")
	num := 1
	for i := true; i && score < len(words); {
		question := words[score]
		fmt.Print("[質問", num, "]次の単語を入力してください:", question, "\n")
		fmt.Print("[答え]")
		select {
		case x := <-ch_rcv:
			if question == x {
				fmt.Println("正解です！")
				score++
				num++
			} else {
				fmt.Println("不正解です！")
			}
		case <-tm:
			fmt.Println("\n制限時間を過ぎました")
			i = false
		}

	}
	fmt.Println("あなたの点数:", score, "点 / ", len(words), " 点")
	n := score
	switch {
	case n <= 10:
		fmt.Println("判定 F")
	case n <= 20:
		fmt.Println("判定 D")
	case n <= 30:
		fmt.Println("判定 C")
	case n <= 35:
		fmt.Println("判定 B")
	case n <= 40:
		fmt.Println("判定 A")
	case n <= 45:
		fmt.Println("判定 S")
	case n > 45:
		fmt.Println("判定 SSS")
	default:
		fmt.Println("判定 F")
	}
}

func input(Stdin io.Reader) <-chan string {
	channel := make(chan string)
	go func() {
		strings := bufio.NewScanner(Stdin)
		for strings.Scan() {
			channel <- strings.Text()
		}
	}()
	return channel
}
