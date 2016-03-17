package main 

/*
	Doc here : https://helpcenter.veeam.com/backup/rest/overview.html
*/

import (
	"log"
	"io/ioutil"
	"encoding/xml"
	"bytes"
	"os"
	"strings"
	"fmt"
	"regexp"
	"flag"
)

/*
	General types
*/
type Element struct {
	Name string
	Type string
	Multiple bool
	IsSimple bool
	Choice *Choice
}
type Attribute struct {
	Name string
	Type string

}

type ComplexType struct {
	Name string
	Extensions []*ComplexType
	ExtensionsName []string
	Elements []*Element
	Attributes []*Attribute
}
type SimpleType struct {
	Name string
}

/*
	Some generic functions for helping
*/
//find attribute based on name
func getAttribute(name string,attrs []xml.Attr) (string) {
	val := ""
	for _,attr := range attrs { if strings.EqualFold(attr.Name.Local,name) {val = attr.Value }  }
	return val
}
//check if we can convert simple type something goes knows
//if not we just convert to string, up to the user to do something useful with it
func convertType(name string) (string,bool,string) {
	conversion := "string"
	directConversion := false
	printfstr := "%s"
	
	switch name {
			case "int":
				conversion = "int"
				directConversion = true
				printfstr = "%d"
			case "long":
				conversion = "int64"
				directConversion = true
				printfstr = "%d"
			case "ulong":
				conversion = "uint64"
				directConversion = true
				printfstr = "%d"
			case "ushort":
				conversion = "uint16"
				directConversion = true
				printfstr = "%d"
			case "bool":
				conversion = "bool"
				directConversion = true
				printfstr = "%t"
			case "string":
				conversion = "string"
				directConversion = true
				printfstr = "%s"
			case "TimeSpan":
				conversion = "int"
				directConversion = true
				printfstr = "%d"
			break
	}
	return conversion,directConversion,printfstr
}



/*
	First stage parse the XSD
*/

type ResultObject struct {
	complexTypes []*ComplexType
	simpleTypes []*SimpleType
	rootElements []*Element
	nameSpace string
}

