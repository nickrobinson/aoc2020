package handheld

import (
	"testing"
)

func TestAddingInstructions(t *testing.T) {
	var tests = []struct {
		name        string
		operation   string
		argument    int
		accumulator int
		retVal      error
	}{
		{"nop operation", "nop", 0, 0, nil},
		{"acc operation", "acc", 10, 10, nil},
		{"jmp operation", "jmp", -3, 0, nil},
		{"foo operation", "foo", -3, 0, InvalidOperationError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			console := NewGameConsole()
			if console.AddInstruction(&Instruction{Operation: tt.operation, Argument: tt.argument}) != tt.retVal {
				t.Errorf("Detected invalid operation")
			}
		})
	}
}

func TestProgramStep(t *testing.T) {
	console := NewGameConsole()
	console.AddInstruction(&Instruction{Operation: "nop", Argument: 0})
	console.AddInstruction(&Instruction{Operation: "acc", Argument: 1})
	console.AddInstruction(&Instruction{Operation: "jmp", Argument: 4})
	console.AddInstruction(&Instruction{Operation: "acc", Argument: 3})
	console.AddInstruction(&Instruction{Operation: "jmp", Argument: -3})
	console.AddInstruction(&Instruction{Operation: "acc", Argument: -99})
	console.AddInstruction(&Instruction{Operation: "acc", Argument: 1})
	console.AddInstruction(&Instruction{Operation: "jmp", Argument: -4})
	console.AddInstruction(&Instruction{Operation: "acc", Argument: 6})
	console.Step()
	if console.InstructionPointer != 1 {
		t.Errorf("Expected Instruction Pointer to be 1, was %d", console.InstructionPointer)
	}
	console.Step()
	if console.InstructionPointer != 2 {
		t.Errorf("Expected Instruction Pointer to be 2, was %d", console.InstructionPointer)
	}
	if console.Accumulator != 1 {
		t.Errorf("Expected Accumulator to be 1, was %d", console.Accumulator)
	}
	console.Step()
	if console.InstructionPointer != 6 {
		t.Errorf("Expected Instruction Pointer to be 6, was %d", console.InstructionPointer)
	}
}

func TestExecutingProgram(t *testing.T) {
	console := NewGameConsole()
	console.AddInstruction(&Instruction{Operation: "nop", Argument: 0})
	console.AddInstruction(&Instruction{Operation: "acc", Argument: 1})
	console.AddInstruction(&Instruction{Operation: "jmp", Argument: 4})
	console.AddInstruction(&Instruction{Operation: "acc", Argument: 3})
	console.AddInstruction(&Instruction{Operation: "jmp", Argument: -3})
	console.AddInstruction(&Instruction{Operation: "acc", Argument: -99})
	console.AddInstruction(&Instruction{Operation: "acc", Argument: 1})
	console.AddInstruction(&Instruction{Operation: "jmp", Argument: -4})
	console.AddInstruction(&Instruction{Operation: "acc", Argument: 6})
	console.ExecuteProgram()
	if console.Accumulator != 5 {
		t.Errorf("Accumulator value at end of execution not correct. Exepected 5, got %d", console.Accumulator)
	}
}
