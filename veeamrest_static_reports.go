package veeamrestapi

import (
"fmt"
"encoding/xml"
)

/*
 * SummaryReport 
 * Not validated 
 */
func (v * VeeamRestServer) GetSummaryReport() (SummaryReportType,error) { 
  var returnerr error 
  SummaryReport := SummaryReportType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("reports/summary")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&SummaryReport)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error SummaryReport",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking SummaryReport",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return SummaryReport,returnerr
}


/*
 * OverviewReportFrame 
 * Not validated 
 */
func (v * VeeamRestServer) GetOverviewReportFrame() (OverviewReportFrameType,error) { 
  var returnerr error 
  OverviewReportFrame := OverviewReportFrameType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("reports/summary/%s","Overview")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&OverviewReportFrame)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error Overview",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking Overview",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return OverviewReportFrame,returnerr
}


/*
 * VmsOverviewReportFrame 
 * Not validated 
 */
func (v * VeeamRestServer) GetVmsOverviewReportFrame() (VmsOverviewReportFrameType,error) { 
  var returnerr error 
  VmsOverviewReportFrame := VmsOverviewReportFrameType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("reports/summary/%s","VmsOverview")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&VmsOverviewReportFrame)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error VmsOverview",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking VmsOverview",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return VmsOverviewReportFrame,returnerr
}


/*
 * JobStatisticsReportFrame 
 * Not validated 
 */
func (v * VeeamRestServer) GetJobStatisticsReportFrame() (JobStatisticsReportFrameType,error) { 
  var returnerr error 
  JobStatisticsReportFrame := JobStatisticsReportFrameType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("reports/summary/%s","JobStatistics")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&JobStatisticsReportFrame)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error JobStatistics",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking JobStatistics",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return JobStatisticsReportFrame,returnerr
}


/*
 * RepositoryReportFrame 
 * Not validated 
 */
func (v * VeeamRestServer) GetRepositoryReportFrame() (RepositoryReportFrameType,error) { 
  var returnerr error 
  RepositoryReportFrame := RepositoryReportFrameType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("reports/summary/%s","Repository")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&RepositoryReportFrame)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error Repository",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking Repository",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return RepositoryReportFrame,returnerr
}


/*
 * ProcessedVmsReportFrame 
 * Not validated 
 */
func (v * VeeamRestServer) GetProcessedVmsReportFrame() (ProcessedVmsReportFrameType,error) { 
  var returnerr error 
  ProcessedVmsReportFrame := ProcessedVmsReportFrameType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("reports/summary/%s","ProcessedVms")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&ProcessedVmsReportFrame)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error ProcessedVms",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking ProcessedVms",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return ProcessedVmsReportFrame,returnerr
}
