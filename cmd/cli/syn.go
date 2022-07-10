package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
		controller := cli.NewApp()

		controller.Name = "Syndicate"
		controller.Usage = "Network CLI tool for gathering information across multiple areas"

		version := "v0.1.4"
		author := "Hifumi1337 (https://github.com/Hifumi1337)"

		defaultValues := []cli.Flag {
			cli.StringFlag {
				// Default values for commands
				Name: "host",
				Value: "google.com",
			},
		}

		// Command List
		controller.Commands = []cli.Command {
		{
			Name: "version",
			Usage: "Version information",
			Flags: defaultValues,

		Action: func(cmd *cli.Context) error {
				fmt.Println("Syndicate |", version)
				fmt.Println("Author:", author)
				
				return nil
			},
		},
		{
			Name: "ns",
			Usage: "Gathers nameserver information about the specified host",
			Flags: defaultValues,

			// ns command activated
			Action: func(cmd *cli.Context) error {
				ns, err := net.LookupNS(cmd.String("host"))

				if err != nil {
					return err
				}

				// Nameservers
				fmt.Println("Nameserver(s):")
				fmt.Println("==============================")

				// Logs results to terminal
				for i := 0; i < len(ns); i++ {
					fmt.Println("[*]", ns[i].Host)
				}

				fmt.Println("==============================")

				return nil
			},
		},
		{
			Name: "ip",
			Usage: "Gathers IP addresses about the specified host",
			Flags:defaultValues,

			// ip command activated
			Action: func(cmd *cli.Context) error {
				ip, err := net.LookupIP(cmd.String("host"))

				if err != nil {
					fmt.Println(err)
				}

				// IP Address(es)
				fmt.Println("IP Address(es):")
				fmt.Println("==============================")

				// Logs results to terminal
				for i := 0; i < len(ip); i++ {
					fmt.Println("[*]", ip[i])
				}

				fmt.Println("==============================")

				return nil
			},
		},
		{
			Name: "cname",
			Usage: "Gathers CNAME information about the specified host",
			Flags:defaultValues,

			// cname command activated
			Action: func(cmd *cli.Context) error {
				cname, err := net.LookupCNAME(cmd.String("host"))

				if err != nil {
					fmt.Println(err)
				}

				// CNAME
				fmt.Println("CNAME:")
				fmt.Println("==============================")

				// Logs results to terminal
				fmt.Println("[*]", cname)

				fmt.Println("==============================")
				
				return nil
			},
		},
		{
			Name: "mx",
			Usage: "Gathers MX records from the specified host",
			Flags:defaultValues,

			// mx command activated
			Action: func(cmd *cli.Context) error {
				mx, err := net.LookupMX(cmd.String("host"))

				if err != nil {
					fmt.Println(err)
				}

				// MX Record
				fmt.Println("MX Record(s):")
				fmt.Println("==============================")

				// Log results to terminal
				for i := 0; i < len(mx); i++ {
					fmt.Println("Host:", mx[i].Host)
					fmt.Println("Priority:", mx[i].Pref)
				}

				fmt.Println("==============================")

				return nil
			},
		},
		{
			Name: "txt",
			Usage: "Gathers TXT records from the specified host",
			Flags:defaultValues,

			// txt command activated
			Action: func(cmd *cli.Context) error {
				txt, err := net.LookupTXT(cmd.String("host"))

				if err != nil {
					fmt.Println(err)
				}

				// TXT Record
				fmt.Println("TXT Record(s):")
				fmt.Println("==============================")

				// Logs results to terminal
				for i := 0; i < len(txt); i++ {
					fmt.Println("[*]", txt[i])
				}

				fmt.Println("==============================")

				return nil
			},
		},
		{
			Name: "all",
			Usage: "Prints every command in a single output",
			Flags:defaultValues,

			// all command activated
			Action: func(cmd *cli.Context) error {
				ns, err := net.LookupNS(cmd.String("host"))
				ip, err := net.LookupIP(cmd.String("host"))
				cname, err := net.LookupCNAME(cmd.String("host"))
				mx, err := net.LookupMX(cmd.String("host"))
				txt, err := net.LookupTXT(cmd.String("host"))

				if err != nil {
					fmt.Println(err)
				}
				
				fmt.Println("Syndicate |", version)
				fmt.Println("Author:", author)

				fmt.Println("\n")

				// Nameservers
				fmt.Println("Nameserver(s):")
				fmt.Println("==============================")

				for i := 0; i < len(ns); i++ {
					fmt.Println("[*]", ns[i].Host)
				}

				fmt.Println("\n")

				// IP Address(es)
				fmt.Println("IP Address(es):")
				fmt.Println("==============================")

				for i := 0; i < len(ip); i++ {
					fmt.Println("[*]", ip[i])
				}

				fmt.Println("\n")

				// CNAME
				fmt.Println("CNAME:")
				fmt.Println("==============================")
				
				fmt.Println("[*]", cname)
				
				fmt.Println("\n")

				// MX Record
				fmt.Println("MX Record(s):")
				fmt.Println("==============================")

				for i := 0; i < len(mx); i++ {
					fmt.Println("Host:", mx[i].Host)
					fmt.Println("Priority:", mx[i].Pref)
				}

				fmt.Println("\n")

				// TXT Record
				fmt.Println("TXT Record(s):")
				fmt.Println("==============================")

				// Logs results to terminal
				for i := 0; i < len(txt); i++ {
					fmt.Println("[*]", txt[i])
				}

				return nil
			},
		},
	}

	// Initializing Engine...
	err := controller.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}