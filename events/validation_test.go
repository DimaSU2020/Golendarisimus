package events

import (
	// "fmt"
	"testing"
)

func TestIsValidPriority(t *testing.T) {
	p := PriorityLow
	err := IsValidPriority(p)
	if err != nil {
		t.Errorf("Expected no error for PriorityLow, got %v", err)
	}

	p  = PriorityMedium
	err  = IsValidPriority(p)
	if err != nil {
		t.Errorf("Expected no error for PriorityMedium, got %v", err)
	}
	
	p  = PriorityHigh
	err  = IsValidPriority(p)
	if err != nil {
		t.Errorf("Expected no error for PriorityHigh, got %v", err)
	}

	p = "LOW"
	err = IsValidPriority(p)
	if err != nil {
		t.Errorf("Expected no error for PriorityLow, got %v", err)
	}

	p = "MeDiUm"
	err = IsValidPriority(p)
	if err != nil {
		t.Errorf("Expected no error for PriorityMedium, got %v", err)
	}

	p = "higH"
	err = IsValidPriority(p)
	if err != nil {
		t.Errorf("Expected no error for PriorityHigh, got %v", err)
	}

	p = "urgent"
	err  = IsValidPriority(p)
	if err == nil {
		t.Errorf("Expected an error for \"%v\" priority, got none",p)
	}

	p = "срочно"
	err  = IsValidPriority(p)
	if err == nil {
		t.Errorf("Expected an error for \"%v\" priority, got none",p)
	}

	p = "несрочно"
	err  = IsValidPriority(p)
	if err == nil {
		t.Errorf("Expected an error for \"%v\" priority, got none",p)
	}

	p = "l ow"
	err  = IsValidPriority(p)
	if err == nil {
		t.Errorf("Expected an error for \"%v\" priority, got none",p)
	}
}

func TestIsValidTitle(t *testing.T) {
	title := "Правильный заголовок"
	matched, _ := IsValidTitle(title)
	if !matched {
		t.Errorf("Expected no error for \"%v\", got error",title)
	}

	title = "Заголовок с ,.#"
	matched, _ = IsValidTitle(title)
	if !matched {
		t.Errorf("Expected no error for \"%v\", got error",title)
	}

	title = "Заголовок с символами @!"
	matched, _ = IsValidTitle(title)
	if matched {
		t.Errorf("Expected an error for \"%v\" special sign, got none",title)
	}
}	

