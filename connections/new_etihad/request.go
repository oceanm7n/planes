package new_etihad

type requestBody struct {
	OperationName string `json:"operationName"`
	Variables     struct {
		AirSearchInput struct {
			CabinClass     string        `json:"cabinClass"`
			AwardBooking   bool          `json:"awardBooking"`
			PromoCodes     []interface{} `json:"promoCodes"`
			SearchType     string        `json:"searchType"`
			ItineraryParts []struct {
				From struct {
					UseNearbyLocations bool   `json:"useNearbyLocations"`
					Code               string `json:"code"`
				} `json:"from"`
				To struct {
					UseNearbyLocations bool   `json:"useNearbyLocations"`
					Code               string `json:"code"`
				} `json:"to"`
				When struct {
					Date string `json:"date"`
				} `json:"when"`
			} `json:"itineraryParts"`
			Passengers struct {
				Adt int `json:"ADT"`
			} `json:"passengers"`
		} `json:"airSearchInput"`
	} `json:"variables"`
	Extensions struct {
	} `json:"extensions"`
	Query string `json:"query"`
}
