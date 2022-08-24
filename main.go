package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("email verifier program: ......")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain, hasMx, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")

	for scanner.Scan() {
		check_domain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from the input %v\n", err)
	}
}

func check_domain(domain string) {

	var hasMx, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecord, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(mxRecord) > 0 {
		hasMx = true
	}

	txtRecord, err1 := net.LookupTXT(domain)

	if err1 != nil {
		log.Printf("Error: %v\n", err1)
	}
	for _, record := range txtRecord {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err2 := net.LookupTXT("_dmarc." + domain)
	if err2 != nil {
		log.Printf("Error: %v\n", err2)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMx, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
