package main

import (
	"bufio"
	"countdown/internal"
	"countdown/sevseg"
	"errors"
	"math/rand"
	"path"
	"strconv"
	"strings"

	"fmt"
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const (
	DirectoryWithSounds = "sounds"
	SoundExtension      = ".mp3"
	SecondsInHour       = 3600
	SecondsInMinute     = 60
	CountDelayInMs      = 1000
)

var (
	in    = bufio.NewReader(os.Stdin)
	out   = os.Stdout
	input = internal.NewInput(out, in)
)

func main() {
	hours := input.EnterHours()
	minutes := input.EnterMinutes()
	seconds := input.EnterSeconds()
	input.PressEnter()

	summarySeconds := hours*SecondsInHour + minutes*SecondsInMinute + seconds

	err := termbox.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go bgthread(summarySeconds)

	timeEnd := false

	for {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			if ev.Key == termbox.KeyCtrlC {
				break
			}

			termbox.Flush()
		}
		if ev.Type == termbox.EventInterrupt {
			timeEnd = true
			break
		}
	}

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
	termbox.Close()

	if timeEnd {
		done()
	} else {
		fmt.Println("The countdown stopped by user.")
	}
}

func bgthread(summarySeconds int) {
	seconds := summarySeconds
	separator := ":"
	ticker := time.NewTicker(CountDelayInMs * time.Millisecond)
	defer ticker.Stop()

	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

		str := getTimeAsString(seconds, separator)
		segments, err := sevseg.GetSegmentsInRows(str, 2)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(1)
		}

		// Remove '\n'
		rows := make([]string, 3)
		for i := range segments {
			rows[i] = strings.TrimSuffix(segments[i], "\n")
		}

		tbprint(1, 1, termbox.ColorDefault, termbox.ColorDefault, rows[0])
		tbprint(1, 2, termbox.ColorDefault, termbox.ColorDefault, rows[1])
		tbprint(1, 3, termbox.ColorDefault, termbox.ColorDefault, rows[2])
		tbprint(1, 5, termbox.ColorDefault, termbox.ColorDefault, "Press 'Ctrl-C' to quit.")
		termbox.Flush()
		<-ticker.C // wait for ticker

		seconds--

		// Blink
		if separator == ":" {
			separator = " "
		} else {
			separator = ":"
		}

		// Quit on time's out git
		if seconds < 0 {
			termbox.Interrupt()
		}
	}
}

func getTimeAsString(summarySeconds int, separator string) []string {
	hours := summarySeconds / SecondsInHour
	secondsLeft := summarySeconds - hours*SecondsInHour
	minutes := secondsLeft / SecondsInMinute
	seconds := secondsLeft - minutes*SecondsInMinute
	return []string{
		strconv.Itoa(hours),
		separator,
		strconv.Itoa(minutes),
		separator,
		strconv.Itoa(seconds),
	}
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

func done() {
	fmt.Println("Time's up!")

	fileToOpen, err := selectSoundFile()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	file, err := os.Open(fileToOpen)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	streamer, format, err := mp3.Decode(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}

func selectSoundFile() (string, error) {
	files, err := os.ReadDir(DirectoryWithSounds)
	if err != nil {
		return "", err
	}

	sounds := []string{}
	for _, file := range files {
		filename := file.Name()
		fileExtension := strings.ToLower(path.Ext(filename))
		if fileExtension == SoundExtension {
			sounds = append(sounds, filename)
		}
	}

	cnt := len(sounds)
	if cnt < 1 {
		return "", errors.New("no sound files found")
	}

	if cnt == 1 {
		filePath := path.Join(DirectoryWithSounds, sounds[0])
		return filePath, nil
	}

	filePath := path.Join(DirectoryWithSounds, sounds[rand.Intn(cnt)])
	return filePath, nil
}
