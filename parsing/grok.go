package parsing

// GrokParser is a parser for less structured data using Grok patterns
type GrokParser struct {
	data *[]byte
}

// Parse parses the data using Grok patterns, returning an array of objects
func (p GrokParser) Parse(data *[]byte) (entries []map[string]interface{}, err error) {
	//TODO: implement Grok parsing
	return entries, err
}
