package language

import "testing"

func TestEnglish(t *testing.T) {
	except := "English"
	actual := English()

	if except != actual {
		t.Errorf("%s != %s", except, actual)
	}
}

func TestJapanese(t *testing.T) {
	except := "Japanese"
	actual := Japanese()

	if except != actual {
		t.Errorf("%s != %s", except, actual)
	}
}
