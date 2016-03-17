package veeamrestapi

import(
	"regexp"
	"time"
)

func SplitUrn(urnstr string) []string {
	idregexp := regexp.MustCompile("^urn\\:([a-zA-Z_0-9\\-]+)\\:([a-zA-Z_0-9\\-]+)\\:([a-zA-Z_0-9\\-\\.]+)$")
	
	splitted := []string{"","",""}
	if matches := idregexp.FindStringSubmatch(urnstr);len(matches) > 3 {
		splitted = matches[1:4]
	}
	return splitted
}

func (o * UrnType) Split() []string { return SplitUrn(string(*o))}
func (o * UidType) Split() []string { return SplitUrn(string(*o))}
func (o * HierarchyObjRefType) Split() []string { return SplitUrn(string(*o))}


func (o * DateTime) ParseTime() *time.Time { 
	parse,err := time.Parse(time.RFC3339,string(*o))
	if err == nil {
		return &parse
	}
	return nil
}
func (o * DateTime) TimeString() string { 
	text := ""
	if p := o.ParseTime();p!= nil {
		text = p.Format(time.RFC822)
	}
	return text
}