type Choice struct {
	min string
	max string
	block bool
}
func parseXSD(filename *string) (*ResultObject)  {
	r := ResultObject{}
	complexStack := []*ComplexType{}

	log.Printf("Generating from %s",*filename)
	dat, _ := ioutil.ReadFile(*filename)
	bytereader := bytes.NewReader(dat)
	decoder := xml.NewDecoder(bytereader)

	
	 
	/*
		Parsing XSD first stage
	*/
	
	depth := 0
	choiceStack := []*Choice{}
	
	for t, _ := decoder.Token(); t != nil; t,_ = decoder.Token() {
		switch se := t.(type) { 
			case xml.StartElement: 
				depth++
				
				switch (se.Name.Local) {
					case "element": {
						elname := getAttribute("name",se.Attr)
						eltype := getAttribute("type",se.Attr)
						
						minOccurs := getAttribute("minOccurs",se.Attr)
						maxOccurs := getAttribute("maxOccurs",se.Attr)
						
						if len(complexStack) > 0 {
							complexNow := complexStack[len(complexStack)-1]
							
							var elchoice *Choice
							elmultiple := false
							
							if maxOccurs != "" {
								if maxOccurs == "unbounded" {
									elchoice = &Choice{min:minOccurs,max:maxOccurs,block:false}
								} else if maxOccurs == "1" {
									elchoice = &Choice{min:minOccurs,max:"1",block:false}
								}
								
								if maxOccurs != "1" {
									elmultiple = true
								}
							} else {
								
								if len(choiceStack)>0 && !(choiceStack[len(choiceStack)-1].block) {
									elchoice = &Choice{min:choiceStack[len(choiceStack)-1].min,max:choiceStack[len(choiceStack)-1].max,block:false}
									if elchoice.max != "1" {
										elmultiple = true
									} 
								} 
							}
							
							complexNow.Elements = append(complexNow.Elements,&Element{Name:elname,Type:eltype,Multiple:elmultiple,Choice:elchoice})
							

						} else {
							r.rootElements = append(r.rootElements,&Element{Name:elname,Type:eltype})
						}
						choiceStack = append(choiceStack,&Choice{min:"",max:"",block:true})
						fmt.Printf("\nStacking %d %v",len(choiceStack),choiceStack[len(choiceStack)-1])
					} 
					break;
					case "attribute": {
						elname := getAttribute("name",se.Attr)
						eltype := getAttribute("type",se.Attr)
						
						if len(complexStack) > 0 {
							complexNow := complexStack[len(complexStack)-1]

							complexNow.Attributes = append(complexNow.Attributes,&Attribute{Name:elname,Type:eltype})
							
						} else {
							log.Printf("Can not handle root attribute")
						}
					}
					case  "complexType": {
						complexName := getAttribute("name",se.Attr)
						if complexName == "" {
							//nested complex ugh
							if len(complexStack) > 0 {
								complexNow := complexStack[len(complexStack)-1]
								if len(complexNow.Elements) > 0 {
									complexName := fmt.Sprintf("%sNested%s",complexNow.Name,complexNow.Elements[len(complexNow.Elements)-1].Name)
									complexNow.Elements[len(complexNow.Elements)-1].Type = complexName
									
									complexStack = append(complexStack,&ComplexType{Name:complexName})
									r.complexTypes = append(r.complexTypes,complexStack[len(complexStack)-1])
									
								} else {
									log.Printf("Need element name to derive new class name")
								}
							} else {
								log.Printf("Can not handle anonymous complextype without parent, sorry, .. , to complex, .. , master didn't teach me this trick")
							}
							
						} else {
							complexStack = append(complexStack,&ComplexType{Name:complexName})
							r.complexTypes = append(r.complexTypes,complexStack[len(complexStack)-1])
						}
						fmt.Printf("\nComplexName "+complexName)
						
					} 
					break;
					case "extension": {
						complexNow := complexStack[len(complexStack)-1]
						complexNow.ExtensionsName = append(complexNow.ExtensionsName,getAttribute("base",se.Attr))
					} 
					break;
					case "choice":  {
						max := getAttribute("maxOccurs",se.Attr)
						min := getAttribute("minOccurs",se.Attr)
					    if max == "unbounded" {
							choiceStack = append(choiceStack,&Choice{min:min,max:max,block:false})
						} else if max == "1" {
							choiceStack = append(choiceStack,&Choice{min:min,max:"1",block:false})
						} else {
							choiceStack = append(choiceStack,&Choice{min:"",max:"",block:true})
						}
						fmt.Printf("\nStacking %d %v",len(choiceStack),choiceStack[len(choiceStack)-1])
					} 
					break;
					case "simpleType": {
						elname := getAttribute("name",se.Attr)
						r.simpleTypes = append(r.simpleTypes,&SimpleType{elname})
					} 
					break;
					case "schema":
						r.nameSpace = getAttribute("xmlns",se.Attr)
						log.Printf("Namespace %s",r.nameSpace)
					break
					case "enumeration":
					case "pattern":
					case "sequence":
						max := getAttribute("maxOccurs",se.Attr)
						min := getAttribute("minOccurs",se.Attr)
					    if max == "unbounded" {
							choiceStack = append(choiceStack,&Choice{min:min,max:max,block:false})
						} else if max == "1" {
							choiceStack = append(choiceStack,&Choice{min:min,max:"1",block:false})
						} else {
							choiceStack = append(choiceStack,&Choice{min:"",max:"",block:true})
						}		
						fmt.Printf("\nStacking seq %d %v",len(choiceStack),choiceStack[len(choiceStack)-1])
					break
					case "complexContent":
					case "simpleContent":
					case "restriction":
					break;
					default:
					log.Printf("Unknown %s at %d",se.Name.Local,depth)
					break;
				}
				break;
			case xml.EndElement:
				if se.Name.Local == "complexType" && len(complexStack) > 0 {
				  complexStack = complexStack[:len(complexStack)-1]
				  
				}
				if len(choiceStack) > 0 && ( se.Name.Local == "choice" || se.Name.Local == "element" || se.Name.Local == "sequence") {
					choiceStack = choiceStack[:len(choiceStack)-1]
					
					fmt.Printf("\nDestacking %d",len(choiceStack))
				}
				
				depth--
				break;
		}
	}
	 
	 
	 
	 return &r
}
func inheritanceLinking(result *ResultObject) {
		/*
			Inheritance linking
		*/
		commonCacheTypes := []*ComplexType{}
		
		for _,c := range result.complexTypes {
			if len(c.ExtensionsName) > 0 {
				for _,e := range c.ExtensionsName {
					found := false
					for _,potentialparent := range commonCacheTypes {
						if strings.EqualFold(potentialparent.Name,e){
							found = true
							c.Extensions = append(c.Extensions,potentialparent)
							//log.Printf("Cache hit")
						}
					}
					if !found {
						for _,potentialparent := range result.complexTypes {
							if strings.EqualFold(potentialparent.Name,e){
								found = true
								c.Extensions = append(c.Extensions,potentialparent)
								commonCacheTypes = append(commonCacheTypes,potentialparent)
								//log.Printf("Deep hit")
							}
						}
						if !found {
							log.Printf("Could not find extension:%s for %s",c.ExtensionsName,c.Name)
						}
					}
					
				}
			}
		}
}
func simpleCalc(result *ResultObject) {
	for _,c := range result.complexTypes {
		for _,e := range c.Elements {
			IsSimple := false
			for _,s := range result.simpleTypes {
				if strings.EqualFold(e.Type,s.Name) {
					IsSimple = true
				}
			}
			if IsSimple {
				//fmt.Println(e.Type)
			}
			e.IsSimple = IsSimple
		}
	}
	
	for _,e := range result.rootElements {
		IsSimple := false
		for _,s := range result.simpleTypes {
			if strings.EqualFold(e.Type,s.Name) {
				IsSimple = true
			}
		}
		e.IsSimple = IsSimple
	}
}

