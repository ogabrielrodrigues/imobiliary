import { useSortable } from "@dnd-kit/sortable"
import { Button } from "./ui/button"
import { GripVertical } from "lucide-react"

export function DragHandle({ id }: { id: number }) {
  const { attributes, listeners } = useSortable({
    id,
  })

  return (
    <Button
      {...attributes}
      {...listeners}
      variant="ghost"
      size="icon"
      className="text-muted-foreground size-7 hover:bg-transparent"
    >
      <GripVertical className="text-muted-foreground size-3" />
      <span className="sr-only">Arrastar para re-ordenar</span>
    </Button>
  )
}