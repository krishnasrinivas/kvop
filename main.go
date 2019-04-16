package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	configPath_ := flag.String("c", "", "nkv_config.json's path")
	path_ := flag.String("i", "", "IPADDR")
	key_ := flag.String("k", "", "key")
	value_ := flag.String("v", "", "value")
	get_ := flag.Bool("g", false, "get")
	put_ := flag.Bool("p", false, "put")
	list_ := flag.Bool("l", false, "list")
	del_ := flag.Bool("d", false, "delete")

	flag.Parse()

	configPath := *configPath_
	key := *key_
	value := *value_
	path := *path_
	get := *get_
	put := *put_
	list := *list_
	del := *del_

	if configPath == "" {
		fmt.Println("config path null")
	}
	if key == "" {
		fmt.Println("key null")
	}

	err := minio_nkv_open(configPath)
	if err != nil {
		log.Fatal(err)
	}
	kv, err := newKV(path, true)
	if err != nil {
		log.Fatal(err)
	}

	if put {
		err = kv.Put(key, []byte(value))
		if err != nil {
			log.Fatal(err)
		}
	}
	if get {
		b := make([]byte, 2*1024*1024)
		b, err = kv.Get(key, b)
		fmt.Println(string(b), err)
	}
	if del {
		err = kv.Delete(key)
		if err != nil {
			log.Fatal(err)
		}
	}
	if list {
		entries, err := kv.List(key)
		fmt.Println(entries, err)
	}
}
