package veeamrestapi


 
import (
 "bytes"
 "fmt"
 )
 
/*
 * LoginSpecType 
 * Not validated 
 */
func (o LoginSpecType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VMwareSSOToken                 : ");buffer.WriteString(fmt.Sprintf("%s",o.VMwareSSOToken))
  if full && o.TenantCredentials != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("TenantCredentials              : ");buffer.WriteString(o.TenantCredentials.FullString(full,depth+1))
  }
   //Inhereting from SpecType
   //Inhereting from ParamsType
  return buffer.String()
}
func (o LoginSpecType) String() (string) { return o.FullString(false,0) }

/*
 * LogonSessionListType 
 * Not validated 
 */
func (o LogonSessionListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.LogonSession {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("LogonSession[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o LogonSessionListType) String() (string) { return o.FullString(false,0) }

/*
 * LogonSessionType 
 * Not validated 
 */
func (o LogonSessionType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("UserName                       : ");buffer.WriteString(fmt.Sprintf("%s",o.UserName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("SessionId                      : ");buffer.WriteString(fmt.Sprintf("%s",o.SessionId))
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o LogonSessionType) String() (string) { return o.FullString(false,0) }

/*
 * SummaryReportType 
 * Not validated 
 */
func (o SummaryReportType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   //Inhereting from ReportResourceType
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o SummaryReportType) String() (string) { return o.FullString(false,0) }

/*
 * OverviewReportFrameType 
 * Not validated 
 */
func (o OverviewReportFrameType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("BackupServers                  : ");buffer.WriteString(fmt.Sprintf("%d",o.BackupServers))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ProxyServers                   : ");buffer.WriteString(fmt.Sprintf("%d",o.ProxyServers))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("RepositoryServers              : ");buffer.WriteString(fmt.Sprintf("%d",o.RepositoryServers))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("RunningJobs                    : ");buffer.WriteString(fmt.Sprintf("%d",o.RunningJobs))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ScheduledJobs                  : ");buffer.WriteString(fmt.Sprintf("%d",o.ScheduledJobs))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("SuccessfulVmLastestStates      : ");buffer.WriteString(fmt.Sprintf("%d",o.SuccessfulVmLastestStates))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("WarningVmLastestStates         : ");buffer.WriteString(fmt.Sprintf("%d",o.WarningVmLastestStates))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FailedVmLastestStates          : ");buffer.WriteString(fmt.Sprintf("%d",o.FailedVmLastestStates))
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o OverviewReportFrameType) String() (string) { return o.FullString(false,0) }

/*
 * VmsOverviewReportFrameType 
 * Not validated 
 */
func (o VmsOverviewReportFrameType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ProtectedVms                   : ");buffer.WriteString(fmt.Sprintf("%d",o.ProtectedVms))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("BackedUpVms                    : ");buffer.WriteString(fmt.Sprintf("%d",o.BackedUpVms))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ReplicatedVms                  : ");buffer.WriteString(fmt.Sprintf("%d",o.ReplicatedVms))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("RestorePoints                  : ");buffer.WriteString(fmt.Sprintf("%d",o.RestorePoints))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FullBackupPointsSize           : ");buffer.WriteString(fmt.Sprintf("%d",o.FullBackupPointsSize))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("IncrementalBackupPointsSize    : ");buffer.WriteString(fmt.Sprintf("%d",o.IncrementalBackupPointsSize))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ReplicaRestorePointsSize       : ");buffer.WriteString(fmt.Sprintf("%d",o.ReplicaRestorePointsSize))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("SourceVmsSize                  : ");buffer.WriteString(fmt.Sprintf("%d",o.SourceVmsSize))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("SuccessBackupPercents          : ");buffer.WriteString(fmt.Sprintf("%d",o.SuccessBackupPercents))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ProtectedVmsReportLink         : ");buffer.WriteString(fmt.Sprintf("%s",o.ProtectedVmsReportLink))
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o VmsOverviewReportFrameType) String() (string) { return o.FullString(false,0) }

/*
 * JobStatisticsReportFrameType 
 * Not validated 
 */
func (o JobStatisticsReportFrameType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("RunningJobs                    : ");buffer.WriteString(fmt.Sprintf("%d",o.RunningJobs))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ScheduledJobs                  : ");buffer.WriteString(fmt.Sprintf("%d",o.ScheduledJobs))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ScheduledBackupJobs            : ");buffer.WriteString(fmt.Sprintf("%d",o.ScheduledBackupJobs))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ScheduledReplicaJobs           : ");buffer.WriteString(fmt.Sprintf("%d",o.ScheduledReplicaJobs))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("TotalJobRuns                   : ");buffer.WriteString(fmt.Sprintf("%d",o.TotalJobRuns))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("SuccessfulJobRuns              : ");buffer.WriteString(fmt.Sprintf("%d",o.SuccessfulJobRuns))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("WarningsJobRuns                : ");buffer.WriteString(fmt.Sprintf("%d",o.WarningsJobRuns))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FailedJobRuns                  : ");buffer.WriteString(fmt.Sprintf("%d",o.FailedJobRuns))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("MaxJobDuration                 : ");buffer.WriteString(fmt.Sprintf("%d",o.MaxJobDuration))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("MaxBackupJobDuration           : ");buffer.WriteString(fmt.Sprintf("%d",o.MaxBackupJobDuration))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("MaxReplicaJobDuration          : ");buffer.WriteString(fmt.Sprintf("%d",o.MaxReplicaJobDuration))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("MaxDurationBackupJobName       : ");buffer.WriteString(fmt.Sprintf("%s",o.MaxDurationBackupJobName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("MaxDurationReplicaJobName      : ");buffer.WriteString(fmt.Sprintf("%s",o.MaxDurationReplicaJobName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("BackupJobStatusReportLink      : ");buffer.WriteString(fmt.Sprintf("%s",o.BackupJobStatusReportLink))
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o JobStatisticsReportFrameType) String() (string) { return o.FullString(false,0) }

/*
 * InfoType 
 * Not validated 
 */
func (o InfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  return buffer.String()
}
func (o InfoType) String() (string) { return o.FullString(false,0) }

/*
 * RepositoryReportFrameType 
 * Not validated 
 */
func (o RepositoryReportFrameType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.Period {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("Period[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("CapacityPlanningReportLink     : ");buffer.WriteString(fmt.Sprintf("%s",o.CapacityPlanningReportLink))
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o RepositoryReportFrameType) String() (string) { return o.FullString(false,0) }

/*
 * RepositoryReportFrameTypeNestedPeriod 
 * Not validated 
 */
func (o RepositoryReportFrameTypeNestedPeriod) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Capacity                       : ");buffer.WriteString(fmt.Sprintf("%d",o.Capacity))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FreeSpace                      : ");buffer.WriteString(fmt.Sprintf("%d",o.FreeSpace))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("BackupSize                     : ");buffer.WriteString(fmt.Sprintf("%d",o.BackupSize))
  return buffer.String()
}
func (o RepositoryReportFrameTypeNestedPeriod) String() (string) { return o.FullString(false,0) }

/*
 * ProcessedVmsReportFrameType 
 * Not validated 
 */
func (o ProcessedVmsReportFrameType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.Day {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("Day[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o ProcessedVmsReportFrameType) String() (string) { return o.FullString(false,0) }

/*
 * ProcessedVmsReportFrameTypeNestedDay 
 * Not validated 
 */
func (o ProcessedVmsReportFrameTypeNestedDay) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Timestamp                      : ");buffer.WriteString(o.Timestamp.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("ReplicatedVms                  : ");buffer.WriteString(fmt.Sprintf("%d",o.ReplicatedVms))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("BackupedVms                    : ");buffer.WriteString(fmt.Sprintf("%d",o.BackupedVms))
  return buffer.String()
}
func (o ProcessedVmsReportFrameTypeNestedDay) String() (string) { return o.FullString(false,0) }

/*
 * EnterpriseManagerType 
 * Not validated 
 */
func (o EnterpriseManagerType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.SupportedVersions != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("SupportedVersions              : ");buffer.WriteString(o.SupportedVersions.FullString(full,depth+1))
  }
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o EnterpriseManagerType) String() (string) { return o.FullString(false,0) }

/*
 * SupportedVersionListType 
 * Not validated 
 */
func (o SupportedVersionListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.SupportedVersion {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("SupportedVersion[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o SupportedVersionListType) String() (string) { return o.FullString(false,0) }

/*
 * SupportedVersionType 
 * Not validated 
 */
func (o SupportedVersionType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
   //Inhereting from InfoType
  return buffer.String()
}
func (o SupportedVersionType) String() (string) { return o.FullString(false,0) }

/*
 * ListType 
 * Not validated 
 */
func (o ListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  return buffer.String()
}
func (o ListType) String() (string) { return o.FullString(false,0) }

/*
 * ReportResourceType 
 * Not validated 
 */
func (o ReportResourceType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o ReportResourceType) String() (string) { return o.FullString(false,0) }

/*
 * LinkListType 
 * Not validated 
 */
func (o LinkListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.Link {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("Link[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
  return buffer.String()
}
func (o LinkListType) String() (string) { return o.FullString(false,0) }

/*
 * LinkType 
 * Not validated 
 */
func (o LinkType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Rel                            : ");buffer.WriteString(fmt.Sprintf("%s",o.Rel))
   //Inhereting from ReferenceType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o LinkType) String() (string) { return o.FullString(false,0) }

/*
 * ReferenceType 
 * Not validated 
 */
func (o ReferenceType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o ReferenceType) String() (string) { return o.FullString(false,0) }

/*
 * ResourceType 
 * Not validated 
 */
func (o ResourceType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o ResourceType) String() (string) { return o.FullString(false,0) }

/*
 * EntityType 
 * Not validated 
 */
func (o EntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o EntityType) String() (string) { return o.FullString(false,0) }

/*
 * EntitiesType 
 * Not validated 
 */
func (o EntitiesType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.Jobs != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Jobs                           : ");buffer.WriteString(o.Jobs.FullString(full,depth+1))
  }
  if full && o.FailoverPlans != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FailoverPlans                  : ");buffer.WriteString(o.FailoverPlans.FullString(full,depth+1))
  }
  if full && o.Backups != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Backups                        : ");buffer.WriteString(o.Backups.FullString(full,depth+1))
  }
  if full && o.Replicas != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Replicas                       : ");buffer.WriteString(o.Replicas.FullString(full,depth+1))
  }
  if full && o.Repositories != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Repositories                   : ");buffer.WriteString(o.Repositories.FullString(full,depth+1))
  }
  if full && o.RestorePoints != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("RestorePoints                  : ");buffer.WriteString(o.RestorePoints.FullString(full,depth+1))
  }
  if full && o.VmRestorePoints != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VmRestorePoints                : ");buffer.WriteString(o.VmRestorePoints.FullString(full,depth+1))
  }
  if full && o.VAppRestorePoints != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VAppRestorePoints              : ");buffer.WriteString(o.VAppRestorePoints.FullString(full,depth+1))
  }
  if full && o.VmReplicaPoints != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VmReplicaPoints                : ");buffer.WriteString(o.VmReplicaPoints.FullString(full,depth+1))
  }
  if full && o.BackupJobSessions != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("BackupJobSessions              : ");buffer.WriteString(o.BackupJobSessions.FullString(full,depth+1))
  }
  if full && o.ReplicaJobSessions != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ReplicaJobSessions             : ");buffer.WriteString(o.ReplicaJobSessions.FullString(full,depth+1))
  }
  if full && o.ReplicaTaskSessions != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ReplicaTaskSessions            : ");buffer.WriteString(o.ReplicaTaskSessions.FullString(full,depth+1))
  }
  if full && o.RestoreSessions != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("RestoreSessions                : ");buffer.WriteString(o.RestoreSessions.FullString(full,depth+1))
  }
  if full && o.HierarchyRoots != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("HierarchyRoots                 : ");buffer.WriteString(o.HierarchyRoots.FullString(full,depth+1))
  }
  if full && o.BackupTaskSessions != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("BackupTaskSessions             : ");buffer.WriteString(o.BackupTaskSessions.FullString(full,depth+1))
  }
  if full && o.BackupServers != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("BackupServers                  : ");buffer.WriteString(o.BackupServers.FullString(full,depth+1))
  }
  if full && o.ManagedServers != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ManagedServers                 : ");buffer.WriteString(o.ManagedServers.FullString(full,depth+1))
  }
  if full && o.EnterpiseRoles != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("EnterpiseRoles                 : ");buffer.WriteString(o.EnterpiseRoles.FullString(full,depth+1))
  }
  if full && o.EnterpiseAccounts != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("EnterpiseAccounts              : ");buffer.WriteString(o.EnterpiseAccounts.FullString(full,depth+1))
  }
  if full && o.WanAccelerators != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("WanAccelerators                : ");buffer.WriteString(o.WanAccelerators.FullString(full,depth+1))
  }
  if full && o.CloudGateways != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CloudGateways                  : ");buffer.WriteString(o.CloudGateways.FullString(full,depth+1))
  }
  if full && o.CloudTenants != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CloudTenants                   : ");buffer.WriteString(o.CloudTenants.FullString(full,depth+1))
  }
  if full && o.Passwords != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Passwords                      : ");buffer.WriteString(o.Passwords.FullString(full,depth+1))
  }
  if full && o.Credentials != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Credentials                    : ");buffer.WriteString(o.Credentials.FullString(full,depth+1))
  }
  return buffer.String()
}
func (o EntitiesType) String() (string) { return o.FullString(false,0) }

/*
 * ResourcesType 
 * Not validated 
 */
func (o ResourcesType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.Files != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Files                          : ");buffer.WriteString(o.Files.FullString(full,depth+1))
  }
  if full && o.Directories != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Directories                    : ");buffer.WriteString(o.Directories.FullString(full,depth+1))
  }
  if full && o.Tasks != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Tasks                          : ");buffer.WriteString(o.Tasks.FullString(full,depth+1))
  }
  if full && o.CredentialsList != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CredentialsList                : ");buffer.WriteString(o.CredentialsList.FullString(full,depth+1))
  }
  if full && o.sInObjectJob != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("sInObjectJob                   : ");buffer.WriteString(o.sInObjectJob.FullString(full,depth+1))
  }
  if full && o.AccountRole != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("AccountRole                    : ");buffer.WriteString(o.AccountRole.FullString(full,depth+1))
  }
  if full && o.AccountHierarchyScope != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("AccountHierarchyScope          : ");buffer.WriteString(o.AccountHierarchyScope.FullString(full,depth+1))
  }
  if full && o.PasswordKeyInfoList != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("PasswordKeyInfoList            : ");buffer.WriteString(o.PasswordKeyInfoList.FullString(full,depth+1))
  }
  return buffer.String()
}
func (o ResourcesType) String() (string) { return o.FullString(false,0) }

/*
 * ParamsType 
 * Not validated 
 */
func (o ParamsType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  return buffer.String()
}
func (o ParamsType) String() (string) { return o.FullString(false,0) }

/*
 * SpecType 
 * Not validated 
 */
func (o SpecType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   //Inhereting from ParamsType
  return buffer.String()
}
func (o SpecType) String() (string) { return o.FullString(false,0) }

/*
 * EntityReferenceListType 
 * Not validated 
 */
