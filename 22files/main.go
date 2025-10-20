package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	// f,err := os.Open("example.txt")
	// if err != nil {
	// 	// fmt.Println("oopsie theres some error")
	// 	panic(err)
	// }else{

	// 	//getting the info of the file
	// 	fileInfo, err := f.Stat()
	// 	if err != nil{
	// 		panic(err)
	// 	}else{
	// 		fmt.Println("file name : ", fileInfo.Name())
	// 		fmt.Println("file size : ", fileInfo.Size())
	// 		fmt.Println("file modified at : ", fileInfo.ModTime())
	// 	}
	// }


	// READ FILE
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// buf := make([]byte,5)

	// d,err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// } 

	// for i:=0; i< len(buf); i++ {
	// 	fmt.Println("data",d, string(buf[i]))
	// }

	//working with folders (reading folders)
	// dir,err := os.Open("../")
	// if err != nil{
	// 	panic(err)
	// }

	// defer dir.Close()
	
	// fileInfo,err := dir.ReadDir(5)

	// for _,fi := range fileInfo{
	// 	fmt.Println(fi.Name())
	// }

	// create a file
	f,err := os.Create("example2.txt")
	if err != nil{
		panic(err)
	}

	defer f.Close()
	fmt.Println("-- Writing Files --")
	f.WriteString("Hi Muawiyah\n, this file has been created using Golang file handling...")
	f.WriteString("lets see if this appends or overwrites the text file")


	// read and write to another file (streaming fashion)
	sourceFile,err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	destinationFile,err := os.Create("example2.txt")
	if err != nil{
		panic(err)
	}

	defer destinationFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destinationFile)

	for {
		b,err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}

			break
		}

		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}
	}

	writer.Flush()

	fmt.Println("written to new file successfully")

}