package adalight

import (
	"log"
	"time"

	"github.com/tarm/serial"
)

type Controller struct {
	portcfg *serial.Config
	strip   *Strip
	fps     int
}

func New(port string, baud int, ledCount int) *Controller {
	c := &Controller{&serial.Config{Name: port, Baud: baud}, newStrip(ledCount), estimateFps(ledCount, baud)}
	return c
}

func (c *Controller) Strip() *Strip {
	return c.strip
}

func (c *Controller) Run(e Effect) error {
	port, err := serial.OpenPort(c.portcfg)
	if err != nil {
		return err
	}
	defer func() {
		if ferr := port.Close(); ferr != nil {
			err = ferr
		}
	}()

	// Initialize the Effect and Frame Time param
	effectName := e.Init(c.strip, c.fps)
	t := time.Now()
	hdr := buildHeader(c.strip.length)
	ft := 1.0 / float64(c.fps)
	fd := time.Duration(ft * float64(time.Second))
	log.Println("Initialized Effect", effectName, "at", t, "with duration", fd.Seconds())

	for f := 0; ; f++ {
		// Get the next frame from the effect
		done := e.Frame(f)
		t = t.Add(fd)

		// Wait for the timer
		time.Sleep(time.Until(t))

		// Write the header and then the pixels
		if _, err := port.Write(hdr); err != nil {
			return err
		}
		if _, err := port.Write(c.strip.pixels); err != nil {
			return err
		}

		if done {
			break
		}
	}
	return nil
}

func buildHeader(cnt int) []byte {
	countHi := (byte)((cnt - 1) >> 8)
	countLo := (byte)((cnt - 1) & 0xFF)
	checksum := countHi ^ countLo ^ 0x55
	return []byte{'A', 'd', 'a', countHi, countLo, checksum}
}

func estimateFps(pixels int, bps int) int {
	bits := (6 + (pixels * 3)) * 8
	max := float64(bps) / float64(bits)
	scaled := int(max * .75)
	return (scaled / 5) * 5
}
