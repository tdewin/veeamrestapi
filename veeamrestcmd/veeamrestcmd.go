package main

import(
	"github.com/tdewin/veeamrestapi"
	"log"
	"flag"
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"net/url"
)


//cmd rest logger remove all verbose info unless requested with -v
type VeeamCMDRestLogger struct {
	verbose bool
}
func (l * VeeamCMDRestLogger) LogRequest(request string) {
	if l.verbose {
		log.Printf(request)
	}
}
func (l * VeeamCMDRestLogger) LogInfo(info string) {
	if l.verbose {
		log.Printf(info)
	}
}
func (l * VeeamCMDRestLogger) LogResponse(code int,xmlout string) {
	if l.verbose {
		log.Printf("Got return code %d",code)
		log.Printf(xmlout)
	}
}




//error
type cmdError struct { e string }
func (e cmdError) Error() (string) { return e.e }

//cmdfunctions
func LsLookupSvc(rest * veeamrestapi.VeeamRestServer) {
	l,err := rest.GetLookupSvc() 
	 if err == nil {
		fmt.Printf("%s\n\n",l.FullString(true,0))
	} else {
		fmt.Printf("Error %s",err)
	}
}

func LsQuerySvc(rest * veeamrestapi.VeeamRestServer) {
	qrp,err := rest.GetQuerySvc() 
	 if err == nil {
		fmt.Printf("%s\n\n",qrp.FullString(true,0))
	} else {
		fmt.Printf("Error %s",err)
	}
}

func LsReportingSvc(rest * veeamrestapi.VeeamRestServer) {
	qrp,err := rest.GetReportingSvc()
	 if err == nil {
		fmt.Printf("%s\n\n",qrp.FullString(true,0))
	} else {
		fmt.Printf("Error %s",err)
	}
}

func LsReport(rest * veeamrestapi.VeeamRestServer,report string) {
	var result fmt.Stringer
	var err error
	
	if strings.EqualFold(report,"summary") {
		result,err = rest.GetSummaryReport()
	} else if  strings.EqualFold(report,"overview") {
		result,err = rest.GetOverviewReportFrame()
	} else if strings.EqualFold(report,"vmsoverview") {
		result,err = rest.GetVmsOverviewReportFrame()
	} else if strings.EqualFold(report,"JobStatistics") {
		result,err =  rest.GetJobStatisticsReportFrame()
	} else if strings.EqualFold(report,"repositories") {
		result,err = rest.GetRepositoryReportFrame()
	} else if strings.EqualFold(report,"processedvms") {
		result,err = rest.GetProcessedVmsReportFrame()
	} 
	
	if result != nil && err == nil {
		fmt.Printf(result.String())
	} else if err == nil {
		log.Printf("Unknown report")
		LsReportingSvc(rest)
	} else {
		log.Printf("Error :"+err.Error())
	}
}


func LsQuery(rest * veeamrestapi.VeeamRestServer, vals url.Values) {
  qrt,err := rest.GetQueryResult(vals) //(QueryResultType,error)
  if err == nil {
	fmt.Printf("%s\n\n",qrt.FullString(true,0))
  } else {
	fmt.Printf("Error %s",err)
  }
}

func LsLookup(rest * veeamrestapi.VeeamRestServer, vals url.Values) {
  x,err := rest.GetLookup(vals)
  if err == nil {
	fmt.Printf("%s\n\n",x.FullString(true,0))
  } else {
	fmt.Printf("Error %s",err)
  }
}



