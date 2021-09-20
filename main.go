package main

import (
	"bufio"
	"fmt"
	"github.com/smallnest/rpcx/log"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/share"
	"io"
	"net"
	"unsafe"
)

func main() {
	s := server.NewServer()
	transfer := server.NewFileTransfer(":8081", saveFilehandler, nil, 100)
	s.EnableFileTransfer(share.SendFileServiceName, transfer)
	if err := s.Serve("tcp", ":8080"); err != nil {
		log.Fatalf("start server err:%#v", err)
	}
}

func saveFilehandler(conn net.Conn, args *share.FileTransferArgs) {
	fmt.Printf("received file name: %s, size: %d, meta: %v\n", args.FileName, args.FileSize, args.Meta)
	//data, err := ioutil.ReadAll(conn)
	//if err != nil {
	//	fmt.Printf("error read: %v\n", err)
	//	return
	//}
	tmp := bufio.NewReadWriter(bufio.NewReaderSize(conn, 1000), nil)
	//for _, v := range data {
	//	if v != 0 {
	//		tmp.WriteByte(v)
	//	}
	//}
	for {
		bytes, err := tmp.ReadBytes('\x00')
		if err == io.EOF {
			log.Errorf("read upload file err:%#v", err)
			return
		}
		if *(*string)(unsafe.Pointer(&bytes)) != string(byte('\x00')) {
			log.Infof("read data:%#v", string(bytes))
		}

	}

	//ioutil.WriteFile("./upload.txt", bytes, 0777)
	//fmt.Printf("file content: %#v\n", string(bytes))
}
