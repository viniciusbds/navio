package container

import (
	"fmt"
	"strings"

	"github.com/viniciusbds/navio/constants"
)

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

// ToStr ...
func (c *Container) ToStr() string {
	name := c.Name + strings.Repeat(" ", constants.MaxContainerNameLength-len(c.Name))
	image := c.Image + strings.Repeat(" ", constants.MaxImageNameLength-len(c.Image))
	return fmt.Sprintf("%s\t%s %s\t%s\t\t\t%s", c.ID, name, image, c.Command, c.Status)
}

// IsRunning ...
func (c *Container) IsRunning() bool {
	return c.Status == "Running"
}

// GetStatus ...
func (c *Container) GetStatus() string {
	return c.Status
}

// SetStatus ...
func (c *Container) SetStatus(status string) {
	c.Status = status
}
