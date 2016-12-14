package mailgun

import (
	"testing"
)

func TestGetDomains(t *testing.T) {
	mg, err := NewMailgunFromEnv()
	if err != nil {
		t.Fatalf("NewMailgunFromEnv() error - %s", err.Error())
	}
	n, domains, err := mg.GetDomains(DefaultLimit, DefaultSkip)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("TestGetDomains: %d domains retrieved\n", n)
	for _, d := range domains {
		t.Logf("TestGetDomains: %#v\n", d)
	}
}

func TestGetSingleDomain(t *testing.T) {
	mg, err := NewMailgunFromEnv()
	if err != nil {
		t.Fatalf("NewMailgunFromEnv() error - %s", err.Error())
	}
	_, domains, err := mg.GetDomains(DefaultLimit, DefaultSkip)
	if err != nil {
		t.Fatal(err)
	}
	dr, rxDnsRecords, txDnsRecords, err := mg.GetSingleDomain(domains[0].Name)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("TestGetSingleDomain: %#v\n", dr)
	for _, rxd := range rxDnsRecords {
		t.Logf("TestGetSingleDomains:   %#v\n", rxd)
	}
	for _, txd := range txDnsRecords {
		t.Logf("TestGetSingleDomains:   %#v\n", txd)
	}
}

func TestGetSingleDomainNotExist(t *testing.T) {
	mg, err := NewMailgunFromEnv()
	if err != nil {
		t.Fatalf("NewMailgunFromEnv() error - %s", err.Error())
	}
	_, _, _, err = mg.GetSingleDomain(randomString(32, "com.edu.org.") + ".com")
	if err == nil {
		t.Fatal("Did not expect a domain to exist")
	}
	ure, ok := err.(*UnexpectedResponseError)
	if !ok {
		t.Fatal("Expected UnexpectedResponseError")
	}
	if ure.Actual != 404 {
		t.Fatalf("Expected 404 response code; got %d", ure.Actual)
	}
}

func TestAddDeleteDomain(t *testing.T) {
	mg, err := NewMailgunFromEnv()
	if err != nil {
		t.Fatalf("NewMailgunFromEnv() error - %s", err.Error())
	}
	// First, we need to add the domain.
	randomDomainName := randomString(16, "DOMAIN") + ".example.com"
	randomPassword := randomString(16, "PASSWD")
	err = mg.CreateDomain(randomDomainName, randomPassword, Tag, false)
	if err != nil {
		t.Fatal(err)
	}

	// Next, we delete it.
	err = mg.DeleteDomain(randomDomainName)
	if err != nil {
		t.Fatal(err)
	}
}
