package badgerdb

import (
	"context"

	"github.com/FleekHQ/space-daemon/core/store"
	db "github.com/FleekHQ/space-daemon/core/store"
)

type BucketArchive struct {
	BucketKey            string         `bson:"_id"`
	Archives             Archives       `bson:"archives"`
	DefaultArchiveConfig *ArchiveConfig `bson:"default_archive_config"`
}

type Archives struct {
	Current Archive   `bson:"current"`
	History []Archive `bson:"history"`
}

type Archive struct {
	Cid        []byte `bson:"cid"`
	JobID      string `bson:"job_id"`
	JobStatus  int    `bson:"job_status"`
	Aborted    bool   `bson:"aborted"`
	AbortedMsg string `bson:"aborted_msg"`
	FailureMsg string `bson:"failure_msg"`
	CreatedAt  int64  `bson:"created_at"`
}

// ArchiveConfig is the desired state of a Cid in the Filecoin network.
type ArchiveConfig struct {
	// RepFactor (ignored in Filecoin mainnet) indicates the desired amount of active deals
	// with different miners to store the data. While making deals
	// the other attributes of FilConfig are considered for miner selection.
	RepFactor int `bson:"rep_factor"`
	// DealMinDuration indicates the duration to be used when making new deals.
	DealMinDuration int64 `bson:"deal_min_duration"`
	// ExcludedMiners (ignored in Filecoin mainnet) is a set of miner addresses won't be ever be selected
	// when making new deals, even if they comply to other filters.
	ExcludedMiners []string `bson:"excluded_miners"`
	// TrustedMiners (ignored in Filecoin mainnet) is a set of miner addresses which will be forcibly used
	// when making new deals. An empty/nil list disables this feature.
	TrustedMiners []string `bson:"trusted_miners"`
	// CountryCodes (ignored in Filecoin mainnet) indicates that new deals should select miners on specific
	// countries.
	CountryCodes []string `bson:"country_codes"`
	// Renew indicates deal-renewal configuration.
	Renew ArchiveRenew `bson:"renew"`
	// Addr is the wallet address used to store the data in filecoin
	Addr string `bson:"addr"`
	// MaxPrice is the maximum price that will be spent to store the data
	MaxPrice uint64 `bson:"max_price"`
	// FastRetrieval indicates that created deals should enable the
	// fast retrieval feature.
	FastRetrieval bool `bson:"fast_retrieval"`
	// DealStartOffset indicates how many epochs in the future impose a
	// deadline to new deals being active on-chain. This value might influence
	// if miners accept deals, since they should seal fast enough to satisfy
	// this constraint.
	DealStartOffset int64 `bson:"deal_start_offset"`
}

// ArchiveRenew contains renew configuration for a ArchiveConfig.
type ArchiveRenew struct {
	// Enabled indicates that deal-renewal is enabled for this Cid.
	Enabled bool `bson:"enabled"`
	// Threshold indicates how many epochs before expiring should trigger
	// deal renewal. e.g: 100 epoch before expiring.
	Threshold int `bson:"threshold"`
}

type BucketArchives struct {
	st db.Store
}

func NewBucketArchives(_ context.Context, st store.Store) (*BucketArchives, error) {
	k := &BucketArchives{st: st}
	return k, nil
}

func (k *BucketArchives) Create(ctx context.Context, bucketKey string) (*BucketArchive, error) {
	return nil, errNotImplemented
}

func (k *BucketArchives) Replace(ctx context.Context, ba *BucketArchive) error {
	return errNotImplemented
}

func (k *BucketArchives) GetOrCreate(ctx context.Context, bucketKey string) (*BucketArchive, error) {
	return nil, errNotImplemented
}
