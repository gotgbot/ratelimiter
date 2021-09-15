package ratelimiter

import (
	"fmt"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func (l Limiter) Name() string {
	return fmt.Sprintf("limiter_%p", l.HandleUpdate)
}

func (l *Limiter) CheckUpdate(b *gotgbot.Bot, u *gotgbot.Update) bool {
	return l.isException(u)
}

func (l *Limiter) HandleUpdate(b *gotgbot.Bot, ctx *ext.Context) error {
	var status *UserStatus
	var id int64
	if l.ConsiderUser && ctx.EffectiveUser != nil {
		id = ctx.EffectiveUser.Id
	} else if ctx.EffectiveChat != nil {
		id = ctx.EffectiveChat.Id
	} else {
		return ext.ContinueGroups
	}

	l.mutex.Lock()
	status = l.userMap[id]
	if status == nil {
		status = new(UserStatus)
		status.Last = time.Now()
		status.count++
		l.userMap[id] = status
		l.mutex.Unlock()
		return ext.ContinueGroups
	}

	if status.limited {
		l.mutex.Unlock()
		return ext.EndGroups
	}

	if time.Since(status.Last) > l.timeout {
		status.count = 0
	}

	status.count++

	if status.count > l.maxCount {
		status.limited = true
		l.mutex.Unlock()
		if l.trigger != nil {
			go l.trigger(b, ctx)
		}

		return ext.EndGroups
	}

	l.mutex.Unlock()

	return ext.ContinueGroups
}
