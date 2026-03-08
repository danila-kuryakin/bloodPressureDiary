package constants

type State string

// menu states
const (
	StateIdle     State = ""
	StatePressure State = "state_pressure"
	StateTag      State = "state_tag"
)

// pressure
const (
	StatePressureCreate State = "state_create"
	StatePressureEdit   State = "state_pressure_edit"
	StatePressureDelete State = "state_pressure_delete"

	StatePressureSys   State = "state_pressure_sys"
	StatePressureDia   State = "state_pressure_dia"
	StatePressurePulse State = "state_pressure_pulse"
	StatePressureTags  State = "state_pressure_tags"

	StatePressureDiaFast   State = "state_pressure_dia_fast"
	StatePressurePulseFast State = "state_pressure_pulse_fast"
	StatePressureTagsFast  State = "state_pressure_tags_fast"

	StatePressureEditSys   State = "state_pressure_edit_sys"
	StatePressureEditDia   State = "state_pressure_edit_dia"
	StatePressureEditPulse State = "state_pressure_edit_pulse"
	StatePressureEditTags  State = "state_pressure_edit_tags"
)

// tags
const (
	StateTagCreate State = "state_tag_create"
	StateTagDelete State = "state_tag_delete"

	StateTagName State = "state_tag_name"
)
