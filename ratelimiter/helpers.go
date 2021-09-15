package ratelimiter

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

type NewLimiterOpts struct {
	maxTimeout time.Duration
	timeout time.Duration
	maxCount int
}

func New(opts *NewLimiterOpts) *Limiter {
	l := new(Limiter)

	if opts == nil {
		l.timeout = DEFAULT_TIME
		l.maxCount = DEFAULT_COUNT
		l.maxTimeout = DEFAULT_MAX_TIMEOUT
	} else {
		if opts.timeout == 0 {
		l.timeout = DEFAULT_TIME
		} else {
			l.timeout = opts.timeout
		}

		if opts.maxCount == 0 {
			l.maxCount = DEFAULT_COUNT
		} else {
			l.maxCount = opts.maxCount
		}

		if opts.maxTimeout == 0 {
			l.maxTimeout = DEFAULT_MAX_TIMEOUT
		} else {
			l.maxTimeout = opts.maxTimeout
		}
	}
	
	return l
}



func getFromId(u *gotgbot.Update) int64 {
	if u.ChannelPost != nil {
		return u.ChannelPost.From.Id
	}

	if u.EditedMessage != nil {
		return u.EditedMessage.From.Id
	}

	if u.Message != nil {
		return u.Message.From.Id
	}

	if u.CallbackQuery != nil {
		return u.ChannelPost.From.Id
	}

	return 0
}

func getChatId(u *gotgbot.Update) int64 {
	if u.ChannelPost != nil {
		return u.ChannelPost.Chat.Id
	}

	if u.EditedMessage != nil {
		return u.EditedMessage.Chat.Id
	}

	if u.Message != nil {
		return u.Message.Chat.Id
	}

	if u.CallbackQuery != nil {
		return u.CallbackQuery.Message.Chat.Id
	}

	return 0
}
