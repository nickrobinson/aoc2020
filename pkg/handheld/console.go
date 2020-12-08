package handheld

import (
	"errors"

	log "github.com/sirupsen/logrus"
)

type Instruction struct {
	Operation string
	Argument  int
}

type GameConsole struct {
	Accumulator        int
	InstructionPointer int
	SeenInstructions   map[int]bool
	Instructions       []Instruction
}

var InvalidOperationError = errors.New("Invalid Operation Type")

func NewGameConsole() GameConsole {
	return GameConsole{
		SeenInstructions: make(map[int]bool),
	}
}

func (g *GameConsole) AddInstruction(i *Instruction) error {
	if !isValidOperation(i.Operation) {
		log.Errorf("Invalid operation type: %s", i.Operation)
		return InvalidOperationError
	}
	g.Instructions = append(g.Instructions, *i)
	return nil
}

func (g *GameConsole) AddInstructions(instructions []Instruction) error {
	for _, i := range instructions {
		err := g.AddInstruction(&i)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *GameConsole) ExecuteProgram() {
	for {
		if g.InstructionPointer == len(g.Instructions) {
			log.Infof("Completed program execution")
			return
		}
		if _, ok := g.SeenInstructions[g.InstructionPointer]; ok {
			log.Infof("Exiting program execution as duplicate operation seen at %d", g.InstructionPointer)
			break
		}
		g.Step()
	}
}

func (g *GameConsole) Step() {
	var instruction Instruction
	if g.InstructionPointer == len(g.Instructions) {
		log.Infof("Completed program execution")
		return
	}
	switch instruction = g.Instructions[g.InstructionPointer]; instruction.Operation {
	case "nop":
		g.SeenInstructions[g.InstructionPointer] = true
		g.InstructionPointer++
	case "acc":
		g.SeenInstructions[g.InstructionPointer] = true
		g.Accumulator += instruction.Argument
		g.InstructionPointer++
	case "jmp":
		g.SeenInstructions[g.InstructionPointer] = true
		g.InstructionPointer += instruction.Argument
	default:
		log.Errorf("Unsupported operation type: %s", instruction.Operation)
	}
}

func isValidOperation(operation string) bool {
	switch operation {
	case
		"acc",
		"nop",
		"jmp":
		return true
	}
	return false
}
