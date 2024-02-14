/*
Copyright © 2024 Abdullah Bagyapan <abdullahbagyapan@hotmail.com>
*/
package cmd

import (
	"log"
	"net"
	"net/http"
	"net/netip"
	"os"

	"github.com/spf13/cobra"
)

var addr net.IP
var port uint16
var directory string

var rootCmd = &cobra.Command{
	Use:   "httpserver",
	Short: "Simple HTTP server application",
	Long: `httpserver is a simple HTTP server CLI application, that written in Go.
This application is a tool to create basic HTTP server.`,

	Run: func(cmd *cobra.Command, args []string) {

		// Directory is exist or not
		_, err := os.Stat(directory)
		if err != nil {
			os.Stderr.WriteString(err.Error())
			os.Exit(1)
		}

		ipAddr, ok := netip.AddrFromSlice(addr)
		if !ok {
			os.Stderr.WriteString("error parsing IP address")
			os.Exit(1)
		}

		laddr := netip.AddrPortFrom(ipAddr, port)

		http.Handle("/", http.FileServer(http.Dir(directory)))

		ch := make(chan struct{}, 1)

		go func() {
			ln, err := net.Listen("tcp", laddr.String())

			if err != nil {
				log.Fatalf("error serving HTTP server %v", err)
			}
			ch <- struct{}{}

			if err := http.Serve(ln, nil); err != nil {
				log.Fatalf("error listening HTTP server %v", err)
			}
		}()

		<-ch
		log.Printf("HTTP server listening on %s:%d", addr, port)

		<-ch
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	// define flags
	rootCmd.Flags().IPVarP(&addr, "addr", "a", net.IP{127, 0, 0, 1}, "IP address for HTTP server")
	rootCmd.Flags().Uint16VarP(&port, "port", "p", 8080, "Port number for HTTP server")

	homeDir := os.Getenv("HOME")

	rootCmd.Flags().StringVarP(&directory, "directory", "d", homeDir, "Directory for HTTP server")

}
