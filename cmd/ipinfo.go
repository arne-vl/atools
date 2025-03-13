package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
	"net/http"
	"io/ioutil"
)

// ipinfoCmd represents the ipinfo command
var ipinfoCmd = &cobra.Command{
	Use:   "ipinfo",
	Short: "Display the IP info of the current machine",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		privateIP, err := getPrivateIP()
		if err != nil {
			fmt.Printf("Error getting private IP: %v\n", err)
			return
		}

		publicIP, err := getPublicIP()
		if err != nil {
			fmt.Printf("Error getting public IP: %v\n", err)
			return
		}

		hostname, err := os.Hostname()
		if err != nil {
			fmt.Printf("Error getting hostname: %v\n", err)
			return
		}

		fmt.Printf("%-15s %s\n", "Hostname:", hostname)
		fmt.Printf("%-15s %s\n", "Private IP:", privateIP)
		fmt.Printf("%-15s %s\n", "Public IP:", publicIP)
	},
}

// Gets the private IP address of the machine
func getPrivateIP() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}

	ips, err := net.LookupIP(hostname)
	if err != nil {
		return "", err
	}

	for _, ip := range ips {
		if ip.To4() != nil {
			return ip.String(), nil
		}
	}
	return "", fmt.Errorf("could not find a valid private IP address")
}

// Gets the public IP address of the machine
func getPublicIP() (string, error) {
	resp, err := http.Get("https://api64.ipify.org")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func init() {
	rootCmd.AddCommand(ipinfoCmd)
}
