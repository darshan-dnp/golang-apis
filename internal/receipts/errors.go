package receipts

type ReceiptExistsError struct {
	Msg string
}

func (e *ReceiptExistsError) Error() string {
	return e.Msg
}

func NewReceiptExistsError(msg string) error {
	return &ReceiptExistsError{Msg: msg}
}
