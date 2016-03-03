package veeamrestapi

import(
	"regexp"
)

func SplitUrn(urnstr string) []string {
	idregexp := regexp.MustCompile("^urn\\:([a-zA-Z_0-9\\-]+)\\:([a-zA-Z_0-9\\-]+)\\:([a-zA-Z_0-9\\-]+)$")
	
	splitted := []string{"","",""}
	if matches := idregexp.FindStringSubmatch(urnstr);len(matches) > 3 {
		splitted = matches[1:4]
	}
	return splitted
}

func (o * UrnType) Split() []string { return SplitUrn(string(*o))}
func (o * UidType) Split() []string { return SplitUrn(string(*o))}
func (o * HierarchyObjRefType) Split() []string { return SplitUrn(string(*o))}
