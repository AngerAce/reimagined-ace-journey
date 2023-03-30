package main

import (
    "flag"
    "fmt"
    "os"
)

// 列出当前目录中的所有文件和子目录
func listFiles(args []string) error {
    // 解析命令行参数
    flags := flag.NewFlagSet("list-files", flag.ExitOnError)
    path := flags.String("path", ".", "directory to list files in")
    if err := flags.Parse(args); err != nil {
        return err
    }

    // 列出当前目录下的所有文件和子目录
    files, err := os.ReadDir(*path)
    if err != nil {
        return err
    }
    for _, file := range files {
        fmt.Println(file.Name())
    }
    return nil
}

// 计算文件大小
func getFileSize(args []string) error {
    // 解析命令行参数
    flags := flag.NewFlagSet("get-file-size", flag.ExitOnError)
    path := flags.String("path", ".", "path to file")
    if err := flags.Parse(args); err != nil {
        return err
    }

    // 获取文件大小
    info, err := os.Stat(*path)
    if err != nil {
        return err
    }
    size := info.Size()
    if size < 1024 {
        // 文件大小小于1KB
        fmt.Printf("%d bytes\n", size)
    } else if size < 1048576 {
        // 文件大小小于1MB
        fmt.Printf("%.2f KB\n", float64(size)/1024)
    } else if size < 1073741824 {
        // 文件大小小于1GB
        fmt.Printf("%.2f MB\n", float64(size)/1048576)
    } else {
        // 文件大小大于等于1GB
        fmt.Printf("%.2f GB\n", float64(size)/1073741824)
    }
    return nil
}

// 查找文件
func findFiles(args []string) error {
    // 解析命令行参数
    flags := flag.NewFlagSet("find-files", flag.ExitOnError)
    pattern := flags.String("pattern", "*", "search pattern")
    path := flags.String("path", ".", "directory to search in")
    if err := flags.Parse(args); err != nil {
        return err
    }

    // 查找文件
    files, err := os.ReadDir(*path)
    if err != nil {
        return err
    }
    for _, file := range files {
        if file.IsDir() {
            // 如果是子目录，则递归查找
            subpath := fmt.Sprintf("%s%c%s", *path, os.PathSeparator, file.Name())
            if err := findFiles([]string{"-pattern", *pattern, "-path", subpath}); err != nil {
                return err
            }
        } else if match, err := file.Match(*pattern); err != nil {
            return err
        } else if match {
            fmt.Printf("%s%c%s\n", *path, os.PathSeparator, file.Name())
        }
    }
    return nil
}

// 显示帮助信息
func showHelp() {
    // 显示帮助信息
    fmt.Println("Usage: myapp <command> [options]")
   
