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
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Domain Name")
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	err := scanner.Err()
	if err != nil {
		log.Fatal("Error: Could Not read from Input")
	}
}

func checkDomain(domain string) {
	/*
		domain: This refers to the domain name being checked.
		hasMX: This indicates whether the domain has an MX (mail exchange) record in its DNS settings, which specifies the mail servers that are responsible for accepting emails sent to the domain.
		mxRecords: This gives the actual MX record for the domain, if it exists
		hasSPF: This indicates whether the domain has an SPF (Sender Policy Framework) record in its DNS settings, which specifies the IP addresses or domains that are authorized to send emails on behalf of the domain.
		spfRecord: This gives the actual SPF record for the domain, if it exists.
		hasDMARC: This indicates whether the domain has a DMARC (Domain-based Message Authentication, Reporting and Conformance) record in its DNS settings, which specifies how to handle emails that fail SPF or DKIM (DomainKeys Identified Mail) authentication checks.
		dmarcRecord: This gives the actual DMARC record for the domain, if it exists.
	*/
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v \n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("error : %v \n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc" + domain)
	if err != nil {
		log.Printf("error %v", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("Domain: %v \n", domain)
	fmt.Printf("hasMX: %v \n", hasMX)
	fmt.Printf("mxRecords: %v \n", mxRecords)
	fmt.Printf("hasSPF: %v \n", hasSPF)
	fmt.Printf("spfRecord: %v \n", spfRecord)
	fmt.Printf("hasDMARC: %v \n", hasDMARC)
	fmt.Printf("dmarcRecord: %v \n", dmarcRecord)
}
