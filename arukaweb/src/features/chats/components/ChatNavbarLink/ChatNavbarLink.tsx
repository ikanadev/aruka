import { AnthropicIcon } from "@assets/brands/AnthropicIcon";
import { GeminiIcon } from "@assets/brands/GeminiIcon";
import { OpenAIIcon } from "@assets/brands/OpenAIIcon";
import { Button } from "@mantine/core";
import { Link, useMatchRoute } from "@tanstack/react-router";
import { type JSX } from "react";

interface Props {
  chatId: string;
  providerName?: string;
  text: string;
}

const providerIconMap: Record<string, JSX.Element> = {
  Anthropic: <AnthropicIcon width={16} height={16} fill="currentColor" />,
  Google: <GeminiIcon width={16} height={16} />,
  OpenAI: <OpenAIIcon width={16} height={16} />,
};



export function ChatNavbarLink(props: Props) {
  const { chatId, providerName = '', text } = props;
  const matchRoute = useMatchRoute();
  const match = matchRoute({ to: '/chat/$chatId', params: { chatId: chatId } });

  return (
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
      }}
    >
      {text}
    </Button>
  );
}
