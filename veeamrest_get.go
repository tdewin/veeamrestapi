package veeamrestapi


import (
"fmt"
"encoding/xml"
)

/*
 * LogonSession 
 * Not validated 
 */
func (v * VeeamRestServer) GetLogonSession(idstring string) (LogonSessionType,error) { 
  var returnerr error 
  LogonSession := LogonSessionType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","LogonSessions",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&LogonSession)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error LogonSession",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking LogonSession",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return LogonSession,returnerr
}


/*
 * LogonSessions 
 * Not validated 
 */
func (v * VeeamRestServer) GetLogonSessions() (LogonSessionListType,error) { 
  var returnerr error 
  LogonSessions := LogonSessionListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","LogonSessions")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&LogonSessions)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error LogonSessions",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking LogonSessions",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return LogonSessions,returnerr
}


/*
 * EnterpriseManager 
 * Not validated 
 */
func (v * VeeamRestServer) GetEnterpriseManager(idstring string) (EnterpriseManagerType,error) { 
  var returnerr error 
  EnterpriseManager := EnterpriseManagerType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","EnterpriseManager",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&EnterpriseManager)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error EnterpriseManager",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking EnterpriseManager",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return EnterpriseManager,returnerr
}


/*
 * BackupServer 
 * Not validated 
 */
func (v * VeeamRestServer) GetBackupServer(idstring string) (BackupServerEntityType,error) { 
  var returnerr error 
  BackupServer := BackupServerEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","BackupServers",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&BackupServer)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error BackupServer",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking BackupServer",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return BackupServer,returnerr
}


/*
 * BackupServers 
 * Not validated 
 */
func (v * VeeamRestServer) GetBackupServers() (BackupServerEntityListType,error) { 
  var returnerr error 
  BackupServers := BackupServerEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","BackupServers")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&BackupServers)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error BackupServers",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking BackupServers",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return BackupServers,returnerr
}


/*
 * ManagedServer 
 * Not validated 
 */
func (v * VeeamRestServer) GetManagedServer(idstring string) (ManagedServerEntityType,error) { 
  var returnerr error 
  ManagedServer := ManagedServerEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","ManagedServers",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&ManagedServer)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error ManagedServer",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking ManagedServer",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return ManagedServer,returnerr
}


/*
 * ManagedServers 
 * Not validated 
 */
func (v * VeeamRestServer) GetManagedServers() (ManagedServerEntityListType,error) { 
  var returnerr error 
  ManagedServers := ManagedServerEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","ManagedServers")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&ManagedServers)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error ManagedServers",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking ManagedServers",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return ManagedServers,returnerr
}


/*
 * Job 
 * Not validated 
 */
func (v * VeeamRestServer) GetJob(idstring string) (JobEntityType,error) { 
  var returnerr error 
  Job := JobEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","Jobs",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&Job)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error Job",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking Job",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return Job,returnerr
}


/*
 * FailoverPlan 
 * Not validated 
 */
func (v * VeeamRestServer) GetFailoverPlan(idstring string) (FailoverPlanEntityType,error) { 
  var returnerr error 
  FailoverPlan := FailoverPlanEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","FailoverPlans",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&FailoverPlan)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error FailoverPlan",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking FailoverPlan",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return FailoverPlan,returnerr
}


/*
 * Jobs 
 * Not validated 
 */
func (v * VeeamRestServer) GetJobs() (JobEntityListType,error) { 
  var returnerr error 
  Jobs := JobEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","Jobs")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&Jobs)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error Jobs",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking Jobs",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return Jobs,returnerr
}


/*
 * FailoverPlans 
 * Not validated 
 */
func (v * VeeamRestServer) GetFailoverPlans() (FailoverPlanEntityListType,error) { 
  var returnerr error 
  FailoverPlans := FailoverPlanEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","FailoverPlans")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&FailoverPlans)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error FailoverPlans",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking FailoverPlans",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return FailoverPlans,returnerr
}


/*
 * ObjectInJob 
 * Not validated 
 */
func (v * VeeamRestServer) GetObjectInJob(idstring string) (ObjectInJobType,error) { 
  var returnerr error 
  ObjectInJob := ObjectInJobType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","ObjectsInJob",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&ObjectInJob)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error ObjectInJob",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking ObjectInJob",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return ObjectInJob,returnerr
}


/*
 * ObjectsInJob 
 * Not validated 
 */
func (v * VeeamRestServer) GetObjectsInJob() (ObjectInJobListType,error) { 
  var returnerr error 
  ObjectsInJob := ObjectInJobListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","ObjectsInJob")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&ObjectsInJob)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error ObjectsInJob",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking ObjectsInJob",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return ObjectsInJob,returnerr
}


/*
 * FailoveredVm 
 * Not validated 
 */
func (v * VeeamRestServer) GetFailoveredVm(idstring string) (FailoveredVmType,error) { 
  var returnerr error 
  FailoveredVm := FailoveredVmType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","FailoveredVms",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&FailoveredVm)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error FailoveredVm",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking FailoveredVm",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return FailoveredVm,returnerr
}


/*
 * FailoveredVms 
 * Not validated 
 */
func (v * VeeamRestServer) GetFailoveredVms() (FailoveredVmListType,error) { 
  var returnerr error 
  FailoveredVms := FailoveredVmListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","FailoveredVms")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&FailoveredVms)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error FailoveredVms",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking FailoveredVms",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return FailoveredVms,returnerr
}


/*
 * Backup 
 * Not validated 
 */
func (v * VeeamRestServer) GetBackup(idstring string) (BackupEntityType,error) { 
  var returnerr error 
  Backup := BackupEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","Backups",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&Backup)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error Backup",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking Backup",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return Backup,returnerr
}


/*
 * Backups 
 * Not validated 
 */