//writepackage header
func writePackage(file * os.File) {
	fmt.Fprintf(file,"package veeamrestapi\r\n\r\n")
}

func writeGetters(result *ResultObject,filename string) {
		f, e := os.Create(filename)
		if e == nil {
			
			defer f.Close()
			
			relisttype := regexp.MustCompile("^(.*?)ListType$")
			retype := regexp.MustCompile("^(.*?)Type$")	
			
			writePackage(f)
			
			importv := "\r\nimport (\r\n\"fmt\"\r\n\"encoding/xml\"\r\n)\r\n\r\n"

			staticlist := []string{"EntityReferenceType","EntityReferenceListType","QuerySvcType","QueryResultType","LookupSvcType","ReportingSvcType","SummaryReportType","OverviewReportFrameType","VmsOverviewReportFrameType","JobStatisticsReportFrameType","RepositoryReportFrameType","ProcessedVmsReportFrameType","ReplicaJobSessionEntityType","ReplicaJobSessionEntityListType","BackupJobSessionEntityType","BackupJobSessionEntityListType"}
			
			fmt.Fprint(f,importv)
			for _,rootElement := range result.rootElements {
				n := rootElement.Name
				t := rootElement.Type
				
				isstatic := false
				for _,static := range staticlist {
					if strings.EqualFold(t,static) {
						isstatic = true
						fmt.Println("Ignoring "+static)
					}
				}
				
				if !isstatic {
					islist := relisttype.MatchString(t)
					//fmt.Println(t)
					singleresourcequery := n
					if !islist {
						if t := retype.FindStringSubmatch(t);len(t) > 1 {
							listtypename := fmt.Sprintf("%sListType",t[1])
							for _,testm := range result.rootElements {
								if strings.EqualFold(testm.Type,listtypename) {
									singleresourcequery = testm.Name	
								}
							}
						} 
					}
					
					/*
					func (v * VeeamRestServer) GetJobs() (JobEntityListType,error) { 
					  var returnerr error 
					  Jobs := JobEntityListType{}
					  
					  if (v.SessionId != "") {
							vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","Jobs")),"GET")
							xmlin, err := v.Request(vrr)
							if err == nil {	
								err = xml.Unmarshal([]byte(xmlin),&Jobs)
								if err != nil {
									returnerr = &VeeamRestError{"Unmarshal error jobs",err.Error()}
								}
							} else {
								returnerr = &VeeamRestError{"Invalid request asking jobs",err.Error()}
							}
					  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
					  
					  return Jobs,returnerr
					}*/
								
					if !islist {
						fmt.Fprintf(f,"/*\r\n * %s \r\n * Not validated \r\n */\r\nfunc (v * VeeamRestServer) Get%s(idstring string) (%s,error) { ",n,n,t)
					} else {
						fmt.Fprintf(f,"/*\r\n * %s \r\n * Not validated \r\n */\r\nfunc (v * VeeamRestServer) Get%s() (%s,error) { ",n,n,t)
					}
					fmt.Fprintf(f,"\r\n  var returnerr error ")
					fmt.Fprintf(f,"\r\n  %s := %s{} ",n,t)
					fmt.Fprintf(f,"\r\n  if (v.SessionId != \"\") { ")
					if !islist {
						fmt.Fprintf(f,"\r\n     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf(\"%%s/%%s?format=Entity\",\"%s\",idstring)),\"GET\") ",singleresourcequery)
					} else {
						fmt.Fprintf(f,"\r\n     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf(\"%%s?format=Entity\",\"%s\")),\"GET\") ",n)
					}
					fmt.Fprintf(f,"\r\n     xmlin, err := v.Request(vrr)")
					fmt.Fprintf(f,"\r\n     if err == nil {	")
					fmt.Fprintf(f,"\r\n     	err = xml.Unmarshal([]byte(xmlin),&%s)",n)
					fmt.Fprintf(f,"\r\n     	if err != nil {",)
					fmt.Fprintf(f,"\r\n     	    returnerr = &VeeamRestError{\"Unmarshal error %s\",err.Error()}",n)
					fmt.Fprintf(f,"\r\n     	}",)
					fmt.Fprintf(f,"\r\n     } else {	")
					fmt.Fprintf(f,"\r\n        returnerr = &VeeamRestError{\"Invalid request asking %s\",err.Error()}",n)
					fmt.Fprintf(f,"\r\n     }")
					fmt.Fprintf(f,"\r\n  } else { returnerr = &VeeamRestError{\"No logon session set, did you login?\",\"\"} }")
					fmt.Fprintf(f,"\r\n  return %s,returnerr",n)
					fmt.Fprintf(f,"\r\n}\r\n\r\n\r\n")
				}
				
			}
					
		} else {
			log.Printf("FS error, could not write to %s",filename)
		}
}

