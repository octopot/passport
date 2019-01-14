package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
)

func main() {
	stopper := exec.Command("make", "stop-service")
	stopper.Stderr, stopper.Stdout = os.Stderr, os.Stdout
	if err := stopper.Run(); err != nil {
		fmt.Println("service down with the error", err)
	}

	dev := exec.Command("make", "dev-server")
	if err := dev.Start(); err != nil {
		fmt.Println("dev server starts with the error", err)
	}

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		{
			if err := dev.Process.Kill(); err != nil {
				fmt.Println("dev server down with the error", err)
			}

			starter := exec.Command("make", "start-service")
			starter.Stderr, starter.Stdout = os.Stderr, os.Stdout
			if err := starter.Run(); err != nil {
				fmt.Println("service up with the error", err)
			}
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
