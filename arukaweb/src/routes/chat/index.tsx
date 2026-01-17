import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/chat/")({
  component: Index,
});

function Index() {
  return (
    <div>
      <h1>Index</h1>
    </div>
  );
}
