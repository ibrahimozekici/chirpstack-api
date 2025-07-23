# PowerShell script to clean up unnecessary generated files for kitchen-only API

Write-Host "Cleaning up unnecessary files for kitchen-only API..." -ForegroundColor Green

# Remove unnecessary Go generated files (keep only kitchen and internal)
$goApiPath = "c:\Users\Utkan\Documents\GitHub\chirpstack-api\go\as\external\api"
Get-ChildItem -Path $goApiPath -Name | Where-Object { 
    $_ -match "\.pb\.(go|gw\.go)$" -and $_ -notmatch "(kitchen|internal)" 
} | ForEach-Object {
    $filePath = Join-Path $goApiPath $_
    Write-Host "Removing: $filePath" -ForegroundColor Yellow
    Remove-Item $filePath -Force
}

# Remove unnecessary Swagger files (keep only kitchen and internal)
$swaggerApiPath = "c:\Users\Utkan\Documents\GitHub\chirpstack-api\swagger\as\external\api"
Get-ChildItem -Path $swaggerApiPath -Name | Where-Object { 
    $_ -match "\.swagger\.json$" -and $_ -notmatch "(kitchen|internal)" 
} | ForEach-Object {
    $filePath = Join-Path $swaggerApiPath $_
    Write-Host "Removing: $filePath" -ForegroundColor Yellow
    Remove-Item $filePath -Force
}

# Remove other unused Go directories
$goBasePath = "c:\Users\Utkan\Documents\GitHub\chirpstack-api\go"
$dirsToRemove = @("ad", "als", "common", "fuota", "geo", "gw", "nc", "ns")

foreach ($dir in $dirsToRemove) {
    $dirPath = Join-Path $goBasePath $dir
    if (Test-Path $dirPath) {
        Write-Host "Removing directory: $dirPath" -ForegroundColor Yellow
        Remove-Item $dirPath -Recurse -Force
    }
}

# Remove as.pb.go and integration directory if they exist
$asFile = Join-Path $goBasePath "as\as.pb.go"
if (Test-Path $asFile) {
    Write-Host "Removing: $asFile" -ForegroundColor Yellow
    Remove-Item $asFile -Force
}

$integrationDir = Join-Path $goBasePath "as\integration"
if (Test-Path $integrationDir) {
    Write-Host "Removing directory: $integrationDir" -ForegroundColor Yellow
    Remove-Item $integrationDir -Recurse -Force
}

Write-Host "Cleanup completed!" -ForegroundColor Green
Write-Host "Remaining files:" -ForegroundColor Cyan
Write-Host "Go API files:" -ForegroundColor White
Get-ChildItem -Path $goApiPath -Name | Sort-Object

Write-Host "`nSwagger files:" -ForegroundColor White
Get-ChildItem -Path $swaggerApiPath -Name | Sort-Object
