package applemsg

import (
	"context"
	"time"
)

type Handler func(Message)

type Service struct {
	reader *Reader

	interval time.Duration
}

func NewService(
	reader *Reader,
) *Service {

	return &Service{
		reader:   reader,
		interval: time.Second,
	}
}

func (s *Service) Run(
	ctx context.Context,
	cursor CursorStore,
	handler Handler,
) error {

	lastID, err := cursor.Load()

	if err != nil {
		return err
	}

	for {

		select {

		case <-ctx.Done():

			return ctx.Err()

		case <-time.After(s.interval):

			messages, err := s.reader.ReadAfter(
				ctx,
				lastID,
			)

			if err != nil {
				return err
			}

			for _, raw := range messages {

				msg := Parse(raw)

				// 业务处理成功后
				handler(msg)

				if msg.ID > lastID {

					lastID = msg.ID

					err := cursor.Save(
						lastID,
					)

					if err != nil {
						return err
					}
				}
			}
		}
	}
}
