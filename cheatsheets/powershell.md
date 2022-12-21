# Powershell Cheat Sheet

Reference: [Powershell Commands](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.management/remove-item?view=powershell-7.3)

1. Get-ChildItem - list directory contents

- `Get-ChildItem` - list contents of current directory
- `Get-ChildItem -Path C:\` - list contents of C: drive
- `Get-ChildItem -Path C:\ -Recurse` - list contents of C: drive and subdirectories

2. Set-Location - change directory

- `Set-Location C:\` - change to C: drive
- `Set-Location ..` - move up one directory level
- `Set-Location -` - go back to previous directory

3. Get-Location - print working directory

- `Get-Location` - print current directory
- `Get-Location -PSProvider FileSystem` - print current directory using filesystem provider
- `Get-Location -PSProvider Registry` - print current directory using registry provider

4. Get-Content - display contents of a file

- `Get-Content file1.txt` - display contents of file1.txt
- `Get-Content file1.txt file2.txt` - display contents of file1.txt and file2.txt
- `Get-Content file1.txt | Out-File file2.txt` - write contents of file1.txt to file2.txt

5. Select-String - search for a pattern in a file or stream

- `Select-String "pattern" file1.txt` - search for "pattern" in file1.txt
- `Select-String -Path file1.txt -Pattern "pattern" -CaseSensitive` - search for "pattern" in file1.txt, case sensitive
- `Select-String -Path file1.txt -Pattern "pattern" -NotMatch` - search for lines that do not match "pattern" in file1.txt

6. Get-Process - display a list of running processes

- `Get-Process` - display a list of running processes
- `Get-Process -Name "pattern"` - display a list of running processes matching "pattern"
- `Get-Process -Id 1234` - display a list of running processes with process id 1234

7. Get-Help - display help for a command

- `Get-Help Get-ChildItem` - display help for Get-ChildItem command
- `Get-Help Get-ChildItem -Full` - display full help for Get-ChildItem command
- `Get-Help Get-ChildItem -Examples` - display examples for Get-ChildItem command

8. Get-Command - display a list of available commands

- `Get-Command` - display a list of available commands
- `Get-Command -Name "pattern"` - display a list of available commands matching "pattern"
- `Get-Command -Verb "pattern"` - display a list of available commands with verb matching "pattern"

9. Get-Module - display a list of available modules

- `Get-Module` - display a list of available modules
- `Get-Module -Name "pattern"` - display a list of available modules matching "pattern"

10. Get-History - display a list of previously executed commands

- `Get-History` - display a list of previously executed commands
- `Get-History -Count 10` - display a list of the last 10 previously executed commands

11. Invoke-WebRequest - download a file

- `Invoke-WebRequest -Uri "https://example.com/file.zip" -OutFile "file.zip"` - download file.zip from https://example.com/file.zip

12. Invoke-Expression - execute a command

- `Invoke-Expression "command"` - execute "command"

13. Get-Date - display the current date and time

- `Get-Date` - display the current date and time

14. Get-EventLog - display a list of event logs

- `Get-EventLog` - display a list of event logs

15. Get-Event - display a list of events

- `Get-Event` - display a list of events

16. Get-EventSubscriber - display a list of event subscribers

- `Get-EventSubscriber` - display a list of event subscribers

17. Get-FormatData - display a list of format data

- `Get-FormatData` - display a list of format data

18. Get-Host - display information about the host

- `Get-Host` - display information about the host

19. Get-Item - display a list of items

- `Get-Item` - display a list of items

20. Get-ItemProperty - display a list of item properties

- `Get-ItemProperty` - display a list of item properties

21. Get-ItemPropertyValue - display a list of item property values

- `Get-ItemPropertyValue` - display a list of item property values

22. Create-Item - create an item

- `Create-Item` - create an item
- `Create-Item -Path "path"` - create an item at "path"
- `Create-Item -Path "path" -ItemType "type"` - create an item at "path" of type "type"

23. Get-ItemPropertyValue - display a list of item property values

- `Get-ItemPropertyValue` - display a list of item property values
- `Get-ItemPropertyValue -Path "path"` - display a list of item property values at "path"

24. Remove-Item - remove an item

- `Remove-Item` - remove an item
- `Remove-Item -Path "path"` - remove an item at "path"
- `Remove-Item -Path "path" -Recurse` - remove an item at "path" and subdirectories

25. Rename-Item - rename an item

- `Rename-Item` - rename an item
- `Rename-Item -Path "path" -NewName "newname"` - rename an item at "path" to "newname"

26. Set-Item - set an item

- `Set-Item` - set an item
- `Set-Item -Path "path" -Value "value"` - set an item at "path" to "value"
