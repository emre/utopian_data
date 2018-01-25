package client

import "time"

type Moderator struct {
	Id                           string  `json:"_id" bson:"_id"`
	Account                      string  `json:"account" bson:"account"`
	TotalPaidRewards             float64 `json:"total_paid_rewards" bson:"total_paid_rewards"`
	ShouldReceiveRewards         float64 `json:"should_receive_rewards" bson:"should_receive_rewards"`
	TotalModerated               int     `json:"total_moderated" bson:"total_moderated"`
	PercentageOnModeratorRewards float64 `json:"percentage_total_rewards_moderators" bson:"percentage_total_rewards_moderators"`
	Reviewed                     bool    `json:"reviewed" bson:"reviewed"`
	Banned                       bool    `json:"banned" bson:"banned"`
	SuperModerator               bool    `json:"supermoderator" bson:"supermoderator"`
	TotalPaidRewardsInSteem      float64 `json:"total_paid_rewards_steem" bson:"total_paid_rewards_steem"`
}

type ModeratorResponse struct {
	Total      int         `json:"total"`
	Moderators []Moderator `json:"results"`
}

type Sponsor struct {
	ID                           string        `json:"_id" bson:"_id"`
	Account                      string        `json:"account" bson:"account"`
	VestingShares                int64         `json:"vesting_shares" bson:"vesting_shares"`
	PercentageTotalVestingShares float64       `json:"percentage_total_vesting_shares" bson:"percentage_total_vesting_shares"`
	TotalPaidRewards             float64       `json:"total_paid_rewards,omitempty" bson:"total_paid_rewards,omitempty"`
	ShouldReceiveRewards         float64       `json:"should_receive_rewards,omitempty" bson:"should_receive_rewards,omitempty"`
	IsWitness                    bool          `json:"is_witness" bson:"is_witness"`
	V                            int           `json:"__v" bson:"__v"`
	TotalPaidRewardsSteem        float64       `json:"total_paid_rewards_steem,omitempty" bson:"total_paid_rewards_steem,omitempty"`
	Projects                     []interface{} `json:"projects" bson:"projects"`
	OptedOut                     bool          `json:"opted_out,omitempty" bson:"opted_out,omitempty"`
	JSONMetadata struct {
		Profile struct {
			Name         string `json:"name" bson:"name"`
			Location     string `json:"location" bson:"location"`
			ProfileImage string `json:"profile_image" bson:"profile_image"`
		} `json:"profile" bson:"profile"`
	} `json:"json_metadata,omitempty" bson:"json_metadata,omitempty"`
}

type SponsorResponse struct {
	Total    int       `json:"total"`
	Sponsors []Sponsor `json:"results"`
}

