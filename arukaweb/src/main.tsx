import { createRoot } from 'react-dom/client'
import { localStorageColorSchemeManager, MantineProvider } from '@mantine/core'
import { createRouter, RouterProvider } from '@tanstack/react-router'

import { routeTree } from './routeTree.gen.ts'

import '@mantine/core/styles.css';
import './index.css'
import { theme } from '@common/utils/theme.ts';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

const router = createRouter({
  routeTree: routeTree,
  defaultPreload: 'intent',
  scrollRestoration: true,
});

declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router;
  }
}

const rootElement = document.getElementById('root')!;

const queryClient = new QueryClient();

const colorSchemeManager = localStorageColorSchemeManager({
  key: 'mantine-color-scheme',
});

if (!rootElement.innerHTML) {
  const root = createRoot(rootElement);
  root.render(
    <QueryClientProvider client={queryClient}>
      <MantineProvider theme={theme} colorSchemeManager={colorSchemeManager} defaultColorScheme='auto'>
        <RouterProvider router={router} />
      </MantineProvider>
    </QueryClientProvider>
  );
}