type ImplementInterface struct {
	Links bool
	Spec bool
}

func writeComplexType(f * os.File,complexType *ComplexType,checkinterfaces *ImplementInterface) {

			for _,element := range complexType.Elements {
				XMLstr := fmt.Sprintf("xml:\"%s\"",element.Name)
				
				if element.Choice != nil && element.Choice.max == "1"   {
					//fmt.Println(element.Name+" "+string(element.Choice.max))
					XMLstr = fmt.Sprintf("xml:\"%s,omitempty\"",element.Name)
				}
				
				typestr := element.Type
				if !element.IsSimple {
					typestr = "*"+typestr
				} else {
					converttype,direct,_ := convertType(typestr)
					if direct { typestr = converttype}
				}
				
				if element.Multiple {
					typestr = "[]"+typestr
				}
				
				fmt.Fprintf(f,"\r\n   %s %s `%s`",element.Name,typestr,XMLstr)
				
				if element.Type == "LinkListType" {
					checkinterfaces.Links = true
				}
			}
			for _,attr := range complexType.Attributes {
				XMLstr := fmt.Sprintf("xml:\"%s,attr\"",attr.Name)
				
				typestr := attr.Type
				converttype,direct,_ := convertType(typestr)
				if direct { typestr = converttype}
				
				fmt.Fprintf(f,"\r\n   %s %s `%s`",attr.Name,typestr,XMLstr)
			}
			for _,ext := range complexType.Extensions {
				fmt.Fprintf(f,"\r\n   //Inhereting from %s",ext.Name)
				switch ext.Name {
					case "SpecType":
					 checkinterfaces.Spec = true
					break
				}
				writeComplexType(f,ext,checkinterfaces)
			}
}

