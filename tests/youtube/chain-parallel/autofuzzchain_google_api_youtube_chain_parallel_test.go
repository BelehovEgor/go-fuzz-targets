package chain_parallel

// Edit if desired. Code generated by "fzgen -chain -parallel google.golang.org/api/youtube/v3".

import (
	"testing"

	"github.com/thepudds/fzgen/fuzzer"
	"google.golang.org/api/youtube/v3"
)

func Fuzz_NewAbuseReportsService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewAbuseReportsService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_AbuseReportsService_Insert",
				Func: func(part []string, abusereport *youtube.AbuseReport) *youtube.AbuseReportsInsertCall {
					return target.Insert(part, abusereport)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewActivitiesService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewActivitiesService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_ActivitiesService_List",
				Func: func(part []string) *youtube.ActivitiesListCall {
					return target.List(part)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewCaptionsService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewCaptionsService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_CaptionsService_Delete",
				Func: func(id string) *youtube.CaptionsDeleteCall {
					return target.Delete(id)
				},
			},
			{
				Name: "Fuzz_CaptionsService_Download",
				Func: func(id string) *youtube.CaptionsDownloadCall {
					return target.Download(id)
				},
			},
			{
				Name: "Fuzz_CaptionsService_Insert",
				Func: func(part []string, caption *youtube.Caption) *youtube.CaptionsInsertCall {
					return target.Insert(part, caption)
				},
			},
			{
				Name: "Fuzz_CaptionsService_List",
				Func: func(part []string, videoId string) *youtube.CaptionsListCall {
					return target.List(part, videoId)
				},
			},
			{
				Name: "Fuzz_CaptionsService_Update",
				Func: func(part []string, caption *youtube.Caption) *youtube.CaptionsUpdateCall {
					return target.Update(part, caption)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewChannelBannersService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewChannelBannersService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_ChannelBannersService_Insert",
				Func: func(channelbannerresource *youtube.ChannelBannerResource) *youtube.ChannelBannersInsertCall {
					return target.Insert(channelbannerresource)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewChannelSectionsService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewChannelSectionsService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_ChannelSectionsService_Delete",
				Func: func(id string) *youtube.ChannelSectionsDeleteCall {
					return target.Delete(id)
				},
			},
			{
				Name: "Fuzz_ChannelSectionsService_Insert",
				Func: func(part []string, channelsection *youtube.ChannelSection) *youtube.ChannelSectionsInsertCall {
					return target.Insert(part, channelsection)
				},
			},
			{
				Name: "Fuzz_ChannelSectionsService_List",
				Func: func(part []string) *youtube.ChannelSectionsListCall {
					return target.List(part)
				},
			},
			{
				Name: "Fuzz_ChannelSectionsService_Update",
				Func: func(part []string, channelsection *youtube.ChannelSection) *youtube.ChannelSectionsUpdateCall {
					return target.Update(part, channelsection)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewChannelsService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewChannelsService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_ChannelsService_List",
				Func: func(part []string) *youtube.ChannelsListCall {
					return target.List(part)
				},
			},
			{
				Name: "Fuzz_ChannelsService_Update",
				Func: func(part []string, channel *youtube.Channel) *youtube.ChannelsUpdateCall {
					return target.Update(part, channel)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewCommentThreadsService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewCommentThreadsService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_CommentThreadsService_Insert",
				Func: func(part []string, commentthread *youtube.CommentThread) *youtube.CommentThreadsInsertCall {
					return target.Insert(part, commentthread)
				},
			},
			{
				Name: "Fuzz_CommentThreadsService_List",
				Func: func(part []string) *youtube.CommentThreadsListCall {
					return target.List(part)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewCommentsService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewCommentsService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_CommentsService_Delete",
				Func: func(id string) *youtube.CommentsDeleteCall {
					return target.Delete(id)
				},
			},
			{
				Name: "Fuzz_CommentsService_Insert",
				Func: func(part []string, comment *youtube.Comment) *youtube.CommentsInsertCall {
					return target.Insert(part, comment)
				},
			},
			{
				Name: "Fuzz_CommentsService_List",
				Func: func(part []string) *youtube.CommentsListCall {
					return target.List(part)
				},
			},
			{
				Name: "Fuzz_CommentsService_MarkAsSpam",
				Func: func(id []string) *youtube.CommentsMarkAsSpamCall {
					return target.MarkAsSpam(id)
				},
			},
			{
				Name: "Fuzz_CommentsService_SetModerationStatus",
				Func: func(id []string, moderationStatus string) *youtube.CommentsSetModerationStatusCall {
					return target.SetModerationStatus(id, moderationStatus)
				},
			},
			{
				Name: "Fuzz_CommentsService_Update",
				Func: func(part []string, comment *youtube.Comment) *youtube.CommentsUpdateCall {
					return target.Update(part, comment)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewI18nLanguagesService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewI18nLanguagesService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_I18nLanguagesService_List",
				Func: func(part []string) *youtube.I18nLanguagesListCall {
					return target.List(part)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewI18nRegionsService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewI18nRegionsService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_I18nRegionsService_List",
				Func: func(part []string) *youtube.I18nRegionsListCall {
					return target.List(part)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewLiveBroadcastsService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewLiveBroadcastsService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_LiveBroadcastsService_Bind",
				Func: func(id string, part []string) *youtube.LiveBroadcastsBindCall {
					return target.Bind(id, part)
				},
			},
			{
				Name: "Fuzz_LiveBroadcastsService_Delete",
				Func: func(id string) *youtube.LiveBroadcastsDeleteCall {
					return target.Delete(id)
				},
			},
			{
				Name: "Fuzz_LiveBroadcastsService_Insert",
				Func: func(part []string, livebroadcast *youtube.LiveBroadcast) *youtube.LiveBroadcastsInsertCall {
					return target.Insert(part, livebroadcast)
				},
			},
			{
				Name: "Fuzz_LiveBroadcastsService_InsertCuepoint",
				Func: func(cuepoint *youtube.Cuepoint) *youtube.LiveBroadcastsInsertCuepointCall {
					return target.InsertCuepoint(cuepoint)
				},
			},
			{
				Name: "Fuzz_LiveBroadcastsService_List",
				Func: func(part []string) *youtube.LiveBroadcastsListCall {
					return target.List(part)
				},
			},
			{
				Name: "Fuzz_LiveBroadcastsService_Transition",
				Func: func(broadcastStatus string, id string, part []string) *youtube.LiveBroadcastsTransitionCall {
					return target.Transition(broadcastStatus, id, part)
				},
			},
			{
				Name: "Fuzz_LiveBroadcastsService_Update",
				Func: func(part []string, livebroadcast *youtube.LiveBroadcast) *youtube.LiveBroadcastsUpdateCall {
					return target.Update(part, livebroadcast)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewLiveChatBansService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewLiveChatBansService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_LiveChatBansService_Delete",
				Func: func(id string) *youtube.LiveChatBansDeleteCall {
					return target.Delete(id)
				},
			},
			{
				Name: "Fuzz_LiveChatBansService_Insert",
				Func: func(part []string, livechatban *youtube.LiveChatBan) *youtube.LiveChatBansInsertCall {
					return target.Insert(part, livechatban)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewLiveChatMessagesService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewLiveChatMessagesService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_LiveChatMessagesService_Delete",
				Func: func(id string) *youtube.LiveChatMessagesDeleteCall {
					return target.Delete(id)
				},
			},
			{
				Name: "Fuzz_LiveChatMessagesService_Insert",
				Func: func(part []string, livechatmessage *youtube.LiveChatMessage) *youtube.LiveChatMessagesInsertCall {
					return target.Insert(part, livechatmessage)
				},
			},
			{
				Name: "Fuzz_LiveChatMessagesService_List",
				Func: func(liveChatId string, part []string) *youtube.LiveChatMessagesListCall {
					return target.List(liveChatId, part)
				},
			},
			{
				Name: "Fuzz_LiveChatMessagesService_Transition",
				Func: func() *youtube.LiveChatMessagesTransitionCall {
					return target.Transition()
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewLiveChatModeratorsService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewLiveChatModeratorsService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_LiveChatModeratorsService_Delete",
				Func: func(id string) *youtube.LiveChatModeratorsDeleteCall {
					return target.Delete(id)
				},
			},
			{
				Name: "Fuzz_LiveChatModeratorsService_Insert",
				Func: func(part []string, livechatmoderator *youtube.LiveChatModerator) *youtube.LiveChatModeratorsInsertCall {
					return target.Insert(part, livechatmoderator)
				},
			},
			{
				Name: "Fuzz_LiveChatModeratorsService_List",
				Func: func(liveChatId string, part []string) *youtube.LiveChatModeratorsListCall {
					return target.List(liveChatId, part)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewLiveStreamsService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewLiveStreamsService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_LiveStreamsService_Delete",
				Func: func(id string) *youtube.LiveStreamsDeleteCall {
					return target.Delete(id)
				},
			},
			{
				Name: "Fuzz_LiveStreamsService_Insert",
				Func: func(part []string, livestream *youtube.LiveStream) *youtube.LiveStreamsInsertCall {
					return target.Insert(part, livestream)
				},
			},
			{
				Name: "Fuzz_LiveStreamsService_List",
				Func: func(part []string) *youtube.LiveStreamsListCall {
					return target.List(part)
				},
			},
			{
				Name: "Fuzz_LiveStreamsService_Update",
				Func: func(part []string, livestream *youtube.LiveStream) *youtube.LiveStreamsUpdateCall {
					return target.Update(part, livestream)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewMembersService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewMembersService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_MembersService_List",
				Func: func(part []string) *youtube.MembersListCall {
					return target.List(part)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewMembershipsLevelsService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewMembershipsLevelsService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_MembershipsLevelsService_List",
				Func: func(part []string) *youtube.MembershipsLevelsListCall {
					return target.List(part)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewPlaylistImagesService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewPlaylistImagesService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_PlaylistImagesService_Delete",
				Func: func() *youtube.PlaylistImagesDeleteCall {
					return target.Delete()
				},
			},
			{
				Name: "Fuzz_PlaylistImagesService_Insert",
				Func: func(playlistimage *youtube.PlaylistImage) *youtube.PlaylistImagesInsertCall {
					return target.Insert(playlistimage)
				},
			},
			{
				Name: "Fuzz_PlaylistImagesService_List",
				Func: func() *youtube.PlaylistImagesListCall {
					return target.List()
				},
			},
			{
				Name: "Fuzz_PlaylistImagesService_Update",
				Func: func(playlistimage *youtube.PlaylistImage) *youtube.PlaylistImagesUpdateCall {
					return target.Update(playlistimage)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewPlaylistItemsService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewPlaylistItemsService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_PlaylistItemsService_Delete",
				Func: func(id string) *youtube.PlaylistItemsDeleteCall {
					return target.Delete(id)
				},
			},
			{
				Name: "Fuzz_PlaylistItemsService_Insert",
				Func: func(part []string, playlistitem *youtube.PlaylistItem) *youtube.PlaylistItemsInsertCall {
					return target.Insert(part, playlistitem)
				},
			},
			{
				Name: "Fuzz_PlaylistItemsService_List",
				Func: func(part []string) *youtube.PlaylistItemsListCall {
					return target.List(part)
				},
			},
			{
				Name: "Fuzz_PlaylistItemsService_Update",
				Func: func(part []string, playlistitem *youtube.PlaylistItem) *youtube.PlaylistItemsUpdateCall {
					return target.Update(part, playlistitem)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewPlaylistsService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewPlaylistsService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_PlaylistsService_Delete",
				Func: func(id string) *youtube.PlaylistsDeleteCall {
					return target.Delete(id)
				},
			},
			{
				Name: "Fuzz_PlaylistsService_Insert",
				Func: func(part []string, playlist *youtube.Playlist) *youtube.PlaylistsInsertCall {
					return target.Insert(part, playlist)
				},
			},
			{
				Name: "Fuzz_PlaylistsService_List",
				Func: func(part []string) *youtube.PlaylistsListCall {
					return target.List(part)
				},
			},
			{
				Name: "Fuzz_PlaylistsService_Update",
				Func: func(part []string, playlist *youtube.Playlist) *youtube.PlaylistsUpdateCall {
					return target.Update(part, playlist)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewSearchService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewSearchService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_SearchService_List",
				Func: func(part []string) *youtube.SearchListCall {
					return target.List(part)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewSubscriptionsService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewSubscriptionsService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_SubscriptionsService_Delete",
				Func: func(id string) *youtube.SubscriptionsDeleteCall {
					return target.Delete(id)
				},
			},
			{
				Name: "Fuzz_SubscriptionsService_Insert",
				Func: func(part []string, subscription *youtube.Subscription) *youtube.SubscriptionsInsertCall {
					return target.Insert(part, subscription)
				},
			},
			{
				Name: "Fuzz_SubscriptionsService_List",
				Func: func(part []string) *youtube.SubscriptionsListCall {
					return target.List(part)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewSuperChatEventsService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewSuperChatEventsService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_SuperChatEventsService_List",
				Func: func(part []string) *youtube.SuperChatEventsListCall {
					return target.List(part)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewTestsService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewTestsService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_TestsService_Insert",
				Func: func(part []string, testitem *youtube.TestItem) *youtube.TestsInsertCall {
					return target.Insert(part, testitem)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewThirdPartyLinksService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewThirdPartyLinksService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_ThirdPartyLinksService_Delete",
				Func: func(linkingToken string, type_ string) *youtube.ThirdPartyLinksDeleteCall {
					return target.Delete(linkingToken, type_)
				},
			},
			{
				Name: "Fuzz_ThirdPartyLinksService_Insert",
				Func: func(part []string, thirdpartylink *youtube.ThirdPartyLink) *youtube.ThirdPartyLinksInsertCall {
					return target.Insert(part, thirdpartylink)
				},
			},
			{
				Name: "Fuzz_ThirdPartyLinksService_List",
				Func: func(part []string) *youtube.ThirdPartyLinksListCall {
					return target.List(part)
				},
			},
			{
				Name: "Fuzz_ThirdPartyLinksService_Update",
				Func: func(part []string, thirdpartylink *youtube.ThirdPartyLink) *youtube.ThirdPartyLinksUpdateCall {
					return target.Update(part, thirdpartylink)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewThumbnailsService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewThumbnailsService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_ThumbnailsService_Set",
				Func: func(videoId string) *youtube.ThumbnailsSetCall {
					return target.Set(videoId)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewVideoAbuseReportReasonsService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewVideoAbuseReportReasonsService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_VideoAbuseReportReasonsService_List",
				Func: func(part []string) *youtube.VideoAbuseReportReasonsListCall {
					return target.List(part)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewVideoCategoriesService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewVideoCategoriesService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_VideoCategoriesService_List",
				Func: func(part []string) *youtube.VideoCategoriesListCall {
					return target.List(part)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewVideosService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewVideosService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_VideosService_Delete",
				Func: func(id string) *youtube.VideosDeleteCall {
					return target.Delete(id)
				},
			},
			{
				Name: "Fuzz_VideosService_GetRating",
				Func: func(id []string) *youtube.VideosGetRatingCall {
					return target.GetRating(id)
				},
			},
			{
				Name: "Fuzz_VideosService_Insert",
				Func: func(part []string, video *youtube.Video) *youtube.VideosInsertCall {
					return target.Insert(part, video)
				},
			},
			{
				Name: "Fuzz_VideosService_List",
				Func: func(part []string) *youtube.VideosListCall {
					return target.List(part)
				},
			},
			{
				Name: "Fuzz_VideosService_Rate",
				Func: func(id string, rating string) *youtube.VideosRateCall {
					return target.Rate(id, rating)
				},
			},
			{
				Name: "Fuzz_VideosService_ReportAbuse",
				Func: func(videoabusereport *youtube.VideoAbuseReport) *youtube.VideosReportAbuseCall {
					return target.ReportAbuse(videoabusereport)
				},
			},
			{
				Name: "Fuzz_VideosService_Update",
				Func: func(part []string, video *youtube.Video) *youtube.VideosUpdateCall {
					return target.Update(part, video)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewWatermarksService_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewWatermarksService(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_WatermarksService_Set",
				Func: func(channelId string, invideobranding *youtube.InvideoBranding) *youtube.WatermarksSetCall {
					return target.Set(channelId, invideobranding)
				},
			},
			{
				Name: "Fuzz_WatermarksService_Unset",
				Func: func(channelId string) *youtube.WatermarksUnsetCall {
					return target.Unset(channelId)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}

func Fuzz_NewYoutubeV3Service_Chain(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var s *youtube.Service
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s)
		if s == nil {
			return
		}

		target := youtube.NewYoutubeV3Service(s)

		steps := []fuzzer.Step{
			{
				Name: "Fuzz_YoutubeV3Service_UpdateCommentThreads",
				Func: func(commentthread *youtube.CommentThread) *youtube.YoutubeV3UpdateCommentThreadsCall {
					return target.UpdateCommentThreads(commentthread)
				},
			},
		}

		// Execute a specific chain of steps, with the count, sequence and arguments controlled by fz.Chain
		fz.Chain(steps, fuzzer.ChainParallel)
	})
}
