package main

import(
	"github.com/tdewin/veeamrestapi"
	"log"
	"flag"
	"fmt"
	"encoding/xml"
	//"io/ioutil"
	//"os"
	//"regexp"
	"strings"
	"net/url"
	"sync"
	"net/http"
	"html/template"
	//"time"
	"time"
	"strconv"
	"bytes"
)


//cmd rest logger remove all verbose info unless requested with -v
type VeeamCMDRestLogger struct {
	verbose bool
}
func (l * VeeamCMDRestLogger) LogRequest(request string) {
	if l.verbose {
		log.Println(request)
	}
}
func (l * VeeamCMDRestLogger) LogInfo(info string) {
	if l.verbose {
		log.Println(info)
	}
}
func (l * VeeamCMDRestLogger) LogResponse(code int,xmlout string) {
	if l.verbose {
		log.Printf("Got return code %d",code)
		if code > 399 {
		  log.Println(xmlout)
		}
	}
}




//error
type cmdError struct { e string }
func (e cmdError) Error() (string) { return e.e }

type JobEditType struct { 
   XMLName xml.Name
   ScheduleConfigured bool `xml:"ScheduleConfigured"`
   Name string `xml:"Name,attr"`
   UID string `xml:"UID,attr"`
}
func NewJobEditType() (*JobEditType) {
  varType := JobEditType{}
  varType.XMLName.Local = "Job"
  varType.XMLName.Space = "http://www.veeam.com/ent/v1.0"
  return &varType
}


type HttpJobTaskSession struct {
	JobSessionUID string
	VmName string
	JobTaskSessionUID string
	CreateUTC string
	EndUTC string
	State string
	Result string
	Reason string
	TotalSize int64
}					
type HttpJobSession struct {
	SessionUID string
	JobUID string
	JobName string
	JobType string
	CreateUTC string
	EndUTC string
	State string
	Result string
	Progress int
	IsRetry bool
}
type HttpRestoreSession struct {
	SessionUID string
	Name string
	JobType string
	CreateUTC string
	EndUTC string
	State string
	Result string
	Progress int
}
type HttpJob struct {
	Name string
	UID string
	JobType string
	RealJob * veeamrestapi.JobEntityType
	VMs []*VM
}
type VM struct {
	Added bool
	Name string
	UID string
	Root *veeamrestapi.HierarchyRootEntityType
}

type HttpBackupVMRestorePoint struct {
	Name string
	UID string
	CreateUTC string
	Algorithm string
	PointType string
}
type HttpBackupRestorePoint struct {
	Name string
	UID string

}
type HttpBackup struct {
	Name string
	UID string
}

type MainTemplate struct {
	Title string
	Customer string
	Autorefresh string
	Content string
	Jobs []*HttpJob 
	Backups []*HttpBackup
	Sessions []*HttpJobSession
	JobTaskSessions []*HttpJobTaskSession
	BackupRestorePoints []*HttpBackupRestorePoint
	BackupVMRestorePoints []*HttpBackupVMRestorePoint
	RestoreSessions []*HttpRestoreSession
	Edit *HttpJob
	Tasks []*veeamrestapi.TaskType
	QuickBackup []*VM
	BackupServer string
}
type WorkStatus struct {
	result string
	success bool
	completed bool
	idstring string
	redirect string
}
type Work struct {
	Method string
	Url string
	Post string
	ResultTask bool
}
type SelfService struct {
	rest * veeamrestapi.VeeamRestServer
	l sync.RWMutex
	maintemplate * template.Template
	files *map[string][]byte
	HierarchyRootCache *[]*veeamrestapi.HierarchyRootEntityType
	VMs *[]*VM
	CacheTime time.Time
	WorkListMutex sync.RWMutex
	WorkList *map[string]*WorkStatus
	WorkListCounter int
}

