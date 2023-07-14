package main

import (
	"fmt"
	"sync"
)

func sum(start, end int, wg *sync.WaitGroup, ch chan int) {

	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	ch <- sum
	defer wg.Done()
}

func get(index int) (ret int) {
	// go 的异常处理
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Some error happened!", r)
			ret = -1
		}
	}()
	arr := [3]int{2, 3, 4}
	return arr[index]
}

type student struct {
	name    string
	age     int
	address string
}

func (s student) getInfo() string {
	return fmt.Sprintf("my name is %s, i am %d years old, and i live in %s", s.name, s.age, s.address)
}

type Person interface {
	getName() string
}

type Student struct {
	name string
	age  int
}

/*
*
方法可以具有接收者（receiver），它指定了方法是属于哪个类型的。
接收者可以是值类型或指针类型。这使得方法可以对接收者进行修改，而函数则不能。
下面的getName接受者是指针类型 下面强转接口（接受者）就需要对应的指针类型 所以是&
前后对应就好
*/
func (stu *Student) getName() string {
	return stu.name
}

type Worker struct {
	name   string
	gender string
}

func (w *Worker) getName() string {
	return w.name
}

//func main() {
//	var p Person = &Student{
//		name: "Tom",
//		age:  18,
//	}
//	fmt.Println(p.getName()) // Tom
//	add := calc.Add(1, 2)
//	fmt.Println(add)
//}

//func main() {
//	start := time.Now()
//	numWorkers := 100
//	total := 100000000
//
//	ch := make(chan int)
//	wg := sync.WaitGroup{}
//
//	for i := 0; i < numWorkers; i++ {
//		wg.Add(1)
//		go sum(i*(total/numWorkers)+1, (i+1)*(total/numWorkers), &wg, ch)
//	}
//
//	go func() {
//		wg.Wait()
//		close(ch)
//	}()
//
//	result := 0
//	for sum := range ch {
//		result += sum
//	}
//	duration := time.Since(start)
//	fmt.Println("Sum:", result, duration)
//
//	second := time.Now()
//	tmp := 0
//	for i := 0; i <= total; i++ {
//		tmp += i
//	}
//	secondEnd := time.Since(second)
//	fmt.Println("串行花费时间", tmp, secondEnd)
//
//}
