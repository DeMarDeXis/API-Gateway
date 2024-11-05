package courses

type InputCourse struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
}

type UpdateCourse struct {
	Name        *string `json:"name"`
	Description *string `json:"desc"`
}

type Courses struct {
	CourseID    int    `json:"course_id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	OwnerID     int    `json:"owner_id"`
}
