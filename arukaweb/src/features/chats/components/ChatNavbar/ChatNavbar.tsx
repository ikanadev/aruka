import { AddIcon } from "@assets/icons/add";
import { ChevronDownIcon } from "@assets/icons/chevron-down";
import { useChats } from "@features/chats/data/use-chats";
import { useDeleteChat } from "@features/chats/data/use-delete-chat";
import { useUpdateChatTitle } from "@features/chats/data/use-update-chat-title";
import { AppShell, Box, Button, ScrollArea, Stack, Title } from "@mantine/core";
import { Link, useMatchRoute, useNavigate } from "@tanstack/react-router";
import { ChatNavbarLink } from "../ChatNavbarLink/ChatNavbarLink";

export function ChatNavbar() {
  const {
    chats,
    fetchNextPageChats,
    hasNextPageChats,
    fetchingNextPageChats,
  } = useChats();
  const { updateChatTitle, updatingChatTitle, updatingChatId } = useUpdateChatTitle();
  const { deleteChat } = useDeleteChat();
  const matchRoute = useMatchRoute();
  const navigate = useNavigate();

  const handleDelete = (chatId: string) => {
    const isActive = matchRoute({ to: "/chat/$chatId", params: { chatId } });
    deleteChat(chatId, {
      onSuccess: () => {
        if (isActive) navigate({ to: "/chat/new" });
      },
    });
  };

  return (
    <>
      <AppShell.Section p="sm">
        <Title order={1}>Aruka</Title>
        <Box h="40px" />
        <Button
          component={Link}
          variant="outline"
          to="/chat/new"
          leftSection={<AddIcon width={18} height={18} />}
        >
          New Chat
        </Button>
      </AppShell.Section>
      <AppShell.Section grow component={ScrollArea} p="sm" pt={0}>
        <Stack gap={0}>
          {chats.map((chat) => (
            <ChatNavbarLink
              key={chat.id}
              chatId={chat.id}
              providerName={chat.provider?.name}
              text={chat.title}
              onUpdateTitle={() => updateChatTitle(chat.id)}
              updatingTitle={updatingChatTitle && updatingChatId === chat.id}
              onDelete={() => handleDelete(chat.id)}
            />
          ))}
        </Stack>
        {hasNextPageChats && (
          <Button
            variant="subtle"
            size="sm"
            fullWidth
            mt="xs"
            leftSection={<ChevronDownIcon width={16} height={16} />}
            loading={fetchingNextPageChats}
            onClick={() => fetchNextPageChats()}
          >
            Load more
          </Button>
        )}
      </AppShell.Section>
      <AppShell.Section />
    </>
  );
}
