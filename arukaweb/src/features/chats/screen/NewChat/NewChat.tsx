import { SendIcon } from "@assets/icons/send";
import { useProviders } from "@common/data/use-providers";
import {
  ActionIcon,
  Box,
  Card,
  Container,
  Fieldset,
  Group,
  NativeSelect,
  SimpleGrid,
  Textarea,
} from "@mantine/core";
import { useNavigate } from "@tanstack/react-router";
import { useNewChatForm, type NewChatFormValues } from "./useNewChatForm";
import { useNewChatMutation } from "./useNewChatMutation";
import { DEFAULT_SELECT_MODEL } from "./utils";

export function NewChat() {
  const navigate = useNavigate();
  const { providers } = useProviders();
  const form = useNewChatForm();
  const { mutate, isPending } = useNewChatMutation();

  const handleSubmit = (values: NewChatFormValues) => {
    mutate(values, {
      onSuccess: (data) => {
        if (!data.chat) return;
        navigate({ to: "/chat/$chatId", params: { chatId: data.chat.id } });
      },
    });
  };

  return (
    <Container>
      <Box h="200px" />
      <Card shadow="sm">
        <form onSubmit={form.onSubmit(handleSubmit)}>
          <Textarea
            key={form.key("userText")}
            placeholder="Ask anything..."
            variant="unstyled"
            autosize
            minRows={3}
            maxRows={8}
            size="md"
            autoFocus
            onKeyDown={(e) => {
              if (e.key === "Enter" && e.ctrlKey) {
                e.preventDefault();
                form.onSubmit(handleSubmit)();
              }
            }}
            {...form.getInputProps("userText")}
          />

          <Fieldset legend="General settings" variant="unstyled">
            <Textarea
              key={form.key("general.customPrompt")}
              placeholder="Custom prompt..."
              label="Custom prompt"
              autosize
              minRows={3}
              maxRows={6}
              size="sm"
              {...form.getInputProps("general.customPrompt")}
            />
            <SimpleGrid cols={2}>
              <NativeSelect
                size="sm"
                label="Provider"
                key={form.key("general.modelId")}
                {...form.getInputProps("general.modelId")}
              >
                <option value={DEFAULT_SELECT_MODEL} disabled>
                  Select Provider
                </option>
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

          <Group justify="end">
            <ActionIcon size="xl" radius="sm" type="submit" loading={isPending}>
              <SendIcon width={24} height={24} />
            </ActionIcon>
          </Group>
        </form>
      </Card>
    </Container>
  );
}
