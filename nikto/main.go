package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("docker", "run", "frapsoft/nikto", "-h", "http://10.45.253.151:9090/WebWolf/")

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	fmt.Printf("Combined Output :: \n%s\n", string(out))
}
