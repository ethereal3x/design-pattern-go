package functionaloption

import (
	"fmt"
	"testing"
	"time"
)

func TestFunctionalOption(t *testing.T) {
	client := NewHttpClient(WithTimeout(5 * time.Second))
	// 使用自定义HTTP客户端发起GET请求
	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.Status)
}
