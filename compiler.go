package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func parseRegister(register string) uint8 {
	registers := map[string]uint8{
		"r0": 0x00,
		"r1": 0x01,
		"r2": 0x02,
		"r3": 0x03,
		"r4": 0x04,
		"r5": 0x05,
		"r6": 0x06,
		"r7": 0x07,
		"r8": 0x08,
		"r9": 0x09,
		"r10": 0x0A,
		"r11": 0x0B,
		"r12": 0x0C,
		"r13": 0x0D,
		"r14": 0x0E,
		"r15": 0x0F,
	}

	if val, ok := registers[register]; ok {
		return val
	} else {
		value, _ := strconv.Atoi(register)
		return uint8(value)
	}
}

func instParser(instruction string, token []string) []byte {
	var register string
	var value interface{}

	switch instruction {
	case "mov":
		register = strings.ReplaceAll(token[1], ",", "")
		value = token[2]

	case "add", "sub", "xor":
		register = strings.ReplaceAll(token[1], ",", "")
		value = token[2]

		if strValue, ok := value.(string); ok {
			value = parseRegister(strValue)
		}

	case "out", "inc", "dec":
		register = strings.ReplaceAll(token[1], ",", "")
		value = nil

	case "exit":
		return []byte{0x00}
	}

	if strValue, ok := value.(string); ok {
		value = parseRegister(strValue)
	}

	newReg := parseRegister(register)
	switch instruction {
	case "mov":
		return []byte{0x01, newReg, value.(uint8)}
	case "add":
		return []byte{0x02, newReg, value.(uint8)}
	case "sub":
		return []byte{0x03, newReg, value.(uint8)}
	case "out":
		return []byte{0x04, newReg}
	case "xor":
		return []byte{0x05, newReg, value.(uint8)}
	case "inc":
		return []byte{0x06, newReg}
	case "dec":
		return []byte{0x07, newReg}
	default:
		return nil
	}
}

func syntaxChecker(value string, lineNum int, position int) {
	instructions := map[string]bool{"mov": true, "add": true, "sub": true, "xor": true, "inc": true, "dec": true, "out": true, "exit": true}
	registers := map[string]bool{"r0": true, "r1": true, "r2": true, "r3": true, "r4": true, "r5": true, "r6": true, "r7": true, "r8": true, "r9": true, "r10": true, "r11": true, "r12": true, "r13": true, "r14": true, "r15": true}

	if !instructions[value] && !registers[value] {
		fmt.Printf("[line %d:%d] Syntax Error: %s\n", lineNum+1, position+1, value)
		os.Exit(1)
	}
}

func main() {

	filePtr := flag.String("f", "", "input file")
	flag.Parse()


	output := make([]byte, 0)

	file, err := ioutil.ReadFile(*filePtr)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	lines := strings.Split(string(file), "\n")

	for lineNum, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, ";") || line == "" {
			continue
		}

		tokens := strings.Split(line, " ")

		for i := 0; i < len(tokens); i++ {
			token := strings.ReplaceAll(tokens[i], ",", "")

			if intValue, err := strconv.Atoi(token); err != nil {
				syntaxChecker(token, lineNum, i)
			} else {
				tokens[i] = strconv.Itoa(intValue)
			}
		}

		code := strings.TrimSpace(line)
		tokens = strings.Split(code, " ")

		if code == "" || strings.HasPrefix(code, ";") {
			continue
		}

		instruction := tokens[0]
		newInst := instParser(instruction, tokens)
		output = append(output, newInst...)
	}


	fileName := *filePtr
	if strings.Contains(fileName, ".") {
		fileName = strings.Split(fileName, ".")[0]
	}

	bytecode := fmt.Sprintf("%s.bin", fileName)

	err = ioutil.WriteFile(bytecode, output, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}
}
