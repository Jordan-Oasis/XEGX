	"flag"
	"log"
	"os"
	"strconv"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/unisat-wallet/libbrc20-indexer/conf"
@@ -37,6 +38,12 @@ func init() {
	if id := os.Getenv("MODULE_SWAP_SOURCE_INSCRIPTION_ID"); id != "" {
		conf.MODULE_SWAP_SOURCE_INSCRIPTION_ID = id
	}

	if heightStr := os.Getenv("BRC20_ENABLE_SELF_MINT_HEIGHT"); heightStr != "" {
		if h, err := strconv.Atoi(heightStr); err != nil {
			conf.ENABLE_SELF_MINT_HEIGHT = uint32(h)
		}
	}
}

func main() {
  9 changes: 5 additions & 4 deletions9  
conf/conf.go
Original file line number	Diff line number	Diff line change
@@ -5,8 +5,9 @@ import (
)

var (
	DEBUG                             = false
	MODULE_SWAP_SOURCE_INSCRIPTION_ID = "d2a30f6131324e06b1366876c8c089d7ad2a9c2b0ea971c5b0dc6198615bda2ei0"
	GlobalNetParams                   = &chaincfg.MainNetParams
	TICKS_ENABLED                     = ""
	DEBUG                                    = false
	MODULE_SWAP_SOURCE_INSCRIPTION_ID        = "d2a30f6131324e06b1366876c8c089d7ad2a9c2b0ea971c5b0dc6198615bda2ei0"
	GlobalNetParams                          = &chaincfg.MainNetParams
	TICKS_ENABLED                            = ""
	ENABLE_SELF_MINT_HEIGHT           uint32 = 837090
)
  2 changes: 1 addition & 1 deletion2  
indexer/brc20_deploy.go
Original file line number	Diff line number	Diff line change
@@ -31,7 +31,7 @@ func (g *BRC20ModuleIndexer) ProcessDeploy(data *model.InscriptionBRC20Data) err
			return nil
			// return errors.New("deploy, tick length 5, but not self_mint")
		}
		if data.Height < g.EnableSelfMintHeight {
		if data.Height < conf.ENABLE_SELF_MINT_HEIGHT {
			return nil
			// return errors.New("deploy, tick length 5, but not enabled")
		}
  7 changes: 3 additions & 4 deletions7  
indexer/init.go
Original file line number	Diff line number	Diff line change
@@ -10,9 +10,8 @@ import (
)

type BRC20ModuleIndexer struct {
	BestHeight           uint32
	EnableSelfMintHeight uint32
	EnableHistory        bool
	BestHeight    uint32
	EnableHistory bool

	HistoryCount uint32
	HistoryData  [][]byte
@@ -108,6 +107,7 @@ func (g *BRC20ModuleIndexer) UpdateHistoryHeightAndGetHistoryIndex(historyObj *m
func (g *BRC20ModuleIndexer) initBRC20() {
	g.EnableHistory = true
	g.BestHeight = 0

	g.HistoryCount = 0
	g.HistoryData = make([][]byte, 0)

@@ -237,7 +237,6 @@ func (g *BRC20ModuleIndexer) GenerateApproveEventsByApprove(owner string, balanc
func (copyDup *BRC20ModuleIndexer) deepCopyBRC20Data(base *BRC20ModuleIndexer) {
	// history
	copyDup.BestHeight = base.BestHeight
	copyDup.EnableSelfMintHeight = base.EnableSelfMintHeight
	copyDup.EnableHistory = base.EnableHistory
	copyDup.HistoryCount = base.HistoryCount
