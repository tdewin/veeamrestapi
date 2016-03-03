package veeamrestapi

import (
	"fmt"
	"encoding/xml"
)

/*
 * EntityRef 
 * Since it is not a generic function call, it is not generated correctly
 * Extracted and modified 
 */
func (v * VeeamRestServer) GetEntityRef(resource string,idstring string) (EntityReferenceType,error) { 
  var returnerr error 
  EntityRef := EntityReferenceType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s",resource,idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&EntityRef)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error EntityRef",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking EntityRef",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return EntityRef,returnerr
}



/*
 * EntityReferences 
 * Since it is not a generic function call, it is not generated correctly
 * Extracted and modified 
 */
func (v * VeeamRestServer) GetEntityReferences(resource string) (EntityReferenceListType,error) { 
  var returnerr error 
  EntityReferences := EntityReferenceListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s",resource)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&EntityReferences)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error EntityReferences",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking EntityReferences",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return EntityReferences,returnerr
}
