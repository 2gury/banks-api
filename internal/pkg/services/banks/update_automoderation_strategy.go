package banks

import (
	"context"
)

func (w *BanksHandler) UpdateAutomoderationStrategy(ctx context.Context, automoderationEnable bool) error {
	if err := w.repo.UpdateAutomoderationStrategy(ctx, automoderationEnable); err != nil {
		return err
	}

	return nil
}
