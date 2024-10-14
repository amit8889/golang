package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("=========Go files=========")

	//=================************** File reading **************=================
	// it load comple file in ram avoid for large file
	// data, err := os.ReadFile("tech.txt")
	// if err != nil {
	// 	fmt.Println("=error in file reading==>", err)
	// 	panic(err)
	// }
	// fmt.Println("=========File Content=========")
	// fmt.Println(string(data))

	//==================*********** open file meta data ***********========================
	// fileRef, err := os.Open("tech.txt")
	// if err != nil {
	// 	fmt.Println("Error in opening file", err)
	// 	panic(err)
	// }
	// defer fileRef.Close()
	// fmt.Println("=========File Content=========")
	// fileInfo, err := fileRef.Stat()
	// if err != nil {
	// 	fmt.Println("=====error in file info======>", err)
	// 	panic(err)
	// }
	// fmt.Println(fileInfo)
	// // name
	// fmt.Println("===name==>", fileInfo.Name())
	// // size
	// fmt.Println("===size===>", fileInfo.Size())
	// // mode
	// fmt.Println("=====mode====>", fileInfo.Mode())
	// // modTime  last modified time
	// fmt.Println("=====mode time===>", fileInfo.ModTime())
	// // isDir
	// fmt.Println("===isDir====>", fileInfo.IsDir())
	// // sys
	// fmt.Println("===sys====>", fileInfo.Sys())

	//==================*********** read file data ***********========================
	// fileRef, err := os.Open("tech.txt")
	// if err != nil {
	// 	fmt.Println("Error in opening file", err)
	// 	panic(err)
	// }
	// defer fileRef.Close()
	// fmt.Println("=========File Content=========")
	// fileInfo, err := fileRef.Stat()
	// if err != nil {
	// 	fmt.Println("=====error in file info======>", err)
	// 	panic(err)
	// }
	// buf := make([]byte, fileInfo.Size())
	// n, err := fileRef.Read(buf)
	// if err != nil {
	// 	fmt.Println("=====error in file reading======>", err)
	// 	panic(err)
	// }
	// // n is buffer lenth//
	// fmt.Println("=====bytes====>", n)

	// fmt.Println("=====read data====>", string(buf))

	//=========================*********** Read folder ************=======================
	// current folder
	// dir, err := os.Open("../")
	// if err != nil {
	// 	fmt.Println("=====error in opening folder======>", err)
	// 	panic(err)
	// }
	// defer dir.Close()
	// filesList, err := dir.ReadDir(0)
	// if err != nil {
	// 	fmt.Println("=====error in reading folder======>", err)
	// 	panic(err)
	// }
	// fmt.Println(filesList)
	// for _, file := range filesList {
	// 	fmt.Println(file, file.IsDir())
	// }

	//=========================*********** create file ************=======================

	// fileRef, err := os.Create("text.txt")
	// if err != nil {
	// 	fmt.Println("=====error in creating file======>", err)
	// 	panic(err)
	// }
	// defer fileRef.Close()
	// //=========================*********** write file ************=======================
	// //fileRef.WriteString("Hello, !")
	// //=========================*********** write file ************=======================
	// fileRef.Write([]byte("Hello, world!"))

	//====================*************Read & Write in stream fashion ********=================

	// Rfile, err := os.Open("main.go")
	// if err != nil {
	// 	fmt.Println("=====error in opening file======>", err)
	// 	panic(err)
	// }
	// defer Rfile.Close()
	// //writer file
	// Wfile, err := os.Create("copyMain.go")
	// if err != nil {
	// 	fmt.Println("=====error in opening file======>", err)
	// 	panic(err)
	// }
	// defer Wfile.Close()
	// // make reader and write
	// reader := bufio.NewReader(Rfile)
	// writer := bufio.NewWriter(Wfile)
	// // read and write
	// for {
	// 	// read
	// 	buf, err := reader.ReadByte()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println("=====error in reading file======>", err)
	// 		panic(err)
	// 	}
	// 	// write
	// 	writer.WriteByte(buf)
	// }
	// // flush
	// writer.Flush()
	// fmt.Println("========file write success fully, name : ", Wfile.Name())

	// ============********* delete a file ********================
	err := os.Remove("copyMain.go")
	if err != nil {
		fmt.Println("=====error in deleting file======>", err)
		panic(err)
	}
	fmt.Println("========file deleted successfully, name : ", "copyMain.go")

}
