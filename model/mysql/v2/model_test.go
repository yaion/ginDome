package v2

import (
	"fmt"
	"strconv"
	"sync"
	"testing"

)

type User struct {
	ID int64 `gorm:"primarkKey"`
	Name string
	Age int
}


func TestModel(t *testing.T){
	user := []User{
		User{Name: "jack1" ,Age: 32},
		User{Name: "jack2" ,Age: 32},
		User{Name: "jack3" ,Age: 32},
		User{Name: "jack4" ,Age: 32},
	}
	res := DB.Table("user").Create(user)
	t.Log(res)
	sync.Pool{}
}

func TestSelect(t *testing.T){

		// 创建通道
		intChan := make(chan int)
		stringChan := make(chan string)

		// 写入
		go func() {
			for i := 1; i < 4; i++ {
				intChan <- i
				stringChan <- "张xx-" + strconv.Itoa(i)
			}
		}()
		// 使用select 没有default,则会阻塞等待。（会随机运行一次）
	for i := 1; i < 4; i++ {
		select {
		case n := <-intChan:
			fmt.Printf("接收到intChan中的数据: %v\n", n)
		case s := <-stringChan:
			fmt.Printf("接收到stringChan中的数据: %v\n", s)
		}
	}
		fmt.Printf("运行结束!")

}
