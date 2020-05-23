package container

import "fmt"

// Container holds the structure defining a container object.
type Container struct {
	ID      string
	Name    string
	Image   string
	Status  string
	Root    string
	Command string
	Params  []string
}

// NewContainer creates a new container with its basic configuration.
func NewContainer(id, name, image, status, root, command string, params []string) *Container {
	return &Container{
		ID:      id,
		Name:    name,
		Image:   image,
		Status:  status,
		Root:    root,
		Command: command,
		Params:  params,
	}
}

// ToStr refatorar isso, pelo amor de Dios
// Essa foi a minha maior vigarice (https://www.youtube.com/watch?v=PK0c_n5EDhk)
func (c *Container) ToStr() string {
	msg := ""
	if len(c.Name) < 7 {
		msg = fmt.Sprintf("%s\t%s           \t\t\t%s   \t\t\t%s \t\t\t%s ", c.ID, c.Name, c.Command, c.Image, c.Status)
	} else if len(c.Name) <= 12 {
		msg = fmt.Sprintf("%s\t%s           \t\t%s   \t\t\t%s \t\t\t%s", c.ID, c.Name, c.Command, c.Image, c.Status)
	} else {
		msg = fmt.Sprintf("%s\t%s    \t\t%s         \t\t%s\t\t\t%s", c.ID, c.Name, c.Command, c.Image, c.Status)
	}
	return msg
}
