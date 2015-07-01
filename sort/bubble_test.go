package sort

import (
	"github.com/Congenital/log/v0.2/log"
	"testing"
)

func TestBubble(t *testing.T) {
	var value []interface{} = []interface{}{321, 32, 42, 1312, 321}

	Bubble(value, func(va, vb interface{}) bool {
		a, ok := va.(int)
		if !ok {
			log.Info(a)
			return false
		}

		b, ok := vb.(int)
		if !ok {
			log.Info(b)
			return false
		}

		if a > b {
			log.Info(a, b)
			return true
		}

		log.Info(a, b)
		return false
	})

	log.Info(value)

	Bubble(value, func(va, vb interface{}) bool {
		a, ok := va.(int)
		if !ok {
			log.Info(a)
			return false
		}

		b, ok := vb.(int)
		if !ok {
			log.Info(b)
			return false
		}

		if a < b {
			log.Info(a, b)
			return true
		}

		log.Info(a, b)
		return false
	})

	log.Info(value)
}
