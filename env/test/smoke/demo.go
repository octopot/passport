package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
)

func main() {
	stoper := exec.Command("make", "stop-service")
	stoper.Stderr, stoper.Stdout = os.Stderr, os.Stdout
	_ = stoper.Run()
	fmt.Println("service down")

	dev := exec.Command("make", "dev-server")
	_ = dev.Start()
	fmt.Println("dev server start")

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		{
			_ = dev.Process.Kill()
			fmt.Println("dev server down")

			starter := exec.Command("make", "start-service")
			starter.Stderr, starter.Stdout = os.Stderr, os.Stdout
			_ = starter.Run()
			fmt.Println("service up")
		}
		signal.Stop(c)
		fmt.Println()
		os.Exit(0)
	}()

	_ = http.ListenAndServe(":9000", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		_, _ = rw.Write([]byte(`
<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Passport</title>
    <script src="//cdnjs.cloudflare.com/ajax/libs/fingerprintjs2/1.8.6/fingerprint2.min.js"></script>
</head>
<body>
    <h1>Demo Page</h1>
    <script src="//localhost:8080/api/v1/tracker/instruction"></script>
</body>
</html>
`))
	}))
}