func (o EntityReferenceListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.Ref {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("Ref[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o EntityReferenceListType) String() (string) { return o.FullString(false,0) }

/*
 * EntityReferenceType 
 * Not validated 
 */
func (o EntityReferenceType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o EntityReferenceType) String() (string) { return o.FullString(false,0) }

/*
 * BackupServerEntityListType 
 * Not validated 
 */
func (o BackupServerEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.BackupServer {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("BackupServer[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o BackupServerEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * ManagedServerEntityListType 
 * Not validated 
 */
func (o ManagedServerEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.ManagedServer {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("ManagedServer[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o ManagedServerEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * BackupServerEntityType 
 * Not validated 
 */
func (o BackupServerEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Description                    : ");buffer.WriteString(fmt.Sprintf("%s",o.Description))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Port                           : ");buffer.WriteString(fmt.Sprintf("%d",o.Port))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Version                        : ");buffer.WriteString(o.Version.String())
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o BackupServerEntityType) String() (string) { return o.FullString(false,0) }

/*
 * ManagedServerEntityType 
 * Not validated 
 */
func (o ManagedServerEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Description                    : ");buffer.WriteString(fmt.Sprintf("%s",o.Description))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ManagedServerType              : ");buffer.WriteString(fmt.Sprintf("%s",o.ManagedServerType))
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o ManagedServerEntityType) String() (string) { return o.FullString(false,0) }

/*
 * BackupServerSpecType 
 * Not validated 
 */
func (o BackupServerSpecType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Description                    : ");buffer.WriteString(fmt.Sprintf("%s",o.Description))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("DnsNameOrIpAddress             : ");buffer.WriteString(fmt.Sprintf("%s",o.DnsNameOrIpAddress))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Port                           : ");buffer.WriteString(fmt.Sprintf("%d",o.Port))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Username                       : ");buffer.WriteString(fmt.Sprintf("%s",o.Username))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Password                       : ");buffer.WriteString(fmt.Sprintf("%s",o.Password))
   //Inhereting from SpecType
   //Inhereting from ParamsType
  return buffer.String()
}
func (o BackupServerSpecType) String() (string) { return o.FullString(false,0) }

/*
 * JobSessionEntityType 
 * Not validated 
 */
func (o JobSessionEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("JobUid                         : ");buffer.WriteString(o.JobUid.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("JobName                        : ");buffer.WriteString(fmt.Sprintf("%s",o.JobName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("JobType                        : ");buffer.WriteString(fmt.Sprintf("%s",o.JobType))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CreationTimeUTC                : ");buffer.WriteString(o.CreationTimeUTC.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("EndTimeUTC                     : ");buffer.WriteString(o.EndTimeUTC.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("State                          : ");buffer.WriteString(fmt.Sprintf("%s",o.State))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Result                         : ");buffer.WriteString(fmt.Sprintf("%s",o.Result))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Progress                       : ");buffer.WriteString(fmt.Sprintf("%d",o.Progress))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FailureMessage                 : ");buffer.WriteString(fmt.Sprintf("%s",o.FailureMessage))
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o JobSessionEntityType) String() (string) { return o.FullString(false,0) }

/*
 * BackupJobSessionEntityType 
 * Not validated 
 */
func (o BackupJobSessionEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("IsRetry                        : ");buffer.WriteString(fmt.Sprintf("%t",o.IsRetry))
   //Inhereting from JobSessionEntityType
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("JobUid                         : ");buffer.WriteString(o.JobUid.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("JobName                        : ");buffer.WriteString(fmt.Sprintf("%s",o.JobName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("JobType                        : ");buffer.WriteString(fmt.Sprintf("%s",o.JobType))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CreationTimeUTC                : ");buffer.WriteString(o.CreationTimeUTC.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("EndTimeUTC                     : ");buffer.WriteString(o.EndTimeUTC.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("State                          : ");buffer.WriteString(fmt.Sprintf("%s",o.State))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Result                         : ");buffer.WriteString(fmt.Sprintf("%s",o.Result))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Progress                       : ");buffer.WriteString(fmt.Sprintf("%d",o.Progress))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FailureMessage                 : ");buffer.WriteString(fmt.Sprintf("%s",o.FailureMessage))
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o BackupJobSessionEntityType) String() (string) { return o.FullString(false,0) }

/*
 * ReplicaJobSessionEntityType 
 * Not validated 
 */
func (o ReplicaJobSessionEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("IsRetry                        : ");buffer.WriteString(fmt.Sprintf("%t",o.IsRetry))
   //Inhereting from JobSessionEntityType
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("JobUid                         : ");buffer.WriteString(o.JobUid.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("JobName                        : ");buffer.WriteString(fmt.Sprintf("%s",o.JobName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("JobType                        : ");buffer.WriteString(fmt.Sprintf("%s",o.JobType))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CreationTimeUTC                : ");buffer.WriteString(o.CreationTimeUTC.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("EndTimeUTC                     : ");buffer.WriteString(o.EndTimeUTC.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("State                          : ");buffer.WriteString(fmt.Sprintf("%s",o.State))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Result                         : ");buffer.WriteString(fmt.Sprintf("%s",o.Result))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Progress                       : ");buffer.WriteString(fmt.Sprintf("%d",o.Progress))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FailureMessage                 : ");buffer.WriteString(fmt.Sprintf("%s",o.FailureMessage))
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o ReplicaJobSessionEntityType) String() (string) { return o.FullString(false,0) }

/*
 * RestoreSessionEntityType 
 * Not validated 
 */
func (o RestoreSessionEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("RestoredObjRef                 : ");buffer.WriteString(o.RestoredObjRef.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("VmDisplayName                  : ");buffer.WriteString(fmt.Sprintf("%s",o.VmDisplayName))
   //Inhereting from JobSessionEntityType
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("JobUid                         : ");buffer.WriteString(o.JobUid.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("JobName                        : ");buffer.WriteString(fmt.Sprintf("%s",o.JobName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("JobType                        : ");buffer.WriteString(fmt.Sprintf("%s",o.JobType))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CreationTimeUTC                : ");buffer.WriteString(o.CreationTimeUTC.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("EndTimeUTC                     : ");buffer.WriteString(o.EndTimeUTC.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("State                          : ");buffer.WriteString(fmt.Sprintf("%s",o.State))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Result                         : ");buffer.WriteString(fmt.Sprintf("%s",o.Result))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Progress                       : ");buffer.WriteString(fmt.Sprintf("%d",o.Progress))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FailureMessage                 : ");buffer.WriteString(fmt.Sprintf("%s",o.FailureMessage))
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o RestoreSessionEntityType) String() (string) { return o.FullString(false,0) }

/*
 * BackupTaskSessionEntityType 
 * Not validated 
 */
func (o BackupTaskSessionEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("JobSessionUid                  : ");buffer.WriteString(o.JobSessionUid.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CreationTimeUTC                : ");buffer.WriteString(o.CreationTimeUTC.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("EndTimeUTC                     : ");buffer.WriteString(o.EndTimeUTC.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("State                          : ");buffer.WriteString(fmt.Sprintf("%s",o.State))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Result                         : ");buffer.WriteString(fmt.Sprintf("%s",o.Result))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Reason                         : ");buffer.WriteString(fmt.Sprintf("%s",o.Reason))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("TotalSize                      : ");buffer.WriteString(fmt.Sprintf("%d",o.TotalSize))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("VmDisplayName                  : ");buffer.WriteString(fmt.Sprintf("%s",o.VmDisplayName))
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o BackupTaskSessionEntityType) String() (string) { return o.FullString(false,0) }

/*
 * BackupTaskSessionEntityListType 
 * Not validated 
 */
func (o BackupTaskSessionEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.BackupTaskSession {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("BackupTaskSession[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o BackupTaskSessionEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * ReplicaTaskSessionEntityType 
 * Not validated 
 */
func (o ReplicaTaskSessionEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("JobSessionUid                  : ");buffer.WriteString(o.JobSessionUid.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CreationTimeUTC                : ");buffer.WriteString(o.CreationTimeUTC.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("EndTimeUTC                     : ");buffer.WriteString(o.EndTimeUTC.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("State                          : ");buffer.WriteString(fmt.Sprintf("%s",o.State))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Result                         : ");buffer.WriteString(fmt.Sprintf("%s",o.Result))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Reason                         : ");buffer.WriteString(fmt.Sprintf("%s",o.Reason))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("TotalSize                      : ");buffer.WriteString(fmt.Sprintf("%d",o.TotalSize))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("VmDisplayName                  : ");buffer.WriteString(fmt.Sprintf("%s",o.VmDisplayName))
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o ReplicaTaskSessionEntityType) String() (string) { return o.FullString(false,0) }

/*
 * ReplicaTaskSessionEntityListType 
 * Not validated 
 */
func (o ReplicaTaskSessionEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.ReplicaTaskSession {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("ReplicaTaskSession[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o ReplicaTaskSessionEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * PlainCredentialsType 
 * Not validated 
 */
func (o PlainCredentialsType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("UserName                       : ");buffer.WriteString(fmt.Sprintf("%s",o.UserName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Password                       : ");buffer.WriteString(fmt.Sprintf("%s",o.Password))
   //Inhereting from InfoType
  return buffer.String()
}
func (o PlainCredentialsType) String() (string) { return o.FullString(false,0) }

/*
 * CredentialsInfoType 
 * Not validated 
 */
func (o CredentialsInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Id                             : ");buffer.WriteString(fmt.Sprintf("%s",o.Id))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Username                       : ");buffer.WriteString(fmt.Sprintf("%s",o.Username))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Description                    : ");buffer.WriteString(fmt.Sprintf("%s",o.Description))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Password                       : ");buffer.WriteString(fmt.Sprintf("%s",o.Password))
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o CredentialsInfoType) String() (string) { return o.FullString(false,0) }

/*
 * CredentialsInfoListType 
 * Not validated 
 */
func (o CredentialsInfoListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.CredentialsInfo {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("CredentialsInfo[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o CredentialsInfoListType) String() (string) { return o.FullString(false,0) }

/*
 * CredentialsInfoSpecType 
 * Not validated 
 */
func (o CredentialsInfoSpecType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Username                       : ");buffer.WriteString(fmt.Sprintf("%s",o.Username))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Description                    : ");buffer.WriteString(fmt.Sprintf("%s",o.Description))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Password                       : ");buffer.WriteString(fmt.Sprintf("%s",o.Password))
   //Inhereting from SpecType
   //Inhereting from ParamsType
  return buffer.String()
}
func (o CredentialsInfoSpecType) String() (string) { return o.FullString(false,0) }

/*
 * PasswordKeyInfoType 
 * Not validated 
 */
func (o PasswordKeyInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Id                             : ");buffer.WriteString(fmt.Sprintf("%s",o.Id))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Hint                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Hint))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("LastModificationDate           : ");buffer.WriteString(o.LastModificationDate.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o PasswordKeyInfoType) String() (string) { return o.FullString(false,0) }

/*
 * PasswordKeyInfoListType 
 * Not validated 
 */
func (o PasswordKeyInfoListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.PasswordKeyInfo {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("PasswordKeyInfo[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o PasswordKeyInfoListType) String() (string) { return o.FullString(false,0) }

/*
 * PasswordKeyInfoSpecType 
 * Not validated 
 */
func (o PasswordKeyInfoSpecType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Hint                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Hint))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Password                       : ");buffer.WriteString(fmt.Sprintf("%s",o.Password))
   //Inhereting from SpecType
   //Inhereting from ParamsType
  return buffer.String()
}
func (o PasswordKeyInfoSpecType) String() (string) { return o.FullString(false,0) }

/*
 * JobEntityListType 
 * Not validated 
 */
func (o JobEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.Job {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("Job[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o JobEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * FailoverPlanEntityListType 
 * Not validated 
 */
func (o FailoverPlanEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.FailoverPlan {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("FailoverPlan[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o FailoverPlanEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * JobEntityType 
 * Not validated 
 */
func (o JobEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("JobType                        : ");buffer.WriteString(fmt.Sprintf("%s",o.JobType))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Platform                       : ");buffer.WriteString(fmt.Sprintf("%s",o.Platform))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Description                    : ");buffer.WriteString(fmt.Sprintf("%s",o.Description))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ScheduleConfigured             : ");buffer.WriteString(fmt.Sprintf("%t",o.ScheduleConfigured))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ScheduleEnabled                : ");buffer.WriteString(fmt.Sprintf("%t",o.ScheduleEnabled))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("NextRun                        : ");buffer.WriteString(o.NextRun.String())
  if full && o.JobScheduleOptions != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("JobScheduleOptions             : ");buffer.WriteString(o.JobScheduleOptions.FullString(full,depth+1))
  }
  if full && o.JobInfo != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("JobInfo                        : ");buffer.WriteString(o.JobInfo.FullString(full,depth+1))
  }
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o JobEntityType) String() (string) { return o.FullString(false,0) }

/*
 * JobEntityTypeNestedJobInfo 
 * Not validated 
 */
func (o JobEntityTypeNestedJobInfo) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.BackupJobInfo != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("BackupJobInfo                  : ");buffer.WriteString(o.BackupJobInfo.FullString(full,depth+1))
  }
  if full && o.FileCopyJobInfo != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FileCopyJobInfo                : ");buffer.WriteString(o.FileCopyJobInfo.FullString(full,depth+1))
  }
  if full && o.ReplicaJobInfo != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ReplicaJobInfo                 : ");buffer.WriteString(o.ReplicaJobInfo.FullString(full,depth+1))
  }
  return buffer.String()
}
func (o JobEntityTypeNestedJobInfo) String() (string) { return o.FullString(false,0) }

/*
 * FailoverPlanEntityType 
 * Not validated 
 */
func (o FailoverPlanEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Description                    : ");buffer.WriteString(fmt.Sprintf("%s",o.Description))
  if full && o.FailoverPlanInfo != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FailoverPlanInfo               : ");buffer.WriteString(o.FailoverPlanInfo.FullString(full,depth+1))
  }
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o FailoverPlanEntityType) String() (string) { return o.FullString(false,0) }

/*
 * JobManagementSpecType 
 * Not validated 
 */
func (o JobManagementSpecType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.Credentials != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Credentials                    : ");buffer.WriteString(o.Credentials.FullString(full,depth+1))
  }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Force                          : ");buffer.WriteString(fmt.Sprintf("%t",o.Force))
   //Inhereting from SpecType
   //Inhereting from ParamsType
  return buffer.String()
}
func (o JobManagementSpecType) String() (string) { return o.FullString(false,0) }

/*
 * FailoverPlanManagementSpecType 
 * Not validated 
 */
func (o FailoverPlanManagementSpecType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("StartNow                       : ");buffer.WriteString(fmt.Sprintf("%t",o.StartNow))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("StartDate                      : ");buffer.WriteString(o.StartDate.String())
   //Inhereting from SpecType
   //Inhereting from ParamsType
  return buffer.String()
}
func (o FailoverPlanManagementSpecType) String() (string) { return o.FullString(false,0) }

/*
 * BackupJobInfoType 
 * Not validated 
 */
func (o BackupJobInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.Includes != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Includes                       : ");buffer.WriteString(o.Includes.FullString(full,depth+1))
  }
  if full && o.GuestProcessingOptions != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("GuestProcessingOptions         : ");buffer.WriteString(o.GuestProcessingOptions.FullString(full,depth+1))
  }
  if full && o.AdvancedStorageOptions != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("AdvancedStorageOptions         : ");buffer.WriteString(o.AdvancedStorageOptions.FullString(full,depth+1))
  }
   //Inhereting from InfoType
  return buffer.String()
}
func (o BackupJobInfoType) String() (string) { return o.FullString(false,0) }

/*
 * FailoverPlanInfoType 
 * Not validated 
 */
func (o FailoverPlanInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.Includes != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Includes                       : ");buffer.WriteString(o.Includes.FullString(full,depth+1))
  }
   //Inhereting from InfoType
  return buffer.String()
}
func (o FailoverPlanInfoType) String() (string) { return o.FullString(false,0) }

/*
 * ObjectInJobListType 
 * Not validated 
 */
func (o ObjectInJobListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.ObjectInJob {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("ObjectInJob[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o ObjectInJobListType) String() (string) { return o.FullString(false,0) }

/*
 * FailoveredVmListType 
 * Not validated 
 */
func (o FailoveredVmListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.FailoveredVm {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("FailoveredVm[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o FailoveredVmListType) String() (string) { return o.FullString(false,0) }

/*
 * ObjectInJobType 
 * Not validated 
 */
func (o ObjectInJobType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ObjectInJobId                  : ");buffer.WriteString(fmt.Sprintf("%s",o.ObjectInJobId))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("HierarchyObjRef                : ");buffer.WriteString(o.HierarchyObjRef.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("DisplayName                    : ");buffer.WriteString(fmt.Sprintf("%s",o.DisplayName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Order                          : ");buffer.WriteString(fmt.Sprintf("%d",o.Order))
  if full && o.GuestProcessingOptions != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("GuestProcessingOptions         : ");buffer.WriteString(o.GuestProcessingOptions.FullString(full,depth+1))
  }
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o ObjectInJobType) String() (string) { return o.FullString(false,0) }

/*
 * FailoveredVmType 
 * Not validated 
 */
func (o FailoveredVmType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FailoveredVmId                 : ");buffer.WriteString(fmt.Sprintf("%s",o.FailoveredVmId))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("HierarchyObjRef                : ");buffer.WriteString(o.HierarchyObjRef.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("DisplayName                    : ");buffer.WriteString(fmt.Sprintf("%s",o.DisplayName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Order                          : ");buffer.WriteString(fmt.Sprintf("%d",o.Order))
  if full && o.GuestProcessingOptions != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("GuestProcessingOptions         : ");buffer.WriteString(o.GuestProcessingOptions.FullString(full,depth+1))
  }
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o FailoveredVmType) String() (string) { return o.FullString(false,0) }

/*
 * CreateObjectInJobSpecType 
 * Not validated 
 */
func (o CreateObjectInJobSpecType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("HierarchyObjRef                : ");buffer.WriteString(o.HierarchyObjRef.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("HierarchyObjName               : ");buffer.WriteString(fmt.Sprintf("%s",o.HierarchyObjName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Order                          : ");buffer.WriteString(fmt.Sprintf("%d",o.Order))
  if full && o.GuestProcessingOptions != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("GuestProcessingOptions         : ");buffer.WriteString(o.GuestProcessingOptions.FullString(full,depth+1))
  }
   //Inhereting from SpecType
   //Inhereting from ParamsType
  return buffer.String()
}
func (o CreateObjectInJobSpecType) String() (string) { return o.FullString(false,0) }

/*
 * JobItemVssOptionsType 
 * Not validated 
 */
func (o JobItemVssOptionsType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Enabled                        : ");buffer.WriteString(fmt.Sprintf("%t",o.Enabled))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("IgnoreErrors                   : ");buffer.WriteString(fmt.Sprintf("%t",o.IgnoreErrors))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("GuestFSIndexingType            : ");buffer.WriteString(fmt.Sprintf("%s",o.GuestFSIndexingType))
  if full && o.Credentials != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Credentials                    : ");buffer.WriteString(o.Credentials.FullString(full,depth+1))
  }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("TransactionLogsTruncation      : ");buffer.WriteString(fmt.Sprintf("%s",o.TransactionLogsTruncation))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("IsFirstUsage                   : ");buffer.WriteString(fmt.Sprintf("%t",o.IsFirstUsage))
  if full && o.IncludedIndexingFolders != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("IncludedIndexingFolders        : ");buffer.WriteString(o.IncludedIndexingFolders.FullString(full,depth+1))
  }
  if full && o.ExcludedIndexingFolders != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ExcludedIndexingFolders        : ");buffer.WriteString(o.ExcludedIndexingFolders.FullString(full,depth+1))
  }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CredentialsId                  : ");buffer.WriteString(fmt.Sprintf("%s",o.CredentialsId))
   //Inhereting from InfoType
  return buffer.String()
}
func (o JobItemVssOptionsType) String() (string) { return o.FullString(false,0) }

/*
 * GuestProcessingOptionsType 
 * Not validated 
 */
func (o GuestProcessingOptionsType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("AppAwareProcessingMode         : ");buffer.WriteString(fmt.Sprintf("%s",o.AppAwareProcessingMode))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FileSystemIndexingMode         : ");buffer.WriteString(fmt.Sprintf("%s",o.FileSystemIndexingMode))
  if full && o.IncludedIndexingFolders != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("IncludedIndexingFolders        : ");buffer.WriteString(o.IncludedIndexingFolders.FullString(full,depth+1))
  }
  if full && o.ExcludedIndexingFolders != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ExcludedIndexingFolders        : ");buffer.WriteString(o.ExcludedIndexingFolders.FullString(full,depth+1))
  }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CredentialsId                  : ");buffer.WriteString(fmt.Sprintf("%s",o.CredentialsId))
  if full && o.VssSnapshotOptions != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VssSnapshotOptions             : ");buffer.WriteString(o.VssSnapshotOptions.FullString(full,depth+1))
  }
  if full && o.WindowsGuestFSIndexingOptions != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("WindowsGuestFSIndexingOptions  : ");buffer.WriteString(o.WindowsGuestFSIndexingOptions.FullString(full,depth+1))
  }
  if full && o.LinuxGuestFSIndexingOptions != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("LinuxGuestFSIndexingOptions    : ");buffer.WriteString(o.LinuxGuestFSIndexingOptions.FullString(full,depth+1))
  }
  if full && o.SqlBackupOptions != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("SqlBackupOptions               : ");buffer.WriteString(o.SqlBackupOptions.FullString(full,depth+1))
  }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("WindowsCredentialsId           : ");buffer.WriteString(fmt.Sprintf("%s",o.WindowsCredentialsId))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("LinuxCredentialsId             : ");buffer.WriteString(fmt.Sprintf("%s",o.LinuxCredentialsId))
  if full && o.FSFileExcludeOptions != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FSFileExcludeOptions           : ");buffer.WriteString(o.FSFileExcludeOptions.FullString(full,depth+1))
  }
  if full && o.OracleBackupOptions != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("OracleBackupOptions            : ");buffer.WriteString(o.OracleBackupOptions.FullString(full,depth+1))
  }
   //Inhereting from InfoType
  return buffer.String()
}
func (o GuestProcessingOptionsType) String() (string) { return o.FullString(false,0) }

/*
 * VssSnapshotOptionsType 
 * Not validated 
 */
func (o VssSnapshotOptionsType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VssSnapshotMode                : ");buffer.WriteString(fmt.Sprintf("%s",o.VssSnapshotMode))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("IsCopyOnly                     : ");buffer.WriteString(fmt.Sprintf("%t",o.IsCopyOnly))
   //Inhereting from InfoType
  return buffer.String()
}
func (o VssSnapshotOptionsType) String() (string) { return o.FullString(false,0) }

/*
 * WindowsGuestFSIndexingOptionsType 
 * Not validated 
 */
func (o WindowsGuestFSIndexingOptionsType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FileSystemIndexingMode         : ");buffer.WriteString(fmt.Sprintf("%s",o.FileSystemIndexingMode))
  if full && o.IncludedIndexingFolders != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("IncludedIndexingFolders        : ");buffer.WriteString(o.IncludedIndexingFolders.FullString(full,depth+1))
  }
  if full && o.ExcludedIndexingFolders != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ExcludedIndexingFolders        : ");buffer.WriteString(o.ExcludedIndexingFolders.FullString(full,depth+1))
  }
   //Inhereting from InfoType
  return buffer.String()
}
func (o WindowsGuestFSIndexingOptionsType) String() (string) { return o.FullString(false,0) }

/*
 * LinuxGuestFSIndexingOptionsType 
 * Not validated 
 */
func (o LinuxGuestFSIndexingOptionsType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FileSystemIndexingMode         : ");buffer.WriteString(fmt.Sprintf("%s",o.FileSystemIndexingMode))
  if full && o.IncludedIndexingFolders != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("IncludedIndexingFolders        : ");buffer.WriteString(o.IncludedIndexingFolders.FullString(full,depth+1))
  }
  if full && o.ExcludedIndexingFolders != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ExcludedIndexingFolders        : ");buffer.WriteString(o.ExcludedIndexingFolders.FullString(full,depth+1))
  }
   //Inhereting from InfoType
  return buffer.String()
}
func (o LinuxGuestFSIndexingOptionsType) String() (string) { return o.FullString(false,0) }

/*
 * SqlBackupOptionsType 
 * Not validated 
 */
func (o SqlBackupOptionsType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("TransactionLogsProcessing      : ");buffer.WriteString(fmt.Sprintf("%s",o.TransactionLogsProcessing))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("BackupLogsFrequencyMin         : ");buffer.WriteString(fmt.Sprintf("%d",o.BackupLogsFrequencyMin))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("UseDbBackupRetention           : ");buffer.WriteString(fmt.Sprintf("%t",o.UseDbBackupRetention))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("RetainDays                     : ");buffer.WriteString(fmt.Sprintf("%d",o.RetainDays))
   //Inhereting from InfoType
  return buffer.String()
}
func (o SqlBackupOptionsType) String() (string) { return o.FullString(false,0) }

/*
 * AdvancedStorageOptionsType 
 * Not validated 
 */
func (o AdvancedStorageOptionsType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("PasswordKeyId                  : ");buffer.WriteString(fmt.Sprintf("%s",o.PasswordKeyId))
   //Inhereting from InfoType
  return buffer.String()
}
func (o AdvancedStorageOptionsType) String() (string) { return o.FullString(false,0) }

/*
 * FSFileExcludeOptionsType 
 * Not validated 
 */
func (o FSFileExcludeOptionsType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("BackupScope                    : ");buffer.WriteString(fmt.Sprintf("%d",o.BackupScope))
  if full && o.IncludeList != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("IncludeList                    : ");buffer.WriteString(o.IncludeList.FullString(full,depth+1))
  }
  if full && o.ExcludeList != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ExcludeList                    : ");buffer.WriteString(o.ExcludeList.FullString(full,depth+1))
  }
   //Inhereting from InfoType
  return buffer.String()
}
func (o FSFileExcludeOptionsType) String() (string) { return o.FullString(false,0) }

