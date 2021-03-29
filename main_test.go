package mapper

import (
	"testing"
)

func TestCapitalizeEveryThirdAlphanumericChar(t *testing.T) {
	t.Run("TestAspiration", func(t *testing.T) {
		expected := "asPirAtiOn.cOm"
		input := "Aspiration.com"
		output := CapitalizeEveryThirdAlphanumericChar(input)
		if expected != output {
			t.Errorf("Expected %s got %s", expected, output)
		}

	})

	t.Run("TestAspiration", func(t *testing.T) {
		expected := "abX@axA"
		input := "Abx@Axa"
		output := CapitalizeEveryThirdAlphanumericChar(input)
		if expected != output {
			t.Errorf("Expected %s got %s", expected, output)
		}

	})

}
