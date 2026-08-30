package component

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestComponentEncoders(t *testing.T) {
	tests := []struct {
		name     string
		comp     Component
		wantKeys []string
	}{
		{
			name: "wearable",
			comp: Wearable{
				Slot:                SlotArmorHead,
				Protection:          3,
				HidesPlayerLocation: true,
				Dispensable:         false,
			},
			wantKeys: []string{"slot", "protection", "hides_player_location", "dispensable"},
		},
		{
			name: "food",
			comp: Food{
				Nutrition:          4,
				SaturationModifier: 0.6,
				CanAlwaysEat:       false,
			},
			wantKeys: []string{"nutrition", "saturation_modifier", "can_always_eat"},
		},
		{
			name: "durability",
			comp: Durability{
				MaxDurability: 100,
				DamageChance:  [2]int32{1, 100},
			},
			wantKeys: []string{"max_durability", "damage_chance"},
		},
		{
			name: "kinetic_weapon",
			comp: KineticWeapon{
				Reach:            [2]float32{3, 4},
				CreativeReach:    [2]float32{5, 6},
				HitboxMargin:     0.5,
				DamageMultiplier: 1.5,
				DamageModifier:   2.0,
				Delay:            10,
				DamageConditions: WeaponConditions{
					MaxDuration:      100,
					MinSpeed:         0.1,
					MinRelativeSpeed: 0,
				},
			},
			wantKeys: []string{"creative_reach", "damage_conditions", "damage_modifier", "damage_multiplier", "delay", "hitbox_margin", "knockback_conditions", "reach"},
		},
		{
			name: "use_modifiers",
			comp: UseModifiers{
				MovementModifier: 0.2,
				UseDuration:      1.0,
				EmitVibrations:   true,
			},
			wantKeys: []string{"movement_modifier", "use_duration", "emit_vibrations"},
		},
		{
			name: "storage_item",
			comp: StorageItem{
				AllowNestedStorageItems: false,
				AllowedItems:            []string{"minecraft:diamond"},
				BannedItems:             []string{"minecraft:bedrock"},
				MaxSlots:                27,
			},
			wantKeys: []string{"allow_nested_storage_items", "allowed_items", "banned_items", "max_slots"},
		},
		{
			name: "shooter",
			comp: Shooter{
				Ammunition: []Ammunition{
					{Item: "minecraft:arrow", SearchInventory: true, UseInCreative: true, UseOffHand: true},
				},
				ChargeOnDraw:             true,
				MaxDrawDuration:          1.0,
				ScalePowerByDrawDuration: true,
			},
			wantKeys: []string{"ammunition", "charge_on_draw", "max_draw_duration", "scale_power_by_draw_duration"},
		},
		{
			name:     "damage",
			comp:     Damage{Value: 6.0},
			wantKeys: []string{"value"},
		},
		{
			name: "cooldown",
			comp: Cooldown{
				Category: "test",
				Duration: 5.0,
			},
			wantKeys: []string{"category", "duration"},
		},
		{
			name: "enchantable",
			comp: Enchantable{
				Slot:  "sword",
				Value: 10,
			},
			wantKeys: []string{"slot", "value"},
		},
		{
			name: "repairable",
			comp: Repairable{
				RepairItems: []RepairEntry{
					{Items: []RepairItemEntry{{Name: "minecraft:diamond"}}, RepairAmount: 100},
				},
			},
			wantKeys: []string{"repair_items"},
		},
		{
			name:     "item_tags",
			comp:     ItemTags{Tags: []string{"tag1", "tag2"}},
			wantKeys: []string{"tags"},
		},
		{
			name: "seed",
			comp: Seed{
				CropResult:             "minecraft:wheat",
				PlantAt:                []string{"minecraft:farmland"},
				PlantAtAnySolidSurface: false,
			},
			wantKeys: []string{"crop_result", "plant_at", "plant_at_any_solid_surface"},
		},
		{
			name:     "fuel",
			comp:     Fuel{Duration: 100},
			wantKeys: []string{"duration"},
		},
		{
			name:     "fire_resistant",
			comp:     FireResistant{Value: true},
			wantKeys: []string{"value"},
		},
		{
			name:     "glint",
			comp:     Glint{Value: true},
			wantKeys: []string{"value"},
		},
		{
			name:     "swing_duration",
			comp:     SwingDuration{Value: 0.5},
			wantKeys: []string{"value"},
		},
		{
			name: "swing_sounds",
			comp: SwingSounds{
				AttackCriticalHit: "sound1",
				AttackHit:         "sound2",
				AttackMiss:        "sound3",
			},
			wantKeys: []string{"attack_critical_hit", "attack_hit", "attack_miss"},
		},
		{
			name: "piercing_weapon",
			comp: PiercingWeapon{
				CreativeReach: [2]float32{5, 6},
				HitboxMargin:  0.5,
				Reach:         [2]float32{3, 4},
			},
			wantKeys: []string{"creative_reach", "hitbox_margin", "reach"},
		},
		{
			name: "camera",
			comp: Camera{
				BlackBarsDuration:    0.5,
				BlackBarsScreenRatio: 0.3,
				PictureDuration:      1.0,
				ShutterDuration:      0.1,
				ShutterScreenRatio:   0.5,
				SlideAwayDuration:    0.5,
			},
			wantKeys: []string{"black_bars_duration", "black_bars_screen_ratio", "picture_duration", "shutter_duration", "shutter_screen_ratio", "slide_away_duration"},
		},
		{
			name: "block_placer",
			comp: BlockPlacer{
				Block:            "minecraft:dirt",
				ReplaceBlockItem: "minecraft:dirt",
				AlignedPlacement: true,
				UseOn:            []string{"minecraft:grass_block"},
			},
			wantKeys: []string{"block", "replace_block_item", "aligned_placement", "use_on"},
		},
		{
			name:     "compostable",
			comp:     Compostable{CompostingChance: 50},
			wantKeys: []string{"composting_chance"},
		},
		{
			name:     "damage_absorption",
			comp:     DamageAbsorption{AbsorbableCauses: []string{"fire", "explosion"}},
			wantKeys: []string{"absorbable_causes"},
		},
		{
			name: "digger",
			comp: Digger{
				DestroySpeeds: []DestroySpeed{{Block: "minecraft:stone", Speed: 5.0}},
				UseEfficiency: true,
			},
			wantKeys: []string{"destroy_speeds", "use_efficiency"},
		},
		{
			name: "durability_sensor",
			comp: DurabilitySensor{
				SoundEvent: "sound1",
				DurabilityThresholds: []DurabilityThreshold{
					{Durability: 50, ParticleType: "particle1", SoundEvent: "sound2"},
				},
			},
			wantKeys: []string{"sound_event", "durability_thresholds"},
		},
		{
			name:     "dyeable",
			comp:     Dyeable{DefaultColor: [3]int32{255, 0, 0}},
			wantKeys: []string{"default_color"},
		},
		{
			name: "entity_placer",
			comp: EntityPlacer{
				Entity:     "minecraft:pig",
				UseOn:      []string{"minecraft:grass_block"},
				DispenseOn: []string{"minecraft:dirt"},
			},
			wantKeys: []string{"entity", "use_on", "dispense_on"},
		},
		{
			name:     "hover_text_color",
			comp:     HoverTextColor{Value: 0xFF0000},
			wantKeys: []string{"value"},
		},
		{
			name:     "interact_button",
			comp:     InteractButton{Value: "test"},
			wantKeys: []string{"value"},
		},
		{
			name:     "liquid_clipped",
			comp:     LiquidClipped{Value: true},
			wantKeys: []string{"value"},
		},
		{
			name:     "rarity",
			comp:     Rarity{Value: "epic"},
			wantKeys: []string{"value"},
		},
		{
			name: "record",
			comp: Record{
				ComparatorSignal: 1,
				Duration:         10.0,
				SoundEvent:       "music.record.test",
			},
			wantKeys: []string{"comparator_signal", "duration", "sound_event"},
		},
		{
			name:     "should_despawn",
			comp:     ShouldDespawn{Value: true},
			wantKeys: []string{"value"},
		},
		{
			name:     "bundle_interaction",
			comp:     BundleInteraction{NumViewableSlots: 12},
			wantKeys: []string{"num_viewable_slots"},
		},
		{
			name:     "storage_weight_limit",
			comp:     StorageWeightLimit{MaxWeightLimit: 100},
			wantKeys: []string{"max_weight_limit"},
		},
		{
			name:     "storage_weight_modifier",
			comp:     StorageWeightModifier{WeightInStorageItem: 5},
			wantKeys: []string{"weight_in_storage_item"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data, err := tc.comp.Encode()
			require.NoError(t, err)
			require.Equal(t, "minecraft:"+tc.name, tc.comp.ComponentName())
			for _, k := range tc.wantKeys {
				require.Contains(t, data, k, "missing key %s in component %s", k, tc.name)
			}
		})
	}
}

func TestRawComponent(t *testing.T) {
	raw := RawComponent{
		Name: "minecraft:test",
		Data: map[string]any{"key": "value"},
	}
	require.Equal(t, "minecraft:test", raw.ComponentName())
	data, err := raw.Encode()
	require.NoError(t, err)
	require.Equal(t, map[string]any{"key": "value"}, data)
}