func (h *SelfService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path,"/")
	
	action := strings.Split(path,"/")
	
	
	if path ==  "favicon.ico" {
		fmt.Fprintf(w,"all your base are ignored by us")
	} else if len(path) > 3 && path[:3] == "___" {
		if len(path) > 15 &&  path[0:14] == "___staticfiles"  {
			w.Header().Add("Content-Type","image/png")
			w.Write((*h.files)[path[14:]])
		} else if  path == "___showvmcache" {
			var buf bytes.Buffer
			h.l.RLock()
			defer h.l.RUnlock()
			
			buf.WriteString("\n<table>")
			for _,vm := range (*h.VMs) {
					buf.WriteString("\n  <tr><td>")
					buf.WriteString(vm.Name)
					buf.WriteString("</td><td>")
					buf.WriteString(vm.UID)
					if vm.Root != nil {
						buf.WriteString("</td><td>")
						buf.WriteString(vm.Root.Name)
						buf.WriteString("</td><td>")
						buf.WriteString(vm.Root.HierarchyRootId)
						buf.WriteString("</td><td>")
						buf.WriteString(vm.Root.UID.S())
					} else {
						buf.WriteString("</td><td colspan='3'>")
					}
					buf.WriteString("</td></tr>")
			}
			buf.WriteString("\n</table>")
			h.maintemplate.Execute(w,MainTemplate{Title:"Cache",Customer:action[0],Content:buf.String()})
		} else if path == "___tasks" {
			h.l.Lock()
			defer h.l.Unlock()
			rest := h.rest
			
			sessionvalid := false
			var err error
			
			if rest.RestSvcSessionId != "" { sessionvalid = rest.HeartbeatSession()}
			if !sessionvalid { err = rest.Authenticate()} else { log.Println("Reusing Session")}	
			
			q,err := h.rest.GetTasks()
			if err == nil {
				h.maintemplate.Execute(w,MainTemplate{Title:"Tasks",Tasks:q.Task})
			} else {
				h.maintemplate.Execute(w,MainTemplate{Title:"Error",Content:(fmt.Sprintf("Error %s",err))})
			}

		}
	} else if len(action) > 2 && action[1] == "workprogress" {
		h.WorkListMutex.RLock()
		defer h.WorkListMutex.RUnlock()
		
		var dots bytes.Buffer
		dots.WriteString("Working on it")
		waited := "/1"
		if len(action) > 3 {
			if i, err := strconv.ParseInt(action[3], 10, 64); err == nil {
				for j := int64(0);j < i;j++ {
					dots.WriteRune('.')
				}
				waited = fmt.Sprintf("/%d",(i+1))
			}
		}
		workstatus := (*h.WorkList)[action[2]]
		if  workstatus != nil {
			if workstatus.completed {
				jobmsg := "Job executed succesfully"
				if !workstatus.success {
					jobmsg = "Job didnt't execute succesfully. Message returned by worker : "+workstatus.result
				}
				h.maintemplate.Execute(w,MainTemplate{Title:"Execution result",Customer:action[0],Content:jobmsg,Autorefresh:"2;URL="+workstatus.redirect})
			} else {
				h.maintemplate.Execute(w,MainTemplate{Title:"Working on it",Customer:action[0],Content:dots.String(),Autorefresh:"2;URL=/"+action[0]+"/workprogress/"+workstatus.idstring+waited})
			}
		} else {
			h.maintemplate.Execute(w,MainTemplate{Title:"Error",Customer:action[0],Content:"Error retrieving job",Autorefresh:"3;URL=/"+action[0]})
		}
		
		
		
	}  else {		
			prefix := strings.TrimSpace(action[0])
			log.Printf("New Call : %s",path)
			if len(path) > 1 {
				h.l.Lock()
				defer h.l.Unlock()
				rest := h.rest
				
				sessionvalid := false
				var err error
				
				if rest.RestSvcSessionId != "" { sessionvalid = rest.HeartbeatSession()}
				if !sessionvalid { err = rest.Authenticate()} else { log.Println("Reusing Session")}
				
				
				if len(action) == 1 {

					if err == nil {
						q,err := rest.GetJobs()
						b,err2 := rest.GetBackups()
						if err == nil && err2 == nil {
							httpjobs := []*HttpJob{}
							for _,job := range q.Job {
								if len(job.Name) >= len(prefix) && strings.EqualFold(prefix,job.Name[0:len(prefix)]) {
									httpjobs = append(httpjobs,&HttpJob{Name:job.Name,UID:job.UID.Split()[2],JobType:job.JobType})
								} 
							}
							httpbackups := []*HttpBackup{}
							for _,backup := range b.Backup {
								if len(backup.Name) >= len(prefix) && strings.EqualFold(prefix,backup.Name[0:len(prefix)]) {
									httpbackups = append(httpbackups,&HttpBackup{backup.Name,backup.UID.Split()[2]})
								}
							}
							h.maintemplate.Execute(w,MainTemplate{Title:"Jobs",Customer:prefix,Jobs:httpjobs,Backups:httpbackups})
						} 

						if err != nil || err2 != nil {
							str := ""
							if err != nil  { str += err.Error()}
							if err2 != nil { str += err2.Error() }
							h.maintemplate.Execute(w,MainTemplate{Title:"Error",Content:(fmt.Sprintf("Error %s",str))})
						}
						
					}
				}  else if len(action) > 2 {
					switch action[1] {
						case "startjob":
							q,err := rest.GetJob(action[2])
							if err == nil {
								if links := veeamrestapi.FindLinkByRel(q,"start");len(links) > 0 {
									task := veeamrestapi.TaskType{}
									err = rest.GenericPostRequest(links[0].Href.String(),"",&task)
									if err == nil {
										h.maintemplate.Execute(w,MainTemplate{Title:"Job Started",Customer:prefix,Autorefresh:"2;URL=/"+prefix+"/showjobsessions/"+action[2],Content:(fmt.Sprintf("Succesfully started job %s",q.Name))})
									}
								}
							} else {
								h.maintemplate.Execute(w,MainTemplate{Title:"Error",Content:(fmt.Sprintf("Error %s",err))})
							}
						break;
						case "stopjob":
							q,err := rest.GetJob(action[2])
							if err == nil {
								if links := veeamrestapi.FindLinkByRel(q,"stop");len(links) > 0 {
									task := veeamrestapi.TaskType{}
									err = rest.GenericPostRequest(links[0].Href.String(),"",&task)
									if err == nil {
										h.maintemplate.Execute(w,MainTemplate{Title:"Job Started",Customer:prefix,Autorefresh:"2;URL=/"+prefix+"/showjobsessions/"+action[2],Content:(fmt.Sprintf("Succesfully stopped job %s",q.Name))})
									}
								}
							} else {
								h.maintemplate.Execute(w,MainTemplate{Title:"Error",Content:(fmt.Sprintf("Error %s",err))})
							}
						break;
						case "editjob":
							q,err := rest.GetJob(action[2])
							if err == nil {
								vms := []*VM{}
								if q.JobInfo.BackupJobInfo  != nil {
									for _,vm := range q.JobInfo.BackupJobInfo.Includes.ObjectInJob  {
										uid := vm.HierarchyObjRef
										if uid.Split()[1] == "Vm" {
											var root *veeamrestapi.HierarchyRootEntityType
											for _,vmq := range (*h.VMs) {
												if uid.String()== vmq.UID {
													root = vmq.Root
												}
											}
											vms = append(vms,&VM{Added:true,Name:vm.Name,UID:uid.String(),Root:root})
										}
									}
									
									for _,vmq := range (*h.VMs) {
										if len(vmq.Name) >= len(prefix) && strings.EqualFold(vmq.Name[:len(prefix)],prefix) {
											already := false
											for _,t := range vms {
												if t.UID == vmq.UID {
													already = true
												}
											}
											if !already {
												vms = append(vms,&VM{Added:false,Name:vmq.Name,UID:vmq.UID,Root:vmq.Root})
											}
										}
									}
								}
								job := HttpJob{q.Name,q.UID.Split()[2],q.JobType,&q,vms}
								h.maintemplate.Execute(w,MainTemplate{Title:"Edit",Customer:prefix,Edit:&job})
								
								
									
							} else {
								h.maintemplate.Execute(w,MainTemplate{Title:"Error",Content:(fmt.Sprintf("Error %s",err))})
							}
						break;
						case "editjobexec":
							q,err := rest.GetJob(action[2])
							
							if err == nil {
								r.ParseForm()
								
								
								
								addVM := []*VM{}
								removeVM := []*VM{}
								
								if len(r.PostForm["___vmlist"]) > 0 {
									for _,vmuid := range r.PostForm["___vmlist"] {
										if len(r.PostForm[vmuid])> 0 && r.PostForm[vmuid][0] == "on" {
											//log.Println("On "+vmuid)
											added := false
											for _,vm := range q.JobInfo.BackupJobInfo.Includes.ObjectInJob  {
												if vm.HierarchyObjRef.String() == vmuid {
													added = true
												}
											}
											
											if !added {
												addVM = append(addVM,&VM{Name:r.PostForm[vmuid+"_vmname"][0],UID:vmuid})
												log.Print("Adding "+vmuid)
											}
											
										} else {
											for _,vm := range q.JobInfo.BackupJobInfo.Includes.ObjectInJob {
												if vm.HierarchyObjRef.String() == vmuid {
													removeVM = append(removeVM,&VM{Name:r.PostForm[vmuid+"_vmname"][0],UID:vmuid})
												}
											}
										}
									}
								}
								
								
								ScheduleConfigured := false
								if len(r.PostForm["ScheduleConfigured"]) > 0 && r.PostForm["ScheduleConfigured"][0] == "on" {
									ScheduleConfigured = true
								} 
								
								
								
								
								h.WorkListMutex.Lock()
								defer h.WorkListMutex.Unlock()
								h.WorkListCounter = h.WorkListCounter+1
								idstring := fmt.Sprintf("%s-%d",prefix,h.WorkListCounter)

								status := WorkStatus{result:"Executing",success:false,completed:false,idstring:idstring,redirect:"/"+prefix+"/editjob/"+q.UID.Split()[2]}
								(*h.WorkList)[idstring] = &status
								
							
								
								h.maintemplate.Execute(w,MainTemplate{Title:"Edit",Customer:prefix,Content:"Job edit submited",Autorefresh:"1;URL=/"+prefix+"/workprogress/"+idstring})
								
								
								go func(h *SelfService,workstatus * WorkStatus,jobid string,addVMP *[]*VM,removeVMP *[]*VM,ScheduleConfigured bool) {
									allsuccess := true
									var errortext bytes.Buffer
									errortext.WriteString("Partial edit failure : ")
									
									h.l.Lock()
									defer h.l.Unlock()
									
									naptime := 4000* time.Millisecond
									shortnaptime := 1000*time.Millisecond
									refreshtasktime := 100* time.Millisecond
									
									var successaction bool;
									
									for _,vm := range (*addVMP) {
										j,err := h.rest.GetJob(jobid) 
										
										successaction := false;
										if err == nil {
											spec := veeamrestapi.NewCreateObjectInJobSpecType()
											spec.HierarchyObjRef =  veeamrestapi.HierarchyObjRefType(vm.UID)
											spec.HierarchyObjName = vm.Name
											spec.Order = len(j.JobInfo.BackupJobInfo.Includes.ObjectInJob)
											
											xmlpost,err := xml.MarshalIndent(spec," "," ")
											if err == nil {
												task := veeamrestapi.TaskType{}
												err = h.rest.GenericRequest("jobs/"+jobid+"/includes",&task,"POST",xml.Header+string(xmlpost))
												if err == nil {
													refreshok := true
													for !strings.EqualFold(task.State,"Finished") && refreshok {
														time.Sleep(refreshtasktime)
														task,err = h.rest.GetTask(task.TaskId)
														if err != nil {
															refreshok = false
														}
													}
													if task.Result != nil && !task.Result.Success {
														allsuccess = false
														errortext.WriteString("-- Task failed for adding VM ");errortext.WriteString(vm.Name)
														log.Printf("Error : "+task.Result.Message)
													} else if task.Result == nil {
														allsuccess = false
														errortext.WriteString("-- Unable to wait succesfully ");errortext.WriteString(vm.Name)
													} else {
														log.Printf(task.Result.Message)
														successaction = true
													}
												} else {
													allsuccess = false
													errortext.WriteString("-- Unsuccesful adding VM ");errortext.WriteString(vm.Name)
												}
											} else {
												allsuccess = false
												errortext.WriteString("-- Could not marshal for adding VM ");errortext.WriteString(vm.Name)
											}
											
										} else {
											allsuccess = false
											errortext.WriteString("-- Could not get job for adding VM ");errortext.WriteString(vm.Name)
										}
										if successaction {
											//shortnaptime
											added := false
											retry := 10
											
											for retry > 0 && !added {
												j,err := h.rest.GetJob(jobid)
												if err == nil {
													found := false
													for _,ovm := range j.JobInfo.BackupJobInfo.Includes.ObjectInJob {
														if ovm.HierarchyObjRef.String() == vm.UID {
															 found = true
														}
													}
													added = found
												}
												retry = retry - 1
												time.Sleep(shortnaptime)
												log.Println("zzzzzZZZZZZzzzzzzzzz: short sleep")
											}
											if !added {
												errortext.WriteString("-- Unsuccesful adding VM ");errortext.WriteString(vm.Name)
												log.Print("Waited for VM to be added but it didn't happen")
											} else {
												log.Print("Job is really updated")
											}
										} else {
											log.Println("zzzzzZZZZZZzzzzzzzzz: long sleep")
											time.Sleep(naptime)
										}
									}
									
									for _,vm := range (*removeVMP) {
										j,err := h.rest.GetJob(jobid)
										
										successaction := false;
										if err == nil {
											link := ""
											for _,ovm := range j.JobInfo.BackupJobInfo.Includes.ObjectInJob {
												if ovm.HierarchyObjRef.String() == vm.UID {
													link = ovm.Href.String()
												}
											}
											
											if link != "" {
												task := veeamrestapi.TaskType{}
												err = h.rest.GenericRequest(link,&task,"DELETE","")
												if err == nil {
													refreshok := true
													for !strings.EqualFold(task.State,"Finished") && refreshok {
														time.Sleep(refreshtasktime)
														task,err = h.rest.GetTask(task.TaskId)
														if err != nil {
															refreshok = false
														}
													}
													if task.Result != nil && !task.Result.Success {
														allsuccess = false
														errortext.WriteString("-- Task failed for deleting VM ");errortext.WriteString(vm.Name)
														log.Printf("Error : "+task.Result.Message)
													}  else if task.Result == nil {
														allsuccess = false
														errortext.WriteString("-- Unable to wait succesfully ");errortext.WriteString(vm.Name)
													} else {
														log.Printf(task.Result.Message)
														successaction = true
													}
												} else {
													allsuccess = false
													errortext.WriteString("-- Unsuccesful deleteing VM ");errortext.WriteString(vm.Name)
												}
											} else {
												allsuccess = false
												errortext.WriteString("-- Could not find ref for removing VM ");errortext.WriteString(vm.Name)
											}
										} else {
											allsuccess = false
											errortext.WriteString("-- Could not get job for removing VM ");errortext.WriteString(vm.Name)
										}
										
										if successaction {
											//shortnaptime
											deleted := false
											retry := 10
											
											for retry > 0 && !deleted {
												j,err := h.rest.GetJob(jobid)
												if err == nil {
													found := false
													for _,ovm := range j.JobInfo.BackupJobInfo.Includes.ObjectInJob {
														if ovm.HierarchyObjRef.String() == vm.UID {
															 found = true
														}
													}
													deleted = !found
												}
												retry = retry - 1
												time.Sleep(shortnaptime)
												log.Println("zzzzzZZZZZZzzzzzzzzz: short sleep")
											}
											if !deleted {
												errortext.WriteString("-- Unsuccesful deleteing VM ");errortext.WriteString(vm.Name)
												log.Print("Waited for VM to be deleted but it didn't happen")
											} else {
												log.Print("Job is really updated")
											}
										} else {
											log.Println("zzzzzZZZZZZzzzzzzzzz: long sleep")
											time.Sleep(naptime)
										}
									}										
									
									successaction = false
									
									j,err := h.rest.GetJob(jobid) 
									je := NewJobEditType()
									je.Name = j.Name
									je.UID = j.UID.String()
									je.ScheduleConfigured = ScheduleConfigured
									
									xmlpost,err := xml.MarshalIndent(je," "," ")
									if err == nil {
										task := veeamrestapi.TaskType{}
										err = h.rest.GenericRequest("jobs/"+jobid,&task,"PUT",xml.Header+string(xmlpost))
										if err == nil {
											refreshok := true
											for !strings.EqualFold(task.State,"Finished") && refreshok {
												time.Sleep(refreshtasktime)
												task,err = h.rest.GetTask(task.TaskId)
												if err != nil {
													refreshok = false
												}
											}
											if task.Result != nil && !task.Result.Success {
												allsuccess = false
												errortext.WriteString("-- Task failed for  en/disabling schedule ")
												log.Printf("Error : "+task.Result.Message)
											} else if task.Result == nil {
												allsuccess = false
												errortext.WriteString("-- Unable to wait succesfully ")
											} else {
												log.Printf(task.Result.Message)
												successaction = true
											}
										} else {
											allsuccess = false
											errortext.WriteString("-- Unsuccesful en/disabling schedule ")
										}
										
									} else {
										allsuccess = false
										errortext.WriteString("-- Unsuccesful marshaling for en/disabling schedule")
									}
									if successaction {
										mod := false
										retry := 10
										
										for retry > 0 && !mod {
											j,err := h.rest.GetJob(jobid)
											if err == nil {
												if j.ScheduleConfigured ==  ScheduleConfigured {
													mod = true
												}
											}
											retry = retry - 1
											time.Sleep(shortnaptime)
											log.Println("zzzzzZZZZZZzzzzzzzzz: short sleep")
										}
										if !mod {
											errortext.WriteString("-- Unsuccesful modding job ")
											log.Print("Waited for job to be edited but it did not happen")
										} else {
											log.Print("Job is really updated")
										}
									} else {
										log.Println("zzzzzZZZZZZzzzzzzzzz: long sleep")
										time.Sleep(naptime)
									}
									
									h.WorkListMutex.Lock()
									defer h.WorkListMutex.Unlock()
									status.completed = true
									if allsuccess {
										status.success = true
										status.result = "All ok"
									}  else {
										status.result = errortext.String()
									}
									
								} (h,&status,q.UID.Split()[2],&addVM,&removeVM,ScheduleConfigured)

							} else {
								h.maintemplate.Execute(w,MainTemplate{Title:"Error",Content:(fmt.Sprintf("Error %s",err))})
							}
						break;	
						case "showbackuppoints":
							rps := veeamrestapi.RestorePointEntityListType{}
							err = rest.GenericGetRequest("backups/"+action[2]+"/restorePoints?format=Entity",&rps)
							if err == nil {
								hrps := []*HttpBackupRestorePoint {}
								for _,rp := range rps.RestorePoint {
									hrps = append(hrps,&HttpBackupRestorePoint{rp.Name,rp.UID.Split()[2]})
								}
								h.maintemplate.Execute(w,MainTemplate{Title:"Restorepoints",Customer:prefix,BackupRestorePoints:hrps})
							} else {
								h.maintemplate.Execute(w,MainTemplate{Title:"Error",Content:(fmt.Sprintf("Error %s",err))})
							}
						break;
						case "showbackuppointvm":
							os := veeamrestapi.VmRestorePointEntityListType{}
							//https://localhost:9398/api/restorePoints/0766c4c2-79da-4a38-b554-7e33be04b5b7/vmRestorePoints
							/*
								<VmRestorePoint Type="VmRestorePoint" Href="https://localhost:9398/api/vmRestorePoints/3e9ac9c0-dd79-4d3e-82e9-df78ef437034?format=Entity" Name="timocbt@2016-03-10 19:28:21" UID="urn:veeam:VmRestorePoint:3e9ac9c0-dd79-4d3e-82e9-df78ef437034">
								<Links>
								  <Link Href="https://localhost:9398/api/vmRestorePoints/3e9ac9c0-dd79-4d3e-82e9-df78ef437034?action=restore" Rel="Restore"/>
								  <Link Href="https://localhost:9398/api/backupServers/1c47570a-27ac-4aab-9a84-8d88df1b86c8" Name="localhost" Type="BackupServerReference" Rel="Up"/>
								  <Link Href="https://localhost:9398/api/restorePoints/0766c4c2-79da-4a38-b554-7e33be04b5b7" Name="Mar 10 2016  7:27PM" Type="RestorePointReference" Rel="Up"/>
								  <Link Href="https://localhost:9398/api/vmRestorePoints/3e9ac9c0-dd79-4d3e-82e9-df78ef437034" Name="timocbt@2016-03-10 19:28:21" Type="VmRestorePointReference" Rel="Alternate"/>
								  <Link Href="https://localhost:9398/api/vmRestorePoints/3e9ac9c0-dd79-4d3e-82e9-df78ef437034/mounts" Type="VmRestorePointMountList" Rel="Down"/>
								  <Link Href="https://localhost:9398/api/vmRestorePoints/3e9ac9c0-dd79-4d3e-82e9-df78ef437034/mounts" Type="VmRestorePointMount" Rel="Create"/>
								</Links>
								<CreationTimeUTC>2016-03-10T19:28:21Z</CreationTimeUTC>
								<Algorithm>Full</Algorithm>
								<PointType>Full</PointType>
								<HierarchyObjRef>urn:VMware:Vm:f49942ee-2098-4a6c-af36-be7e1c82458a.vm-14765</HierarchyObjRef>
							*/
							err = rest.GenericGetRequest("restorePoints/"+action[2]+"/vmRestorePoints?format=Entity",&os)
							if err == nil {
								backupVMRestorePoints  := []*HttpBackupVMRestorePoint{}
								for _,vrp := range os.VmRestorePoint {
									namesplit := strings.Split(vrp.Name,"@")
									if len(namesplit) > 1 {
										backupVMRestorePoints = append(backupVMRestorePoints,&HttpBackupVMRestorePoint{Name:namesplit[0],CreateUTC:vrp.CreationTimeUTC.TimeString(),Algorithm:vrp.Algorithm,PointType:vrp.PointType,UID:vrp.UID.Split()[2]})
									} else {
										backupVMRestorePoints = append(backupVMRestorePoints,&HttpBackupVMRestorePoint{Name:vrp.Name,CreateUTC:vrp.CreationTimeUTC.TimeString(),Algorithm:vrp.Algorithm,PointType:vrp.PointType,UID:vrp.UID.Split()[2]})
									}
								}
								h.maintemplate.Execute(w,MainTemplate{Title:"Restorepoints",Customer:prefix,BackupVMRestorePoints:backupVMRestorePoints})
							} else {
								h.maintemplate.Execute(w,MainTemplate{Title:"Error",Content:(fmt.Sprintf("Error %s",err))})
							}
						break;
						case "fullvmrestore":
							h.maintemplate.Execute(w,MainTemplate{Title:"Are you sure",Customer:prefix,Content:(fmt.Sprintf("Are you sure you sure? <a href='/%s/fullvmrestoreconfirm/%s'>YES</a> /  <a href='/%s'>NO</a>",prefix,action[2],prefix))})
						break;
						case "fullvmrestoreconfirm":
							//<Link Href="https://localhost:9398/api/vmRestorePoints/070ce19e-0a90-49a4-97f9-0082fdf138fa?action=restore" Rel="Restore"/>
							task := veeamrestapi.TaskType{}
							err = rest.GenericPostRequest("vmRestorePoints/"+action[2]+"?action=restore","",&task)
							if err == nil {
								h.WorkListMutex.Lock()
								defer h.WorkListMutex.Unlock()
								h.WorkListCounter = h.WorkListCounter+1
								idstring := fmt.Sprintf("%s-%d",prefix,h.WorkListCounter)

								status := WorkStatus{result:"Executing",success:false,completed:false,idstring:idstring,redirect:"/"+prefix}
								(*h.WorkList)[idstring] = &status
								
							
								
								h.maintemplate.Execute(w,MainTemplate{Title:"Job started",Customer:prefix,Content:"Job start submited",Autorefresh:"1;URL=/"+prefix+"/workprogress/"+idstring})
			
								go func(h *SelfService,workstatus * WorkStatus,taskid string) {									
									h.l.Lock()
									defer h.l.Unlock()
								
									task,err := h.rest.GetTask(taskid)
									for err == nil && !strings.EqualFold(task.State,"Finished") {
										time.Sleep(500*time.Millisecond)
										task,err = h.rest.GetTask(taskid)
									}
									

									h.WorkListMutex.Lock()
									defer h.WorkListMutex.Unlock()
									status.completed = true
									
									if err == nil {
										if task.Result != nil && !task.Result.Success {
											status.result = "Fail : "+task.Result.Message
										} else if task.Result == nil {
											status.result = "Fail but no message"
										} else {
											status.success = true
											links := veeamrestapi.FindLinkByType(task,"RestoreSession")
											if len(links) > 0 {
												parts := strings.Split(strings.Trim(links[0].Href.S(),"?format=Entity"),"/")
												status.redirect = "/"+prefix+"/showrestoresession/"+parts[(len(parts)-1)]
											} else {
												status.result = "Session started but could not find follow up link"
											}
											
											
										}
									} else {
										status.result = "Fail but internal error"
									}
								} (h,&status,task.TaskId)
							} else {
								h.maintemplate.Execute(w,MainTemplate{Title:"Error",Content:(fmt.Sprintf("Error %s",err))})
							}
						break
						case "showrestoresession":
							session,err := rest.GetRestoreSession(action[2])
							if err == nil {
								rs := []*HttpRestoreSession{}
								rs = append(rs,&HttpRestoreSession{SessionUID:session.UID.Split()[2],Name:session.Name,JobType:session.JobType,CreateUTC:session.CreationTimeUTC.TimeString(),EndUTC:session.EndTimeUTC.TimeString(),State:session.State,Result:session.Result,Progress:session.Progress})
								h.maintemplate.Execute(w,MainTemplate{Title:"Result",Customer:prefix,RestoreSessions:rs,Autorefresh:"5"})
							} else {
								h.maintemplate.Execute(w,MainTemplate{Title:"Error",Content:(fmt.Sprintf("Error %s",err))})
							}
						break;
						case "showrestoresessions":
							sessions,err := rest.GetRestoreSessions()
							if err == nil {
								rs := []*HttpRestoreSession{}
								for _,session := range sessions.RestoreSession {
									if len(session.Name) > len(prefix) && strings.EqualFold(session.Name[:len(prefix)],prefix) {
										rs = append(rs,&HttpRestoreSession{SessionUID:session.UID.Split()[2],Name:session.Name,JobType:session.JobType,CreateUTC:session.CreationTimeUTC.TimeString(),EndUTC:session.EndTimeUTC.TimeString(),State:session.State,Result:session.Result,Progress:session.Progress})
									}
								}
								h.maintemplate.Execute(w,MainTemplate{Title:"Result",Customer:prefix,RestoreSessions:rs,Autorefresh:"5"})
							} else {
								h.maintemplate.Execute(w,MainTemplate{Title:"Error",Content:(fmt.Sprintf("Error %s",err))})
							}
						break;
						case "quickbackup":
							q,err := rest.GetJob(action[2])
							if err == nil {
								links := veeamrestapi.FindLinkByType(q,"BackupServerReference")
								if len(links) > 0 {
									BackupServerLink := strings.Split(strings.TrimLeft(links[0].Href.S(),"?format=Entity"),"/")
									
									vms := []*VM{}
									if q.JobInfo.BackupJobInfo  != nil {
										for _,vm := range q.JobInfo.BackupJobInfo.Includes.ObjectInJob  {
											uid := vm.HierarchyObjRef
											if uid.Split()[1] == "Vm" {
												var root *veeamrestapi.HierarchyRootEntityType
												for _,vmq := range (*h.VMs) {
													if uid.String()== vmq.UID {
														root = vmq.Root
													}
												}
												vms = append(vms,&VM{Added:true,Name:vm.Name,UID:uid.String(),Root:root})
											}
										}
									}
									h.maintemplate.Execute(w,MainTemplate{Title:"Quick Backup",Customer:prefix,QuickBackup:vms,BackupServer:BackupServerLink[len(BackupServerLink)-1]})
								} else {
									h.maintemplate.Execute(w,MainTemplate{Title:"Error",Content:"Can not find backup server link"})
								}
									
							} else {
								h.maintemplate.Execute(w,MainTemplate{Title:"Error",Content:(fmt.Sprintf("Error %s",err))})
							}
							
						break;
						case "quickbackupexec":
							if len(action) > 3 {
								spec := veeamrestapi.NewQuickBackupStartupSpecType()
								spec.VmRef =  veeamrestapi.HierarchyObjRefType(action[3])
								
								xmlpost,err := xml.MarshalIndent(spec," "," ")
								if err == nil {
									task := veeamrestapi.TaskType{}
									err = h.rest.GenericRequest("/backupServers/"+action[2]+"?action=quickbackup",&task,"POST",xml.Header+string(xmlpost))
									if err == nil {
										h.WorkListMutex.Lock()
										defer h.WorkListMutex.Unlock()
										h.WorkListCounter = h.WorkListCounter+1
										idstring := fmt.Sprintf("%s-%d",prefix,h.WorkListCounter)

										status := WorkStatus{result:"Executing",success:false,completed:false,idstring:idstring,redirect:"/"+prefix}
										(*h.WorkList)[idstring] = &status
										
										h.maintemplate.Execute(w,MainTemplate{Title:"Job started",Customer:prefix,Content:"Job start submited",Autorefresh:"1;URL=/"+prefix+"/workprogress/"+idstring})
										
										
										go func(h *SelfService,workstatus * WorkStatus,taskid string) {									
											h.l.Lock()
											defer h.l.Unlock()
										
											task,err := h.rest.GetTask(taskid)
											for err == nil && !strings.EqualFold(task.State,"Finished") {
												time.Sleep(500*time.Millisecond)
												task,err = h.rest.GetTask(taskid)
											}
											

											h.WorkListMutex.Lock()
											defer h.WorkListMutex.Unlock()
											status.completed = true
											
											if err == nil {
												if task.Result != nil && !task.Result.Success {
													status.result = "Fail : "+task.Result.Message
												} else if task.Result == nil {
													status.result = "Fail but no message"
												} else {
													status.success = true
													links := veeamrestapi.FindLinkByType(task,"BackupJobSession")
													log.Printf("%v",links)
													if len(links) > 0 {
														parts := strings.Split(strings.Trim(links[0].Href.S(),"?format=Entity"),"/")
														status.redirect = "/"+prefix+"/showbackupjobsession/"+parts[(len(parts)-1)]
													} else {
														status.result = "Session started but could not find follow up link"
													}	
												}
											} else {
												status.result = "Fail but internal error"
											}
										} (h,&status,task.TaskId)

									} else {
										h.maintemplate.Execute(w,MainTemplate{Title:"Error",Content:(fmt.Sprintf("Error %s",err))})
									} 
								} else {
									h.maintemplate.Execute(w,MainTemplate{Title:"Error",Content:(fmt.Sprintf("Marshal Error %s",err))})
								}

							} else {
								h.maintemplate.Execute(w,MainTemplate{Title:"Error",Content:"You need to supply backup and vm reference for quickbackup"})
							}
						break;
						case "showbackupjobsession":
							session,err := rest.GetBackupJobSession(action[2])
							if err == nil {
								sessions := []*HttpJobSession{}
								if len(session.JobName) >= len(prefix) && strings.EqualFold(prefix,session.JobName[0:len(prefix)]) {
									sessions = append(sessions,&HttpJobSession{SessionUID:session.UID.Split()[2],JobUID:session.JobUid.Split()[2],JobName:session.JobName,JobType:session.JobType,CreateUTC:session.CreationTimeUTC.TimeString(),EndUTC:session.EndTimeUTC.TimeString(),State:session.State,Result:session.Result,Progress:session.Progress,IsRetry:session.IsRetry})
								}
								h.maintemplate.Execute(w,MainTemplate{Title:"Result",Customer:prefix,Sessions:sessions,Autorefresh:"5"})
							} else {
								h.maintemplate.Execute(w,MainTemplate{Title:"Error",Content:(fmt.Sprintf("Error %s",err))})
							}
						break
						case "showjobsessions":
							//https://localhost:9398/web/#/api/query?type=BackupJobSession&format=Entities&sortDesc=name&pageSize=10&page=1&filter=(JobUID==f7d731be-53f7-40ca-9c45-cbdaf29e2d99)
							qget := url.Values{}
							qget.Add("type","BackupJobSession")
							qget.Add("format","Entities")
							qget.Add("sortDesc","name")
							qget.Add("pageSize","10")
							qget.Add("page","1")
							qget.Add("filter","(JobUID=="+action[2]+")")
							
							q,err := rest.GetQueryResult(qget)
							if err == nil {
								/*Entities : BackupJobSessions : BackupJobSession
									<JobUid>urn:veeam:Job:f8568011-3542-470f-ab38-46b27e427682</JobUid>
									<JobName>Backup Job 8</JobName>
									<JobType>Backup</JobType>
									<CreationTimeUTC>2016-01-04T18:15:43Z</CreationTimeUTC>
									<EndTimeUTC>2016-01-04T18:17:05Z</EndTimeUTC>
									<State>Stopped</State>
									<Result>Success</Result>
									<Progress>100</Progress>
									<IsRetry>false</IsRetry>
								*/
								sessions := []*HttpJobSession{}
								for _,session := range q.Entities.BackupJobSessions.BackupJobSession {
									if len(session.JobName) >= len(prefix) && strings.EqualFold(prefix,session.JobName[0:len(prefix)]) {
										sessions = append(sessions,&HttpJobSession{SessionUID:session.UID.Split()[2],JobUID:session.JobUid.Split()[2],JobName:session.JobName,JobType:session.JobType,CreateUTC:session.CreationTimeUTC.TimeString(),EndUTC:session.EndTimeUTC.TimeString(),State:session.State,Result:session.Result,Progress:session.Progress,IsRetry:session.IsRetry})
									}
								}
								h.maintemplate.Execute(w,MainTemplate{Title:"Result",Customer:prefix,Sessions:sessions,Autorefresh:"5"})
							} else {
								h.maintemplate.Execute(w,MainTemplate{Title:"Error",Content:(fmt.Sprintf("Error %s",err))})
							}
						break;
						case "showjobtasksessions":
							//https://localhost:9398/web/#/api/query?type=BackupTaskSession&format=Entities&sortDesc=name&filter=(JobSessionUid==52a15e33-f210-462c-b103-fb1d6d5dd789)
							qget := url.Values{}
							qget.Add("type","BackupTaskSession")
							qget.Add("format","Entities")
							qget.Add("sortDesc","name")
							qget.Add("filter","(JobSessionUID=="+action[2]+")")
							
							q,err := rest.GetQueryResult(qget)
							if err == nil {
								/*Entities : BackupTaskSessions : BackupTaskSession
									<BackupTaskSession Type="BackupTaskSession" Href="https://localhost:9398/api/backupTaskSessions/605fcf67-c9a2-43e5-bc66-0d08ef8c25bf?format=Entity" Name="timocbt@2016-03-10 18:23:13" UID="urn:veeam:BackupTaskSession:605fcf67-c9a2-43e5-bc66-0d08ef8c25bf" VmDisplayName="timocbt">
									<JobSessionUid>urn:veeam:BackupJobSession:52a15e33-f210-462c-b103-fb1d6d5dd789</JobSessionUid>
									<CreationTimeUTC>2016-03-10T18:23:13Z</CreationTimeUTC>
									<EndTimeUTC>2016-03-10T18:25:00Z</EndTimeUTC>
									<State>Completed</State>
									<Result>Success</Result>
									<Reason/>
									<TotalSize>9743368192</TotalSize>
								*/
								tasks := []*HttpJobTaskSession{}
								for _,session := range q.Entities.BackupTaskSessions.BackupTaskSession {
									tasks = append(tasks,&HttpJobTaskSession{JobSessionUID:session.JobSessionUid.Split()[2],VmName:session.VmDisplayName,JobTaskSessionUID:session.UID.Split()[2],CreateUTC:session.CreationTimeUTC.TimeString(),EndUTC:session.EndTimeUTC.TimeString(),State:session.State,Result:session.Result,Reason:session.Reason,TotalSize:session.TotalSize})					
								}	
								h.maintemplate.Execute(w,MainTemplate{Title:"Result",Customer:prefix,JobTaskSessions:tasks,Autorefresh:"5"})
							} else {
								h.maintemplate.Execute(w,MainTemplate{Title:"Error",Content:(fmt.Sprintf("Error %s",err))})
							}
						break;
						default:
							h.maintemplate.Execute(w,MainTemplate{Title:"Error",Content:"Unknown page"})
						break
						
					}
				} else {
					h.maintemplate.Execute(w,MainTemplate{Title:"Error",Content:"Please supply /prefix for your tenant"})
				}
				
				
				if err != nil {
					h.maintemplate.Execute(w,MainTemplate{Title:"Error",Content:(fmt.Sprintf("Error %s",err))})
					
				}
			} else {
				h.maintemplate.Execute(w,MainTemplate{Title:"Error",Content:"Please supply 1 or 3 arguments e.g /prefix/action/id"})
			}
	}
}

