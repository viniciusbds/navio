package container

// Container ...
type Container struct {
	ID      string
	Name    string
	ImageID string
	Status  string
	Root    string
}

// NewContainer ...
func NewContainer(id, name, imageID, status, root string) *Container {
	return &Container{
		ID:      id,
		Name:    name,
		ImageID: imageID,
		Status:  status,
		Root:    root,
	}
}

// ToStr ...
func (c *Container) ToStr() string {
	return c.ID + "\t\t\t" + c.Name + "\t\t\t" + c.ImageID + "\t\t\t" + c.Status
}
