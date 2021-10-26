package sensor

type Message struct {
	IdCapteur int
	IATA      string
	TypeValue string
	Value     float32
	Timestamp int64
}