func (v * VeeamRestServer) GetBackups() (BackupEntityListType,error) { 
  var returnerr error 
  Backups := BackupEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","Backups")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&Backups)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error Backups",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking Backups",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return Backups,returnerr
}


/*
 * Replica 
 * Not validated 
 */
func (v * VeeamRestServer) GetReplica(idstring string) (ReplicaEntityType,error) { 
  var returnerr error 
  Replica := ReplicaEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","Replicas",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&Replica)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error Replica",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking Replica",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return Replica,returnerr
}


/*
 * Replicas 
 * Not validated 
 */
func (v * VeeamRestServer) GetReplicas() (ReplicaEntityListType,error) { 
  var returnerr error 
  Replicas := ReplicaEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","Replicas")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&Replicas)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error Replicas",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking Replicas",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return Replicas,returnerr
}


/*
 * RestorePoint 
 * Not validated 
 */
func (v * VeeamRestServer) GetRestorePoint(idstring string) (RestorePointEntityType,error) { 
  var returnerr error 
  RestorePoint := RestorePointEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","RestorePoints",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&RestorePoint)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error RestorePoint",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking RestorePoint",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return RestorePoint,returnerr
}


/*
 * RestorePoints 
 * Not validated 
 */
func (v * VeeamRestServer) GetRestorePoints() (RestorePointEntityListType,error) { 
  var returnerr error 
  RestorePoints := RestorePointEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","RestorePoints")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&RestorePoints)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error RestorePoints",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking RestorePoints",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return RestorePoints,returnerr
}


/*
 * VmRestorePoint 
 * Not validated 
 */
func (v * VeeamRestServer) GetVmRestorePoint(idstring string) (VmRestorePointEntityType,error) { 
  var returnerr error 
  VmRestorePoint := VmRestorePointEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","VmRestorePoints",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&VmRestorePoint)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error VmRestorePoint",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking VmRestorePoint",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return VmRestorePoint,returnerr
}


/*
 * VmRestorePoints 
 * Not validated 
 */
func (v * VeeamRestServer) GetVmRestorePoints() (VmRestorePointEntityListType,error) { 
  var returnerr error 
  VmRestorePoints := VmRestorePointEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","VmRestorePoints")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&VmRestorePoints)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error VmRestorePoints",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking VmRestorePoints",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return VmRestorePoints,returnerr
}


/*
 * VAppRestorePoint 
 * Not validated 
 */
func (v * VeeamRestServer) GetVAppRestorePoint(idstring string) (VAppRestorePointEntityType,error) { 
  var returnerr error 
  VAppRestorePoint := VAppRestorePointEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","VAppRestorePoints",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&VAppRestorePoint)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error VAppRestorePoint",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking VAppRestorePoint",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return VAppRestorePoint,returnerr
}


/*
 * VAppRestorePoints 
 * Not validated 
 */
func (v * VeeamRestServer) GetVAppRestorePoints() (VAppRestorePointEntityListType,error) { 
  var returnerr error 
  VAppRestorePoints := VAppRestorePointEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","VAppRestorePoints")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&VAppRestorePoints)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error VAppRestorePoints",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking VAppRestorePoints",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return VAppRestorePoints,returnerr
}


/*
 * VmReplicaPoint 
 * Not validated 
 */
func (v * VeeamRestServer) GetVmReplicaPoint(idstring string) (VmReplicaPointEntityType,error) { 
  var returnerr error 
  VmReplicaPoint := VmReplicaPointEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","VmReplicaPoints",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&VmReplicaPoint)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error VmReplicaPoint",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking VmReplicaPoint",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return VmReplicaPoint,returnerr
}


/*
 * VmReplicaPoints 
 * Not validated 
 */
func (v * VeeamRestServer) GetVmReplicaPoints() (VmReplicaPointEntityListType,error) { 
  var returnerr error 
  VmReplicaPoints := VmReplicaPointEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","VmReplicaPoints")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&VmReplicaPoints)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error VmReplicaPoints",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking VmReplicaPoints",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return VmReplicaPoints,returnerr
}


/*
 * VmRestorePointMount 
 * Not validated 
 */
func (v * VeeamRestServer) GetVmRestorePointMount(idstring string) (VmRestorePointMountType,error) { 
  var returnerr error 
  VmRestorePointMount := VmRestorePointMountType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","VmRestorePointMounts",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&VmRestorePointMount)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error VmRestorePointMount",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking VmRestorePointMount",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return VmRestorePointMount,returnerr
}


/*
 * VmRestorePointMounts 
 * Not validated 
 */
func (v * VeeamRestServer) GetVmRestorePointMounts() (VmRestorePointMountListType,error) { 
  var returnerr error 
  VmRestorePointMounts := VmRestorePointMountListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","VmRestorePointMounts")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&VmRestorePointMounts)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error VmRestorePointMounts",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking VmRestorePointMounts",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return VmRestorePointMounts,returnerr
}


/*
 * VmReplicaPointMount 
 * Not validated 
 */
func (v * VeeamRestServer) GetVmReplicaPointMount(idstring string) (VmReplicaPointMountType,error) { 
  var returnerr error 
  VmReplicaPointMount := VmReplicaPointMountType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","VmReplicaPointMounts",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&VmReplicaPointMount)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error VmReplicaPointMount",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking VmReplicaPointMount",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return VmReplicaPointMount,returnerr
}


/*
 * VmReplicaPointMounts 
 * Not validated 
 */
func (v * VeeamRestServer) GetVmReplicaPointMounts() (VmReplicaPointMountListType,error) { 
  var returnerr error 
  VmReplicaPointMounts := VmReplicaPointMountListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","VmReplicaPointMounts")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&VmReplicaPointMounts)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error VmReplicaPointMounts",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking VmReplicaPointMounts",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return VmReplicaPointMounts,returnerr
}


/*
 * FileSystemEntry 
 * Not validated 
 */
