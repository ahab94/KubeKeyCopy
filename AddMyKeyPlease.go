package main

import (
	"os"
	"flag"
)

func main() {
	publicKey := flag.String("key", "", "Public key to be copied (required)")
	username := flag.String("username", "root", "Username (default: root")
	flag.Parse()

	if *publicKey == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	sshFile := "/root/.ssh/authorized_keys"

	if *username != "root" {
		sshFile = "/home/" + *username + "/.ssh/authorized_keys"
	}

	f, err := os.OpenFile(sshFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(*publicKey); err != nil {
		panic(err)
	}
}
