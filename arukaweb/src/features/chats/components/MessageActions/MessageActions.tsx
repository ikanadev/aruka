import { ExpandHorizontal } from "@assets/icons/expand-horizontal";
import { ActionIcon, Container, Group } from "@mantine/core";

interface Props {
  toggleExpand: () => void;
}

export function MessageActions(props: Props) {
  const { toggleExpand } = props;
  return (
    <Container size="md">
      <Group justify="flex-end">
        <ActionIcon onClick={toggleExpand} variant="light" size="md">
          <ExpandHorizontal width="1.2em" height="1.2em" />
        </ActionIcon>
      </Group>
    </Container>
  );
}
