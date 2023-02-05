package main

import "github.com/lupprl/goshortener/internal/app/server"

var UrlInputForm = `
<html>
    <head>
    <title></title>
    </head>
    <body>
        <form method="post">
            <label>Enter URL to shorten: </label><input type="text" name="url">
            <input type="submit" value="OK">
        </form>
    </body>
</html>
`

func main() {
	server.Run()
}
