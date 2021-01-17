package main

import (
	"bufio"
	"fmt"
	"os"
	"pipeline"
	"strconv"
)

func main() {
	p := createNetworkPipe("large.txt", 80000000, 4)
	//time.Sleep(time.Hour)
	//p:=createPipe("large.txt",80000000,8)
	WriteResult("large_out.txt", p)
	printResult("large_out.txt")
}

func printResult(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderBuff(bufio.NewReader(file), -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count > 10 {
			break
		}
	}
	fmt.Println(count)

}

func WriteResult(fileName string, in <-chan int) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	pipeline.WriteSink(writer, in)
	defer writer.Flush()

}

func createPipe(fileName string, fileSize int, chunkSize int) <-chan int {
	getAllData := []<-chan int{}
	offset := fileSize / chunkSize
	pipeline.Init()
	for i := 0; i < chunkSize; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*offset), 0)
		source := pipeline.ReaderBuff(bufio.NewReader(file), offset)
		p := pipeline.ImMemSort(source)
		getAllData = append(getAllData, p)
	}
	return pipeline.MergeN(getAllData...)
}

func createNetworkPipe(fileName string, fileSize int, chunkSize int) <-chan int {
	var getAllData []<-chan int
	offset := fileSize / chunkSize
	var addrString []string
	pipeline.Init()
	for i := 0; i < chunkSize; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*offset), 0)
		addr := ":" + strconv.Itoa(7000+i)
		source := pipeline.ReaderBuff(bufio.NewReader(file), offset)
		pipeline.NetWorkSink(addr, pipeline.ImMemSort(source))
		addrString = append(addrString, addr)
	}

	for i := 0; i < chunkSize; i++ {
		p := pipeline.NetWorkSource(addrString[i])
		getAllData = append(getAllData, p)
	}

	return pipeline.MergeN(getAllData...)
}
