package main

import(
	"github.com/tdewin/veeamrestapi"
	"log"
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"encoding/xml"
)


type VeeamInteractiveRestLogger struct {
	responenum int
}
func (l * VeeamInteractiveRestLogger) LogRequest(request string) {
  log.Print(request)
}
func (l * VeeamInteractiveRestLogger) LogInfo(info string) {
  log.Print(info)
}
func (l * VeeamInteractiveRestLogger) LogResponse(code int,xml string) {
  log.Printf("Status Code: %d",code)
  
  l.responenum = l.responenum+1
  filename := fmt.Sprintf("%d-%d.xml",l.responenum,code)
  
  xmlbyte := []byte(xml)
  
  err := ioutil.WriteFile(filename, xmlbyte, 0666)
  if err != nil {
	log.Printf("Could not write to %s",filename)
  }
}


func main() {
	fserver := flag.String("server", "localhost", "Server hosting rest")
	fport := flag.Int("port", 9398, "Specify port")
	fnotsecure := flag.Bool("notsecure", false, "Set true to use http instead of https")
	fusername := flag.String("username", "rest", "username")
	fpassword := flag.String("password", "", "password")
	
	flag.Parse()

	idregexp := regexp.MustCompile("^urn\\:veeam\\:[a-zA-Z_0-9\\-]+\\:([a-zA-Z_0-9\\-]+)$")
	_ = idregexp
	
	rest := veeamrestapi.VeeamRestServer{
		UserName:*fusername,
		Password:*fpassword,
		Server:*fserver,
		Port:*fport,
		NotSecure:*fnotsecure,
		IgnoreSelfSigned:true,
		Logger:&VeeamInteractiveRestLogger{0}}
	
	rest.Init()
	
	err := rest.Authenticate()
	if err == nil {
		jobs,err := rest.GetJobs()
		jobuid := ""
		if err == nil {
			for i,job := range jobs.Job {
				xml.MarshalIndent(&job,"","  ")
				log.Printf("\n%s",job.Name)
				if i == 0 {
					if t := idregexp.FindStringSubmatch(job.UID.S());len(t) > 1 {
						jobuid = t[1]
					}
				}
			}
		} else {
			log.Printf("%v",err)
		}
		
		if jobuid != "" {
			log.Printf("Got single backup UID, try singlemethod")
			job, err := rest.GetJob(jobuid)
			if err == nil {
				log.Printf("\n%s",job.Name)
			}  else {
				log.Printf("%v",err)
			}
		}
		
		backups,err := rest.GetBackups()
		if err == nil {
			for _,backup := range backups.Backup {
				xml.MarshalIndent(&backup,"","  ")
				log.Printf("\n%s",backup.Name)
			}
		} else {
			log.Printf("%v",err)
		}
		
		replicasref,err := rest.GetEntityReferences("replicas")
		if err == nil {
			for _,replicaref := range replicasref.Ref {

				if t := idregexp.FindStringSubmatch(replicaref.UID.S());len(t) > 1 {
					log.Printf("Found replica id, trying to query ref %s",t[1])
					replicasingleref,err := rest.GetEntityRef("replicas",t[1])
					if err == nil {
						log.Printf("\nSingle ref test %s %s",replicasingleref.Name,replicasingleref.Href)
					}
				}
				log.Printf("\n%s %s",replicaref.Name,replicaref.Href)
			}
		} else {
			log.Printf("%v",err)
		}		
		err = rest.LogoffAllExceptThisSession()
		if err != nil {
			log.Printf("%s",err)
		}
		err = rest.Logoff()
		if err != nil {
			log.Printf("%s",err)
		}
	} else {
		log.Printf("%s",err)
	}
	
}