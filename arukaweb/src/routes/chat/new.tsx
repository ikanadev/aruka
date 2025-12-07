import { NewChat } from "@features/chats/screen/NewChat/NewChat";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/chat/new")({
  component: RouteComponent,
});

function RouteComponent() {
  return <NewChat />;
}
