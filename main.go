package main

import (
	"fmt"
	"log"

	"golang.org/x/sys/windows/registry"
)

func main() {
	var c rune

	fmt.Print("1) Downloads\n")
	fmt.Print("2) 3D Objects\n")
	fmt.Print("3) Pictures\n")
	fmt.Print("4) Desktop\n")
	fmt.Print("5) Documents\n")
	fmt.Print("6) Music\n")
	fmt.Print("7) Videos\n")
	for c != 'q' {
		fmt.Println("Type in number of folder you want to disable. To terminate press q")
		fmt.Scanf("%c\n", &c)
		getCase(c)
	}

	//close registry key
}

//get what folder to modify
func getCase(c rune) {

	switch c {
	case '1':
		exist := changeValue("{088e3905-0323-4b02-9826-5d99428e115f}")
		deleteValue("{088e3905-0323-4b02-9826-5d99428e115f}", exist)
		fmt.Println("Download is disabled")
		break
	case '2':
		exist := changeValue("{0DB7E03F-FC29-4DC6-9020-FF41B59E513A}") //or {31C0DD25-9439-4F12-BF41-7FF4EDA38722}
		deleteValue("{0DB7E03F-FC29-4DC6-9020-FF41B59E513A}", exist)
		fmt.Println("3D Objects Disabled")
		break
	case '3':
		exist := changeValue("{24ad3ad4-a569-4530-98e1-ab02f9417aa8}")
		deleteValue("{24ad3ad4-a569-4530-98e1-ab02f9417aa8}", exist)
		fmt.Println("Pictures is disabled")
		break
	case '4':
		exist := changeValue("{B4BFCC3A-DB2C-424C-B029-7FE99A87C641}")
		deleteValue("{B4BFCC3A-DB2C-424C-B029-7FE99A87C641}", exist)
		fmt.Println("Desktop is disabled")
		break
	case '5':
		exist := changeValue("{d3162b92-9365-467a-956b-92703aca08af}")
		deleteValue("{d3162b92-9365-467a-956b-92703aca08af}", exist)
		fmt.Println("Documents is disabled")
		break
	case '6':
		exist := changeValue("{3dfdf296-dbec-4fb4-81d1-6a3438bcf4de}")
		deleteValue("{3dfdf296-dbec-4fb4-81d1-6a3438bcf4de}", exist)
		fmt.Println("Music is disabled")
		break
	case '7':
		exist := changeValue("{f86fa3ab-70d2-4fc7-9c99-fcbf05467f3a}")
		deleteValue("{f86fa3ab-70d2-4fc7-9c99-fcbf05467f3a}", exist)
		fmt.Println("Videos is disabled")
		break
	case 'q':
		fmt.Println("Program is being terminated")
		break
	default:
		fmt.Println("Invalid input")
		break
	}
}

//change the key
func changeValue(value string) bool {
	k, exist, err := registry.CreateKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\MyComputer\NameSpace\-`+value, registry.ALL_ACCESS)
	check(err)
	v, exist, err := registry.CreateKey(registry.LOCAL_MACHINE, `SOFTWARE\Wow6432Node\Microsoft\Windows\CurrentVersion\Explorer\MyComputer\NameSpace\-`+value, registry.ALL_ACCESS)
	check(err)
	k.Close()
	v.Close()
	return exist
}

//delete the old key
func deleteValue(value string, exist bool) {
	k, kExist, e := registry.CreateKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\MyComputer\NameSpace\`+value, registry.ALL_ACCESS)
	check(e)
	v, vExist, e := registry.CreateKey(registry.LOCAL_MACHINE, `SOFTWARE\Wow6432Node\Microsoft\Windows\CurrentVersion\Explorer\MyComputer\NameSpace\`+value, registry.ALL_ACCESS)
	check(e)
	if !exist || !kExist {
		err := registry.DeleteKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\MyComputer\NameSpace\`+value)
		check(err)
	}
	if !exist || !vExist {
		err := registry.DeleteKey(registry.LOCAL_MACHINE, `SOFTWARE\Wow6432Node\Microsoft\Windows\CurrentVersion\Explorer\MyComputer\NameSpace\`+value)
		check(err)
	}
	k.Close()
	v.Close()
}

//check for error
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
