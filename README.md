# Simple Register-based VM

I've created this simple Register-based VM only for learning purpose, I've also created a simple compiler to compile the script to bytecode, this simple VM has only 16 registers and few operators.

## Quick Start 

To build the VM and compiler, you can run the following code:
```bash
$ go mod init register-vm
$ go mod tidy
$ go build compiler.go
$ go build vm.go
```
Now you can compile the script and run the bytecode, which will have the same name as the script file but with the `.bin` extension added at the end.
```bash
$ ./compiler -f script
$ ./vm -f script.bin
```

## Compiler

To use the compiler, you just have to provide the script file, and it will create a bytecode file with the same name as the script.
```bash
$ ./compiler --help
Usage of ./compiler:
  -f string
    	input file
```

The compiler can also check for syntax errors. if it finds any errors, it will show the line and column where the errors occurs
```bash
$ ./compiler -f script
[line 1:1] Syntax Error: muv
```

## VM and Instructions 

I've implemented only a few instructions that allows us to perform basics mathematical operations, such as `add`, `xor`, `inc` etc...
```bash
add r1, r2 # add the value of the register r2 to register r1
sub r1, r2 # subtract the value of r2 from r1
xor r1, r2 # do a xor operation 
inc r1 # increment the register r1 in 1
dec r2 # decrement the register r2 in 1
out r1 # print the value of the r1
mov r1, 10 # move 10 to the register r1
exit # exit the program  
```

## Example of script 

I'll leave an example `script` here to show how the syntax accepted by compiler works
```bash
mov r0, 10
mov r1, 38

sub r1, 8
mov r0, r1

xor r1, 2

; This is a comment
;add r0, r1
;xor r0, 10
inc r0
dec r1

; print the value
out r0
out r1
exit
```