/*
 * OracleBackupOptionsType 
 * Not validated 
 */
func (o OracleBackupOptionsType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("BackupLogsEnabled              : ");buffer.WriteString(fmt.Sprintf("%t",o.BackupLogsEnabled))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("BackupLogsFrequencyMin         : ");buffer.WriteString(fmt.Sprintf("%d",o.BackupLogsFrequencyMin))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("UseDbBackupRetention           : ");buffer.WriteString(fmt.Sprintf("%t",o.UseDbBackupRetention))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("RetainDays                     : ");buffer.WriteString(fmt.Sprintf("%d",o.RetainDays))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ArchivedLogsTruncation         : ");buffer.WriteString(fmt.Sprintf("%s",o.ArchivedLogsTruncation))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ArchivedLogsMaxAgeHours        : ");buffer.WriteString(fmt.Sprintf("%d",o.ArchivedLogsMaxAgeHours))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ArchivedLogsMaxSizeMb          : ");buffer.WriteString(fmt.Sprintf("%d",o.ArchivedLogsMaxSizeMb))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("SysdbaCredsId                  : ");buffer.WriteString(fmt.Sprintf("%s",o.SysdbaCredsId))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ProxyAutoSelect                : ");buffer.WriteString(fmt.Sprintf("%t",o.ProxyAutoSelect))
   //Inhereting from InfoType
  return buffer.String()
}
func (o OracleBackupOptionsType) String() (string) { return o.FullString(false,0) }

/*
 * StringListType 
 * Not validated 
 */
func (o StringListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.Path {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("Path[%d]",i))
   buffer.Write(identbuffer.Bytes())
    buffer.WriteString("Path                           : ");buffer.WriteString(fmt.Sprintf("%s",c))
   }
  }
  return buffer.String()
}
func (o StringListType) String() (string) { return o.FullString(false,0) }

/*
 * JobScheduleOptionsInfoType 
 * Not validated 
 */
func (o JobScheduleOptionsInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.RetryOptions != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("RetryOptions                   : ");buffer.WriteString(o.RetryOptions.FullString(full,depth+1))
  }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("WaitForBackupCompletion        : ");buffer.WriteString(fmt.Sprintf("%t",o.WaitForBackupCompletion))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("BackupCompetitionWaitingPeriodMin : ");buffer.WriteString(fmt.Sprintf("%d",o.BackupCompetitionWaitingPeriodMin))
  if full && o.OptionsDaily != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("OptionsDaily                   : ");buffer.WriteString(o.OptionsDaily.FullString(full,depth+1))
  }
  if full && o.OptionsMonthly != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("OptionsMonthly                 : ");buffer.WriteString(o.OptionsMonthly.FullString(full,depth+1))
  }
  if full && o.OptionsPeriodically != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("OptionsPeriodically            : ");buffer.WriteString(o.OptionsPeriodically.FullString(full,depth+1))
  }
  if full && o.OptionsContinuous != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("OptionsContinuous              : ");buffer.WriteString(o.OptionsContinuous.FullString(full,depth+1))
  }
  if full && o.OptionsBackupWindow != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("OptionsBackupWindow            : ");buffer.WriteString(o.OptionsBackupWindow.FullString(full,depth+1))
  }
  if full && o.OptionsDaisyChaining != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("OptionsDaisyChaining           : ");buffer.WriteString(o.OptionsDaisyChaining.FullString(full,depth+1))
  }
   //Inhereting from InfoType
  return buffer.String()
}
func (o JobScheduleOptionsInfoType) String() (string) { return o.FullString(false,0) }

/*
 * JobScheduleRetryOptionsType 
 * Not validated 
 */
func (o JobScheduleRetryOptionsType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("RetryTimes                     : ");buffer.WriteString(fmt.Sprintf("%d",o.RetryTimes))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("RetryTimeout                   : ");buffer.WriteString(fmt.Sprintf("%d",o.RetryTimeout))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("RetrySpecified                 : ");buffer.WriteString(fmt.Sprintf("%t",o.RetrySpecified))
   //Inhereting from InfoType
  return buffer.String()
}
func (o JobScheduleRetryOptionsType) String() (string) { return o.FullString(false,0) }

/*
 * JobScheduleDailyOptionsType 
 * Not validated 
 */
func (o JobScheduleDailyOptionsType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Kind                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Kind))
  if full {
   for i,c := range o.Days {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("Days[%d]",i))
   buffer.Write(identbuffer.Bytes())
    buffer.WriteString("Days                           : ");buffer.WriteString(c.String())
   }
  }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Time                           : ");buffer.WriteString(o.Time.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("TimeOffsetUtc                  : ");buffer.WriteString(fmt.Sprintf("%d",o.TimeOffsetUtc))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Enabled                        : ");buffer.WriteString(fmt.Sprintf("%t",o.Enabled))
   //Inhereting from InfoType
  return buffer.String()
}
func (o JobScheduleDailyOptionsType) String() (string) { return o.FullString(false,0) }

/*
 * JobScheduleMonthlyOptionsType 
 * Not validated 
 */
