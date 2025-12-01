import { useComputedColorScheme } from "@mantine/core";

export function useIsDark() {
  const theme = useComputedColorScheme();
  return theme === "dark";
}
