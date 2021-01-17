package main

import (
	"bufio"
	"fmt"
	"os"
	"pipeline"
)

func main() {
	//p := pipeline.Merge(pipeline.ImMemSort(pipeline.ArraySource(1,23,20,123,9)),
	//			pipeline.ImMemSort(pipeline.ArraySource(3,41,40,54,1,31,30)))
	//printP1(p)

	const filename = "large.txt"
	const count = 10000

	file,err:= os.Create(filename)
	if err !=nil{
		panic(err)
	}
	defer file.Close()

	writer:=bufio.NewWriter(file)
	pipeline.WriteSink(writer,pipeline.RadomSource(count))
	writer.Flush()

	fileReader,err:=os.Open(filename)
	if err!=nil{
		panic(err)
	}
	defer fileReader.Close()
	p:=pipeline.ReaderBuff(bufio.NewReader(fileReader),-1)
	i   :=0
	for v:= range p{
		fmt.Println(v)
		i++
		if i > 10 {
			break;
		}
	}
}


func printP1(p <- chan int){
	for v := range p{
		fmt.Println(v)
	}
}

func printP2(p <- chan int){
	for {
		if num ,ok := <-p; ok{
			fmt.Println(num)
		}else{
			break
		}
	}
}