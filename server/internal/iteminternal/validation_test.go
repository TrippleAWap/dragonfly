package iteminternal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidation(t *testing.T) {
	tests := []struct {
		name       string
		components map[string]any
		wantErr    bool
		errMsg     string
	}{
		{
			name: "kinetic_weapon without use_modifiers",
			components: map[string]any{
				"minecraft:kinetic_weapon": map[string]any{
					"damage_conditions": map[string]any{"max_duration": int32(100)},
				},
			},
			wantErr: true,
			errMsg:  "minecraft:kinetic_weapon requires minecraft:use_modifiers component",
		},
		{
			name: "kinetic_weapon with use_modifiers",
			components: map[string]any{
				"minecraft:kinetic_weapon": map[string]any{
					"damage_conditions": map[string]any{"max_duration": int32(100)},
				},
				"minecraft:use_modifiers": map[string]any{
					"use_duration":      float32(1.0),
					"movement_modifier": float32(0.2),
				},
			},
			wantErr: false,
		},
		{
			name: "kinetic_weapon with invalid conditions (both min_speed and min_relative_speed)",
			components: map[string]any{
				"minecraft:kinetic_weapon": map[string]any{
					"damage_conditions": map[string]any{
						"max_duration":       int32(100),
						"min_speed":          float32(0.1),
						"min_relative_speed": float32(0.2),
					},
				},
				"minecraft:use_modifiers": map[string]any{
					"use_duration": float32(1.0),
				},
			},
			wantErr: true,
			errMsg:  "min_speed and min_relative_speed are mutually exclusive",
		},
		{
			name: "kinetic_weapon with zero max_duration",
			components: map[string]any{
				"minecraft:kinetic_weapon": map[string]any{
					"damage_conditions": map[string]any{
						"max_duration": int32(0),
					},
				},
				"minecraft:use_modifiers": map[string]any{
					"use_duration": float32(1.0),
				},
			},
			wantErr: true,
			errMsg:  "at least one condition with max_duration > 0 required",
		},
		{
			name: "bundle_interaction without storage_item",
			components: map[string]any{
				"minecraft:bundle_interaction": map[string]any{
					"num_viewable_slots": int32(12),
				},
			},
			wantErr: true,
			errMsg:  "minecraft:bundle_interaction requires minecraft:storage_item component",
		},
		{
			name: "bundle_interaction with storage_item",
			components: map[string]any{
				"minecraft:bundle_interaction": map[string]any{
					"num_viewable_slots": int32(12),
				},
				"minecraft:storage_item": map[string]any{
					"max_slots": int32(27),
				},
			},
			wantErr: false,
		},
		{
			name: "shooter without use_modifiers",
			components: map[string]any{
				"minecraft:shooter": map[string]any{
					"ammunition": []any{},
				},
			},
			wantErr: true,
			errMsg:  "minecraft:shooter requires minecraft:use_modifiers component with non-zero use_duration",
		},
		{
			name: "shooter with use_modifiers but zero use_duration",
			components: map[string]any{
				"minecraft:shooter": map[string]any{
					"ammunition": []any{},
				},
				"minecraft:use_modifiers": map[string]any{
					"use_duration": float32(0),
				},
			},
			wantErr: true,
			errMsg:  "minecraft:shooter requires non-zero use_duration in minecraft:use_modifiers",
		},
		{
			name: "shooter with valid use_modifiers",
			components: map[string]any{
				"minecraft:shooter": map[string]any{
					"ammunition": []any{},
				},
				"minecraft:use_modifiers": map[string]any{
					"use_duration": float32(1.0),
				},
			},
			wantErr: false,
		},
		{
			name: "repairable with invalid repair_items",
			components: map[string]any{
				"minecraft:repairable": map[string]any{
					"repair_items": []any{
						map[string]any{"invalid": "entry"},
					},
				},
			},
			wantErr: true,
			errMsg:  "minecraft:repairable repair_items must have 'items' field",
		},
		{
			name: "repairable with valid repair_items",
			components: map[string]any{
				"minecraft:repairable": map[string]any{
					"repair_items": []any{
						map[string]any{"items": []any{map[string]any{"name": "minecraft:diamond"}}, "repair_amount": int32(100)},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateComponents(tc.components)
			if tc.wantErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.errMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
