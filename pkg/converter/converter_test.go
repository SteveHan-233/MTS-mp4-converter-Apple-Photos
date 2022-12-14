package converter

import (
	"testing"
)

var expectedCommand string = "mkdir -p \"dest\" && ffmpeg -i \"original/0.MTS\" -c:v libx265 -preset fast -crf 28 -tag:v hvc1 -c:a eac3 -b:a 224k \"dest/0.mp4\" && exiftool -tagsFromFile \"original/0.MTS\" -time:all -overwrite_original \"dest/0.mp4\""

type mockCmd struct{}

func (*mockCmd) CombinedOutput() ([]byte, error) {
	return nil, nil
}

func TestRightCommand(t *testing.T) {
	err := ConvertVideo("original/0.MTS", "dest/0.mp4",
		func(command string, args ...string) Command {
			if command != "zsh" {
				t.Errorf("Expected command zsh, received %s", command)
			}
			if args[0] != "-c" {
				t.Errorf("Expected flag -c, received %s", args[0])
			}
			if args[1] != expectedCommand {
				t.Errorf(
					"Expected command %s but received %s",
					expectedCommand,
					args[1],
				)
			}
			return &mockCmd{}
		})
	if err != nil {
		t.Error("Method resulted in error")
	}
}
