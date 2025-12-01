import { useChats } from "@features/chats/data/use-chats";
import { Box, Button, Stack, Title } from "@mantine/core";
import { Link } from "@tanstack/react-router";
import { ChatNavbarLink } from "../ChatNavbarLink/ChatNavbarLink";
import { AddIcon } from "@assets/icons/add";

export function ChatNavbar() {
  const { chats } = useChats();

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
          <ChatNavbarLink key={chat.id} chatId={chat.id} providerName={chat.provider?.name} text={chat.title} />
        ))}
      </Stack>
    </Stack>
  );
}
