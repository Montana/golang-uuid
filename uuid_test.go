package uuid

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	uuid, err := Generate()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(uuid) != 36 {
		t.Errorf("Expected UUID of length 36, got %d", len(uuid))
	}
}

func TestValidate(t *testing.T) {
	validUUID := "550e8400-e29b-41d4-a716-446655440000"
	invalidUUID := "550e8400-e29b-41d4-a716-44665544000X"

	if err := Validate(validUUID); err != nil {
		t.Errorf("Expected valid UUID, got error: %v", err)
	}

	if err := Validate(invalidUUID); err == nil {
		t.Errorf("Expected error for invalid UUID, got none")
	}
}

func TestToBinary(t *testing.T) {
	uuidStr := "550e8400-e29b-41d4-a716-446655440000"
	binaryUUID, err := ToBinary(uuidStr)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(binaryUUID) != 16 {
		t.Errorf("Expected binary UUID of length 16, got %d", len(binaryUUID))
	}
}

func TestGetParts(t *testing.T) {
	uuidStr := "550e8400-e29b-41d4-a716-446655440000"
	parts, err := GetParts(uuidStr)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(parts) != 5 {
		t.Errorf("Expected 5 parts in UUID, got %d", len(parts))
	}
}
