package main

import "fmt"

func main() {
	const name = "Nathnael Zelalem"
	html := `
        <!DOCTYPE html>
        <html lang="en">
        <head>
        <meta charset="UTF-8">
        <title>Hello World!</title>
        </head>
        <body>
        <h1>` + name + `</h1>
        </body>
         </html>`
	fmt.Println(html)
}
