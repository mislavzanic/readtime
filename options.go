package readtime

import "errors"

type Options struct {
    wordLength          uint32
	wpm                 uint32
	isTechnical         bool
	technicalDifficulty uint8
	totalWords          uint64
	totalSeconds        uint64
}

func NewOption() *Options {
	return &Options{
		wordLength: 4,
		wpm: 265,
		isTechnical: false,
		technicalDifficulty: 3,
		totalWords: 0,
		totalSeconds: 0,
	}
}

func (o *Options) WordLength(wordLength uint32) *Options {
	o.wordLength = wordLength
	return o
}

func (o *Options) IsTechnical(isTechnical bool) *Options {
	o.isTechnical = isTechnical
	return o
}

func (o *Options) OldReadTime(rt ReadTime) *Options {
	o.totalWords = rt.WordCount
	o.totalSeconds = rt.Seconds
	return o
}

func (o *Options) TechnicalDifficulty(technicalDifficulty uint8) (*Options, error) {
	if technicalDifficulty < 0 || technicalDifficulty > 5 {
		return o, errors.New("technicalDifficulty is a value in [0, 5]!")

	}
	o.technicalDifficulty = technicalDifficulty
	return o, nil
}

func (o *Options) CalculateWPM() uint32 {
	newWPM := o.wpm

	if o.totalWords > 0 && o.totalSeconds > 0 {
		newWPM = ((uint32(o.totalWords) * 60) / (uint32(o.totalSeconds)))
	}

	if o.isTechnical {
		tWPM := o.wpm - (65 + (30 * uint32(o.technicalDifficulty)))
		if tWPM <= 0 {return 50} else {return tWPM}
	}

	if o.totalWords <= 0 || o.totalSeconds <= 0 {
		if !o.isTechnical {
			newWPM = 265
		}
	}

	return newWPM
}
