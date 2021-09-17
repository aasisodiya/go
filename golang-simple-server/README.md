# Simple Go Lang Server

## Powershell Command to Test

```ps
$response = Invoke-RestMethod 'http://localhost:8090/get-sample' -Method 'GET' -Headers $headers
$response | ConvertTo-Json
```
