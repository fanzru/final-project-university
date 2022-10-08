package outbound

import "encoding/xml"

// type TEI struct {
// 	Text Text `json:"text"`
// }

// type Text struct {
// 	Body Body `json:"body"`
// }

// type Body struct {
// 	Div []Div `json:"div"`
// }

// type Div struct {
// 	Head string `json:"head"`
// 	P    []P    `json:"p"`
// }

// type P struct {
// 	Text string `json:"#text"`
// 	Ref  string `json:"ref"`
// }

type TEI struct {
	XMLName        xml.Name `xml:"TEI"`
	Chardata       string   `xml:",chardata"`
	Space          string   `xml:"space,attr"`
	Xmlns          string   `xml:"xmlns,attr"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Xlink          string   `xml:"xlink,attr"`
	TeiHeader      struct {
		Text     string `xml:",chardata"`
		Lang     string `xml:"lang,attr"`
		FileDesc struct {
			Text      string `xml:",chardata"`
			TitleStmt struct {
				Text  string `xml:",chardata"`
				Title struct {
					Text  string `xml:",chardata"`
					Level string `xml:"level,attr"`
					Type  string `xml:"type,attr"`
				} `xml:"title"`
			} `xml:"titleStmt"`
			PublicationStmt struct {
				Text         string `xml:",chardata"`
				Publisher    string `xml:"publisher"`
				Availability struct {
					Text    string `xml:",chardata"`
					Status  string `xml:"status,attr"`
					Licence string `xml:"licence"`
				} `xml:"availability"`
			} `xml:"publicationStmt"`
			SourceDesc struct {
				Text       string `xml:",chardata"`
				BiblStruct struct {
					Text     string `xml:",chardata"`
					Analytic struct {
						Text   string `xml:",chardata"`
						Author []struct {
							Text     string `xml:",chardata"`
							PersName struct {
								Text     string `xml:",chardata"`
								Forename struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"forename"`
								Surname string `xml:"surname"`
							} `xml:"persName"`
							Affiliation []struct {
								Text    string `xml:",chardata"`
								Key     string `xml:"key,attr"`
								OrgName []struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
									Key  string `xml:"key,attr"`
								} `xml:"orgName"`
								Address struct {
									Text       string `xml:",chardata"`
									PostCode   string `xml:"postCode"`
									Settlement string `xml:"settlement"`
									Country    struct {
										Text string `xml:",chardata"`
										Key  string `xml:"key,attr"`
									} `xml:"country"`
								} `xml:"address"`
							} `xml:"affiliation"`
						} `xml:"author"`
						Title struct {
							Text  string `xml:",chardata"`
							Level string `xml:"level,attr"`
							Type  string `xml:"type,attr"`
						} `xml:"title"`
					} `xml:"analytic"`
					Monogr struct {
						Text    string `xml:",chardata"`
						Imprint struct {
							Text string `xml:",chardata"`
							Date string `xml:"date"`
						} `xml:"imprint"`
					} `xml:"monogr"`
					Idno []struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"idno"`
					Note struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"note"`
				} `xml:"biblStruct"`
			} `xml:"sourceDesc"`
		} `xml:"fileDesc"`
		EncodingDesc struct {
			Text    string `xml:",chardata"`
			AppInfo struct {
				Text        string `xml:",chardata"`
				Application struct {
					Text    string `xml:",chardata"`
					Version string `xml:"version,attr"`
					Ident   string `xml:"ident,attr"`
					When    string `xml:"when,attr"`
					Desc    string `xml:"desc"`
					Ref     struct {
						Text   string `xml:",chardata"`
						Target string `xml:"target,attr"`
					} `xml:"ref"`
				} `xml:"application"`
			} `xml:"appInfo"`
		} `xml:"encodingDesc"`
		ProfileDesc struct {
			Text      string `xml:",chardata"`
			TextClass struct {
				Text     string `xml:",chardata"`
				Keywords struct {
					Text string   `xml:",chardata"`
					Term []string `xml:"term"`
				} `xml:"keywords"`
			} `xml:"textClass"`
			Abstract struct {
				Text string `xml:",chardata"`
				Div  struct {
					Text  string `xml:",chardata"`
					Xmlns string `xml:"xmlns,attr"`
					P     string `xml:"p"`
				} `xml:"div"`
			} `xml:"abstract"`
		} `xml:"profileDesc"`
	} `xml:"teiHeader"`
	Text struct {
		Text string `xml:",chardata"`
		Lang string `xml:"lang,attr"`
		Body struct {
			Text string `xml:",chardata"`
			Div  []struct {
				Text  string `xml:",chardata"`
				Xmlns string `xml:"xmlns,attr"`
				Head  struct {
					Text string `xml:",chardata"`
					N    string `xml:"n,attr"`
				} `xml:"head"`
				P []struct {
					Text string `xml:",chardata"`
					Ref  []struct {
						Text   string `xml:",chardata"`
						Type   string `xml:"type,attr"`
						Target string `xml:"target,attr"`
					} `xml:"ref"`
				} `xml:"p"`
				Formula []struct {
					Text  string `xml:",chardata"`
					ID    string `xml:"id,attr"`
					Label string `xml:"label"`
				} `xml:"formula"`
			} `xml:"div"`
			Figure []struct {
				Text    string `xml:",chardata"`
				Xmlns   string `xml:"xmlns,attr"`
				ID      string `xml:"id,attr"`
				Type    string `xml:"type,attr"`
				Head    string `xml:"head"`
				Label   string `xml:"label"`
				FigDesc string `xml:"figDesc"`
				Graphic struct {
					Text   string `xml:",chardata"`
					Coords string `xml:"coords,attr"`
					Type   string `xml:"type,attr"`
				} `xml:"graphic"`
				Table struct {
					Text string `xml:",chardata"`
					Row  []struct {
						Text string `xml:",chardata"`
						Cell []struct {
							Text string `xml:",chardata"`
							Cols string `xml:"cols,attr"`
						} `xml:"cell"`
					} `xml:"row"`
				} `xml:"table"`
				Note string `xml:"note"`
			} `xml:"figure"`
			Note struct {
				Text  string `xml:",chardata"`
				Xmlns string `xml:"xmlns,attr"`
				Place string `xml:"place,attr"`
			} `xml:"note"`
		} `xml:"body"`
		Back struct {
			Text string `xml:",chardata"`
			Div  []struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
				Div  struct {
					Text  string   `xml:",chardata"`
					Xmlns string   `xml:"xmlns,attr"`
					P     []string `xml:"p"`
				} `xml:"div"`
				ListBibl struct {
					Text       string `xml:",chardata"`
					BiblStruct []struct {
						Text   string `xml:",chardata"`
						ID     string `xml:"id,attr"`
						Monogr struct {
							Text  string `xml:",chardata"`
							Title struct {
								Text  string `xml:",chardata"`
								Level string `xml:"level,attr"`
								Type  string `xml:"type,attr"`
							} `xml:"title"`
							Author []struct {
								Text     string `xml:",chardata"`
								PersName struct {
									Text     string `xml:",chardata"`
									Forename []struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr"`
									} `xml:"forename"`
									Surname string `xml:"surname"`
								} `xml:"persName"`
							} `xml:"author"`
							Imprint struct {
								Text string `xml:",chardata"`
								Date struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
									When string `xml:"when,attr"`
								} `xml:"date"`
								BiblScope []struct {
									Text string `xml:",chardata"`
									Unit string `xml:"unit,attr"`
									From string `xml:"from,attr"`
									To   string `xml:"to,attr"`
								} `xml:"biblScope"`
								Publisher string `xml:"publisher"`
								PubPlace  string `xml:"pubPlace"`
							} `xml:"imprint"`
							Idno struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"idno"`
							Meeting  string `xml:"meeting"`
							RespStmt struct {
								Text    string `xml:",chardata"`
								OrgName string `xml:"orgName"`
							} `xml:"respStmt"`
						} `xml:"monogr"`
						Analytic struct {
							Text  string `xml:",chardata"`
							Title struct {
								Text  string `xml:",chardata"`
								Level string `xml:"level,attr"`
								Type  string `xml:"type,attr"`
							} `xml:"title"`
							Author []struct {
								Text     string `xml:",chardata"`
								PersName struct {
									Text     string `xml:",chardata"`
									Forename []struct {
										Text string `xml:",chardata"`
										Type string `xml:"type,attr"`
									} `xml:"forename"`
									Surname string `xml:"surname"`
								} `xml:"persName"`
							} `xml:"author"`
							Idno struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
							} `xml:"idno"`
							Ptr struct {
								Text   string `xml:",chardata"`
								Target string `xml:"target,attr"`
							} `xml:"ptr"`
						} `xml:"analytic"`
						Note struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"note"`
					} `xml:"biblStruct"`
				} `xml:"listBibl"`
			} `xml:"div"`
		} `xml:"back"`
	} `xml:"text"`
}
