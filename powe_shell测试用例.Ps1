# 列出当前目录中的所有文件和子目录
function List-Files($path=".") {
    # 列出当前目录下的所有文件和子目录
    Get-ChildItem -Path $path -Recurse
}

# 计算文件大小
function Get-FileSize($path) {
    # 获取文件大小
    $size = (Get-Item $path).Length
    if ($size -lt 1024) {
        # 文件大小小于1KB
        return "$size bytes"
    }
    elseif ($size -lt 1048576) {
        # 文件大小小于1MB
        return "{0:N2} KB" -f ($size / 1KB)
    }
    elseif ($size -lt 1073741824) {
        # 文件大小小于1GB
        return "{0:N2} MB" -f ($size / 1MB)
    }
    else {
        # 文件大小大于等于1GB
        return "{0:N2} GB" -f ($size / 1GB)
    }
}

# 查找文件
function Find-Files($pattern, $path=".") {
    # 查找文件
    Get-ChildItem -Path $path -Filter $pattern -Recurse
}

# 查找进程
function Find-Process($name) {
    # 查找进程
    Get-Process -Name $name
}

# 查看系统信息
function Get-SystemInfo {
    # 查看系统信息
    $os = Get-WmiObject -Class Win32_OperatingSystem
    $cpu = Get-WmiObject -Class Win32_Processor
    $mem = Get-WmiObject -Class Win32_PhysicalMemory | Measure-Object -Property Capacity -Sum
    Write-Host "Operating System: $($os.Caption) $($os.Version)"
    Write-Host "CPU: $($cpu.Name) $($cpu.NumberOfCores) cores"
    Write-Host "Memory: $($mem.Sum / 1GB) GB"
}

# 显示帮助信息
function Show-Help {
    # 显示帮助信息
    Write-Host "List-Files - 列出当前目录中的所有文件和子目录"
    Write-Host "Get-FileSize - 计算文件大小"
    Write-Host "Find-Files - 查找文件"
    Write-Host "Find-Process - 查找进程"
    Write-Host "Get-SystemInfo - 查看系统信息"
}

# 主程序
function main {
    # 显示帮助信息
    Show-Help

    # 创建一个输入对象
    $choice = 0
    while ($choice -ne 6) {
        # 显示菜单
        Write-Host "1. List files"
        Write-Host "2. Get file size"
        Write-Host "3. Find files"
        Write-Host "4. Find process"
        Write-Host "5. Get system info"
        Write-Host "6. Exit"
        $choice = Read-Host "Enter your choice"

        switch ($choice) {
            1 {
                # 列出当前目录中的所有文件和子目录
                List-Files
            }
            2 {
                # 计算
