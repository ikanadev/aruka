import { Chat } from '@features/chats/screen/Chat/Chat'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/chat/$chatId')({
  component: RouteComponent,
})

function RouteComponent() {
  const { chatId } = Route.useParams()
  return <Chat chatId={chatId} />
}