func (o JobScheduleMonthlyOptionsType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Time                           : ");buffer.WriteString(o.Time.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("TimeOffsetUtc                  : ");buffer.WriteString(fmt.Sprintf("%d",o.TimeOffsetUtc))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("DayNumberInMonth               : ");buffer.WriteString(fmt.Sprintf("%s",o.DayNumberInMonth))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("DayOfWeek                      : ");buffer.WriteString(fmt.Sprintf("%s",o.DayOfWeek))
  if full {
   for i,c := range o.Months {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("Months[%d]",i))
   buffer.Write(identbuffer.Bytes())
    buffer.WriteString("Months                         : ");buffer.WriteString(c.String())
   }
  }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("DayOfMonth                     : ");buffer.WriteString(fmt.Sprintf("%d",o.DayOfMonth))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Enabled                        : ");buffer.WriteString(fmt.Sprintf("%t",o.Enabled))
   //Inhereting from InfoType
  return buffer.String()
}
func (o JobScheduleMonthlyOptionsType) String() (string) { return o.FullString(false,0) }

/*
 * JobSchedulePeriodicallyOptionsType 
 * Not validated 
 */
func (o JobSchedulePeriodicallyOptionsType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Kind                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Kind))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FullPeriod                     : ");buffer.WriteString(fmt.Sprintf("%d",o.FullPeriod))
  if full && o.Schedule != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Schedule                       : ");buffer.WriteString(o.Schedule.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Enabled                        : ");buffer.WriteString(fmt.Sprintf("%t",o.Enabled))
   //Inhereting from InfoType
  return buffer.String()
}
func (o JobSchedulePeriodicallyOptionsType) String() (string) { return o.FullString(false,0) }

/*
 * TimePeriodsType 
 * Not validated 
 */
func (o TimePeriodsType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.Day {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("Day[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from InfoType
  return buffer.String()
}
func (o TimePeriodsType) String() (string) { return o.FullString(false,0) }

/*
 * TimePeriodsTypeNestedDay 
 * Not validated 
 */
func (o TimePeriodsTypeNestedDay) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  return buffer.String()
}
func (o TimePeriodsTypeNestedDay) String() (string) { return o.FullString(false,0) }

/*
 * JobScheduleContinuousOptionsType 
 * Not validated 
 */
func (o JobScheduleContinuousOptionsType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Enabled                        : ");buffer.WriteString(fmt.Sprintf("%t",o.Enabled))
   //Inhereting from InfoType
  return buffer.String()
}
func (o JobScheduleContinuousOptionsType) String() (string) { return o.FullString(false,0) }

/*
 * JobScheduleBackupWindowOptionsType 
 * Not validated 
 */
func (o JobScheduleBackupWindowOptionsType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.TimePeriods != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("TimePeriods                    : ");buffer.WriteString(o.TimePeriods.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Enabled                        : ");buffer.WriteString(fmt.Sprintf("%t",o.Enabled))
   //Inhereting from InfoType
  return buffer.String()
}
func (o JobScheduleBackupWindowOptionsType) String() (string) { return o.FullString(false,0) }

/*
 * JobScheduleDaisyChainingOptionsType 
 * Not validated 
 */
func (o JobScheduleDaisyChainingOptionsType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("PreviousJobUid                 : ");buffer.WriteString(o.PreviousJobUid.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Enabled                        : ");buffer.WriteString(fmt.Sprintf("%t",o.Enabled))
   //Inhereting from InfoType
  return buffer.String()
}
func (o JobScheduleDaisyChainingOptionsType) String() (string) { return o.FullString(false,0) }

/*
 * FileCopyJobInfoType 
 * Not validated 
 */
func (o FileCopyJobInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   //Inhereting from InfoType
  return buffer.String()
}
func (o FileCopyJobInfoType) String() (string) { return o.FullString(false,0) }

/*
 * ReplicaJobInfoType 
 * Not validated 
 */
func (o ReplicaJobInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.Includes != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Includes                       : ");buffer.WriteString(o.Includes.FullString(full,depth+1))
  }
  if full && o.GuestProcessingOptions != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("GuestProcessingOptions         : ");buffer.WriteString(o.GuestProcessingOptions.FullString(full,depth+1))
  }
   //Inhereting from InfoType
  return buffer.String()
}
func (o ReplicaJobInfoType) String() (string) { return o.FullString(false,0) }

/*
 * TaskType 
 * Not validated 
 */
func (o TaskType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("TaskId                         : ");buffer.WriteString(fmt.Sprintf("%s",o.TaskId))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("State                          : ");buffer.WriteString(fmt.Sprintf("%s",o.State))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Operation                      : ");buffer.WriteString(fmt.Sprintf("%s",o.Operation))
  if full && o.Result != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Result                         : ");buffer.WriteString(o.Result.FullString(full,depth+1))
  }
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o TaskType) String() (string) { return o.FullString(false,0) }

/*
 * TaskListType 
 * Not validated 
 */
func (o TaskListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.Task {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("Task[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o TaskListType) String() (string) { return o.FullString(false,0) }

/*
 * QueryResultType 
 * Not validated 
 */
func (o QueryResultType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.Refs != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Refs                           : ");buffer.WriteString(o.Refs.FullString(full,depth+1))
  }
  if full && o.Entities != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Entities                       : ");buffer.WriteString(o.Entities.FullString(full,depth+1))
  }
  if full && o.Resources != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Resources                      : ");buffer.WriteString(o.Resources.FullString(full,depth+1))
  }
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  if full && o.PagingInfo != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("PagingInfo                     : ");buffer.WriteString(o.PagingInfo.FullString(full,depth+1))
  }
  return buffer.String()
}
func (o QueryResultType) String() (string) { return o.FullString(false,0) }

/*
 * PagingInfoType 
 * Not validated 
 */
func (o PagingInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("PageNum                        : ");buffer.WriteString(fmt.Sprintf("%d",o.PageNum))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("PageSize                       : ");buffer.WriteString(fmt.Sprintf("%d",o.PageSize))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("PagesCount                     : ");buffer.WriteString(fmt.Sprintf("%d",o.PagesCount))
   //Inhereting from InfoType
  return buffer.String()
}
func (o PagingInfoType) String() (string) { return o.FullString(false,0) }

/*
 * BackupEntityType 
 * Not validated 
 */
func (o BackupEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.Ref != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Ref                            : ");buffer.WriteString(o.Ref.FullString(full,depth+1))
  }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Platform                       : ");buffer.WriteString(fmt.Sprintf("%s",o.Platform))
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o BackupEntityType) String() (string) { return o.FullString(false,0) }

/*
 * BackupEntityListType 
 * Not validated 
 */
