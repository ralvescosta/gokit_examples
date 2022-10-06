package consumers

type (
	BasicMessage struct{}
)

func (BasicMessage) String() string {
	return "BasicMessage"
}
