package main

import (
	"os"
	"flag"
	"log"
	"encoding/base64"
	"io/ioutil"
	"strings"
)

func main() {
	publicKey := flag.String("key", "none", "B64 Public key to be copied (required)")
	username := flag.String("username", "root", "User to whom the key should be copied to")

	flag.Parse()

	if *publicKey == "none" {
		log.Println("Usage: ")
		flag.PrintDefaults()
		os.Exit(1)
	}

	sshFile := "/root/.ssh/authorized_keys"

	if *username != "root" {
		sshFile = "/home/" + *username + "/.ssh/authorized_keys"
	}

	if _, err := os.Stat(sshFile); os.IsNotExist(err) {
		log.Println("error: No such file or directory: ", sshFile)
		log.Println("Make sure the authorized_keys file exist")
		os.Exit(2)
	}

	f, err := os.OpenFile(sshFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	b64Key, err := base64.StdEncoding.DecodeString(*publicKey)
	if err != nil {
		panic(err)
	}

	read, err := ioutil.ReadFile(sshFile)

	if err != nil {
		panic(err)
	}

	if !strings.Contains(string(read), string(b64Key)) {
		if _, err = f.WriteString("\n" + string(b64Key)); err != nil {
			panic(err)
		}
	} else {
		log.Println("Key already exists")
		os.Exit(0)
	}

}

