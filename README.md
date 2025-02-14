# golang uuid 

## Overview

This Golang module provides functions for generating, validating, formatting, and converting UUIDs. The implementation follows the UUID v4 standard.

## Features
- Generate UUID v4
- Validate UUID format
- Convert UUID to binary
- Extract UUID components

## Usage

### Generate a UUID
```go
uuid, err := uuid.Generate()
if err != nil {
    log.Fatal(err)
}
fmt.Println("Generated UUID:", uuid)
```

### Validate a UUID
```go
err := uuid.Validate("550e8400-e29b-41d4-a716-446655440000")
if err != nil {
    fmt.Println("Invalid UUID")
} else {
    fmt.Println("Valid UUID")
}
```

### Convert UUID to Binary
```go
binaryUUID, err := uuid.ToBinary("550e8400-e29b-41d4-a716-446655440000")
if err != nil {
    log.Fatal(err)
}
fmt.Println("Binary UUID:", binaryUUID)
```

### Extract Parts from UUID
```go
parts, err := uuid.GetParts("550e8400-e29b-41d4-a716-446655440000")
if err != nil {
    log.Fatal(err)
}
fmt.Println("UUID Parts:", parts)
```

## Author 

Michael Mendy (c) 2025.
