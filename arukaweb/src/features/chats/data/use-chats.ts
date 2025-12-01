import { chatClient } from "@common/utils/clients";
import { useInfiniteQuery } from "@tanstack/react-query";
import { chatQueryKeys } from "./chat-query-keys";

const ITEMS_PER_PAGE = 10;

export function useChats() {
  const query = useInfiniteQuery({
    queryFn: async ({ pageParam }) => {
      return chatClient.listChats({
        pagination: {
          page: pageParam,
          limit: ITEMS_PER_PAGE,
        },
      });
    },
    initialPageParam: 1,
    getNextPageParam: (lastPage) => {
      if (!lastPage.pagination) return null;
      if (!lastPage.pagination.hasNext) return null;
      return lastPage.pagination.page + 1;
    },
    queryKey: chatQueryKeys.chats(),
  });

  const chats = query.data?.pages.flat().map((chat) => chat.chats).flat() || [];

  return {
    chats,
    loadingChats: query.isLoading,
    fetchingChats: query.isFetching,
    fetchingNextPageChats: query.isFetchingNextPage,
    fetchNextPageChats: query.fetchNextPage,
    hasNextPageChats: query.hasNextPage,
  }
}