func writeComplexTypes(result *ResultObject,filename string) {
		f, e := os.Create(filename)
		if e == nil {
			defer f.Close()
			writePackage(f)
			fmt.Fprintf(f,"\r\n \r\nimport (\r\n \"encoding/xml\"\r\n )\r\n \r\n ")
			for _,complexType := range result.complexTypes {
				checkinterfaces := &ImplementInterface{}
			
				fmt.Fprintf(f,"/*\r\n * %s \r\n * Not validated \r\n */\r\ntype %s struct { ",complexType.Name,complexType.Name)
				//more important during creating new objects at client side
				fmt.Fprintf(f,"\r\n   XMLName xml.Name")
				writeComplexType(f,complexType,checkinterfaces)
				fmt.Fprintf(f,"\r\n}")
				
				if checkinterfaces.Links {
					fmt.Fprintf(f,"\r\nfunc (l %s) GetLinks() (*LinkListType) { return (l.Links) }",complexType.Name)
				}
				if checkinterfaces.Spec {
					fmt.Fprintf(f,"\r\nfunc New%s() (*%s) {",complexType.Name,complexType.Name)
					fmt.Fprintf(f,"\r\n  var%s := %s{}",complexType.Name,complexType.Name)
					fmt.Fprintf(f,"\r\n  var%s.XMLName.Local = \"%s\"",complexType.Name,strings.TrimRight(complexType.Name,"Type"))
					fmt.Fprintf(f,"\r\n  var%s.XMLName.Space = \"%s\"",complexType.Name,result.nameSpace)
					fmt.Fprintf(f,"\r\n  return &var%s",complexType.Name)
					fmt.Fprintf(f,"\r\n}",)
				}
				fmt.Fprintf(f,"\r\n\r\n")
			}
			
		} else {
			log.Printf("FS error, could not write to %s",filename)
		}
}
//write simpletypes, if we can not convert make a type that basically overwrites string
func writeSimpleTypes(result *ResultObject,filename string) {
		f, e := os.Create(filename)
		if e == nil {
			defer f.Close()
			writePackage(f)
			for _,simpleType := range result.simpleTypes {
				_,direct,_ := convertType(simpleType.Name)
				if !direct {
					fmt.Fprintf(f,"/* Unconvertable simple type * %s */",simpleType.Name)
					fmt.Fprintf(f,"\r\ntype %s string",simpleType.Name)
					fmt.Fprintf(f,"\r\nfunc (o * %s) String() string { return string(*o) }",simpleType.Name)
					fmt.Fprintf(f,"\r\nfunc (o * %s) S() string { return string(*o) }",simpleType.Name)
					fmt.Fprintf(f,"\r\n\r\n\r\n")
				}
			}
		} else {
			log.Printf("FS error, could not write to %s",filename)
		}
}
func writeStringer(f * os.File,complexType * ComplexType) {

			for _,element := range complexType.Elements {
				
				if element.Multiple {
					fmt.Fprintf(f,"\r\n  if full {")
					fmt.Fprintf(f,"\r\n   for i,c := range o.%s {",element.Name)
					fmt.Fprintf(f,"\r\n    buffer.Write(identbuffer.Bytes())")
					fmt.Fprintf(f,"\r\n    buffer.WriteString(fmt.Sprintf(\"%s[%%d]\",i))",element.Name)
					
					if element.IsSimple {
						_,direct,printfstr := convertType(element.Type)
						fmt.Fprintf(f,"\r\n   buffer.Write(identbuffer.Bytes())")
						if !direct {
							fmt.Fprintf(f,"\r\n    buffer.WriteString(\"%-30s : \");buffer.WriteString(c.String())",element.Name)
						}  else {
							fmt.Fprintf(f,"\r\n    buffer.WriteString(\"%-30s : \");buffer.WriteString(fmt.Sprintf(\"%s\",c))",element.Name,printfstr)
						}
					} else {
						fmt.Fprintf(f,"\r\n    buffer.WriteString(c.FullString(full,depth+1))")
					}
					fmt.Fprintf(f,"\r\n   }\r\n  }")
				} else {
					if element.IsSimple {
						fmt.Fprintf(f,"\r\n   buffer.Write(identbuffer.Bytes())")
						_,direct,printfstr := convertType(element.Type)
						if !direct {
							
							fmt.Fprintf(f,"\r\n   buffer.WriteString(\"%-30s : \");buffer.WriteString(o.%s.String())",element.Name,element.Name)
						}  else {

							fmt.Fprintf(f,"\r\n   buffer.WriteString(\"%-30s : \");buffer.WriteString(fmt.Sprintf(\"%s\",o.%s))",element.Name,printfstr,element.Name)
						}
					} else {
						fmt.Fprintf(f,"\r\n  if full && o.%s != nil {",element.Name)
						fmt.Fprintf(f,"\r\n   buffer.Write(identbuffer.Bytes())")
						fmt.Fprintf(f,"\r\n   buffer.WriteString(\"%-30s : \");buffer.WriteString(o.%s.FullString(full,depth+1))",element.Name,element.Name)
						fmt.Fprintf(f,"\r\n  }")
					}
				}
				
			}
			for _,attr := range complexType.Attributes {			
				_,direct,printfstr := convertType(attr.Type)
				fmt.Fprintf(f,"\r\n  buffer.Write(identbuffer.Bytes())")
				if !direct {
						fmt.Fprintf(f,"\r\n  buffer.WriteString(\"%-30s : \");buffer.WriteString(o.%s.String())",attr.Name,attr.Name)
				} else {
						fmt.Fprintf(f,"\r\n  buffer.WriteString(\"%-30s : \");buffer.WriteString(fmt.Sprintf(\"%s\",o.%s))",attr.Name,printfstr,attr.Name)
				}
			}
			for _,ext := range complexType.Extensions {
				fmt.Fprintf(f,"\r\n   //Inhereting from %s",ext.Name)
				writeStringer(f,ext)
			}
}
func writeStringers(result *ResultObject,filename string) {
		f, e := os.Create(filename)
		if e == nil {
			defer f.Close()
			writePackage(f)
			fmt.Fprintf(f,"\r\n \r\nimport (\r\n \"bytes\"\r\n \"fmt\"\r\n )\r\n \r\n")
			
			for _,complexType := range result.complexTypes {
				fmt.Fprintf(f,"/*\r\n * %s \r\n * Not validated \r\n */\r\nfunc (o %s) FullString(full bool,depth int) (string) { ",complexType.Name,complexType.Name)
				fmt.Fprintf(f,"\r\n  var buffer bytes.Buffer")
				fmt.Fprintf(f,"\r\n  var identbuffer bytes.Buffer")
				fmt.Fprintf(f,"\r\n  identbuffer.WriteString(\"\\r\\n\");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }")
				writeStringer(f,complexType)
				fmt.Fprintf(f,"\r\n  return buffer.String()")
				fmt.Fprintf(f,"\r\n}")
				fmt.Fprintf(f,"\r\nfunc (o %s) String() (string) { return o.FullString(false,0) }",complexType.Name)
				fmt.Fprintf(f,"\r\n\r\n")
			}
		} else {
			log.Printf("FS error, could not write to %s",filename)
		}
}

//start the fun
func main() {
	xsdfilep := flag.String("x", `C:\Program Files\Veeam\Backup and Replication\Enterprise Manager\schemas\RestAPI.xsd`, "RestAPI.xsd, should be fine if you run it on the enterprise manager")
	objectfilep := flag.String("o","veeamrest_object.go","object file")
	getfilep := flag.String("g","veeamrest_get.go","getters file")
	simplefilep := flag.String("s","veeamrest_simple.go","simpletypeconversion file")
	stringfilep := flag.String("p","veeamrest_string.go","string conversion")
	
	flag.Parse()
	
	result := parseXSD(xsdfilep)
	inheritanceLinking(result)
	simpleCalc(result)

	
	writeSimpleTypes(result,*simplefilep)
	writeComplexTypes(result,*objectfilep)
	writeGetters(result,*getfilep)
	writeStringers(result,*stringfilep)
	

}