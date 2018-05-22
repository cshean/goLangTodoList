# Creates a shortcut into the shortcuts folder from the go location
#These need to be replaced with paths relative to the user and the executables location
$WshShell = New-Object -comObject WScript.Shell
$Shortcut = $WshShell.CreateShortcut("C:\Users\carte\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup\todo.lnk")
$Shortcut.TargetPath = "C:\Users\carte\go\src\todo\todo.exe"
$Shortcut.Save()