func (v * VeeamRestServer) GetFileSystemEntry(idstring string) (FileSystemEntryType,error) { 
  var returnerr error 
  FileSystemEntry := FileSystemEntryType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","FileSystemEntry",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&FileSystemEntry)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error FileSystemEntry",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking FileSystemEntry",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return FileSystemEntry,returnerr
}


/*
 * FileSystemEntries 
 * Not validated 
 */
func (v * VeeamRestServer) GetFileSystemEntries(idstring string) (FileSystemEntriesType,error) { 
  var returnerr error 
  FileSystemEntries := FileSystemEntriesType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","FileSystemEntries",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&FileSystemEntries)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error FileSystemEntries",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking FileSystemEntries",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return FileSystemEntries,returnerr
}


/*
 * HierarchyRoot 
 * Not validated 
 */
func (v * VeeamRestServer) GetHierarchyRoot(idstring string) (HierarchyRootEntityType,error) { 
  var returnerr error 
  HierarchyRoot := HierarchyRootEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","HierarchyRoots",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&HierarchyRoot)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error HierarchyRoot",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking HierarchyRoot",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return HierarchyRoot,returnerr
}


/*
 * HierarchyRoots 
 * Not validated 
 */
func (v * VeeamRestServer) GetHierarchyRoots() (HierarchyRootEntityListType,error) { 
  var returnerr error 
  HierarchyRoots := HierarchyRootEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","HierarchyRoots")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&HierarchyRoots)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error HierarchyRoots",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking HierarchyRoots",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return HierarchyRoots,returnerr
}


/*
 * Repository 
 * Not validated 
 */
func (v * VeeamRestServer) GetRepository(idstring string) (RepositoryEntityType,error) { 
  var returnerr error 
  Repository := RepositoryEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","Repositories",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&Repository)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error Repository",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking Repository",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return Repository,returnerr
}


/*
 * Repositories 
 * Not validated 
 */
func (v * VeeamRestServer) GetRepositories() (RepositoryEntityListType,error) { 
  var returnerr error 
  Repositories := RepositoryEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","Repositories")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&Repositories)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error Repositories",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking Repositories",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return Repositories,returnerr
}




/*
 * RestoreSession 
 * Not validated 
 */
func (v * VeeamRestServer) GetRestoreSession(idstring string) (RestoreSessionEntityType,error) { 
  var returnerr error 
  RestoreSession := RestoreSessionEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","RestoreSessions",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&RestoreSession)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error RestoreSession",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking RestoreSession",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return RestoreSession,returnerr
}


/*
 * RestoreSessions 
 * Not validated 
 */
func (v * VeeamRestServer) GetRestoreSessions() (RestoreSessionEntityListType,error) { 
  var returnerr error 
  RestoreSessions := RestoreSessionEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","RestoreSessions")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&RestoreSessions)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error RestoreSessions",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking RestoreSessions",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return RestoreSessions,returnerr
}


/*
 * BackupTaskSession 
 * Not validated 
 */
func (v * VeeamRestServer) GetBackupTaskSession(idstring string) (BackupTaskSessionEntityType,error) { 
  var returnerr error 
  BackupTaskSession := BackupTaskSessionEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","BackupTaskSessions",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&BackupTaskSession)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error BackupTaskSession",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking BackupTaskSession",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return BackupTaskSession,returnerr
}


/*
 * BackupTaskSessions 
 * Not validated 
 */
func (v * VeeamRestServer) GetBackupTaskSessions() (BackupTaskSessionEntityListType,error) { 
  var returnerr error 
  BackupTaskSessions := BackupTaskSessionEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","BackupTaskSessions")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&BackupTaskSessions)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error BackupTaskSessions",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking BackupTaskSessions",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return BackupTaskSessions,returnerr
}


/*
 * ReplicaTaskSession 
 * Not validated 
 */
func (v * VeeamRestServer) GetReplicaTaskSession(idstring string) (ReplicaTaskSessionEntityType,error) { 
  var returnerr error 
  ReplicaTaskSession := ReplicaTaskSessionEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","ReplicaTaskSessions",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&ReplicaTaskSession)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error ReplicaTaskSession",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking ReplicaTaskSession",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return ReplicaTaskSession,returnerr
}


/*
 * ReplicaTaskSessions 
 * Not validated 
 */
func (v * VeeamRestServer) GetReplicaTaskSessions() (ReplicaTaskSessionEntityListType,error) { 
  var returnerr error 
  ReplicaTaskSessions := ReplicaTaskSessionEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","ReplicaTaskSessions")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&ReplicaTaskSessions)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error ReplicaTaskSessions",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking ReplicaTaskSessions",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return ReplicaTaskSessions,returnerr
}


/*
 * LoginSpec 
 * Not validated 
 */
func (v * VeeamRestServer) GetLoginSpec(idstring string) (LoginSpecType,error) { 
  var returnerr error 
  LoginSpec := LoginSpecType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","LoginSpec",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&LoginSpec)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error LoginSpec",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking LoginSpec",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return LoginSpec,returnerr
}


/*
 * RestoreSpec 
 * Not validated 
 */
func (v * VeeamRestServer) GetRestoreSpec(idstring string) (RestoreSpecType,error) { 
  var returnerr error 
  RestoreSpec := RestoreSpecType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","RestoreSpec",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&RestoreSpec)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error RestoreSpec",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking RestoreSpec",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return RestoreSpec,returnerr
}


/*
 * FileRestoreSpec 
 * Not validated 
 */
func (v * VeeamRestServer) GetFileRestoreSpec(idstring string) (FileRestoreSpecType,error) { 
  var returnerr error 
  FileRestoreSpec := FileRestoreSpecType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","FileRestoreSpec",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&FileRestoreSpec)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error FileRestoreSpec",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking FileRestoreSpec",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return FileRestoreSpec,returnerr
}


/*
 * JobCloneSpec 
 * Not validated 
 */
