package chain_of_responsibility

import (
	"testing"
)

func TestChainOfResponsiblity(t *testing.T) {
	boardingProcessor := BuildBoardingProcessorChain()
	passenger := &Passenger{
		name:                  "李四",
		hasBoardingPass:       false,
		hasLuggage:            true,
		isPassIdentityCheck:   false,
		isPassSecurityCheck:   false,
		isCompleteForBoarding: false,
	}
	boardingProcessor.ProcessFor(passenger)
}
