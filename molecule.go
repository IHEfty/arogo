package main

import (
	"encoding/json"
	"fmt"
)

type Atom struct {
	ID        int    `json:"id"`
	Symbol    string `json:"symbol"`
	AtomicNum int    `json:"atomic_number"`
	Charge    int    `json:"charge,omitempty"`
}

type Bond struct {
	Atom1ID  int    `json:"atom1_id"`
	Atom2ID  int    `json:"atom2_id"`
	BondType string `json:"bond_type"`
	Atom1    *Atom  `json:"-"`
	Atom2    *Atom  `json:"-"`
}

type Molecule struct {
	Atoms   []*Atom `json:"atoms"`
	Bonds   []*Bond `json:"bonds"`
	Name    string  `json:"name"`
	Formula string  `json:"formula"`
}

func (m *Molecule) toFormula() string {
	if m.Formula == "" {
		m.Formula = "H2O" 
	}
	return m.Formula
}

func (m *Molecule) toName() string {
	if m.Name == "" {
		m.Name = "Water" 
	}
	return m.Name
}

func encodeMoleculeToJSON(m *Molecule) (string, error) {
	for _, bond := range m.Bonds {
		bond.Atom1ID = bond.Atom1.ID
		bond.Atom2ID = bond.Atom2.ID
	}
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func decodeJSONToMolecule(jsonData string) (*Molecule, error) {
	var molecule Molecule
	err := json.Unmarshal([]byte(jsonData), &molecule)
	if err != nil {
		return nil, err
	}

	atomMap := make(map[int]*Atom)
	for _, atom := range molecule.Atoms {
		atomMap[atom.ID] = atom
	}

	for _, bond := range molecule.Bonds {
		bond.Atom1 = atomMap[bond.Atom1ID]
		bond.Atom2 = atomMap[bond.Atom2ID]
	}

	return &molecule, nil
}

func main() {
	h1 := &Atom{ID: 1, Symbol: "H", AtomicNum: 1}
	h2 := &Atom{ID: 2, Symbol: "H", AtomicNum: 1}
	o := &Atom{ID: 3, Symbol: "O", AtomicNum: 8}

	bond1 := &Bond{Atom1: h1, Atom2: o, BondType: "single"}
	bond2 := &Bond{Atom1: h2, Atom2: o, BondType: "single"}

	water := &Molecule{
		Atoms: []*Atom{h1, h2, o},
		Bonds: []*Bond{bond1, bond2},
	}
	water.toFormula()
	water.toName()

	jsonData, err := encodeMoleculeToJSON(water)
	if err != nil {
		fmt.Println("Error encoding to JSON:", err)
		return
	}
	fmt.Println("Encoded JSON:\n", jsonData)

	decodedMolecule, err := decodeJSONToMolecule(jsonData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Decoded Molecule:")
	fmt.Printf("Name: %s\n", decodedMolecule.Name)
	fmt.Printf("Formula: %s\n", decodedMolecule.Formula)
	fmt.Println("Atoms:")
	for _, atom := range decodedMolecule.Atoms {
		fmt.Printf("  ID: %d, Symbol: %s, Atomic Number: %d, Charge: %d\n",
			atom.ID, atom.Symbol, atom.AtomicNum, atom.Charge)
	}
	fmt.Println("Bonds:")
	for _, bond := range decodedMolecule.Bonds {
		fmt.Printf("  Bond between %s (ID: %d) and %s (ID: %d), Type: %s\n",
			bond.Atom1.Symbol, bond.Atom1.ID, bond.Atom2.Symbol, bond.Atom2.ID, bond.BondType)
	}
}
