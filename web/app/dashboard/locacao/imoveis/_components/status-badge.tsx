import { Badge } from "@/components/ui/badge"
import { cn } from "@/lib/utils"
import { Property } from "@/types/property"
import { CircleCheck, CircleMinus, CircleX, Hammer, LockKeyhole } from "lucide-react"

type StatusBadgeProps = Pick<Property, 'status'>

export function colorStatusDetail(status: string) {
  switch (status) {
    case "Disponível":
      return "emerald-500"
    case "Ocupado":
      return "yellow-500"
    case "Indisponível":
      return "red-500"
    case "Reformando":
      return "orange-500"
    case "Reservado":
      return "sky-500"
  }
}

export function StatusBadge({ status }: StatusBadgeProps) {
  switch (status) {
    case "Disponível":
      return <Badge variant="outline" className="h-8 gap-2">
        <CircleCheck className={cn(["!size-4", `text-${colorStatusDetail(status)}`])} />{status}
      </Badge >
    case "Ocupado":
      return <Badge variant="outline" className="h-8 gap-2">
        <CircleMinus className={cn(["!size-4", `text-${colorStatusDetail(status)}`])} />{status}
      </Badge >
    case "Indisponível":
      return <Badge variant="outline" className="h-8 gap-2">
        <CircleX className={cn(["!size-4", `text-${colorStatusDetail(status)}`])} />{status}
      </Badge >
    case "Reservado":
      return <Badge variant="outline" className="h-8 gap-2">
        <LockKeyhole className={cn(["!size-4", `text-${colorStatusDetail(status)}`])} />{status}
      </Badge >
    case "Reformando":
      return <Badge variant="outline" className="h-8 gap-2">
        <Hammer className={cn(["!size-4", `text-${colorStatusDetail(status)}`])} />{status}
      </Badge >
  }
}