func (h *SelfService) setTemplates() (error) {
	var err error
	tpl := `{{$customer :=  .Customer}}<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{$customer}} {{.Title}}</title>
		{{if .Autorefresh}}<meta http-equiv="refresh" content="{{.Autorefresh}}"/>{{end}}
		<style>
			.header {
				background-color:#33cc33;
				color:#ffffff;
				width:100hv;
				height:50px;
				font-size:40px;
				font-family:verdana;
				margin: 0px 0px 0px 0px;
				padding: 10px 10px 10px 10px;
			}
			.content {
				width:100hv;
				font-size:15px;
				font-family:verdana;
				margin: 0px 0px 0px 0px;
				padding: 5px 10px 5px 10px;
				color:#333;
			}
			body {
				margin: 0px 0px 0px 0px;
				padding: 0px 0px 0px 0px;
			}
			td {
				font-size:12px;
			}
			.jobtable td.iconwidth,.sessiontable td.iconwidth {
				margin: 0px 0px 0px 0px;
				padding: 0px 1px 0px 1px;
				width:20px;
			}
			.jobtable td {
				width: 100px;
			}
			.jobtable td.namew {
				width: 300px;
			}
			.sessiontable td {
				width: 150px;
			}
			.sessiontable td.progress {
				width: 120px;
				height: 10px;
				padding: 0px 10px 0px 10px;
			}
			.sessiontable td.namew {
				width: 300px;
			}
			a,a:hover,a:link,a:active,a:visited {
				color:#333;
			}
			div.progress {
				display: inline-block;
				height:10px;
				
				margin: 0px 0px 0px 0px;
			}
			
			div.green { background-color:#33cc33; }
			div.red { background-color:#cc3333; }
			div.gray { background-color:#333333; }
		</style>
	</head>
	<body>
		<div class="header">Self Service Demo</div>
		<div class="content">
		{{if .Content }} {{raw .Content }} {{end}}
		{{if .Tasks}}
			<table class="sessiontable">
			{{ range $index, $element := .Tasks}}
				<tr>
				<td>{{.TaskId}}</td>
				<td>{{.State}}</td>
				<td>{{.Operation}}</td>
				{{if .Result}}
				<td>{{.Result.Message}}</td>
				<td>{{.Result.Success}}</td>
				{{else}}
				<td colspan=2></td>
				{{end}}
				</tr>
			{{end}}
			</table>
		{{end}}
		{{if .Jobs}}
			<h1>Jobs</h1>
			<table class="jobtable">
			{{ range $index, $element := .Jobs }}
				<tr><td class="namew">{{.Name}}</td>
					<td class="iconwidth"><a href="/{{$customer}}/startjob/{{.UID}}"><img src="/___staticfilesstart.png"/></a></td>
					<td class="iconwidth"><a href="/{{$customer}}/stopjob/{{.UID}}"><img src="/___staticfilesstop.png"/></a></td>
					<td class="iconwidth"><a href="/{{$customer}}/editjob/{{.UID}}"><img src="/___staticfilesedit.png"/></a></td>
					<td class="iconwidth"><a href="/{{$customer}}/showjobsessions/{{.UID}}"><img src="/___staticfileslist.png"/></a></td>
					<td class="iconwidth"><a href="/{{$customer}}/quickbackup/{{.UID}}"><img src="/___staticfilesquickbackup.png"/></a></td>
					</tr>
			{{end}}
			</table>
		{{end}}
		{{if .Backups}}
			<h1>Backups</h1>
			<table class="jobtable">
			{{ range $index, $element := .Backups }}
				<tr><td class="namew">{{.Name}}</td>
				<td class="iconwidth"><a href="/{{$customer}}/showbackuppoints/{{.UID}}"><img src="/___staticfileslist.png"/></a></td>
					</tr>
			{{end}}
			</table>
		{{end}}
		{{if .Edit}}
		{{if .Edit.RealJob}}
			<div><a href="/{{$customer}}">Main</a></div>
			<h1>Job {{.Edit.Name}} ({{.Edit.JobType}})</h1>
			<p>Description : {{.Edit.RealJob.Description}}</p>
			{{if .Edit.RealJob.JobInfo.BackupJobInfo}}
			<form action="/{{$customer}}/editjobexec/{{.Edit.UID}}" method="post">
				<h1>General</h1>
				<table>
					<tr><td>Schedule Configured</td><td><input name="ScheduleConfigured" type="checkbox" {{if .Edit.RealJob.ScheduleConfigured}}checked{{end}}/></td></tr>
				</table>
				<h1>VMs</h1>
				<table>
					{{ range $index, $vm := .Edit.VMs}}
					
					<tr><td>
						<input name="{{$vm.UID}}" type="checkbox" {{if $vm.Added}}checked{{end}}/>{{if $vm.Root}}{{$vm.Root.Name}}\{{end}}{{$vm.Name}}
						<input type="hidden" name="{{$vm.UID}}_vmname" value="{{$vm.Name}}" />
						<input type="hidden" name="___vmlist" value="{{$vm.UID}}" />
						
					</td></tr>
					{{end}}
				</table>
				<input type="submit" value="Save"/>
			</form>
			{{else}}
				Not implemented yet!
			{{end}}
		{{end}}
		{{end}}
		{{if .Sessions}}
			<div><a href="/{{$customer}}">Main</a></div>
			<h1>Sessions</h1>
			<table class="sessiontable">
			{{ range $index, $element := .Sessions}}
				<tr>
				<td class="namew">{{.JobName}}</td>
				<td>{{.JobType}}</td>
				<td>{{.CreateUTC}}</td>
				
				<td class="progress"><div class="progress {{ if .Result | eq "Failed" }}red{{else}}green{{end}}" style="width:{{.Progress}}px;"></div><div class="progress gray" style="width:{{barminh .Progress}}px;"></div></td>
				<td>{{.State}}</td>
				<td>{{.Result}}</td>
				<td class="iconwidth"><a href="/{{$customer}}/showjobtasksessions/{{.SessionUID}}"><img src="/___staticfileslist.png"/></a></td>
				</tr>
			{{end}}
			</table>
		{{end}}
		{{if .RestoreSessions}}
			<div><a href="/{{$customer}}">Main</a></div>
			<h1>Restore Sessions</h1>
			<table class="sessiontable">
			{{ range $index, $element := .RestoreSessions}}
				<tr>
				<td class="namew">{{.Name}}</td>
				<td>{{.JobType}}</td>
				<td>{{.CreateUTC}}</td>
				<td class="progress"><div class="progress {{ if .Result | eq "Failed"  }}red{{else}}green{{end}}" style="width:{{.Progress}}px;"></div><div class="progress gray" style="width:{{barminh .Progress}}px;"></div></td>
				<td>{{.State}}</td>
				<td>{{.Result}}</td>
				</tr>
			{{end}}
			</table>
		{{end}}
		{{if .JobTaskSessions}}
			<div><a href="/{{$customer}}">Main</a></div>
			<h1>Sessions</h1>
			<table class="sessiontable">
			{{ range $index, $element := .JobTaskSessions}}
				<tr>
				<td>{{.VmName}}</td>
				<td>{{.CreateUTC}}</td>
				<td>{{.State}}</td>
				<td>{{.Result}}</td>
				<td>{{GB .TotalSize}} MB</td>
				</tr>
			{{end}}
			</table>
		{{end}}
		{{if .BackupRestorePoints}}
			<div><a href="/{{$customer}}">Main</a></div>
			<h1>RestorePoints</h1>
			<table class="jobtable">
			{{ range $index, $element := .BackupRestorePoints }}
				<tr><td class="namew">{{.Name}}</td>
				<td class="iconwidth"><a href="/{{$customer}}/showbackuppointvm/{{.UID}}"><img src="/___staticfilesvm.png"/></a></td>
					</tr>
			{{end}}
			</table>
		{{end}}
		{{if .BackupVMRestorePoints}}
			<div><a href="/{{$customer}}">Main</a></div>
			<h1>Restore Points</h1>
			<table class="jobtable">
			{{ range $index, $element := .BackupVMRestorePoints }}
				<tr><td class="namew">{{.Name}}</td>
				<td style="width:200px">{{.CreateUTC}}</td>
				<td>{{.Algorithm}}</td>
				<td>{{.PointType}}</td>
				<td class="iconwidth"><a href="/{{$customer}}/fullvmrestore/{{.UID}}"><img src="/___staticfilesrestore.png"/></a></td>
				</tr>
			{{end}}
			</table>
		{{end}}
		{{$bse := .BackupServer}}
		{{if .QuickBackup}}
			{{if $bse}}
			<div><a href="/{{$customer}}">Main</a></div>
			<h1>QuickBackup</h1>
			<table class="jobtable">
			{{ range $index, $element := .QuickBackup }}
				<tr><td class="namew">{{.Name}}</td>
				<td class="iconwidth"><a href="/{{$customer}}/quickbackupexec/{{$bse}}/{{.UID}}"><img src="/___staticfilesquickbackup.png"/></a></td>
				</tr>
			{{end}}
			</table>
			{{else}}
				Backup server not defined, should not happen
			{{end}}
		{{end}}		
		</div>
	</body>
</html>`

	funcMap := template.FuncMap{
        "barminh": func(i int) int {
            return 100-i
        },
		"GB": func(i int64) int64 {
            return i/1024/1024
        },
		"raw": func(s string) template.HTML {
			return template.HTML(s)
		},
    }
	_ = funcMap
	h.maintemplate, err = template.New("main").Funcs(funcMap).Parse(tpl)
	
	return err
}

