package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"text/template"
)

func main() {
	idA, idB, idC := make([]byte, 1), make([]byte, 2), make([]byte, 2)
	_, _ = rand.Read(idA)
	_, _ = rand.Read(idB)
	_, _ = rand.Read(idC)

	address := fmt.Sprintf("fd%x:%x:%x", idA, idB, idC)

	tpl := `Address (Prefix):
{{.}}::/48

First subnet:
{{.}}::/64

Last subnet:
{{.}}:ffff::/64
`
	_ = template.Must(template.New("Address").Parse(tpl)).Execute(os.Stdout, address)
}
