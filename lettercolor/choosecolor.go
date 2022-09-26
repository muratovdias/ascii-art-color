package lettercolor

import "errors"

func ChooseColor(color string) (string, error) {
	switch color {
	case "white":
		return "\u001b[38;2;255;255;255m", nil
	case "black":
		return "\u001b[38;2;0;0;0m", nil
	case "red":
		return "\u001b[38;2;255;0;0m", nil
	case "green":
		return "\u001b[38;2;0;255;0m", nil
	case "blue":
		return "\u001b[38;2;0;0;255m", nil
	case "yellow":
		return "\u001b[38;2;255;255;0m", nil
	case "pink":
		return "\u001b[38;2;255;0;255m", nil
	case "grey":
		return "\u001b[38;2;128;128;128m", nil
	case "purple":
		return "\u001b[38;2;160;32;255m", nil
	case "brown":
		return "\u001b[38;2;160;128;96m", nil
	case "orange":
		return "\u001b[38;2;255;160;16m", nil
	case "cyan":
		return "\u001b[38;2;0;183;235m", nil
	}
	return "", errors.New("ERROR: Check the color or choose a different color. For example: <white> \u001b[38;2;0;0;0m <black> \u001b[38;2;255;0;0m <red> \u001b[38;2;0;255;0m <green> \u001b[38;2;0;0;255m <blue> \u001b[38;2;255;255;0m <yellow> \u001b[38;2;255;0;255m <pink> \u001b[38;2;128;128;128m <grey> \u001b[38;2;160;32;255m <purple> \u001b[38;2;160;128;96m <brown> \u001b[38;2;255;160;16m <orange> \u001b[38;2;0;183;235m<cyan>")
}
