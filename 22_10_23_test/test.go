package main

import (
	"fmt"
	"os"
)

func main() {
	dirName := "/Users/man9o/Desktop/Code/Go/22_10_23_test/aust"
	fileName := "/Users/man9o/Desktop/Code/Go/22_10_23_test/aust/2021304892.txt"

	// 创建目录
	err := os.Mkdir(dirName, os.ModePerm)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println("%s 目录创建成功！", dirName)
	}

	// 创建文件
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Println("%s 文件创建成功！%v", fileName, file)
	}
}

//------------------------------
//package main
//
//import (
//	"fmt"
//	"io"
//	"os"
//)
//
//func main() {
//	//dirName := "/Users/man9o/Desktop/Code/Go/22_10_23_test/aust"
//	fileName := "/Users/man9o/Desktop/Code/Go/22_10_23_test/aust/2021304892.txt"
//
//	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
//	if err != nil {
//		fmt.Println("打开文件失败！", err.Error())
//	} else {
//		n, err := file.Write([]byte("姓名：肖越\n"))
//		if err != nil {
//			fmt.Println("写入文件失败！", err.Error())
//		} else {
//			fmt.Println("写入成功！", n)
//		}
//		n, err = file.WriteString("班级：区块链21-1")
//		if err != nil {
//			fmt.Println("写入文件失败！", err.Error())
//		} else {
//			fmt.Println("写入成功！", n)
//		}
//	}
//
//	file1, err := os.Open(fileName)
//	if err != nil {
//		fmt.Println("打开文件失败！", err.Error())
//	} else {
//		bs := make([]byte, 1024*8, 1024*8)
//		n := 1
//		for {
//			n, err = file1.Read(bs)
//			if n == 0 || err == io.EOF {
//				fmt.Println("读取文件结束！")
//				break
//			}
//			fmt.Println(string(bs[:n]))
//		}
//	}
//	file1.Close()
//}
//
//// ------------------------------
//package main
//
//import (
//	"fmt"
//	"io"
//	"os"
//)
//
//func copyFile(srcFile, destFile string) (int64, error) {
//	file1, err := os.Open(srcFile)
//	if err != nil {
//		return 0, err
//	}
//	file2, err := os.OpenFile(destFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
//	if err != nil {
//		return 0, err
//	}
//	defer file1.Close()
//	defer file2.Close()
//
//	return io.Copy(file2, file1)
//}
//func main() {
//	// 源文件和路径
//	srcFile := "/Users/man9o/Desktop/Code/Go/22_10_23_test/aust/2021304892.txt"
//	// 目标文件的路径
//	destFile := "/Users/man9o/Desktop/Code/Go/22_10_23_test/aust/computer.txt"
//	total, err := copyFile(srcFile, destFile)
//	if err != nil {
//		fmt.Println(err.Error())
//	} else {
//		fmt.Println("复制成功：", total)
//	}
//}
