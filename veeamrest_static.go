package veeamrestapi

import (
	"fmt"
	"net/http"
	"log"
	"crypto/tls"
	"io/ioutil"
	"io"
	"encoding/xml"
	"strings"
)


//default logger, can be overwritten by implementing VeeamRestLogger
type VeeamRestLogger interface {
	LogRequest(request string)
	LogResponse(code int,xml string)
	LogInfo(info string)
}
type VeeamDefaultRestLogger struct {}
func (l * VeeamDefaultRestLogger) LogRequest(request string) {
  log.Print(request)
}
func (l * VeeamDefaultRestLogger) LogInfo(info string) {
  log.Print(info)
}
func (l * VeeamDefaultRestLogger) LogResponse(code int,xml string) {
  log.Printf("Status Code: %d",code)
  log.Printf("\n\n%s\n\n",xml)
}


//Resterror to create some custom errors
type VeeamRestError struct {
	errormsg string
	originalmsg string
}
func (e *VeeamRestError) Error() (string) {
	if e.originalmsg != "" {
		return fmt.Sprintf("%s : %s",e.errormsg,e.originalmsg)
	}
	return e.errormsg
}

type VeeamRestServer struct {
	Server string `xml:"Server,attr"`
	UserName string `xml:"UserName"`
	Password string `xml:"Password"`
	SessionId string `xml:"SessionId"`
	RestSvcSessionId string `xml:"RestSvcSessionId"`
	Client * http.Client `xml:"-"`
	Port int `xml:"Port,attr"`
	NotSecure bool `xml:"NotSecure,attr"`
	IgnoreSelfSigned bool `xml:"IgnoreSelfSigned,attr"`
	Logger VeeamRestLogger `xml:"-"`
}
type VeeamRestRequest struct {
	Url string
	Method string
	Post string
	SetBasicAuth bool
}

