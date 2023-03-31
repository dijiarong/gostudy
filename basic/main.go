package main

import (
	"container/heap"
	"context"
	"fmt"
	datastruct "gostudy/basic/data_struct"

	"github.com/sashabaranov/go-openai"
)

func main() {
	// conf := gogpt.DefaultConfig("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")
	// // proxyAddress, _ := url.Parse("http://127.0.0.1:8081")
	// // conf.HTTPClient = &http.Client{
	// // 	Transport: &http.Transport{
	// // 		Proxy: http.ProxyURL(proxyAddress),
	// // 	},
	// // }

	// gptClient := gogpt.NewClientWithConfig(conf)

	// req := gogpt.ChatCompletionRequest{
	// 	Model:     gogpt.GPT3Dot5Turbo0301,
	// 	MaxTokens: 100,
	// 	Stream:    false,
	// 	Messages: []gogpt.ChatCompletionMessage{
	// 		{
	// 			Role:    "assistant",
	// 			Content: "用golang写一个堆排序",
	// 		},
	// 	},
	// }
	// resp, err := gptClient.CreateChatCompletion(context.Background(), req)
	// if err != nil {
	// 	return
	// }
	// fmt.Printf("%+v", resp)
	client := openai.NewClient("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")
	resp, err := client.CreateChatCompletion(
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

	fmt.Println(resp.Choices[0].Message.Content)
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
