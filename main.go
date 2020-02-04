package main

import (
	"fmt"
	"log"

	"golang.org/x/sys/windows/registry"
)

func main() {
	var c rune
	//open registry key
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\MyComputer\NameSpace`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	v, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Wow6432Node\Microsoft\Windows\CurrentVersion\Explorer\MyComputer\NameSpace`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
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
		getCase(c, k, v)
	}

	//close registry key
	if err := k.Close(); err != nil {
		log.Fatal(err)
	}
	if err := v.Close(); err != nil {
		log.Fatal(err)
	}
}

//get what folder to modify
func getCase(c rune, k registry.Key, v registry.Key) {

	switch c {
	case '1':
		changeValue("{088e3905-0323-4b02-9826-5d99428e115f}", "-{088e3905-0323-4b02-9826-5d99428e115f}", k)
		changeValue("{088e3905-0323-4b02-9826-5d99428e115f}", "-{088e3905-0323-4b02-9826-5d99428e115f}", v)
		fmt.Println("Download is disabled")
		break
	case '2':
		changeValue("{0DB7E03F-FC29-4DC6-9020-FF41B59E513A}", "-{0DB7E03F-FC29-4DC6-9020-FF41B59E513A}", k) //or {31C0DD25-9439-4F12-BF41-7FF4EDA38722}
		changeValue("{0DB7E03F-FC29-4DC6-9020-FF41B59E513A}", "-{0DB7E03F-FC29-4DC6-9020-FF41B59E513A}", v) //or {31C0DD25-9439-4F12-BF41-7FF4EDA38722}
		fmt.Println("3D Objects Disabled")
		break
	case '3':
		changeValue("{24ad3ad4-a569-4530-98e1-ab02f9417aa8}", "-{24ad3ad4-a569-4530-98e1-ab02f9417aa8}", k)
		changeValue("{24ad3ad4-a569-4530-98e1-ab02f9417aa8}", "-{24ad3ad4-a569-4530-98e1-ab02f9417aa8}", v)
		fmt.Println("Pictures is disabled")
		break
	case '4':
		changeValue("{B4BFCC3A-DB2C-424C-B029-7FE99A87C641}", "-{B4BFCC3A-DB2C-424C-B029-7FE99A87C641}", k)
		changeValue("{B4BFCC3A-DB2C-424C-B029-7FE99A87C641}", "-{B4BFCC3A-DB2C-424C-B029-7FE99A87C641}", v)
		fmt.Println("Desktop is disabled")
		break
	case '5':
		changeValue("{d3162b92-9365-467a-956b-92703aca08af}", "-{d3162b92-9365-467a-956b-92703aca08af}", k)
		changeValue("{d3162b92-9365-467a-956b-92703aca08af}", "-{d3162b92-9365-467a-956b-92703aca08af}", v)
		fmt.Println("Documents is disabled")
		break
	case '6':
		changeValue("{3dfdf296-dbec-4fb4-81d1-6a3438bcf4de}", "-{3dfdf296-dbec-4fb4-81d1-6a3438bcf4de}", k)
		changeValue("{3dfdf296-dbec-4fb4-81d1-6a3438bcf4de}", "-{3dfdf296-dbec-4fb4-81d1-6a3438bcf4de}", v)
		fmt.Println("Music is disabled")
		break
	case '7':
		changeValue("{f86fa3ab-70d2-4fc7-9c99-fcbf05467f3a}", "-{f86fa3ab-70d2-4fc7-9c99-fcbf05467f3a}", k)
		changeValue("{f86fa3ab-70d2-4fc7-9c99-fcbf05467f3a}", "-{f86fa3ab-70d2-4fc7-9c99-fcbf05467f3a}", v)
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

//changing the value of the key
func changeValue(name, value string, k registry.Key) {
	if err := k.SetStringValue(name, value); err != nil {
		log.Fatal(err)
	}
}
