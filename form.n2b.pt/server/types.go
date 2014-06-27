package main

type AgentTransfer struct {
	OldName  string
	NewName  string
	OldNif   string
	NewNif   string
	Hardlock string
	Proof    []byte
}

type AgentUpgrade struct {
	Name     string
	Nif      string
	Software string
	Proof    []byte
}
