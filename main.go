package main

import (
	"context"
	"fmt"	
	"time"
	"encoding/json"
	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"syscall/js"
)

func listBuckets() js.Func {
	
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// Sets your Google Cloud Platform project ID.
		projectID := "orai-305206"
		keybytes := []byte(args[0].String())
		ctx := context.Background()
		client, err := storage.NewClient(ctx,option.WithCredentialsJSON(keybytes))
		if err != nil {
			return err.Error()			
		}
		defer client.Close()

		ctx, cancel := context.WithTimeout(ctx, time.Second*30)
		defer cancel()

		var buckets []string
		it := client.Buckets(ctx, projectID)
		for {
				battrs, err := it.Next()
				if err == iterator.Done {
						break
				}
				if err != nil {						
						return err.Error()
				}
				buckets = append(buckets, battrs.Name)				
		}
		pretty, _ := json.MarshalIndent(buckets, "", "  ")
		return string(pretty)
	})
	return jsonFunc
}

func main() {  
    fmt.Println("Go Web Assembly")
    js.Global().Set("listBuckets", listBuckets())
	<-make(chan bool)
}