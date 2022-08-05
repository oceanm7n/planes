package new_etihad

type responseBody struct {
	Data struct {
		BookingAirSearch struct {
			OriginalResponse struct {
				SearchResultMetaData struct {
					Branded            bool `json:"branded"`
					MultipleDateResult bool `json:"multipleDateResult"`
					ComposedResult     bool `json:"composedResult"`
					InterlineRoute     bool `json:"interlineRoute"`
					ContextShopping    bool `json:"contextShopping"`
				} `json:"searchResultMetaData"`
				FareFamilies []struct {
					BrandID    string `json:"brandId"`
					BrandLabel []struct {
						ProgramID     string `json:"programId"`
						LanguageID    string `json:"languageId"`
						MarketingText string `json:"marketingText"`
					} `json:"brandLabel"`
					MarketingTexts []struct {
						ProgramID  string `json:"programId"`
						LanguageID string `json:"languageId"`
					} `json:"marketingTexts"`
					BrandAncillaries struct {
						FlightAncillary []interface{} `json:"flightAncillary"`
					} `json:"brandAncillaries"`
					FareFamilyRemarkRPH int `json:"fareFamilyRemarkRPH"`
				} `json:"fareFamilies"`
				UnbundledOffers [][]struct {
					ShoppingBasketHashCode int    `json:"shoppingBasketHashCode"`
					BrandID                string `json:"brandId"`
					Soldout                bool   `json:"soldout"`
					BundlePrice            struct {
					} `json:"bundlePrice"`
					AvailableObFees []interface{} `json:"availableObFees"`
					SeatsRemaining  struct {
						Count           int  `json:"count"`
						LowAvailability bool `json:"lowAvailability"`
					} `json:"seatsRemaining"`
					ItineraryPart []struct {
						Type     string `json:"@type"`
						ID       string `json:"@id"`
						Segments []struct {
							Type                    string `json:"@type"`
							ID                      string `json:"@id"`
							SegmentOfferInformation struct {
								FlightsMiles int  `json:"flightsMiles"`
								AwardFare    bool `json:"awardFare"`
							} `json:"segmentOfferInformation"`
							Duration   int    `json:"duration"`
							CabinClass string `json:"cabinClass"`
							Equipment  string `json:"equipment"`
							Flight     struct {
								FlightNumber          int           `json:"flightNumber"`
								OperatingFlightNumber int           `json:"operatingFlightNumber"`
								AirlineCode           string        `json:"airlineCode"`
								OperatingAirlineCode  string        `json:"operatingAirlineCode"`
								StopAirports          []interface{} `json:"stopAirports"`
								Advisories            []interface{} `json:"advisories"`
								DepartureTerminal     string        `json:"departureTerminal"`
								ArrivalTerminal       string        `json:"arrivalTerminal"`
							} `json:"flight"`
							Origin                      string `json:"origin"`
							Destination                 string `json:"destination"`
							Departure                   string `json:"departure"`
							Arrival                     string `json:"arrival"`
							BookingClass                string `json:"bookingClass"`
							LayoverDuration             int    `json:"layoverDuration"`
							FareBasis                   string `json:"fareBasis"`
							SubjectToGovernmentApproval bool   `json:"subjectToGovernmentApproval"`
							FareComponentBeginAirport   string `json:"fareComponentBeginAirport,omitempty"`
							FareComponentEndAirport     string `json:"fareComponentEndAirport,omitempty"`
							FareComponentDirectionality string `json:"fareComponentDirectionality,omitempty"`
							Meals                       []struct {
								PassengerType string `json:"passengerType"`
								MealCode      string `json:"mealCode"`
							} `json:"meals"`
						} `json:"segments"`
						Stops                  int `json:"stops"`
						TotalDuration          int `json:"totalDuration"`
						ConnectionInformations []struct {
							ID              string `json:"@id"`
							Duration        int    `json:"duration"`
							ChangeOfAirport bool   `json:"changeOfAirport"`
						} `json:"connectionInformations"`
						BookingClass string        `json:"bookingClass"`
						ProgramIDs   []string      `json:"programIDs"`
						ProgramCodes []string      `json:"programCodes"`
						Advisories   []interface{} `json:"advisories"`
					} `json:"itineraryPart"`
					CabinClass       string `json:"cabinClass"`
					OfferInformation struct {
						Discounted     bool   `json:"discounted"`
						Negotiated     bool   `json:"negotiated"`
						NegotiatedType string `json:"negotiatedType"`
					} `json:"offerInformation"`
					Advisories []interface{} `json:"advisories"`
					Total      struct {
						Alternatives [][]struct {
							Amount   int    `json:"amount"`
							Currency string `json:"currency"`
						} `json:"alternatives"`
					} `json:"total"`
					Fare struct {
						Alternatives [][]struct {
							Amount   int    `json:"amount"`
							Currency string `json:"currency"`
						} `json:"alternatives"`
					} `json:"fare"`
					Taxes struct {
						Alternatives [][]struct {
							Amount   int    `json:"amount"`
							Currency string `json:"currency"`
						} `json:"alternatives"`
					} `json:"taxes"`
				} `json:"unbundledOffers"`
				UnbundledAlternateDateOffers [][]struct {
					BrandID     string `json:"brandId"`
					Soldout     bool   `json:"soldout"`
					BundlePrice struct {
					} `json:"bundlePrice"`
					AvailableObFees []interface{} `json:"availableObFees"`
					SeatsRemaining  struct {
						Count           int  `json:"count"`
						LowAvailability bool `json:"lowAvailability"`
					} `json:"seatsRemaining"`
					ItineraryPart []struct {
						Type     string `json:"@type"`
						ID       string `json:"@id"`
						Segments []struct {
							Type                    string `json:"@type"`
							ID                      string `json:"@id"`
							SegmentOfferInformation struct {
								FlightsMiles int  `json:"flightsMiles"`
								AwardFare    bool `json:"awardFare"`
							} `json:"segmentOfferInformation"`
							Duration   int    `json:"duration"`
							CabinClass string `json:"cabinClass"`
							Equipment  string `json:"equipment"`
							Flight     struct {
								FlightNumber          int           `json:"flightNumber"`
								OperatingFlightNumber int           `json:"operatingFlightNumber"`
								AirlineCode           string        `json:"airlineCode"`
								OperatingAirlineCode  string        `json:"operatingAirlineCode"`
								StopAirports          []interface{} `json:"stopAirports"`
								DepartureTerminal     string        `json:"departureTerminal"`
								ArrivalTerminal       string        `json:"arrivalTerminal"`
							} `json:"flight"`
							Origin                      string `json:"origin"`
							Destination                 string `json:"destination"`
							Departure                   string `json:"departure"`
							Arrival                     string `json:"arrival"`
							BookingClass                string `json:"bookingClass"`
							LayoverDuration             int    `json:"layoverDuration"`
							FareBasis                   string `json:"fareBasis"`
							SubjectToGovernmentApproval bool   `json:"subjectToGovernmentApproval"`
							FareComponentBeginAirport   string `json:"fareComponentBeginAirport,omitempty"`
							FareComponentEndAirport     string `json:"fareComponentEndAirport,omitempty"`
							FareComponentDirectionality string `json:"fareComponentDirectionality,omitempty"`
							Meals                       []struct {
								PassengerType string `json:"passengerType"`
								MealCode      string `json:"mealCode"`
							} `json:"meals"`
						} `json:"segments"`
						Stops                  int `json:"stops"`
						TotalDuration          int `json:"totalDuration"`
						ConnectionInformations []struct {
							ID              string `json:"@id"`
							Duration        int    `json:"duration"`
							ChangeOfAirport bool   `json:"changeOfAirport"`
						} `json:"connectionInformations"`
						BookingClass string   `json:"bookingClass"`
						ProgramIDs   []string `json:"programIDs"`
						ProgramCodes []string `json:"programCodes"`
					} `json:"itineraryPart"`
					CabinClass       string `json:"cabinClass"`
					OfferInformation struct {
						Discounted     bool   `json:"discounted"`
						Negotiated     bool   `json:"negotiated"`
						NegotiatedType string `json:"negotiatedType"`
					} `json:"offerInformation"`
					Status         string   `json:"status"`
					DepartureDates []string `json:"departureDates"`
					Total          struct {
						Alternatives [][]struct {
							Amount   int    `json:"amount"`
							Currency string `json:"currency"`
						} `json:"alternatives"`
					} `json:"total"`
					Fare struct {
						Alternatives [][]struct {
							Amount   int    `json:"amount"`
							Currency string `json:"currency"`
						} `json:"alternatives"`
					} `json:"fare"`
					Taxes struct {
						Alternatives [][]struct {
							Amount   int    `json:"amount"`
							Currency string `json:"currency"`
						} `json:"alternatives"`
					} `json:"taxes"`
				} `json:"unbundledAlternateDateOffers"`
				BundledOffers              []interface{} `json:"bundledOffers"`
				BundledAlternateDateOffers []interface{} `json:"bundledAlternateDateOffers"`
				BrandedResults             struct {
					ItineraryPartBrands [][]struct {
						ItineraryPart struct {
							Ref string `json:"@ref"`
						} `json:"itineraryPart"`
						BrandOffers []struct {
							ShoppingBasketHashCode int    `json:"shoppingBasketHashCode"`
							BrandID                string `json:"brandId"`
							Soldout                bool   `json:"soldout"`
							BundlePrice            struct {
							} `json:"bundlePrice"`
							AvailableObFees []interface{} `json:"availableObFees"`
							SeatsRemaining  struct {
								Count           int  `json:"count"`
								LowAvailability bool `json:"lowAvailability"`
							} `json:"seatsRemaining"`
							CabinClass       string `json:"cabinClass"`
							OfferInformation struct {
								Discounted     bool   `json:"discounted"`
								Negotiated     bool   `json:"negotiated"`
								NegotiatedType string `json:"negotiatedType"`
							} `json:"offerInformation"`
							Advisories []interface{} `json:"advisories"`
							Total      struct {
								Alternatives [][]struct {
									Amount   int    `json:"amount"`
									Currency string `json:"currency"`
								} `json:"alternatives"`
							} `json:"total"`
							TotalMandatoryObFees struct {
								Alternatives []interface{} `json:"alternatives"`
							} `json:"totalMandatoryObFees"`
							Fare struct {
								Alternatives [][]struct {
									Amount   int    `json:"amount"`
									Currency string `json:"currency"`
								} `json:"alternatives"`
							} `json:"fare"`
							Taxes struct {
								Alternatives [][]struct {
									Amount   int    `json:"amount"`
									Currency string `json:"currency"`
								} `json:"alternatives"`
							} `json:"taxes"`
						} `json:"brandOffers"`
						Duration  int    `json:"duration"`
						Departure string `json:"departure"`
						Arrival   string `json:"arrival"`
					} `json:"itineraryPartBrands"`
				} `json:"brandedResults"`
				TravelPartAdvisories       [][]interface{} `json:"travelPartAdvisories"`
				SoldOutDatesOutbound       []interface{}   `json:"soldOutDatesOutbound"`
				SoldOutDatesInbound        []interface{}   `json:"soldOutDatesInbound"`
				NoneScheduledDatesOutbound []interface{}   `json:"noneScheduledDatesOutbound"`
				NoneScheduledDatesInbound  []interface{}   `json:"noneScheduledDatesInbound"`
				Warnings                   []interface{}   `json:"warnings"`
				Currency                   string          `json:"currency"`
				PromocodeValid             bool            `json:"promocodeValid"`
				NegotiateFarePresent       bool            `json:"negotiateFarePresent"`
				ConversionRatesFound       bool            `json:"conversionRatesFound"`
			} `json:"originalResponse"`
			Typename string `json:"__typename"`
		} `json:"bookingAirSearch"`
	} `json:"data"`
	Extensions struct {
	} `json:"extensions"`
}
