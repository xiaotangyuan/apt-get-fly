package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

const APT_SOURCE_FILE = "/etc/apt/sources.list"
const ALTERNATIVE_SOURCE_PATH = "./sources"
const ORINGIN_SOURCE_FILE = "/etc/apt/sources.list.bak.by_apt-get-fly"

func createBakSourceFile() {
	_, err := os.Stat(ORINGIN_SOURCE_FILE)
	if err == nil {
		//log.Println("bak file already exist")
		return
	}
	cmd := exec.Command("cp", APT_SOURCE_FILE, ORINGIN_SOURCE_FILE)
	buf, err := cmd.Output()
	if err != nil {
		log.Fatal("make bak file error")
		log.Fatal(err)
	}
	log.Println(buf)
}

func useSourceFile(sourcename string) {
	sourcefilepath := fmt.Sprintf("%s/%s.sources.list", ALTERNATIVE_SOURCE_PATH, sourcename)
	log.Println("copy file ok:", sourcefilepath)
	cmd := exec.Command("cp", sourcefilepath, APT_SOURCE_FILE)
	_, err := cmd.Output()
	if err != nil {
		log.Fatal("useSourceFile file error")
		log.Fatal(err)
	}
	log.Println("updated sources.list for:", sourcename)
}

func main() {
	createBakSourceFile()
	useSourceFile("aliyun")
}
