package main

import (
	"fmt"
	// "go-dag-scheduler/dag"
	"go-dag-scheduler/goairflow"
	// "reflect"
	"time"
	// "github.com/robfig/cron/v3"
)

type BaseFunction struct {}

func (f BaseFunction) Test() (string, error){
	fmt.Println("hk")
	time.Sleep(time.Second * 3)
	return "ok", nil
}

func (f BaseFunction) Test2() (string, error){
	fmt.Println("hk2")
	time.Sleep(time.Second * 6)
	return "ok", nil
}

var k = goairflow.BaseFunction{}


func complexFunction() *goairflow.Job {
	j := &goairflow.Job{
		Name: "test",
		Schedule: "* * * * *",
	}

	j.Add(&goairflow.Task{
			BaseFunction: k,
			FunctionName: "Test",
			Name: "haha",
	})

	j.Add(&goairflow.Task{
		BaseFunction: k,
		FunctionName: "Test2",
		Name: "haha2",
	})

	j.SetDownstream(j.Task("haha"), j.Task("haha2"))
	return j
	
}

func main() {

	// s := complexFunction()
	// fmt.Println(s)
	goairflow.Complexunction()

	// fmt.Println(dag.H)
	// ls := make([]cron.EntryID, 0)
	// fmt.Println(ls)
	// c := cron.New()
	// c.Start()
	// res := reflect.ValueOf(k).MethodByName("Test")
	// ress := res.Call(nil)
	// resss := ress[0]
	// f := ress[1]
	// fmt.Println(resss.String())
	// if f.IsNil() {
	// 	fmt.Print(f.String())
	// }
	time.Sleep(time.Second * 10)
	// for {
	// 	count := 0
	// 	entryID, _ := c.AddFunc("@every 1s", func() {
	// 		count = count + 1
	// 		fmt.Println(count)
	// 		fmt.Println("tick every 1 second")
	// 	})
	// 	ls = append(ls, entryID)
	// 	fmt.Printf("entryID: %v \n", entryID)
	// 	fmt.Printf("list: %v", ls)
		
	// 	time.Sleep(time.Second * 3)
	// 	go func(en cron.EntryID) {
	// 		time.Sleep(time.Second * 2)
	// 		c.Remove(en)
	// 		fmt.Printf("remove entryID %v\n", en)
	// 	}(entryID)

	// }
	
}