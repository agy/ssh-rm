package main

import (
	"errors"
	"flag"
	"net"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

func Fingerprint(fp string) (func(ssh.PublicKey) string, error) {
	if fp == "" {
		return nil, errors.New("invalid fingerprint")
	}

	hash := strings.Split(fp, ":")[0]
	switch hash {
	case "SHA256":
		return ssh.FingerprintSHA256, nil
	case "MD5":
		return ssh.FingerprintLegacyMD5, nil
	}

	return nil, errors.New("invalid key fingerprint")
}

func main() {
	var (
		fingerprint = flag.String("fingerprint", "", "fingerprint")
	)
	flag.Parse()

	if *fingerprint == "" {
		flag.Usage()
		os.Exit(0)
	}

	conn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	sshAgent := agent.NewClient(conn)

	// Grab keys in the agent
	keys, err := sshAgent.List()
	if err != nil {
		panic(err)
	}

	if len(keys) == 0 {
		os.Exit(0)
	}

	// Combine this with the below logic in a separate function
	fpFunc, err := Fingerprint(*fingerprint)
	if err != nil {
		panic(err)
	}

	for _, key := range keys {
		pub, err := ssh.ParsePublicKey(key.Blob)
		if err != nil {
			panic(err)
		}

		switch pub.(type) {
		case *ssh.Certificate:
			sshCert, ok := pub.(*ssh.Certificate)
			if !ok {
				panic(errors.New("invalid SSH certificate"))
			}

			k, err := ssh.ParsePublicKey(sshCert.Key.Marshal())
			if err != nil {
				panic(err)
			}

			if fpFunc(k) != *fingerprint {
				continue
			}
		default:
			if fpFunc(pub) != *fingerprint {
				continue
			}
		}

		if err := sshAgent.Remove(pub); err != nil {
			panic(err)
		}

	}
}
