package prscontroller

type RunBody struct {
	ID string `json:"id" validate:"required"`
}

type ProfilingParams struct {
	ID string `param:"id" validate:"required"`
}
