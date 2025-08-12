package quote

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Picker handles quote selection and filtering
type Picker struct {
	quotes []Quote
}

// NewPicker creates a new picker instance
func NewPicker(quotes []Quote) *Picker {
	return &Picker{quotes: quotes}
}

// PickRandom selects a random quote
func (p *Picker) PickRandom() (*Quote, error) {
	if len(p.quotes) == 0 {
		return nil, fmt.Errorf("no quotes available")
	}

	index := rand.Intn(len(p.quotes))
	return &p.quotes[index], nil
}

// PickDaily selects a quote based on the current date (same quote for the same day)
func (p *Picker) PickDaily() (*Quote, error) {
	if len(p.quotes) == 0 {
		return nil, fmt.Errorf("no quotes available")
	}

	// Use current date as seed for consistent daily selection
	now := time.Now()
	seed := now.Year()*10000 + int(now.Month())*100 + now.Day()
	rand.Seed(int64(seed))

	index := rand.Intn(len(p.quotes))
	return &p.quotes[index], nil
}

// FilterByTags filters quotes by specified tags
func (p *Picker) FilterByTags(tags string) *Picker {
	if tags == "" {
		return p
	}

	tagList := strings.Split(tags, ",")
	for i, tag := range tagList {
		tagList[i] = strings.TrimSpace(tag)
	}

	var filtered []Quote
	for _, quote := range p.quotes {
		if p.hasAnyTag(quote, tagList) {
			filtered = append(filtered, quote)
		}
	}

	return &Picker{quotes: filtered}
}

// FilterByLang filters quotes by language
func (p *Picker) FilterByLang(lang string) *Picker {
	if lang == "" {
		return p
	}

	var filtered []Quote
	for _, quote := range p.quotes {
		if strings.EqualFold(quote.Lang, lang) {
			filtered = append(filtered, quote)
		}
	}

	return &Picker{quotes: filtered}
}

// hasAnyTag checks if a quote has any of the specified tags
func (p *Picker) hasAnyTag(quote Quote, tags []string) bool {
	for _, tag := range tags {
		for _, quoteTag := range quote.Tags {
			if strings.EqualFold(quoteTag, tag) {
				return true
			}
		}
	}
	return false
}

// Count returns the number of available quotes
func (p *Picker) Count() int {
	return len(p.quotes)
}
