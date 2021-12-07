package utils

import (
	"context"

	"github.com/0B1t322/Magic-Circle/ent"
	"github.com/0B1t322/Magic-Circle/ent/predicate"
)

func DeleteVariant(ctx context.Context, client *ent.Client, pred predicate.Variant) error {
	_, err := client.Variant.Delete().Where(pred).Exec(ctx)
	return err
}