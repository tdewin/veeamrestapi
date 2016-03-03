package veeamrestapi

/*
 * LoginSpecType 
 * Not validated 
 */
type LoginSpecType struct { 
   VMwareSSOToken string `xml:"VMwareSSOToken"`
   TenantCredentials TenantCredentialsInfoType `xml:"TenantCredentials"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}

/*
 * LogonSessionListType 
 * Not validated 
 */
type LogonSessionListType struct { 
   LogonSession []LogonSessionType `xml:"LogonSession"`
   //Inhereting from ListType
}

/*
 * LogonSessionType 
 * Not validated 
 */
type LogonSessionType struct { 
   UserName string `xml:"UserName"`
   SessionId string `xml:"SessionId"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l LogonSessionType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * SummaryReportType 
 * Not validated 
 */
type SummaryReportType struct { 
   //Inhereting from ReportResourceType
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l SummaryReportType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * OverviewReportFrameType 
 * Not validated 
 */
type OverviewReportFrameType struct { 
   BackupServers int `xml:"BackupServers"`
   ProxyServers int `xml:"ProxyServers"`
   RepositoryServers int `xml:"RepositoryServers"`
   RunningJobs int `xml:"RunningJobs"`
   ScheduledJobs int `xml:"ScheduledJobs"`
   SuccessfulVmLastestStates int `xml:"SuccessfulVmLastestStates"`
   WarningVmLastestStates int `xml:"WarningVmLastestStates"`
   FailedVmLastestStates int `xml:"FailedVmLastestStates"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l OverviewReportFrameType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * VmsOverviewReportFrameType 
 * Not validated 
 */
type VmsOverviewReportFrameType struct { 
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
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l VmsOverviewReportFrameType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * JobStatisticsReportFrameType 
 * Not validated 
 */
type JobStatisticsReportFrameType struct { 
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
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l JobStatisticsReportFrameType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * InfoType 
 * Not validated 
 */
type InfoType struct { 
}

/*
 * RepositoryReportFrameType 
 * Not validated 
 */
type RepositoryReportFrameType struct { 
   Period RepositoryReportFrameTypeNestedPeriod `xml:"Period"`
   CapacityPlanningReportLink string `xml:"CapacityPlanningReportLink,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l RepositoryReportFrameType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * RepositoryReportFrameTypeNestedPeriod 
 * Not validated 
 */
type RepositoryReportFrameTypeNestedPeriod struct { 
   Name []string `xml:"Name"`
   Capacity []int64 `xml:"Capacity"`
   FreeSpace []int64 `xml:"FreeSpace"`
   BackupSize []int64 `xml:"BackupSize"`
}

/*
 * ProcessedVmsReportFrameType 
 * Not validated 
 */
type ProcessedVmsReportFrameType struct { 
   Day ProcessedVmsReportFrameTypeNestedDay `xml:"Day"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l ProcessedVmsReportFrameType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * ProcessedVmsReportFrameTypeNestedDay 
 * Not validated 
 */
type ProcessedVmsReportFrameTypeNestedDay struct { 
   Timestamp DateTime `xml:"Timestamp,attr"`
   ReplicatedVms int `xml:"ReplicatedVms,attr"`
   BackupedVms int `xml:"BackupedVms,attr"`
}

/*
 * EnterpriseManagerType 
 * Not validated 
 */
type EnterpriseManagerType struct { 
   SupportedVersions SupportedVersionListType `xml:"SupportedVersions"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l EnterpriseManagerType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * SupportedVersionListType 
 * Not validated 
 */
type SupportedVersionListType struct { 
   SupportedVersion []SupportedVersionType `xml:"SupportedVersion"`
   //Inhereting from ListType
}

/*
 * SupportedVersionType 
 * Not validated 
 */
type SupportedVersionType struct { 
   Links LinkListType `xml:"Links"`
   Name string `xml:"Name,attr"`
   //Inhereting from InfoType
}
func (l SupportedVersionType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * ListType 
 * Not validated 
 */
type ListType struct { 
}

/*
 * ReportResourceType 
 * Not validated 
 */
type ReportResourceType struct { 
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l ReportResourceType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * LinkListType 
 * Not validated 
 */
type LinkListType struct { 
   Link []LinkType `xml:"Link"`
}

/*
 * LinkType 
 * Not validated 
 */
type LinkType struct { 
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
   Href UrlType `xml:"Href,attr"`
   Name string `xml:"Name,attr"`
   Type string `xml:"Type,attr"`
}

/*
 * ResourceType 
 * Not validated 
 */
type ResourceType struct { 
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l ResourceType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * EntityType 
 * Not validated 
 */
type EntityType struct { 
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l EntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * EntitiesType 
 * Not validated 
 */
type EntitiesType struct { 
   Jobs JobEntityListType `xml:"Jobs"`
   FailoverPlans FailoverPlanEntityListType `xml:"FailoverPlans"`
   Backups BackupEntityListType `xml:"Backups"`
   Replicas ReplicaEntityListType `xml:"Replicas"`
   Repositories RepositoryEntityListType `xml:"Repositories"`
   RestorePoints RestorePointEntityListType `xml:"RestorePoints"`
   VmRestorePoints VmRestorePointEntityListType `xml:"VmRestorePoints"`
   VAppRestorePoints VAppRestorePointEntityListType `xml:"VAppRestorePoints"`
   VmReplicaPoints VmReplicaPointEntityListType `xml:"VmReplicaPoints"`
   BackupJobSessions BackupJobSessionEntityListType `xml:"BackupJobSessions"`
   ReplicaJobSessions ReplicaJobSessionEntityListType `xml:"ReplicaJobSessions"`
   ReplicaTaskSessions ReplicaTaskSessionEntityListType `xml:"ReplicaTaskSessions"`
   RestoreSessions RestoreSessionEntityListType `xml:"RestoreSessions"`
   HierarchyRoots HierarchyRootEntityListType `xml:"HierarchyRoots"`
   BackupTaskSessions BackupTaskSessionEntityListType `xml:"BackupTaskSessions"`
   BackupServers BackupServerEntityListType `xml:"BackupServers"`
   ManagedServers ManagedServerEntityListType `xml:"ManagedServers"`
   EnterpiseRoles EnterpriseRoleEntityListType `xml:"EnterpiseRoles"`
   EnterpiseAccounts EnterpriseAccountEntityListType `xml:"EnterpiseAccounts"`
   WanAccelerators WanAcceleratorEntityListType `xml:"WanAccelerators"`
   CloudGateways CloudGatewayEntityListType `xml:"CloudGateways"`
   CloudTenants CloudTenantEntityListType `xml:"CloudTenants"`
   Passwords PasswordKeyInfoListType `xml:"Passwords"`
   Credentials CredentialsInfoListType `xml:"Credentials"`
}

/*
 * ResourcesType 
 * Not validated 
 */
type ResourcesType struct { 
   Files FileEntryListType `xml:"Files"`
   Directories DirectoryEntryListType `xml:"Directories"`
   Tasks TaskListType `xml:"Tasks"`
   CredentialsList CredentialsInfoListType `xml:"CredentialsList"`
   sInObjectJob ObjectInJobListType `xml:"sInObjectJob"`
   AccountRole EnterpriseAccountInRoleListType `xml:"AccountRole"`
   AccountHierarchyScope EnterpriseAccountHierarchyScopeListType `xml:"AccountHierarchyScope"`
   PasswordKeyInfoList PasswordKeyInfoListType `xml:"PasswordKeyInfoList"`
}

/*
 * ParamsType 
 * Not validated 
 */
type ParamsType struct { 
}

/*
 * SpecType 
 * Not validated 
 */
type SpecType struct { 
   //Inhereting from ParamsType
}

/*
 * EntityReferenceListType 
 * Not validated 
 */
type EntityReferenceListType struct { 
   Ref []EntityReferenceType `xml:"Ref"`
   //Inhereting from ListType
}

/*
 * EntityReferenceType 
 * Not validated 
 */
type EntityReferenceType struct { 
   Links LinkListType `xml:"Links"`
   UID UidType `xml:"UID,attr"`
   Name string `xml:"Name,attr"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l EntityReferenceType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * BackupServerEntityListType 
 * Not validated 
 */
type BackupServerEntityListType struct { 
   BackupServer []BackupServerEntityType `xml:"BackupServer"`
   //Inhereting from ListType
}

/*
 * ManagedServerEntityListType 
 * Not validated 
 */
type ManagedServerEntityListType struct { 
   ManagedServer []ManagedServerEntityType `xml:"ManagedServer"`
   //Inhereting from ListType
}

/*
 * BackupServerEntityType 
 * Not validated 
 */
type BackupServerEntityType struct { 
   Description string `xml:"Description"`
   Port uint16 `xml:"Port"`
   Version VersionType `xml:"Version"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l BackupServerEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * ManagedServerEntityType 
 * Not validated 
 */
type ManagedServerEntityType struct { 
   Description string `xml:"Description"`
   ManagedServerType string `xml:"ManagedServerType"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l ManagedServerEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * BackupServerSpecType 
 * Not validated 
 */
type BackupServerSpecType struct { 
   Description string `xml:"Description"`
   DnsNameOrIpAddress string `xml:"DnsNameOrIpAddress"`
   Port uint16 `xml:"Port"`
   Username string `xml:"Username"`
   Password string `xml:"Password"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}

/*
 * JobSessionEntityType 
 * Not validated 
 */
type JobSessionEntityType struct { 
   JobUid UidType `xml:"JobUid"`
   JobName string `xml:"JobName"`
   JobType string `xml:"JobType"`
   CreationTimeUTC DateTime `xml:"CreationTimeUTC"`
   EndTimeUTC DateTime `xml:"EndTimeUTC"`
   State string `xml:"State"`
   Result string `xml:"Result"`
   Progress int `xml:"Progress"`
   FailureMessage string `xml:"FailureMessage"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l JobSessionEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * BackupJobSessionEntityType 
 * Not validated 
 */
type BackupJobSessionEntityType struct { 
   IsRetry bool `xml:"IsRetry"`
   //Inhereting from JobSessionEntityType
   JobUid UidType `xml:"JobUid"`
   JobName string `xml:"JobName"`
   JobType string `xml:"JobType"`
   CreationTimeUTC DateTime `xml:"CreationTimeUTC"`
   EndTimeUTC DateTime `xml:"EndTimeUTC"`
   State string `xml:"State"`
   Result string `xml:"Result"`
   Progress int `xml:"Progress"`
   FailureMessage string `xml:"FailureMessage"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l BackupJobSessionEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * ReplicaJobSessionEntityType 
 * Not validated 
 */
type ReplicaJobSessionEntityType struct { 
   IsRetry bool `xml:"IsRetry"`
   //Inhereting from JobSessionEntityType
   JobUid UidType `xml:"JobUid"`
   JobName string `xml:"JobName"`
   JobType string `xml:"JobType"`
   CreationTimeUTC DateTime `xml:"CreationTimeUTC"`
   EndTimeUTC DateTime `xml:"EndTimeUTC"`
   State string `xml:"State"`
   Result string `xml:"Result"`
   Progress int `xml:"Progress"`
   FailureMessage string `xml:"FailureMessage"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l ReplicaJobSessionEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * RestoreSessionEntityType 
 * Not validated 
 */
type RestoreSessionEntityType struct { 
   RestoredObjRef HierarchyObjRefType `xml:"RestoredObjRef"`
   VmDisplayName string `xml:"VmDisplayName,attr"`
   //Inhereting from JobSessionEntityType
   JobUid UidType `xml:"JobUid"`
   JobName string `xml:"JobName"`
   JobType string `xml:"JobType"`
   CreationTimeUTC DateTime `xml:"CreationTimeUTC"`
   EndTimeUTC DateTime `xml:"EndTimeUTC"`
   State string `xml:"State"`
   Result string `xml:"Result"`
   Progress int `xml:"Progress"`
   FailureMessage string `xml:"FailureMessage"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l RestoreSessionEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * BackupTaskSessionEntityType 
 * Not validated 
 */
type BackupTaskSessionEntityType struct { 
   JobSessionUid UidType `xml:"JobSessionUid"`
   CreationTimeUTC DateTime `xml:"CreationTimeUTC"`
   EndTimeUTC DateTime `xml:"EndTimeUTC"`
   State string `xml:"State"`
   Result string `xml:"Result"`
   Reason string `xml:"Reason"`
   TotalSize int64 `xml:"TotalSize"`
   VmDisplayName string `xml:"VmDisplayName,attr"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l BackupTaskSessionEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * BackupTaskSessionEntityListType 
 * Not validated 
 */
type BackupTaskSessionEntityListType struct { 
   BackupTaskSession []BackupTaskSessionEntityType `xml:"BackupTaskSession"`
   //Inhereting from ListType
}

/*
 * ReplicaTaskSessionEntityType 
 * Not validated 
 */
type ReplicaTaskSessionEntityType struct { 
   JobSessionUid UidType `xml:"JobSessionUid"`
   CreationTimeUTC DateTime `xml:"CreationTimeUTC"`
   EndTimeUTC DateTime `xml:"EndTimeUTC"`
   State string `xml:"State"`
   Result string `xml:"Result"`
   Reason string `xml:"Reason"`
   TotalSize int64 `xml:"TotalSize"`
   VmDisplayName string `xml:"VmDisplayName,attr"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l ReplicaTaskSessionEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * ReplicaTaskSessionEntityListType 
 * Not validated 
 */
type ReplicaTaskSessionEntityListType struct { 
   ReplicaTaskSession []ReplicaTaskSessionEntityType `xml:"ReplicaTaskSession"`
   //Inhereting from ListType
}

/*
 * PlainCredentialsType 
 * Not validated 
 */
type PlainCredentialsType struct { 
   UserName string `xml:"UserName"`
   Password string `xml:"Password"`
   //Inhereting from InfoType
}

/*
 * CredentialsInfoType 
 * Not validated 
 */
type CredentialsInfoType struct { 
   Id string `xml:"Id"`
   Username string `xml:"Username"`
   Description string `xml:"Description"`
   Password string `xml:"Password"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CredentialsInfoType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * CredentialsInfoListType 
 * Not validated 
 */
type CredentialsInfoListType struct { 
   CredentialsInfo []CredentialsInfoType `xml:"CredentialsInfo"`
   //Inhereting from ListType
}

/*
 * CredentialsInfoSpecType 
 * Not validated 
 */
type CredentialsInfoSpecType struct { 
   Username string `xml:"Username"`
   Description string `xml:"Description"`
   Password string `xml:"Password"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}

/*
 * PasswordKeyInfoType 
 * Not validated 
 */
type PasswordKeyInfoType struct { 
   Id string `xml:"Id"`
   Hint string `xml:"Hint"`
   LastModificationDate DateTime `xml:"LastModificationDate"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l PasswordKeyInfoType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * PasswordKeyInfoListType 
 * Not validated 
 */
type PasswordKeyInfoListType struct { 
   PasswordKeyInfo []PasswordKeyInfoType `xml:"PasswordKeyInfo"`
   //Inhereting from ListType
}

/*
 * PasswordKeyInfoSpecType 
 * Not validated 
 */
type PasswordKeyInfoSpecType struct { 
   Hint string `xml:"Hint"`
   Password string `xml:"Password"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}

/*
 * JobEntityListType 
 * Not validated 
 */
type JobEntityListType struct { 
   Job []JobEntityType `xml:"Job"`
   //Inhereting from ListType
}

/*
 * FailoverPlanEntityListType 
 * Not validated 
 */
type FailoverPlanEntityListType struct { 
   FailoverPlan []FailoverPlanEntityType `xml:"FailoverPlan"`
   //Inhereting from ListType
}

/*
 * JobEntityType 
 * Not validated 
 */
type JobEntityType struct { 
   JobType string `xml:"JobType"`
   Platform string `xml:"Platform"`
   Description string `xml:"Description"`
   ScheduleConfigured bool `xml:"ScheduleConfigured"`
   ScheduleEnabled bool `xml:"ScheduleEnabled"`
   NextRun DateTime `xml:"NextRun"`
   JobScheduleOptions JobScheduleOptionsInfoType `xml:"JobScheduleOptions"`
   JobInfo JobEntityTypeNestedJobInfo `xml:"JobInfo"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l JobEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * JobEntityTypeNestedJobInfo 
 * Not validated 
 */
type JobEntityTypeNestedJobInfo struct { 
   BackupJobInfo BackupJobInfoType `xml:"BackupJobInfo"`
   FileCopyJobInfo FileCopyJobInfoType `xml:"FileCopyJobInfo"`
   ReplicaJobInfo ReplicaJobInfoType `xml:"ReplicaJobInfo"`
}

/*
 * FailoverPlanEntityType 
 * Not validated 
 */
type FailoverPlanEntityType struct { 
   Description string `xml:"Description"`
   FailoverPlanInfo FailoverPlanInfoType `xml:"FailoverPlanInfo"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l FailoverPlanEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * JobManagementSpecType 
 * Not validated 
 */
type JobManagementSpecType struct { 
   Credentials PlainCredentialsType `xml:"Credentials"`
   Force bool `xml:"Force"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}

/*
 * FailoverPlanManagementSpecType 
 * Not validated 
 */
type FailoverPlanManagementSpecType struct { 
   StartNow bool `xml:"StartNow"`
   StartDate DateTime `xml:"StartDate"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}

/*
 * BackupJobInfoType 
 * Not validated 
 */
type BackupJobInfoType struct { 
   Includes ObjectInJobListType `xml:"Includes"`
   GuestProcessingOptions GuestProcessingOptionsType `xml:"GuestProcessingOptions"`
   AdvancedStorageOptions AdvancedStorageOptionsType `xml:"AdvancedStorageOptions"`
   //Inhereting from InfoType
}

/*
 * FailoverPlanInfoType 
 * Not validated 
 */
type FailoverPlanInfoType struct { 
   Includes FailoveredVmListType `xml:"Includes"`
   //Inhereting from InfoType
}

/*
 * ObjectInJobListType 
 * Not validated 
 */
type ObjectInJobListType struct { 
   ObjectInJob []ObjectInJobType `xml:"ObjectInJob"`
   //Inhereting from ListType
}

/*
 * FailoveredVmListType 
 * Not validated 
 */
type FailoveredVmListType struct { 
   FailoveredVm []FailoveredVmType `xml:"FailoveredVm"`
   //Inhereting from ListType
}

/*
 * ObjectInJobType 
 * Not validated 
 */
type ObjectInJobType struct { 
   ObjectInJobId string `xml:"ObjectInJobId"`
   HierarchyObjRef HierarchyObjRefType `xml:"HierarchyObjRef"`
   Name string `xml:"Name"`
   DisplayName string `xml:"DisplayName"`
   Order int `xml:"Order"`
   GuestProcessingOptions GuestProcessingOptionsType `xml:"GuestProcessingOptions"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l ObjectInJobType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * FailoveredVmType 
 * Not validated 
 */
type FailoveredVmType struct { 
   FailoveredVmId string `xml:"FailoveredVmId"`
   HierarchyObjRef HierarchyObjRefType `xml:"HierarchyObjRef"`
   Name string `xml:"Name"`
   DisplayName string `xml:"DisplayName"`
   Order int `xml:"Order"`
   GuestProcessingOptions GuestProcessingOptionsType `xml:"GuestProcessingOptions"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l FailoveredVmType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * CreateObjectInJobSpecType 
 * Not validated 
 */
type CreateObjectInJobSpecType struct { 
   HierarchyObjRef HierarchyObjRefType `xml:"HierarchyObjRef"`
   HierarchyObjName string `xml:"HierarchyObjName"`
   Order int `xml:"Order"`
   GuestProcessingOptions GuestProcessingOptionsType `xml:"GuestProcessingOptions"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}

/*
 * JobItemVssOptionsType 
 * Not validated 
 */
type JobItemVssOptionsType struct { 
   Enabled bool `xml:"Enabled"`
   IgnoreErrors bool `xml:"IgnoreErrors"`
   GuestFSIndexingType string `xml:"GuestFSIndexingType"`
   Credentials PlainCredentialsType `xml:"Credentials"`
   TransactionLogsTruncation string `xml:"TransactionLogsTruncation"`
   IsFirstUsage bool `xml:"IsFirstUsage"`
   IncludedIndexingFolders StringListType `xml:"IncludedIndexingFolders"`
   ExcludedIndexingFolders StringListType `xml:"ExcludedIndexingFolders"`
   CredentialsId string `xml:"CredentialsId"`
   //Inhereting from InfoType
}

/*
 * GuestProcessingOptionsType 
 * Not validated 
 */
type GuestProcessingOptionsType struct { 
   AppAwareProcessingMode string `xml:"AppAwareProcessingMode"`
   FileSystemIndexingMode string `xml:"FileSystemIndexingMode"`
   IncludedIndexingFolders StringListType `xml:"IncludedIndexingFolders"`
   ExcludedIndexingFolders StringListType `xml:"ExcludedIndexingFolders"`
   CredentialsId string `xml:"CredentialsId"`
   VssSnapshotOptions VssSnapshotOptionsType `xml:"VssSnapshotOptions"`
   WindowsGuestFSIndexingOptions WindowsGuestFSIndexingOptionsType `xml:"WindowsGuestFSIndexingOptions"`
   LinuxGuestFSIndexingOptions LinuxGuestFSIndexingOptionsType `xml:"LinuxGuestFSIndexingOptions"`
   SqlBackupOptions SqlBackupOptionsType `xml:"SqlBackupOptions"`
   WindowsCredentialsId string `xml:"WindowsCredentialsId"`
   LinuxCredentialsId string `xml:"LinuxCredentialsId"`
   FSFileExcludeOptions FSFileExcludeOptionsType `xml:"FSFileExcludeOptions"`
   OracleBackupOptions OracleBackupOptionsType `xml:"OracleBackupOptions"`
   //Inhereting from InfoType
}

/*
 * VssSnapshotOptionsType 
 * Not validated 
 */
type VssSnapshotOptionsType struct { 
   VssSnapshotMode string `xml:"VssSnapshotMode"`
   IsCopyOnly bool `xml:"IsCopyOnly"`
   //Inhereting from InfoType
}

/*
 * WindowsGuestFSIndexingOptionsType 
 * Not validated 
 */
type WindowsGuestFSIndexingOptionsType struct { 
   FileSystemIndexingMode string `xml:"FileSystemIndexingMode"`
   IncludedIndexingFolders StringListType `xml:"IncludedIndexingFolders"`
   ExcludedIndexingFolders StringListType `xml:"ExcludedIndexingFolders"`
   //Inhereting from InfoType
}

/*
 * LinuxGuestFSIndexingOptionsType 
 * Not validated 
 */
type LinuxGuestFSIndexingOptionsType struct { 
   FileSystemIndexingMode string `xml:"FileSystemIndexingMode"`
   IncludedIndexingFolders StringListType `xml:"IncludedIndexingFolders"`
   ExcludedIndexingFolders StringListType `xml:"ExcludedIndexingFolders"`
   //Inhereting from InfoType
}

/*
 * SqlBackupOptionsType 
 * Not validated 
 */
type SqlBackupOptionsType struct { 
   TransactionLogsProcessing string `xml:"TransactionLogsProcessing"`
   BackupLogsFrequencyMin int `xml:"BackupLogsFrequencyMin"`
   UseDbBackupRetention bool `xml:"UseDbBackupRetention"`
   RetainDays int `xml:"RetainDays"`
   //Inhereting from InfoType
}

/*
 * AdvancedStorageOptionsType 
 * Not validated 
 */
type AdvancedStorageOptionsType struct { 
   PasswordKeyId string `xml:"PasswordKeyId"`
   //Inhereting from InfoType
}

/*
 * FSFileExcludeOptionsType 
 * Not validated 
 */
type FSFileExcludeOptionsType struct { 
   BackupScope int `xml:"BackupScope"`
   IncludeList StringListType `xml:"IncludeList"`
   ExcludeList StringListType `xml:"ExcludeList"`
   //Inhereting from InfoType
}

/*
 * OracleBackupOptionsType 
 * Not validated 
 */
type OracleBackupOptionsType struct { 
   BackupLogsEnabled bool `xml:"BackupLogsEnabled"`
   BackupLogsFrequencyMin int `xml:"BackupLogsFrequencyMin"`
   UseDbBackupRetention bool `xml:"UseDbBackupRetention"`
   RetainDays int `xml:"RetainDays"`
   ArchivedLogsTruncation string `xml:"ArchivedLogsTruncation"`
   ArchivedLogsMaxAgeHours int `xml:"ArchivedLogsMaxAgeHours"`
   ArchivedLogsMaxSizeMb int `xml:"ArchivedLogsMaxSizeMb"`
   SysdbaCredsId string `xml:"SysdbaCredsId"`
   ProxyAutoSelect bool `xml:"ProxyAutoSelect"`
   //Inhereting from InfoType
}

/*
 * StringListType 
 * Not validated 
 */
type StringListType struct { 
   Path []string `xml:"Path"`
}

/*
 * JobScheduleOptionsInfoType 
 * Not validated 
 */
type JobScheduleOptionsInfoType struct { 
   RetryOptions JobScheduleRetryOptionsType `xml:"RetryOptions"`
   WaitForBackupCompletion bool `xml:"WaitForBackupCompletion"`
   BackupCompetitionWaitingPeriodMin int `xml:"BackupCompetitionWaitingPeriodMin"`
   OptionsDaily JobScheduleDailyOptionsType `xml:"OptionsDaily"`
   OptionsMonthly JobScheduleMonthlyOptionsType `xml:"OptionsMonthly"`
   OptionsPeriodically JobSchedulePeriodicallyOptionsType `xml:"OptionsPeriodically"`
   OptionsContinuous JobScheduleContinuousOptionsType `xml:"OptionsContinuous"`
   OptionsBackupWindow JobScheduleBackupWindowOptionsType `xml:"OptionsBackupWindow"`
   OptionsDaisyChaining JobScheduleDaisyChainingOptionsType `xml:"OptionsDaisyChaining"`
   //Inhereting from InfoType
}

/*
 * JobScheduleRetryOptionsType 
 * Not validated 
 */
type JobScheduleRetryOptionsType struct { 
   RetryTimes int `xml:"RetryTimes"`
   RetryTimeout int `xml:"RetryTimeout"`
   RetrySpecified bool `xml:"RetrySpecified"`
   //Inhereting from InfoType
}

/*
 * JobScheduleDailyOptionsType 
 * Not validated 
 */
type JobScheduleDailyOptionsType struct { 
   Kind string `xml:"Kind"`
   Days DaysOfWeekEnumeration `xml:"Days"`
   Time Time `xml:"Time"`
   TimeOffsetUtc int `xml:"TimeOffsetUtc"`
   Enabled bool `xml:"Enabled,attr"`
   //Inhereting from InfoType
}

/*
 * JobScheduleMonthlyOptionsType 
 * Not validated 
 */
type JobScheduleMonthlyOptionsType struct { 
   Time Time `xml:"Time"`
   TimeOffsetUtc int `xml:"TimeOffsetUtc"`
   DayNumberInMonth string `xml:"DayNumberInMonth"`
   DayOfWeek string `xml:"DayOfWeek"`
   Months JobScheduleMonthEnumeration `xml:"Months"`
   DayOfMonth int `xml:"DayOfMonth"`
   Enabled bool `xml:"Enabled,attr"`
   //Inhereting from InfoType
}

/*
 * JobSchedulePeriodicallyOptionsType 
 * Not validated 
 */
type JobSchedulePeriodicallyOptionsType struct { 
   Kind string `xml:"Kind"`
   FullPeriod int `xml:"FullPeriod"`
   Schedule TimePeriodsType `xml:"Schedule"`
   Enabled bool `xml:"Enabled,attr"`
   //Inhereting from InfoType
}

/*
 * TimePeriodsType 
 * Not validated 
 */
type TimePeriodsType struct { 
   Day TimePeriodsTypeNestedDay `xml:"Day"`
   //Inhereting from InfoType
}

/*
 * TimePeriodsTypeNestedDay 
 * Not validated 
 */
type TimePeriodsTypeNestedDay struct { 
   Name string `xml:"Name,attr"`
}

/*
 * JobScheduleContinuousOptionsType 
 * Not validated 
 */
type JobScheduleContinuousOptionsType struct { 
   Enabled bool `xml:"Enabled,attr"`
   //Inhereting from InfoType
}

/*
 * JobScheduleBackupWindowOptionsType 
 * Not validated 
 */
type JobScheduleBackupWindowOptionsType struct { 
   TimePeriods TimePeriodsType `xml:"TimePeriods"`
   Enabled bool `xml:"Enabled,attr"`
   //Inhereting from InfoType
}

/*
 * JobScheduleDaisyChainingOptionsType 
 * Not validated 
 */
type JobScheduleDaisyChainingOptionsType struct { 
   PreviousJobUid UidType `xml:"PreviousJobUid"`
   Enabled bool `xml:"Enabled,attr"`
   //Inhereting from InfoType
}

/*
 * FileCopyJobInfoType 
 * Not validated 
 */
type FileCopyJobInfoType struct { 
   //Inhereting from InfoType
}

/*
 * ReplicaJobInfoType 
 * Not validated 
 */
type ReplicaJobInfoType struct { 
   Includes ObjectInJobListType `xml:"Includes"`
   GuestProcessingOptions GuestProcessingOptionsType `xml:"GuestProcessingOptions"`
   //Inhereting from InfoType
}

/*
 * TaskType 
 * Not validated 
 */
type TaskType struct { 
   TaskId string `xml:"TaskId"`
   State string `xml:"State"`
   Operation string `xml:"Operation"`
   Result TaskResultInfoType `xml:"Result"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l TaskType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * TaskListType 
 * Not validated 
 */
type TaskListType struct { 
   Task []TaskType `xml:"Task"`
   //Inhereting from ListType
}

/*
 * QueryResultType 
 * Not validated 
 */
type QueryResultType struct { 
   Refs EntityReferenceListType `xml:"Refs"`
   Entities EntitiesType `xml:"Entities"`
   Resources ResourcesType `xml:"Resources"`
   Links LinkListType `xml:"Links"`
   PagingInfo PagingInfoType `xml:"PagingInfo"`
}
func (l QueryResultType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * PagingInfoType 
 * Not validated 
 */
type PagingInfoType struct { 
   Links LinkListType `xml:"Links"`
   PageNum int `xml:"PageNum,attr"`
   PageSize int `xml:"PageSize,attr"`
   PagesCount int `xml:"PagesCount,attr"`
   //Inhereting from InfoType
}
func (l PagingInfoType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * BackupEntityType 
 * Not validated 
 */
type BackupEntityType struct { 
   Ref EntityReferenceType `xml:"Ref"`
   Platform string `xml:"Platform"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l BackupEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * BackupEntityListType 
 * Not validated 
 */
type BackupEntityListType struct { 
   Backup []BackupEntityType `xml:"Backup"`
   //Inhereting from ListType
}

/*
 * ReplicaEntityType 
 * Not validated 
 */
type ReplicaEntityType struct { 
   Ref EntityReferenceType `xml:"Ref"`
   Platform string `xml:"Platform"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l ReplicaEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * ReplicaEntityListType 
 * Not validated 
 */
type ReplicaEntityListType struct { 
   Replica []ReplicaEntityType `xml:"Replica"`
   //Inhereting from ListType
}

/*
 * RestorePointEntityType 
 * Not validated 
 */
type RestorePointEntityType struct { 
   Ref EntityReferenceType `xml:"Ref"`
   BackupDateUTC DateTime `xml:"BackupDateUTC"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l RestorePointEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * RestorePointEntityListType 
 * Not validated 
 */
type RestorePointEntityListType struct { 
   RestorePoint []RestorePointEntityType `xml:"RestorePoint"`
   //Inhereting from ListType
}

/*
 * VmRestorePointEntityType 
 * Not validated 
 */
type VmRestorePointEntityType struct { 
   CreationTimeUTC DateTime `xml:"CreationTimeUTC"`
   VmName string `xml:"VmName"`
   Algorithm string `xml:"Algorithm"`
   PointType string `xml:"PointType"`
   HierarchyObjRef HierarchyObjRefType `xml:"HierarchyObjRef"`
   SqlInfo VmRestorePointSqlInfoType `xml:"SqlInfo"`
   VmDisplayName string `xml:"VmDisplayName,attr"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l VmRestorePointEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * VAppRestorePointEntityType 
 * Not validated 
 */
type VAppRestorePointEntityType struct { 
   CreationTimeUTC DateTime `xml:"CreationTimeUTC"`
   VAppName string `xml:"VAppName"`
   Algorithm string `xml:"Algorithm"`
   PointType string `xml:"PointType"`
   HierarchyObjRef HierarchyObjRefType `xml:"HierarchyObjRef"`
   VAppDisplayName string `xml:"VAppDisplayName,attr"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l VAppRestorePointEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * VmRestorePointEntityListType 
 * Not validated 
 */
type VmRestorePointEntityListType struct { 
   VmRestorePoint []VmRestorePointEntityType `xml:"VmRestorePoint"`
   //Inhereting from ListType
}

/*
 * VAppRestorePointEntityListType 
 * Not validated 
 */
type VAppRestorePointEntityListType struct { 
   VAppRestorePoint []VAppRestorePointEntityType `xml:"VAppRestorePoint"`
   //Inhereting from ListType
}

/*
 * VmReplicaPointEntityType 
 * Not validated 
 */
type VmReplicaPointEntityType struct { 
   CreationTimeUTC DateTime `xml:"CreationTimeUTC"`
   VmName string `xml:"VmName"`
   Algorithm string `xml:"Algorithm"`
   PointType string `xml:"PointType"`
   VmDisplayName string `xml:"VmDisplayName,attr"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l VmReplicaPointEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * VmReplicaPointEntityListType 
 * Not validated 
 */
type VmReplicaPointEntityListType struct { 
   VmRestorePoint []VmReplicaPointEntityType `xml:"VmRestorePoint"`
   //Inhereting from ListType
}

/*
 * VmRestorePointMountType 
 * Not validated 
 */
type VmRestorePointMountType struct { 
   FSRoots DirectoryEntryListType `xml:"FSRoots"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l VmRestorePointMountType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * VmRestorePointMountListType 
 * Not validated 
 */
type VmRestorePointMountListType struct { 
   VmRestorePointMount []VmRestorePointMountType `xml:"VmRestorePointMount"`
   //Inhereting from ListType
}

/*
 * VmReplicaPointMountType 
 * Not validated 
 */
type VmReplicaPointMountType struct { 
   FSRoots DirectoryEntryListType `xml:"FSRoots"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l VmReplicaPointMountType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * VmReplicaPointMountListType 
 * Not validated 
 */
type VmReplicaPointMountListType struct { 
   VmReplicaPointMount []VmReplicaPointMountType `xml:"VmReplicaPointMount"`
   //Inhereting from ListType
}

/*
 * CatalogVmEntityType 
 * Not validated 
 */
type CatalogVmEntityType struct { 
   VmDisplayName string `xml:"VmDisplayName,attr"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CatalogVmEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * CatalogVmEntityListType 
 * Not validated 
 */
type CatalogVmEntityListType struct { 
   CatalogVm []CatalogVmEntityType `xml:"CatalogVm"`
   //Inhereting from ListType
}

/*
 * CatalogVmRestorePointEntityType 
 * Not validated 
 */
type CatalogVmRestorePointEntityType struct { 
   BackupDateUTC DateTime `xml:"BackupDateUTC"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CatalogVmRestorePointEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * CatalogVmRestorePointEntityListType 
 * Not validated 
 */
type CatalogVmRestorePointEntityListType struct { 
   CatalogVmRestorePoint []CatalogVmRestorePointEntityType `xml:"CatalogVmRestorePoint"`
   //Inhereting from ListType
}

/*
 * FileSystemEntryType 
 * Not validated 
 */
type FileSystemEntryType struct { 
   FileEntry []FileEntryType `xml:"FileEntry"`
   DirectoryEntry []DirectoryEntryType `xml:"DirectoryEntry"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l FileSystemEntryType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * DirectoryEntryType 
 * Not validated 
 */
type DirectoryEntryType struct { 
   Path string `xml:"Path"`
   Name string `xml:"Name"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l DirectoryEntryType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * FileEntryType 
 * Not validated 
 */
type FileEntryType struct { 
   Path string `xml:"Path"`
   Name string `xml:"Name"`
   Size int64 `xml:"Size"`
   Owner string `xml:"Owner"`
   ModifiedDateUTC DateTime `xml:"ModifiedDateUTC"`
   Actions LinkListType `xml:"Actions"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l FileEntryType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * FileEntryListType 
 * Not validated 
 */
type FileEntryListType struct { 
   FileEntry FileEntryType `xml:"FileEntry"`
   //Inhereting from ListType
}

/*
 * HierarchyRootEntityType 
 * Not validated 
 */
type HierarchyRootEntityType struct { 
   HierarchyRootId string `xml:"HierarchyRootId"`
   UniqueId string `xml:"UniqueId"`
   HostType string `xml:"HostType"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l HierarchyRootEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * RepositoryEntityType 
 * Not validated 
 */
type RepositoryEntityType struct { 
   Capacity int64 `xml:"Capacity"`
   FreeSpace int64 `xml:"FreeSpace"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l RepositoryEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * HierarchyRootEntityListType 
 * Not validated 
 */
type HierarchyRootEntityListType struct { 
   HierarchyRoot HierarchyRootEntityType `xml:"HierarchyRoot"`
   //Inhereting from ListType
}

/*
 * RepositoryEntityListType 
 * Not validated 
 */
type RepositoryEntityListType struct { 
   Repository RepositoryEntityType `xml:"Repository"`
   //Inhereting from ListType
}

/*
 * DirectoryEntryListType 
 * Not validated 
 */
type DirectoryEntryListType struct { 
   DirectoryEntry DirectoryEntryType `xml:"DirectoryEntry"`
   //Inhereting from ListType
}

/*
 * FileSystemEntriesType 
 * Not validated 
 */
type FileSystemEntriesType struct { 
   Files FileEntryListType `xml:"Files"`
   Directories DirectoryEntryListType `xml:"Directories"`
   PagingInfo PagingInfoType `xml:"PagingInfo"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l FileSystemEntriesType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * HierarchyItemListType 
 * Not validated 
 */
type HierarchyItemListType struct { 
   HierarchyItem HierarchyItemType `xml:"HierarchyItem"`
   //Inhereting from ListType
}

/*
 * BackupJobSessionEntityListType 
 * Not validated 
 */
type BackupJobSessionEntityListType struct { 
   BackupJobSession []BackupJobSessionEntityType `xml:"BackupJobSession"`
   //Inhereting from ListType
}

/*
 * ReplicaJobSessionEntityListType 
 * Not validated 
 */
type ReplicaJobSessionEntityListType struct { 
   ReplicaJobSession []ReplicaJobSessionEntityType `xml:"ReplicaJobSession"`
   //Inhereting from ListType
}

/*
 * RestoreSessionEntityListType 
 * Not validated 
 */
type RestoreSessionEntityListType struct { 
   RestoreSession []RestoreSessionEntityType `xml:"RestoreSession"`
   //Inhereting from ListType
}

/*
 * FileRestoreSpecType 
 * Not validated 
 */
type FileRestoreSpecType struct { 
   ToOriginalLocation FileRestoreSpecTypeNestedToOriginalLocation `xml:"ToOriginalLocation"`
   ForDirectDownload FileRestoreSpecTypeNestedForDirectDownload `xml:"ForDirectDownload"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}

/*
 * FileRestoreSpecTypeNestedToOriginalLocation 
 * Not validated 
 */
type FileRestoreSpecTypeNestedToOriginalLocation struct { 
}

/*
 * FileRestoreSpecTypeNestedForDirectDownload 
 * Not validated 
 */
type FileRestoreSpecTypeNestedForDirectDownload struct { 
}

/*
 * RestoreSpecType 
 * Not validated 
 */
type RestoreSpecType struct { 
   VmRestoreSpec []VmRestoreSpecInfoType `xml:"VmRestoreSpec"`
   vCloudVmRestoreSpec []vCloudVmRestoreSpecInfoType `xml:"vCloudVmRestoreSpec"`
   vCloudVAppRestoreSpec []vCloudVAppRestoreSpecInfoType `xml:"vCloudVAppRestoreSpec"`
   SqlItemRestoreSpec []SqlItemRestoreSpecInfoType `xml:"SqlItemRestoreSpec"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}

/*
 * VmRestoreSpecInfoType 
 * Not validated 
 */
type VmRestoreSpecInfoType struct { 
   PowerOnAfterRestore bool `xml:"PowerOnAfterRestore"`
   QuickRollback bool `xml:"QuickRollback"`
   //Inhereting from InfoType
}

/*
 * vCloudVmRestoreSpecInfoType 
 * Not validated 
 */
type vCloudVmRestoreSpecInfoType struct { 
   PowerOnAfterRestore bool `xml:"PowerOnAfterRestore"`
   HierarchyRootUid UidType `xml:"HierarchyRootUid"`
   vAppRef HierarchyObjRefType `xml:"vAppRef"`
   VmRestoreParameters vCloudVmRestoreParametersInfoType `xml:"VmRestoreParameters"`
   //Inhereting from InfoType
}

/*
 * vCloudVAppRestoreSpecInfoType 
 * Not validated 
 */
type vCloudVAppRestoreSpecInfoType struct { 
   PowerOnAfterRestore bool `xml:"PowerOnAfterRestore"`
   HierarchyRootUid UidType `xml:"HierarchyRootUid"`
   OrgVdcRef HierarchyObjRefType `xml:"OrgVdcRef"`
   vAppNewName string `xml:"vAppNewName"`
   VmsRestoreParameters vCloudVAppRestoreSpecInfoTypeNestedVmsRestoreParameters `xml:"VmsRestoreParameters"`
   //Inhereting from InfoType
}

/*
 * vCloudVAppRestoreSpecInfoTypeNestedVmsRestoreParameters 
 * Not validated 
 */
type vCloudVAppRestoreSpecInfoTypeNestedVmsRestoreParameters struct { 
   Vm vCloudVmRestoreParametersInfoType `xml:"Vm"`
}

/*
 * vCloudVmRestoreParametersInfoType 
 * Not validated 
 */
type vCloudVmRestoreParametersInfoType struct { 
   VmRestorePointUid UidType `xml:"VmRestorePointUid"`
   VmNewName string `xml:"VmNewName"`
   DatastoreRef HierarchyObjRefType `xml:"DatastoreRef"`
   OrgVdcStorageProfileRef HierarchyObjRefType `xml:"OrgVdcStorageProfileRef"`
   LinkedCloneVmTemplateRef HierarchyObjRefType `xml:"LinkedCloneVmTemplateRef"`
   //Inhereting from InfoType
}

/*
 * BackupJobCloneInfoType 
 * Not validated 
 */
type BackupJobCloneInfoType struct { 
   JobName string `xml:"JobName"`
   FolderName string `xml:"FolderName"`
   RepositoryUid UidType `xml:"RepositoryUid"`
   //Inhereting from InfoType
}

/*
 * ReplicaJobCloneInfoType 
 * Not validated 
 */
type ReplicaJobCloneInfoType struct { 
   JobName string `xml:"JobName"`
   VmSuffix string `xml:"VmSuffix"`
   //Inhereting from InfoType
}

/*
 * JobCloneSpecType 
 * Not validated 
 */
type JobCloneSpecType struct { 
   BackupJobCloneInfo BackupJobCloneInfoType `xml:"BackupJobCloneInfo"`
   ReplicaJobCloneInfo ReplicaJobCloneInfoType `xml:"ReplicaJobCloneInfo"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}

/*
 * TaskResultInfoType 
 * Not validated 
 */
type TaskResultInfoType struct { 
   Message string `xml:"Message"`
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
   FirstChanceExceptionMessage string `xml:"FirstChanceExceptionMessage"`
   StackTrace string `xml:"StackTrace"`
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
   ObjectRef HierarchyObjRefType `xml:"ObjectRef"`
   ObjectType string `xml:"ObjectType"`
   ObjectName string `xml:"ObjectName"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l HierarchyItemType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * QuerySvcType 
 * Not validated 
 */
type QuerySvcType struct { 
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l QuerySvcType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * LookupSvcType 
 * Not validated 
 */
type LookupSvcType struct { 
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l LookupSvcType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * ReportingSvcType 
 * Not validated 
 */
type ReportingSvcType struct { 
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l ReportingSvcType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * EnterpriseRoleEntityType 
 * Not validated 
 */
type EnterpriseRoleEntityType struct { 
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l EnterpriseRoleEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * EnterpriseRoleEntityListType 
 * Not validated 
 */
type EnterpriseRoleEntityListType struct { 
   EnterpriseRoleEntity []EnterpriseRoleEntityType `xml:"EnterpriseRoleEntity"`
   //Inhereting from ListType
}

/*
 * EnterpriseAccountEntityType 
 * Not validated 
 */
type EnterpriseAccountEntityType struct { 
   AccountType AccountTypeEnumeration `xml:"AccountType"`
   Roles EnterpriseRoleEntityListType `xml:"Roles"`
   AllowRestoreAllVms bool `xml:"AllowRestoreAllVms"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l EnterpriseAccountEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * EnterpriseAccountEntityListType 
 * Not validated 
 */
type EnterpriseAccountEntityListType struct { 
   EnterpriseAccountEntity []EnterpriseAccountEntityType `xml:"EnterpriseAccountEntity"`
   //Inhereting from ListType
}

/*
 * EnterpriseAccountHierarchyScopeType 
 * Not validated 
 */
type EnterpriseAccountHierarchyScopeType struct { 
   Name string `xml:"Name"`
   HierarchyRootName string `xml:"HierarchyRootName"`
   Platform string `xml:"Platform"`
   HierarchyObjectType string `xml:"HierarchyObjectType"`
   State string `xml:"State"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l EnterpriseAccountHierarchyScopeType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * EnterpriseAccountHierarchyScopeListType 
 * Not validated 
 */
type EnterpriseAccountHierarchyScopeListType struct { 
   EnterpriseAccountHierarchyScope []EnterpriseAccountHierarchyScopeType `xml:"EnterpriseAccountHierarchyScope"`
   //Inhereting from ListType
}

/*
 * EnterpriseAccountInRoleType 
 * Not validated 
 */
type EnterpriseAccountInRoleType struct { 
   RoleName string `xml:"RoleName"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l EnterpriseAccountInRoleType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * EnterpriseAccountInRoleListType 
 * Not validated 
 */
type EnterpriseAccountInRoleListType struct { 
   EnterpriseAccountInRole []EnterpriseAccountInRoleType `xml:"EnterpriseAccountInRole"`
   //Inhereting from ListType
}

/*
 * HierarchyScopeCreateSpecType 
 * Not validated 
 */
type HierarchyScopeCreateSpecType struct { 
   HierarchyScopeItem []HierarchyScopeCreateSpecItemType `xml:"HierarchyScopeItem"`
   //Inhereting from ListType
}

/*
 * HierarchyScopeCreateSpecItemType 
 * Not validated 
 */
type HierarchyScopeCreateSpecItemType struct { 
   HierarchyObjRef HierarchyObjRefType `xml:"HierarchyObjRef"`
   ObjectName string `xml:"ObjectName"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}

/*
 * EnterpriseAccountInRoleCreateSpecType 
 * Not validated 
 */
type EnterpriseAccountInRoleCreateSpecType struct { 
   EnterpriseRoleUid UidType `xml:"EnterpriseRoleUid"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}

/*
 * EnterpriseAccountInRoleCreateSpecListType 
 * Not validated 
 */
type EnterpriseAccountInRoleCreateSpecListType struct { 
   EnterpriseRole []EnterpriseAccountInRoleCreateSpecType `xml:"EnterpriseRole"`
   //Inhereting from ListType
}

/*
 * EnterpriseSecuritySettingsType 
 * Not validated 
 */
type EnterpriseSecuritySettingsType struct { 
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l EnterpriseSecuritySettingsType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * EnterpriseAccountCreateSpecType 
 * Not validated 
 */
type EnterpriseAccountCreateSpecType struct { 
   AccountType AccountTypeEnumeration `xml:"AccountType"`
   AccountName string `xml:"AccountName"`
   Roles EnterpriseAccountInRoleCreateSpecListType `xml:"Roles"`
   AllowRestoreAllVms bool `xml:"AllowRestoreAllVms"`
   HierarchyScopeObjects HierarchyScopeCreateSpecType `xml:"HierarchyScopeObjects"`
   FlrSettings FileRestoreSettingsSpecsType `xml:"FlrSettings"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}

/*
 * FileRestoreSettingsSpecsType 
 * Not validated 
 */
type FileRestoreSettingsSpecsType struct { 
   FlrInlaceOnly bool `xml:"FlrInlaceOnly"`
   FlrExtentionRestrictions string `xml:"FlrExtentionRestrictions"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}

/*
 * RebuildScopeJobSpecType 
 * Not validated 
 */
type RebuildScopeJobSpecType struct { 
   RebuildAll RebuildScopeJobSpecTypeNestedRebuildAll `xml:"RebuildAll"`
   RebuildUnprocessed RebuildScopeJobSpecTypeNestedRebuildUnprocessed `xml:"RebuildUnprocessed"`
   RebuildForCurrentUser RebuildScopeJobSpecTypeNestedRebuildForCurrentUser `xml:"RebuildForCurrentUser"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}

/*
 * RebuildScopeJobSpecTypeNestedRebuildAll 
 * Not validated 
 */
type RebuildScopeJobSpecTypeNestedRebuildAll struct { 
}

/*
 * RebuildScopeJobSpecTypeNestedRebuildUnprocessed 
 * Not validated 
 */
type RebuildScopeJobSpecTypeNestedRebuildUnprocessed struct { 
}

/*
 * RebuildScopeJobSpecTypeNestedRebuildForCurrentUser 
 * Not validated 
 */
type RebuildScopeJobSpecTypeNestedRebuildForCurrentUser struct { 
}

/*
 * WanAcceleratorEntityType 
 * Not validated 
 */
type WanAcceleratorEntityType struct { 
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
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l WanAcceleratorEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * WanAcceleratorEntityListType 
 * Not validated 
 */
type WanAcceleratorEntityListType struct { 
   WanAccelerator []WanAcceleratorEntityType `xml:"WanAccelerator"`
   //Inhereting from ListType
}

/*
 * CloudGatewayEntityType 
 * Not validated 
 */
type CloudGatewayEntityType struct { 
   Enabled bool `xml:"Enabled"`
   DnsNameOrIpAddress string `xml:"DnsNameOrIpAddress"`
   NetworkMode CloudGatewayNetworkingMode `xml:"NetworkMode"`
   ExternalIP string `xml:"ExternalIP"`
   ExternalPort uint16 `xml:"ExternalPort"`
   InternalPort uint16 `xml:"InternalPort"`
   Description string `xml:"Description"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CloudGatewayEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * CloudGatewayEntityListType 
 * Not validated 
 */
type CloudGatewayEntityListType struct { 
   CloudGateway []CloudGatewayEntityType `xml:"CloudGateway"`
   //Inhereting from ListType
}

/*
 * CloudTenantEntityListType 
 * Not validated 
 */
type CloudTenantEntityListType struct { 
   CloudTenant []CloudTenantEntityType `xml:"CloudTenant"`
   //Inhereting from ListType
}

/*
 * CloudTenantEntityType 
 * Not validated 
 */
type CloudTenantEntityType struct { 
   Password string `xml:"Password"`
   Description string `xml:"Description"`
   Enabled bool `xml:"Enabled"`
   LeaseOptions CloudTenantLeaseOptionsType `xml:"LeaseOptions"`
   Resources CloudTenantResourceListType `xml:"Resources"`
   LastResult string `xml:"LastResult"`
   LastActive DateTime `xml:"LastActive"`
   VmCount int `xml:"VmCount"`
   ComputeResources CloudTenantComputeResourceListType `xml:"ComputeResources"`
   ThrottlingEnabled bool `xml:"ThrottlingEnabled"`
   ThrottlingSpeedLimit int `xml:"ThrottlingSpeedLimit"`
   ThrottlingSpeedUnit string `xml:"ThrottlingSpeedUnit"`
   PublicIpCount int `xml:"PublicIpCount"`
   BackupCount int `xml:"BackupCount"`
   ReplicaCount int `xml:"ReplicaCount"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CloudTenantEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * CloudTenantLeaseOptionsType 
 * Not validated 
 */
type CloudTenantLeaseOptionsType struct { 
   Enabled bool `xml:"Enabled,attr"`
   ExpirationDate DateTime `xml:"ExpirationDate,attr"`
   //Inhereting from InfoType
}

/*
 * CloudTenantResourceType 
 * Not validated 
 */
type CloudTenantResourceType struct { 
   RepositoryQuota CloudTenantRepositoryQuotaInfoType `xml:"RepositoryQuota"`
   Id string `xml:"Id,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CloudTenantResourceType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * CloudTenantResourceListType 
 * Not validated 
 */
type CloudTenantResourceListType struct { 
   CloudTenantResource []CloudTenantResourceType `xml:"CloudTenantResource"`
   //Inhereting from ListType
}

/*
 * CloudTenantRepositoryQuotaInfoType 
 * Not validated 
 */
type CloudTenantRepositoryQuotaInfoType struct { 
   DisplayName string `xml:"DisplayName"`
   RepositoryUid UidType `xml:"RepositoryUid"`
   WanAcceleratorUid UidType `xml:"WanAcceleratorUid"`
   Quota int64 `xml:"Quota"`
   UsedQuota int64 `xml:"UsedQuota"`
   //Inhereting from InfoType
}

/*
 * CreateCloudGatewaySpecType 
 * Not validated 
 */
type CreateCloudGatewaySpecType struct { 
   BackupServerIdOrName string `xml:"BackupServerIdOrName"`
   ServerHostName string `xml:"ServerHostName"`
   Description string `xml:"Description"`
   IncomingPort uint16 `xml:"IncomingPort"`
   ExternalIp string `xml:"ExternalIp"`
   ExternalPort uint16 `xml:"ExternalPort"`
   NetworkType CloudGatewayNetworkingMode `xml:"NetworkType"`
   InternalPort uint16 `xml:"InternalPort"`
   BackupServerUid UidType `xml:"BackupServerUid"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}

/*
 * CreateCloudTenantSpecType 
 * Not validated 
 */
type CreateCloudTenantSpecType struct { 
   BackupServerIdOrName string `xml:"BackupServerIdOrName"`
   Name string `xml:"Name"`
   Description string `xml:"Description"`
   Password string `xml:"Password"`
   Enabled bool `xml:"Enabled"`
   LeaseExpirationDate DateTime `xml:"LeaseExpirationDate"`
   Resources CreateCloudTenantResourceListType `xml:"Resources"`
   ComputeResources CloudTenantComputeResourceCreateListType `xml:"ComputeResources"`
   VmCount int `xml:"VmCount"`
   ThrottlingEnabled bool `xml:"ThrottlingEnabled"`
   ThrottlingSpeedLimit int `xml:"ThrottlingSpeedLimit"`
   ThrottlingSpeedUnit string `xml:"ThrottlingSpeedUnit"`
   PublicIpCount int `xml:"PublicIpCount"`
   BackupServerUid UidType `xml:"BackupServerUid"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}

/*
 * CreateCloudTenantResourceListType 
 * Not validated 
 */
type CreateCloudTenantResourceListType struct { 
   BackupResource []CreateCloudTenantResourceSpecType `xml:"BackupResource"`
   //Inhereting from ListType
}

/*
 * CreateCloudTenantResourceSpecType 
 * Not validated 
 */
type CreateCloudTenantResourceSpecType struct { 
   Name string `xml:"Name"`
   RepositoryUid UidType `xml:"RepositoryUid"`
   QuotaMb int `xml:"QuotaMb"`
   WanAcceleratorUid UidType `xml:"WanAcceleratorUid"`
   Folder string `xml:"Folder"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}

/*
 * CloudHardwarePlanEntityListType 
 * Not validated 
 */
type CloudHardwarePlanEntityListType struct { 
   CloudHardwarePlan []CloudHardwarePlanEntityType `xml:"CloudHardwarePlan"`
   //Inhereting from ListType
}

/*
 * CloudHardwarePlanEntityType 
 * Not validated 
 */
type CloudHardwarePlanEntityType struct { 
   Description string `xml:"Description"`
   ProcessorUsageLimitMhz int `xml:"ProcessorUsageLimitMhz"`
   MemoryUsageLimitMb int `xml:"MemoryUsageLimitMb"`
   HardwarePlanDetails CloudHardwarePlanEntityTypeNestedHardwarePlanDetails `xml:"HardwarePlanDetails"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CloudHardwarePlanEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * CloudHardwarePlanEntityTypeNestedHardwarePlanDetails 
 * Not validated 
 */
type CloudHardwarePlanEntityTypeNestedHardwarePlanDetails struct { 
   ViCloudHardwarePlan ViCloudHardwarePlanInfoType `xml:"ViCloudHardwarePlan"`
   HvCloudHardwarePlan HvCloudHardwarePlanInfoType `xml:"HvCloudHardwarePlan"`
}

/*
 * CloudHardwarePlanCreateSpecType 
 * Not validated 
 */
type CloudHardwarePlanCreateSpecType struct { 
   BackupServerUid UidType `xml:"BackupServerUid"`
   Name string `xml:"Name"`
   Description string `xml:"Description"`
   ProcessorUsageLimitMhz int `xml:"ProcessorUsageLimitMhz"`
   MemoryUsageLimitMb int `xml:"MemoryUsageLimitMb"`
   HardwarePlanDetails CloudHardwarePlanCreateSpecTypeNestedHardwarePlanDetails `xml:"HardwarePlanDetails"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}

/*
 * CloudHardwarePlanCreateSpecTypeNestedHardwarePlanDetails 
 * Not validated 
 */
type CloudHardwarePlanCreateSpecTypeNestedHardwarePlanDetails struct { 
   ViCloudHardwarePlan ViCloudHardwarePlanInfoType `xml:"ViCloudHardwarePlan"`
   HvCloudHardwarePlan HvCloudHardwarePlanInfoType `xml:"HvCloudHardwarePlan"`
}

/*
 * ViCloudHardwarePlanInfoType 
 * Not validated 
 */
type ViCloudHardwarePlanInfoType struct { 
   HypervisorHostRef UidType `xml:"HypervisorHostRef"`
   ParentType string `xml:"ParentType"`
   ParentName string `xml:"ParentName"`
   Datastores ViCloudHardwarePlanDatastoreInfoListType `xml:"Datastores"`
   Network CloudHardwarePlanNetworkInfo `xml:"Network"`
   //Inhereting from InfoType
}

/*
 * ViCloudHardwarePlanDatastoreInfoListType 
 * Not validated 
 */
type ViCloudHardwarePlanDatastoreInfoListType struct { 
   Datastore []ViCloudHardwarePlanDatastoreInfoType `xml:"Datastore"`
   //Inhereting from ListType
}

/*
 * ViCloudHardwarePlanDatastoreInfoType 
 * Not validated 
 */
type ViCloudHardwarePlanDatastoreInfoType struct { 
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
   HypervisorHostRef HierarchyObjRefType `xml:"HypervisorHostRef"`
   Volumes HvCloudHardwarePlanVolumesInfoListType `xml:"Volumes"`
   Network CloudHardwarePlanNetworkInfo `xml:"Network"`
   //Inhereting from InfoType
}

/*
 * HvCloudHardwarePlanVolumesInfoListType 
 * Not validated 
 */
type HvCloudHardwarePlanVolumesInfoListType struct { 
   Volume []HvCloudHardwarePlanVolumeInfoType `xml:"Volume"`
   //Inhereting from ListType
}

/*
 * HvCloudHardwarePlanVolumeInfoType 
 * Not validated 
 */
type HvCloudHardwarePlanVolumeInfoType struct { 
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
   CloudFailoverTasks CloudFailoverTaskSessionInfoListType `xml:"CloudFailoverTasks"`
   CloudFailoverPlanName string `xml:"CloudFailoverPlanName,attr"`
   //Inhereting from JobSessionEntityType
   JobUid UidType `xml:"JobUid"`
   JobName string `xml:"JobName"`
   JobType string `xml:"JobType"`
   CreationTimeUTC DateTime `xml:"CreationTimeUTC"`
   EndTimeUTC DateTime `xml:"EndTimeUTC"`
   State string `xml:"State"`
   Result string `xml:"Result"`
   Progress int `xml:"Progress"`
   FailureMessage string `xml:"FailureMessage"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CloudFailoverSessionEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * CloudFailoverSessionEntityListType 
 * Not validated 
 */
type CloudFailoverSessionEntityListType struct { 
   CloudFailoverSession []CloudFailoverSessionEntityType `xml:"CloudFailoverSession"`
   //Inhereting from ListType
}

/*
 * CloudFailoverTaskSessionInfoType 
 * Not validated 
 */
type CloudFailoverTaskSessionInfoType struct { 
   VmReplicaPointLink LinkType `xml:"VmReplicaPointLink"`
   CreationTimeUTC DateTime `xml:"CreationTimeUTC"`
   EndTimeUTC DateTime `xml:"EndTimeUTC"`
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
   CloudFailoverTasks []CloudFailoverTaskSessionInfoType `xml:"CloudFailoverTasks"`
   //Inhereting from ListType
}

/*
 * CloudPublicIpAddressEntityListType 
 * Not validated 
 */
type CloudPublicIpAddressEntityListType struct { 
   CloudPublicIp []CloudPublicIpAddressEntityType `xml:"CloudPublicIp"`
   //Inhereting from ListType
}

/*
 * CloudPublicIpAddressEntityType 
 * Not validated 
 */
type CloudPublicIpAddressEntityType struct { 
   IpAddress string `xml:"IpAddress"`
   TenantUid UidType `xml:"TenantUid"`
   BackupServerUid UidType `xml:"BackupServerUid,attr"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CloudPublicIpAddressEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * CloudPublicIpAddressCreateSpecType 
 * Not validated 
 */
type CloudPublicIpAddressCreateSpecType struct { 
   BackupServerUid UidType `xml:"BackupServerUid"`
   IpAddressLowerBound string `xml:"IpAddressLowerBound"`
   IpAddressUpperBound string `xml:"IpAddressUpperBound"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}

/*
 * CloudTenantComputeResourceListType 
 * Not validated 
 */
type CloudTenantComputeResourceListType struct { 
   CloudTenantComputeResource []CloudTenantComputeResourceType `xml:"CloudTenantComputeResource"`
   //Inhereting from ListType
}

/*
 * CloudTenantComputeResourceType 
 * Not validated 
 */
type CloudTenantComputeResourceType struct { 
   CloudHardwarePlanUid UidType `xml:"CloudHardwarePlanUid"`
   WanAcceleratorUid UidType `xml:"WanAcceleratorUid"`
   PlatformType string `xml:"PlatformType"`
   UseNetworkFailoverResources bool `xml:"UseNetworkFailoverResources"`
   NetworkAppliance NetworkApplianceInfoType `xml:"NetworkAppliance"`
   ComputeResourceStats ComputeResourceStatsInfoType `xml:"ComputeResourceStats"`
   Id string `xml:"Id,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CloudTenantComputeResourceType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * ComputeResourceStatsInfoType 
 * Not validated 
 */
type ComputeResourceStatsInfoType struct { 
   MemoryUsageMb int `xml:"MemoryUsageMb"`
   CPUCount int `xml:"CPUCount"`
   StorageResourceStats StorageResourceStatsListType `xml:"StorageResourceStats"`
   //Inhereting from InfoType
}

/*
 * StorageResourceStatsListType 
 * Not validated 
 */
type StorageResourceStatsListType struct { 
   StorageResourceStat []StorageResourceStatInfoType `xml:"StorageResourceStat"`
   //Inhereting from ListType
}

/*
 * StorageResourceStatInfoType 
 * Not validated 
 */
type StorageResourceStatInfoType struct { 
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
   Enabled bool `xml:"Enabled,attr"`
   ExpirationDate DateTime `xml:"ExpirationDate,attr"`
   //Inhereting from InfoType
}

/*
 * CloudTenantComputeResourceCreateListType 
 * Not validated 
 */
type CloudTenantComputeResourceCreateListType struct { 
   ComputeResource []CloudTenantComputeResourceCreateSpecType `xml:"ComputeResource"`
   //Inhereting from ListType
}

/*
 * CloudTenantComputeResourceCreateSpecType 
 * Not validated 
 */
type CloudTenantComputeResourceCreateSpecType struct { 
   CloudHardwarePlanUid UidType `xml:"CloudHardwarePlanUid"`
   WanAcceleratorUid UidType `xml:"WanAcceleratorUid"`
   PlatformType string `xml:"PlatformType"`
   UseNetworkFailoverResources bool `xml:"UseNetworkFailoverResources"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}

/*
 * CloudVmReplicaPointEntityType 
 * Not validated 
 */
type CloudVmReplicaPointEntityType struct { 
   CreationTimeUTC DateTime `xml:"CreationTimeUTC"`
   VmName string `xml:"VmName"`
   PointType string `xml:"PointType"`
   State string `xml:"State"`
   VmDisplayName string `xml:"VmDisplayName,attr"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CloudVmReplicaPointEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * CloudVmReplicaPointEntityListType 
 * Not validated 
 */
type CloudVmReplicaPointEntityListType struct { 
   CloudReplica []CloudVmReplicaPointEntityType `xml:"CloudReplica"`
   //Inhereting from ListType
}

/*
 * CloudFailoverPlanEntityListType 
 * Not validated 
 */
type CloudFailoverPlanEntityListType struct { 
   CloudFailoverPlan []CloudFailoverPlanEntityType `xml:"CloudFailoverPlan"`
   //Inhereting from ListType
}

/*
 * CloudFailoverPlanEntityType 
 * Not validated 
 */
type CloudFailoverPlanEntityType struct { 
   TenantUid UidType `xml:"TenantUid"`
   TenantName string `xml:"TenantName"`
   Description string `xml:"Description"`
   CloudFailoverPlanOptions CloudFailoverPlanOptionsInfoType `xml:"CloudFailoverPlanOptions"`
   CloudFailoverPlanInfo CloudFailoverPlanInfoType `xml:"CloudFailoverPlanInfo"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CloudFailoverPlanEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * CloudFailoverPlanManagementSpecType 
 * Not validated 
 */
type CloudFailoverPlanManagementSpecType struct { 
   StartNow bool `xml:"StartNow"`
   StartDate DateTime `xml:"StartDate"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}

/*
 * CloudFailoverPlanInfoType 
 * Not validated 
 */
type CloudFailoverPlanInfoType struct { 
   Includes CloudFailoveredVmListType `xml:"Includes"`
   //Inhereting from InfoType
}

/*
 * CloudFailoveredVmType 
 * Not validated 
 */
type CloudFailoveredVmType struct { 
   FailoverPlanVMId string `xml:"FailoverPlanVMId"`
   Name string `xml:"Name"`
   Order int `xml:"Order"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CloudFailoveredVmType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * CloudFailoveredVmListType 
 * Not validated 
 */
type CloudFailoveredVmListType struct { 
   CloudFailoveredVm []CloudFailoveredVmType `xml:"CloudFailoveredVm"`
   //Inhereting from ListType
}

/*
 * CloudFailoverPlanOptionsInfoType 
 * Not validated 
 */
type CloudFailoverPlanOptionsInfoType struct { 
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
   Ref EntityReferenceType `xml:"Ref"`
   Platform string `xml:"Platform"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CloudReplicaEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * CloudReplicaEntityListType 
 * Not validated 
 */
type CloudReplicaEntityListType struct { 
   CloudReplica []CloudReplicaEntityType `xml:"CloudReplica"`
   //Inhereting from ListType
}

/*
 * CloudConnectServiceType 
 * Not validated 
 */
type CloudConnectServiceType struct { 
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l CloudConnectServiceType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * TenantCredentialsInfoType 
 * Not validated 
 */
type TenantCredentialsInfoType struct { 
   Username string `xml:"Username"`
   Password string `xml:"Password"`
   //Inhereting from InfoType
}

/*
 * VlanConfigurationEntityListType 
 * Not validated 
 */
type VlanConfigurationEntityListType struct { 
   Vlans []VlanConfigurationEntityType `xml:"Vlans"`
   //Inhereting from ListType
}

/*
 * VlanConfigurationEntityType 
 * Not validated 
 */
type VlanConfigurationEntityType struct { 
   HostRef HierarchyObjRefType `xml:"HostRef"`
   PlatformType string `xml:"PlatformType"`
   VlanIdsWithInternetLeftBound int `xml:"VlanIdsWithInternetLeftBound"`
   VlanIdsWithInternetRightBound int `xml:"VlanIdsWithInternetRightBound"`
   VlanIdsWithoutInternetLeftBound int `xml:"VlanIdsWithoutInternetLeftBound"`
   VlanIdsWithoutInternetRightBound int `xml:"VlanIdsWithoutInternetRightBound"`
   SwitchName string `xml:"SwitchName"`
   SwitchId string `xml:"SwitchId"`
   ClusterName string `xml:"ClusterName"`
   ClusterId string `xml:"ClusterId"`
   //Inhereting from EntityType
   Name string `xml:"Name,attr"`
   UID UidType `xml:"UID,attr"`
   //Inhereting from ResourceType
   Links LinkListType `xml:"Links"`
   Href UrlType `xml:"Href,attr"`
   Type string `xml:"Type,attr"`
}
func (l VlanConfigurationEntityType) GetLinks() (*LinkListType) { return &(l.Links) }

/*
 * CloudVlanConfigurationCreateSpecType 
 * Not validated 
 */
type CloudVlanConfigurationCreateSpecType struct { 
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

/*
 * NetworkApplianceInfoType 
 * Not validated 
 */
type NetworkApplianceInfoType struct { 
   Name string `xml:"Name"`
   ProductionNetwork string `xml:"ProductionNetwork"`
   ObtainIPAddressAutomatically bool `xml:"ObtainIPAddressAutomatically"`
   ManualIpAddressSettingsInfoType ManualIpAddressSettingsInfoType `xml:"ManualIpAddressSettingsInfoType"`
   //Inhereting from InfoType
}

/*
 * ManualIpAddressSettingsInfoType 
 * Not validated 
 */
type ManualIpAddressSettingsInfoType struct { 
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
   Databases VmRestorePointSqlDatabaseInfoListType `xml:"Databases"`
   //Inhereting from InfoType
}

/*
 * VmRestorePointSqlDatabaseInfoListType 
 * Not validated 
 */
type VmRestorePointSqlDatabaseInfoListType struct { 
   Database []VmRestorePointSqlDatabaseInfoType `xml:"Database"`
   //Inhereting from ListType
}

/*
 * VmRestorePointSqlDatabaseInfoType 
 * Not validated 
 */
type VmRestorePointSqlDatabaseInfoType struct { 
   Items VmRestorePointSqlDatabaseItemListInfoType `xml:"Items"`
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
   Item []VmRestorePointSqlDatabaseItemInfoType `xml:"Item"`
   //Inhereting from ListType
}

/*
 * VmRestorePointSqlDatabaseItemInfoType 
 * Not validated 
 */
type VmRestorePointSqlDatabaseItemInfoType struct { 
   Type string `xml:"Type,attr"`
   DataPath string `xml:"DataPath,attr"`
   //Inhereting from InfoType
}

/*
 * SqlItemRestoreSpecInfoType 
 * Not validated 
 */
type SqlItemRestoreSpecInfoType struct { 
   IsRestoreToOriginal bool `xml:"IsRestoreToOriginal"`
   ServerName string `xml:"ServerName"`
   VmName string `xml:"VmName"`
   InstanceName string `xml:"InstanceName"`
   FamilyUid string `xml:"FamilyUid"`
   Credentials SqlCredentialsInfoType `xml:"Credentials"`
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
   SqlCredentials PlainCredentialsType `xml:"SqlCredentials"`
   ServerCredentials PlainCredentialsType `xml:"ServerCredentials"`
   UserCredentials PlainCredentialsType `xml:"UserCredentials"`
   UseSqlAuth bool `xml:"UseSqlAuth"`
   //Inhereting from InfoType
}

/*
 * VeeamZipStartupSpecType 
 * Not validated 
 */
type VeeamZipStartupSpecType struct { 
   VmRef HierarchyObjRefType `xml:"VmRef"`
   RepositoryUid UidType `xml:"RepositoryUid"`
   PasswordKeyId string `xml:"PasswordKeyId"`
   CompressionLevel int `xml:"CompressionLevel"`
   DisableGuestQuiescence bool `xml:"DisableGuestQuiescence"`
   BackupRetention FreeBackupRetentionEnumType `xml:"BackupRetention"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}

/*
 * QuickBackupStartupSpecType 
 * Not validated 
 */
type QuickBackupStartupSpecType struct { 
   VmRef HierarchyObjRefType `xml:"VmRef"`
   //Inhereting from SpecType
   //Inhereting from ParamsType
}

