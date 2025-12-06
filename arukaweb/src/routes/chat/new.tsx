import { SendIcon } from '@assets/icons/send';
import { useProviders } from '@common/data/use-providers';
import { ChatTextAreaInput } from '@features/chats/components/ChatTextAreaInput/ChatTextAreaInput';
import { ActionIcon, Box, Card, Container, Fieldset, Group, NativeSelect, SimpleGrid, Textarea } from '@mantine/core';
import { useForm } from '@mantine/form';
import { createFileRoute } from '@tanstack/react-router'
import { useEffect, useRef } from 'react';

export const Route = createFileRoute('/chat/new')({
  component: RouteComponent,
})

const DEFAULT_SELECT_MODEL = 'DEFAULT_SELECT_MODEL';

function RouteComponent() {
  const textAreaRef = useRef<HTMLTextAreaElement>(null);
  const { providers } = useProviders();

  const form = useForm({
    mode: 'uncontrolled',
    initialValues: {
      userText: '',
      general: {
        modelId: DEFAULT_SELECT_MODEL,
        customPrompt: '',
      },
    },
    validate: {
      userText: (value) => (value ? null : 'Please enter a message'),
      general: {
        modelId: (value) => (value === DEFAULT_SELECT_MODEL ? 'Please select a provider' : null),
      },
    },
  });

  const handleSubmit = (values: typeof form.values) => {
    console.log(values);
  }

  useEffect(() => {
    if (!textAreaRef.current) return;
    textAreaRef.current.focus();
  }, []);
  return (
    <Container>
      <Box h="200px" />
      <Card shadow="sm">
        <form onSubmit={form.onSubmit(handleSubmit)}>
          <ChatTextAreaInput textAreaRef={textAreaRef} />

          <Fieldset legend="General settings" variant='unstyled'>
            <Textarea
              key={form.key('general.customPrompt')}
              placeholder="Custom prompt..."
              label="Custom prompt"
              autosize
              minRows={3}
              maxRows={6}
              size="sm"
              {...form.getInputProps('general.customPrompt')}
            />
            <SimpleGrid cols={2}>
              <NativeSelect
                size="sm"
                label="Provider"
                key={form.key('general.modelId')}
                {...form.getInputProps('general.modelId')}
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
            <ActionIcon size="xl" radius="sm" type="submit">
              <SendIcon width={24} height={24} />
            </ActionIcon>
          </Group>
        </form>
      </Card>
    </Container>
  );
}
