package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

// errors
var (
	ErrSetPermissions             = errors.Register(ModuleName, 2, "error setting permissions")
	ErrEmptyProposerAccAddress    = errors.Register(ModuleName, 3, "empty proposer key")
	ErrEmptyPermissionsAccAddress = errors.Register(ModuleName, 4, "empty address to set the permissions")
	ErrNotEnoughPermissions       = errors.Register(ModuleName, 5, "not enough permissions")
	ErrCouncilorEmptyAddress      = errors.Register(ModuleName, 6, "empty councilor address")
	ErrRoleDoesNotExist           = errors.Register(ModuleName, 7, "role does not exist")
	ErrWhitelisting               = errors.Register(ModuleName, 8, "error adding to whitelist")
	ErrEmptyPermissions           = errors.Register(ModuleName, 9, "empty permissions")
	ErrRoleExist                  = errors.Register(ModuleName, 12, "role already exist")
	ErrRoleAlreadyAssigned        = errors.Register(ModuleName, 13, "role already assigned")
	ErrRoleNotAssigned            = errors.Register(ModuleName, 14, "role not assigned")
	ErrCouncilorNotFound          = errors.Register(ModuleName, 15, "councilor not found")
	ErrProposalDoesNotExist       = errors.Register(ModuleName, 17, "proposal does not exist")
	ErrActorIsNotActive           = errors.Register(ModuleName, 18, "actor is not active")
	ErrInvalidNetworkProperty     = errors.Register(ModuleName, 19, "invalid network property")
	ErrFeeNotExist                = errors.Register(ModuleName, 20, "fee does not exist")
	ErrPoorNetworkMessagesNotSet  = errors.Register(ModuleName, 21, "poor network messages not set")
	ErrGettingProposals           = errors.Register(ModuleName, 23, "error getting proposals")
	ErrGettingProposalVotes       = errors.Register(ModuleName, 24, "error getting votes for proposal")
	ErrVotingTimeEnded            = errors.Register(ModuleName, 25, "voting time has ended")
	ErrInvalidDate                = errors.Register(ModuleName, 26, "invalid date")
	ErrEmptyVerifierAccAddress    = errors.Register(ModuleName, 27, "empty verifier account address")
	ErrInvalidRecordId            = errors.Register(ModuleName, 28, "invalid record id")
	ErrInvalidVerifyRequestId     = errors.Register(ModuleName, 29, "invalid verify request id")
	ErrInvalidTip                 = errors.Register(ModuleName, 30, "invalid tip")
	ErrInvalidRecordIds           = errors.Register(ModuleName, 31, "invalid record ids")
	ErrEmptyInfos                 = errors.Register(ModuleName, 32, "invalid record ids")
	ErrUnknownProposal            = errors.Register(ModuleName, 33, "unknown proposal")
	ErrInactiveProposal           = errors.Register(ModuleName, 34, "inactive proposal")
	ErrAlreadyActiveProposal      = errors.Register(ModuleName, 35, "proposal already active")
	ErrInvalidProposalContent     = errors.Register(ModuleName, 36, "invalid proposal content")
	ErrInvalidProposalType        = errors.Register(ModuleName, 37, "invalid proposal type")
	ErrInvalidVote                = errors.Register(ModuleName, 38, "invalid vote option")
	ErrInvalidGenesis             = errors.Register(ModuleName, 39, "invalid genesis state")
	ErrNoProposalHandlerExists    = errors.Register(ModuleName, 40, "no handler exists for proposal type")
	ErrInvalidIdentityRecordKey   = errors.Register(ModuleName, 41, "invalid identity record key")
	ErrMonikerDeletionNotAllowed  = errors.Register(ModuleName, 42, "moniker field is not allowed to delete")
	ErrInvalidIdentityRecordId    = errors.Register(ModuleName, 43, "invalid identity record id")
	ErrInvalidApprovalTip         = errors.Register(ModuleName, 44, "invalid approval tip")
	ErrKeyShouldBeUnique          = errors.Register(ModuleName, 45, "key should be unique")
)