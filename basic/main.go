package main

import (
	"container/heap"
	"context"
	"fmt"
	datastruct "gostudy/basic/data_struct"

	"github.com/sashabaranov/go-openai"
)

func main() {
	client := openai.NewClient("")
	resp, err := client.CreateChatCompletionStream(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}
	fmt.Println(resp)
}

func Test1() {
	type my struct {
		ID int
	}
	m := map[int]my{}
	m[1] = my{
		ID: 1,
	}
	a := m[1]
	a.ID = 2
	print(fmt.Sprintf("%+v", m))
}

func Test2() {
	myHeap := &datastruct.Myheap{8, 2, 10, 4, 5}
	heap.Init(myHeap)
	println(fmt.Sprintf("%+v", myHeap))
	heap.Push(myHeap, 0)
	println(fmt.Sprintf("%+v", myHeap))
	fmt.Printf("%+v", heap.Pop(myHeap))
	println(fmt.Sprintf("%+v", myHeap))
	fmt.Printf("%+v", heap.Pop(myHeap))
	println(fmt.Sprintf("%+v", myHeap))
}
