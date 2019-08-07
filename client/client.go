package main

import (
	"bytes"
	"io"
	"log"
	"net"

	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("port", "12001")
	viper.SetDefault("max", "20")

	viper.AutomaticEnv()
}

func main() {
	for i := 0; i < viper.GetInt("max"); i++ {
		conn, err := net.Dial("tcp", "127.0.0.1:"+viper.GetString("port"))
		if err != nil {
			log.Fatal(err)
		}

		buff := bytes.NewBufferString("simpler")
		io.Copy(conn, buff)
		conn.Close()
	}
}