func (h *SelfService) cacheStatic() {
var files map[string][]byte
files = make(map[string][]byte)
files["start.png"] = []byte{137,80,78,71,13,10,26,10,0,0,0,13,73,72,68,82,0,0,0,21,0,0,0,21,8,6,0,0,0,169,23,165,150,0,0,0,4,115,66,73,84,8,8,8,8,124,8,100,136,0,0,0,9,112,72,89,115,0,0,13,215,0,0,13,215,1,66,40,155,120,0,0,0,25,116,69,88,116,83,111,102,116,119,97,114,101,0,119,119,119,46,105,110,107,115,99,97,112,101,46,111,114,103,155,238,60,26,0,0,1,222,73,68,65,84,56,141,181,213,189,75,91,81,24,199,241,239,189,177,73,184,67,138,166,109,204,75,211,169,40,183,127,129,224,166,155,147,196,64,6,69,28,92,196,136,131,248,130,131,147,186,56,41,8,154,65,36,75,64,20,23,55,71,193,127,192,4,181,165,212,86,239,109,90,99,105,2,151,196,122,239,233,80,114,81,106,106,222,250,219,14,28,62,156,231,57,207,225,72,66,8,65,147,35,223,95,108,234,155,20,204,66,115,209,253,235,125,6,210,3,28,228,14,154,135,2,228,126,229,88,248,180,192,228,135,73,244,91,189,57,104,57,71,63,143,136,166,163,108,127,221,198,18,86,227,168,207,233,195,231,244,81,180,138,172,93,173,49,120,58,72,198,200,52,134,134,92,33,118,223,237,50,220,62,140,44,201,156,27,231,140,156,142,176,242,101,5,195,50,234,67,1,220,178,155,120,48,78,178,51,137,170,168,152,194,36,245,45,69,44,19,227,56,127,92,31,90,78,135,210,193,86,231,22,83,175,167,80,100,5,173,164,17,127,31,103,246,227,44,55,119,55,245,161,0,14,201,65,236,85,140,148,154,162,203,211,5,192,225,143,67,34,39,17,246,174,247,16,60,124,63,85,161,229,4,92,1,86,223,174,210,211,218,3,64,193,44,176,116,177,68,66,79,60,216,215,82,11,170,149,52,150,63,47,219,61,245,56,60,140,135,198,233,127,209,95,59,106,10,147,157,239,59,172,95,173,219,183,223,219,218,203,116,120,154,182,150,182,191,246,63,137,158,25,103,44,94,44,218,115,26,112,5,152,11,207,217,189,125,44,21,209,162,85,36,161,39,72,102,147,88,194,194,33,57,136,190,140,50,22,28,67,145,149,127,30,228,81,244,178,116,73,36,29,33,123,155,5,254,140,213,252,155,121,84,69,125,170,176,202,104,25,115,203,110,70,253,163,12,249,134,144,165,234,7,165,98,249,221,207,187,153,9,207,224,119,250,171,198,42,162,222,103,94,38,130,19,244,121,251,106,198,236,136,123,217,208,54,68,254,46,47,26,141,244,63,254,168,223,189,184,244,159,131,112,122,53,0,0,0,0,73,69,78,68,174,66,96,130}
files["stop.png"] = []byte{137,80,78,71,13,10,26,10,0,0,0,13,73,72,68,82,0,0,0,21,0,0,0,21,8,6,0,0,0,169,23,165,150,0,0,0,4,115,66,73,84,8,8,8,8,124,8,100,136,0,0,0,9,112,72,89,115,0,0,13,215,0,0,13,215,1,66,40,155,120,0,0,0,25,116,69,88,116,83,111,102,116,119,97,114,101,0,119,119,119,46,105,110,107,115,99,97,112,101,46,111,114,103,155,238,60,26,0,0,0,153,73,68,65,84,56,141,99,252,255,255,255,127,6,42,3,22,6,6,6,134,31,143,30,49,124,57,127,158,98,195,120,12,13,25,56,228,228,32,134,126,57,127,158,225,97,115,51,197,134,202,215,214,50,112,200,201,49,48,81,108,18,22,192,130,46,32,153,156,204,192,107,108,76,180,1,159,207,158,101,120,62,119,46,126,67,57,85,84,24,120,205,204,136,54,244,207,199,143,24,98,52,241,254,168,161,163,134,142,88,67,49,178,233,247,59,119,24,88,248,249,137,54,224,251,157,59,132,13,125,62,119,46,70,1,65,42,160,137,247,25,255,255,255,255,159,218,37,63,35,45,234,40,0,133,154,49,104,146,150,48,242,0,0,0,0,73,69,78,68,174,66,96,130}
files["list.png"] = []byte{137,80,78,71,13,10,26,10,0,0,0,13,73,72,68,82,0,0,0,21,0,0,0,21,8,6,0,0,0,169,23,165,150,0,0,0,4,115,66,73,84,8,8,8,8,124,8,100,136,0,0,0,9,112,72,89,115,0,0,13,215,0,0,13,215,1,66,40,155,120,0,0,0,25,116,69,88,116,83,111,102,116,119,97,114,101,0,119,119,119,46,105,110,107,115,99,97,112,101,46,111,114,103,155,238,60,26,0,0,1,37,73,68,65,84,56,141,213,148,61,107,194,80,20,134,31,53,26,77,154,218,54,131,46,209,82,226,7,232,36,14,46,153,92,252,69,58,56,136,131,131,46,254,51,113,17,68,165,31,100,113,113,16,52,98,91,147,14,129,75,179,25,48,131,239,248,28,120,224,220,247,112,99,158,231,121,220,56,241,91,11,1,36,128,222,123,143,197,113,33,96,231,165,131,149,181,232,127,244,67,201,106,106,141,241,219,216,151,238,126,118,108,191,183,98,184,255,221,115,118,207,1,118,77,242,169,60,16,209,250,209,189,105,89,41,19,255,231,47,164,11,104,146,70,227,161,17,74,102,42,38,0,177,40,78,74,2,24,126,14,89,58,75,1,219,207,109,90,143,45,70,95,163,80,178,170,82,101,240,58,240,165,246,217,102,125,90,139,97,93,173,227,92,156,0,187,38,106,66,5,238,174,125,67,54,56,94,142,2,230,82,57,148,132,66,41,83,10,37,51,100,3,136,178,253,137,61,97,227,108,4,180,158,44,154,90,147,169,61,13,37,51,21,147,174,209,245,165,43,103,197,252,48,23,195,98,186,72,37,83,97,118,152,133,146,186,184,192,221,181,175,39,117,241,109,1,100,165,44,114,92,14,176,107,162,39,117,32,162,246,35,89,255,15,134,54,85,244,233,3,126,31,0,0,0,0,73,69,78,68,174,66,96,130}
files["edit.png"] = []byte{137,80,78,71,13,10,26,10,0,0,0,13,73,72,68,82,0,0,0,21,0,0,0,21,8,6,0,0,0,169,23,165,150,0,0,0,4,115,66,73,84,8,8,8,8,124,8,100,136,0,0,0,9,112,72,89,115,0,0,13,215,0,0,13,215,1,66,40,155,120,0,0,0,25,116,69,88,116,83,111,102,116,119,97,114,101,0,119,119,119,46,105,110,107,115,99,97,112,101,46,111,114,103,155,238,60,26,0,0,1,190,73,68,65,84,56,141,213,213,191,75,35,65,20,192,241,239,172,147,203,114,97,13,196,35,114,151,59,225,14,174,210,74,176,179,244,31,176,14,88,9,138,96,32,149,165,118,166,180,16,196,210,214,82,196,214,255,193,242,14,35,172,222,173,138,49,138,151,205,238,206,58,115,133,176,92,48,187,106,180,185,233,102,222,123,31,222,252,128,17,198,24,195,27,15,235,173,193,87,161,218,104,86,154,43,172,187,235,40,163,122,98,114,80,116,235,247,22,135,237,67,52,154,227,238,49,155,223,55,177,45,123,240,78,247,90,123,236,94,238,82,126,87,198,182,108,220,192,37,210,81,18,127,49,122,244,231,136,237,95,219,24,12,231,209,57,69,89,164,241,173,193,176,28,30,12,245,34,143,213,147,85,2,19,224,107,31,128,249,143,243,76,58,147,61,121,201,153,42,163,104,199,237,84,240,58,190,166,246,179,70,94,228,185,141,111,1,168,142,86,153,253,48,251,40,55,65,67,29,210,236,54,251,130,26,205,218,201,26,57,43,135,23,121,0,140,23,198,169,85,106,125,243,31,221,254,152,61,198,136,28,233,89,91,252,177,136,35,29,220,192,125,40,18,146,133,79,11,12,137,161,231,161,82,72,242,86,62,153,111,156,109,112,161,46,184,137,111,48,24,4,130,250,151,58,5,171,208,23,236,139,254,59,14,90,7,236,183,246,81,90,37,79,102,249,243,50,19,239,39,136,77,156,90,151,122,251,161,14,217,185,220,65,32,232,232,14,0,51,165,25,230,70,231,178,250,200,70,79,195,83,238,212,29,165,92,9,120,56,235,198,215,198,147,32,100,108,223,139,60,186,166,75,39,234,80,45,87,89,170,44,61,11,204,68,109,203,166,94,169,51,93,156,78,186,125,53,58,229,76,129,243,34,43,29,189,82,87,248,247,126,102,81,160,3,164,72,127,56,73,68,32,144,66,226,223,251,79,162,64,38,42,254,155,239,228,47,76,161,165,175,238,3,197,28,0,0,0,0,73,69,78,68,174,66,96,130}
files["restore.png"] = []byte{137,80,78,71,13,10,26,10,0,0,0,13,73,72,68,82,0,0,0,21,0,0,0,21,8,6,0,0,0,169,23,165,150,0,0,0,4,115,66,73,84,8,8,8,8,124,8,100,136,0,0,0,9,112,72,89,115,0,0,13,215,0,0,13,215,1,66,40,155,120,0,0,0,25,116,69,88,116,83,111,102,116,119,97,114,101,0,119,119,119,46,105,110,107,115,99,97,112,101,46,111,114,103,155,238,60,26,0,0,2,14,73,68,65,84,56,141,213,213,191,75,27,97,24,192,241,239,93,2,77,114,33,13,105,41,196,210,40,87,2,197,130,63,22,193,77,93,4,21,255,1,55,39,137,99,150,44,78,78,1,103,227,34,18,17,116,16,116,10,184,100,144,44,130,154,205,37,180,18,181,73,35,73,27,147,11,38,241,210,183,131,244,236,113,137,45,141,14,125,166,123,238,121,222,207,221,243,114,188,39,9,33,4,79,28,118,128,139,250,5,105,45,221,53,54,236,30,38,224,8,220,163,105,45,205,114,118,185,107,116,169,119,137,128,35,128,220,181,212,38,254,31,212,254,55,77,62,187,15,255,11,63,138,172,112,213,188,34,215,200,253,59,218,231,232,35,244,54,196,216,203,49,100,233,97,168,204,109,134,149,203,21,78,170,39,109,215,117,28,127,212,51,74,252,67,156,9,239,132,9,4,8,58,131,196,130,49,230,253,243,143,163,242,111,190,234,84,137,190,143,162,216,148,142,83,200,146,76,168,39,196,148,111,202,82,51,198,15,191,11,51,224,30,32,249,61,201,184,119,28,151,236,50,154,178,245,44,91,133,45,26,162,193,140,111,134,17,207,136,81,139,244,70,56,170,30,81,186,43,89,223,52,173,165,81,29,42,11,61,11,244,43,253,70,67,241,174,72,248,83,152,189,226,30,137,82,130,197,204,34,137,82,194,168,187,100,23,115,111,230,218,143,127,92,61,70,23,58,54,201,102,106,240,216,60,236,126,220,101,210,55,9,128,64,16,189,140,82,209,43,70,207,236,235,89,211,58,3,45,235,101,110,244,27,203,254,52,69,147,245,252,58,7,223,14,140,123,181,86,141,253,210,190,145,123,237,94,6,149,65,43,10,144,170,164,16,60,28,90,90,75,99,231,122,135,88,46,102,121,216,97,249,208,148,15,185,135,58,160,229,20,90,75,51,192,237,235,109,214,114,107,22,16,224,172,118,134,46,116,35,87,157,170,113,109,250,248,79,181,83,108,146,141,90,171,70,188,16,103,35,191,209,22,132,251,109,89,253,178,138,64,160,181,52,178,245,44,211,175,166,173,104,253,71,157,243,219,115,146,229,36,241,175,241,142,224,175,216,44,108,154,242,182,40,64,228,115,132,124,51,255,71,240,177,144,132,16,226,169,79,126,233,57,254,81,63,1,34,2,199,116,21,146,251,10,0,0,0,0,73,69,78,68,174,66,96,130}
files["vm.png"] = []byte{137,80,78,71,13,10,26,10,0,0,0,13,73,72,68,82,0,0,0,21,0,0,0,21,8,6,0,0,0,169,23,165,150,0,0,0,4,115,66,73,84,8,8,8,8,124,8,100,136,0,0,0,9,112,72,89,115,0,0,13,215,0,0,13,215,1,66,40,155,120,0,0,0,25,116,69,88,116,83,111,102,116,119,97,114,101,0,119,119,119,46,105,110,107,115,99,97,112,101,46,111,114,103,155,238,60,26,0,0,1,133,73,68,65,84,56,141,237,213,177,75,2,97,24,199,241,239,213,89,28,210,101,121,100,16,5,129,137,139,75,139,138,180,138,46,37,8,65,131,244,55,180,54,251,23,56,11,14,130,56,137,32,216,208,226,208,152,131,16,36,20,52,134,118,137,151,89,150,94,215,80,72,71,150,193,229,16,244,27,159,231,225,195,243,190,240,242,10,134,97,24,252,114,68,128,59,253,142,174,222,181,140,41,54,5,155,96,123,67,211,215,105,114,141,156,101,52,227,205,224,179,251,152,178,44,141,200,63,58,97,212,103,247,177,191,188,111,26,216,118,110,19,144,3,68,23,163,36,215,147,8,8,0,72,83,18,41,119,10,191,236,255,30,189,120,188,32,174,196,17,5,113,88,219,93,218,165,254,80,103,101,118,5,143,228,97,203,177,5,192,142,178,131,107,198,133,98,83,190,71,123,47,61,78,239,79,9,202,65,0,60,146,135,198,115,131,246,160,13,64,190,153,103,111,105,15,81,16,9,205,135,168,180,43,227,143,15,80,190,45,19,117,70,1,136,56,35,148,91,229,97,79,211,53,46,31,47,57,92,59,228,184,117,140,142,254,51,180,218,169,178,33,109,32,79,203,132,228,16,39,218,137,169,159,109,100,217,156,219,228,168,117,52,18,132,247,183,255,49,6,6,149,118,133,131,213,3,106,221,26,79,47,79,166,126,243,185,73,236,44,246,37,56,114,83,120,187,2,183,228,166,164,150,134,53,181,175,210,25,116,76,115,106,95,69,27,104,227,55,5,184,234,93,145,56,79,152,106,69,181,248,105,174,112,83,248,249,166,86,243,119,80,17,192,43,121,9,47,132,45,99,14,209,1,128,48,137,63,234,21,150,81,122,212,46,98,175,109,0,0,0,0,73,69,78,68,174,66,96,130}
files["quickbackup.png"] = []byte{137,80,78,71,13,10,26,10,0,0,0,13,73,72,68,82,0,0,0,21,0,0,0,21,8,6,0,0,0,169,23,165,150,0,0,0,4,115,66,73,84,8,8,8,8,124,8,100,136,0,0,0,9,112,72,89,115,0,0,13,215,0,0,13,215,1,66,40,155,120,0,0,0,25,116,69,88,116,83,111,102,116,119,97,114,101,0,119,119,119,46,105,110,107,115,99,97,112,101,46,111,114,103,155,238,60,26,0,0,1,214,73,68,65,84,56,141,181,213,191,75,27,97,24,192,241,239,123,189,75,82,78,75,18,194,149,107,109,196,45,32,209,128,224,166,131,136,14,14,14,14,138,32,8,130,32,66,80,16,196,33,241,31,17,197,65,7,7,23,39,39,7,225,192,69,106,32,165,132,20,179,136,181,134,66,82,47,254,188,59,135,210,134,96,174,77,46,237,51,190,47,239,231,125,159,231,225,125,95,225,56,142,195,63,14,25,160,108,149,49,45,179,101,44,162,68,80,132,242,19,221,184,220,96,247,106,183,101,116,43,182,69,92,141,35,181,130,8,68,221,113,79,168,132,196,112,104,152,129,224,64,221,121,185,89,176,255,77,63,203,29,203,40,66,97,234,211,84,107,104,183,218,77,242,125,146,190,246,62,0,146,249,36,79,206,147,55,52,26,136,178,248,110,145,161,208,208,239,26,26,101,3,163,100,16,86,194,220,89,119,84,236,74,99,168,250,74,101,82,155,100,44,60,134,44,100,174,31,175,209,20,13,128,174,64,23,71,137,35,140,146,65,170,144,122,177,214,181,81,166,101,178,121,185,201,68,118,130,212,121,138,138,85,61,141,238,211,57,253,113,202,122,97,29,219,177,155,75,223,47,249,153,215,231,153,121,59,131,36,170,251,159,148,79,88,59,95,107,190,166,137,182,4,233,206,52,209,64,180,102,252,236,230,140,149,47,43,60,216,15,174,135,169,155,126,143,218,195,82,199,18,154,79,123,49,151,173,100,241,75,126,87,208,21,205,152,25,102,63,207,50,248,113,144,133,220,2,14,213,55,103,90,155,230,32,126,192,156,62,231,138,254,177,166,182,99,51,30,25,175,185,142,197,199,34,251,197,125,246,190,237,121,67,123,219,122,25,13,143,2,144,187,205,177,115,181,195,225,247,67,215,6,253,21,149,132,196,234,135,85,50,55,25,182,191,110,115,92,58,174,41,131,39,84,247,233,164,11,105,242,183,249,134,160,134,208,139,251,139,166,177,26,52,246,58,198,72,104,196,51,242,43,130,114,16,0,241,63,254,168,103,163,248,149,153,9,33,69,227,0,0,0,0,73,69,78,68,174,66,96,130}
h.files = &files
}

