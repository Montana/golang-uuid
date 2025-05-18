package uuid

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	t.Run("DefaultGeneration", func(t *testing.T) {
		uuid, err := Generate()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if len(uuid) != 36 {
			t.Errorf("Expected UUID length 36, got %d", len(uuid))
		}
		if err := Validate(uuid); err != nil {
			t.Errorf("Generated UUID is invalid: %v", err)
		}
	})
}

func TestGenerate_WithRandPool(t *testing.T) {
	EnableRandPool()
	defer DisableRandPool()

	uuid, err := Generate()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(uuid) != 36 {
		t.Errorf("Expected UUID of length 36, got %d", len(uuid))
	}
}

func TestMustGenerate(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Expected no panic, but got: %v", r)
		}
	}()
	uuid := MustGenerate()
	if len(uuid) != 36 {
		t.Errorf("Expected UUID of length 36, got %d", len(uuid))
	}
}

func TestValidate(t *testing.T) {
	validUUID := "550e8400-e29b-41d4-a716-446655440000"
	invalidUUID := "550e8400-e29b-41d4-a716-44665544000X"
	shortUUID := "550e8400-e29b"

	if err := Validate(validUUID); err != nil {
		t.Errorf("Expected valid UUID, got error: %v", err)
	}
	if err := Validate(invalidUUID); err == nil {
		t.Errorf("Expected error for invalid UUID, got none")
	}
	if err := Validate(shortUUID); err == nil {
		t.Errorf("Expected error for short UUID, got none")
	}
}

func TestToBinary(t *testing.T) {
	uuidStr := "550e8400-e29b-41d4-a716-446655440000"
	bin, err := ToBinary(uuidStr)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(bin) != 16 {
		t.Errorf("Expected binary UUID length 16, got %d", len(bin))
	}
}

func TestToBinary_Invalid(t *testing.T) {
	_, err := ToBinary("invalid-uuid")
	if err == nil {
		t.Errorf("Expected error for invalid UUID string")
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
	if parts[2][0] != '4' {
		t.Errorf("Expected version 4 UUID, got: %s", parts[2])
	}
}

func TestGetParts_Invalid(t *testing.T) {
	_, err := GetParts("invalid-uuid")
	if err == nil {
		t.Errorf("Expected error for invalid UUID")
	}
}
