#Veeam Rest CMD

Test case for Veeam Rest API bindings for GOLANG
````
Usage of veeamrestcmd:
  -notsecure
        Set true to use http instead of https
  -password string
        password
  -port int
        Specify port (default 9398)
  -s    Save Session
  -server string
        Server hosting rest (default "localhost")
  -sessionfile string
        .veeamrestcmdsave (default ".veeamrestcmdsave")
  -u    Save Session and Password (cleartext so not secure)
  -username string
        username (default "rest")
  -v    Set true to get some verbose message
````
To use persistent session, use -s or -u. -u will save your password in plaintext so it is less secure.

For example:
Start the session
````
veeamrestcmd -s -server enterprisemanager.local -username rest -password 123456
````
Next command, you can ommit server settings, your session will be reused. Watch out for timeouts though
````
veeamrestcmd -s
veeamrestcmd -s lsjobs
veeamrestcmd -s lsbackups
veeamrestcmd -s startjob <id>
````

If you even do unsecure save with plaintext pass, you can omit -u. However a session will be rebuild every time
````
veeamrestcmd -u -server enterprisemanager.local -username rest -password 123456
veeamrestcmd lsjobs
veeamrestcmd lsbackups
````

You can also use non persistent mode
````
veeamrestcmd -server enterprisemanager.local -username rest -password 123456 lsjobs
````


