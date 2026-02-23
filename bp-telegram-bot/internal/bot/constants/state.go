package constants

type State string

const (
	// menu states
	StateIdle     State = ""
	StatePressure State = "state_pressure"
	StateTag      State = "state_tag"

	// pressure
	StatePressureCreate State = "state_create"
	StatePressureEdit   State = "state_pressure_edit"
	StatePressureDelete State = "state_pressure_delete"

	StatePressureSys   State = "state_pressure_sys"
	StatePressureDia   State = "state_pressure_dia"
	StatePressurePulse State = "state_pressure_pulse"
	StatePressureTags  State = "state_pressure_tags"

	StatePressureEditSys   State = "state_pressure_edit_sys"
	StatePressureEditDia   State = "state_pressure_edit_dia"
	StatePressureEditPulse State = "state_pressure_edit_pulse"
	StatePressureEditTags  State = "state_pressure_edit_tags"

	// tags
	StateTagCreate State = "state_tag_create"
	StateTagEdit   State = "state_tag_edit"
	StateTagDelete State = "state_tag_delete"

	StateTagName     State = "state_tag_name"
	StateTagEditName State = "state_tag_edit_name"
)