type Post struct {
	Moderator      string `json:"moderator" bson:"moderator"`
	Flagged        bool   `json:"flagged" bson:"flagged"`
	Reviewed       bool   `json:"reviewed" bson:"reviewed"`
	Pending        bool   `json:"pending" bson:"pending"`
	MongoId        string `json:"_id" bson:"_id"`
	ID             int    `json:"id" bson:"id"`
	Author         string `json:"author" bson:"author"`
	Permlink       string `json:"permlink" bson:"permlink"`
	Category       string `json:"category" bson:"category"`
	ParentAuthor   string `json:"parent_author" bson:"parent_author"`
	ParentPermlink string `json:"parent_permlink" bson:"parent_permlink"`
	Title          string `json:"title" bson:"title"`
	Body           string `json:"body" bson:"body"`
	JSONMetadata struct {
		Moderator struct {
			Flagged  bool      `json:"flagged" bson:"flagged"`
			Pending  bool      `json:"pending" bson:"pending"`
			Reviewed bool      `json:"reviewed" bson:"reviewed"`
			Time     time.Time `json:"time" bson:"time"`
			Account  string    `json:"account" bson:"account"`
		} `json:"moderator"`
		Users        []string      `json:"users" bson:"users"`
		Tags         []string      `json:"tags" bson:"tags"`
		Type         string        `json:"type" bson:"type"`
		Platform     string        `json:"platform" bson:"platform"`
		PullRequests []interface{} `json:"pullRequests" bson:"pullRequests"`
		Repository struct {
			Owner struct {
				Login string `json:"login" bson:"login"`
			} `json:"owner" bson:"owner"`
			Fork     bool   `json:"fork" bson:"fork"`
			HTMLURL  string `json:"html_url" bson:"html_url"`
			FullName string `json:"full_name" bson:"full_name"`
			Name     string `json:"name" bson:"name"`
			ID       int    `json:"id" bson:"id"`
		} `json:"repository" bson:"repository"`
		Format    string `json:"format" bson:"format"`
		App       string `json:"app" bson:"app"`
		Community string `json:"community" bson:"community"`
	} `json:"json_metadata" bson:"json_metadata"`
	LastUpdate              string        `json:"last_update" bson:"last_update"`
	Created                 string        `json:"created" bson:"created"`
	Active                  string        `json:"active" bson:"active"`
	LastPayout              string        `json:"last_payout" bson:"last_payout"`
	Depth                   int           `json:"depth" bson:"depth"`
	Children                int           `json:"children" bson:"children"`
	NetRshares              int64         `json:"net_rshares" bson:"net_rshares"`
	AbsRshares              int64         `json:"abs_rshares" bson:"abs_rshares"`
	VoteRshares             int64         `json:"vote_rshares" bson:"vote_rshares"`
	ChildrenAbsRshares      int64         `json:"children_abs_rshares" bson:"children_abs_rshares"`
	CashoutTime             string        `json:"cashout_time" bson:"cashout_time"`
	MaxCashoutTime          string        `json:"max_cashout_time" bson:"max_cashout_time"`
	TotalVoteWeight         int           `json:"total_vote_weight" bson:"total_vote_weight"`
	RewardWeight            int           `json:"reward_weight" bson:"reward_weight"`
	TotalPayoutValue        string        `json:"total_payout_value" bson:"total_payout_value"`
	CuratorPayoutValue      string        `json:"curator_payout_value" bson:"curator_payout_value"`
	AuthorRewards           int           `json:"author_rewards" bson:"author_rewards"`
	NetVotes                int           `json:"net_votes" bson:"net_votes"`
	RootComment             int           `json:"root_comment" bson:"root_comment"`
	MaxAcceptedPayout       string        `json:"max_accepted_payout" bson:"max_accepted_payout"`
	PercentSteemDollars     int           `json:"percent_steem_dollars" bson:"percent_steem_dollars"`
	AllowReplies            bool          `json:"allow_replies" bson:"allow_replies"`
	AllowVotes              bool          `json:"allow_votes" bson:"allow_votes"`
	AllowCurationRewards    bool          `json:"allow_curation_rewards" bson:"allow_curation_rewards"`
	URL                     string        `json:"url" bson:"url"`
	RootTitle               string        `json:"root_title" bson:"root_title"`
	PendingPayoutValue      string        `json:"pending_payout_value" bson:"pending_payout_value"`
	TotalPendingPayoutValue string        `json:"total_pending_payout_value" bson:"total_pending_payout_value"`
	AuthorReputation        int64         `json:"author_reputation" bson:"author_reputation"`
	Promoted                string        `json:"promoted" bson:"promoted"`
	BodyLength              int           `json:"body_length" bson:"body_length"`
	V                       int           `json:"__v" bson:"__v"`
	Reserved                bool          `json:"reserved" bson:"reserved"`
	Replies                 []interface{} `json:"replies" bson:"replies"`
	RebloggedBy             []interface{} `json:"reblogged_by" bson:"reblogged_by"`
	Beneficiaries []struct {
		Account string `json:"account" bson:"account"`
		Weight  int    `json:"weight" bson:"weight"`
	} `json:"beneficiaries" bson:"beneficiaries"`
	ActiveVotes []struct {
		Time       string      `json:"time" bson:"time"`
		Reputation interface{} `json:"reputation" bson:"reputation"`
		Percent    int         `json:"percent" bson:"percent"`
		Rshares    interface{} `json:"rshares,omitempty" bson:"rshares,omitempty"`
		Weight     int         `json:"weight" bson:"weight"`
		Voter      string      `json:"voter" json:"voter"`
	} `json:"active_votes" bson:"active_votes"`
}

type PostResponse struct {
	Total int    `json:"total"`
	Posts []Post `json:"results"`
}
