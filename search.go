package openpeoplesearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type phoneSearchRequest struct {
	PhoneNumber string `json:"phoneNumber"`
}

type emailAddressSearchRequest struct {
	EmailAddress string `json:"emailAddress"`
}

type nameSearchRequest struct {
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	City       string `json:"city"`
	State      string `json:"state"`
}

type nameAddressSearchRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Address   string `json:"address"`
	Unit      string `json:"unit"`
	City      string `json:"city"`
	State     string `json:"state"`
}

type nameDobSearchRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Dob       string `json:"dob"`
}

type businessSearchRequest struct {
	BusinessName string `json:"businessName"`
	City         string `json:"city"`
	State        string `json:"state"`
}

type addressSearchRequest struct {
	Address string `json:"address"`
	Unit    string `json:"unit"`
	City    string `json:"city"`
	State   string `json:"state"`
}

type poBoxSearchRequest struct {
	PoBox string `json:"poBox"`
	City  string `json:"city"`
	State string `json:"state"`
}

func (s *Session) PhoneSearch(phoneNumber string) (SearchResult, error) {
	payload := phoneSearchRequest{PhoneNumber: phoneNumber}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return SearchResult{}, err
	}
	return s.search(jsonPayload, SEARCH_PHONE)
}

func (s *Session) EmailSearch(emailAddress string) (SearchResult, error) {
	payload := emailAddressSearchRequest{EmailAddress: emailAddress}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return SearchResult{}, err
	}
	return s.search(jsonPayload, SEARCH_EMAIL)
}

func (s *Session) NameSearch(firstName, middleName, lastName, city, state string) (SearchResult, error) {
	payload := nameSearchRequest{FirstName: firstName, MiddleName: middleName, LastName: lastName, City: city, State: state}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return SearchResult{}, err
	}
	return s.search(jsonPayload, SEARCH_NAME)
}

func (s *Session) NameAddressSearch(firstName, lastName, address, unit, city, state string) (SearchResult, error) {
	payload := nameAddressSearchRequest{FirstName: firstName, LastName: lastName, Address: address, Unit: unit, City: city, State: state}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return SearchResult{}, err
	}
	return s.search(jsonPayload, SEARCH_NAMEADDRESS)
}

func (s *Session) NameDobSearch(firstName, lastName, dob string) (SearchResult, error) {
	payload := nameDobSearchRequest{FirstName: firstName, LastName: lastName, Dob: dob}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return SearchResult{}, err
	}
	return s.search(jsonPayload, SEARCH_NAMEDOB)
}

func (s *Session) BusinessSearch(businessName, city, state string) (SearchResult, error) {
	payload := businessSearchRequest{BusinessName: businessName, City: city, State: state}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return SearchResult{}, err
	}
	return s.search(jsonPayload, SEARCH_BUSINESS)
}

func (s *Session) AddressSearch(address, unit, city, state string) (SearchResult, error) {
	payload := addressSearchRequest{Address: address, Unit: unit, City: city, State: state}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return SearchResult{}, err
	}
	return s.search(jsonPayload, SEARCH_ADDRESS)
}

func (s *Session) PoBoxSearch(poBox, city, state string) (SearchResult, error) {
	payload := poBoxSearchRequest{PoBox: poBox, City: city, State: state}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return SearchResult{}, err
	}
	return s.search(jsonPayload, SEARCH_POBOX)
}

func (s *Session) search(body []byte, endpoint string) (SearchResult, error) {

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		return SearchResult{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.Token)
	req.Header.Set("Accept", "text/plain")
	req.Header.Set("User-Agent", USERAGENT)

	resp, err := opsClient.Do(req)
	if err != nil {
		return SearchResult{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return SearchResult{}, fmt.Errorf("PhoneSearch failed with status code %d", resp.StatusCode)
	}

	var result SearchResult
	err = json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}
