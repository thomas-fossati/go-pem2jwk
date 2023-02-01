package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	pem, err := io.ReadAll(r)
	if err != nil {
		log.Fatalf("reading PEM file from stdin: %s", err)
	}

	k, err := jwk.ParseKey(pem, jwk.WithPEM(true))
	if err != nil {
		log.Fatalf("parsing PEM file into JWK: %s", err)
	}

	buf, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal key into JSON: %s", err)
	}

	fmt.Println(string(buf))
}
