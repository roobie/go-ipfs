package commands

import (
	path "github.com/ipfs/go-ipfs/path"
)

type ArgumentType int

const (
	ArgString ArgumentType = iota
	ArgFile
	ArgPath
)

type Argument struct {
	Name          string
	Type          ArgumentType
	Required      bool // error if no value is specified
	Variadic      bool // unlimited values can be specfied
	SupportsStdin bool // can accept stdin as a value
	Recursive     bool // supports recursive file adding (with '-r' flag)
	Description   string
}

func StringArg(name string, required, variadic bool, description string) Argument {
	return Argument{
		Name:        name,
		Type:        ArgString,
		Required:    required,
		Variadic:    variadic,
		Description: description,
	}
}

func FileArg(name string, required, variadic bool, description string) Argument {
	return Argument{
		Name:        name,
		Type:        ArgFile,
		Required:    required,
		Variadic:    variadic,
		Description: description,
	}
}

func PathArg(name string, required, variadic bool, description string) Argument {
	return Argument{
		Name:        name,
		Type:        ArgPath,
		Required:    required,
		Variadic:    variadic,
		Description: description,
	}
}

type ArgumentValue struct {
	value String
	def   Argument
}

func (a ArgumentValue) File() (value string, found bool, err error) {
	if a.def != ArgFile {
		return "", false, util.ErrCast()
	}
	return a.value, len(a.value) > 0, nil
}

func (a ArgumentValue) String() (value string, found bool, err error) {
	if a.def != ArgString {
		return "", false, util.ErrCast()
	}
	return a.value, len(a.value) > 0, nil
}

func (a ArgumentValue) Path() (value path.Path, found bool, err error) {
	if a.def != ArgPath {
		return "", false, util.ErrCast()
	}
	val, err := path.ParsePath(a.value)

	return val, a.found, err
}

// TODO: modifiers might need a different API?
//       e.g. passing enum values into arg constructors variadically
//       (`FileArg("file", ArgRequired, ArgStdin, ArgRecursive)`)

func (a Argument) EnableStdin() Argument {
	a.SupportsStdin = true
	return a
}

func (a Argument) EnableRecursive() Argument {
	if a.Type != ArgFile {
		panic("Only ArgFile arguments can enable recursive")
	}

	a.Recursive = true
	return a
}
