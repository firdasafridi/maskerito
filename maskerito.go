package maskerito

import (
	"github.com/firdasafridi/maskerito/masker"
)

// MaskerEnergy is function that used for masking logical function
type MaskerEnergy func(input string) (output string)

// Maskerito is structure to hold masking energy
type Maskerito struct {
	maskers map[string]MaskerEnergy
}

// Config is used for add config on maskerito
type Config struct {
	// replacing or add logical masking
	Others map[string]MaskerEnergy

	// DefaultMask using *. This field using for change the default mask of maskerito
	DefaultMask string
}

// New create new maskerito data
func New(config *Config) (*Maskerito, error) {
	if config == nil {
		return nil, ErrInputNil
	}

	defaultMaskerito := defaultMaskerito()

	if config.DefaultMask != "" {
		masker.ChangeDefaultMask(config.DefaultMask)
	}

	// replacing  or add new logical masking
	for index, value := range config.Others {
		defaultMaskerito.maskers[index] = value
	}

	return defaultMaskerito, nil
}

// DefaultConfig will return default configuration
func DefaultConfig() *Config {
	return &Config{}
}

func defaultMaskerito() *Maskerito {
	maskers := map[string]MaskerEnergy{
		Mobile:  MaskerEnergy(masker.Mobile),
		Name:    MaskerEnergy(masker.Name),
		Email:   MaskerEnergy(masker.Email),
		ID:      MaskerEnergy(masker.ID),
		Bank:    MaskerEnergy(masker.BankAccountNumber),
		NPWP:    MaskerEnergy(masker.NPWP),
		SIUP:    MaskerEnergy(masker.SIUP),
		Address: MaskerEnergy(masker.Address),
	}

	return &Maskerito{
		maskers: maskers,
	}
}