func LsJobs(rest * veeamrestapi.VeeamRestServer) {
	jobs,err := rest.GetJobs()
	if err == nil {
		for _,job := range jobs.Job {
			fmt.Printf("%s\n\n",job.String())
		}
	} else {
		log.Printf("%s",err)
	}
}
func LsJob(rest * veeamrestapi.VeeamRestServer,id string) {
	job,err := rest.GetJob(id)
	if err == nil {
		fmt.Printf("%s\n\n",job.FullString(true,0))
	} else {
		log.Printf("Error catching job %s",err)
	}
}
func LsBackups(rest * veeamrestapi.VeeamRestServer) {
	backups,err := rest.GetBackups()
	if err == nil {
		for _,backup := range backups.Backup {
			fmt.Printf("%s\n\n",backup.String())
		}
	} else {
		log.Printf("%s",err)
	}
}
func LsTasks(rest * veeamrestapi.VeeamRestServer) {
	tasks,err := rest.GetTasks()
	if err == nil {
		for _,task := range tasks.Task {
			fmt.Printf("%s",task.String())
			if task.Result != nil {
				fmt.Printf("%s",task.Result.String())
			}
			fmt.Print("\n\n")
		}
	} else {
		log.Printf("%s",err)
	}
}
func LsOverview(rest * veeamrestapi.VeeamRestServer) {
	frame,err := rest.GetOverviewReportFrame()
	if err == nil {
		fmt.Printf("------------------------------------------------\n")
		fmt.Printf(" Connected to : %s\n",rest.Server)
		fmt.Printf("------------------------------------------------\n")
		fmt.Println(frame.String())
		
	} else {
		log.Printf("%s",err)
	}
}
func ActionJob(rest * veeamrestapi.VeeamRestServer,action string,jobid string) {
	job,err := rest.GetJob(jobid)
	if err == nil {
		if links := veeamrestapi.FindLinkByRel(job,action);len(links) > 0 {
			task := veeamrestapi.TaskType{}
			var err error
			
			if strings.EqualFold(action,"clone") {
				spec := veeamrestapi.NewJobCloneSpecType()
				jobnameclone := job.Name+"_restclone"
				known := false
				if strings.EqualFold(job.JobType,"Backup") {
					known = false
					spec.BackupJobCloneInfo =  &veeamrestapi.BackupJobCloneInfoType{
						JobName:jobnameclone,
						FolderName:"restfoldername",
						RepositoryUid:"urn:veeam:Repository:uid"}
				} else if  strings.EqualFold(job.JobType,"Replica") {
					known = true
					spec.ReplicaJobCloneInfo =  &veeamrestapi.ReplicaJobCloneInfoType{JobName:jobnameclone,VmSuffix:"_replica"}
				} 
				
				if known {
					err = rest.GenericPostRequest(links[0].Href.String(),veeamrestapi.MarshalForPost(spec),&task)
				} else {
					err = cmdError{"Canot not clone type "+job.JobType}
					 
				}
			} else {
				err = rest.GenericPostRequest(links[0].Href.String(),"",&task)
			}
			
			if err == nil {
				fmt.Printf("%-30s %10s\n",task.TaskId,task.State)
			} else {
				log.Printf("Error executing %s",err)
			}
		} else {
			log.Printf("Could not find matching link %s",err)
		}
	} else {
		log.Printf("Error catching job %s",err)
	}
}


