import { SendIcon } from '@assets/icons/send';
import { useProviders } from '@common/data/use-providers';
import { ChatTextAreaInput } from '@features/chats/components/ChatTextAreaInput/ChatTextAreaInput';
import { ActionIcon, Box, Button, Card, Container, Fieldset, Group, NativeSelect, SimpleGrid, Textarea } from '@mantine/core';
import { createFileRoute } from '@tanstack/react-router'
import { useEffect, useRef, useState } from 'react';

export const Route = createFileRoute('/chat/new')({
  component: RouteComponent,
})

const DEFAULT_SELECT_MODEL = 'DEFAULT_SELECT_MODEL';

function RouteComponent() {
  const textAreaRef = useRef<HTMLTextAreaElement>(null);
  const [modelId, setModelId] = useState(DEFAULT_SELECT_MODEL);
  const { providers } = useProviders();
  console.log({ modelId });

  useEffect(() => {
    if (!textAreaRef.current) return;
    textAreaRef.current.focus();
  }, []);
  return (
    <Container>
      <Box h="200px" />
      <Card shadow="sm">
        <ChatTextAreaInput textAreaRef={textAreaRef} />

        <Fieldset legend="General settings" variant='unstyled'>
          <Textarea placeholder="Custom prompt..." label="Custom prompt" autosize minRows={3} maxRows={6} size="sm" />
          <SimpleGrid cols={2}>
            <NativeSelect
              size="sm"
              label="Provider"
              value={modelId}
              onChange={(e) => setModelId(e.currentTarget.value)}
            >
              <option value={DEFAULT_SELECT_MODEL} disabled>Select Provider</option>
              {providers.map((provider) => (
                <optgroup key={provider.id} label={provider.name}>
                  {provider.models.map((model) => (
                    <option key={model.id} value={model.id}>
                      {model.name}
                    </option>
                  ))}
                </optgroup>
              ))}
            </NativeSelect>
          </SimpleGrid>
        </Fieldset>

        <Group justify='end'>
          <ActionIcon size="xl" radius="sm">
            <SendIcon width={24} height={24} />
          </ActionIcon>
        </Group>
      </Card>
    </Container>
  );
}
