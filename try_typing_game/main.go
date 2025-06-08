package main

import (
	"fmt"
)

// 標準出力でランダムに英単語を表示する
// ユーザーの入力を受け付ける
// 入力した単語と、表示した単語が一致していればOK

func inputWord() string {
	return "hello"
}

// func main() {
// 	typ := flag.String("type", "default", "type of word")
// 	flag.Parse()
	
// 	wordSets := map[string][]string{
// 		"default": {"hello", "world", "golang", "programming"},
// 		"animals": {"cat", "dog", "elephant", "tiger", "lion"},
// 		"fruits":  {"apple", "banana", "orange", "grape", "mango"},
// 	}
	
// 	words, exists := wordSets[*typ]
// 	if !exists {
// 		words = wordSets["default"]
// 	}
	
// 	rand.Seed(time.Now().UnixNano())
	
// 	userInput := inputWord()
// 	fmt.Println("ユーザーの入力:", userInput)
// 	// ランダムに選択
// 	randomWord := words[rand.Intn(len(words))]
// 	fmt.Println("ランダムな単語:", randomWord)

// }

func main() {
	totalScore := 0
	ask(1, "hello", &totalScore)
	ask(2, "world", &totalScore)
	ask(3, "golang", &totalScore)
	ask(4, "programming", &totalScore)
	ask(5, "typescript", &totalScore)
	ask(6, "javascript", &totalScore)
	ask(7, "python", &totalScore)
	ask(8, "java", &totalScore)
	ask(9, "c++", &totalScore)
	ask(10, "c#", &totalScore)
	
	fmt.Println("あなたの合計得点は", totalScore, "点です")
}

func ask(number int, question string, scorePtr *int) {
	var ans string
	fmt.Printf("[質問%d] 次の単語を入力してください: %s\n", number, question)
	fmt.Scanln(&ans)
	if ans == question {
		*scorePtr += 10
		fmt.Println("正解です")
	} else {
		fmt.Println("不正解です")
	}
}

