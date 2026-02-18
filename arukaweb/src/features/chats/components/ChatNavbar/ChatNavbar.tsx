import { AddIcon } from "@assets/icons/add";
import { useChats } from "@features/chats/data/use-chats";
import { useUpdateChatTitle } from "@features/chats/data/use-update-chat-title";
import { Box, Button, Stack, Title } from "@mantine/core";
import { Link } from "@tanstack/react-router";
import { ChatNavbarLink } from "../ChatNavbarLink/ChatNavbarLink";

export function ChatNavbar() {
  const { chats } = useChats();
  const { updateChatTitle, updatingChatTitle, updatingChatId } = useUpdateChatTitle();

  return (
    <Stack p="sm">
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
      <Stack gap={0}>
        {chats.map((chat) => (
          <ChatNavbarLink
            key={chat.id}
            chatId={chat.id}
            providerName={chat.provider?.name}
            text={chat.title}
            onUpdateTitle={() => updateChatTitle(chat.id)}
            updatingTitle={updatingChatTitle && updatingChatId === chat.id}
          />
        ))}
      </Stack>
    </Stack>
  );
}
