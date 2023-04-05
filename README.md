# OpenPeopleSearch

OpenPeopleSearch is a simple library for querying the OpenPeopleSearch API to find information on people and businesses.

## Table of Contents
- [Introduction](#introduction)
- [Installation](#installation)
- [Usage](#usage)
- [Available Search Methods](#available-search-methods)
- [Contributing](#contributing)
- [License](#license)

## Introduction
This library provides easy-to-use search methods for accessing OpenPeopleSearch data. The searches are performed using various parameters such as name, phone number, email address, address, business name, or PO Box.

## Installation
To include this library in your project, simply import the package:
```go
go get "github.com/Minagoroshi/OpenPeopleSearch-API-Wrapper"
```

## Usage
First, create a session with your API token:
```go
session := openpeoplesearch.NewSession("your_api_token")
```
Then, you can use the session to call search functions with the required parameters.
```go
result, err := session.PhoneSearch("555-123-4567")
  if err != nil {
  log.Fatal(err)
}

fmt.Println(result)
```


## Available Search Methods
The following search methods are available in this library:
- PhoneSearch(phoneNumber string) (SearchResult, error)
- EmailSearch(emailAddress string) (SearchResult, error)
- NameSearch(firstName, middleName, lastName, city, state string) (SearchResult, error)
- NameAddressSearch(firstName, lastName, address, unit, city, state string) (SearchResult, error)
- NameDobSearch(firstName, lastName, dob string) (SearchResult, error)
- BusinessSearch(businessName, city, state string) (SearchResult, error)
- AddressSearch(address, unit, city, state string) (SearchResult, error)
- PoBoxSearch(poBox, city, state string) (SearchResult, error)

## Contributing
Contributions are welcome! Please submit a pull request or create an issue with any improvements or suggestions.

## License
This library is distributed under the MIT License. See [License](LICENSE) for more information.
