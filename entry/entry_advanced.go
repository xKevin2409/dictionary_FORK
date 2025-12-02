package entry

import "strings"

// FromWordPairCsv erzeugt ein neues Entry-Objekt aus einem String,
// der ein Wortpaar enthält. Das deutsche und das englische Wort sind
// durch ein Komma getrennt.
// Gibt es keines oder mehrere Kommas im String, wird ein leerer Eintrag zurückgegeben.
func FromWordPairCsv(pair string) Entry {
	parts := strings.Split(pair, ",")
	if len(parts) != 2 {
		return Empty()
	}

	de := strings.TrimSpace(parts[0])
	en := strings.TrimSpace(parts[1])

	return New(de, en)

}
