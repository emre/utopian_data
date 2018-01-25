package client

import "time"

type Moderator struct {
	Id                           string  `json:"_id"`
	Account                      string  `json:"account"`
	TotalPaidRewards             float64 `json:"total_paid_rewards"`
	ShouldReceiveRewards         float64 `json:"should_receive_rewards"`
	TotalModerated               int     `json:"total_moderated"`
	PercentageOnModeratorRewards float64 `json:"percentage_total_rewards_moderators"`
	Reviewed                     bool    `json:"reviewed"`
	Banned                       bool    `json:"banned"`
	SuperModerator               bool    `json:"supermoderator"`
	TotalPaidRewardsInSteem      float64 `json:"total_paid_rewards_steem"`
}

type ModeratorResponse struct {
	Total      int         `json:"total"`
	Moderators []Moderator `json:"results"`
}

type Sponsor struct {
	ID                           string        `json:"_id"`
	Account                      string        `json:"account"`
	VestingShares                int64         `json:"vesting_shares"`
	PercentageTotalVestingShares float64       `json:"percentage_total_vesting_shares"`
	TotalPaidRewards             float64       `json:"total_paid_rewards,omitempty"`
	ShouldReceiveRewards         float64       `json:"should_receive_rewards,omitempty"`
	IsWitness                    bool          `json:"is_witness"`
	V                            int           `json:"__v"`
	TotalPaidRewardsSteem        float64       `json:"total_paid_rewards_steem,omitempty"`
	Projects                     []interface{} `json:"projects"`
	OptedOut                     bool          `json:"opted_out,omitempty"`
	JSONMetadata struct {
		Profile struct {
			Name         string `json:"name"`
			Location     string `json:"location"`
			ProfileImage string `json:"profile_image"`
		} `json:"profile"`
	} `json:"json_metadata,omitempty"`
}

type SponsorResponse struct {
	Total    int       `json:"total"`
	Sponsors []Sponsor `json:"results"`
}

type Post struct {
	Moderator      string `json:"moderator"`
	Flagged        bool   `json:"flagged"`
	Reviewed       bool   `json:"reviewed"`
	Pending        bool   `json:"pending"`
	MongoId        string `json:"_id"`
	ID             int    `json:"id"`
	Author         string `json:"author"`
	Permlink       string `json:"permlink"`
	Category       string `json:"category"`
	ParentAuthor   string `json:"parent_author"`
	ParentPermlink string `json:"parent_permlink"`
	Title          string `json:"title"`
	Body           string `json:"body"`
	JSONMetadata struct {
		Moderator struct {
			Flagged  bool      `json:"flagged"`
			Pending  bool      `json:"pending"`
			Reviewed bool      `json:"reviewed"`
			Time     time.Time `json:"time"`
			Account  string    `json:"account"`
		} `json:"moderator"`
		Users        []string      `json:"users"`
		Tags         []string      `json:"tags"`
		Type         string        `json:"type"`
		Platform     string        `json:"platform"`
		PullRequests []interface{} `json:"pullRequests"`
		Repository struct {
			Owner struct {
				Login string `json:"login"`
			} `json:"owner"`
			Fork     bool   `json:"fork"`
			HTMLURL  string `json:"html_url"`
			FullName string `json:"full_name"`
			Name     string `json:"name"`
			ID       int    `json:"id"`
		} `json:"repository"`
		Format    string `json:"format"`
		App       string `json:"app"`
		Community string `json:"community"`
	} `json:"json_metadata"`
	LastUpdate              string        `json:"last_update"`
	Created                 string        `json:"created"`
	Active                  string        `json:"active"`
	LastPayout              string        `json:"last_payout"`
	Depth                   int           `json:"depth"`
	Children                int           `json:"children"`
	NetRshares              int64         `json:"net_rshares"`
	AbsRshares              int64         `json:"abs_rshares"`
	VoteRshares             int64         `json:"vote_rshares"`
	ChildrenAbsRshares      int64         `json:"children_abs_rshares"`
	CashoutTime             string        `json:"cashout_time"`
	MaxCashoutTime          string        `json:"max_cashout_time"`
	TotalVoteWeight         int           `json:"total_vote_weight"`
	RewardWeight            int           `json:"reward_weight"`
	TotalPayoutValue        string        `json:"total_payout_value"`
	CuratorPayoutValue      string        `json:"curator_payout_value"`
	AuthorRewards           int           `json:"author_rewards"`
	NetVotes                int           `json:"net_votes"`
	RootComment             int           `json:"root_comment"`
	MaxAcceptedPayout       string        `json:"max_accepted_payout"`
	PercentSteemDollars     int           `json:"percent_steem_dollars"`
	AllowReplies            bool          `json:"allow_replies"`
	AllowVotes              bool          `json:"allow_votes"`
	AllowCurationRewards    bool          `json:"allow_curation_rewards"`
	URL                     string        `json:"url"`
	RootTitle               string        `json:"root_title"`
	PendingPayoutValue      string        `json:"pending_payout_value"`
	TotalPendingPayoutValue string        `json:"total_pending_payout_value"`
	AuthorReputation        int64         `json:"author_reputation"`
	Promoted                string        `json:"promoted"`
	BodyLength              int           `json:"body_length"`
	V                       int           `json:"__v"`
	Reserved                bool          `json:"reserved"`
	Replies                 []interface{} `json:"replies"`
	RebloggedBy             []interface{} `json:"reblogged_by"`
	Beneficiaries []struct {
		Account string `json:"account"`
		Weight  int    `json:"weight"`
	} `json:"beneficiaries"`
	ActiveVotes []struct {
		Time       string      `json:"time"`
		Reputation interface{} `json:"reputation"`
		Percent    int         `json:"percent"`
		Rshares    interface{} `json:"rshares,omitempty"`
		Weight     int         `json:"weight"`
		Voter      string      `json:"voter"`
	} `json:"active_votes"`
}

type PostResponse struct {
	Total int    `json:"total"`
	Posts []Post `json:"results"`
}
