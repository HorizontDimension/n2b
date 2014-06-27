package main

type AgentTransfer struct {
	OldName  string `schema:"OldName"`
	NewName  string `schema:"NewName"`
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
