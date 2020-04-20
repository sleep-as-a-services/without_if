package captcha

import (
	"errors"
	"fmt"
	"sleep-as-a-serive/strcon"
)

var (
	ErrUnsupportMathOperator   = errors.New("unsupport math operator")
	ErrUnsupportFormatOperator = errors.New("unsupport format operator")

	mathOperator = map[int]string{
		0: "+",
		1: "-",
		2: "*",
	}
)

const (
	Format string = "%v %s %v"

	LeftFormatOperator  = 0
	RightFormatOperator = 1
)

type Captcha struct {
	Right        int
	Left         int
	MathOperator int
	Format       int
}

func NewCaptcha(format, left, mathOperator, right int) *Captcha {
	return &Captcha{
		right,
		left,
		mathOperator,
		format,
	}
}

func (c *Captcha) Validate() error {
	if _, ok := mathOperator[c.MathOperator]; !ok {
		return ErrUnsupportMathOperator
	}
	if c.Format != LeftFormatOperator && c.Format != RightFormatOperator {
		return ErrUnsupportFormatOperator
	}
	return nil
}

func (c *Captcha) GenerateCaptcha() string {
	if c.Format == LeftFormatOperator {
		return c.doLeft()
	}
	return c.doRight()
}

func (c *Captcha) mathOperatorSymbol() string {
	return mathOperator[c.MathOperator]
}

func (c *Captcha) doLeft() string {
	return fmt.Sprintf(Format, c.Left, c.mathOperatorSymbol(), strcon.Ntot(c.Right))
}

func (c *Captcha) doRight() string {
	return fmt.Sprintf(Format, strcon.Ntot(c.Left), c.mathOperatorSymbol(), c.Right)
}
