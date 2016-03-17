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
	 execurl := v.ConstructUrl("query?")+getstr
     vrr := v.MakeRequest(v.ConstructUrl(execurl),"GET") 
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
 * Lookup returns HierarchyItems 
 * https://helpcenter.veeam.com/backup/rest/get_lookupsvc.html
 * /lookup?host={hostUID}&hierachyRef={hierachyRef}&name={objName}&type={objType}
 */
func (v * VeeamRestServer) GetLookup(vals url.Values) (HierarchyItemListType,error) { 
  var returnerr error 
  HierarchyItems := HierarchyItemListType{} 
  if (v.SessionId != "") { 
	 getstr := vals.Encode()
	 execurl := v.ConstructUrl("lookup?")+getstr
     vrr := v.MakeRequest(execurl,"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&HierarchyItems)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error HierarchyItems",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking HierarchyItems",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return HierarchyItems,returnerr
}

/*
 * LookupSvc 
 * Not validated 
 */
func (v * VeeamRestServer) GetLookupSvc() (LookupSvcType,error) { 
  var returnerr error 
  LookupSvc := LookupSvcType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s","LookupSvc")),"GET") 
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



/*
 * BackupJobSession 
 * Not validated 
 */
func (v * VeeamRestServer) GetBackupJobSession(idstring string) (BackupJobSessionEntityType,error) { 
  var returnerr error 
  BackupJobSession := BackupJobSessionEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","BackupSessions",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&BackupJobSession)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error BackupJobSession",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking BackupJobSession",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return BackupJobSession,returnerr
}


/*
 * BackupJobSessions 
 * Not validated 
 */
func (v * VeeamRestServer) GetBackupJobSessions() (BackupJobSessionEntityListType,error) { 
  var returnerr error 
  BackupJobSessions := BackupJobSessionEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","BackupSessions")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&BackupJobSessions)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error BackupJobSessions",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking BackupJobSessions",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return BackupJobSessions,returnerr
}


/*
 * ReplicaJobSession 
 * Not validated 
 */
func (v * VeeamRestServer) GetReplicaJobSession(idstring string) (ReplicaJobSessionEntityType,error) { 
  var returnerr error 
  ReplicaJobSession := ReplicaJobSessionEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","ReplicaSessions",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&ReplicaJobSession)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error ReplicaJobSession",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking ReplicaJobSession",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return ReplicaJobSession,returnerr
}


/*
 * ReplicaJobSessions 
 * Not validated 
 */
func (v * VeeamRestServer) GetReplicaJobSessions() (ReplicaJobSessionEntityListType,error) { 
  var returnerr error 
  ReplicaJobSessions := ReplicaJobSessionEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","ReplicaSessions")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&ReplicaJobSessions)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error ReplicaJobSessions",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking ReplicaJobSessions",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return ReplicaJobSessions,returnerr
}

