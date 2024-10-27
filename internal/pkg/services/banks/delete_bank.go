package banks

import (
	"context"
)

func (w *BanksHandler) DeleteBank(ctx context.Context, id int64) error {
	if err := w.repo.DeleteBank(ctx, id); err != nil {
		return err
	}

	return nil
}
