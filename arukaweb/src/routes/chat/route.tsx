import { ChevronDoubleLeft } from '@assets/icons/chevron-double-left';
import { MiChevronDoubleRight } from '@assets/icons/chevron-double-right';
import { ChatNavbar } from '@features/chats/components/ChatNavbar/ChatNavbar';
import { ActionIcon, AppShell, Box } from '@mantine/core'
import { useDisclosure } from '@mantine/hooks'
import { createFileRoute, Outlet } from '@tanstack/react-router'

export const Route = createFileRoute('/chat')({
  component: RouteComponent,
})

function RouteComponent() {
  const [opened, { toggle }] = useDisclosure();
  return (
    <AppShell
      padding={{ base: 10, sm: 15, lg: 'md' }}
      navbar={{
        width: 350,
        breakpoint: 'sm',
        collapsed: { mobile: !opened, desktop: !opened },
      }}
    >
      <AppShell.Navbar>
        <ChatNavbar />
      </AppShell.Navbar>
      <AppShell.Main style={{ '--app-shell-padding': '0rem', height: '100dvh' }}>
        <Box pos="relative">
          <Box bg="var(--mantine-color-body)" pos="absolute" top={20} left={20} style={{ zIndex: 10 }}>
            <ActionIcon onClick={toggle} variant="light">
              {opened ? (
                <ChevronDoubleLeft />
              ) : (
                <MiChevronDoubleRight />
              )}
            </ActionIcon>
          </Box>
        </Box>
        <Outlet />
      </AppShell.Main>
    </AppShell>
  );
}
