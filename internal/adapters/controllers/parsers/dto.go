package prscontroller

type RunBody struct {
	ID string `json:"id" validate:"required"`
}

type ProfilingParams struct {
	ID string `param:"id" validate:"required"`
}

type OptionsBody struct {
	ID   string `json:"id" validate:"required"`
	Wrks int    `json:"wrks" validate:"required"`
}
