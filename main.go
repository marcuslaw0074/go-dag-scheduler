package main

import (
	"fmt"
	// "go-dag-scheduler/dag"
	"go-dag-scheduler/goairflow"
	// "reflect"
	// "time"
	// "github.com/robfig/cron/v3"
)

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
	fmt.Print("finished")
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