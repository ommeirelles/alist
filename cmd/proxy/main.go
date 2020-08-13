package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, time.Second * 5)
		defer cancel()

		// endpoint de queries permitidas validar se é uma query permitida

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:8888/graphql", r.Body)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("Content-type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("err Do", err.Error())
			return
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Println("err status code")
			return
		}

		_, err = io.Copy(rw, resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	})

	_ = http.ListenAndServe(":8080", nil)
}