func (o BackupEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.Backup {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("Backup[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o BackupEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * ReplicaEntityType 
 * Not validated 
 */
func (o ReplicaEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.Ref != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Ref                            : ");buffer.WriteString(o.Ref.FullString(full,depth+1))
  }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Platform                       : ");buffer.WriteString(fmt.Sprintf("%s",o.Platform))
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o ReplicaEntityType) String() (string) { return o.FullString(false,0) }

/*
 * ReplicaEntityListType 
 * Not validated 
 */
func (o ReplicaEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.Replica {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("Replica[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o ReplicaEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * RestorePointEntityType 
 * Not validated 
 */
func (o RestorePointEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.Ref != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Ref                            : ");buffer.WriteString(o.Ref.FullString(full,depth+1))
  }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("BackupDateUTC                  : ");buffer.WriteString(o.BackupDateUTC.String())
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o RestorePointEntityType) String() (string) { return o.FullString(false,0) }

/*
 * RestorePointEntityListType 
 * Not validated 
 */
func (o RestorePointEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.RestorePoint {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("RestorePoint[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o RestorePointEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * VmRestorePointEntityType 
 * Not validated 
 */
func (o VmRestorePointEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CreationTimeUTC                : ");buffer.WriteString(o.CreationTimeUTC.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VmName                         : ");buffer.WriteString(fmt.Sprintf("%s",o.VmName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Algorithm                      : ");buffer.WriteString(fmt.Sprintf("%s",o.Algorithm))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("PointType                      : ");buffer.WriteString(fmt.Sprintf("%s",o.PointType))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("HierarchyObjRef                : ");buffer.WriteString(o.HierarchyObjRef.String())
  if full && o.SqlInfo != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("SqlInfo                        : ");buffer.WriteString(o.SqlInfo.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("VmDisplayName                  : ");buffer.WriteString(fmt.Sprintf("%s",o.VmDisplayName))
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o VmRestorePointEntityType) String() (string) { return o.FullString(false,0) }

/*
 * VAppRestorePointEntityType 
 * Not validated 
 */
func (o VAppRestorePointEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CreationTimeUTC                : ");buffer.WriteString(o.CreationTimeUTC.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VAppName                       : ");buffer.WriteString(fmt.Sprintf("%s",o.VAppName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Algorithm                      : ");buffer.WriteString(fmt.Sprintf("%s",o.Algorithm))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("PointType                      : ");buffer.WriteString(fmt.Sprintf("%s",o.PointType))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("HierarchyObjRef                : ");buffer.WriteString(o.HierarchyObjRef.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("VAppDisplayName                : ");buffer.WriteString(fmt.Sprintf("%s",o.VAppDisplayName))
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o VAppRestorePointEntityType) String() (string) { return o.FullString(false,0) }

/*
 * VmRestorePointEntityListType 
 * Not validated 
 */
func (o VmRestorePointEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.VmRestorePoint {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("VmRestorePoint[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o VmRestorePointEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * VAppRestorePointEntityListType 
 * Not validated 
 */
func (o VAppRestorePointEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.VAppRestorePoint {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("VAppRestorePoint[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o VAppRestorePointEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * VmReplicaPointEntityType 
 * Not validated 
 */
func (o VmReplicaPointEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CreationTimeUTC                : ");buffer.WriteString(o.CreationTimeUTC.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VmName                         : ");buffer.WriteString(fmt.Sprintf("%s",o.VmName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Algorithm                      : ");buffer.WriteString(fmt.Sprintf("%s",o.Algorithm))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("PointType                      : ");buffer.WriteString(fmt.Sprintf("%s",o.PointType))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("VmDisplayName                  : ");buffer.WriteString(fmt.Sprintf("%s",o.VmDisplayName))
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o VmReplicaPointEntityType) String() (string) { return o.FullString(false,0) }

/*
 * VmReplicaPointEntityListType 
 * Not validated 
 */
func (o VmReplicaPointEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.VmRestorePoint {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("VmRestorePoint[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o VmReplicaPointEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * VmRestorePointMountType 
 * Not validated 
 */
func (o VmRestorePointMountType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.FSRoots != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FSRoots                        : ");buffer.WriteString(o.FSRoots.FullString(full,depth+1))
  }
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o VmRestorePointMountType) String() (string) { return o.FullString(false,0) }

/*
 * VmRestorePointMountListType 
 * Not validated 
 */
func (o VmRestorePointMountListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.VmRestorePointMount {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("VmRestorePointMount[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o VmRestorePointMountListType) String() (string) { return o.FullString(false,0) }

/*
 * VmReplicaPointMountType 
 * Not validated 
 */
func (o VmReplicaPointMountType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.FSRoots != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FSRoots                        : ");buffer.WriteString(o.FSRoots.FullString(full,depth+1))
  }
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o VmReplicaPointMountType) String() (string) { return o.FullString(false,0) }

/*
 * VmReplicaPointMountListType 
 * Not validated 
 */
func (o VmReplicaPointMountListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.VmReplicaPointMount {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("VmReplicaPointMount[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o VmReplicaPointMountListType) String() (string) { return o.FullString(false,0) }

/*
 * CatalogVmEntityType 
 * Not validated 
 */
func (o CatalogVmEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("VmDisplayName                  : ");buffer.WriteString(fmt.Sprintf("%s",o.VmDisplayName))
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o CatalogVmEntityType) String() (string) { return o.FullString(false,0) }

/*
 * CatalogVmEntityListType 
 * Not validated 
 */
func (o CatalogVmEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.CatalogVm {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("CatalogVm[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o CatalogVmEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * CatalogVmRestorePointEntityType 
 * Not validated 
 */
func (o CatalogVmRestorePointEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("BackupDateUTC                  : ");buffer.WriteString(o.BackupDateUTC.String())
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o CatalogVmRestorePointEntityType) String() (string) { return o.FullString(false,0) }

/*
 * CatalogVmRestorePointEntityListType 
 * Not validated 
 */
func (o CatalogVmRestorePointEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.CatalogVmRestorePoint {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("CatalogVmRestorePoint[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o CatalogVmRestorePointEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * FileSystemEntryType 
 * Not validated 
 */
func (o FileSystemEntryType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.FileEntry != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FileEntry                      : ");buffer.WriteString(o.FileEntry.FullString(full,depth+1))
  }
  if full && o.DirectoryEntry != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("DirectoryEntry                 : ");buffer.WriteString(o.DirectoryEntry.FullString(full,depth+1))
  }
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o FileSystemEntryType) String() (string) { return o.FullString(false,0) }

/*
 * DirectoryEntryType 
 * Not validated 
 */
func (o DirectoryEntryType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Path                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Path))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o DirectoryEntryType) String() (string) { return o.FullString(false,0) }

/*
 * FileEntryType 
 * Not validated 
 */
func (o FileEntryType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Path                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Path))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Size                           : ");buffer.WriteString(fmt.Sprintf("%d",o.Size))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Owner                          : ");buffer.WriteString(fmt.Sprintf("%s",o.Owner))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ModifiedDateUTC                : ");buffer.WriteString(o.ModifiedDateUTC.String())
  if full && o.Actions != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Actions                        : ");buffer.WriteString(o.Actions.FullString(full,depth+1))
  }
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o FileEntryType) String() (string) { return o.FullString(false,0) }

/*
 * FileEntryListType 
 * Not validated 
 */
func (o FileEntryListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.FileEntry {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("FileEntry[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o FileEntryListType) String() (string) { return o.FullString(false,0) }

/*
 * HierarchyRootEntityType 
 * Not validated 
 */
func (o HierarchyRootEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("HierarchyRootId                : ");buffer.WriteString(fmt.Sprintf("%s",o.HierarchyRootId))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("UniqueId                       : ");buffer.WriteString(fmt.Sprintf("%s",o.UniqueId))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("HostType                       : ");buffer.WriteString(fmt.Sprintf("%s",o.HostType))
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o HierarchyRootEntityType) String() (string) { return o.FullString(false,0) }

/*
 * RepositoryEntityType 
 * Not validated 
 */
func (o RepositoryEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Capacity                       : ");buffer.WriteString(fmt.Sprintf("%d",o.Capacity))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FreeSpace                      : ");buffer.WriteString(fmt.Sprintf("%d",o.FreeSpace))
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o RepositoryEntityType) String() (string) { return o.FullString(false,0) }

/*
 * HierarchyRootEntityListType 
 * Not validated 
 */
func (o HierarchyRootEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.HierarchyRoot {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("HierarchyRoot[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o HierarchyRootEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * RepositoryEntityListType 
 * Not validated 
 */
func (o RepositoryEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.Repository {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("Repository[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o RepositoryEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * DirectoryEntryListType 
 * Not validated 
 */
func (o DirectoryEntryListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.DirectoryEntry {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("DirectoryEntry[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o DirectoryEntryListType) String() (string) { return o.FullString(false,0) }

/*
 * FileSystemEntriesType 
 * Not validated 
 */
func (o FileSystemEntriesType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.Files != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Files                          : ");buffer.WriteString(o.Files.FullString(full,depth+1))
  }
  if full && o.Directories != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Directories                    : ");buffer.WriteString(o.Directories.FullString(full,depth+1))
  }
  if full && o.PagingInfo != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("PagingInfo                     : ");buffer.WriteString(o.PagingInfo.FullString(full,depth+1))
  }
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o FileSystemEntriesType) String() (string) { return o.FullString(false,0) }

/*
 * HierarchyItemListType 
 * Not validated 
 */
func (o HierarchyItemListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.HierarchyItem {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("HierarchyItem[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o HierarchyItemListType) String() (string) { return o.FullString(false,0) }

/*
 * BackupJobSessionEntityListType 
 * Not validated 
 */
func (o BackupJobSessionEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.BackupJobSession {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("BackupJobSession[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o BackupJobSessionEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * ReplicaJobSessionEntityListType 
 * Not validated 
 */
func (o ReplicaJobSessionEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.ReplicaJobSession {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("ReplicaJobSession[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o ReplicaJobSessionEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * RestoreSessionEntityListType 
 * Not validated 
 */
func (o RestoreSessionEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.RestoreSession {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("RestoreSession[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o RestoreSessionEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * FileRestoreSpecType 
 * Not validated 
 */
func (o FileRestoreSpecType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.ToOriginalLocation != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ToOriginalLocation             : ");buffer.WriteString(o.ToOriginalLocation.FullString(full,depth+1))
  }
  if full && o.ForDirectDownload != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ForDirectDownload              : ");buffer.WriteString(o.ForDirectDownload.FullString(full,depth+1))
  }
   //Inhereting from SpecType
   //Inhereting from ParamsType
  return buffer.String()
}
func (o FileRestoreSpecType) String() (string) { return o.FullString(false,0) }

/*
 * FileRestoreSpecTypeNestedToOriginalLocation 
 * Not validated 
 */
func (o FileRestoreSpecTypeNestedToOriginalLocation) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  return buffer.String()
}
func (o FileRestoreSpecTypeNestedToOriginalLocation) String() (string) { return o.FullString(false,0) }

/*
 * FileRestoreSpecTypeNestedForDirectDownload 
 * Not validated 
 */
func (o FileRestoreSpecTypeNestedForDirectDownload) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  return buffer.String()
}
func (o FileRestoreSpecTypeNestedForDirectDownload) String() (string) { return o.FullString(false,0) }

/*
 * RestoreSpecType 
 * Not validated 
 */
func (o RestoreSpecType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.VmRestoreSpec != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VmRestoreSpec                  : ");buffer.WriteString(o.VmRestoreSpec.FullString(full,depth+1))
  }
  if full && o.vCloudVmRestoreSpec != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("vCloudVmRestoreSpec            : ");buffer.WriteString(o.vCloudVmRestoreSpec.FullString(full,depth+1))
  }
  if full && o.vCloudVAppRestoreSpec != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("vCloudVAppRestoreSpec          : ");buffer.WriteString(o.vCloudVAppRestoreSpec.FullString(full,depth+1))
  }
  if full && o.SqlItemRestoreSpec != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("SqlItemRestoreSpec             : ");buffer.WriteString(o.SqlItemRestoreSpec.FullString(full,depth+1))
  }
   //Inhereting from SpecType
   //Inhereting from ParamsType
  return buffer.String()
}
func (o RestoreSpecType) String() (string) { return o.FullString(false,0) }

/*
 * VmRestoreSpecInfoType 
 * Not validated 
 */
func (o VmRestoreSpecInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("PowerOnAfterRestore            : ");buffer.WriteString(fmt.Sprintf("%t",o.PowerOnAfterRestore))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("QuickRollback                  : ");buffer.WriteString(fmt.Sprintf("%t",o.QuickRollback))
   //Inhereting from InfoType
  return buffer.String()
}
func (o VmRestoreSpecInfoType) String() (string) { return o.FullString(false,0) }

/*
 * vCloudVmRestoreSpecInfoType 
 * Not validated 
 */
func (o vCloudVmRestoreSpecInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("PowerOnAfterRestore            : ");buffer.WriteString(fmt.Sprintf("%t",o.PowerOnAfterRestore))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("HierarchyRootUid               : ");buffer.WriteString(o.HierarchyRootUid.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("vAppRef                        : ");buffer.WriteString(o.vAppRef.String())
  if full && o.VmRestoreParameters != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VmRestoreParameters            : ");buffer.WriteString(o.VmRestoreParameters.FullString(full,depth+1))
  }
   //Inhereting from InfoType
  return buffer.String()
}
func (o vCloudVmRestoreSpecInfoType) String() (string) { return o.FullString(false,0) }

/*
 * vCloudVAppRestoreSpecInfoType 
 * Not validated 
 */
func (o vCloudVAppRestoreSpecInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("PowerOnAfterRestore            : ");buffer.WriteString(fmt.Sprintf("%t",o.PowerOnAfterRestore))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("HierarchyRootUid               : ");buffer.WriteString(o.HierarchyRootUid.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("OrgVdcRef                      : ");buffer.WriteString(o.OrgVdcRef.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("vAppNewName                    : ");buffer.WriteString(fmt.Sprintf("%s",o.vAppNewName))
  if full && o.VmsRestoreParameters != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VmsRestoreParameters           : ");buffer.WriteString(o.VmsRestoreParameters.FullString(full,depth+1))
  }
   //Inhereting from InfoType
  return buffer.String()
}
func (o vCloudVAppRestoreSpecInfoType) String() (string) { return o.FullString(false,0) }

/*
 * vCloudVAppRestoreSpecInfoTypeNestedVmsRestoreParameters 
 * Not validated 
 */
func (o vCloudVAppRestoreSpecInfoTypeNestedVmsRestoreParameters) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.Vm {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("Vm[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
  return buffer.String()
}
func (o vCloudVAppRestoreSpecInfoTypeNestedVmsRestoreParameters) String() (string) { return o.FullString(false,0) }

/*
 * vCloudVmRestoreParametersInfoType 
 * Not validated 
 */
func (o vCloudVmRestoreParametersInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VmRestorePointUid              : ");buffer.WriteString(o.VmRestorePointUid.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VmNewName                      : ");buffer.WriteString(fmt.Sprintf("%s",o.VmNewName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("DatastoreRef                   : ");buffer.WriteString(o.DatastoreRef.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("OrgVdcStorageProfileRef        : ");buffer.WriteString(o.OrgVdcStorageProfileRef.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("LinkedCloneVmTemplateRef       : ");buffer.WriteString(o.LinkedCloneVmTemplateRef.String())
   //Inhereting from InfoType
  return buffer.String()
}
func (o vCloudVmRestoreParametersInfoType) String() (string) { return o.FullString(false,0) }

/*
 * BackupJobCloneInfoType 
 * Not validated 
 */
func (o BackupJobCloneInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("JobName                        : ");buffer.WriteString(fmt.Sprintf("%s",o.JobName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FolderName                     : ");buffer.WriteString(fmt.Sprintf("%s",o.FolderName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("RepositoryUid                  : ");buffer.WriteString(o.RepositoryUid.String())
   //Inhereting from InfoType
  return buffer.String()
}
func (o BackupJobCloneInfoType) String() (string) { return o.FullString(false,0) }

/*
 * ReplicaJobCloneInfoType 
 * Not validated 
 */
func (o ReplicaJobCloneInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("JobName                        : ");buffer.WriteString(fmt.Sprintf("%s",o.JobName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VmSuffix                       : ");buffer.WriteString(fmt.Sprintf("%s",o.VmSuffix))
   //Inhereting from InfoType
  return buffer.String()
}
func (o ReplicaJobCloneInfoType) String() (string) { return o.FullString(false,0) }

/*
 * JobCloneSpecType 
 * Not validated 
 */
func (o JobCloneSpecType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.BackupJobCloneInfo != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("BackupJobCloneInfo             : ");buffer.WriteString(o.BackupJobCloneInfo.FullString(full,depth+1))
  }
  if full && o.ReplicaJobCloneInfo != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ReplicaJobCloneInfo            : ");buffer.WriteString(o.ReplicaJobCloneInfo.FullString(full,depth+1))
  }
   //Inhereting from SpecType
   //Inhereting from ParamsType
  return buffer.String()
}
func (o JobCloneSpecType) String() (string) { return o.FullString(false,0) }

/*
 * TaskResultInfoType 
 * Not validated 
 */
func (o TaskResultInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Message                        : ");buffer.WriteString(fmt.Sprintf("%s",o.Message))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Success                        : ");buffer.WriteString(fmt.Sprintf("%t",o.Success))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("CredentialsNeeded              : ");buffer.WriteString(fmt.Sprintf("%t",o.CredentialsNeeded))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("ConfirmationNeeded             : ");buffer.WriteString(fmt.Sprintf("%t",o.ConfirmationNeeded))
   //Inhereting from InfoType
  return buffer.String()
}
func (o TaskResultInfoType) String() (string) { return o.FullString(false,0) }

/*
 * ErrorType 
 * Not validated 
 */
func (o ErrorType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FirstChanceExceptionMessage    : ");buffer.WriteString(fmt.Sprintf("%s",o.FirstChanceExceptionMessage))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("StackTrace                     : ");buffer.WriteString(fmt.Sprintf("%s",o.StackTrace))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Message                        : ");buffer.WriteString(fmt.Sprintf("%s",o.Message))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("StatusCode                     : ");buffer.WriteString(fmt.Sprintf("%d",o.StatusCode))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Status                         : ");buffer.WriteString(fmt.Sprintf("%s",o.Status))
   //Inhereting from InfoType
  return buffer.String()
}
func (o ErrorType) String() (string) { return o.FullString(false,0) }

/*
 * HierarchyItemType 
 * Not validated 
 */
func (o HierarchyItemType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ObjectRef                      : ");buffer.WriteString(o.ObjectRef.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ObjectType                     : ");buffer.WriteString(fmt.Sprintf("%s",o.ObjectType))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ObjectName                     : ");buffer.WriteString(fmt.Sprintf("%s",o.ObjectName))
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o HierarchyItemType) String() (string) { return o.FullString(false,0) }

/*
 * QuerySvcType 
 * Not validated 
 */
func (o QuerySvcType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o QuerySvcType) String() (string) { return o.FullString(false,0) }

/*
 * LookupSvcType 
 * Not validated 
 */
func (o LookupSvcType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o LookupSvcType) String() (string) { return o.FullString(false,0) }

/*
 * ReportingSvcType 
 * Not validated 
 */
func (o ReportingSvcType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o ReportingSvcType) String() (string) { return o.FullString(false,0) }

/*
 * EnterpriseRoleEntityType 
 * Not validated 
 */
func (o EnterpriseRoleEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o EnterpriseRoleEntityType) String() (string) { return o.FullString(false,0) }

/*
 * EnterpriseRoleEntityListType 
 * Not validated 
 */
func (o EnterpriseRoleEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.EnterpriseRoleEntity {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("EnterpriseRoleEntity[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o EnterpriseRoleEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * EnterpriseAccountEntityType 
 * Not validated 
 */
func (o EnterpriseAccountEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("AccountType                    : ");buffer.WriteString(o.AccountType.String())
  if full && o.Roles != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Roles                          : ");buffer.WriteString(o.Roles.FullString(full,depth+1))
  }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("AllowRestoreAllVms             : ");buffer.WriteString(fmt.Sprintf("%t",o.AllowRestoreAllVms))
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o EnterpriseAccountEntityType) String() (string) { return o.FullString(false,0) }

/*
 * EnterpriseAccountEntityListType 
 * Not validated 
 */
func (o EnterpriseAccountEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.EnterpriseAccountEntity {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("EnterpriseAccountEntity[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o EnterpriseAccountEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * EnterpriseAccountHierarchyScopeType 
 * Not validated 
 */
func (o EnterpriseAccountHierarchyScopeType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("HierarchyRootName              : ");buffer.WriteString(fmt.Sprintf("%s",o.HierarchyRootName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Platform                       : ");buffer.WriteString(fmt.Sprintf("%s",o.Platform))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("HierarchyObjectType            : ");buffer.WriteString(fmt.Sprintf("%s",o.HierarchyObjectType))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("State                          : ");buffer.WriteString(fmt.Sprintf("%s",o.State))
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o EnterpriseAccountHierarchyScopeType) String() (string) { return o.FullString(false,0) }

/*
 * EnterpriseAccountHierarchyScopeListType 
 * Not validated 
 */
func (o EnterpriseAccountHierarchyScopeListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.EnterpriseAccountHierarchyScope {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("EnterpriseAccountHierarchyScope[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o EnterpriseAccountHierarchyScopeListType) String() (string) { return o.FullString(false,0) }

/*
 * EnterpriseAccountInRoleType 
 * Not validated 
 */
func (o EnterpriseAccountInRoleType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("RoleName                       : ");buffer.WriteString(fmt.Sprintf("%s",o.RoleName))
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o EnterpriseAccountInRoleType) String() (string) { return o.FullString(false,0) }

/*
 * EnterpriseAccountInRoleListType 
 * Not validated 
 */
func (o EnterpriseAccountInRoleListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.EnterpriseAccountInRole {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("EnterpriseAccountInRole[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o EnterpriseAccountInRoleListType) String() (string) { return o.FullString(false,0) }

/*
 * HierarchyScopeCreateSpecType 
 * Not validated 
 */
func (o HierarchyScopeCreateSpecType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.HierarchyScopeItem {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("HierarchyScopeItem[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o HierarchyScopeCreateSpecType) String() (string) { return o.FullString(false,0) }

/*
 * HierarchyScopeCreateSpecItemType 
 * Not validated 
 */
func (o HierarchyScopeCreateSpecItemType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("HierarchyObjRef                : ");buffer.WriteString(o.HierarchyObjRef.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ObjectName                     : ");buffer.WriteString(fmt.Sprintf("%s",o.ObjectName))
   //Inhereting from SpecType
   //Inhereting from ParamsType
  return buffer.String()
}
func (o HierarchyScopeCreateSpecItemType) String() (string) { return o.FullString(false,0) }

/*
 * EnterpriseAccountInRoleCreateSpecType 
 * Not validated 
 */
func (o EnterpriseAccountInRoleCreateSpecType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("EnterpriseRoleUid              : ");buffer.WriteString(o.EnterpriseRoleUid.String())
   //Inhereting from SpecType
   //Inhereting from ParamsType
  return buffer.String()
}
func (o EnterpriseAccountInRoleCreateSpecType) String() (string) { return o.FullString(false,0) }

/*
 * EnterpriseAccountInRoleCreateSpecListType 
 * Not validated 
 */
func (o EnterpriseAccountInRoleCreateSpecListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.EnterpriseRole {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("EnterpriseRole[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o EnterpriseAccountInRoleCreateSpecListType) String() (string) { return o.FullString(false,0) }

/*
 * EnterpriseSecuritySettingsType 
 * Not validated 
 */
func (o EnterpriseSecuritySettingsType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o EnterpriseSecuritySettingsType) String() (string) { return o.FullString(false,0) }

/*
 * EnterpriseAccountCreateSpecType 
 * Not validated 
 */
func (o EnterpriseAccountCreateSpecType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("AccountType                    : ");buffer.WriteString(o.AccountType.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("AccountName                    : ");buffer.WriteString(fmt.Sprintf("%s",o.AccountName))
  if full && o.Roles != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Roles                          : ");buffer.WriteString(o.Roles.FullString(full,depth+1))
  }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("AllowRestoreAllVms             : ");buffer.WriteString(fmt.Sprintf("%t",o.AllowRestoreAllVms))
  if full && o.HierarchyScopeObjects != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("HierarchyScopeObjects          : ");buffer.WriteString(o.HierarchyScopeObjects.FullString(full,depth+1))
  }
  if full && o.FlrSettings != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FlrSettings                    : ");buffer.WriteString(o.FlrSettings.FullString(full,depth+1))
  }
   //Inhereting from SpecType
   //Inhereting from ParamsType
  return buffer.String()
}
func (o EnterpriseAccountCreateSpecType) String() (string) { return o.FullString(false,0) }

/*
 * FileRestoreSettingsSpecsType 
 * Not validated 
 */
func (o FileRestoreSettingsSpecsType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FlrInlaceOnly                  : ");buffer.WriteString(fmt.Sprintf("%t",o.FlrInlaceOnly))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FlrExtentionRestrictions       : ");buffer.WriteString(fmt.Sprintf("%s",o.FlrExtentionRestrictions))
   //Inhereting from SpecType
   //Inhereting from ParamsType
  return buffer.String()
}
func (o FileRestoreSettingsSpecsType) String() (string) { return o.FullString(false,0) }

/*
 * RebuildScopeJobSpecType 
 * Not validated 
 */
func (o RebuildScopeJobSpecType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.RebuildAll != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("RebuildAll                     : ");buffer.WriteString(o.RebuildAll.FullString(full,depth+1))
  }
  if full && o.RebuildUnprocessed != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("RebuildUnprocessed             : ");buffer.WriteString(o.RebuildUnprocessed.FullString(full,depth+1))
  }
  if full && o.RebuildForCurrentUser != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("RebuildForCurrentUser          : ");buffer.WriteString(o.RebuildForCurrentUser.FullString(full,depth+1))
  }
   //Inhereting from SpecType
   //Inhereting from ParamsType
  return buffer.String()
}
func (o RebuildScopeJobSpecType) String() (string) { return o.FullString(false,0) }

/*
 * RebuildScopeJobSpecTypeNestedRebuildAll 
 * Not validated 
 */
func (o RebuildScopeJobSpecTypeNestedRebuildAll) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  return buffer.String()
}
func (o RebuildScopeJobSpecTypeNestedRebuildAll) String() (string) { return o.FullString(false,0) }

/*
 * RebuildScopeJobSpecTypeNestedRebuildUnprocessed 
 * Not validated 
 */
func (o RebuildScopeJobSpecTypeNestedRebuildUnprocessed) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  return buffer.String()
}
func (o RebuildScopeJobSpecTypeNestedRebuildUnprocessed) String() (string) { return o.FullString(false,0) }

/*
 * RebuildScopeJobSpecTypeNestedRebuildForCurrentUser 
 * Not validated 
 */
func (o RebuildScopeJobSpecTypeNestedRebuildForCurrentUser) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  return buffer.String()
}
func (o RebuildScopeJobSpecTypeNestedRebuildForCurrentUser) String() (string) { return o.FullString(false,0) }

/*
 * WanAcceleratorEntityType 
 * Not validated 
 */
func (o WanAcceleratorEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Description                    : ");buffer.WriteString(fmt.Sprintf("%s",o.Description))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("OutOfDate                      : ");buffer.WriteString(fmt.Sprintf("%t",o.OutOfDate))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Version                        : ");buffer.WriteString(fmt.Sprintf("%s",o.Version))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Capacity                       : ");buffer.WriteString(fmt.Sprintf("%d",o.Capacity))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("TrafficPort                    : ");buffer.WriteString(fmt.Sprintf("%d",o.TrafficPort))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ConnectionsCount               : ");buffer.WriteString(fmt.Sprintf("%d",o.ConnectionsCount))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CachePath                      : ");buffer.WriteString(fmt.Sprintf("%s",o.CachePath))
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o WanAcceleratorEntityType) String() (string) { return o.FullString(false,0) }

/*
 * WanAcceleratorEntityListType 
 * Not validated 
 */
func (o WanAcceleratorEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.WanAccelerator {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("WanAccelerator[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o WanAcceleratorEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * CloudGatewayEntityType 
 * Not validated 
 */
func (o CloudGatewayEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Enabled                        : ");buffer.WriteString(fmt.Sprintf("%t",o.Enabled))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("DnsNameOrIpAddress             : ");buffer.WriteString(fmt.Sprintf("%s",o.DnsNameOrIpAddress))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("NetworkMode                    : ");buffer.WriteString(o.NetworkMode.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ExternalIP                     : ");buffer.WriteString(fmt.Sprintf("%s",o.ExternalIP))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ExternalPort                   : ");buffer.WriteString(fmt.Sprintf("%d",o.ExternalPort))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("InternalPort                   : ");buffer.WriteString(fmt.Sprintf("%d",o.InternalPort))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Description                    : ");buffer.WriteString(fmt.Sprintf("%s",o.Description))
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o CloudGatewayEntityType) String() (string) { return o.FullString(false,0) }

/*
 * CloudGatewayEntityListType 
 * Not validated 
 */
func (o CloudGatewayEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.CloudGateway {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("CloudGateway[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o CloudGatewayEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * CloudTenantEntityListType 
 * Not validated 
 */
func (o CloudTenantEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.CloudTenant {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("CloudTenant[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o CloudTenantEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * CloudTenantEntityType 
 * Not validated 
 */
func (o CloudTenantEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Password                       : ");buffer.WriteString(fmt.Sprintf("%s",o.Password))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Description                    : ");buffer.WriteString(fmt.Sprintf("%s",o.Description))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Enabled                        : ");buffer.WriteString(fmt.Sprintf("%t",o.Enabled))
  if full && o.LeaseOptions != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("LeaseOptions                   : ");buffer.WriteString(o.LeaseOptions.FullString(full,depth+1))
  }
  if full && o.Resources != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Resources                      : ");buffer.WriteString(o.Resources.FullString(full,depth+1))
  }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("LastResult                     : ");buffer.WriteString(fmt.Sprintf("%s",o.LastResult))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("LastActive                     : ");buffer.WriteString(o.LastActive.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VmCount                        : ");buffer.WriteString(fmt.Sprintf("%d",o.VmCount))
  if full && o.ComputeResources != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ComputeResources               : ");buffer.WriteString(o.ComputeResources.FullString(full,depth+1))
  }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ThrottlingEnabled              : ");buffer.WriteString(fmt.Sprintf("%t",o.ThrottlingEnabled))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ThrottlingSpeedLimit           : ");buffer.WriteString(fmt.Sprintf("%d",o.ThrottlingSpeedLimit))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ThrottlingSpeedUnit            : ");buffer.WriteString(fmt.Sprintf("%s",o.ThrottlingSpeedUnit))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("PublicIpCount                  : ");buffer.WriteString(fmt.Sprintf("%d",o.PublicIpCount))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("BackupCount                    : ");buffer.WriteString(fmt.Sprintf("%d",o.BackupCount))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ReplicaCount                   : ");buffer.WriteString(fmt.Sprintf("%d",o.ReplicaCount))
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o CloudTenantEntityType) String() (string) { return o.FullString(false,0) }

/*
 * CloudTenantLeaseOptionsType 
 * Not validated 
 */
func (o CloudTenantLeaseOptionsType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Enabled                        : ");buffer.WriteString(fmt.Sprintf("%t",o.Enabled))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("ExpirationDate                 : ");buffer.WriteString(o.ExpirationDate.String())
   //Inhereting from InfoType
  return buffer.String()
}
func (o CloudTenantLeaseOptionsType) String() (string) { return o.FullString(false,0) }

/*
 * CloudTenantResourceType 
 * Not validated 
 */
func (o CloudTenantResourceType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.RepositoryQuota != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("RepositoryQuota                : ");buffer.WriteString(o.RepositoryQuota.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Id                             : ");buffer.WriteString(fmt.Sprintf("%s",o.Id))
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o CloudTenantResourceType) String() (string) { return o.FullString(false,0) }

/*
 * CloudTenantResourceListType 
 * Not validated 
 */
func (o CloudTenantResourceListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.CloudTenantResource {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("CloudTenantResource[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o CloudTenantResourceListType) String() (string) { return o.FullString(false,0) }

/*
 * CloudTenantRepositoryQuotaInfoType 
 * Not validated 
 */
func (o CloudTenantRepositoryQuotaInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("DisplayName                    : ");buffer.WriteString(fmt.Sprintf("%s",o.DisplayName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("RepositoryUid                  : ");buffer.WriteString(o.RepositoryUid.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("WanAcceleratorUid              : ");buffer.WriteString(o.WanAcceleratorUid.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Quota                          : ");buffer.WriteString(fmt.Sprintf("%d",o.Quota))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("UsedQuota                      : ");buffer.WriteString(fmt.Sprintf("%d",o.UsedQuota))
   //Inhereting from InfoType
  return buffer.String()
}
func (o CloudTenantRepositoryQuotaInfoType) String() (string) { return o.FullString(false,0) }

/*
 * CreateCloudGatewaySpecType 
 * Not validated 
 */
func (o CreateCloudGatewaySpecType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("BackupServerIdOrName           : ");buffer.WriteString(fmt.Sprintf("%s",o.BackupServerIdOrName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ServerHostName                 : ");buffer.WriteString(fmt.Sprintf("%s",o.ServerHostName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Description                    : ");buffer.WriteString(fmt.Sprintf("%s",o.Description))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("IncomingPort                   : ");buffer.WriteString(fmt.Sprintf("%d",o.IncomingPort))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ExternalIp                     : ");buffer.WriteString(fmt.Sprintf("%s",o.ExternalIp))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ExternalPort                   : ");buffer.WriteString(fmt.Sprintf("%d",o.ExternalPort))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("NetworkType                    : ");buffer.WriteString(o.NetworkType.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("InternalPort                   : ");buffer.WriteString(fmt.Sprintf("%d",o.InternalPort))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("BackupServerUid                : ");buffer.WriteString(o.BackupServerUid.String())
   //Inhereting from SpecType
   //Inhereting from ParamsType
  return buffer.String()
}
func (o CreateCloudGatewaySpecType) String() (string) { return o.FullString(false,0) }

/*
 * CreateCloudTenantSpecType 
 * Not validated 
 */
func (o CreateCloudTenantSpecType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("BackupServerIdOrName           : ");buffer.WriteString(fmt.Sprintf("%s",o.BackupServerIdOrName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Description                    : ");buffer.WriteString(fmt.Sprintf("%s",o.Description))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Password                       : ");buffer.WriteString(fmt.Sprintf("%s",o.Password))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Enabled                        : ");buffer.WriteString(fmt.Sprintf("%t",o.Enabled))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("LeaseExpirationDate            : ");buffer.WriteString(o.LeaseExpirationDate.String())
  if full && o.Resources != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Resources                      : ");buffer.WriteString(o.Resources.FullString(full,depth+1))
  }
  if full && o.ComputeResources != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ComputeResources               : ");buffer.WriteString(o.ComputeResources.FullString(full,depth+1))
  }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VmCount                        : ");buffer.WriteString(fmt.Sprintf("%d",o.VmCount))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ThrottlingEnabled              : ");buffer.WriteString(fmt.Sprintf("%t",o.ThrottlingEnabled))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ThrottlingSpeedLimit           : ");buffer.WriteString(fmt.Sprintf("%d",o.ThrottlingSpeedLimit))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ThrottlingSpeedUnit            : ");buffer.WriteString(fmt.Sprintf("%s",o.ThrottlingSpeedUnit))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("PublicIpCount                  : ");buffer.WriteString(fmt.Sprintf("%d",o.PublicIpCount))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("BackupServerUid                : ");buffer.WriteString(o.BackupServerUid.String())
   //Inhereting from SpecType
   //Inhereting from ParamsType
  return buffer.String()
}
func (o CreateCloudTenantSpecType) String() (string) { return o.FullString(false,0) }

/*
 * CreateCloudTenantResourceListType 
 * Not validated 
 */
func (o CreateCloudTenantResourceListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.BackupResource {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("BackupResource[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o CreateCloudTenantResourceListType) String() (string) { return o.FullString(false,0) }

/*
 * CreateCloudTenantResourceSpecType 
 * Not validated 
 */
func (o CreateCloudTenantResourceSpecType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("RepositoryUid                  : ");buffer.WriteString(o.RepositoryUid.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("QuotaMb                        : ");buffer.WriteString(fmt.Sprintf("%d",o.QuotaMb))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("WanAcceleratorUid              : ");buffer.WriteString(o.WanAcceleratorUid.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Folder                         : ");buffer.WriteString(fmt.Sprintf("%s",o.Folder))
   //Inhereting from SpecType
   //Inhereting from ParamsType
  return buffer.String()
}
func (o CreateCloudTenantResourceSpecType) String() (string) { return o.FullString(false,0) }

/*
 * CloudHardwarePlanEntityListType 
 * Not validated 
 */
func (o CloudHardwarePlanEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.CloudHardwarePlan {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("CloudHardwarePlan[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o CloudHardwarePlanEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * CloudHardwarePlanEntityType 
 * Not validated 
 */
func (o CloudHardwarePlanEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Description                    : ");buffer.WriteString(fmt.Sprintf("%s",o.Description))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ProcessorUsageLimitMhz         : ");buffer.WriteString(fmt.Sprintf("%d",o.ProcessorUsageLimitMhz))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("MemoryUsageLimitMb             : ");buffer.WriteString(fmt.Sprintf("%d",o.MemoryUsageLimitMb))
  if full && o.HardwarePlanDetails != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("HardwarePlanDetails            : ");buffer.WriteString(o.HardwarePlanDetails.FullString(full,depth+1))
  }
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o CloudHardwarePlanEntityType) String() (string) { return o.FullString(false,0) }

/*
 * CloudHardwarePlanEntityTypeNestedHardwarePlanDetails 
 * Not validated 
 */
func (o CloudHardwarePlanEntityTypeNestedHardwarePlanDetails) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.ViCloudHardwarePlan != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ViCloudHardwarePlan            : ");buffer.WriteString(o.ViCloudHardwarePlan.FullString(full,depth+1))
  }
  if full && o.HvCloudHardwarePlan != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("HvCloudHardwarePlan            : ");buffer.WriteString(o.HvCloudHardwarePlan.FullString(full,depth+1))
  }
  return buffer.String()
}
func (o CloudHardwarePlanEntityTypeNestedHardwarePlanDetails) String() (string) { return o.FullString(false,0) }

/*
 * CloudHardwarePlanCreateSpecType 
 * Not validated 
 */
func (o CloudHardwarePlanCreateSpecType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("BackupServerUid                : ");buffer.WriteString(o.BackupServerUid.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Description                    : ");buffer.WriteString(fmt.Sprintf("%s",o.Description))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ProcessorUsageLimitMhz         : ");buffer.WriteString(fmt.Sprintf("%d",o.ProcessorUsageLimitMhz))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("MemoryUsageLimitMb             : ");buffer.WriteString(fmt.Sprintf("%d",o.MemoryUsageLimitMb))
  if full && o.HardwarePlanDetails != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("HardwarePlanDetails            : ");buffer.WriteString(o.HardwarePlanDetails.FullString(full,depth+1))
  }
   //Inhereting from SpecType
   //Inhereting from ParamsType
  return buffer.String()
}
func (o CloudHardwarePlanCreateSpecType) String() (string) { return o.FullString(false,0) }

/*
 * CloudHardwarePlanCreateSpecTypeNestedHardwarePlanDetails 
 * Not validated 
 */
func (o CloudHardwarePlanCreateSpecTypeNestedHardwarePlanDetails) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.ViCloudHardwarePlan != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ViCloudHardwarePlan            : ");buffer.WriteString(o.ViCloudHardwarePlan.FullString(full,depth+1))
  }
  if full && o.HvCloudHardwarePlan != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("HvCloudHardwarePlan            : ");buffer.WriteString(o.HvCloudHardwarePlan.FullString(full,depth+1))
  }
  return buffer.String()
}
func (o CloudHardwarePlanCreateSpecTypeNestedHardwarePlanDetails) String() (string) { return o.FullString(false,0) }

/*
 * ViCloudHardwarePlanInfoType 
 * Not validated 
 */
func (o ViCloudHardwarePlanInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("HypervisorHostRef              : ");buffer.WriteString(o.HypervisorHostRef.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ParentType                     : ");buffer.WriteString(fmt.Sprintf("%s",o.ParentType))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ParentName                     : ");buffer.WriteString(fmt.Sprintf("%s",o.ParentName))
  if full && o.Datastores != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Datastores                     : ");buffer.WriteString(o.Datastores.FullString(full,depth+1))
  }
  if full && o.Network != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Network                        : ");buffer.WriteString(o.Network.FullString(full,depth+1))
  }
   //Inhereting from InfoType
  return buffer.String()
}
func (o ViCloudHardwarePlanInfoType) String() (string) { return o.FullString(false,0) }

/*
 * ViCloudHardwarePlanDatastoreInfoListType 
 * Not validated 
 */
func (o ViCloudHardwarePlanDatastoreInfoListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.Datastore {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("Datastore[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o ViCloudHardwarePlanDatastoreInfoListType) String() (string) { return o.FullString(false,0) }

/*
 * ViCloudHardwarePlanDatastoreInfoType 
 * Not validated 
 */
func (o ViCloudHardwarePlanDatastoreInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FriendlyName                   : ");buffer.WriteString(fmt.Sprintf("%s",o.FriendlyName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("DatastoreType                  : ");buffer.WriteString(fmt.Sprintf("%s",o.DatastoreType))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Reference                      : ");buffer.WriteString(o.Reference.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("RootPath                       : ");buffer.WriteString(fmt.Sprintf("%s",o.RootPath))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("QuotaGb                        : ");buffer.WriteString(fmt.Sprintf("%d",o.QuotaGb))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("PbmProfileId                   : ");buffer.WriteString(fmt.Sprintf("%s",o.PbmProfileId))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Id                             : ");buffer.WriteString(fmt.Sprintf("%s",o.Id))
   //Inhereting from InfoType
  return buffer.String()
}
func (o ViCloudHardwarePlanDatastoreInfoType) String() (string) { return o.FullString(false,0) }

/*
 * HvCloudHardwarePlanInfoType 
 * Not validated 
 */
func (o HvCloudHardwarePlanInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("HypervisorHostRef              : ");buffer.WriteString(o.HypervisorHostRef.String())
  if full && o.Volumes != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Volumes                        : ");buffer.WriteString(o.Volumes.FullString(full,depth+1))
  }
  if full && o.Network != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Network                        : ");buffer.WriteString(o.Network.FullString(full,depth+1))
  }
   //Inhereting from InfoType
  return buffer.String()
}
func (o HvCloudHardwarePlanInfoType) String() (string) { return o.FullString(false,0) }

/*
 * HvCloudHardwarePlanVolumesInfoListType 
 * Not validated 
 */
func (o HvCloudHardwarePlanVolumesInfoListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.Volume {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("Volume[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o HvCloudHardwarePlanVolumesInfoListType) String() (string) { return o.FullString(false,0) }

/*
 * HvCloudHardwarePlanVolumeInfoType 
 * Not validated 
 */
func (o HvCloudHardwarePlanVolumeInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FriendlyName                   : ");buffer.WriteString(fmt.Sprintf("%s",o.FriendlyName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VolumePath                     : ");buffer.WriteString(fmt.Sprintf("%s",o.VolumePath))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("QuotaGb                        : ");buffer.WriteString(fmt.Sprintf("%d",o.QuotaGb))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Id                             : ");buffer.WriteString(fmt.Sprintf("%s",o.Id))
   //Inhereting from InfoType
  return buffer.String()
}
func (o HvCloudHardwarePlanVolumeInfoType) String() (string) { return o.FullString(false,0) }

/*
 * CloudHardwarePlanNetworkInfo 
 * Not validated 
 */
func (o CloudHardwarePlanNetworkInfo) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CountWithInternet              : ");buffer.WriteString(fmt.Sprintf("%d",o.CountWithInternet))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CountWithoutInternet           : ");buffer.WriteString(fmt.Sprintf("%d",o.CountWithoutInternet))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Id                             : ");buffer.WriteString(fmt.Sprintf("%s",o.Id))
   //Inhereting from InfoType
  return buffer.String()
}
func (o CloudHardwarePlanNetworkInfo) String() (string) { return o.FullString(false,0) }

/*
 * CloudFailoverSessionEntityType 
 * Not validated 
 */
func (o CloudFailoverSessionEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.CloudFailoverTasks != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CloudFailoverTasks             : ");buffer.WriteString(o.CloudFailoverTasks.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("CloudFailoverPlanName          : ");buffer.WriteString(fmt.Sprintf("%s",o.CloudFailoverPlanName))
   //Inhereting from JobSessionEntityType
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("JobUid                         : ");buffer.WriteString(o.JobUid.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("JobName                        : ");buffer.WriteString(fmt.Sprintf("%s",o.JobName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("JobType                        : ");buffer.WriteString(fmt.Sprintf("%s",o.JobType))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CreationTimeUTC                : ");buffer.WriteString(o.CreationTimeUTC.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("EndTimeUTC                     : ");buffer.WriteString(o.EndTimeUTC.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("State                          : ");buffer.WriteString(fmt.Sprintf("%s",o.State))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Result                         : ");buffer.WriteString(fmt.Sprintf("%s",o.Result))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Progress                       : ");buffer.WriteString(fmt.Sprintf("%d",o.Progress))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FailureMessage                 : ");buffer.WriteString(fmt.Sprintf("%s",o.FailureMessage))
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o CloudFailoverSessionEntityType) String() (string) { return o.FullString(false,0) }

/*
 * CloudFailoverSessionEntityListType 
 * Not validated 
 */
func (o CloudFailoverSessionEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.CloudFailoverSession {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("CloudFailoverSession[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o CloudFailoverSessionEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * CloudFailoverTaskSessionInfoType 
 * Not validated 
 */
func (o CloudFailoverTaskSessionInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.VmReplicaPointLink != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VmReplicaPointLink             : ");buffer.WriteString(o.VmReplicaPointLink.FullString(full,depth+1))
  }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CreationTimeUTC                : ");buffer.WriteString(o.CreationTimeUTC.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("EndTimeUTC                     : ");buffer.WriteString(o.EndTimeUTC.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("State                          : ");buffer.WriteString(fmt.Sprintf("%s",o.State))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Result                         : ");buffer.WriteString(fmt.Sprintf("%s",o.Result))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Progress                       : ");buffer.WriteString(fmt.Sprintf("%d",o.Progress))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FailureMessage                 : ");buffer.WriteString(fmt.Sprintf("%s",o.FailureMessage))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("VmName                         : ");buffer.WriteString(fmt.Sprintf("%s",o.VmName))
   //Inhereting from InfoType
  return buffer.String()
}
func (o CloudFailoverTaskSessionInfoType) String() (string) { return o.FullString(false,0) }

/*
 * CloudFailoverTaskSessionInfoListType 
 * Not validated 
 */
func (o CloudFailoverTaskSessionInfoListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.CloudFailoverTasks {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("CloudFailoverTasks[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o CloudFailoverTaskSessionInfoListType) String() (string) { return o.FullString(false,0) }

/*
 * CloudPublicIpAddressEntityListType 
 * Not validated 
 */
func (o CloudPublicIpAddressEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.CloudPublicIp {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("CloudPublicIp[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o CloudPublicIpAddressEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * CloudPublicIpAddressEntityType 
 * Not validated 
 */
func (o CloudPublicIpAddressEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("IpAddress                      : ");buffer.WriteString(fmt.Sprintf("%s",o.IpAddress))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("TenantUid                      : ");buffer.WriteString(o.TenantUid.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("BackupServerUid                : ");buffer.WriteString(o.BackupServerUid.String())
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o CloudPublicIpAddressEntityType) String() (string) { return o.FullString(false,0) }

/*
 * CloudPublicIpAddressCreateSpecType 
 * Not validated 
 */
func (o CloudPublicIpAddressCreateSpecType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("BackupServerUid                : ");buffer.WriteString(o.BackupServerUid.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("IpAddressLowerBound            : ");buffer.WriteString(fmt.Sprintf("%s",o.IpAddressLowerBound))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("IpAddressUpperBound            : ");buffer.WriteString(fmt.Sprintf("%s",o.IpAddressUpperBound))
   //Inhereting from SpecType
   //Inhereting from ParamsType
  return buffer.String()
}
func (o CloudPublicIpAddressCreateSpecType) String() (string) { return o.FullString(false,0) }

/*
 * CloudTenantComputeResourceListType 
 * Not validated 
 */
func (o CloudTenantComputeResourceListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.CloudTenantComputeResource {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("CloudTenantComputeResource[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o CloudTenantComputeResourceListType) String() (string) { return o.FullString(false,0) }

/*
 * CloudTenantComputeResourceType 
 * Not validated 
 */
func (o CloudTenantComputeResourceType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CloudHardwarePlanUid           : ");buffer.WriteString(o.CloudHardwarePlanUid.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("WanAcceleratorUid              : ");buffer.WriteString(o.WanAcceleratorUid.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("PlatformType                   : ");buffer.WriteString(fmt.Sprintf("%s",o.PlatformType))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("UseNetworkFailoverResources    : ");buffer.WriteString(fmt.Sprintf("%t",o.UseNetworkFailoverResources))
  if full && o.NetworkAppliance != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("NetworkAppliance               : ");buffer.WriteString(o.NetworkAppliance.FullString(full,depth+1))
  }
  if full && o.ComputeResourceStats != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ComputeResourceStats           : ");buffer.WriteString(o.ComputeResourceStats.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Id                             : ");buffer.WriteString(fmt.Sprintf("%s",o.Id))
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o CloudTenantComputeResourceType) String() (string) { return o.FullString(false,0) }

/*
 * ComputeResourceStatsInfoType 
 * Not validated 
 */
func (o ComputeResourceStatsInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("MemoryUsageMb                  : ");buffer.WriteString(fmt.Sprintf("%d",o.MemoryUsageMb))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CPUCount                       : ");buffer.WriteString(fmt.Sprintf("%d",o.CPUCount))
  if full && o.StorageResourceStats != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("StorageResourceStats           : ");buffer.WriteString(o.StorageResourceStats.FullString(full,depth+1))
  }
   //Inhereting from InfoType
  return buffer.String()
}
func (o ComputeResourceStatsInfoType) String() (string) { return o.FullString(false,0) }

/*
 * StorageResourceStatsListType 
 * Not validated 
 */
func (o StorageResourceStatsListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.StorageResourceStat {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("StorageResourceStat[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o StorageResourceStatsListType) String() (string) { return o.FullString(false,0) }

/*
 * StorageResourceStatInfoType 
 * Not validated 
 */
func (o StorageResourceStatInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("StorageName                    : ");buffer.WriteString(fmt.Sprintf("%s",o.StorageName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("StorageUsageGb                 : ");buffer.WriteString(fmt.Sprintf("%d",o.StorageUsageGb))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("StorageLimitGb                 : ");buffer.WriteString(fmt.Sprintf("%d",o.StorageLimitGb))
   //Inhereting from InfoType
  return buffer.String()
}
func (o StorageResourceStatInfoType) String() (string) { return o.FullString(false,0) }

/*
 * CloudTenantComputeResourceLeaseOptionsType 
 * Not validated 
 */
func (o CloudTenantComputeResourceLeaseOptionsType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Enabled                        : ");buffer.WriteString(fmt.Sprintf("%t",o.Enabled))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("ExpirationDate                 : ");buffer.WriteString(o.ExpirationDate.String())
   //Inhereting from InfoType
  return buffer.String()
}
func (o CloudTenantComputeResourceLeaseOptionsType) String() (string) { return o.FullString(false,0) }

/*
 * CloudTenantComputeResourceCreateListType 
 * Not validated 
 */
func (o CloudTenantComputeResourceCreateListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.ComputeResource {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("ComputeResource[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o CloudTenantComputeResourceCreateListType) String() (string) { return o.FullString(false,0) }

/*
 * CloudTenantComputeResourceCreateSpecType 
 * Not validated 
 */
func (o CloudTenantComputeResourceCreateSpecType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CloudHardwarePlanUid           : ");buffer.WriteString(o.CloudHardwarePlanUid.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("WanAcceleratorUid              : ");buffer.WriteString(o.WanAcceleratorUid.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("PlatformType                   : ");buffer.WriteString(fmt.Sprintf("%s",o.PlatformType))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("UseNetworkFailoverResources    : ");buffer.WriteString(fmt.Sprintf("%t",o.UseNetworkFailoverResources))
   //Inhereting from SpecType
   //Inhereting from ParamsType
  return buffer.String()
}
func (o CloudTenantComputeResourceCreateSpecType) String() (string) { return o.FullString(false,0) }

/*
 * CloudVmReplicaPointEntityType 
 * Not validated 
 */
func (o CloudVmReplicaPointEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CreationTimeUTC                : ");buffer.WriteString(o.CreationTimeUTC.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VmName                         : ");buffer.WriteString(fmt.Sprintf("%s",o.VmName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("PointType                      : ");buffer.WriteString(fmt.Sprintf("%s",o.PointType))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("State                          : ");buffer.WriteString(fmt.Sprintf("%s",o.State))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("VmDisplayName                  : ");buffer.WriteString(fmt.Sprintf("%s",o.VmDisplayName))
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o CloudVmReplicaPointEntityType) String() (string) { return o.FullString(false,0) }

/*
 * CloudVmReplicaPointEntityListType 
 * Not validated 
 */
func (o CloudVmReplicaPointEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.CloudReplica {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("CloudReplica[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o CloudVmReplicaPointEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * CloudFailoverPlanEntityListType 
 * Not validated 
 */
func (o CloudFailoverPlanEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.CloudFailoverPlan {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("CloudFailoverPlan[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o CloudFailoverPlanEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * CloudFailoverPlanEntityType 
 * Not validated 
 */
func (o CloudFailoverPlanEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("TenantUid                      : ");buffer.WriteString(o.TenantUid.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("TenantName                     : ");buffer.WriteString(fmt.Sprintf("%s",o.TenantName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Description                    : ");buffer.WriteString(fmt.Sprintf("%s",o.Description))
  if full && o.CloudFailoverPlanOptions != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CloudFailoverPlanOptions       : ");buffer.WriteString(o.CloudFailoverPlanOptions.FullString(full,depth+1))
  }
  if full && o.CloudFailoverPlanInfo != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CloudFailoverPlanInfo          : ");buffer.WriteString(o.CloudFailoverPlanInfo.FullString(full,depth+1))
  }
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o CloudFailoverPlanEntityType) String() (string) { return o.FullString(false,0) }

/*
 * CloudFailoverPlanManagementSpecType 
 * Not validated 
 */
func (o CloudFailoverPlanManagementSpecType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("StartNow                       : ");buffer.WriteString(fmt.Sprintf("%t",o.StartNow))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("StartDate                      : ");buffer.WriteString(o.StartDate.String())
   //Inhereting from SpecType
   //Inhereting from ParamsType
  return buffer.String()
}
func (o CloudFailoverPlanManagementSpecType) String() (string) { return o.FullString(false,0) }

/*
 * CloudFailoverPlanInfoType 
 * Not validated 
 */
func (o CloudFailoverPlanInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.Includes != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Includes                       : ");buffer.WriteString(o.Includes.FullString(full,depth+1))
  }
   //Inhereting from InfoType
  return buffer.String()
}
func (o CloudFailoverPlanInfoType) String() (string) { return o.FullString(false,0) }

/*
 * CloudFailoveredVmType 
 * Not validated 
 */
func (o CloudFailoveredVmType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FailoverPlanVMId               : ");buffer.WriteString(fmt.Sprintf("%s",o.FailoverPlanVMId))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Order                          : ");buffer.WriteString(fmt.Sprintf("%d",o.Order))
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o CloudFailoveredVmType) String() (string) { return o.FullString(false,0) }

/*
 * CloudFailoveredVmListType 
 * Not validated 
 */
func (o CloudFailoveredVmListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.CloudFailoveredVm {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("CloudFailoveredVm[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o CloudFailoveredVmListType) String() (string) { return o.FullString(false,0) }

/*
 * CloudFailoverPlanOptionsInfoType 
 * Not validated 
 */
func (o CloudFailoverPlanOptionsInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("PostFailoverPlanCommandEnabled : ");buffer.WriteString(fmt.Sprintf("%t",o.PostFailoverPlanCommandEnabled))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("PostFailoverPlanCommand        : ");buffer.WriteString(fmt.Sprintf("%s",o.PostFailoverPlanCommand))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("PreFailoverPlanCommandEnabled  : ");buffer.WriteString(fmt.Sprintf("%t",o.PreFailoverPlanCommandEnabled))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("PreFailoverPlanCommand         : ");buffer.WriteString(fmt.Sprintf("%s",o.PreFailoverPlanCommand))
   //Inhereting from InfoType
  return buffer.String()
}
func (o CloudFailoverPlanOptionsInfoType) String() (string) { return o.FullString(false,0) }

/*
 * CloudReplicaEntityType 
 * Not validated 
 */
func (o CloudReplicaEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.Ref != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Ref                            : ");buffer.WriteString(o.Ref.FullString(full,depth+1))
  }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Platform                       : ");buffer.WriteString(fmt.Sprintf("%s",o.Platform))
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o CloudReplicaEntityType) String() (string) { return o.FullString(false,0) }

/*
 * CloudReplicaEntityListType 
 * Not validated 
 */
func (o CloudReplicaEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.CloudReplica {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("CloudReplica[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o CloudReplicaEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * CloudConnectServiceType 
 * Not validated 
 */
func (o CloudConnectServiceType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o CloudConnectServiceType) String() (string) { return o.FullString(false,0) }

/*
 * TenantCredentialsInfoType 
 * Not validated 
 */
func (o TenantCredentialsInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Username                       : ");buffer.WriteString(fmt.Sprintf("%s",o.Username))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Password                       : ");buffer.WriteString(fmt.Sprintf("%s",o.Password))
   //Inhereting from InfoType
  return buffer.String()
}
func (o TenantCredentialsInfoType) String() (string) { return o.FullString(false,0) }

/*
 * VlanConfigurationEntityListType 
 * Not validated 
 */
func (o VlanConfigurationEntityListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.Vlans {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("Vlans[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o VlanConfigurationEntityListType) String() (string) { return o.FullString(false,0) }

/*
 * VlanConfigurationEntityType 
 * Not validated 
 */
func (o VlanConfigurationEntityType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("HostRef                        : ");buffer.WriteString(o.HostRef.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("PlatformType                   : ");buffer.WriteString(fmt.Sprintf("%s",o.PlatformType))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VlanIdsWithInternetLeftBound   : ");buffer.WriteString(fmt.Sprintf("%d",o.VlanIdsWithInternetLeftBound))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VlanIdsWithInternetRightBound  : ");buffer.WriteString(fmt.Sprintf("%d",o.VlanIdsWithInternetRightBound))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VlanIdsWithoutInternetLeftBound : ");buffer.WriteString(fmt.Sprintf("%d",o.VlanIdsWithoutInternetLeftBound))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VlanIdsWithoutInternetRightBound : ");buffer.WriteString(fmt.Sprintf("%d",o.VlanIdsWithoutInternetRightBound))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("SwitchName                     : ");buffer.WriteString(fmt.Sprintf("%s",o.SwitchName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("SwitchId                       : ");buffer.WriteString(fmt.Sprintf("%s",o.SwitchId))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ClusterName                    : ");buffer.WriteString(fmt.Sprintf("%s",o.ClusterName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ClusterId                      : ");buffer.WriteString(fmt.Sprintf("%s",o.ClusterId))
   //Inhereting from EntityType
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("UID                            : ");buffer.WriteString(o.UID.String())
   //Inhereting from ResourceType
  if full && o.Links != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Links                          : ");buffer.WriteString(o.Links.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Href                           : ");buffer.WriteString(o.Href.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  return buffer.String()
}
func (o VlanConfigurationEntityType) String() (string) { return o.FullString(false,0) }

/*
 * CloudVlanConfigurationCreateSpecType 
 * Not validated 
 */
func (o CloudVlanConfigurationCreateSpecType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("BackupServerUid                : ");buffer.WriteString(o.BackupServerUid.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("HostRef                        : ");buffer.WriteString(o.HostRef.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("PlatformType                   : ");buffer.WriteString(fmt.Sprintf("%s",o.PlatformType))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VlanIdsWithInternetLeftBound   : ");buffer.WriteString(fmt.Sprintf("%d",o.VlanIdsWithInternetLeftBound))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VlanIdsWithInternetRightBound  : ");buffer.WriteString(fmt.Sprintf("%d",o.VlanIdsWithInternetRightBound))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VlanIdsWithoutInternetLeftBound : ");buffer.WriteString(fmt.Sprintf("%d",o.VlanIdsWithoutInternetLeftBound))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VlanIdsWithoutInternetRightBound : ");buffer.WriteString(fmt.Sprintf("%d",o.VlanIdsWithoutInternetRightBound))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("viClusterRef                   : ");buffer.WriteString(fmt.Sprintf("%s",o.viClusterRef))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("viClusterName                  : ");buffer.WriteString(fmt.Sprintf("%s",o.viClusterName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("viSwitchName                   : ");buffer.WriteString(fmt.Sprintf("%s",o.viSwitchName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("viSwitchType                   : ");buffer.WriteString(fmt.Sprintf("%d",o.viSwitchType))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("viSwitchId                     : ");buffer.WriteString(fmt.Sprintf("%s",o.viSwitchId))
   //Inhereting from SpecType
   //Inhereting from ParamsType
  return buffer.String()
}
func (o CloudVlanConfigurationCreateSpecType) String() (string) { return o.FullString(false,0) }

/*
 * NetworkApplianceInfoType 
 * Not validated 
 */
func (o NetworkApplianceInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ProductionNetwork              : ");buffer.WriteString(fmt.Sprintf("%s",o.ProductionNetwork))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ObtainIPAddressAutomatically   : ");buffer.WriteString(fmt.Sprintf("%t",o.ObtainIPAddressAutomatically))
  if full && o.ManualIpAddressSettingsInfoType != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ManualIpAddressSettingsInfoType : ");buffer.WriteString(o.ManualIpAddressSettingsInfoType.FullString(full,depth+1))
  }
   //Inhereting from InfoType
  return buffer.String()
}
func (o NetworkApplianceInfoType) String() (string) { return o.FullString(false,0) }

/*
 * ManualIpAddressSettingsInfoType 
 * Not validated 
 */
func (o ManualIpAddressSettingsInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("IpAddress                      : ");buffer.WriteString(o.IpAddress.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("SubnetMask                     : ");buffer.WriteString(o.SubnetMask.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("DefaultGateway                 : ");buffer.WriteString(o.DefaultGateway.String())
   //Inhereting from InfoType
  return buffer.String()
}
func (o ManualIpAddressSettingsInfoType) String() (string) { return o.FullString(false,0) }

/*
 * VmRestorePointSqlInfoType 
 * Not validated 
 */
func (o VmRestorePointSqlInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.Databases != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Databases                      : ");buffer.WriteString(o.Databases.FullString(full,depth+1))
  }
   //Inhereting from InfoType
  return buffer.String()
}
func (o VmRestorePointSqlInfoType) String() (string) { return o.FullString(false,0) }

/*
 * VmRestorePointSqlDatabaseInfoListType 
 * Not validated 
 */
func (o VmRestorePointSqlDatabaseInfoListType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.Database {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("Database[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o VmRestorePointSqlDatabaseInfoListType) String() (string) { return o.FullString(false,0) }

/*
 * VmRestorePointSqlDatabaseInfoType 
 * Not validated 
 */
func (o VmRestorePointSqlDatabaseInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.Items != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Items                          : ");buffer.WriteString(o.Items.FullString(full,depth+1))
  }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Instance                       : ");buffer.WriteString(fmt.Sprintf("%s",o.Instance))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("InstanceName                   : ");buffer.WriteString(fmt.Sprintf("%s",o.InstanceName))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Name                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Name))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("FullVersion                    : ");buffer.WriteString(fmt.Sprintf("%s",o.FullVersion))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("ProductVersion                 : ");buffer.WriteString(fmt.Sprintf("%s",o.ProductVersion))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("CheckpointLsn                  : ");buffer.WriteString(fmt.Sprintf("%d",o.CheckpointLsn))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("DatabaseBackupLsn              : ");buffer.WriteString(fmt.Sprintf("%d",o.DatabaseBackupLsn))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("FirstLsn                       : ");buffer.WriteString(fmt.Sprintf("%d",o.FirstLsn))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("LastLsn                        : ");buffer.WriteString(fmt.Sprintf("%d",o.LastLsn))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("StartCreationTime              : ");buffer.WriteString(o.StartCreationTime.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("EndCreationTime                : ");buffer.WriteString(o.EndCreationTime.String())
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("TimeZone                       : ");buffer.WriteString(fmt.Sprintf("%d",o.TimeZone))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("RecoveryModel                  : ");buffer.WriteString(fmt.Sprintf("%s",o.RecoveryModel))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("CompatibilityLevel             : ");buffer.WriteString(fmt.Sprintf("%d",o.CompatibilityLevel))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("CertName                       : ");buffer.WriteString(fmt.Sprintf("%s",o.CertName))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("CertSerialNumber               : ");buffer.WriteString(fmt.Sprintf("%s",o.CertSerialNumber))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("ReadOnly                       : ");buffer.WriteString(fmt.Sprintf("%s",o.ReadOnly))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("BindingID                      : ");buffer.WriteString(fmt.Sprintf("%s",o.BindingID))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("FamilyGUID                     : ");buffer.WriteString(fmt.Sprintf("%s",o.FamilyGUID))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("DatabaseId                     : ");buffer.WriteString(fmt.Sprintf("%d",o.DatabaseId))
   //Inhereting from InfoType
  return buffer.String()
}
func (o VmRestorePointSqlDatabaseInfoType) String() (string) { return o.FullString(false,0) }

/*
 * VmRestorePointSqlDatabaseItemListInfoType 
 * Not validated 
 */
func (o VmRestorePointSqlDatabaseItemListInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full {
   for i,c := range o.Item {
    buffer.Write(identbuffer.Bytes())
    buffer.WriteString(fmt.Sprintf("Item[%d]",i))
    buffer.WriteString(c.FullString(full,depth+1))
   }
  }
   //Inhereting from ListType
  return buffer.String()
}
func (o VmRestorePointSqlDatabaseItemListInfoType) String() (string) { return o.FullString(false,0) }

/*
 * VmRestorePointSqlDatabaseItemInfoType 
 * Not validated 
 */
func (o VmRestorePointSqlDatabaseItemInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("Type                           : ");buffer.WriteString(fmt.Sprintf("%s",o.Type))
  buffer.Write(identbuffer.Bytes())
  buffer.WriteString("DataPath                       : ");buffer.WriteString(fmt.Sprintf("%s",o.DataPath))
   //Inhereting from InfoType
  return buffer.String()
}
func (o VmRestorePointSqlDatabaseItemInfoType) String() (string) { return o.FullString(false,0) }

/*
 * SqlItemRestoreSpecInfoType 
 * Not validated 
 */
func (o SqlItemRestoreSpecInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("IsRestoreToOriginal            : ");buffer.WriteString(fmt.Sprintf("%t",o.IsRestoreToOriginal))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ServerName                     : ");buffer.WriteString(fmt.Sprintf("%s",o.ServerName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VmName                         : ");buffer.WriteString(fmt.Sprintf("%s",o.VmName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("InstanceName                   : ");buffer.WriteString(fmt.Sprintf("%s",o.InstanceName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("FamilyUid                      : ");buffer.WriteString(fmt.Sprintf("%s",o.FamilyUid))
  if full && o.Credentials != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("Credentials                    : ");buffer.WriteString(o.Credentials.FullString(full,depth+1))
  }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("DatabaseName                   : ");buffer.WriteString(fmt.Sprintf("%s",o.DatabaseName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("SourceDatabaseName             : ");buffer.WriteString(fmt.Sprintf("%s",o.SourceDatabaseName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("SourceInstanceName             : ");buffer.WriteString(fmt.Sprintf("%s",o.SourceInstanceName))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("RestoreToDate                  : ");buffer.WriteString(o.RestoreToDate.String())
   //Inhereting from InfoType
  return buffer.String()
}
func (o SqlItemRestoreSpecInfoType) String() (string) { return o.FullString(false,0) }

/*
 * SqlCredentialsInfoType 
 * Not validated 
 */
func (o SqlCredentialsInfoType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
  if full && o.SqlCredentials != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("SqlCredentials                 : ");buffer.WriteString(o.SqlCredentials.FullString(full,depth+1))
  }
  if full && o.ServerCredentials != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("ServerCredentials              : ");buffer.WriteString(o.ServerCredentials.FullString(full,depth+1))
  }
  if full && o.UserCredentials != nil {
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("UserCredentials                : ");buffer.WriteString(o.UserCredentials.FullString(full,depth+1))
  }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("UseSqlAuth                     : ");buffer.WriteString(fmt.Sprintf("%t",o.UseSqlAuth))
   //Inhereting from InfoType
  return buffer.String()
}
func (o SqlCredentialsInfoType) String() (string) { return o.FullString(false,0) }

/*
 * VeeamZipStartupSpecType 
 * Not validated 
 */
func (o VeeamZipStartupSpecType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VmRef                          : ");buffer.WriteString(o.VmRef.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("RepositoryUid                  : ");buffer.WriteString(o.RepositoryUid.String())
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("PasswordKeyId                  : ");buffer.WriteString(fmt.Sprintf("%s",o.PasswordKeyId))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("CompressionLevel               : ");buffer.WriteString(fmt.Sprintf("%d",o.CompressionLevel))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("DisableGuestQuiescence         : ");buffer.WriteString(fmt.Sprintf("%t",o.DisableGuestQuiescence))
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("BackupRetention                : ");buffer.WriteString(o.BackupRetention.String())
   //Inhereting from SpecType
   //Inhereting from ParamsType
  return buffer.String()
}
func (o VeeamZipStartupSpecType) String() (string) { return o.FullString(false,0) }

/*
 * QuickBackupStartupSpecType 
 * Not validated 
 */
func (o QuickBackupStartupSpecType) FullString(full bool,depth int) (string) { 
  var buffer bytes.Buffer
  var identbuffer bytes.Buffer
  identbuffer.WriteString("\r\n");for i:=0;i <depth;i++ { identbuffer.WriteRune(' ') }
   buffer.Write(identbuffer.Bytes())
   buffer.WriteString("VmRef                          : ");buffer.WriteString(o.VmRef.String())
   //Inhereting from SpecType
   //Inhereting from ParamsType
  return buffer.String()
}
func (o QuickBackupStartupSpecType) String() (string) { return o.FullString(false,0) }

