package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	const name = "Nathnael Zelalem"
	str := fmt.Sprintf(`
<!DOCTYPE html>
        <html lang="en">
        <head>
        <meta charset="UTF-8">
        <title>Hello World!</title>
        </head>
        <body>
        <h1>` + name + `</h1>
        </body>
         </html>
`)
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
}
