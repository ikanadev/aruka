import { AnthropicIcon } from "@assets/brands/AnthropicIcon";
import { GeminiIcon } from "@assets/brands/GeminiIcon";
import { OpenAIIcon } from "@assets/brands/OpenAIIcon";
import { DotsVerticalIcon } from "@assets/icons/dots-vertical";
import { ActionIcon, Button, Group, Loader, Menu } from "@mantine/core";
import { Link, useMatchRoute } from "@tanstack/react-router";
import { type JSX, type MouseEvent } from "react";

interface Props {
  chatId: string;
  providerName?: string;
  text: string;
  onUpdateTitle: () => void;
  updatingTitle: boolean;
  onDelete: () => void;
}

const providerIconMap: Record<string, JSX.Element> = {
  Anthropic: <AnthropicIcon width={16} height={16} fill="currentColor" />,
  Google: <GeminiIcon width={16} height={16} />,
  OpenAI: <OpenAIIcon width={16} height={16} />,
};

export function ChatNavbarLink(props: Props) {
  const { chatId, providerName = "", text, onUpdateTitle, updatingTitle, onDelete } = props;
  const matchRoute = useMatchRoute();
  const match = matchRoute({ to: "/chat/$chatId", params: { chatId: chatId } });

  const handleMenuClick = (e: MouseEvent) => {
    e.preventDefault();
    e.stopPropagation();
  };

  return (
    <Group align="center" gap={0}>
      <Button
        key={chatId}
        component={Link}
        variant={match ? "light" : "subtle"}
        justify="start"
        color="gray"
        size="xs"
        to="/chat/$chatId"
        // @ts-expect-error error due using Button as container
        params={{ chatId: chatId }}
        leftSection={providerIconMap[providerName]}
        styles={{
          label: { fontWeight: 400 },
          root: { flex: 1, overflow: "hidden" },
        }}
      >
        {text}
      </Button>
      <Menu position="bottom-end" withArrow>
        <Menu.Target>
          <ActionIcon
            variant="subtle"
            color="gray"
            size="xs"
            onClick={handleMenuClick}
          >
            {updatingTitle ? (
              <Loader size={12} />
            ) : (
              <DotsVerticalIcon width={14} height={14} />
            )}
          </ActionIcon>
        </Menu.Target>
        <Menu.Dropdown>
          <Menu.Item onClick={onUpdateTitle} disabled={updatingTitle}>
            Update chat title
          </Menu.Item>
          <Menu.Item color="red" onClick={onDelete}>
            Remove
          </Menu.Item>
        </Menu.Dropdown>
      </Menu>
    </Group>
  );
}
