package pipeline

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)
var startTime  time.Time

func Init()  {
	startTime=time.Now()
}

func ArraySource(a ...int) chan int{
	out := make(chan int)
	go func(){
		for _,v := range a{
			out <- v
		}
		close(out)
	}()
	return out
}

func ImMemSort(arrIn <- chan int) <- chan int{
	out := make(chan int,1024)
	go func() {
		//read in
		a := []int{}
		for v := range arrIn{
			a=append(a, v)
		}
		fmt.Println("Read Done:",time.Now().Sub(startTime),len(a))
		//Sort
		sort.Ints(a)
		fmt.Println("Sort Done:",time.Now().Sub(startTime))
		//send out
		for _,v := range a{
			out <- v
		}
		close(out)
	}()
	return out
}

func Merge(in1,in2 <- chan int) <- chan int{
	out := make(chan int,1024)
	go func() {
		v1,ok1 := <- in1
		v2,ok2 := <- in2
		for ok1 || ok2 {
		if !ok2 || (ok1 && v1 <= v2){
			out <- v1
			v1,ok1 = <- in1
		}else{
			out <- v2
			v2,ok2 = <- in2
		}
		}
		close(out)
		fmt.Println("Merge Done:",time.Now().Sub(startTime))
	}()
	return out
}

func ReaderBuff(reader io.Reader,ChunkSize int) <- chan int{
	out := make (chan int,1024)
	go func() {
		buffer := make([]byte,8)
		bufferCount:=0
		for{
			n,err:=reader.Read(buffer)
			bufferCount+=n
			if n>0{
				v  := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			if err != nil || (ChunkSize !=-1 && bufferCount >= ChunkSize) {
				break
			}
		}
		close(out)
	}()
	return out
}

func WriteSink(writer io.Writer,in <- chan int){
		for v := range in{
			buffer:= make([]byte,8)
			binary.BigEndian.PutUint64(buffer,uint64(v))
			writer.Write(buffer)
		}
}

func RadomSource(count int) <- chan int{
	out := make(chan int )
	go func() {
		for i:=0;i<count;i++{
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

func MergeN(inputs ... <- chan int) <- chan int{
	if len(inputs) ==1 {
		return inputs[0]
	}
	m:=len(inputs)/2
	return Merge(MergeN(inputs[:m]...),MergeN(inputs[m:]...))
}