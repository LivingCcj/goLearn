package pipeline

import (
	"bufio"
	"net"
)

func NetWorkSink(addr string,in <- chan int){
	linsen,err:=net.Listen("tcp",addr)
	if err!=nil{
		panic(err)
	}

	go func() {
		defer linsen.Close()
		conn,err:=linsen.Accept()
		if err!=nil{
			panic(err)
		}
		defer conn.Close()

		writer:=bufio.NewWriter(conn)
		defer writer.Flush()
		WriteSink(writer,in)
	}()
}

func NetWorkSource(addr string) <- chan int{
	out:=make(chan int,1024)
	go func() {
		conn,err:=net.Dial("tcp",addr)
		if err!=nil{
			panic(err)
		}
		r:=ReaderBuff(bufio.NewReader(conn),-1)
		for v:=range r{
			out <- v
		}
		close(out)
	}()
	return out
}