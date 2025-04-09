package bernabeu

import (
	"encoding/json"
	"event-notifier/utils"
	"fmt"
	"io"
	"net/http"
	"time"
)

type FootballManager struct {
	MatchesToday []Match
}

func NewFootballManager() *FootballManager {
	return &FootballManager{MatchesToday: []Match{}}
}

type FootballEventsResponse struct {
	Data Data `json:"data"`
}

type Data struct {
	MatchList MatchList `json:"matchList"`
}

type MatchList struct {
	Items []Match `json:"items"`
}

type Match struct {
	HideMatchFeed       bool   `json:"hideMatchFeed"`
	HideMatchTickets    bool   `json:"hideMatchTickets"`
	HideMatchCalendar   bool   `json:"hideMatchCalendar"`
	ReducedDisplay      bool   `json:"reducedDisplay"`
	ID                  string `json:"id"`
	OptaID              string `json:"optaId"`
	OptaLegacyID        string `json:"optaLegacyId"`
	CartoExternalID     any    `json:"cartoExternalId"`
	MonterosaExternalID string `json:"monterosaExternalId"`
	Slug                string `json:"slug"`
	Path                string `json:"_path"`
	Competition         struct {
		OptaID        string `json:"optaId"`
		OptaLegacyID  string `json:"optaLegacyId"`
		CompetitionID string `json:"competitionId"`
		Slug          string `json:"slug"`
		Logo          struct {
			PublishURL string `json:"_publishUrl"`
			DmS7URL    string `json:"_dmS7Url"`
		} `json:"logo"`
		Name string `json:"name"`
	} `json:"competition"`
	Subtitle    any `json:"subtitle"`
	Description struct {
		Plaintext string `json:"plaintext"`
	} `json:"description"`
	Week        string    `json:"week"`
	WeekAsText  any       `json:"weekAsText"`
	DateTime    time.Time `json:"dateTime"`
	DateTxt     any       `json:"dateTxt"`
	IsScheduled bool      `json:"isScheduled"`
	Status      string    `json:"status"`
	Venue       struct {
		Path      string `json:"_path"`
		Variation string `json:"_variation"`
		OptaID    string `json:"optaId"`
		Name      string `json:"name"`
		City      any    `json:"city"`
		Country   any    `json:"country"`
		Address   struct {
			Plaintext any `json:"plaintext"`
		} `json:"address"`
	} `json:"venue"`
	Tv struct {
		HTML string `json:"html"`
	} `json:"tv"`
	RelatedNews struct {
		Slug string   `json:"slug"`
		Tag  []string `json:"tag"`
	} `json:"relatedNews"`
	RelatedPublication any `json:"relatedPublication"`
	HomeTeam           struct {
		OptaID       string `json:"optaId"`
		OptaLegacyID string `json:"optaLegacyId"`
		Logo         struct {
			PublishURL string `json:"_publishUrl"`
			DynamicURL string `json:"_dynamicUrl"`
			DmS7URL    string `json:"_dmS7Url"`
		} `json:"logo"`
		Name         string `json:"name"`
		ShortName    string `json:"shortName"`
		OfficialName string `json:"officialName"`
		Tag          any    `json:"tag"`
	} `json:"homeTeam"`
	HomeTeamScoreTotal          string   `json:"homeTeamScoreTotal"`
	HomeTeamPenaltiesScoreTotal any      `json:"homeTeamPenaltiesScoreTotal"`
	HomeScorers                 []string `json:"homeScorers"`
	AwayTeam                    struct {
		OptaID       string `json:"optaId"`
		OptaLegacyID string `json:"optaLegacyId"`
		Logo         struct {
			PublishURL string `json:"_publishUrl"`
			DynamicURL string `json:"_dynamicUrl"`
			DmS7URL    string `json:"_dmS7Url"`
		} `json:"logo"`
		Name         string `json:"name"`
		ShortName    string `json:"shortName"`
		OfficialName string `json:"officialName"`
		Tag          any    `json:"tag"`
	} `json:"awayTeam"`
	AwayTeamScoreTotal          string   `json:"awayTeamScoreTotal"`
	AwayTeamPenaltiesScoreTotal any      `json:"awayTeamPenaltiesScoreTotal"`
	AwayScorers                 []string `json:"awayScorers"`
	Squad                       struct {
		ProvisionalSquad      bool     `json:"provisionalSquad"`
		Name                  string   `json:"name"`
		ShortName             string   `json:"shortName"`
		OfficialName          string   `json:"officialName"`
		SquadLabel            string   `json:"squadLabel"`
		ProvisionalSquadLabel any      `json:"provisionalSquadLabel"`
		Tag                   []string `json:"tag"`
		SellTicketsLimitTime  int      `json:"sellTicketsLimitTime"`
	} `json:"squad"`
	SquadList   []any `json:"squadList"`
	LineUp      []any `json:"lineUp"`
	Bench       []any `json:"bench"`
	RivalLineUp struct {
		HTML any `json:"html"`
	} `json:"rivalLineUp"`
	PlayAsHome        bool   `json:"playAsHome"`
	SoldOut           bool   `json:"soldOut"`
	TicketsLink       string `json:"ticketsLink"`
	FromPrice         int    `json:"fromPrice"`
	VipTickets        bool   `json:"vipTickets"`
	SoldOutVIP        bool   `json:"soldOutVIP"`
	VipTicketsLink    string `json:"vipTicketsLink"`
	FromPriceVIP      int    `json:"fromPriceVIP"`
	AdditionalLinks   []any  `json:"additionalLinks"`
	TicketInformation struct {
		PriceDateSaleLink             any    `json:"priceDateSaleLink"`
		ComingSoon                    bool   `json:"comingSoon"`
		ComingSoonVIP                 bool   `json:"comingSoonVIP"`
		GeneralSaleStartLabel         string `json:"generalSaleStartLabel"`
		VipTicketsCollective          any    `json:"vipTicketsCollective"`
		AwardedMembersCollective      any    `json:"awardedMembersCollective"`
		SeasonTicketHoldersCollective any    `json:"seasonTicketHoldersCollective"`
		ClubMembersCollective         struct {
			Discount      string    `json:"discount"`
			StartDateTime time.Time `json:"startDateTime"`
		} `json:"clubMembersCollective"`
		MadridistaPremiumCollective struct {
			Discount      any       `json:"discount"`
			StartDateTime time.Time `json:"startDateTime"`
		} `json:"madridistaPremiumCollective"`
		GeneralPublicCollective struct {
			Discount      any       `json:"discount"`
			StartDateTime time.Time `json:"startDateTime"`
		} `json:"generalPublicCollective"`
		MembersCollective      any    `json:"membersCollective"`
		AffectedUefaCollective any    `json:"affectedUefaCollective"`
		VipTicketsTitle        string `json:"vipTicketsTitle"`
		VipTicketsSubtitle     string `json:"vipTicketsSubtitle"`
		VipTicketsText         string `json:"vipTicketsText"`
		StadiumImage           struct {
			PublishURL string `json:"_publishUrl"`
			DynamicURL string `json:"_dynamicUrl"`
			DmS7URL    string `json:"_dmS7Url"`
		} `json:"stadiumImage"`
		StadiumZones []struct {
			ZoneName                    string `json:"zoneName"`
			Subzone1Name                string `json:"subzone1name"`
			Subzone1GeneralPublic       string `json:"subzone1generalPublic"`
			Subzone1Madridistas         any    `json:"subzone1madridistas"`
			Subzone1ClubMembers         string `json:"subzone1clubMembers"`
			Subzone1SeasonTicketHolders any    `json:"subzone1seasonTicketHolders"`
			Subzone1AwardedMembers      any    `json:"subzone1awardedMembers"`
			Subzone1Vip                 any    `json:"subzone1vip"`
			Subzone1Companion           any    `json:"subzone1companion"`
			Subzone2Name                string `json:"subzone2name"`
			Subzone2GeneralPublic       string `json:"subzone2generalPublic"`
			Subzone2Madridistas         any    `json:"subzone2madridistas"`
			Subzone2ClubMembers         string `json:"subzone2clubMembers"`
			Subzone2SeasonTicketHolders any    `json:"subzone2seasonTicketHolders"`
			Subzone2AwardedMembers      any    `json:"subzone2awardedMembers"`
			Subzone2Vip                 any    `json:"subzone2vip"`
			Subzone2Companion           any    `json:"subzone2companion"`
			Subzone3Name                string `json:"subzone3name"`
			Subzone3GeneralPublic       string `json:"subzone3generalPublic"`
			Subzone3Madridistas         any    `json:"subzone3madridistas"`
			Subzone3ClubMembers         string `json:"subzone3clubMembers"`
			Subzone3SeasonTicketHolders any    `json:"subzone3seasonTicketHolders"`
			Subzone3AwardedMembers      any    `json:"subzone3awardedMembers"`
			Subzone3Vip                 any    `json:"subzone3vip"`
			Subzone3Companion           any    `json:"subzone3companion"`
			Subzone4Name                string `json:"subzone4name"`
			Subzone4GeneralPublic       string `json:"subzone4generalPublic"`
			Subzone4Madridistas         any    `json:"subzone4madridistas"`
			Subzone4ClubMembers         string `json:"subzone4clubMembers"`
			Subzone4SeasonTicketHolders any    `json:"subzone4seasonTicketHolders"`
			Subzone4AwardedMembers      any    `json:"subzone4awardedMembers"`
			Subzone4Vip                 any    `json:"subzone4vip"`
			Subzone4Companion           any    `json:"subzone4companion"`
			Subzone5Name                string `json:"subzone5name"`
			Subzone5GeneralPublic       string `json:"subzone5generalPublic"`
			Subzone5Madridistas         any    `json:"subzone5madridistas"`
			Subzone5ClubMembers         string `json:"subzone5clubMembers"`
			Subzone5SeasonTicketHolders any    `json:"subzone5seasonTicketHolders"`
			Subzone5AwardedMembers      any    `json:"subzone5awardedMembers"`
			Subzone5Vip                 any    `json:"subzone5vip"`
			Subzone5Companion           any    `json:"subzone5companion"`
			Subzone6Name                string `json:"subzone6name"`
			Subzone6GeneralPublic       string `json:"subzone6generalPublic"`
			Subzone6Madridistas         any    `json:"subzone6madridistas"`
			Subzone6ClubMembers         string `json:"subzone6clubMembers"`
			Subzone6SeasonTicketHolders any    `json:"subzone6seasonTicketHolders"`
			Subzone6AwardedMembers      any    `json:"subzone6awardedMembers"`
			Subzone6Vip                 any    `json:"subzone6vip"`
			Subzone6Companion           any    `json:"subzone6companion"`
			Subzone7Name                string `json:"subzone7name"`
			Subzone7GeneralPublic       string `json:"subzone7generalPublic"`
			Subzone7Madridistas         any    `json:"subzone7madridistas"`
			Subzone7ClubMembers         string `json:"subzone7clubMembers"`
			Subzone7SeasonTicketHolders any    `json:"subzone7seasonTicketHolders"`
			Subzone7AwardedMembers      any    `json:"subzone7awardedMembers"`
			Subzone7Vip                 any    `json:"subzone7vip"`
			Subzone7Companion           any    `json:"subzone7companion"`
			Subzone8Name                any    `json:"subzone8name"`
			Subzone8GeneralPublic       any    `json:"subzone8generalPublic"`
			Subzone8Madridistas         any    `json:"subzone8madridistas"`
			Subzone8ClubMembers         any    `json:"subzone8clubMembers"`
			Subzone8SeasonTicketHolders any    `json:"subzone8seasonTicketHolders"`
			Subzone8AwardedMembers      any    `json:"subzone8awardedMembers"`
			Subzone8Vip                 any    `json:"subzone8vip"`
			Subzone8Companion           any    `json:"subzone8companion"`
			Subzone9Name                any    `json:"subzone9name"`
			Subzone9GeneralPublic       any    `json:"subzone9generalPublic"`
			Subzone9Madridistas         any    `json:"subzone9madridistas"`
			Subzone9ClubMembers         any    `json:"subzone9clubMembers"`
			Subzone9SeasonTicketHolders any    `json:"subzone9seasonTicketHolders"`
			Subzone9AwardedMembers      any    `json:"subzone9awardedMembers"`
			Subzone9Vip                 any    `json:"subzone9vip"`
			Subzone9Companion           any    `json:"subzone9companion"`
		} `json:"stadiumZones"`
		TitleGeneralDocuments string `json:"titleGeneralDocuments"`
		GeneralDocuments      []any  `json:"generalDocuments"`
		TitleFaqs             string `json:"titleFaqs"`
		ListFaqs              []struct {
			Title       string `json:"title"`
			Description struct {
				HTML string `json:"html"`
			} `json:"description"`
		} `json:"listFaqs"`
		AreaVipText struct {
			HTML string `json:"html"`
		} `json:"areaVipText"`
		TitleAreaVipDocuments string `json:"titleAreaVipDocuments"`
		AreaVipDocuments      []any  `json:"areaVipDocuments"`
		SeoTitle              any    `json:"seoTitle"`
		SeoDescription        any    `json:"seoDescription"`
	} `json:"ticketInformation"`
	IsAcademy      bool `json:"isAcademy"`
	HomeTeamName   any  `json:"homeTeamName"`
	HomeTeamLogo   any  `json:"homeTeamLogo"`
	AwayTeamName   any  `json:"awayTeamName"`
	AwayTeamLogo   any  `json:"awayTeamLogo"`
	SeoTitle       any  `json:"seoTitle"`
	SeoDescription any  `json:"seoDescription"`
	SeoRobots      any  `json:"seoRobots"`
	SeoCanonical   any  `json:"seoCanonical"`
	SeoHrefLang    any  `json:"seoHrefLang"`
}

// MatchesToday gets important football matches ocurring the current day
func (fm *FootballManager) GetMatchesToday() error {

	prev, _, next := utils.GetMonthStrings()

	url := fmt.Sprintf("https://publish-p47754-e237306.adobeaemcloud.com/graphql/execute.json/realmadridmastersite/diary;fromDate=2025-%s-31T22:00:00.000Z;toDate=2025-%s-30T23:59:00.000Z;alang=es-es", prev, next)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var events FootballEventsResponse
	if err := json.Unmarshal(body, &events); err != nil {
		return err
	}

	for _, match := range events.Data.MatchList.Items {
		if isImportant(match) {
			fm.MatchesToday = append(fm.MatchesToday, match)
		}
	}

	return nil
}

func isImportant(e Match) bool {
	if utils.IsToday(e.DateTime) &&
		e.Status == "pre_match" &&
		e.IsScheduled &&
		e.Squad.SquadLabel == "Fútbol · Primer Equipo" &&
		e.Venue.Name == "Santiago Bernabéu" {
		return true
	}

	return false
}
