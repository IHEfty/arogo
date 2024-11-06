# AroGO - Molecule JSON Encoder and Decoder

This Go package provides a simple structure for encoding and decoding chemical molecules to and from JSON format. It represents molecules with atoms and bonds, making it easy to serialize molecular structures for storage or transmission and then deserialize them back into usable data with restored atom-to-bond relationships.

## Overview

The code defines three main types:
1. **Atom**: Represents an atom with properties such as symbol, atomic number, and optional charge.
2. **Bond**: Represents a bond between two atoms, specifying the type of bond (e.g., single, double) and linking two atoms by their IDs.
3. **Molecule**: Represents the whole molecule, containing a list of atoms, a list of bonds, and optional fields for the molecule's name and formula.

The package includes functions to:
- **Encode** a `Molecule` to JSON (`encodeMoleculeToJSON`) with references to atoms set as IDs for simple serialization.
- **Decode** JSON back into a `Molecule` structure (`decodeJSONToMolecule`), restoring atom pointers within bonds.
- Convert the molecule into a human-readable **formula** and **name**.

## Why Use This Package?

1. **Data Interchange**: Provides a JSON format that can be shared or stored, making molecular data easy to transmit between systems.
2. **Structure Restoration**: JSON decoding restores all necessary relationships between atoms and bonds, essential for applications that require atomic-level detail.
3. **Easy to Extend**: This package can be extended for more complex molecules and additional properties, making it useful for chemistry-related applications.

## Code Structure

- **`Atom` struct**: Represents individual atoms with attributes like symbol and atomic number.
- **`Bond` struct**: Represents bonds between atoms, with fields to store atom IDs and bond types.
- **`Molecule` struct**: Represents a molecule with atoms and bonds, and includes optional name and formula fields.

### Functions

- **`toFormula`**: Sets or retrieves the formula of the molecule.
- **`toName`**: Sets or retrieves the name of the molecule.
- **`encodeMoleculeToJSON`**: Converts a `Molecule` struct to JSON format.
- **`decodeJSONToMolecule`**: Parses JSON data and restores atom-bond relationships.

## How It Works

1. **Define Atoms and Bonds**: Atoms are defined with unique IDs, symbols, and atomic numbers. Bonds are created between pairs of atoms.
2. **Build the Molecule**: Construct a `Molecule` with the defined atoms and bonds.
3. **Encode to JSON**: The `encodeMoleculeToJSON` function serializes the `Molecule` to JSON, saving the atom-bond relationships using IDs.
4. **Decode from JSON**: The `decodeJSONToMolecule` function deserializes JSON back to a `Molecule` structure, restoring atom-to-bond relationships.

## Running the Code

1. **Install Go**: Ensure Go is installed on your machine.
2. **Save the Code**: Save the provided code into a file, e.g., `molecule.go`.
3. **Run the Program**:
   ```bash
   go run molecule.go
   ```

## Example Output

When run, the program will output the JSON representation of a water molecule (H₂O), then decode it back and print the structure line by line.

Example output:
```
Encoded JSON:
{
  "atoms": [
    { "id": 1, "symbol": "H", "atomic_number": 1 },
    { "id": 2, "symbol": "H", "atomic_number": 1 },
    { "id": 3, "symbol": "O", "atomic_number": 8 }
  ],
  "bonds": [
    { "atom1_id": 1, "atom2_id": 3, "bond_type": "single" },
    { "atom1_id": 2, "atom2_id": 3, "bond_type": "single" }
  ],
  "name": "Water",
  "formula": "H2O"
}

Decoded Molecule:
Name: Water
Formula: H2O
Atoms:
  ID: 1, Symbol: H, Atomic Number: 1, Charge: 0
  ID: 2, Symbol: H, Atomic Number: 1, Charge: 0
  ID: 3, Symbol: O, Atomic Number: 8, Charge: 0
Bonds:
  Bond between H (ID: 1) and O (ID: 3), Type: single
  Bond between H (ID: 2) and O (ID: 3), Type: single
```

## Notes

- **Error Handling**: The code includes error handling for encoding and decoding to ensure robustness.
- **Customization**: You can extend `Atom`, `Bond`, and `Molecule` structures with additional fields (e.g., bond angles, atom positions) for more detailed molecular representation.