func (v * VeeamRestServer) Init() {
	if v.Server == "" {
		v.Server = "localhost"
	}
	if v.Port == 0 {
		if v.NotSecure {
			v.Port = 9399
		} else {
			v.Port = 9398
		}
	} 
	
	if(v.IgnoreSelfSigned) {
		tr := &http.Transport{ TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
		v.Client = &http.Client{Transport: tr}
	} else {
		v.Client = &http.Client{}
	}
	
	if v.Logger == nil {
	  v.Logger = &VeeamDefaultRestLogger{}
	}
}

func (v * VeeamRestServer) ConstructUrl(request string) (string) {
	//skipping if it is already url
	if len(request) > 8 && (strings.EqualFold("https://",request[:8]) || strings.EqualFold("http://",request[:7])) {
		return request
	}

	sec := ""
	if !v.NotSecure { sec = "s" }
	
	
	return fmt.Sprintf("http%s://%s:%d/api/%s",sec,v.Server,v.Port,request)
}
//just a small helper
func ReadStringFromReader(rc io.ReadCloser) (string,error) {
	if b, err := ioutil.ReadAll(rc); err == nil {
		return string(b),nil
	} else {
		return "",err
	}
} 
func (v * VeeamRestServer) MakeRequest(url string, method string) (vrr * VeeamRestRequest) {
	return &VeeamRestRequest{Url:url,Method:method}
}
func (v * VeeamRestServer) Request(vrr * VeeamRestRequest) (string,error) {
	if (v.Client == nil) { v.Init()}
	
	var returnerr error 
	returnxml := ""
	
	var err error
	var req * http.Request
	if vrr.Post != "" {
		req, err = http.NewRequest(vrr.Method, vrr.Url, strings.NewReader(vrr.Post))
	} else {
		req, err = http.NewRequest(vrr.Method, vrr.Url, nil)
	}
	if err == nil {
		if v.RestSvcSessionId != "" {
			req.Header.Add("X-RestSvcSessionId",v.RestSvcSessionId)
		}
		if vrr.SetBasicAuth {
			req.SetBasicAuth(v.UserName,v.Password)
		}
		v.Logger.LogRequest(fmt.Sprintf("Contacting (%s) %s",vrr.Method, vrr.Url))
		resp, err := v.Client.Do(req)
		if err == nil {
			defer resp.Body.Close() 
			if resp.StatusCode < 400 {
				xmlin,err := ReadStringFromReader(resp.Body)
				if err == nil {
					v.Logger.LogResponse(resp.StatusCode,xmlin)
					returnxml = xmlin
					
					restsessionid := ""
					//bug in api that is not respecting cases
					for name,val := range resp.Header {
						if strings.EqualFold("X-RestSvcSessionId",name) {
							restsessionid = val[0]
						}
					}

					if restsessionid != ""  {
						v.RestSvcSessionId = restsessionid
					}
				} else {
					returnerr = &VeeamRestError{"Error while reading response",err.Error()}
				}
		    } else {
				returnerr = &VeeamRestError{fmt.Sprintf("Got high return code from http %d",resp.StatusCode),""}
			}
		} else {
			returnerr = &VeeamRestError{"Error while getting data",err.Error()}
		}
	} else {
		returnerr = &VeeamRestError{"Error constructing url",err.Error()}
	}
	return returnxml,returnerr
}

func (v * VeeamRestServer) GenericGetXMLRequest(url string) (string,error) {
	vrr := v.MakeRequest(v.ConstructUrl(url),"GET")
	return v.Request(vrr)				
}
func (v * VeeamRestServer) GenericGetRequest(url string,o interface{}) (error) {
	xmlin,err := v.GenericGetXMLRequest(url)
	if err == nil {	
		err  = xml.Unmarshal([]byte(xmlin),o)
	}
	return err 
}

func (v * VeeamRestServer) GenericPostXMLRequest(url string,post string) (string,error) {
	vrr := v.MakeRequest(v.ConstructUrl(url),"POST")
	vrr.Post = post
	return v.Request(vrr)				
}
func (v * VeeamRestServer) GenericPostRequest(url string,post string,o interface{}) (error) {
	xmlin,err := v.GenericPostXMLRequest(url,post)
	if err == nil {	
		err  = xml.Unmarshal([]byte(xmlin),o)
	}
	return err 
}

func (v * VeeamRestServer) HeartbeatSession() (bool) {
	if (v.Client == nil) { v.Init()}
	
	noresetneeded := false
	if (v.SessionId != "") {
		vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s","logonSessions")),"GET")
		xmlin, err := v.Request(vrr)
		if err == nil {	
			logonSessions := LogonSessionListType{}
			err  := xml.Unmarshal([]byte(xmlin),&logonSessions)
			if err == nil {
				noresetneeded = true
			}
		}
	}
	if !noresetneeded {		
		v.SessionId = ""
		v.RestSvcSessionId = ""
	}
	return noresetneeded
}
func (v * VeeamRestServer) Authenticate() (error) {
	if (v.Client == nil) { v.Init()}
	var returnerr error 
	
	if v.UserName != "" && v.Password != "" {
		vrr := v.MakeRequest(v.ConstructUrl(""),"GET")
		xmlin, err := v.Request(vrr)
		if err == nil {
			entmgr := EnterpriseManagerType{}
			err  := xml.Unmarshal([]byte(xmlin),&entmgr)
			if err == nil {
				var sv * SupportedVersionType
				for _,t := range entmgr.SupportedVersions.SupportedVersion {
					if t.Name == "v1_2" {
						sv = &t
					}
				}
				if sv != nil {
					apilogin := string(sv.Links.Link[0].Href)
					v.Logger.LogInfo(fmt.Sprintf("Found api %s",apilogin))
					
					vrr = v.MakeRequest(apilogin,"POST")
					vrr.SetBasicAuth = true
					
					xmlin, err = v.Request(vrr)
					if err == nil {
						logonSession := LogonSessionType{}
						err  = xml.Unmarshal([]byte(xmlin),&logonSession)
						if err == nil {
							if v.RestSvcSessionId != "" {
								v.SessionId = logonSession.SessionId
								//so we are sure we get the full name correctly (even if we authenticated without domain prefix)
								v.UserName = logonSession.UserName
								v.Logger.LogInfo(fmt.Sprintf("Succesfully logged in with %s",v.UserName))
							} else {
								returnerr = &VeeamRestError{(fmt.Sprintf("Did not extract RestSvcSessionId from headers, possible no auth")),err.Error()}
							}
						} else {
							returnerr = &VeeamRestError{(fmt.Sprintf("Could not unmarshal logonSession")),""}
						}
					} else {
						returnerr = &VeeamRestError{(fmt.Sprintf("Error loading login page %s",vrr.Url)),""}
					}
				} else {
					returnerr = &VeeamRestError{"Couldn't find API, are you running v9?",""}
				}
				
			} else {
				returnerr = &VeeamRestError{(fmt.Sprintf("Could not unmarshal enterprisemanager, can not find API")),err.Error()}
			}
		} else {
			returnerr = &VeeamRestError{(fmt.Sprintf("Error requesting basic login page %s",vrr.Url)),err.Error()}
		}
	} else {
		returnerr = &VeeamRestError{"Did you set username and password?",""}
	}
	
	return returnerr
}

func (v * VeeamRestServer) LogoffSession(sessionid string) (error) {
	var returnerr error 
	if (v.SessionId != "") {
		vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s/%s","logonSessions",sessionid)),"DELETE")
		_, err := v.Request(vrr)
		if err == nil {	
			v.Logger.LogInfo(fmt.Sprintf("Succesfully logged off session %s",sessionid))
		}
	} else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
	return returnerr
}
func (v * VeeamRestServer) Logoff() (error) {
	return v.LogoffSession(v.SessionId)
}


func (v * VeeamRestServer) LogoffAllExceptThisSession() (error) {
	var returnerr error 
	if (v.SessionId != "") {
		vrr := v.MakeRequest(v.ConstructUrl(fmt.Sprintf("%s","logonSessions")),"GET")
		xmlin, err := v.Request(vrr)
		if err == nil {	
			logonSessions := LogonSessionListType{}
			err  := xml.Unmarshal([]byte(xmlin),&logonSessions)
			if err == nil {
					sessionlogofffail := []string{}
					
					for _,sess := range logonSessions.LogonSession {
						if(sess.UserName == v.UserName && sess.SessionId != v.SessionId) {
							v.Logger.LogInfo(fmt.Sprintf("Found another session :%s",sess.SessionId))
							err = v.LogoffSession(sess.SessionId)
							if err != nil {
								sessionlogofffail = append(sessionlogofffail,sess.SessionId)
							}
						}
					}	
					if len(sessionlogofffail) > 0 {
						returnerr = &VeeamRestError{(fmt.Sprintf("Some session could not be logged of %s",strings.Join(sessionlogofffail,","))),""}
					}
			} else {
				returnerr = &VeeamRestError{(fmt.Sprintf("Could not unmarshal sessions")),err.Error()}
			}
		} else {
			returnerr = &VeeamRestError{(fmt.Sprintf("Error requesting loggon session page %s",vrr.Url)),err.Error()}
		}
	} else { returnerr = &VeeamRestError{"No logon session set, did you login?",""} }
	return returnerr
}




		  
		  
		
			