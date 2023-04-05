package openpeoplesearch

const (
	BASEURL = "https://api.openpeoplesearch.com/api/v1/"

	// User routes
	USER = BASEURL + "User"       // POST
	AUTH = USER + "/authenticate" // POST

	// Consumer routes
	SEARCH             = BASEURL + "Consumer"           // POST
	SEARCH_PHONE       = SEARCH + "/PhoneSearch"        // POST
	SEARCH_EMAIL       = SEARCH + "/EmailAddressSearch" // POST
	SEARCH_NAME        = SEARCH + "/NameSearch"         // POST
	SEARCH_NAMEADDRESS = SEARCH + "/NameAddressSearch"  // POST
	SEARCH_NAMEDOB     = SEARCH + "/NameDOBSearch"      // POST
	SEARCH_BUSINESS    = SEARCH + "/BusinessSearch"     // POST
	SEARCH_ADDRESS     = SEARCH + "/AddressSearch"      // POST
	SEARCH_POBOX       = SEARCH + "/PoBoxSearch"        // POST
)
