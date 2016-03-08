package veeamrestapi

import (
	"fmt"
	"encoding/xml"
	"net/url"
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




/*
 * QuerySvc 
 * https://localhost:9398/web/#/api/querySvc
 */
func (v * VeeamRestServer) GetQuerySvc() (QuerySvcType,error) { 
  var returnerr error 
  QuerySvc := QuerySvcType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s","querySvc")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&QuerySvc)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error QuerySvc",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking QuerySvc",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return QuerySvc,returnerr
}


/*
 * QuerySvc 
 * https://helpcenter.veeam.com/backup/rest/querysvc.html
 * https://localhost:9398/web/#/api/query?type=Repository&format=Entities&sortAsc=name&pageSize=15&page=1&filter=
 */
func (v * VeeamRestServer) GetQueryResult(vals url.Values) (QueryResultType,error) { 
  var returnerr error 
  QueryResult := QueryResultType{} 
  if (v.SessionId != "") { 
	 getstr := vals.Encode()
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?%s","query",getstr)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&QueryResult)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error QueryResult",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking QueryResult",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return QueryResult,returnerr
}


/*
 * LookupSvc 
 * Not validated 
 */
func (v * VeeamRestServer) GetLookupSvc(idstring string) (LookupSvcType,error) { 
  var returnerr error 
  LookupSvc := LookupSvcType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","LookupSvc",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&LookupSvc)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error LookupSvc",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking LookupSvc",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return LookupSvc,returnerr
}

/*
 * ReportingSvc 
 * Not validated 
 */
func (v * VeeamRestServer) GetReportingSvc() (ReportingSvcType,error) { 
  var returnerr error 
  ReportingSvc := ReportingSvcType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s","reports/summary")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&ReportingSvc)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error ReportingSvc",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking ReportingSvc",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return ReportingSvc,returnerr
}

