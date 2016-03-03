package veeamrestapi

import(
	"strings"
)
//interface that every resource with Links should have
type Linker interface {
    GetLinks() (*LinkListType)
}



func (v * VeeamRestServer) FindLinkByRel(o Linker,rel string) ([]LinkType) {
	links := []LinkType{}

	for _,link := range o.GetLinks().Link {
		if strings.EqualFold(rel,link.Rel) {
			links = append(links,link)
		}
	}
	return links
}
func (v * VeeamRestServer) FindLinkByType(o Linker,typestr string) ([]LinkType) {
	links := []LinkType{}

	for _,link := range o.GetLinks().Link {
		if strings.EqualFold(typestr,link.Type) {
			links = append(links,link)
		}
	}
	return links
}