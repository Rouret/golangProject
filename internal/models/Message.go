package models

type Message struct {
	IdCapteur int     `json:idCapteur`
	IATA      string  `json:iata`
	TypeValue string  `json:typeValue`
	Value     float32 `json:value`
	Timestamp int64   `json:timestamp`
}

type Messages []Message