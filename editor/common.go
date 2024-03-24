package editor

import "github.com/gvcgo/asciinema-edit/cast"

func FindIndex(eventStream []*cast.Event, from, to float64) (start, end int) {
	if len(eventStream) == 0 {
		return -1, -1
	}

	if from >= eventStream[len(eventStream)-1].Time {
		return -1, -1
	}
	if to <= eventStream[0].Time {
		return -1, -1
	}
	for i, ev := range eventStream {
		if i == 0 && ev.Time >= from {
			start = 0
			continue
		}
		if i == len(eventStream)-1 && ev.Time <= to {
			end = i
			continue
		}

		if ev.Time < from && eventStream[i+1].Time > from {
			start = i + 1
		}

		if ev.Time < to && eventStream[i+1].Time > to {
			end = i + 1
		}
	}
	return
}