func (h *SelfService) cacheHierList() (error) {
	rs, err := h.rest.GetHierarchyRoots()
	vmlist := []*VM{}
	if err == nil {
		h.HierarchyRootCache = &rs.HierarchyRoot
		for _,root := range (*h.HierarchyRootCache) {
				urlg := url.Values{}
				urlg.Add("host",root.UID.String())
				urlg.Add("name","*")
				urlg.Add("type","Vm")

				vmsq,err := h.rest.GetLookup(urlg)
				if err == nil {
					for _,vmq := range vmsq.HierarchyItem {
							vmlist = append(vmlist,&VM{Added:false,Name:vmq.ObjectName,UID:vmq.ObjectRef.String(),Root:root})
					}
				}
		}
		h.VMs = &vmlist
		h.CacheTime = time.Now()
	}
	return err
}
func main() {
	fserver := flag.String("server", "localhost", "Server hosting rest")
	frport := flag.Int("rport", 9398, "Specify port rest port")
	fport := flag.Int("port", 16180, "Specify port listening")
	fnotsecure := flag.Bool("notsecure", false, "Set true to use http instead of https (for rest calls)")
	fverbose := flag.Bool("v", false, "Set true to get some verbose message")
	fusername := flag.String("username", "rest", "username")
	fpassword := flag.String("password", "Rest123", "password")
	flag.Parse()

	

	var rest * veeamrestapi.VeeamRestServer
	rest = &veeamrestapi.VeeamRestServer{
			UserName:*fusername,
			Password:*fpassword,
			Server:*fserver,
			Port:*frport,
			NotSecure:*fnotsecure,
			IgnoreSelfSigned:true,
			Logger:&VeeamCMDRestLogger{*fverbose}}		
		

	rest.Init()
	err := rest.Authenticate()
	if err == nil {
		selfservice := SelfService{rest:rest}
		err = selfservice.setTemplates()
		if err == nil {
			selfservice.cacheStatic()
			selfservice.cacheHierList()
			initmap := map[string]*WorkStatus{}
			selfservice.WorkList = &initmap
		
			if err == nil {
				log.Printf("Starting on http://localhost:%d",*fport)
				http.Handle("/", &selfservice)
				http.ListenAndServe(fmt.Sprintf(":%d",*fport), nil)
			}
		}
	} 

	if err != nil{
		log.Printf("%s",err)
	}

	
}