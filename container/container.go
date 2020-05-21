package container

import "fmt"

// Container holds the structure defining a container object.
type Container struct {
	ID      string
	Name    string
	ImageID string
	Status  string
	Root    string
	Command string
}

// NewContainer creates a new container with its basic configuration.
func NewContainer(id, name, imageID, status, root, command string) *Container {
	return &Container{
		ID:      id,
		Name:    name,
		ImageID: imageID,
		Status:  status,
		Root:    root,
		Command: command,
	}
}

// ToStr refatorar isso, pelo amor de Dios
func (c *Container) ToStr() string {
	msg := ""
	if len(c.Name) < 7 {
		msg = fmt.Sprintf("%s\t%s           \t\t\t%s   \t\t\t%s \t\t\t%s", c.ID, c.Name, c.Command, c.ImageID, c.Status)
	} else if len(c.Name) <= 12 {
		msg = fmt.Sprintf("%s\t%s           \t\t%s   \t\t\t%s \t\t\t%s", c.ID, c.Name, c.Command, c.ImageID, c.Status)
	} else {
		msg = fmt.Sprintf("%s\t%s    \t\t%s         \t\t%s\t\t\t%s", c.ID, c.Name, c.Command, c.ImageID, c.Status)
	}
	return msg
}
