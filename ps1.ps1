clear-host 
Get-Service | Where-Object { $_.Status -eq 'Running'}