package provider_test

import (
	"testing"

	"github.com/117503445/markdown-translate/internal/provider"
	"github.com/rs/zerolog/log"

	"github.com/stretchr/testify/assert"
)

func TestOpenAIProvider_Translate(t *testing.T) {

	assert := assert.New(t)
	p := provider.NewOpenAIProvider()

	sources := []string{
		"where $K$ is the workload for receiving a proposal from the leader. Finally, we can derive the maximum throughput as",
		"Network asynchrony. During network asynchrony, a proposal is likely to arrive before some of referenced transactions (i.e. missing transactions), which negatively impacts performance. The point of this experiment is to show that Stratus-based",
		"To address these challenges, we use stable time (ST) to es-timate a replica's load status. The stable time of a microblock is measured from when the sender broadcasts the microblock until the time that the microblock becomes stable (receiving $f + 1$ acks). To estimate ST of a replica,the replica calculates the ST of each microblock if it is the sender and takes the $n$ -th (e.g., $n = 95$ ) percentile of the ST values in a window of the latest stable microblocks. Figure 4 shows the estimation process. The estimated ST of a replica is updated when a new microblock becomes stable. The window size is configurable and we use 100 as the default size.",
	}

	for _, source := range sources {
		text, err := p.Translate(source)
		assert.Nil(err)

		log.Debug().Str("source", source).Str("text", text).Msg("translated")
	}

}
