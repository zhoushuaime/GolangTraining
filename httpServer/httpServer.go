package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func logs(r *http.Request) {
	log.Printf("%v %v %v\n",
		r.Host, r.Method, r.URL)

}

func sayHelloServer(w http.ResponseWriter, r *http.Request) {

	//r.ParseForm()
	//for k, v := range r.Form {
	//	fmt.Println("key:", k)
	//	fmt.Println("values:", v)
	//}
	logs(r)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"data":"write something"}`))
	//ctx := context.WithValue(r.Context(), "name", "zhoushuai")
	//createData(w, r.WithContext(ctx))

}

func createData(w http.ResponseWriter, r *http.Request) {
	logs(r)
	name := r.Context().Value("name").(string)
	fmt.Println("name:", name)
	_, _ = w.Write([]byte("createData"))
}

func main() {
	http.HandleFunc("/v1/hello", sayHelloServer)
	http.HandleFunc("/v1/c", createData)

	server := &http.Server{
		Addr:    ":8080",
		Handler: http.DefaultServeMux,
	}
	signalChan := make(chan os.Signal, 1)
	signal.Notify(
		signalChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGKILL,
	)
	go func() {
		select {
		case <-signalChan:
			println("shutdown...")
			timeout := 1 * time.Second
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				println("HTTP Server Shutdown error:", err)
			}
			println("server has stop")

		}

	}()
	fmt.Println("httpServer started at port", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		println("start server error:", err)
	}
	//if err := server.ListenAndServeTLS("apiclient_cert.pem", "apiclient_key.pem"); err != nil {
	//	println("start server error:", err)
	//}
}
