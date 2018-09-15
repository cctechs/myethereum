package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	node2 "github.com/ethereum/go-ethereum/node"
	"github.com/codegangsta/cli"
	"path/filepath"
	"os"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/cmd/utils"
)

var (
	gitCommit = ""
	app = utils.NewApp(gitCommit, "my ethereum")
	
	nodeFlags = []cli.Flag{
		utils.IdentityFlag,
	}
)


func geth(ctx *cli.Context) error{
	//node :=
	return nil
}


func init(){
	app.Action = nil
}


func NewApp(gitCommit, usage string) *cli.App {
	app := cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Author = ""
	//app.Authors = nil
	app.Email = ""
	app.Version = params.Version
	if len(gitCommit) >= 8 {
		app.Version += "-" + gitCommit[:8]
	}
	app.Usage = usage
	return app
}


func main(){

	app := NewApp("", "test123")
	fmt.Println(app)

	config := node2.Config{DataDir:"/Users/wubo/Documents/blockchain/privatechain/data0"}

	fmt.Println(config.NodeDB())
	fmt.Println(config.NodeName())
	fmt.Println(config.IPCEndpoint())

	node, err := node2.New(&config)
	fmt.Println(err)
	fmt.Println(node)


	keyStore := keystore.NewKeyStore(
		"/Users/wubo/Documents/blockchain/privatechain/data0",
		keystore.LightScryptN, keystore.LightScryptP)

	fmt.Println(keyStore)

	fmt.Println("hello world")
}