func (v * VeeamRestServer) GetJobCloneSpec(idstring string) (JobCloneSpecType,error) { 
  var returnerr error 
  JobCloneSpec := JobCloneSpecType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","JobCloneSpec",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&JobCloneSpec)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error JobCloneSpec",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking JobCloneSpec",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return JobCloneSpec,returnerr
}


/*
 * CreateObjectInJobSpec 
 * Not validated 
 */
func (v * VeeamRestServer) GetCreateObjectInJobSpec(idstring string) (CreateObjectInJobSpecType,error) { 
  var returnerr error 
  CreateObjectInJobSpec := CreateObjectInJobSpecType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","CreateObjectInJobSpec",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CreateObjectInJobSpec)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CreateObjectInJobSpec",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CreateObjectInJobSpec",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CreateObjectInJobSpec,returnerr
}


/*
 * Task 
 * Not validated 
 */
func (v * VeeamRestServer) GetTask(idstring string) (TaskType,error) { 
  var returnerr error 
  Task := TaskType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","Tasks",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&Task)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error Task",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking Task",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return Task,returnerr
}


/*
 * Tasks 
 * Not validated 
 */
func (v * VeeamRestServer) GetTasks() (TaskListType,error) { 
  var returnerr error 
  Tasks := TaskListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","Tasks")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&Tasks)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error Tasks",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking Tasks",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return Tasks,returnerr
}


/*
 * Error 
 * Not validated 
 */
func (v * VeeamRestServer) GetError(idstring string) (ErrorType,error) { 
  var returnerr error 
  Error := ErrorType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","Error",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&Error)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error Error",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking Error",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return Error,returnerr
}


/*
 * HierarchyItem 
 * Not validated 
 */
func (v * VeeamRestServer) GetHierarchyItem(idstring string) (HierarchyItemType,error) { 
  var returnerr error 
  HierarchyItem := HierarchyItemType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","HierarchyItems",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&HierarchyItem)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error HierarchyItem",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking HierarchyItem",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return HierarchyItem,returnerr
}


/*
 * HierarchyItems 
 * Not validated 
 */
