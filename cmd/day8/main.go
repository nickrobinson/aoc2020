package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/nickrobinson/aoc2020/pkg/handheld"
	log "github.com/sirupsen/logrus"
)

func main() {
	fp, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(fp)
	//

	var instructions []handheld.Instruction

	for scanner.Scan() {
		lineText := scanner.Text()
		splitLine := strings.Split(lineText, " ")
		operation := splitLine[0]
		argument, _ := strconv.Atoi(splitLine[1])
		log.WithFields(log.Fields{"operation": operation, "argument": argument}).Info("Loading")
		instructions = append(instructions, handheld.Instruction{Operation: operation, Argument: argument})
	}
	// Part 1
	// console.AddInstructions(instructions)
	// console.ExecuteProgram()
	// log.Infof("Last acc value was: %d", console.Accumulator)

	log.Infof("Instruction length: %d", len(instructions))
	for i, instruction := range instructions {
		if instruction.Operation == "jmp" {
			log.Infof("Changing operation at %d from jmp->nop", i)
			instructions[i].Operation = "nop"
			console := handheld.NewGameConsole()
			console.AddInstructions(instructions)
			console.ExecuteProgram()
			if console.InstructionPointer == len(instructions) {
				log.Infof("Changing operation at %d from jmp->nop worked", i)
				log.Infof("Accumulator value is %d", console.Accumulator)
			}
			instructions[i].Operation = "jmp"
		} else if instruction.Operation == "nop" {
			log.Infof("Changing operation at %d from nop->jmp", i)
			instructions[i].Operation = "jmp"
			console := handheld.NewGameConsole()
			console.AddInstructions(instructions)
			console.ExecuteProgram()
			if console.InstructionPointer == len(instructions) {
				log.Infof("Changing operation at %d from nop->jmp worked", i)
				log.Infof("Accumulator value is %d", console.Accumulator)
			}
			instructions[i].Operation = "nop"
		}
	}
}
