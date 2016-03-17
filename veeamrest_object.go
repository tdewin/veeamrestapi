package veeamrestapi


 
import (
 "encoding/xml"
 )
 
 /*
 * LoginSpecType 
 * Not validated 
 */
type LoginSpecType struct { 
   XMLName xml.Name
   VMwareSSOToken string `xml:"VMwareSSOToken,omitempty"`
   TenantCredentials *TenantCredentialsInfoType `xml:"TenantCredentials,omitempty"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}
func NewLoginSpecType() (*LoginSpecType) {
  varLoginSpecType := LoginSpecType{}
  varLoginSpecType.XMLName.Local = "LoginSpec"
  varLoginSpecType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varLoginSpecType
}

/*
 * LogonSessionListType 
 * Not validated 
 */
type LogonSessionListType struct { 
   XMLName xml.Name
   LogonSession []*LogonSessionType `xml:"LogonSession"`
   //Inhereting from ListType
}

/*
 * LogonSessionType 
 * Not validated 
 */
type LogonSessionType struct { 
   XMLName xml.Name
   UserName string `xml:"UserName,omitempty"`
   SessionId string `xml:"SessionId,omitempty"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l LogonSessionType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * SummaryReportType 
 * Not validated 
 */
type SummaryReportType struct { 
   XMLName xml.Name
   //Inhereting from ReportResourceType
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l SummaryReportType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * OverviewReportFrameType 
 * Not validated 
 */
type OverviewReportFrameType struct { 
   XMLName xml.Name
   BackupServers int `xml:"BackupServers"`
   ProxyServers int `xml:"ProxyServers"`
   RepositoryServers int `xml:"RepositoryServers"`
   RunningJobs int `xml:"RunningJobs"`
   ScheduledJobs int `xml:"ScheduledJobs"`
   SuccessfulVmLastestStates int `xml:"SuccessfulVmLastestStates"`
   WarningVmLastestStates int `xml:"WarningVmLastestStates"`
   FailedVmLastestStates int `xml:"FailedVmLastestStates"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l OverviewReportFrameType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * VmsOverviewReportFrameType 
 * Not validated 
 */
type VmsOverviewReportFrameType struct { 
   XMLName xml.Name
   ProtectedVms int `xml:"ProtectedVms"`
   BackedUpVms int `xml:"BackedUpVms"`
   ReplicatedVms int `xml:"ReplicatedVms"`
   RestorePoints int `xml:"RestorePoints"`
   FullBackupPointsSize int64 `xml:"FullBackupPointsSize"`
   IncrementalBackupPointsSize int64 `xml:"IncrementalBackupPointsSize"`
   ReplicaRestorePointsSize int64 `xml:"ReplicaRestorePointsSize"`
   SourceVmsSize int64 `xml:"SourceVmsSize"`
   SuccessBackupPercents int `xml:"SuccessBackupPercents"`
   ProtectedVmsReportLink string `xml:"ProtectedVmsReportLink"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l VmsOverviewReportFrameType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * JobStatisticsReportFrameType 
 * Not validated 
 */
type JobStatisticsReportFrameType struct { 
   XMLName xml.Name
   RunningJobs int `xml:"RunningJobs"`
   ScheduledJobs int `xml:"ScheduledJobs"`
   ScheduledBackupJobs int `xml:"ScheduledBackupJobs"`
   ScheduledReplicaJobs int `xml:"ScheduledReplicaJobs"`
   TotalJobRuns int `xml:"TotalJobRuns"`
   SuccessfulJobRuns int `xml:"SuccessfulJobRuns"`
   WarningsJobRuns int `xml:"WarningsJobRuns"`
   FailedJobRuns int `xml:"FailedJobRuns"`
   MaxJobDuration int `xml:"MaxJobDuration"`
   MaxBackupJobDuration int `xml:"MaxBackupJobDuration"`
   MaxReplicaJobDuration int `xml:"MaxReplicaJobDuration"`
   MaxDurationBackupJobName string `xml:"MaxDurationBackupJobName"`
   MaxDurationReplicaJobName string `xml:"MaxDurationReplicaJobName"`
   BackupJobStatusReportLink string `xml:"BackupJobStatusReportLink"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l JobStatisticsReportFrameType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * InfoType 
 * Not validated 
 */
type InfoType struct { 
   XMLName xml.Name
}

/*
 * RepositoryReportFrameType 
 * Not validated 
 */
type RepositoryReportFrameType struct { 
   XMLName xml.Name
   Period []*RepositoryReportFrameTypeNestedPeriod `xml:"Period"`
   CapacityPlanningReportLink string `xml:"CapacityPlanningReportLink,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l RepositoryReportFrameType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * RepositoryReportFrameTypeNestedPeriod 
 * Not validated 
 */
type RepositoryReportFrameTypeNestedPeriod struct { 
   XMLName xml.Name
   Name string `xml:"Name"`
   Capacity int64 `xml:"Capacity"`
   FreeSpace int64 `xml:"FreeSpace"`
   BackupSize int64 `xml:"BackupSize"`
}

/*
 * ProcessedVmsReportFrameType 
 * Not validated 
 */
type ProcessedVmsReportFrameType struct { 
   XMLName xml.Name
   Day []*ProcessedVmsReportFrameTypeNestedDay `xml:"Day"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l ProcessedVmsReportFrameType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * ProcessedVmsReportFrameTypeNestedDay 
 * Not validated 
 */
type ProcessedVmsReportFrameTypeNestedDay struct { 
   XMLName xml.Name
   Timestamp DateTime `xml:"Timestamp,attr"`
   ReplicatedVms int `xml:"ReplicatedVms,attr"`
   BackupedVms int `xml:"BackupedVms,attr"`
}

/*
 * EnterpriseManagerType 
 * Not validated 
 */
type EnterpriseManagerType struct { 
   XMLName xml.Name
   SupportedVersions *SupportedVersionListType `xml:"SupportedVersions"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l EnterpriseManagerType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * SupportedVersionListType 
 * Not validated 
 */
type SupportedVersionListType struct { 
   XMLName xml.Name
   SupportedVersion []*SupportedVersionType `xml:"SupportedVersion"`
   //Inhereting from ListType
}

/*
 * SupportedVersionType 
 * Not validated 
 */
type SupportedVersionType struct { 
   XMLName xml.Name
   Links *LinkListType `xml:"Links"`
   Name string `xml:"Name,attr"`
   //Inhereting from InfoType
}
func (l SupportedVersionType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * ListType 
 * Not validated 
 */
type ListType struct { 
   XMLName xml.Name
}

/*
 * ReportResourceType 
 * Not validated 
 */
type ReportResourceType struct { 
   XMLName xml.Name
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l ReportResourceType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * LinkListType 
 * Not validated 
 */
type LinkListType struct { 
   XMLName xml.Name
   Link []*LinkType `xml:"Link"`
}

/*
 * LinkType 
 * Not validated 
 */
type LinkType struct { 
   XMLName xml.Name
   Rel string `xml:"Rel,attr"`
   //Inhereting from ReferenceType
   Href UrlType `xml:"Href,attr"`
   Name string `xml:"Name,attr"`
   Type string `xml:"Type,attr"`
}

/*
 * ReferenceType 
 * Not validated 
 */
type ReferenceType struct { 
   XMLName xml.Name
   Href UrlType `xml:"Href,attr"`
   Name string `xml:"Name,attr"`
   Type string `xml:"Type,attr"`
}

/*
 * ResourceType 
 * Not validated 
 */
type ResourceType struct { 
   XMLName xml.Name
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l ResourceType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * EntityType 
 * Not validated 
 */
type EntityType struct { 
   XMLName xml.Name
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l EntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * EntitiesType 
 * Not validated 
 */
type EntitiesType struct { 
   XMLName xml.Name
   Jobs *JobEntityListType `xml:"Jobs,omitempty"`
   FailoverPlans *FailoverPlanEntityListType `xml:"FailoverPlans,omitempty"`
   Backups *BackupEntityListType `xml:"Backups,omitempty"`
   Replicas *ReplicaEntityListType `xml:"Replicas,omitempty"`
   Repositories *RepositoryEntityListType `xml:"Repositories,omitempty"`
   RestorePoints *RestorePointEntityListType `xml:"RestorePoints,omitempty"`
   VmRestorePoints *VmRestorePointEntityListType `xml:"VmRestorePoints,omitempty"`
   VAppRestorePoints *VAppRestorePointEntityListType `xml:"VAppRestorePoints,omitempty"`
   VmReplicaPoints *VmReplicaPointEntityListType `xml:"VmReplicaPoints,omitempty"`
   BackupJobSessions *BackupJobSessionEntityListType `xml:"BackupJobSessions,omitempty"`
   ReplicaJobSessions *ReplicaJobSessionEntityListType `xml:"ReplicaJobSessions,omitempty"`
   ReplicaTaskSessions *ReplicaTaskSessionEntityListType `xml:"ReplicaTaskSessions,omitempty"`
   RestoreSessions *RestoreSessionEntityListType `xml:"RestoreSessions,omitempty"`
   HierarchyRoots *HierarchyRootEntityListType `xml:"HierarchyRoots,omitempty"`
   BackupTaskSessions *BackupTaskSessionEntityListType `xml:"BackupTaskSessions,omitempty"`
   BackupServers *BackupServerEntityListType `xml:"BackupServers,omitempty"`
   ManagedServers *ManagedServerEntityListType `xml:"ManagedServers,omitempty"`
   EnterpiseRoles *EnterpriseRoleEntityListType `xml:"EnterpiseRoles,omitempty"`
   EnterpiseAccounts *EnterpriseAccountEntityListType `xml:"EnterpiseAccounts,omitempty"`
   WanAccelerators *WanAcceleratorEntityListType `xml:"WanAccelerators,omitempty"`
   CloudGateways *CloudGatewayEntityListType `xml:"CloudGateways,omitempty"`
   CloudTenants *CloudTenantEntityListType `xml:"CloudTenants,omitempty"`
   Passwords *PasswordKeyInfoListType `xml:"Passwords,omitempty"`
   Credentials *CredentialsInfoListType `xml:"Credentials,omitempty"`
}

/*
 * ResourcesType 
 * Not validated 
 */
type ResourcesType struct { 
   XMLName xml.Name
   Files *FileEntryListType `xml:"Files,omitempty"`
   Directories *DirectoryEntryListType `xml:"Directories,omitempty"`
   Tasks *TaskListType `xml:"Tasks,omitempty"`
   CredentialsList *CredentialsInfoListType `xml:"CredentialsList,omitempty"`
   sInObjectJob *ObjectInJobListType `xml:"sInObjectJob,omitempty"`
   AccountRole *EnterpriseAccountInRoleListType `xml:"AccountRole,omitempty"`
   AccountHierarchyScope *EnterpriseAccountHierarchyScopeListType `xml:"AccountHierarchyScope,omitempty"`
   PasswordKeyInfoList *PasswordKeyInfoListType `xml:"PasswordKeyInfoList,omitempty"`
}

/*
 * ParamsType 
 * Not validated 
 */
type ParamsType struct { 
   XMLName xml.Name
}

/*
 * SpecType 
 * Not validated 
 */
type SpecType struct { 
   XMLName xml.Name
   //Inhereting from ParamsType
}

/*
 * EntityReferenceListType 
 * Not validated 
 */
type EntityReferenceListType struct { 
   XMLName xml.Name
   Ref []*EntityReferenceType `xml:"Ref"`
   //Inhereting from ListType
}

/*
 * EntityReferenceType 
 * Not validated 
 */
type EntityReferenceType struct { 
   XMLName xml.Name
   Links *LinkListType `xml:"Links,omitempty"`
   UID UidType `xml:"UID,attr"`
   Name string `xml:"Name,attr"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l EntityReferenceType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * BackupServerEntityListType 
 * Not validated 
 */
type BackupServerEntityListType struct { 
   XMLName xml.Name
   BackupServer []*BackupServerEntityType `xml:"BackupServer"`
   //Inhereting from ListType
}

/*
 * ManagedServerEntityListType 
 * Not validated 
 */
type ManagedServerEntityListType struct { 
   XMLName xml.Name
   ManagedServer []*ManagedServerEntityType `xml:"ManagedServer"`
   //Inhereting from ListType
}

/*
 * BackupServerEntityType 
 * Not validated 
 */
type BackupServerEntityType struct { 
   XMLName xml.Name
   Description string `xml:"Description"`
   Port uint16 `xml:"Port"`
   Version VersionType `xml:"Version"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l BackupServerEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * ManagedServerEntityType 
 * Not validated 
 */
type ManagedServerEntityType struct { 
   XMLName xml.Name
   Description string `xml:"Description"`
   ManagedServerType string `xml:"ManagedServerType"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l ManagedServerEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * BackupServerSpecType 
 * Not validated 
 */
type BackupServerSpecType struct { 
   XMLName xml.Name
   Description string `xml:"Description"`
   DnsNameOrIpAddress string `xml:"DnsNameOrIpAddress"`
   Port uint16 `xml:"Port,omitempty"`
   Username string `xml:"Username"`
   Password string `xml:"Password"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}
func NewBackupServerSpecType() (*BackupServerSpecType) {
  varBackupServerSpecType := BackupServerSpecType{}
  varBackupServerSpecType.XMLName.Local = "BackupServerSpec"
  varBackupServerSpecType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varBackupServerSpecType
}

/*
 * JobSessionEntityType 
 * Not validated 
 */
type JobSessionEntityType struct { 
   XMLName xml.Name
   JobUid UidType `xml:"JobUid"`
   JobName string `xml:"JobName"`
   JobType string `xml:"JobType"`
   CreationTimeUTC DateTime `xml:"CreationTimeUTC"`
   EndTimeUTC DateTime `xml:"EndTimeUTC,omitempty"`
   State string `xml:"State"`
   Result string `xml:"Result"`
   Progress int `xml:"Progress"`
   FailureMessage string `xml:"FailureMessage"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l JobSessionEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * BackupJobSessionEntityType 
 * Not validated 
 */
type BackupJobSessionEntityType struct { 
   XMLName xml.Name
   IsRetry bool `xml:"IsRetry"`
   //Inhereting from JobSessionEntityType
   JobUid UidType `xml:"JobUid"`
   JobName string `xml:"JobName"`
   JobType string `xml:"JobType"`
   CreationTimeUTC DateTime `xml:"CreationTimeUTC"`
   EndTimeUTC DateTime `xml:"EndTimeUTC,omitempty"`
   State string `xml:"State"`
   Result string `xml:"Result"`
   Progress int `xml:"Progress"`
   FailureMessage string `xml:"FailureMessage"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l BackupJobSessionEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * ReplicaJobSessionEntityType 
 * Not validated 
 */
type ReplicaJobSessionEntityType struct { 
   XMLName xml.Name
   IsRetry bool `xml:"IsRetry"`
   //Inhereting from JobSessionEntityType
   JobUid UidType `xml:"JobUid"`
   JobName string `xml:"JobName"`
   JobType string `xml:"JobType"`
   CreationTimeUTC DateTime `xml:"CreationTimeUTC"`
   EndTimeUTC DateTime `xml:"EndTimeUTC,omitempty"`
   State string `xml:"State"`
   Result string `xml:"Result"`
   Progress int `xml:"Progress"`
   FailureMessage string `xml:"FailureMessage"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l ReplicaJobSessionEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * RestoreSessionEntityType 
 * Not validated 
 */
type RestoreSessionEntityType struct { 
   XMLName xml.Name
   RestoredObjRef HierarchyObjRefType `xml:"RestoredObjRef,omitempty"`
   VmDisplayName string `xml:"VmDisplayName,attr"`
   //Inhereting from JobSessionEntityType
   JobUid UidType `xml:"JobUid"`
   JobName string `xml:"JobName"`
   JobType string `xml:"JobType"`
   CreationTimeUTC DateTime `xml:"CreationTimeUTC"`
   EndTimeUTC DateTime `xml:"EndTimeUTC,omitempty"`
   State string `xml:"State"`
   Result string `xml:"Result"`
   Progress int `xml:"Progress"`
   FailureMessage string `xml:"FailureMessage"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l RestoreSessionEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * BackupTaskSessionEntityType 
 * Not validated 
 */
type BackupTaskSessionEntityType struct { 
   XMLName xml.Name
   JobSessionUid UidType `xml:"JobSessionUid"`
   CreationTimeUTC DateTime `xml:"CreationTimeUTC"`
   EndTimeUTC DateTime `xml:"EndTimeUTC,omitempty"`
   State string `xml:"State"`
   Result string `xml:"Result"`
   Reason string `xml:"Reason"`
   TotalSize int64 `xml:"TotalSize"`
   VmDisplayName string `xml:"VmDisplayName,attr"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l BackupTaskSessionEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * BackupTaskSessionEntityListType 
 * Not validated 
 */
type BackupTaskSessionEntityListType struct { 
   XMLName xml.Name
   BackupTaskSession []*BackupTaskSessionEntityType `xml:"BackupTaskSession"`
   //Inhereting from ListType
}

/*
 * ReplicaTaskSessionEntityType 
 * Not validated 
 */
type ReplicaTaskSessionEntityType struct { 
   XMLName xml.Name
   JobSessionUid UidType `xml:"JobSessionUid"`
   CreationTimeUTC DateTime `xml:"CreationTimeUTC"`
   EndTimeUTC DateTime `xml:"EndTimeUTC,omitempty"`
   State string `xml:"State"`
   Result string `xml:"Result"`
   Reason string `xml:"Reason"`
   TotalSize int64 `xml:"TotalSize"`
   VmDisplayName string `xml:"VmDisplayName,attr"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l ReplicaTaskSessionEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * ReplicaTaskSessionEntityListType 
 * Not validated 
 */
type ReplicaTaskSessionEntityListType struct { 
   XMLName xml.Name
   ReplicaTaskSession []*ReplicaTaskSessionEntityType `xml:"ReplicaTaskSession"`
   //Inhereting from ListType
}

/*
 * PlainCredentialsType 
 * Not validated 
 */
type PlainCredentialsType struct { 
   XMLName xml.Name
   UserName string `xml:"UserName"`
   Password string `xml:"Password"`
   //Inhereting from InfoType
}

/*
 * CredentialsInfoType 
 * Not validated 
 */
type CredentialsInfoType struct { 
   XMLName xml.Name
   Id string `xml:"Id"`
   Username string `xml:"Username,omitempty"`
   Description string `xml:"Description,omitempty"`
   Password string `xml:"Password,omitempty"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CredentialsInfoType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * CredentialsInfoListType 
 * Not validated 
 */
type CredentialsInfoListType struct { 
   XMLName xml.Name
   CredentialsInfo []*CredentialsInfoType `xml:"CredentialsInfo"`
   //Inhereting from ListType
}

/*
 * CredentialsInfoSpecType 
 * Not validated 
 */
type CredentialsInfoSpecType struct { 
   XMLName xml.Name
   Username string `xml:"Username,omitempty"`
   Description string `xml:"Description,omitempty"`
   Password string `xml:"Password,omitempty"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}
func NewCredentialsInfoSpecType() (*CredentialsInfoSpecType) {
  varCredentialsInfoSpecType := CredentialsInfoSpecType{}
  varCredentialsInfoSpecType.XMLName.Local = "CredentialsInfoSpec"
  varCredentialsInfoSpecType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varCredentialsInfoSpecType
}

/*
 * PasswordKeyInfoType 
 * Not validated 
 */
type PasswordKeyInfoType struct { 
   XMLName xml.Name
   Id string `xml:"Id"`
   Hint string `xml:"Hint"`
   LastModificationDate DateTime `xml:"LastModificationDate"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l PasswordKeyInfoType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * PasswordKeyInfoListType 
 * Not validated 
 */
type PasswordKeyInfoListType struct { 
   XMLName xml.Name
   PasswordKeyInfo []*PasswordKeyInfoType `xml:"PasswordKeyInfo"`
   //Inhereting from ListType
}

/*
 * PasswordKeyInfoSpecType 
 * Not validated 
 */
type PasswordKeyInfoSpecType struct { 
   XMLName xml.Name
   Hint string `xml:"Hint,omitempty"`
   Password string `xml:"Password,omitempty"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}
func NewPasswordKeyInfoSpecType() (*PasswordKeyInfoSpecType) {
  varPasswordKeyInfoSpecType := PasswordKeyInfoSpecType{}
  varPasswordKeyInfoSpecType.XMLName.Local = "PasswordKeyInfoSpec"
  varPasswordKeyInfoSpecType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varPasswordKeyInfoSpecType
}

/*
 * JobEntityListType 
 * Not validated 
 */
type JobEntityListType struct { 
   XMLName xml.Name
   Job []*JobEntityType `xml:"Job"`
   //Inhereting from ListType
}

/*
 * FailoverPlanEntityListType 
 * Not validated 
 */
type FailoverPlanEntityListType struct { 
   XMLName xml.Name
   FailoverPlan []*FailoverPlanEntityType `xml:"FailoverPlan"`
   //Inhereting from ListType
}

/*
 * JobEntityType 
 * Not validated 
 */
type JobEntityType struct { 
   XMLName xml.Name
   JobType string `xml:"JobType,omitempty"`
   Platform string `xml:"Platform,omitempty"`
   Description string `xml:"Description,omitempty"`
   ScheduleConfigured bool `xml:"ScheduleConfigured,omitempty"`
   ScheduleEnabled bool `xml:"ScheduleEnabled,omitempty"`
   NextRun DateTime `xml:"NextRun,omitempty"`
   JobScheduleOptions *JobScheduleOptionsInfoType `xml:"JobScheduleOptions,omitempty"`
   JobInfo *JobEntityTypeNestedJobInfo `xml:"JobInfo,omitempty"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l JobEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * JobEntityTypeNestedJobInfo 
 * Not validated 
 */
type JobEntityTypeNestedJobInfo struct { 
   XMLName xml.Name
   BackupJobInfo *BackupJobInfoType `xml:"BackupJobInfo,omitempty"`
   FileCopyJobInfo *FileCopyJobInfoType `xml:"FileCopyJobInfo,omitempty"`
   ReplicaJobInfo *ReplicaJobInfoType `xml:"ReplicaJobInfo,omitempty"`
}

/*
 * FailoverPlanEntityType 
 * Not validated 
 */
type FailoverPlanEntityType struct { 
   XMLName xml.Name
   Description string `xml:"Description,omitempty"`
   FailoverPlanInfo *FailoverPlanInfoType `xml:"FailoverPlanInfo"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l FailoverPlanEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * JobManagementSpecType 
 * Not validated 
 */
type JobManagementSpecType struct { 
   XMLName xml.Name
   Credentials *PlainCredentialsType `xml:"Credentials"`
   Force bool `xml:"Force"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}
func NewJobManagementSpecType() (*JobManagementSpecType) {
  varJobManagementSpecType := JobManagementSpecType{}
  varJobManagementSpecType.XMLName.Local = "JobManagementSpec"
  varJobManagementSpecType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varJobManagementSpecType
}

/*
 * FailoverPlanManagementSpecType 
 * Not validated 
 */
type FailoverPlanManagementSpecType struct { 
   XMLName xml.Name
   StartNow bool `xml:"StartNow"`
   StartDate DateTime `xml:"StartDate"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}
func NewFailoverPlanManagementSpecType() (*FailoverPlanManagementSpecType) {
  varFailoverPlanManagementSpecType := FailoverPlanManagementSpecType{}
  varFailoverPlanManagementSpecType.XMLName.Local = "FailoverPlanManagementSpec"
  varFailoverPlanManagementSpecType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varFailoverPlanManagementSpecType
}

/*
 * BackupJobInfoType 
 * Not validated 
 */
type BackupJobInfoType struct { 
   XMLName xml.Name
   Includes *ObjectInJobListType `xml:"Includes,omitempty"`
   GuestProcessingOptions *GuestProcessingOptionsType `xml:"GuestProcessingOptions,omitempty"`
   AdvancedStorageOptions *AdvancedStorageOptionsType `xml:"AdvancedStorageOptions,omitempty"`
   //Inhereting from InfoType
}

/*
 * FailoverPlanInfoType 
 * Not validated 
 */
type FailoverPlanInfoType struct { 
   XMLName xml.Name
   Includes *FailoveredVmListType `xml:"Includes,omitempty"`
   //Inhereting from InfoType
}

/*
 * ObjectInJobListType 
 * Not validated 
 */
type ObjectInJobListType struct { 
   XMLName xml.Name
   ObjectInJob []*ObjectInJobType `xml:"ObjectInJob"`
   //Inhereting from ListType
}

/*
 * FailoveredVmListType 
 * Not validated 
 */
type FailoveredVmListType struct { 
   XMLName xml.Name
   FailoveredVm []*FailoveredVmType `xml:"FailoveredVm"`
   //Inhereting from ListType
}

/*
 * ObjectInJobType 
 * Not validated 
 */
type ObjectInJobType struct { 
   XMLName xml.Name
   ObjectInJobId string `xml:"ObjectInJobId,omitempty"`
   HierarchyObjRef HierarchyObjRefType `xml:"HierarchyObjRef,omitempty"`
   Name string `xml:"Name,omitempty"`
   DisplayName string `xml:"DisplayName"`
   Order int `xml:"Order,omitempty"`
   GuestProcessingOptions *GuestProcessingOptionsType `xml:"GuestProcessingOptions"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l ObjectInJobType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * FailoveredVmType 
 * Not validated 
 */
type FailoveredVmType struct { 
   XMLName xml.Name
   FailoveredVmId string `xml:"FailoveredVmId,omitempty"`
   HierarchyObjRef HierarchyObjRefType `xml:"HierarchyObjRef,omitempty"`
   Name string `xml:"Name,omitempty"`
   DisplayName string `xml:"DisplayName"`
   Order int `xml:"Order,omitempty"`
   GuestProcessingOptions *GuestProcessingOptionsType `xml:"GuestProcessingOptions"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l FailoveredVmType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * CreateObjectInJobSpecType 
 * Not validated 
 */
type CreateObjectInJobSpecType struct { 
   XMLName xml.Name
   HierarchyObjRef HierarchyObjRefType `xml:"HierarchyObjRef,omitempty"`
   HierarchyObjName string `xml:"HierarchyObjName,omitempty"`
   Order int `xml:"Order,omitempty"`
   GuestProcessingOptions *GuestProcessingOptionsType `xml:"GuestProcessingOptions,omitempty"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}
func NewCreateObjectInJobSpecType() (*CreateObjectInJobSpecType) {
  varCreateObjectInJobSpecType := CreateObjectInJobSpecType{}
  varCreateObjectInJobSpecType.XMLName.Local = "CreateObjectInJobSpec"
  varCreateObjectInJobSpecType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varCreateObjectInJobSpecType
}

/*
 * JobItemVssOptionsType 
 * Not validated 
 */
type JobItemVssOptionsType struct { 
   XMLName xml.Name
   Enabled bool `xml:"Enabled"`
   IgnoreErrors bool `xml:"IgnoreErrors"`
   GuestFSIndexingType string `xml:"GuestFSIndexingType"`
   Credentials *PlainCredentialsType `xml:"Credentials"`
   TransactionLogsTruncation string `xml:"TransactionLogsTruncation"`
   IsFirstUsage bool `xml:"IsFirstUsage"`
   IncludedIndexingFolders *StringListType `xml:"IncludedIndexingFolders"`
   ExcludedIndexingFolders *StringListType `xml:"ExcludedIndexingFolders"`
   CredentialsId string `xml:"CredentialsId"`
   //Inhereting from InfoType
}

/*
 * GuestProcessingOptionsType 
 * Not validated 
 */
type GuestProcessingOptionsType struct { 
   XMLName xml.Name
   AppAwareProcessingMode string `xml:"AppAwareProcessingMode,omitempty"`
   FileSystemIndexingMode string `xml:"FileSystemIndexingMode,omitempty"`
   IncludedIndexingFolders *StringListType `xml:"IncludedIndexingFolders,omitempty"`
   ExcludedIndexingFolders *StringListType `xml:"ExcludedIndexingFolders,omitempty"`
   CredentialsId string `xml:"CredentialsId,omitempty"`
   VssSnapshotOptions *VssSnapshotOptionsType `xml:"VssSnapshotOptions,omitempty"`
   WindowsGuestFSIndexingOptions *WindowsGuestFSIndexingOptionsType `xml:"WindowsGuestFSIndexingOptions,omitempty"`
   LinuxGuestFSIndexingOptions *LinuxGuestFSIndexingOptionsType `xml:"LinuxGuestFSIndexingOptions,omitempty"`
   SqlBackupOptions *SqlBackupOptionsType `xml:"SqlBackupOptions,omitempty"`
   WindowsCredentialsId string `xml:"WindowsCredentialsId,omitempty"`
   LinuxCredentialsId string `xml:"LinuxCredentialsId,omitempty"`
   FSFileExcludeOptions *FSFileExcludeOptionsType `xml:"FSFileExcludeOptions,omitempty"`
   OracleBackupOptions *OracleBackupOptionsType `xml:"OracleBackupOptions,omitempty"`
   //Inhereting from InfoType
}

/*
 * VssSnapshotOptionsType 
 * Not validated 
 */
type VssSnapshotOptionsType struct { 
   XMLName xml.Name
   VssSnapshotMode string `xml:"VssSnapshotMode"`
   IsCopyOnly bool `xml:"IsCopyOnly,omitempty"`
   //Inhereting from InfoType
}

/*
 * WindowsGuestFSIndexingOptionsType 
 * Not validated 
 */
type WindowsGuestFSIndexingOptionsType struct { 
   XMLName xml.Name
   FileSystemIndexingMode string `xml:"FileSystemIndexingMode"`
   IncludedIndexingFolders *StringListType `xml:"IncludedIndexingFolders"`
   ExcludedIndexingFolders *StringListType `xml:"ExcludedIndexingFolders"`
   //Inhereting from InfoType
}

/*
 * LinuxGuestFSIndexingOptionsType 
 * Not validated 
 */
type LinuxGuestFSIndexingOptionsType struct { 
   XMLName xml.Name
   FileSystemIndexingMode string `xml:"FileSystemIndexingMode"`
   IncludedIndexingFolders *StringListType `xml:"IncludedIndexingFolders"`
   ExcludedIndexingFolders *StringListType `xml:"ExcludedIndexingFolders"`
   //Inhereting from InfoType
}

/*
 * SqlBackupOptionsType 
 * Not validated 
 */
type SqlBackupOptionsType struct { 
   XMLName xml.Name
   TransactionLogsProcessing string `xml:"TransactionLogsProcessing"`
   BackupLogsFrequencyMin int `xml:"BackupLogsFrequencyMin,omitempty"`
   UseDbBackupRetention bool `xml:"UseDbBackupRetention,omitempty"`
   RetainDays int `xml:"RetainDays,omitempty"`
   //Inhereting from InfoType
}

/*
 * AdvancedStorageOptionsType 
 * Not validated 
 */
type AdvancedStorageOptionsType struct { 
   XMLName xml.Name
   PasswordKeyId string `xml:"PasswordKeyId,omitempty"`
   //Inhereting from InfoType
}

/*
 * FSFileExcludeOptionsType 
 * Not validated 
 */
type FSFileExcludeOptionsType struct { 
   XMLName xml.Name
   BackupScope int `xml:"BackupScope"`
   IncludeList *StringListType `xml:"IncludeList"`
   ExcludeList *StringListType `xml:"ExcludeList"`
   //Inhereting from InfoType
}

/*
 * OracleBackupOptionsType 
 * Not validated 
 */
type OracleBackupOptionsType struct { 
   XMLName xml.Name
   BackupLogsEnabled bool `xml:"BackupLogsEnabled"`
   BackupLogsFrequencyMin int `xml:"BackupLogsFrequencyMin,omitempty"`
   UseDbBackupRetention bool `xml:"UseDbBackupRetention,omitempty"`
   RetainDays int `xml:"RetainDays,omitempty"`
   ArchivedLogsTruncation string `xml:"ArchivedLogsTruncation"`
   ArchivedLogsMaxAgeHours int `xml:"ArchivedLogsMaxAgeHours,omitempty"`
   ArchivedLogsMaxSizeMb int `xml:"ArchivedLogsMaxSizeMb,omitempty"`
   SysdbaCredsId string `xml:"SysdbaCredsId"`
   ProxyAutoSelect bool `xml:"ProxyAutoSelect"`
   //Inhereting from InfoType
}

/*
 * StringListType 
 * Not validated 
 */
type StringListType struct { 
   XMLName xml.Name
   Path []string `xml:"Path"`
}

/*
 * JobScheduleOptionsInfoType 
 * Not validated 
 */
type JobScheduleOptionsInfoType struct { 
   XMLName xml.Name
   RetryOptions *JobScheduleRetryOptionsType `xml:"RetryOptions,omitempty"`
   WaitForBackupCompletion bool `xml:"WaitForBackupCompletion,omitempty"`
   BackupCompetitionWaitingPeriodMin int `xml:"BackupCompetitionWaitingPeriodMin,omitempty"`
   OptionsDaily *JobScheduleDailyOptionsType `xml:"OptionsDaily,omitempty"`
   OptionsMonthly *JobScheduleMonthlyOptionsType `xml:"OptionsMonthly,omitempty"`
   OptionsPeriodically *JobSchedulePeriodicallyOptionsType `xml:"OptionsPeriodically,omitempty"`
   OptionsContinuous *JobScheduleContinuousOptionsType `xml:"OptionsContinuous,omitempty"`
   OptionsBackupWindow *JobScheduleBackupWindowOptionsType `xml:"OptionsBackupWindow,omitempty"`
   OptionsDaisyChaining *JobScheduleDaisyChainingOptionsType `xml:"OptionsDaisyChaining,omitempty"`
   //Inhereting from InfoType
}

/*
 * JobScheduleRetryOptionsType 
 * Not validated 
 */
type JobScheduleRetryOptionsType struct { 
   XMLName xml.Name
   RetryTimes int `xml:"RetryTimes,omitempty"`
   RetryTimeout int `xml:"RetryTimeout,omitempty"`
   RetrySpecified bool `xml:"RetrySpecified,omitempty"`
   //Inhereting from InfoType
}

/*
 * JobScheduleDailyOptionsType 
 * Not validated 
 */
type JobScheduleDailyOptionsType struct { 
   XMLName xml.Name
   Kind string `xml:"Kind,omitempty"`
   Days []DaysOfWeekEnumeration `xml:"Days"`
   Time Time `xml:"Time,omitempty"`
   TimeOffsetUtc int `xml:"TimeOffsetUtc,omitempty"`
   Enabled bool `xml:"Enabled,attr"`
   //Inhereting from InfoType
}

/*
 * JobScheduleMonthlyOptionsType 
 * Not validated 
 */
type JobScheduleMonthlyOptionsType struct { 
   XMLName xml.Name
   Time Time `xml:"Time,omitempty"`
   TimeOffsetUtc int `xml:"TimeOffsetUtc,omitempty"`
   DayNumberInMonth string `xml:"DayNumberInMonth,omitempty"`
   DayOfWeek string `xml:"DayOfWeek,omitempty"`
   Months []JobScheduleMonthEnumeration `xml:"Months"`
   DayOfMonth int `xml:"DayOfMonth,omitempty"`
   Enabled bool `xml:"Enabled,attr"`
   //Inhereting from InfoType
}

/*
 * JobSchedulePeriodicallyOptionsType 
 * Not validated 
 */
type JobSchedulePeriodicallyOptionsType struct { 
   XMLName xml.Name
   Kind string `xml:"Kind,omitempty"`
   FullPeriod int `xml:"FullPeriod,omitempty"`
   Schedule *TimePeriodsType `xml:"Schedule,omitempty"`
   Enabled bool `xml:"Enabled,attr"`
   //Inhereting from InfoType
}

/*
 * TimePeriodsType 
 * Not validated 
 */
type TimePeriodsType struct { 
   XMLName xml.Name
   Day []*TimePeriodsTypeNestedDay `xml:"Day"`
   //Inhereting from InfoType
}

/*
 * TimePeriodsTypeNestedDay 
 * Not validated 
 */
type TimePeriodsTypeNestedDay struct { 
   XMLName xml.Name
   Name string `xml:"Name,attr"`
}

/*
 * JobScheduleContinuousOptionsType 
 * Not validated 
 */
type JobScheduleContinuousOptionsType struct { 
   XMLName xml.Name
   Enabled bool `xml:"Enabled,attr"`
   //Inhereting from InfoType
}

/*
 * JobScheduleBackupWindowOptionsType 
 * Not validated 
 */
type JobScheduleBackupWindowOptionsType struct { 
   XMLName xml.Name
   TimePeriods *TimePeriodsType `xml:"TimePeriods"`
   Enabled bool `xml:"Enabled,attr"`
   //Inhereting from InfoType
}

/*
 * JobScheduleDaisyChainingOptionsType 
 * Not validated 
 */
type JobScheduleDaisyChainingOptionsType struct { 
   XMLName xml.Name
   PreviousJobUid UidType `xml:"PreviousJobUid"`
   Enabled bool `xml:"Enabled,attr"`
   //Inhereting from InfoType
}

/*
 * FileCopyJobInfoType 
 * Not validated 
 */
type FileCopyJobInfoType struct { 
   XMLName xml.Name
   //Inhereting from InfoType
}

/*
 * ReplicaJobInfoType 
 * Not validated 
 */
type ReplicaJobInfoType struct { 
   XMLName xml.Name
   Includes *ObjectInJobListType `xml:"Includes,omitempty"`
   GuestProcessingOptions *GuestProcessingOptionsType `xml:"GuestProcessingOptions,omitempty"`
   //Inhereting from InfoType
}

/*
 * TaskType 
 * Not validated 
 */
type TaskType struct { 
   XMLName xml.Name
   TaskId string `xml:"TaskId,omitempty"`
   State string `xml:"State,omitempty"`
   Operation string `xml:"Operation,omitempty"`
   Result *TaskResultInfoType `xml:"Result,omitempty"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l TaskType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * TaskListType 
 * Not validated 
 */
type TaskListType struct { 
   XMLName xml.Name
   Task []*TaskType `xml:"Task"`
   //Inhereting from ListType
}

/*
 * QueryResultType 
 * Not validated 
 */
type QueryResultType struct { 
   XMLName xml.Name
   Refs *EntityReferenceListType `xml:"Refs,omitempty"`
   Entities *EntitiesType `xml:"Entities,omitempty"`
   Resources *ResourcesType `xml:"Resources,omitempty"`
   Links *LinkListType `xml:"Links,omitempty"`
   PagingInfo *PagingInfoType `xml:"PagingInfo,omitempty"`
}
func (l QueryResultType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * PagingInfoType 
 * Not validated 
 */
type PagingInfoType struct { 
   XMLName xml.Name
   Links *LinkListType `xml:"Links,omitempty"`
   PageNum int `xml:"PageNum,attr"`
   PageSize int `xml:"PageSize,attr"`
   PagesCount int `xml:"PagesCount,attr"`
   //Inhereting from InfoType
}
func (l PagingInfoType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * BackupEntityType 
 * Not validated 
 */
type BackupEntityType struct { 
   XMLName xml.Name
   Ref *EntityReferenceType `xml:"Ref"`
   Platform string `xml:"Platform,omitempty"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l BackupEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * BackupEntityListType 
 * Not validated 
 */
type BackupEntityListType struct { 
   XMLName xml.Name
   Backup []*BackupEntityType `xml:"Backup"`
   //Inhereting from ListType
}

/*
 * ReplicaEntityType 
 * Not validated 
 */
type ReplicaEntityType struct { 
   XMLName xml.Name
   Ref *EntityReferenceType `xml:"Ref"`
   Platform string `xml:"Platform,omitempty"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l ReplicaEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * ReplicaEntityListType 
 * Not validated 
 */
type ReplicaEntityListType struct { 
   XMLName xml.Name
   Replica []*ReplicaEntityType `xml:"Replica"`
   //Inhereting from ListType
}

/*
 * RestorePointEntityType 
 * Not validated 
 */
type RestorePointEntityType struct { 
   XMLName xml.Name
   Ref *EntityReferenceType `xml:"Ref"`
   BackupDateUTC DateTime `xml:"BackupDateUTC"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l RestorePointEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * RestorePointEntityListType 
 * Not validated 
 */
type RestorePointEntityListType struct { 
   XMLName xml.Name
   RestorePoint []*RestorePointEntityType `xml:"RestorePoint"`
   //Inhereting from ListType
}

/*
 * VmRestorePointEntityType 
 * Not validated 
 */
type VmRestorePointEntityType struct { 
   XMLName xml.Name
   CreationTimeUTC DateTime `xml:"CreationTimeUTC"`
   VmName string `xml:"VmName"`
   Algorithm string `xml:"Algorithm"`
   PointType string `xml:"PointType"`
   HierarchyObjRef HierarchyObjRefType `xml:"HierarchyObjRef,omitempty"`
   SqlInfo *VmRestorePointSqlInfoType `xml:"SqlInfo,omitempty"`
   VmDisplayName string `xml:"VmDisplayName,attr"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l VmRestorePointEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * VAppRestorePointEntityType 
 * Not validated 
 */
type VAppRestorePointEntityType struct { 
   XMLName xml.Name
   CreationTimeUTC DateTime `xml:"CreationTimeUTC"`
   VAppName string `xml:"VAppName"`
   Algorithm string `xml:"Algorithm"`
   PointType string `xml:"PointType"`
   HierarchyObjRef HierarchyObjRefType `xml:"HierarchyObjRef,omitempty"`
   VAppDisplayName string `xml:"VAppDisplayName,attr"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l VAppRestorePointEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * VmRestorePointEntityListType 
 * Not validated 
 */
type VmRestorePointEntityListType struct { 
   XMLName xml.Name
   VmRestorePoint []*VmRestorePointEntityType `xml:"VmRestorePoint"`
   //Inhereting from ListType
}

/*
 * VAppRestorePointEntityListType 
 * Not validated 
 */
type VAppRestorePointEntityListType struct { 
   XMLName xml.Name
   VAppRestorePoint []*VAppRestorePointEntityType `xml:"VAppRestorePoint"`
   //Inhereting from ListType
}

/*
 * VmReplicaPointEntityType 
 * Not validated 
 */
type VmReplicaPointEntityType struct { 
   XMLName xml.Name
   CreationTimeUTC DateTime `xml:"CreationTimeUTC"`
   VmName string `xml:"VmName"`
   Algorithm string `xml:"Algorithm"`
   PointType string `xml:"PointType"`
   VmDisplayName string `xml:"VmDisplayName,attr"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l VmReplicaPointEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * VmReplicaPointEntityListType 
 * Not validated 
 */
type VmReplicaPointEntityListType struct { 
   XMLName xml.Name
   VmRestorePoint []*VmReplicaPointEntityType `xml:"VmRestorePoint"`
   //Inhereting from ListType
}

/*
 * VmRestorePointMountType 
 * Not validated 
 */
type VmRestorePointMountType struct { 
   XMLName xml.Name
   FSRoots *DirectoryEntryListType `xml:"FSRoots"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l VmRestorePointMountType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * VmRestorePointMountListType 
 * Not validated 
 */
type VmRestorePointMountListType struct { 
   XMLName xml.Name
   VmRestorePointMount []*VmRestorePointMountType `xml:"VmRestorePointMount"`
   //Inhereting from ListType
}

/*
 * VmReplicaPointMountType 
 * Not validated 
 */
type VmReplicaPointMountType struct { 
   XMLName xml.Name
   FSRoots *DirectoryEntryListType `xml:"FSRoots"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l VmReplicaPointMountType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * VmReplicaPointMountListType 
 * Not validated 
 */
type VmReplicaPointMountListType struct { 
   XMLName xml.Name
   VmReplicaPointMount []*VmReplicaPointMountType `xml:"VmReplicaPointMount"`
   //Inhereting from ListType
}

/*
 * CatalogVmEntityType 
 * Not validated 
 */
type CatalogVmEntityType struct { 
   XMLName xml.Name
   VmDisplayName string `xml:"VmDisplayName,attr"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CatalogVmEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * CatalogVmEntityListType 
 * Not validated 
 */
type CatalogVmEntityListType struct { 
   XMLName xml.Name
   CatalogVm []*CatalogVmEntityType `xml:"CatalogVm"`
   //Inhereting from ListType
}

/*
 * CatalogVmRestorePointEntityType 
 * Not validated 
 */
type CatalogVmRestorePointEntityType struct { 
   XMLName xml.Name
   BackupDateUTC DateTime `xml:"BackupDateUTC"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CatalogVmRestorePointEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * CatalogVmRestorePointEntityListType 
 * Not validated 
 */
type CatalogVmRestorePointEntityListType struct { 
   XMLName xml.Name
   CatalogVmRestorePoint []*CatalogVmRestorePointEntityType `xml:"CatalogVmRestorePoint"`
   //Inhereting from ListType
}

/*
 * FileSystemEntryType 
 * Not validated 
 */
type FileSystemEntryType struct { 
   XMLName xml.Name
   FileEntry *FileEntryType `xml:"FileEntry"`
   DirectoryEntry *DirectoryEntryType `xml:"DirectoryEntry"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l FileSystemEntryType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * DirectoryEntryType 
 * Not validated 
 */
type DirectoryEntryType struct { 
   XMLName xml.Name
   Path string `xml:"Path"`
   Name string `xml:"Name"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l DirectoryEntryType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * FileEntryType 
 * Not validated 
 */
type FileEntryType struct { 
   XMLName xml.Name
   Path string `xml:"Path"`
   Name string `xml:"Name"`
   Size int64 `xml:"Size"`
   Owner string `xml:"Owner"`
   ModifiedDateUTC DateTime `xml:"ModifiedDateUTC"`
   Actions *LinkListType `xml:"Actions"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l FileEntryType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * FileEntryListType 
 * Not validated 
 */
type FileEntryListType struct { 
   XMLName xml.Name
   FileEntry []*FileEntryType `xml:"FileEntry"`
   //Inhereting from ListType
}

/*
 * HierarchyRootEntityType 
 * Not validated 
 */
type HierarchyRootEntityType struct { 
   XMLName xml.Name
   HierarchyRootId string `xml:"HierarchyRootId,omitempty"`
   UniqueId string `xml:"UniqueId,omitempty"`
   HostType string `xml:"HostType,omitempty"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l HierarchyRootEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * RepositoryEntityType 
 * Not validated 
 */
type RepositoryEntityType struct { 
   XMLName xml.Name
   Capacity int64 `xml:"Capacity"`
   FreeSpace int64 `xml:"FreeSpace"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l RepositoryEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * HierarchyRootEntityListType 
 * Not validated 
 */
type HierarchyRootEntityListType struct { 
   XMLName xml.Name
   HierarchyRoot []*HierarchyRootEntityType `xml:"HierarchyRoot"`
   //Inhereting from ListType
}

/*
 * RepositoryEntityListType 
 * Not validated 
 */
type RepositoryEntityListType struct { 
   XMLName xml.Name
   Repository []*RepositoryEntityType `xml:"Repository"`
   //Inhereting from ListType
}

/*
 * DirectoryEntryListType 
 * Not validated 
 */
type DirectoryEntryListType struct { 
   XMLName xml.Name
   DirectoryEntry []*DirectoryEntryType `xml:"DirectoryEntry"`
   //Inhereting from ListType
}

/*
 * FileSystemEntriesType 
 * Not validated 
 */
type FileSystemEntriesType struct { 
   XMLName xml.Name
   Files *FileEntryListType `xml:"Files,omitempty"`
   Directories *DirectoryEntryListType `xml:"Directories,omitempty"`
   PagingInfo *PagingInfoType `xml:"PagingInfo,omitempty"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l FileSystemEntriesType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * HierarchyItemListType 
 * Not validated 
 */
type HierarchyItemListType struct { 
   XMLName xml.Name
   HierarchyItem []*HierarchyItemType `xml:"HierarchyItem"`
   //Inhereting from ListType
}

/*
 * BackupJobSessionEntityListType 
 * Not validated 
 */
type BackupJobSessionEntityListType struct { 
   XMLName xml.Name
   BackupJobSession []*BackupJobSessionEntityType `xml:"BackupJobSession"`
   //Inhereting from ListType
}

/*
 * ReplicaJobSessionEntityListType 
 * Not validated 
 */
type ReplicaJobSessionEntityListType struct { 
   XMLName xml.Name
   ReplicaJobSession []*ReplicaJobSessionEntityType `xml:"ReplicaJobSession"`
   //Inhereting from ListType
}

/*
 * RestoreSessionEntityListType 
 * Not validated 
 */
type RestoreSessionEntityListType struct { 
   XMLName xml.Name
   RestoreSession []*RestoreSessionEntityType `xml:"RestoreSession"`
   //Inhereting from ListType
}

/*
 * FileRestoreSpecType 
 * Not validated 
 */
type FileRestoreSpecType struct { 
   XMLName xml.Name
   ToOriginalLocation *FileRestoreSpecTypeNestedToOriginalLocation `xml:"ToOriginalLocation"`
   ForDirectDownload *FileRestoreSpecTypeNestedForDirectDownload `xml:"ForDirectDownload"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}
func NewFileRestoreSpecType() (*FileRestoreSpecType) {
  varFileRestoreSpecType := FileRestoreSpecType{}
  varFileRestoreSpecType.XMLName.Local = "FileRestoreSpec"
  varFileRestoreSpecType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varFileRestoreSpecType
}

/*
 * FileRestoreSpecTypeNestedToOriginalLocation 
 * Not validated 
 */
type FileRestoreSpecTypeNestedToOriginalLocation struct { 
   XMLName xml.Name
}

/*
 * FileRestoreSpecTypeNestedForDirectDownload 
 * Not validated 
 */
type FileRestoreSpecTypeNestedForDirectDownload struct { 
   XMLName xml.Name
}

/*
 * RestoreSpecType 
 * Not validated 
 */
type RestoreSpecType struct { 
   XMLName xml.Name
   VmRestoreSpec *VmRestoreSpecInfoType `xml:"VmRestoreSpec"`
   vCloudVmRestoreSpec *vCloudVmRestoreSpecInfoType `xml:"vCloudVmRestoreSpec"`
   vCloudVAppRestoreSpec *vCloudVAppRestoreSpecInfoType `xml:"vCloudVAppRestoreSpec"`
   SqlItemRestoreSpec *SqlItemRestoreSpecInfoType `xml:"SqlItemRestoreSpec"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}
func NewRestoreSpecType() (*RestoreSpecType) {
  varRestoreSpecType := RestoreSpecType{}
  varRestoreSpecType.XMLName.Local = "RestoreSpec"
  varRestoreSpecType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varRestoreSpecType
}

/*
 * VmRestoreSpecInfoType 
 * Not validated 
 */
type VmRestoreSpecInfoType struct { 
   XMLName xml.Name
   PowerOnAfterRestore bool `xml:"PowerOnAfterRestore,omitempty"`
   QuickRollback bool `xml:"QuickRollback,omitempty"`
   //Inhereting from InfoType
}

/*
 * vCloudVmRestoreSpecInfoType 
 * Not validated 
 */
type vCloudVmRestoreSpecInfoType struct { 
   XMLName xml.Name
   PowerOnAfterRestore bool `xml:"PowerOnAfterRestore,omitempty"`
   HierarchyRootUid UidType `xml:"HierarchyRootUid,omitempty"`
   vAppRef HierarchyObjRefType `xml:"vAppRef,omitempty"`
   VmRestoreParameters *vCloudVmRestoreParametersInfoType `xml:"VmRestoreParameters,omitempty"`
   //Inhereting from InfoType
}

/*
 * vCloudVAppRestoreSpecInfoType 
 * Not validated 
 */
type vCloudVAppRestoreSpecInfoType struct { 
   XMLName xml.Name
   PowerOnAfterRestore bool `xml:"PowerOnAfterRestore,omitempty"`
   HierarchyRootUid UidType `xml:"HierarchyRootUid,omitempty"`
   OrgVdcRef HierarchyObjRefType `xml:"OrgVdcRef,omitempty"`
   vAppNewName string `xml:"vAppNewName,omitempty"`
   VmsRestoreParameters *vCloudVAppRestoreSpecInfoTypeNestedVmsRestoreParameters `xml:"VmsRestoreParameters"`
   //Inhereting from InfoType
}

/*
 * vCloudVAppRestoreSpecInfoTypeNestedVmsRestoreParameters 
 * Not validated 
 */
type vCloudVAppRestoreSpecInfoTypeNestedVmsRestoreParameters struct { 
   XMLName xml.Name
   Vm []*vCloudVmRestoreParametersInfoType `xml:"Vm"`
}

/*
 * vCloudVmRestoreParametersInfoType 
 * Not validated 
 */
type vCloudVmRestoreParametersInfoType struct { 
   XMLName xml.Name
   VmRestorePointUid UidType `xml:"VmRestorePointUid,omitempty"`
   VmNewName string `xml:"VmNewName,omitempty"`
   DatastoreRef HierarchyObjRefType `xml:"DatastoreRef,omitempty"`
   OrgVdcStorageProfileRef HierarchyObjRefType `xml:"OrgVdcStorageProfileRef,omitempty"`
   LinkedCloneVmTemplateRef HierarchyObjRefType `xml:"LinkedCloneVmTemplateRef,omitempty"`
   //Inhereting from InfoType
}

/*
 * BackupJobCloneInfoType 
 * Not validated 
 */
type BackupJobCloneInfoType struct { 
   XMLName xml.Name
   JobName string `xml:"JobName,omitempty"`
   FolderName string `xml:"FolderName,omitempty"`
   RepositoryUid UidType `xml:"RepositoryUid,omitempty"`
   //Inhereting from InfoType
}

/*
 * ReplicaJobCloneInfoType 
 * Not validated 
 */
type ReplicaJobCloneInfoType struct { 
   XMLName xml.Name
   JobName string `xml:"JobName,omitempty"`
   VmSuffix string `xml:"VmSuffix,omitempty"`
   //Inhereting from InfoType
}

/*
 * JobCloneSpecType 
 * Not validated 
 */
type JobCloneSpecType struct { 
   XMLName xml.Name
   BackupJobCloneInfo *BackupJobCloneInfoType `xml:"BackupJobCloneInfo,omitempty"`
   ReplicaJobCloneInfo *ReplicaJobCloneInfoType `xml:"ReplicaJobCloneInfo,omitempty"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}
func NewJobCloneSpecType() (*JobCloneSpecType) {
  varJobCloneSpecType := JobCloneSpecType{}
  varJobCloneSpecType.XMLName.Local = "JobCloneSpec"
  varJobCloneSpecType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varJobCloneSpecType
}

/*
 * TaskResultInfoType 
 * Not validated 
 */
type TaskResultInfoType struct { 
   XMLName xml.Name
   Message string `xml:"Message,omitempty"`
   Success bool `xml:"Success,attr"`
   CredentialsNeeded bool `xml:"CredentialsNeeded,attr"`
   ConfirmationNeeded bool `xml:"ConfirmationNeeded,attr"`
   //Inhereting from InfoType
}

/*
 * ErrorType 
 * Not validated 
 */
type ErrorType struct { 
   XMLName xml.Name
   FirstChanceExceptionMessage string `xml:"FirstChanceExceptionMessage,omitempty"`
   StackTrace string `xml:"StackTrace,omitempty"`
   Message string `xml:"Message,attr"`
   StatusCode int `xml:"StatusCode,attr"`
   Status string `xml:"Status,attr"`
   //Inhereting from InfoType
}

/*
 * HierarchyItemType 
 * Not validated 
 */
type HierarchyItemType struct { 
   XMLName xml.Name
   ObjectRef HierarchyObjRefType `xml:"ObjectRef,omitempty"`
   ObjectType string `xml:"ObjectType,omitempty"`
   ObjectName string `xml:"ObjectName,omitempty"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l HierarchyItemType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * QuerySvcType 
 * Not validated 
 */
type QuerySvcType struct { 
   XMLName xml.Name
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l QuerySvcType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * LookupSvcType 
 * Not validated 
 */
type LookupSvcType struct { 
   XMLName xml.Name
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l LookupSvcType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * ReportingSvcType 
 * Not validated 
 */
type ReportingSvcType struct { 
   XMLName xml.Name
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l ReportingSvcType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * EnterpriseRoleEntityType 
 * Not validated 
 */
type EnterpriseRoleEntityType struct { 
   XMLName xml.Name
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l EnterpriseRoleEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * EnterpriseRoleEntityListType 
 * Not validated 
 */
type EnterpriseRoleEntityListType struct { 
   XMLName xml.Name
   EnterpriseRoleEntity []*EnterpriseRoleEntityType `xml:"EnterpriseRoleEntity"`
   //Inhereting from ListType
}

/*
 * EnterpriseAccountEntityType 
 * Not validated 
 */
type EnterpriseAccountEntityType struct { 
   XMLName xml.Name
   AccountType AccountTypeEnumeration `xml:"AccountType"`
   Roles *EnterpriseRoleEntityListType `xml:"Roles"`
   AllowRestoreAllVms bool `xml:"AllowRestoreAllVms"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l EnterpriseAccountEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * EnterpriseAccountEntityListType 
 * Not validated 
 */
type EnterpriseAccountEntityListType struct { 
   XMLName xml.Name
   EnterpriseAccountEntity []*EnterpriseAccountEntityType `xml:"EnterpriseAccountEntity"`
   //Inhereting from ListType
}

/*
 * EnterpriseAccountHierarchyScopeType 
 * Not validated 
 */
type EnterpriseAccountHierarchyScopeType struct { 
   XMLName xml.Name
   Name string `xml:"Name"`
   HierarchyRootName string `xml:"HierarchyRootName"`
   Platform string `xml:"Platform"`
   HierarchyObjectType string `xml:"HierarchyObjectType"`
   State string `xml:"State"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l EnterpriseAccountHierarchyScopeType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * EnterpriseAccountHierarchyScopeListType 
 * Not validated 
 */
type EnterpriseAccountHierarchyScopeListType struct { 
   XMLName xml.Name
   EnterpriseAccountHierarchyScope []*EnterpriseAccountHierarchyScopeType `xml:"EnterpriseAccountHierarchyScope"`
   //Inhereting from ListType
}

/*
 * EnterpriseAccountInRoleType 
 * Not validated 
 */
type EnterpriseAccountInRoleType struct { 
   XMLName xml.Name
   RoleName string `xml:"RoleName"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l EnterpriseAccountInRoleType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * EnterpriseAccountInRoleListType 
 * Not validated 
 */
type EnterpriseAccountInRoleListType struct { 
   XMLName xml.Name
   EnterpriseAccountInRole []*EnterpriseAccountInRoleType `xml:"EnterpriseAccountInRole"`
   //Inhereting from ListType
}

/*
 * HierarchyScopeCreateSpecType 
 * Not validated 
 */
type HierarchyScopeCreateSpecType struct { 
   XMLName xml.Name
   HierarchyScopeItem []*HierarchyScopeCreateSpecItemType `xml:"HierarchyScopeItem"`
   //Inhereting from ListType
}

/*
 * HierarchyScopeCreateSpecItemType 
 * Not validated 
 */
type HierarchyScopeCreateSpecItemType struct { 
   XMLName xml.Name
   HierarchyObjRef HierarchyObjRefType `xml:"HierarchyObjRef"`
   ObjectName string `xml:"ObjectName"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}
func NewHierarchyScopeCreateSpecItemType() (*HierarchyScopeCreateSpecItemType) {
  varHierarchyScopeCreateSpecItemType := HierarchyScopeCreateSpecItemType{}
  varHierarchyScopeCreateSpecItemType.XMLName.Local = "HierarchyScopeCreateSpecItem"
  varHierarchyScopeCreateSpecItemType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varHierarchyScopeCreateSpecItemType
}

/*
 * EnterpriseAccountInRoleCreateSpecType 
 * Not validated 
 */
type EnterpriseAccountInRoleCreateSpecType struct { 
   XMLName xml.Name
   EnterpriseRoleUid UidType `xml:"EnterpriseRoleUid"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}
func NewEnterpriseAccountInRoleCreateSpecType() (*EnterpriseAccountInRoleCreateSpecType) {
  varEnterpriseAccountInRoleCreateSpecType := EnterpriseAccountInRoleCreateSpecType{}
  varEnterpriseAccountInRoleCreateSpecType.XMLName.Local = "EnterpriseAccountInRoleCreateSpec"
  varEnterpriseAccountInRoleCreateSpecType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varEnterpriseAccountInRoleCreateSpecType
}

/*
 * EnterpriseAccountInRoleCreateSpecListType 
 * Not validated 
 */
type EnterpriseAccountInRoleCreateSpecListType struct { 
   XMLName xml.Name
   EnterpriseRole []*EnterpriseAccountInRoleCreateSpecType `xml:"EnterpriseRole"`
   //Inhereting from ListType
}

/*
 * EnterpriseSecuritySettingsType 
 * Not validated 
 */
type EnterpriseSecuritySettingsType struct { 
   XMLName xml.Name
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l EnterpriseSecuritySettingsType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * EnterpriseAccountCreateSpecType 
 * Not validated 
 */
type EnterpriseAccountCreateSpecType struct { 
   XMLName xml.Name
   AccountType AccountTypeEnumeration `xml:"AccountType,omitempty"`
   AccountName string `xml:"AccountName,omitempty"`
   Roles *EnterpriseAccountInRoleCreateSpecListType `xml:"Roles,omitempty"`
   AllowRestoreAllVms bool `xml:"AllowRestoreAllVms,omitempty"`
   HierarchyScopeObjects *HierarchyScopeCreateSpecType `xml:"HierarchyScopeObjects,omitempty"`
   FlrSettings *FileRestoreSettingsSpecsType `xml:"FlrSettings,omitempty"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}
func NewEnterpriseAccountCreateSpecType() (*EnterpriseAccountCreateSpecType) {
  varEnterpriseAccountCreateSpecType := EnterpriseAccountCreateSpecType{}
  varEnterpriseAccountCreateSpecType.XMLName.Local = "EnterpriseAccountCreateSpec"
  varEnterpriseAccountCreateSpecType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varEnterpriseAccountCreateSpecType
}

/*
 * FileRestoreSettingsSpecsType 
 * Not validated 
 */
type FileRestoreSettingsSpecsType struct { 
   XMLName xml.Name
   FlrInlaceOnly bool `xml:"FlrInlaceOnly,omitempty"`
   FlrExtentionRestrictions string `xml:"FlrExtentionRestrictions"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}
func NewFileRestoreSettingsSpecsType() (*FileRestoreSettingsSpecsType) {
  varFileRestoreSettingsSpecsType := FileRestoreSettingsSpecsType{}
  varFileRestoreSettingsSpecsType.XMLName.Local = "FileRestoreSettingsSpecs"
  varFileRestoreSettingsSpecsType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varFileRestoreSettingsSpecsType
}

/*
 * RebuildScopeJobSpecType 
 * Not validated 
 */
type RebuildScopeJobSpecType struct { 
   XMLName xml.Name
   RebuildAll *RebuildScopeJobSpecTypeNestedRebuildAll `xml:"RebuildAll"`
   RebuildUnprocessed *RebuildScopeJobSpecTypeNestedRebuildUnprocessed `xml:"RebuildUnprocessed"`
   RebuildForCurrentUser *RebuildScopeJobSpecTypeNestedRebuildForCurrentUser `xml:"RebuildForCurrentUser"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}
func NewRebuildScopeJobSpecType() (*RebuildScopeJobSpecType) {
  varRebuildScopeJobSpecType := RebuildScopeJobSpecType{}
  varRebuildScopeJobSpecType.XMLName.Local = "RebuildScopeJobSpec"
  varRebuildScopeJobSpecType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varRebuildScopeJobSpecType
}

/*
 * RebuildScopeJobSpecTypeNestedRebuildAll 
 * Not validated 
 */
type RebuildScopeJobSpecTypeNestedRebuildAll struct { 
   XMLName xml.Name
}

/*
 * RebuildScopeJobSpecTypeNestedRebuildUnprocessed 
 * Not validated 
 */
type RebuildScopeJobSpecTypeNestedRebuildUnprocessed struct { 
   XMLName xml.Name
}

/*
 * RebuildScopeJobSpecTypeNestedRebuildForCurrentUser 
 * Not validated 
 */
type RebuildScopeJobSpecTypeNestedRebuildForCurrentUser struct { 
   XMLName xml.Name
}

/*
 * WanAcceleratorEntityType 
 * Not validated 
 */
type WanAcceleratorEntityType struct { 
   XMLName xml.Name
   Description string `xml:"Description"`
   OutOfDate bool `xml:"OutOfDate"`
   Version string `xml:"Version"`
   Capacity int64 `xml:"Capacity"`
   TrafficPort int `xml:"TrafficPort"`
   ConnectionsCount int `xml:"ConnectionsCount"`
   CachePath string `xml:"CachePath"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l WanAcceleratorEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * WanAcceleratorEntityListType 
 * Not validated 
 */
type WanAcceleratorEntityListType struct { 
   XMLName xml.Name
   WanAccelerator []*WanAcceleratorEntityType `xml:"WanAccelerator"`
   //Inhereting from ListType
}

/*
 * CloudGatewayEntityType 
 * Not validated 
 */
type CloudGatewayEntityType struct { 
   XMLName xml.Name
   Enabled bool `xml:"Enabled,omitempty"`
   DnsNameOrIpAddress string `xml:"DnsNameOrIpAddress"`
   NetworkMode CloudGatewayNetworkingMode `xml:"NetworkMode,omitempty"`
   ExternalIP string `xml:"ExternalIP"`
   ExternalPort uint16 `xml:"ExternalPort,omitempty"`
   InternalPort uint16 `xml:"InternalPort,omitempty"`
   Description string `xml:"Description"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CloudGatewayEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * CloudGatewayEntityListType 
 * Not validated 
 */
type CloudGatewayEntityListType struct { 
   XMLName xml.Name
   CloudGateway []*CloudGatewayEntityType `xml:"CloudGateway"`
   //Inhereting from ListType
}

/*
 * CloudTenantEntityListType 
 * Not validated 
 */
type CloudTenantEntityListType struct { 
   XMLName xml.Name
   CloudTenant []*CloudTenantEntityType `xml:"CloudTenant"`
   //Inhereting from ListType
}

/*
 * CloudTenantEntityType 
 * Not validated 
 */
type CloudTenantEntityType struct { 
   XMLName xml.Name
   Password string `xml:"Password"`
   Description string `xml:"Description"`
   Enabled bool `xml:"Enabled,omitempty"`
   LeaseOptions *CloudTenantLeaseOptionsType `xml:"LeaseOptions,omitempty"`
   Resources *CloudTenantResourceListType `xml:"Resources"`
   LastResult string `xml:"LastResult"`
   LastActive DateTime `xml:"LastActive,omitempty"`
   VmCount int `xml:"VmCount,omitempty"`
   ComputeResources *CloudTenantComputeResourceListType `xml:"ComputeResources,omitempty"`
   ThrottlingEnabled bool `xml:"ThrottlingEnabled,omitempty"`
   ThrottlingSpeedLimit int `xml:"ThrottlingSpeedLimit,omitempty"`
   ThrottlingSpeedUnit string `xml:"ThrottlingSpeedUnit,omitempty"`
   PublicIpCount int `xml:"PublicIpCount,omitempty"`
   BackupCount int `xml:"BackupCount,omitempty"`
   ReplicaCount int `xml:"ReplicaCount,omitempty"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CloudTenantEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * CloudTenantLeaseOptionsType 
 * Not validated 
 */
type CloudTenantLeaseOptionsType struct { 
   XMLName xml.Name
   Enabled bool `xml:"Enabled,attr"`
   ExpirationDate DateTime `xml:"ExpirationDate,attr"`
   //Inhereting from InfoType
}

/*
 * CloudTenantResourceType 
 * Not validated 
 */
type CloudTenantResourceType struct { 
   XMLName xml.Name
   RepositoryQuota *CloudTenantRepositoryQuotaInfoType `xml:"RepositoryQuota,omitempty"`
   Id string `xml:"Id,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CloudTenantResourceType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * CloudTenantResourceListType 
 * Not validated 
 */
type CloudTenantResourceListType struct { 
   XMLName xml.Name
   CloudTenantResource []*CloudTenantResourceType `xml:"CloudTenantResource"`
   //Inhereting from ListType
}

/*
 * CloudTenantRepositoryQuotaInfoType 
 * Not validated 
 */
type CloudTenantRepositoryQuotaInfoType struct { 
   XMLName xml.Name
   DisplayName string `xml:"DisplayName"`
   RepositoryUid UidType `xml:"RepositoryUid"`
   WanAcceleratorUid UidType `xml:"WanAcceleratorUid,omitempty"`
   Quota int64 `xml:"Quota,omitempty"`
   UsedQuota int64 `xml:"UsedQuota,omitempty"`
   //Inhereting from InfoType
}

/*
 * CreateCloudGatewaySpecType 
 * Not validated 
 */
type CreateCloudGatewaySpecType struct { 
   XMLName xml.Name
   BackupServerIdOrName string `xml:"BackupServerIdOrName"`
   ServerHostName string `xml:"ServerHostName"`
   Description string `xml:"Description"`
   IncomingPort uint16 `xml:"IncomingPort,omitempty"`
   ExternalIp string `xml:"ExternalIp"`
   ExternalPort uint16 `xml:"ExternalPort"`
   NetworkType CloudGatewayNetworkingMode `xml:"NetworkType,omitempty"`
   InternalPort uint16 `xml:"InternalPort,omitempty"`
   BackupServerUid UidType `xml:"BackupServerUid"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}
func NewCreateCloudGatewaySpecType() (*CreateCloudGatewaySpecType) {
  varCreateCloudGatewaySpecType := CreateCloudGatewaySpecType{}
  varCreateCloudGatewaySpecType.XMLName.Local = "CreateCloudGatewaySpec"
  varCreateCloudGatewaySpecType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varCreateCloudGatewaySpecType
}

/*
 * CreateCloudTenantSpecType 
 * Not validated 
 */
type CreateCloudTenantSpecType struct { 
   XMLName xml.Name
   BackupServerIdOrName string `xml:"BackupServerIdOrName"`
   Name string `xml:"Name"`
   Description string `xml:"Description"`
   Password string `xml:"Password"`
   Enabled bool `xml:"Enabled"`
   LeaseExpirationDate DateTime `xml:"LeaseExpirationDate,omitempty"`
   Resources *CreateCloudTenantResourceListType `xml:"Resources"`
   ComputeResources *CloudTenantComputeResourceCreateListType `xml:"ComputeResources"`
   VmCount int `xml:"VmCount,omitempty"`
   ThrottlingEnabled bool `xml:"ThrottlingEnabled,omitempty"`
   ThrottlingSpeedLimit int `xml:"ThrottlingSpeedLimit,omitempty"`
   ThrottlingSpeedUnit string `xml:"ThrottlingSpeedUnit,omitempty"`
   PublicIpCount int `xml:"PublicIpCount,omitempty"`
   BackupServerUid UidType `xml:"BackupServerUid"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}
func NewCreateCloudTenantSpecType() (*CreateCloudTenantSpecType) {
  varCreateCloudTenantSpecType := CreateCloudTenantSpecType{}
  varCreateCloudTenantSpecType.XMLName.Local = "CreateCloudTenantSpec"
  varCreateCloudTenantSpecType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varCreateCloudTenantSpecType
}

/*
 * CreateCloudTenantResourceListType 
 * Not validated 
 */
type CreateCloudTenantResourceListType struct { 
   XMLName xml.Name
   BackupResource []*CreateCloudTenantResourceSpecType `xml:"BackupResource"`
   //Inhereting from ListType
}

/*
 * CreateCloudTenantResourceSpecType 
 * Not validated 
 */
type CreateCloudTenantResourceSpecType struct { 
   XMLName xml.Name
   Name string `xml:"Name"`
   RepositoryUid UidType `xml:"RepositoryUid"`
   QuotaMb int `xml:"QuotaMb"`
   WanAcceleratorUid UidType `xml:"WanAcceleratorUid,omitempty"`
   Folder string `xml:"Folder"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}
func NewCreateCloudTenantResourceSpecType() (*CreateCloudTenantResourceSpecType) {
  varCreateCloudTenantResourceSpecType := CreateCloudTenantResourceSpecType{}
  varCreateCloudTenantResourceSpecType.XMLName.Local = "CreateCloudTenantResourceSpec"
  varCreateCloudTenantResourceSpecType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varCreateCloudTenantResourceSpecType
}

/*
 * CloudHardwarePlanEntityListType 
 * Not validated 
 */
type CloudHardwarePlanEntityListType struct { 
   XMLName xml.Name
   CloudHardwarePlan []*CloudHardwarePlanEntityType `xml:"CloudHardwarePlan"`
   //Inhereting from ListType
}

/*
 * CloudHardwarePlanEntityType 
 * Not validated 
 */
type CloudHardwarePlanEntityType struct { 
   XMLName xml.Name
   Description string `xml:"Description,omitempty"`
   ProcessorUsageLimitMhz int `xml:"ProcessorUsageLimitMhz,omitempty"`
   MemoryUsageLimitMb int `xml:"MemoryUsageLimitMb,omitempty"`
   HardwarePlanDetails *CloudHardwarePlanEntityTypeNestedHardwarePlanDetails `xml:"HardwarePlanDetails,omitempty"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CloudHardwarePlanEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * CloudHardwarePlanEntityTypeNestedHardwarePlanDetails 
 * Not validated 
 */
type CloudHardwarePlanEntityTypeNestedHardwarePlanDetails struct { 
   XMLName xml.Name
   ViCloudHardwarePlan *ViCloudHardwarePlanInfoType `xml:"ViCloudHardwarePlan,omitempty"`
   HvCloudHardwarePlan *HvCloudHardwarePlanInfoType `xml:"HvCloudHardwarePlan,omitempty"`
}

/*
 * CloudHardwarePlanCreateSpecType 
 * Not validated 
 */
type CloudHardwarePlanCreateSpecType struct { 
   XMLName xml.Name
   BackupServerUid UidType `xml:"BackupServerUid"`
   Name string `xml:"Name"`
   Description string `xml:"Description"`
   ProcessorUsageLimitMhz int `xml:"ProcessorUsageLimitMhz"`
   MemoryUsageLimitMb int `xml:"MemoryUsageLimitMb"`
   HardwarePlanDetails *CloudHardwarePlanCreateSpecTypeNestedHardwarePlanDetails `xml:"HardwarePlanDetails,omitempty"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}
func NewCloudHardwarePlanCreateSpecType() (*CloudHardwarePlanCreateSpecType) {
  varCloudHardwarePlanCreateSpecType := CloudHardwarePlanCreateSpecType{}
  varCloudHardwarePlanCreateSpecType.XMLName.Local = "CloudHardwarePlanCreateSpec"
  varCloudHardwarePlanCreateSpecType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varCloudHardwarePlanCreateSpecType
}

/*
 * CloudHardwarePlanCreateSpecTypeNestedHardwarePlanDetails 
 * Not validated 
 */
type CloudHardwarePlanCreateSpecTypeNestedHardwarePlanDetails struct { 
   XMLName xml.Name
   ViCloudHardwarePlan *ViCloudHardwarePlanInfoType `xml:"ViCloudHardwarePlan,omitempty"`
   HvCloudHardwarePlan *HvCloudHardwarePlanInfoType `xml:"HvCloudHardwarePlan,omitempty"`
}

/*
 * ViCloudHardwarePlanInfoType 
 * Not validated 
 */
type ViCloudHardwarePlanInfoType struct { 
   XMLName xml.Name
   HypervisorHostRef UidType `xml:"HypervisorHostRef"`
   ParentType string `xml:"ParentType"`
   ParentName string `xml:"ParentName"`
   Datastores *ViCloudHardwarePlanDatastoreInfoListType `xml:"Datastores"`
   Network *CloudHardwarePlanNetworkInfo `xml:"Network"`
   //Inhereting from InfoType
}

/*
 * ViCloudHardwarePlanDatastoreInfoListType 
 * Not validated 
 */
type ViCloudHardwarePlanDatastoreInfoListType struct { 
   XMLName xml.Name
   Datastore []*ViCloudHardwarePlanDatastoreInfoType `xml:"Datastore"`
   //Inhereting from ListType
}

/*
 * ViCloudHardwarePlanDatastoreInfoType 
 * Not validated 
 */
type ViCloudHardwarePlanDatastoreInfoType struct { 
   XMLName xml.Name
   FriendlyName string `xml:"FriendlyName"`
   DatastoreType string `xml:"DatastoreType"`
   Reference HierarchyObjRefType `xml:"Reference"`
   RootPath string `xml:"RootPath"`
   QuotaGb int `xml:"QuotaGb"`
   PbmProfileId string `xml:"PbmProfileId"`
   Id string `xml:"Id,attr"`
   //Inhereting from InfoType
}

/*
 * HvCloudHardwarePlanInfoType 
 * Not validated 
 */
type HvCloudHardwarePlanInfoType struct { 
   XMLName xml.Name
   HypervisorHostRef HierarchyObjRefType `xml:"HypervisorHostRef"`
   Volumes *HvCloudHardwarePlanVolumesInfoListType `xml:"Volumes"`
   Network *CloudHardwarePlanNetworkInfo `xml:"Network"`
   //Inhereting from InfoType
}

/*
 * HvCloudHardwarePlanVolumesInfoListType 
 * Not validated 
 */
type HvCloudHardwarePlanVolumesInfoListType struct { 
   XMLName xml.Name
   Volume []*HvCloudHardwarePlanVolumeInfoType `xml:"Volume"`
   //Inhereting from ListType
}

/*
 * HvCloudHardwarePlanVolumeInfoType 
 * Not validated 
 */
type HvCloudHardwarePlanVolumeInfoType struct { 
   XMLName xml.Name
   FriendlyName string `xml:"FriendlyName"`
   VolumePath string `xml:"VolumePath"`
   QuotaGb int `xml:"QuotaGb"`
   Id string `xml:"Id,attr"`
   //Inhereting from InfoType
}

/*
 * CloudHardwarePlanNetworkInfo 
 * Not validated 
 */
type CloudHardwarePlanNetworkInfo struct { 
   XMLName xml.Name
   CountWithInternet int `xml:"CountWithInternet"`
   CountWithoutInternet int `xml:"CountWithoutInternet"`
   Id string `xml:"Id,attr"`
   //Inhereting from InfoType
}

/*
 * CloudFailoverSessionEntityType 
 * Not validated 
 */
type CloudFailoverSessionEntityType struct { 
   XMLName xml.Name
   CloudFailoverTasks *CloudFailoverTaskSessionInfoListType `xml:"CloudFailoverTasks"`
   CloudFailoverPlanName string `xml:"CloudFailoverPlanName,attr"`
   //Inhereting from JobSessionEntityType
   JobUid UidType `xml:"JobUid"`
   JobName string `xml:"JobName"`
   JobType string `xml:"JobType"`
   CreationTimeUTC DateTime `xml:"CreationTimeUTC"`
   EndTimeUTC DateTime `xml:"EndTimeUTC,omitempty"`
   State string `xml:"State"`
   Result string `xml:"Result"`
   Progress int `xml:"Progress"`
   FailureMessage string `xml:"FailureMessage"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CloudFailoverSessionEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * CloudFailoverSessionEntityListType 
 * Not validated 
 */
type CloudFailoverSessionEntityListType struct { 
   XMLName xml.Name
   CloudFailoverSession []*CloudFailoverSessionEntityType `xml:"CloudFailoverSession"`
   //Inhereting from ListType
}

/*
 * CloudFailoverTaskSessionInfoType 
 * Not validated 
 */
type CloudFailoverTaskSessionInfoType struct { 
   XMLName xml.Name
   VmReplicaPointLink *LinkType `xml:"VmReplicaPointLink"`
   CreationTimeUTC DateTime `xml:"CreationTimeUTC"`
   EndTimeUTC DateTime `xml:"EndTimeUTC,omitempty"`
   State string `xml:"State"`
   Result string `xml:"Result"`
   Progress int `xml:"Progress"`
   FailureMessage string `xml:"FailureMessage"`
   VmName string `xml:"VmName,attr"`
   //Inhereting from InfoType
}

/*
 * CloudFailoverTaskSessionInfoListType 
 * Not validated 
 */
type CloudFailoverTaskSessionInfoListType struct { 
   XMLName xml.Name
   CloudFailoverTasks []*CloudFailoverTaskSessionInfoType `xml:"CloudFailoverTasks"`
   //Inhereting from ListType
}

/*
 * CloudPublicIpAddressEntityListType 
 * Not validated 
 */
type CloudPublicIpAddressEntityListType struct { 
   XMLName xml.Name
   CloudPublicIp []*CloudPublicIpAddressEntityType `xml:"CloudPublicIp"`
   //Inhereting from ListType
}

/*
 * CloudPublicIpAddressEntityType 
 * Not validated 
 */
type CloudPublicIpAddressEntityType struct { 
   XMLName xml.Name
   IpAddress string `xml:"IpAddress"`
   TenantUid UidType `xml:"TenantUid"`
   BackupServerUid UidType `xml:"BackupServerUid,attr"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CloudPublicIpAddressEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * CloudPublicIpAddressCreateSpecType 
 * Not validated 
 */
type CloudPublicIpAddressCreateSpecType struct { 
   XMLName xml.Name
   BackupServerUid UidType `xml:"BackupServerUid"`
   IpAddressLowerBound string `xml:"IpAddressLowerBound"`
   IpAddressUpperBound string `xml:"IpAddressUpperBound"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}
func NewCloudPublicIpAddressCreateSpecType() (*CloudPublicIpAddressCreateSpecType) {
  varCloudPublicIpAddressCreateSpecType := CloudPublicIpAddressCreateSpecType{}
  varCloudPublicIpAddressCreateSpecType.XMLName.Local = "CloudPublicIpAddressCreateSpec"
  varCloudPublicIpAddressCreateSpecType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varCloudPublicIpAddressCreateSpecType
}

/*
 * CloudTenantComputeResourceListType 
 * Not validated 
 */
type CloudTenantComputeResourceListType struct { 
   XMLName xml.Name
   CloudTenantComputeResource []*CloudTenantComputeResourceType `xml:"CloudTenantComputeResource"`
   //Inhereting from ListType
}

/*
 * CloudTenantComputeResourceType 
 * Not validated 
 */
type CloudTenantComputeResourceType struct { 
   XMLName xml.Name
   CloudHardwarePlanUid UidType `xml:"CloudHardwarePlanUid"`
   WanAcceleratorUid UidType `xml:"WanAcceleratorUid,omitempty"`
   PlatformType string `xml:"PlatformType"`
   UseNetworkFailoverResources bool `xml:"UseNetworkFailoverResources"`
   NetworkAppliance *NetworkApplianceInfoType `xml:"NetworkAppliance"`
   ComputeResourceStats *ComputeResourceStatsInfoType `xml:"ComputeResourceStats"`
   Id string `xml:"Id,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CloudTenantComputeResourceType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * ComputeResourceStatsInfoType 
 * Not validated 
 */
type ComputeResourceStatsInfoType struct { 
   XMLName xml.Name
   MemoryUsageMb int `xml:"MemoryUsageMb"`
   CPUCount int `xml:"CPUCount"`
   StorageResourceStats *StorageResourceStatsListType `xml:"StorageResourceStats"`
   //Inhereting from InfoType
}

/*
 * StorageResourceStatsListType 
 * Not validated 
 */
type StorageResourceStatsListType struct { 
   XMLName xml.Name
   StorageResourceStat []*StorageResourceStatInfoType `xml:"StorageResourceStat"`
   //Inhereting from ListType
}

/*
 * StorageResourceStatInfoType 
 * Not validated 
 */
type StorageResourceStatInfoType struct { 
   XMLName xml.Name
   StorageName string `xml:"StorageName"`
   StorageUsageGb int `xml:"StorageUsageGb"`
   StorageLimitGb int `xml:"StorageLimitGb"`
   //Inhereting from InfoType
}

/*
 * CloudTenantComputeResourceLeaseOptionsType 
 * Not validated 
 */
type CloudTenantComputeResourceLeaseOptionsType struct { 
   XMLName xml.Name
   Enabled bool `xml:"Enabled,attr"`
   ExpirationDate DateTime `xml:"ExpirationDate,attr"`
   //Inhereting from InfoType
}

/*
 * CloudTenantComputeResourceCreateListType 
 * Not validated 
 */
type CloudTenantComputeResourceCreateListType struct { 
   XMLName xml.Name
   ComputeResource []*CloudTenantComputeResourceCreateSpecType `xml:"ComputeResource"`
   //Inhereting from ListType
}

/*
 * CloudTenantComputeResourceCreateSpecType 
 * Not validated 
 */
type CloudTenantComputeResourceCreateSpecType struct { 
   XMLName xml.Name
   CloudHardwarePlanUid UidType `xml:"CloudHardwarePlanUid"`
   WanAcceleratorUid UidType `xml:"WanAcceleratorUid,omitempty"`
   PlatformType string `xml:"PlatformType"`
   UseNetworkFailoverResources bool `xml:"UseNetworkFailoverResources"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}
func NewCloudTenantComputeResourceCreateSpecType() (*CloudTenantComputeResourceCreateSpecType) {
  varCloudTenantComputeResourceCreateSpecType := CloudTenantComputeResourceCreateSpecType{}
  varCloudTenantComputeResourceCreateSpecType.XMLName.Local = "CloudTenantComputeResourceCreateSpec"
  varCloudTenantComputeResourceCreateSpecType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varCloudTenantComputeResourceCreateSpecType
}

/*
 * CloudVmReplicaPointEntityType 
 * Not validated 
 */
type CloudVmReplicaPointEntityType struct { 
   XMLName xml.Name
   CreationTimeUTC DateTime `xml:"CreationTimeUTC"`
   VmName string `xml:"VmName"`
   PointType string `xml:"PointType"`
   State string `xml:"State"`
   VmDisplayName string `xml:"VmDisplayName,attr"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CloudVmReplicaPointEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * CloudVmReplicaPointEntityListType 
 * Not validated 
 */
type CloudVmReplicaPointEntityListType struct { 
   XMLName xml.Name
   CloudReplica []*CloudVmReplicaPointEntityType `xml:"CloudReplica"`
   //Inhereting from ListType
}

/*
 * CloudFailoverPlanEntityListType 
 * Not validated 
 */
type CloudFailoverPlanEntityListType struct { 
   XMLName xml.Name
   CloudFailoverPlan []*CloudFailoverPlanEntityType `xml:"CloudFailoverPlan"`
   //Inhereting from ListType
}

/*
 * CloudFailoverPlanEntityType 
 * Not validated 
 */
type CloudFailoverPlanEntityType struct { 
   XMLName xml.Name
   TenantUid UidType `xml:"TenantUid"`
   TenantName string `xml:"TenantName"`
   Description string `xml:"Description,omitempty"`
   CloudFailoverPlanOptions *CloudFailoverPlanOptionsInfoType `xml:"CloudFailoverPlanOptions"`
   CloudFailoverPlanInfo *CloudFailoverPlanInfoType `xml:"CloudFailoverPlanInfo"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CloudFailoverPlanEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * CloudFailoverPlanManagementSpecType 
 * Not validated 
 */
type CloudFailoverPlanManagementSpecType struct { 
   XMLName xml.Name
   StartNow bool `xml:"StartNow"`
   StartDate DateTime `xml:"StartDate"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}
func NewCloudFailoverPlanManagementSpecType() (*CloudFailoverPlanManagementSpecType) {
  varCloudFailoverPlanManagementSpecType := CloudFailoverPlanManagementSpecType{}
  varCloudFailoverPlanManagementSpecType.XMLName.Local = "CloudFailoverPlanManagementSpec"
  varCloudFailoverPlanManagementSpecType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varCloudFailoverPlanManagementSpecType
}

/*
 * CloudFailoverPlanInfoType 
 * Not validated 
 */
type CloudFailoverPlanInfoType struct { 
   XMLName xml.Name
   Includes *CloudFailoveredVmListType `xml:"Includes,omitempty"`
   //Inhereting from InfoType
}

/*
 * CloudFailoveredVmType 
 * Not validated 
 */
type CloudFailoveredVmType struct { 
   XMLName xml.Name
   FailoverPlanVMId string `xml:"FailoverPlanVMId,omitempty"`
   Name string `xml:"Name,omitempty"`
   Order int `xml:"Order,omitempty"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CloudFailoveredVmType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * CloudFailoveredVmListType 
 * Not validated 
 */
type CloudFailoveredVmListType struct { 
   XMLName xml.Name
   CloudFailoveredVm []*CloudFailoveredVmType `xml:"CloudFailoveredVm"`
   //Inhereting from ListType
}

/*
 * CloudFailoverPlanOptionsInfoType 
 * Not validated 
 */
type CloudFailoverPlanOptionsInfoType struct { 
   XMLName xml.Name
   PostFailoverPlanCommandEnabled bool `xml:"PostFailoverPlanCommandEnabled"`
   PostFailoverPlanCommand string `xml:"PostFailoverPlanCommand"`
   PreFailoverPlanCommandEnabled bool `xml:"PreFailoverPlanCommandEnabled"`
   PreFailoverPlanCommand string `xml:"PreFailoverPlanCommand"`
   //Inhereting from InfoType
}

/*
 * CloudReplicaEntityType 
 * Not validated 
 */
type CloudReplicaEntityType struct { 
   XMLName xml.Name
   Ref *EntityReferenceType `xml:"Ref"`
   Platform string `xml:"Platform,omitempty"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CloudReplicaEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * CloudReplicaEntityListType 
 * Not validated 
 */
type CloudReplicaEntityListType struct { 
   XMLName xml.Name
   CloudReplica []*CloudReplicaEntityType `xml:"CloudReplica"`
   //Inhereting from ListType
}

/*
 * CloudConnectServiceType 
 * Not validated 
 */
type CloudConnectServiceType struct { 
   XMLName xml.Name
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CloudConnectServiceType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * TenantCredentialsInfoType 
 * Not validated 
 */
type TenantCredentialsInfoType struct { 
   XMLName xml.Name
   Username string `xml:"Username"`
   Password string `xml:"Password"`
   //Inhereting from InfoType
}

/*
 * VlanConfigurationEntityListType 
 * Not validated 
 */
type VlanConfigurationEntityListType struct { 
   XMLName xml.Name
   Vlans []*VlanConfigurationEntityType `xml:"Vlans"`
   //Inhereting from ListType
}

/*
 * VlanConfigurationEntityType 
 * Not validated 
 */
type VlanConfigurationEntityType struct { 
   XMLName xml.Name
   HostRef HierarchyObjRefType `xml:"HostRef"`
   PlatformType string `xml:"PlatformType"`
   VlanIdsWithInternetLeftBound int `xml:"VlanIdsWithInternetLeftBound"`
   VlanIdsWithInternetRightBound int `xml:"VlanIdsWithInternetRightBound"`
   VlanIdsWithoutInternetLeftBound int `xml:"VlanIdsWithoutInternetLeftBound"`
   VlanIdsWithoutInternetRightBound int `xml:"VlanIdsWithoutInternetRightBound"`
   SwitchName string `xml:"SwitchName"`
   SwitchId string `xml:"SwitchId"`
   ClusterName string `xml:"ClusterName,omitempty"`
   ClusterId string `xml:"ClusterId"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links *LinkListType `xml:"Links,omitempty"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l VlanConfigurationEntityType) GetLinks() (*LinkListType) { return (l.Links) }

/*
 * CloudVlanConfigurationCreateSpecType 
 * Not validated 
 */
type CloudVlanConfigurationCreateSpecType struct { 
   XMLName xml.Name
   BackupServerUid UidType `xml:"BackupServerUid"`
   HostRef HierarchyObjRefType `xml:"HostRef"`
   PlatformType string `xml:"PlatformType"`
   VlanIdsWithInternetLeftBound int `xml:"VlanIdsWithInternetLeftBound"`
   VlanIdsWithInternetRightBound int `xml:"VlanIdsWithInternetRightBound"`
   VlanIdsWithoutInternetLeftBound int `xml:"VlanIdsWithoutInternetLeftBound"`
   VlanIdsWithoutInternetRightBound int `xml:"VlanIdsWithoutInternetRightBound"`
   viClusterRef string `xml:"viClusterRef"`
   viClusterName string `xml:"viClusterName"`
   viSwitchName string `xml:"viSwitchName"`
   viSwitchType int `xml:"viSwitchType"`
   viSwitchId string `xml:"viSwitchId"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}
func NewCloudVlanConfigurationCreateSpecType() (*CloudVlanConfigurationCreateSpecType) {
  varCloudVlanConfigurationCreateSpecType := CloudVlanConfigurationCreateSpecType{}
  varCloudVlanConfigurationCreateSpecType.XMLName.Local = "CloudVlanConfigurationCreateSpec"
  varCloudVlanConfigurationCreateSpecType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varCloudVlanConfigurationCreateSpecType
}

/*
 * NetworkApplianceInfoType 
 * Not validated 
 */
type NetworkApplianceInfoType struct { 
   XMLName xml.Name
   Name string `xml:"Name"`
   ProductionNetwork string `xml:"ProductionNetwork"`
   ObtainIPAddressAutomatically bool `xml:"ObtainIPAddressAutomatically"`
   ManualIpAddressSettingsInfoType *ManualIpAddressSettingsInfoType `xml:"ManualIpAddressSettingsInfoType,omitempty"`
   //Inhereting from InfoType
}

/*
 * ManualIpAddressSettingsInfoType 
 * Not validated 
 */
type ManualIpAddressSettingsInfoType struct { 
   XMLName xml.Name
   IpAddress IPv4 `xml:"IpAddress"`
   SubnetMask IPv4 `xml:"SubnetMask"`
   DefaultGateway IPv4 `xml:"DefaultGateway"`
   //Inhereting from InfoType
}

/*
 * VmRestorePointSqlInfoType 
 * Not validated 
 */
type VmRestorePointSqlInfoType struct { 
   XMLName xml.Name
   Databases *VmRestorePointSqlDatabaseInfoListType `xml:"Databases"`
   //Inhereting from InfoType
}

/*
 * VmRestorePointSqlDatabaseInfoListType 
 * Not validated 
 */
type VmRestorePointSqlDatabaseInfoListType struct { 
   XMLName xml.Name
   Database []*VmRestorePointSqlDatabaseInfoType `xml:"Database"`
   //Inhereting from ListType
}

/*
 * VmRestorePointSqlDatabaseInfoType 
 * Not validated 
 */
type VmRestorePointSqlDatabaseInfoType struct { 
   XMLName xml.Name
   Items *VmRestorePointSqlDatabaseItemListInfoType `xml:"Items"`
   Instance string `xml:"Instance,attr"`
   InstanceName string `xml:"InstanceName,attr"`
   Name string `xml:"Name,attr"`
   FullVersion string `xml:"FullVersion,attr"`
   ProductVersion string `xml:"ProductVersion,attr"`
   CheckpointLsn uint64 `xml:"CheckpointLsn,attr"`
   DatabaseBackupLsn uint64 `xml:"DatabaseBackupLsn,attr"`
   FirstLsn uint64 `xml:"FirstLsn,attr"`
   LastLsn uint64 `xml:"LastLsn,attr"`
   StartCreationTime DateTime `xml:"StartCreationTime,attr"`
   EndCreationTime DateTime `xml:"EndCreationTime,attr"`
   TimeZone int `xml:"TimeZone,attr"`
   RecoveryModel string `xml:"RecoveryModel,attr"`
   CompatibilityLevel int `xml:"CompatibilityLevel,attr"`
   CertName string `xml:"CertName,attr"`
   CertSerialNumber string `xml:"CertSerialNumber,attr"`
   ReadOnly string `xml:"ReadOnly,attr"`
   BindingID string `xml:"BindingID,attr"`
   FamilyGUID string `xml:"FamilyGUID,attr"`
   DatabaseId int `xml:"DatabaseId,attr"`
   //Inhereting from InfoType
}

/*
 * VmRestorePointSqlDatabaseItemListInfoType 
 * Not validated 
 */
type VmRestorePointSqlDatabaseItemListInfoType struct { 
   XMLName xml.Name
   Item []*VmRestorePointSqlDatabaseItemInfoType `xml:"Item"`
   //Inhereting from ListType
}

/*
 * VmRestorePointSqlDatabaseItemInfoType 
 * Not validated 
 */
type VmRestorePointSqlDatabaseItemInfoType struct { 
   XMLName xml.Name
   Type string `xml:"Type,attr"`
   DataPath string `xml:"DataPath,attr"`
   //Inhereting from InfoType
}

/*
 * SqlItemRestoreSpecInfoType 
 * Not validated 
 */
type SqlItemRestoreSpecInfoType struct { 
   XMLName xml.Name
   IsRestoreToOriginal bool `xml:"IsRestoreToOriginal"`
   ServerName string `xml:"ServerName"`
   VmName string `xml:"VmName"`
   InstanceName string `xml:"InstanceName"`
   FamilyUid string `xml:"FamilyUid"`
   Credentials *SqlCredentialsInfoType `xml:"Credentials"`
   DatabaseName string `xml:"DatabaseName"`
   SourceDatabaseName string `xml:"SourceDatabaseName"`
   SourceInstanceName string `xml:"SourceInstanceName"`
   RestoreToDate DateTime `xml:"RestoreToDate"`
   //Inhereting from InfoType
}

/*
 * SqlCredentialsInfoType 
 * Not validated 
 */
type SqlCredentialsInfoType struct { 
   XMLName xml.Name
   SqlCredentials *PlainCredentialsType `xml:"SqlCredentials,omitempty"`
   ServerCredentials *PlainCredentialsType `xml:"ServerCredentials,omitempty"`
   UserCredentials *PlainCredentialsType `xml:"UserCredentials,omitempty"`
   UseSqlAuth bool `xml:"UseSqlAuth"`
   //Inhereting from InfoType
}

/*
 * VeeamZipStartupSpecType 
 * Not validated 
 */
type VeeamZipStartupSpecType struct { 
   XMLName xml.Name
   VmRef HierarchyObjRefType `xml:"VmRef"`
   RepositoryUid UidType `xml:"RepositoryUid"`
   PasswordKeyId string `xml:"PasswordKeyId"`
   CompressionLevel int `xml:"CompressionLevel"`
   DisableGuestQuiescence bool `xml:"DisableGuestQuiescence"`
   BackupRetention FreeBackupRetentionEnumType `xml:"BackupRetention"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}
func NewVeeamZipStartupSpecType() (*VeeamZipStartupSpecType) {
  varVeeamZipStartupSpecType := VeeamZipStartupSpecType{}
  varVeeamZipStartupSpecType.XMLName.Local = "VeeamZipStartupSpec"
  varVeeamZipStartupSpecType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varVeeamZipStartupSpecType
}

/*
 * QuickBackupStartupSpecType 
 * Not validated 
 */
type QuickBackupStartupSpecType struct { 
   XMLName xml.Name
   VmRef HierarchyObjRefType `xml:"VmRef"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}
func NewQuickBackupStartupSpecType() (*QuickBackupStartupSpecType) {
  varQuickBackupStartupSpecType := QuickBackupStartupSpecType{}
  varQuickBackupStartupSpecType.XMLName.Local = "QuickBackupStartupSpec"
  varQuickBackupStartupSpecType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varQuickBackupStartupSpecType
}

