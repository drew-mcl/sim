package fix

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix44/newordersingle"
)

// NewOrderSingle creates a FIX 4.4 New Order Single message.
func NewOrderSingle(clOrdID, symbol string, side rune, orderQty float64, orderType rune) *quickfix.Message {
	msg := newordersingle.New(field.NewClOrdID(clOrdID), field.NewSymbol(symbol), field.NewSide(side), field.NewTransactTime(), field.NewOrdType(orderType))
	msg.Body.Set(field.NewOrderQty(orderQty, 2)) // Setting with precision for Qty
	return msg.ToMessage()
}

// ParseMessage parses a FIX string into a Message object.
func ParseMessage(data string) (*quickfix.Message, quickfix.MessageRejectError) {
	msg := quickfix.NewMessage()
	err := quickfix.ParseMessage(msg, []byte(data))
	return msg, err
}
