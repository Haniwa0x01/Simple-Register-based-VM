package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"os"
)

type CPU struct {
	memory    [256]int
	registers [16]int
	pc        int
}

func (c *CPU) load(program []int) {
	for i, instruction := range program {
		c.memory[i] = instruction
	}
}

func (c * CPU) add(operand1, operand2 int) int {
	if operand2 > len(c.registers) || c.registers[operand2] == 0 {
		return c.registers[operand1] + operand2
	}
	return c.registers[operand1] + c.registers[operand2]
}

func (c * CPU) sub(operand1, operand2 int) int {
	if operand2 > len(c.registers) || c.registers[operand2] == 0 {
		return c.registers[operand1] - operand2
	}
	return c.registers[operand1] - c.registers[operand2]
}

func (c * CPU) xor(operand1, operand2 int) int {
	if operand2 > len(c.registers) || c.registers[operand2] == 0 {
		return c.registers[operand1] ^ operand2
	}
	return c.registers[operand1] ^ c.registers[operand2]
}


func (c *CPU) run() {
	for {
		instruction := c.memory[c.pc]

		switch instruction {
			case 0x01: // 0x01 = MOV
				operand := c.memory[c.pc+1]
				value := c.memory[c.pc+2]
				c.registers[operand] = value
				c.pc += 3
			case 0x02: // ADD
				operand1 := c.memory[c.pc+1]
				operand2 := c.memory[c.pc+2]
				result := c.add(operand1, operand2)
				c.registers[operand1] = result
				c.pc += 3

			case 0x03: // 0x03 = SUB
				operand1 := c.memory[c.pc+1]
				operand2 := c.memory[c.pc+2]
				result := c.sub(operand1, operand2)
				c.registers[operand1] = result
				c.pc += 3

			case 0x04: // 0x04 = OUT (num)
				operand := c.memory[c.pc+1]
				value := c.registers[operand]
				fmt.Println(value)
				c.pc += 2
			case 0x05: // 0x05 = XOR
				operand1 := c.memory[c.pc+1]
				operand2 := c.memory[c.pc+2]
				result := c.xor(operand1, operand2)
				c.registers[operand1] = result
				c.pc += 3

			case 0x06: // 0x06 = INC
				operand := c.memory[c.pc+1]
				result := c.registers[operand] + 1
				c.registers[operand] = result
				c.pc += 2
			case 0x07: // 0x07 = DEC
				operand := c.memory[c.pc+1]
				result := c.registers[operand] - 1
				c.registers[operand] = result
				c.pc += 2

			case 0x00: // 0x00 = EXIT
				os.Exit(0)
				break
			default:
				fmt.Printf("Unknown instruction: %d\n", instruction)
				os.Exit(0)
		}
	}
}

func main() {
	filePtr := flag.String("f", "", "input file")
	flag.Parse()

	fileContent, err := ioutil.ReadFile(*filePtr)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	program := make([]int, len(fileContent))
	for i, byteVal := range fileContent {
		program[i] = int(byteVal)
	}

	vm := CPU{}
	vm.load(program)
	vm.run()
}
