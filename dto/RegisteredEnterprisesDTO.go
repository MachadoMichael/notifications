package dto

type RegisteredEnterprisesDTO struct {
	Lenght  int      `json: "lenght"`
	Content []string `json: "content"`
	ReadAt  string   `json: "readAt"`
}
