package code

import "fmt"

// The Product (Complex Object)
type Computer struct {
	CPU       string
	RAM       string
	HardDrive string
	Graphics  string
}

func (c *Computer) String() string {
	return fmt.Sprintf("Computer [CPU: %s, RAM: %s, HardDrive: %s, Graphics: %s]", c.CPU, c.RAM, c.HardDrive, c.Graphics)
}

// The Builder Interface
type ComputerBuilder interface {
	SetCPU(cpu string) ComputerBuilder
	SetRAM(ram string) ComputerBuilder
	SetHardDrive(hd string) ComputerBuilder
	SetGraphics(graphics string) ComputerBuilder
	Build() *Computer
}

// Concrete Builder (Builds the Computer)
type ComputerBuilderImpl struct {
	computer *Computer
}

func NewComputerBuilder() ComputerBuilder {
	return &ComputerBuilderImpl{computer: &Computer{}}
}

func (b *ComputerBuilderImpl) SetCPU(cpu string) ComputerBuilder {
	b.computer.CPU = cpu
	return b
}

func (b *ComputerBuilderImpl) SetRAM(ram string) ComputerBuilder {
	b.computer.RAM = ram
	return b
}

func (b *ComputerBuilderImpl) SetHardDrive(hd string) ComputerBuilder {
	b.computer.HardDrive = hd
	return b
}

func (b *ComputerBuilderImpl) SetGraphics(graphics string) ComputerBuilder {
	b.computer.Graphics = graphics
	return b
}

func (b *ComputerBuilderImpl) Build() *Computer {
	return b.computer
}

func main() {
	// Using the builder to construct a Computer object
	builder := NewComputerBuilder()
	computer := builder.SetCPU("Intel i9").
		SetRAM("32GB").
		SetHardDrive("1TB SSD").
		SetGraphics("NVIDIA GTX 3080").
		Build()

	// Output the final constructed object
	fmt.Println(computer) // Output: Computer [CPU: Intel i9, RAM: 32GB, HardDrive: 1TB SSD, Graphics: NVIDIA GTX 3080]
}
