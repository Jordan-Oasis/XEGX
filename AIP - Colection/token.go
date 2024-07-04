@@ -4,6 +4,7 @@ import (
	"context"
	"time"

	"github.com/adshao/go-brc721/sig"
	"github.com/go-kratos/kratos/v2/log"
)

@@ -13,17 +14,18 @@ const (

// Token is a Token model.
type Token struct {
	ID             int       `json:"id"`
	P              string    `json:"p"`
	Tick           string    `json:"tick"`
	TokenID        uint64    `json:"token_id"`
	TxHash         string    `json:"tx_hash"`
	BlockHeight    uint64    `json:"block_height"`
	BlockTime      time.Time `json:"block_time"`
	Address        string    `json:"address"`
	InscriptionID  int64     `json:"inscription_id"`
	InscriptionUID string    `json:"inscription_uid"`
	CollectionID   int       `json:"collection_id"`
	ID             int         `json:"id"`
	P              string      `json:"p"`
	Tick           string      `json:"tick"`
	TokenID        uint64      `json:"token_id"`
	TxHash         string      `json:"tx_hash"`
	BlockHeight    uint64      `json:"block_height"`
	BlockTime      time.Time   `json:"block_time"`
	Address        string      `json:"address"`
	InscriptionID  int64       `json:"inscription_id"`
	InscriptionUID string      `json:"inscription_uid"`
	CollectionID   int         `json:"collection_id"`
	Sig            sig.MintSig `json:"sig,omitempty"`
}

type TokenListOption struct {
@@ -40,6 +42,7 @@ type TokenRepo interface {
	Update(context.Context, *Token) (*Token, error)
	FindByTickTokenID(context.Context, string, string, uint64) (*Token, error)
	FindByInscriptionID(context.Context, int64) ([]*Token, error)
	FindByTickSigUID(context.Context, string, string, string) (*Token, error)
	List(context.Context, ...TokenListOption) ([]*Token, error)
	Delete(context.Context, int) error
	Count(context.Context, ...TokenListOption) (int, error)
@@ -80,6 +83,12 @@ func (uc *TokenUsecase) FindByInscriptionID(ctx context.Context, inscriptionID i
	return uc.repo.FindByInscriptionID(ctx, inscriptionID)
}

// FindByTickSigUID finds the Token by Tick and SigUID.
func (uc *TokenUsecase) FindByTickSigUID(ctx context.Context, p, tick, sigUID string) (*Token, error) {
	uc.log.WithContext(ctx).Debugf("FindByTickSigUID for %s %s %s", p, tick, sigUID)
	return uc.repo.FindByTickSigUID(ctx, p, tick, sigUID)
}

// ListTokens lists Tokens.
func (uc *TokenUsecase) ListTokens(ctx context.Context, opt *TokenListOption) ([]*Token, error) {
	uc.log.WithContext(ctx).Debugf("ListTokens for %v", opt)