//main 
//handles all session and argument parsing
func main() {
	fserver := flag.String("server", "localhost", "Server hosting rest")
	fport := flag.Int("port", 9398, "Specify port")
	fnotsecure := flag.Bool("notsecure", false, "Set true to use http instead of https")
	fverbose := flag.Bool("v", false, "Set true to get some verbose message")
	
	
	fsave := flag.Bool("s", false, "Save Session")
	fsaveunsecure := flag.Bool("u", false, "Save Session and Password (cleartext so not secure)")
	fsessionfile := flag.String("sessionfile", ".veeamrestcmdsave", ".veeamrestcmdsave")
	
	fusername := flag.String("username", "rest", "username")
	fpassword := flag.String("password", "", "password")
	flag.Parse()

	argsrem := flag.Args()
	onetimeexec := true
	var unmarshallerror error
	
	idmatch := regexp.MustCompile("^[a-zA-Z\\-0-9]+$")
	
	
	
	
	var rest * veeamrestapi.VeeamRestServer
	if _, err := os.Stat(*fsessionfile); err == nil {
		rest = &veeamrestapi.VeeamRestServer{}
		var data []byte
		
		data,unmarshallerror = ioutil.ReadFile(*fsessionfile)
		if unmarshallerror == nil {
			unmarshallerror = xml.Unmarshal(data,rest)
		}
		if rest.Password == "" && *fpassword != "" {
			rest.Password = *fpassword
		}
		rest.Logger = &VeeamCMDRestLogger{*fverbose}
	} else {
		rest = &veeamrestapi.VeeamRestServer{
			UserName:*fusername,
			Password:*fpassword,
			Server:*fserver,
			Port:*fport,
			NotSecure:*fnotsecure,
			IgnoreSelfSigned:true,
			Logger:&VeeamCMDRestLogger{*fverbose}}		
	}
		
		
		
		
	if unmarshallerror != nil {
		log.Printf("Session file corrupted, please remove %s",*fsessionfile)
	} else {
		rest.Init()
		
		sessionvalid := false
		
		if rest.RestSvcSessionId != "" {
			sessionvalid = rest.HeartbeatSession()
		}
		
		var autherr error
		if !sessionvalid {
			
			autherr = rest.Authenticate()
		}
		
		if autherr == nil {
			//if we were able to login, we can save the session, otherwise there is no point trying
			xmlout := []byte{}
			if *fsave {
				t := rest.Password 
				rest.Password = ""
				xmlout,_ = xml.MarshalIndent(rest,"","  ")
				rest.Password = t
				onetimeexec = false
			} else if *fsaveunsecure {
				xmlout,_ = xml.MarshalIndent(rest,"","  ")
				onetimeexec = false
			}
			
			if len(xmlout) > 0 {
				err := ioutil.WriteFile(*fsessionfile,xmlout,0666)
				if err != nil {
					log.Printf("Was not able to write file %s : %s",*fsessionfile,err.Error())
				}
			}
			
			if len(argsrem) > 0 {
				var err error
				
				switch argsrem[0] {
					case "lsjobs":
							LsJobs(rest)
						
					break;
					case "lsjob":
						if len(argsrem) > 1 && idmatch.MatchString(argsrem[1]) { 
							LsJob(rest,argsrem[1])
						} else {
							log.Printf("Please supply and id (lsjobs)")
						}
					break;
					case "lsbackups":
						LsBackups(rest)
					break;
					case "lsoverview":
						LsOverview(rest)
					break;
					case "lstasks":
						LsTasks(rest)
					break;
					case "lsquerysvc":
						LsQuerySvc(rest)
					break;
					case "lslookupsvc":
						 LsLookupSvc(rest)
					break;
					case "lslookup":
						vals := url.Values{}
						for i:=1;i+1 < len(argsrem);i=i+2 {
							vals.Add(argsrem[i],argsrem[i+1])
						}
						LsLookup(rest,vals)
					break;
					case "lsreportingsvc":
						LsReportingSvc(rest)
					break;
					case "lsquery":
						vals := url.Values{}
						for i:=1;i+1 < len(argsrem);i=i+2 {
							vals.Add(argsrem[i],argsrem[i+1])
						}
						
						LsQuery(rest,vals)
					break;
					case "lsreport":
						if len(argsrem) > 1 {
						 LsReport(rest,argsrem[1])
						} else {
						  log.Printf("Please supply report suffix from url (lsreportingsvc)")
						}
					break;
					case "startjob":
						if len(argsrem) > 1 && idmatch.MatchString(argsrem[1]) { 
							ActionJob(rest,"start",argsrem[1])
						} else {
							log.Printf("Please supply and id (lsjobs)")
						}
					break
					case "retryjob":
						if len(argsrem) > 1 && idmatch.MatchString(argsrem[1]) { 
							ActionJob(rest,"retry",argsrem[1])
						} else {
							log.Printf("Please supply and id (lsjobs)")
						}
					break
					case "stopjob":
						if len(argsrem) > 1 && idmatch.MatchString(argsrem[1]) { 
							ActionJob(rest,"stop",argsrem[1])
						} else {
							log.Printf("Please supply and id (lsjobs)")
						}
					break
					case "clonejob":
						if len(argsrem) > 1 && idmatch.MatchString(argsrem[1]) { 
							ActionJob(rest,"clone",argsrem[1])
						} else {
							log.Printf("Please supply and id (lsjobs)")
						}
					break
					case "exlogoff":
						err = rest.Logoff()
						if err != nil {
							log.Printf("%s",err)
						}
					break
					default:
						log.Printf("Unknown action")
						log.Printf("lsoverview | lstasks | lsjobs | lsjob | startjob | retryjob | stopjob | lsbackups")
					break
				}
				
				if onetimeexec {
						err = rest.Logoff()
						if err != nil {
							log.Printf("%s",err)
						}			
				}
			} else {
				LsOverview(rest)
			}
		} else {
			log.Printf("%s",autherr)
		}
	}
	
}