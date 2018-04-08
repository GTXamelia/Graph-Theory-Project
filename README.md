# Graph Theory Project

The project takes a regular expression and compares it to a string of text.

## Table Of Contents

- [Graph Theory Project](#graph-theory-project)
  * [Table Of Contents](#table-of-contents)
  * [Getting Started](#getting-started)
    + [Installing](#installing)
  * [Test Cases](#test-cases)
    + [True Test Cases](#true-test-cases)
    + [False Test Cases](#false-test-cases)
    + [Menu Option 1 Sample](#menu-option-1-sample)
      - [True](#true)
      - [False](#false)
    + [Menu Option 2 Sample](#menu-option-2-sample)
      - [True](#true-1)
      - [False](#false-1)
      - [File Not Found](#file-not-found)
    + [Menu Option 3 Sample](#menu-option-3-sample)
    + [Invalid Menu Option](#invalid-menu-option)
  * [Built With](#built-with)
  * [Authors](#authors)
  * [License](#license)
  * [Acknowledgments](#acknowledgments)
  
## Features
1. Evaluate a string using a regular expression

Allow the user to evaluate a string using a regular expression.

Example:
```
Regular Expression: ab+ ---> String abbbbbbb ---> True
Regular Expression: ab+ ---> String accccccc ---> False
```

2. Evaluate a txt file contents using a regular expression

The user can select a txt file in the same directory that the launcher is in.
The user just enters the file name not file extenstion, as the program will append that to the file name.

Example: 
```
File: testingFile.txt ----> User enters: testingFile
```

3. No need to concatanate regular expression (Can still concatanate without any problems)

The user doesn't need to concatanate a regular expression (although they still can)

Example: 
```
Regular Expression: abc   ---> String abc ---> True
Regular Expression: a.b.c ---> String abc ---> False
```


## Getting Started

- Right click on project folder
- Git init
- Git remote add origin https://github.com/cian2009/UWPIrishTransportTracker.git
- Git pull origin master
- Import project into Visual Studio Code

### Installing

A step by step series of examples that tell you have to get a development enviroment running

1.Set-up Git Repository

```
Git init
Git remote add origin https://github.com/cian2009/Graph-Theory-Project-.git
Git pull origin master
```

2.Add project to IDE

```
Open Visual Studio Code IDE
Drag folder of git contents into Visual Studio Code
```

3.Add Color package for Go

```
Open git bash for project repository
Enter 'go get github.com/fatih/color'
```

4.Building project

```
Open integrated terminal
cd .\main\
go build .\main.go
```

## Test Cases

### True Test Cases

String:                    |  Regular Expression:      | Status:
:-------------------------:|:-------------------------:|:-------------------------:
abc                        |abc                        |True
abc+                       |abcccc                     |True
abc*                       |abcccc                     |True
abc?                       |ab                         |True
abc&#124;d                 |abc                        |True
abc(a*&#124;d+)            |abdddddddd                 |True
 
### False Test Cases
 
String:                    |  Regular Expression:      | Status:
:-------------------------:|:-------------------------:|:-------------------------:
abc                        |abcd                       |False
abc+                       |ab                         |False
abc*                       |abccccd                    |False
abc?                       |abcd                       |False
abc&#124;d                 |abcd                       |False
abc(a*&#124;d+)            |abddddeeee                 |False

### Menu Option 1 Sample

#### True 
If the user enters a regular expression which matches the entered string. </br>
<a href="https://imgur.com/M1l0RBF"><img src="https://imgur.com/M1l0RBF.png" title="Menu Option 1 True"/></a>

#### False
If the user enters a regular expression which doesn't match the entered string. </br>
<a href="https://imgur.com/GhiE2BG"><img src="https://imgur.com/GhiE2BG.png" title="Menu Option 1 False"/></a>

### Menu Option 2 Sample

#### True 
If the user enters a correct file name and the file contains strings which match the regular expression entered. </br>
<a href="https://imgur.com/pRJ1YgV"><img src="https://imgur.com/pRJ1YgV.png" title="Menu Option 2 True"/></a>

#### False
If the user enters a correct file name and the file doesn't contain any strings which match the regular expression entered. </br>
<a href="https://imgur.com/9K5DtcW"><img src="https://imgur.com/9K5DtcW.png" title="Menu Option 2 False"/></a>

#### File Not Found 
If the user enters the name of a file which doesn't exsist. </br>
<a href="https://imgur.com/rXYuoFD"><img src="https://imgur.com/rXYuoFD.png" title="Menu Option 2 No File"/></a>

### Menu Option 3 Sample

If the user enters enters '3' to exit the program. </br>
<a href="https://imgur.com/HNmqD5v"><img src="https://imgur.com/HNmqD5v.png" title="Menu Option 3"/></a>

### Invalid Menu Option

If the user doesn't enter one of the valid option. </br>
<a href="https://imgur.com/o9z8zlW"><img src="https://imgur.com/o9z8zlW.png" title="Menu Option Fail"/></a>

## Built With

* [Golang](https://golang.org/) - Golang is the code used in development
* [Visual Studio Code](https://code.visualstudio.com/) - A lightweight IDE which was used in development
* [Color package for Go](https://github.com/fatih/color) - Adds colour to GoLang output

## Authors

* **Cian Gannon** - *All source code* - [GitHub](https://github.com/cian2009)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

* **[StackOverflow](https://stackoverflow.com/)** - Used for light problem solving
* **[Table Of Contents Generator](https://ecotrust-canada.github.io/markdown-toc/)** - Used for auto generate the table of contents
