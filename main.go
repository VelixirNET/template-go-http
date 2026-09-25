package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const page = `<!doctype html>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width,initial-scale=1">
<title>Live on velixir</title>
<style>
  :root { color-scheme: dark; }
  body { margin:0; min-height:100vh; display:grid; place-items:center;
         background:#0b0b10; color:#e8e8ef;
         font:16px/1.6 ui-sans-serif,system-ui,-apple-system,Segoe UI,sans-serif; }
  main { max-width:34rem; padding:2.5rem 1.5rem; text-align:center; }
  h1 { font-size:1.6rem; margin:0 0 .5rem; letter-spacing:-.02em; }
  p { color:#a0a0b0; margin:.5rem 0; }
  code { background:#16161f; border:1px solid #26263a; border-radius:6px;
         padding:.15rem .4rem; font-size:.875em; color:#c7d2fe; }
  .dot { display:inline-block; width:.5rem; height:.5rem; border-radius:50%;
         background:#34d399; margin-right:.5rem; vertical-align:middle; }
</style>
<main>
  <p><span class="dot"></span>Go is running on velixir</p>
  <h1>Your first deploy worked.</h1>
  <p>This page is served by your own container, built from source. No Dockerfile was involved.</p>
  <p>Next: edit <code>main.go</code>, then run <code>velixir deploy</code>.</p>
</main>`

func main() {
	// velixir injects PORT. Listening on ":"+port (rather than "localhost:"+port)
	// binds every interface, which is what the platform health check needs.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, page)
	})

	log.Printf("listening on %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
