package main

// =============================================================================
// LESSON 01: PACKAGE DECLARATION & IMPORTS
// =============================================================================
// Concept: Every Go program starts with a package declaration and imports the
// packages it needs. This is the foundation of Go program structure.
//
// 1. package main    → tells Go this is an executable (not a library)
// 2. import (...)    → brings in functionality from Go's standard library
//
// Standard library packages used across the project:
//   - "fmt"      : formatted I/O (printing to console)          [used in: 20]
//   - "math"     : mathematical functions (math.Ceil)            [used in: 13]
//   - "strconv"  : string ↔ number conversions (Atoi, Itoa)      [used in: 11,13,19]
//   - "strings"  : string manipulation (Split, Builder)          [used in: 11,13,19]
//   - "time"     : date/time handling (Date, Weekday, AddDate)   [used in: 13,16,17]
//
// In Go, each FILE declares its own imports (not the package).
// This means if file A uses "strings", it imports "strings" even if
// file B in the same package also uses it. This keeps files self-contained.
// =============================================================================
