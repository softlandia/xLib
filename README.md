# golang util library #

(c) softlandia@gmail.com

download: go get -u github.com/softlandia/xLib
install: go install

## dependences ##

>"golang.org/x/text/encoding/charmap"  
>"golang.org/x/text/transform"

## functions ##

1. FileExists(name string) bool
2. StrContainBackSlash(s string) bool
3. StrIsPrintRune(s string) bool
4. ChangeFileExt(iFileName, newExt string) string
5. SeekFileToString(fileName, strToSearch string) (*bufio.Scanner, error)
6. CodePageDetect(fn string) (int, error)
7. ConvertStrCodePage(s string, fromCP, toCP int64) (string, error)
8. FindFilesExt(fileList *[]string, path, fileNameExt string) (int, error)

## description ##

    func FileExists(name string) bool  
    return true if file exist

    func StrContainBackSlash(s string) bool
    return true if string s consist rune back slash '\'

    func StrIsPrintRune(s string) bool  
    return true if input string consists only of printable rune

    func ChangeFileExt(iFileName, newExt string) string  
    return file name with new extention

    func SeekFileToString(fileName, strToSearch string) (*bufio.Scanner, error)  
    read text file fileName and return Scanner at line strToSearch

    func CodePageDetect(fn string, stopStr ...string) (int, error)  
    read text file fn and return code page, detect only IBM CodePage866 and Windows1251  
    return constant Cp866, Cp1251, CpEmpty
    if string stopStr is present then input file scanned befor appearance stopStr

    func ConvertStrCodePage(s string, fromCP, toCP int64) (string, error)
    convert string from one code page to another


    func FindFilesExt(fileList *[]string, path, fileNameExt string) (int, error) 
    search in path files with extention == fielNameExt and put file name to slice fileList
