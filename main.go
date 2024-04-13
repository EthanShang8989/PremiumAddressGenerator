package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"runtime"
	"strings"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

type Result struct {
	address    string
	privateKey string
}

func main() {
	var prefix, suffix string
	var cores int

	flag.StringVar(&prefix, "prefix", "", "Prefix of the Ethereum address")
	flag.StringVar(&suffix, "suffix", "", "Suffix of the Ethereum address")
	flag.IntVar(&cores, "cores", 1, "Number of CPU cores to use")

	flag.Parse()
	if cores < 1 || cores > runtime.NumCPU() {
		log.Fatalf("Invalid number of cores. Use a number between 1 and %d", runtime.NumCPU())
	}
	runtime.GOMAXPROCS(cores)
	resultChan := make(chan Result)

	var count uint64
	const batchSize = 5000
	ctx, Canecl := context.WithCancel(context.Background())
	go func() {
		for {
			<-time.After(1 * time.Second)
			fmt.Printf("\rAddresses checked: %d", atomic.LoadUint64(&count))
			fmt.Print("     ")
		}
	}()
	for i := 0; i < cores; i++ {
		go func() {
			resultChan <- generateEthAddress(ctx, prefix, suffix, &count, batchSize)
		}()
	}
	result := <-resultChan
	Canecl()
	totalAddresses := atomic.LoadUint64(&count)
	fmt.Printf("\nFound matching address: %s\n", result.address)
	fmt.Printf("Private Key: %s\n", result.privateKey)
	fmt.Printf("Total addresses checked: %d\n", totalAddresses)
}

func generateEthAddress(ctx context.Context, prefix, suffix string, count *uint64, batchSize uint) (result Result) {
	localCount := uint(0)
	for {
		select {
		case <-ctx.Done():
			atomic.AddUint64(count, uint64(localCount))
			return
		default:
			key, err := crypto.GenerateKey()
			if err != nil {
				log.Fatalf("Failed to generate key: %v", err)
			}
			localCount++
			if localCount >= batchSize {
				atomic.AddUint64(count, uint64(localCount))
				localCount = 0
			}
			address := crypto.PubkeyToAddress(key.PublicKey).Hex()
			if strings.HasPrefix(address, "0x"+prefix) && strings.HasSuffix(address, suffix) {
				if localCount > 0 {
					atomic.AddUint64(count, uint64(localCount))
				}
				privateKeyBytes := crypto.FromECDSA(key)
				return Result{
					address:    address,
					privateKey: fmt.Sprintf("%x", privateKeyBytes),
				}
			}
		}
	}
}
