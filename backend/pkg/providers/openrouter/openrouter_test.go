package openrouter

import (
	"testing"

	"pentagi/pkg/config"
	"pentagi/pkg/providers/pconfig"
	"pentagi/pkg/providers/provider"
)

func TestNewProvider(t *testing.T) {
	cfg := &config.Config{
		OpenRouterAPIKey:    "test-key",
		OpenRouterServerURL: "https://openrouter.ai/api/v1",
	}
	providerConfig, err := DefaultProviderConfig()
	if err != nil {
		t.Fatalf("DefaultProviderConfig() error = %v", err)
	}

	prov, err := New(cfg, provider.DefaultProviderNameOpenRouter, providerConfig)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	if prov.Name() != provider.DefaultProviderNameOpenRouter {
		t.Errorf("Expected provider name %v, got %v", provider.DefaultProviderNameOpenRouter, prov.Name())
	}
}

func TestProviderType(t *testing.T) {
	cfg := &config.Config{
		OpenRouterAPIKey:    "test-key",
		OpenRouterServerURL: "https://openrouter.ai/api/v1",
	}
	providerConfig, err := DefaultProviderConfig()
	if err != nil {
		t.Fatalf("DefaultProviderConfig() error = %v", err)
	}

	prov, err := New(cfg, provider.DefaultProviderNameOpenRouter, providerConfig)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	if prov.Type() != provider.ProviderOpenRouter {
		t.Errorf("Expected provider type %v, got %v", provider.ProviderOpenRouter, prov.Type())
	}
}

func TestMissingAPIKey(t *testing.T) {
	cfg := &config.Config{
		OpenRouterServerURL: "https://openrouter.ai/api/v1",
	}
	providerConfig, err := DefaultProviderConfig()
	if err != nil {
		t.Fatalf("DefaultProviderConfig() error = %v", err)
	}

	_, err = New(cfg, provider.DefaultProviderNameOpenRouter, providerConfig)
	if err == nil {
		t.Fatal("Expected error for missing API key")
	}
}

func TestModelWithPrefix(t *testing.T) {
	cfg := &config.Config{
		OpenRouterAPIKey:    "test-key",
		OpenRouterServerURL: "https://openrouter.ai/api/v1",
		OpenRouterProvider:  "openrouter",
	}
	providerConfig, err := DefaultProviderConfig()
	if err != nil {
		t.Fatalf("DefaultProviderConfig() error = %v", err)
	}

	prov, err := New(cfg, provider.DefaultProviderNameOpenRouter, providerConfig)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	if got := prov.ModelWithPrefix(pconfig.OptionsTypeSimple); got != "openrouter/free" {
		t.Errorf("Expected prefixed model, got %s", got)
	}
}
