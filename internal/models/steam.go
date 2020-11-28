package models

import "github.com/go-pg/pg/v10"

type SteamImpl interface {
	Get(id int) (*Steam, bool)
	Create(s *Steam) (*Steam, error)
}

type Steam struct {
	tableName                struct{} `pg:"steams,alias:c"` //nolint
	ID                       int      `json:"id"`
	CommunityVisibilityState int      `json:"community_visibility_state"`
	ProfileState             int      `json:"profile_state"`
	PersonaName              string   `json:"persona_name"`
	CommentPermission        int      `json:"comment_permission"`
	ProfileURL               string   `json:"profile_url"`
	Avatar                   string   `json:"avatar"`
	AvatarMedium             string   `json:"avatar_medium"`
	AvatarFull               string   `json:"avatar_full"`
	AvatarHash               string   `json:"avatar_hash"`
	LastLogoff               int      `json:"last_logoff"`
	PersonaState             int      `json:"persona_state"`
	RealName                 string   `json:"real_name"`
	PrimaryClanID            string   `json:"primary_clan_id"`
	TimeCreated              int      `json:"time_created"`
	PersonaStateFlags        int      `json:"persona_state_flags"`
	LocCountryCode           string   `json:"loc_country_code"`
}

type SteamData struct {
	Response struct {
		Players []struct {
			SteamID                  string `json:"steamid"`
			CommunityVisibilityState int    `json:"communityvisibilitystate"`
			ProfileState             int    `json:"profilestate"`
			PersonaName              string `json:"personaname"`
			CommentPermission        int    `json:"commentpermission"`
			ProfileURL               string `json:"profileurl"`
			Avatar                   string `json:"avatar"`
			AvatarMedium             string `json:"avatarmedium"`
			AvatarFull               string `json:"avatarfull"`
			AvatarHash               string `json:"avatarhash"`
			LastLogoff               int    `json:"lastlogoff"`
			PersonaState             int    `json:"personastate"`
			RealName                 string `json:"realname"`
			PrimaryClanID            string `json:"primaryclanid"`
			TimeCreated              int    `json:"timecreated"`
			PersonaStateFlags        int    `json:"personastateflags"`
			LocCountryCode           string `json:"loccountrycode"`
		} `json:"players"`
	} `json:"response"`
}

type SteamRepo struct {
	db *pg.DB
}

func NewSteamModel(db *pg.DB) *SteamRepo {
	return &SteamRepo{db}
}

func (r *SteamRepo) Get(id int) (*Steam, bool) {
	steam := &Steam{}
	err := r.db.Model(steam).Where("id = ?", id).Select()
	if err != nil {
		return steam, false
	}

	return steam, true
}

func (r *SteamRepo) Create(s *Steam) (*Steam, error) {
	_, err := r.db.Model(s).Insert()
	if err != nil {
		return nil, err
	}

	return s, nil
}
