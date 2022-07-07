package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: showcert <filename>")
	}

	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	for {
		d, rest := pem.Decode(data)
		if d == nil {
			log.Fatal("pem decode failed")
		}

		c, err := x509.ParseCertificate(d.Bytes)
		if err != nil {
			log.Fatal("parse of cert failed")
		}

		fmt.Printf("Subject: %v\nIssuer : %v\n\n", c.Subject, c.Issuer)

		if len(rest) == 0 {
			break
		}

		data = rest
	}
}
