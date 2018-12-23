// Default matcher for searching data
package search

type defaultMatcher struct{}

// init registers the default matcher with the prigram
func init()  {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search implements the behavior for the default matcher
func(m defaultMatcher) Search(feed *Feed, searchTerm string) ([] *Result, error) {
	return nil, nil
}