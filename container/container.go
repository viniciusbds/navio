package container

import (
	"fmt"
	"strings"

	"github.com/viniciusbds/navio/constants"
)

// CGroup holds the structure defining a cgroup object.
type CGroup struct {
	Maxpids   string
	Cpus      string
	Cpushares string
	Memory    string
}

// NewCGroup  creates a new cgroup with its basic configuration.
func NewCGroup(maxpids, cpus, cpushares, memlimit string) *CGroup {
	return &CGroup{
		Maxpids:   maxpids,
		Cpus:      cpus,
		Cpushares: cpushares,
		Memory:    memlimit,
	}
}

// Container holds the structure defining a container object.
type Container struct {
	ID      string
	Name    string
	Image   string
	Status  string
	RootFS  string
	Command string
	Params  []string
	CGroup  *CGroup
}

// NewContainer creates a new container with its basic configuration.
func NewContainer(id, name, image, status, rootfs, command string, params []string, cgroup *CGroup) *Container {
	var pids, cpus, cpushares, memory string

	if cgroup != nil && cgroup.Maxpids != "" {
		pids = cgroup.Maxpids
	} else {
		pids = constants.DefaultMaxProcessCreation
	}
	if cgroup != nil && cgroup.Cpus != "" {
		cpus = cgroup.Cpus
	} else {
		cpus = constants.DefaultCPUS
	}
	if cgroup != nil && cgroup.Cpushares != "" {
		cpushares = cgroup.Cpushares
	} else {
		cpushares = constants.DefaultCPUshares
	}
	if cgroup != nil && cgroup.Memory != "" {
		memory = cgroup.Memory
	} else {
		memory = constants.DefaultMemlimit
	}

	container := &Container{
		ID:      id,
		Name:    name,
		Image:   image,
		Status:  status,
		RootFS:  rootfs,
		Command: command,
		Params:  params,
		CGroup: &CGroup{
			Maxpids:   pids,
			Cpus:      cpus,
			Cpushares: cpushares,
			Memory:    memory,
		},
	}

	return container
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

// GetMaxpids ...
func (c *Container) GetMaxpids() string {
	return c.CGroup.Maxpids
}

// SetMaxpids ...
func (c *Container) SetMaxpids(pids string) {
	c.CGroup.Maxpids = pids
}

// GetCPUS ...
func (c *Container) GetCPUS() string {
	return c.CGroup.Cpus
}

// SetCPUS ...
func (c *Container) SetCPUS(cpus string) {
	c.CGroup.Cpus = cpus
}

// GetCPUshares ...
func (c *Container) GetCPUshares() string {
	return c.CGroup.Cpushares
}

// SetCPUshares ...
func (c *Container) SetCPUshares(cpushares string) {
	c.CGroup.Cpushares = cpushares
}

// GetMemory ...
func (c *Container) GetMemory() string {
	return c.CGroup.Memory
}

// SetMemory ...
func (c *Container) SetMemory(memory string) {
	c.CGroup.Memory = memory
}
