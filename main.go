package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/AElfProject/aelf-sdk.go/client"
)

var (
	allWatching  []*Watching
	port         string
	updates      string
	prefix       string
	loadSeconds  float64
	totalLoaded  int64
	aelfClient   *client.AElfClient
	sleepSeconds int
)

type Watching struct {
	Name    string
	Address string
	Balance string
}

// Reference: https://docs.aelf.com/tools/chain-sdk/go-sdk/
// TODO: Connect to AElf node
func ConnectToAElf(url string) error {
	var err error
	aelfClient, err = client.NewAElfClient(url)
	return err
}

// TODO: Fetch balance from AElf node
func GetAElfBalance(address string) *big.Float {
	balance, err := aelfClient.GetBalance(context.TODO(), address, "ELF")
	if err != nil {
		log.Printf("Error fetching balance for address: %v\n", address)
		return big.NewFloat(0)
	}
	return new(big.Float).SetFloat64(balance.Balance)
}

// HTTP response handler for /metrics
func MetricsHttp(w http.ResponseWriter, r *http.Request) {
	var allOut []string
	total := big.NewFloat(0)
	for _, v := range allWatching {
		if v.Balance == "" {
			v.Balance = "0"
		}
		bal := big.NewFloat(0)
		bal.SetString(v.Balance)
		total.Add(total, bal)
		allOut = append(allOut, fmt.Sprintf("%vaelf_balance{name=\"%v\",address=\"%v\"} %v", prefix, v.Name, v.Address, v.Balance))
	}
	allOut = append(allOut, fmt.Sprintf("%vaelf_balance_total %0.18f", prefix, total))
	allOut = append(allOut, fmt.Sprintf("%vaelf_load_seconds %0.2f", prefix, loadSeconds))
	allOut = append(allOut, fmt.Sprintf("%vaelf_loaded_addresses %v", prefix, totalLoaded))
	allOut = append(allOut, fmt.Sprintf("%vaelf_total_addresses %v", prefix, len(allWatching)))
	fmt.Fprintln(w, strings.Join(allOut, "\n"))
}

// Open the addresses.txt file (name:address)
func OpenAddresses(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		object := strings.Split(scanner.Text(), ":")
		if len(object) != 2 || !isValidAddress(object[1]) {
			continue
		}
		w := &Watching{
			Name:    object[0],
			Address: object[1],
		}
		allWatching = append(allWatching, w)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

// TODO: Check if an address is valid
func isValidAddress(address string) bool {
	return true
}

func main() {
	aelfUrl := os.Getenv("AELF_URL")
	port = os.Getenv("PORT")
	prefix = os.Getenv("PREFIX")
	sleepDuration := os.Getenv("SLEEP_DURATION")

	if aelfUrl == "" || port == "" || prefix == "" || sleepDuration == "" {
		log.Fatal("Environment variables AELF_URL, PORT, PREFIX, and SLEEP_DURATION must be set")
	}

	var err error
	sleepSeconds, err = strconv.Atoi(sleepDuration)
	if err != nil {
		log.Fatal("SLEEP_DURATION must be an integer")
	}

	err = OpenAddresses("addresses.txt")
	if err != nil {
		log.Fatalf("Failed to open addresses.txt: %v", err)
	}

	err = ConnectToAElf(aelfUrl)
	if err != nil {
		log.Fatalf("Failed to connect to AElf node: %v", err)
	}

	// Check address balances
	go func() {
		for {
			totalLoaded = 0
			t1 := time.Now()
			log.Printf("Checking %v wallets...\n", len(allWatching))
			for _, v := range allWatching {
				v.Balance = GetAElfBalance(v.Address).String()
				totalLoaded++
			}
			t2 := time.Now()
			loadSeconds = t2.Sub(t1).Seconds()
			log.Printf("Finished checking %v wallets in %0.2f seconds, sleeping for %v seconds.\n", len(allWatching), loadSeconds, sleepSeconds)
			time.Sleep(time.Duration(sleepSeconds) * time.Second)
		}
	}()

	log.Printf("AELFexporter has started on port %v using AElf node: %v\n", port, aelfUrl)
	http.HandleFunc("/metrics", MetricsHttp)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
