package fix

import (
	"testing"

	"github.com/quickfixgo/quickfix/tag"
)

func TestNewOrderSingle(t *testing.T) {
	var tests = []struct {
		clOrdID      string
		symbol       string
		side         rune
		orderQty     float64
		orderType    rune
		expectedType string
	}{
		{"123456", "AAPL", '1', 100, '2', "D"}, // '1' for Buy side and '2' for Limit Order Type
	}

	for _, tt := range tests {
		t.Run("NewOrderSingle", func(t *testing.T) {
			msg := NewOrderSingle(tt.clOrdID, tt.symbol, tt.side, tt.orderQty, tt.orderType)
			msgType, err := msg.Header.GetString(tag.MsgType)
			if err != nil {
				t.Error("Header does not have MsgType")
			}
			if msgType != tt.expectedType {
				t.Errorf("Expected MsgType %s, got %s", tt.expectedType, msgType)
			}
		})
	}
}

func TestParseMessage(t *testing.T) {
	msg := NewOrderSingle("123456", "AAPL", '1', 100, '2')
	rawMsg := msg.String()

	var tests = []struct {
		description string
		input       string
		shouldError bool
	}{
		{"ValidMessage", rawMsg, false},
		{"InvalidMessage", "8=FIX.4.4\x019=70\x0135=D\x0134=1\x01", true},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			_, err := ParseMessage(test.input)
			if (err != nil) != test.shouldError {
				t.Errorf("ParseMessage() expected error: %v, got error: %v", test.shouldError, err)
			}
		})
	}
}
