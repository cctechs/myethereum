package main

import (
	"io/ioutil"
	"os"
	"github.com/ethereum/go-ethereum/ethdb"
	"fmt"
)

func newLdb() (*ethdb.LDBDatabase, func()) {
	dirname, err := ioutil.TempDir(os.TempDir(), "ethdb_test_")
	if err != nil{
		panic(err)
	}
	db, err := ethdb.NewLDBDatabase(dirname, 0, 0)
	if err != nil{
		panic(err)
	}
	fmt.Printf("dirname=%s\n", dirname)
	return db, func() {
		db.Close()
		os.RemoveAll(dirname)
	}
}

func main(){

	values := []string{"aa", "bb", "cc"}

	db, remove := newLdb()
	defer remove()

	for _, v := range values{
		err := db.Put([]byte(v), []byte(v))
		if err != nil{
			panic(err)
		}
	}

	for _, v := range values{
		str, err := db.Get([]byte(v))
		if err != nil{
			panic(err)
		}
		fmt.Printf("%s\n", str)
	}
}