func (v * VeeamRestServer) GetHierarchyItems() (HierarchyItemListType,error) { 
  var returnerr error 
  HierarchyItems := HierarchyItemListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","HierarchyItems")),"GET") 
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
 * CredentialsInfo 
 * Not validated 
 */
func (v * VeeamRestServer) GetCredentialsInfo(idstring string) (CredentialsInfoType,error) { 
  var returnerr error 
  CredentialsInfo := CredentialsInfoType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","CredentialsInfoList",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CredentialsInfo)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CredentialsInfo",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CredentialsInfo",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CredentialsInfo,returnerr
}


/*
 * CredentialsInfoList 
 * Not validated 
 */
func (v * VeeamRestServer) GetCredentialsInfoList() (CredentialsInfoListType,error) { 
  var returnerr error 
  CredentialsInfoList := CredentialsInfoListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","CredentialsInfoList")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CredentialsInfoList)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CredentialsInfoList",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CredentialsInfoList",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CredentialsInfoList,returnerr
}


/*
 * CredentialsInfoSpec 
 * Not validated 
 */
func (v * VeeamRestServer) GetCredentialsInfoSpec(idstring string) (CredentialsInfoSpecType,error) { 
  var returnerr error 
  CredentialsInfoSpec := CredentialsInfoSpecType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","CredentialsInfoSpec",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CredentialsInfoSpec)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CredentialsInfoSpec",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CredentialsInfoSpec",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CredentialsInfoSpec,returnerr
}


/*
 * PasswordKeyInfo 
 * Not validated 
 */
func (v * VeeamRestServer) GetPasswordKeyInfo(idstring string) (PasswordKeyInfoType,error) { 
  var returnerr error 
  PasswordKeyInfo := PasswordKeyInfoType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","PasswordKeyInfoList",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&PasswordKeyInfo)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error PasswordKeyInfo",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking PasswordKeyInfo",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return PasswordKeyInfo,returnerr
}


/*
 * PasswordKeyInfoList 
 * Not validated 
 */
func (v * VeeamRestServer) GetPasswordKeyInfoList() (PasswordKeyInfoListType,error) { 
  var returnerr error 
  PasswordKeyInfoList := PasswordKeyInfoListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","PasswordKeyInfoList")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&PasswordKeyInfoList)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error PasswordKeyInfoList",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking PasswordKeyInfoList",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return PasswordKeyInfoList,returnerr
}


/*
 * PasswordKeyInfoSpec 
 * Not validated 
 */
func (v * VeeamRestServer) GetPasswordKeyInfoSpec(idstring string) (PasswordKeyInfoSpecType,error) { 
  var returnerr error 
  PasswordKeyInfoSpec := PasswordKeyInfoSpecType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","PasswordKeyInfoSpec",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&PasswordKeyInfoSpec)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error PasswordKeyInfoSpec",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking PasswordKeyInfoSpec",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return PasswordKeyInfoSpec,returnerr
}


/*
 * CatalogVm 
 * Not validated 
 */
func (v * VeeamRestServer) GetCatalogVm(idstring string) (CatalogVmEntityType,error) { 
  var returnerr error 
  CatalogVm := CatalogVmEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","CatalogVms",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CatalogVm)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CatalogVm",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CatalogVm",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CatalogVm,returnerr
}


/*
 * CatalogVms 
 * Not validated 
 */
func (v * VeeamRestServer) GetCatalogVms() (CatalogVmEntityListType,error) { 
  var returnerr error 
  CatalogVms := CatalogVmEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","CatalogVms")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CatalogVms)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CatalogVms",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CatalogVms",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CatalogVms,returnerr
}


/*
 * CatalogVmRestorePoint 
 * Not validated 
 */
func (v * VeeamRestServer) GetCatalogVmRestorePoint(idstring string) (CatalogVmRestorePointEntityType,error) { 
  var returnerr error 
  CatalogVmRestorePoint := CatalogVmRestorePointEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","CatalogVmRestorePoints",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CatalogVmRestorePoint)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CatalogVmRestorePoint",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CatalogVmRestorePoint",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CatalogVmRestorePoint,returnerr
}


/*
 * CatalogVmRestorePoints 
 * Not validated 
 */
func (v * VeeamRestServer) GetCatalogVmRestorePoints() (CatalogVmRestorePointEntityListType,error) { 
  var returnerr error 
  CatalogVmRestorePoints := CatalogVmRestorePointEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","CatalogVmRestorePoints")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CatalogVmRestorePoints)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CatalogVmRestorePoints",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CatalogVmRestorePoints",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CatalogVmRestorePoints,returnerr
}


/*
 * EnterpriseRole 
 * Not validated 
 */
func (v * VeeamRestServer) GetEnterpriseRole(idstring string) (EnterpriseRoleEntityType,error) { 
  var returnerr error 
  EnterpriseRole := EnterpriseRoleEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","EnterpriseRoles",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&EnterpriseRole)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error EnterpriseRole",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking EnterpriseRole",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return EnterpriseRole,returnerr
}


/*
 * EnterpriseRoles 
 * Not validated 
 */
func (v * VeeamRestServer) GetEnterpriseRoles() (EnterpriseRoleEntityListType,error) { 
  var returnerr error 
  EnterpriseRoles := EnterpriseRoleEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","EnterpriseRoles")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&EnterpriseRoles)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error EnterpriseRoles",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking EnterpriseRoles",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return EnterpriseRoles,returnerr
}


/*
 * EnterpriseAccount 
 * Not validated 
 */
func (v * VeeamRestServer) GetEnterpriseAccount(idstring string) (EnterpriseAccountEntityType,error) { 
  var returnerr error 
  EnterpriseAccount := EnterpriseAccountEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","EnterpriseAccounts",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&EnterpriseAccount)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error EnterpriseAccount",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking EnterpriseAccount",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return EnterpriseAccount,returnerr
}


/*
 * EnterpriseAccounts 
 * Not validated 
 */
func (v * VeeamRestServer) GetEnterpriseAccounts() (EnterpriseAccountEntityListType,error) { 
  var returnerr error 
  EnterpriseAccounts := EnterpriseAccountEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","EnterpriseAccounts")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&EnterpriseAccounts)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error EnterpriseAccounts",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking EnterpriseAccounts",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return EnterpriseAccounts,returnerr
}


/*
 * EnterpriseAccountHierarchyScope 
 * Not validated 
 */
func (v * VeeamRestServer) GetEnterpriseAccountHierarchyScope(idstring string) (EnterpriseAccountHierarchyScopeType,error) { 
  var returnerr error 
  EnterpriseAccountHierarchyScope := EnterpriseAccountHierarchyScopeType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","EnterpriseAccountHierarchyScopes",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&EnterpriseAccountHierarchyScope)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error EnterpriseAccountHierarchyScope",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking EnterpriseAccountHierarchyScope",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return EnterpriseAccountHierarchyScope,returnerr
}


/*
 * EnterpriseAccountHierarchyScopes 
 * Not validated 
 */
func (v * VeeamRestServer) GetEnterpriseAccountHierarchyScopes() (EnterpriseAccountHierarchyScopeListType,error) { 
  var returnerr error 
  EnterpriseAccountHierarchyScopes := EnterpriseAccountHierarchyScopeListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","EnterpriseAccountHierarchyScopes")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&EnterpriseAccountHierarchyScopes)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error EnterpriseAccountHierarchyScopes",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking EnterpriseAccountHierarchyScopes",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return EnterpriseAccountHierarchyScopes,returnerr
}


/*
 * EnterpriseAccountInRole 
 * Not validated 
 */
func (v * VeeamRestServer) GetEnterpriseAccountInRole(idstring string) (EnterpriseAccountInRoleType,error) { 
  var returnerr error 
  EnterpriseAccountInRole := EnterpriseAccountInRoleType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","EnterpriseAccountInRoleList",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&EnterpriseAccountInRole)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error EnterpriseAccountInRole",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking EnterpriseAccountInRole",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return EnterpriseAccountInRole,returnerr
}


/*
 * EnterpriseAccountInRoleList 
 * Not validated 
 */
func (v * VeeamRestServer) GetEnterpriseAccountInRoleList() (EnterpriseAccountInRoleListType,error) { 
  var returnerr error 
  EnterpriseAccountInRoleList := EnterpriseAccountInRoleListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","EnterpriseAccountInRoleList")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&EnterpriseAccountInRoleList)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error EnterpriseAccountInRoleList",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking EnterpriseAccountInRoleList",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return EnterpriseAccountInRoleList,returnerr
}


/*
 * HierarchyScopeCreateSpec 
 * Not validated 
 */
func (v * VeeamRestServer) GetHierarchyScopeCreateSpec(idstring string) (HierarchyScopeCreateSpecType,error) { 
  var returnerr error 
  HierarchyScopeCreateSpec := HierarchyScopeCreateSpecType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","HierarchyScopeCreateSpec",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&HierarchyScopeCreateSpec)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error HierarchyScopeCreateSpec",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking HierarchyScopeCreateSpec",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return HierarchyScopeCreateSpec,returnerr
}


/*
 * EnterpriseAccountInRoleCreateSpec 
 * Not validated 
 */
func (v * VeeamRestServer) GetEnterpriseAccountInRoleCreateSpec(idstring string) (EnterpriseAccountInRoleCreateSpecType,error) { 
  var returnerr error 
  EnterpriseAccountInRoleCreateSpec := EnterpriseAccountInRoleCreateSpecType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","EnterpriseAccountInRoleCreateSpec",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&EnterpriseAccountInRoleCreateSpec)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error EnterpriseAccountInRoleCreateSpec",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking EnterpriseAccountInRoleCreateSpec",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return EnterpriseAccountInRoleCreateSpec,returnerr
}


/*
 * EnterpriseSecuritySettings 
 * Not validated 
 */
func (v * VeeamRestServer) GetEnterpriseSecuritySettings(idstring string) (EnterpriseSecuritySettingsType,error) { 
  var returnerr error 
  EnterpriseSecuritySettings := EnterpriseSecuritySettingsType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","EnterpriseSecuritySettings",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&EnterpriseSecuritySettings)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error EnterpriseSecuritySettings",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking EnterpriseSecuritySettings",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return EnterpriseSecuritySettings,returnerr
}


/*
 * EnterpriseAccountCreateSpec 
 * Not validated 
 */
func (v * VeeamRestServer) GetEnterpriseAccountCreateSpec(idstring string) (EnterpriseAccountCreateSpecType,error) { 
  var returnerr error 
  EnterpriseAccountCreateSpec := EnterpriseAccountCreateSpecType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","EnterpriseAccountCreateSpec",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&EnterpriseAccountCreateSpec)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error EnterpriseAccountCreateSpec",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking EnterpriseAccountCreateSpec",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return EnterpriseAccountCreateSpec,returnerr
}


/*
 * RebuildScopeJobSpec 
 * Not validated 
 */
func (v * VeeamRestServer) GetRebuildScopeJobSpec(idstring string) (RebuildScopeJobSpecType,error) { 
  var returnerr error 
  RebuildScopeJobSpec := RebuildScopeJobSpecType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","RebuildScopeJobSpec",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&RebuildScopeJobSpec)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error RebuildScopeJobSpec",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking RebuildScopeJobSpec",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return RebuildScopeJobSpec,returnerr
}


/*
 * JobManagementSpec 
 * Not validated 
 */
func (v * VeeamRestServer) GetJobManagementSpec(idstring string) (JobManagementSpecType,error) { 
  var returnerr error 
  JobManagementSpec := JobManagementSpecType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","JobManagementSpec",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&JobManagementSpec)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error JobManagementSpec",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking JobManagementSpec",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return JobManagementSpec,returnerr
}


/*
 * BackupServerSpec 
 * Not validated 
 */
func (v * VeeamRestServer) GetBackupServerSpec(idstring string) (BackupServerSpecType,error) { 
  var returnerr error 
  BackupServerSpec := BackupServerSpecType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","BackupServerSpec",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&BackupServerSpec)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error BackupServerSpec",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking BackupServerSpec",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return BackupServerSpec,returnerr
}


/*
 * VeeamZipStartupSpec 
 * Not validated 
 */
func (v * VeeamRestServer) GetVeeamZipStartupSpec(idstring string) (VeeamZipStartupSpecType,error) { 
  var returnerr error 
  VeeamZipStartupSpec := VeeamZipStartupSpecType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","VeeamZipStartupSpec",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&VeeamZipStartupSpec)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error VeeamZipStartupSpec",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking VeeamZipStartupSpec",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return VeeamZipStartupSpec,returnerr
}


/*
 * QuickBackupStartupSpec 
 * Not validated 
 */
func (v * VeeamRestServer) GetQuickBackupStartupSpec(idstring string) (QuickBackupStartupSpecType,error) { 
  var returnerr error 
  QuickBackupStartupSpec := QuickBackupStartupSpecType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","QuickBackupStartupSpec",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&QuickBackupStartupSpec)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error QuickBackupStartupSpec",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking QuickBackupStartupSpec",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return QuickBackupStartupSpec,returnerr
}


/*
 * WanAccelerator 
 * Not validated 
 */
func (v * VeeamRestServer) GetWanAccelerator(idstring string) (WanAcceleratorEntityType,error) { 
  var returnerr error 
  WanAccelerator := WanAcceleratorEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","WanAccelerators",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&WanAccelerator)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error WanAccelerator",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking WanAccelerator",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return WanAccelerator,returnerr
}


/*
 * WanAccelerators 
 * Not validated 
 */
func (v * VeeamRestServer) GetWanAccelerators() (WanAcceleratorEntityListType,error) { 
  var returnerr error 
  WanAccelerators := WanAcceleratorEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","WanAccelerators")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&WanAccelerators)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error WanAccelerators",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking WanAccelerators",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return WanAccelerators,returnerr
}


/*
 * CloudGateway 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudGateway(idstring string) (CloudGatewayEntityType,error) { 
  var returnerr error 
  CloudGateway := CloudGatewayEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","CloudGateways",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudGateway)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudGateway",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudGateway",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudGateway,returnerr
}


/*
 * CloudGateways 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudGateways() (CloudGatewayEntityListType,error) { 
  var returnerr error 
  CloudGateways := CloudGatewayEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","CloudGateways")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudGateways)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudGateways",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudGateways",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudGateways,returnerr
}


/*
 * CloudTenant 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudTenant(idstring string) (CloudTenantEntityType,error) { 
  var returnerr error 
  CloudTenant := CloudTenantEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","CloudTenants",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudTenant)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudTenant",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudTenant",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudTenant,returnerr
}


/*
 * CloudTenants 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudTenants() (CloudTenantEntityListType,error) { 
  var returnerr error 
  CloudTenants := CloudTenantEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","CloudTenants")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudTenants)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudTenants",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudTenants",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudTenants,returnerr
}


/*
 * CloudTenantResource 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudTenantResource(idstring string) (CloudTenantResourceType,error) { 
  var returnerr error 
  CloudTenantResource := CloudTenantResourceType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","CloudTenantResources",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudTenantResource)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudTenantResource",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudTenantResource",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudTenantResource,returnerr
}


/*
 * CloudTenantResources 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudTenantResources() (CloudTenantResourceListType,error) { 
  var returnerr error 
  CloudTenantResources := CloudTenantResourceListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","CloudTenantResources")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudTenantResources)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudTenantResources",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudTenantResources",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudTenantResources,returnerr
}


/*
 * CreateCloudGatewaySpec 
 * Not validated 
 */
func (v * VeeamRestServer) GetCreateCloudGatewaySpec(idstring string) (CreateCloudGatewaySpecType,error) { 
  var returnerr error 
  CreateCloudGatewaySpec := CreateCloudGatewaySpecType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","CreateCloudGatewaySpec",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CreateCloudGatewaySpec)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CreateCloudGatewaySpec",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CreateCloudGatewaySpec",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CreateCloudGatewaySpec,returnerr
}


/*
 * CreateCloudTenantSpec 
 * Not validated 
 */
func (v * VeeamRestServer) GetCreateCloudTenantSpec(idstring string) (CreateCloudTenantSpecType,error) { 
  var returnerr error 
  CreateCloudTenantSpec := CreateCloudTenantSpecType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","CreateCloudTenantSpec",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CreateCloudTenantSpec)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CreateCloudTenantSpec",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CreateCloudTenantSpec",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CreateCloudTenantSpec,returnerr
}


/*
 * CreateCloudTenantResourceSpec 
 * Not validated 
 */
func (v * VeeamRestServer) GetCreateCloudTenantResourceSpec(idstring string) (CreateCloudTenantResourceSpecType,error) { 
  var returnerr error 
  CreateCloudTenantResourceSpec := CreateCloudTenantResourceSpecType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","CreateCloudTenantResourceSpec",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CreateCloudTenantResourceSpec)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CreateCloudTenantResourceSpec",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CreateCloudTenantResourceSpec",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CreateCloudTenantResourceSpec,returnerr
}


/*
 * CloudHardwarePlans 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudHardwarePlans() (CloudHardwarePlanEntityListType,error) { 
  var returnerr error 
  CloudHardwarePlans := CloudHardwarePlanEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","CloudHardwarePlans")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudHardwarePlans)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudHardwarePlans",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudHardwarePlans",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudHardwarePlans,returnerr
}


/*
 * CloudHardwarePlan 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudHardwarePlan(idstring string) (CloudHardwarePlanEntityType,error) { 
  var returnerr error 
  CloudHardwarePlan := CloudHardwarePlanEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","CloudHardwarePlans",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudHardwarePlan)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudHardwarePlan",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudHardwarePlan",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudHardwarePlan,returnerr
}


/*
 * CloudHardwarePlanCreateSpec 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudHardwarePlanCreateSpec(idstring string) (CloudHardwarePlanCreateSpecType,error) { 
  var returnerr error 
  CloudHardwarePlanCreateSpec := CloudHardwarePlanCreateSpecType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","CloudHardwarePlanCreateSpec",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudHardwarePlanCreateSpec)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudHardwarePlanCreateSpec",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudHardwarePlanCreateSpec",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudHardwarePlanCreateSpec,returnerr
}


/*
 * CloudPublicIpAddresses 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudPublicIpAddresses() (CloudPublicIpAddressEntityListType,error) { 
  var returnerr error 
  CloudPublicIpAddresses := CloudPublicIpAddressEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","CloudPublicIpAddresses")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudPublicIpAddresses)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudPublicIpAddresses",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudPublicIpAddresses",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudPublicIpAddresses,returnerr
}


/*
 * CloudPublicIpAddress 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudPublicIpAddress(idstring string) (CloudPublicIpAddressEntityType,error) { 
  var returnerr error 
  CloudPublicIpAddress := CloudPublicIpAddressEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","CloudPublicIpAddresses",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudPublicIpAddress)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudPublicIpAddress",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudPublicIpAddress",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudPublicIpAddress,returnerr
}


/*
 * CloudPublicIpAddressCreateSpec 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudPublicIpAddressCreateSpec(idstring string) (CloudPublicIpAddressCreateSpecType,error) { 
  var returnerr error 
  CloudPublicIpAddressCreateSpec := CloudPublicIpAddressCreateSpecType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","CloudPublicIpAddressCreateSpec",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudPublicIpAddressCreateSpec)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudPublicIpAddressCreateSpec",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudPublicIpAddressCreateSpec",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudPublicIpAddressCreateSpec,returnerr
}


/*
 * CloudTenantComputeResources 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudTenantComputeResources() (CloudTenantComputeResourceListType,error) { 
  var returnerr error 
  CloudTenantComputeResources := CloudTenantComputeResourceListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","CloudTenantComputeResources")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudTenantComputeResources)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudTenantComputeResources",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudTenantComputeResources",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudTenantComputeResources,returnerr
}


/*
 * CloudTenantComputeResource 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudTenantComputeResource(idstring string) (CloudTenantComputeResourceType,error) { 
  var returnerr error 
  CloudTenantComputeResource := CloudTenantComputeResourceType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","CloudTenantComputeResources",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudTenantComputeResource)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudTenantComputeResource",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudTenantComputeResource",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudTenantComputeResource,returnerr
}


/*
 * CloudTenantComputeResourceCreateSpec 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudTenantComputeResourceCreateSpec(idstring string) (CloudTenantComputeResourceCreateSpecType,error) { 
  var returnerr error 
  CloudTenantComputeResourceCreateSpec := CloudTenantComputeResourceCreateSpecType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","CloudTenantComputeResourceCreateSpec",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudTenantComputeResourceCreateSpec)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudTenantComputeResourceCreateSpec",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudTenantComputeResourceCreateSpec",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudTenantComputeResourceCreateSpec,returnerr
}


/*
 * CloudVmReplicaPoints 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudVmReplicaPoints() (CloudVmReplicaPointEntityListType,error) { 
  var returnerr error 
  CloudVmReplicaPoints := CloudVmReplicaPointEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","CloudVmReplicaPoints")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudVmReplicaPoints)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudVmReplicaPoints",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudVmReplicaPoints",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudVmReplicaPoints,returnerr
}


/*
 * CloudVmReplicaPoint 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudVmReplicaPoint(idstring string) (CloudVmReplicaPointEntityType,error) { 
  var returnerr error 
  CloudVmReplicaPoint := CloudVmReplicaPointEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","CloudVmReplicaPoints",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudVmReplicaPoint)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudVmReplicaPoint",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudVmReplicaPoint",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudVmReplicaPoint,returnerr
}


/*
 * CloudFailoverPlans 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudFailoverPlans() (CloudFailoverPlanEntityListType,error) { 
  var returnerr error 
  CloudFailoverPlans := CloudFailoverPlanEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","CloudFailoverPlans")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudFailoverPlans)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudFailoverPlans",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudFailoverPlans",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudFailoverPlans,returnerr
}


/*
 * CloudFailoverPlan 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudFailoverPlan(idstring string) (CloudFailoverPlanEntityType,error) { 
  var returnerr error 
  CloudFailoverPlan := CloudFailoverPlanEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","CloudFailoverPlans",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudFailoverPlan)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudFailoverPlan",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudFailoverPlan",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudFailoverPlan,returnerr
}


/*
 * CloudFailoverPlanManagementSpec 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudFailoverPlanManagementSpec(idstring string) (CloudFailoverPlanManagementSpecType,error) { 
  var returnerr error 
  CloudFailoverPlanManagementSpec := CloudFailoverPlanManagementSpecType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","CloudFailoverPlanManagementSpec",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudFailoverPlanManagementSpec)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudFailoverPlanManagementSpec",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudFailoverPlanManagementSpec",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudFailoverPlanManagementSpec,returnerr
}


/*
 * VlanConfigurations 
 * Not validated 
 */
func (v * VeeamRestServer) GetVlanConfigurations() (VlanConfigurationEntityListType,error) { 
  var returnerr error 
  VlanConfigurations := VlanConfigurationEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","VlanConfigurations")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&VlanConfigurations)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error VlanConfigurations",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking VlanConfigurations",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return VlanConfigurations,returnerr
}


/*
 * VlanConfiguration 
 * Not validated 
 */
func (v * VeeamRestServer) GetVlanConfiguration(idstring string) (VlanConfigurationEntityType,error) { 
  var returnerr error 
  VlanConfiguration := VlanConfigurationEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","VlanConfigurations",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&VlanConfiguration)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error VlanConfiguration",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking VlanConfiguration",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return VlanConfiguration,returnerr
}


/*
 * VlanConfigurationCreateSpec 
 * Not validated 
 */
func (v * VeeamRestServer) GetVlanConfigurationCreateSpec(idstring string) (CloudVlanConfigurationCreateSpecType,error) { 
  var returnerr error 
  VlanConfigurationCreateSpec := CloudVlanConfigurationCreateSpecType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","VlanConfigurationCreateSpec",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&VlanConfigurationCreateSpec)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error VlanConfigurationCreateSpec",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking VlanConfigurationCreateSpec",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return VlanConfigurationCreateSpec,returnerr
}


/*
 * CloudFailoverSession 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudFailoverSession(idstring string) (CloudFailoverSessionEntityType,error) { 
  var returnerr error 
  CloudFailoverSession := CloudFailoverSessionEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","CloudFailoverSessions",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudFailoverSession)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudFailoverSession",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudFailoverSession",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudFailoverSession,returnerr
}


/*
 * CloudFailoverSessions 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudFailoverSessions() (CloudFailoverSessionEntityListType,error) { 
  var returnerr error 
  CloudFailoverSessions := CloudFailoverSessionEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","CloudFailoverSessions")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudFailoverSessions)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudFailoverSessions",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudFailoverSessions",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudFailoverSessions,returnerr
}


/*
 * CloudFailoveredVm 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudFailoveredVm(idstring string) (CloudFailoveredVmType,error) { 
  var returnerr error 
  CloudFailoveredVm := CloudFailoveredVmType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","CloudFailoveredVms",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudFailoveredVm)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudFailoveredVm",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudFailoveredVm",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudFailoveredVm,returnerr
}


/*
 * CloudFailoveredVms 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudFailoveredVms() (CloudFailoveredVmListType,error) { 
  var returnerr error 
  CloudFailoveredVms := CloudFailoveredVmListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","CloudFailoveredVms")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudFailoveredVms)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudFailoveredVms",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudFailoveredVms",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudFailoveredVms,returnerr
}


/*
 * CloudConnectService 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudConnectService(idstring string) (CloudConnectServiceType,error) { 
  var returnerr error 
  CloudConnectService := CloudConnectServiceType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","CloudConnectService",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudConnectService)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudConnectService",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudConnectService",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudConnectService,returnerr
}


/*
 * CloudReplicas 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudReplicas() (CloudReplicaEntityListType,error) { 
  var returnerr error 
  CloudReplicas := CloudReplicaEntityListType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s?format=Entity","CloudReplicas")),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudReplicas)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudReplicas",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudReplicas",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudReplicas,returnerr
}


/*
 * CloudReplica 
 * Not validated 
 */
func (v * VeeamRestServer) GetCloudReplica(idstring string) (CloudReplicaEntityType,error) { 
  var returnerr error 
  CloudReplica := CloudReplicaEntityType{} 
  if (v.SessionId != "") { 
     vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s?format=Entity","CloudReplicas",idstring)),"GET") 
     xmlin, err := v.Request(vrr)
     if err == nil {	
     	err = xml.Unmarshal([]byte(xmlin),&CloudReplica)
     	if err != nil {
     	    returnerr = &VeeamRestError{"Unmarshal error CloudReplica",err.Error()}
     	}
     } else {	
        returnerr = &VeeamRestError{"Invalid request asking CloudReplica",err.Error()}
     }
  } else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
  return CloudReplica,returnerr
}


