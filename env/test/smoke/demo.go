package main

import "net/http"

func main() {
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
