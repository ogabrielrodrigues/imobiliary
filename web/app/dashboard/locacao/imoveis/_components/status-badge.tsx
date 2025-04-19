import { Badge } from "@/components/ui/badge"
import { Property } from "@/types/property"
import { CircleCheck, CircleMinus, CircleX, Hammer, LockKeyhole } from "lucide-react"

type StatusBadgeProps = Pick<Property, 'status'>

export function StatusBadge({ status }: StatusBadgeProps) {
  switch (status?.toUpperCase()) {
    case "DISPONÍVEL":
      return <Badge variant="outline" className="h-8 gap-2">
        <CircleCheck className="!size-4 text-green-500" />{status}
      </Badge >
    case "OCUPADO":
      return <Badge variant="outline" className="h-8 gap-2">
        <CircleMinus className="!size-4 text-yellow-500" />{status}
      </Badge >
    case "INDISPONÍVEL":
      return <Badge variant="outline" className="h-8 gap-2">
        <CircleX className="!size-4 text-red-500" />{status}
      </Badge >
    case "RESERVADO":
      return <Badge variant="outline" className="h-8 gap-2">
        <LockKeyhole className="!size-4" />{status}
      </Badge >
    case "REFORMANDO":
      return <Badge variant="outline" className="h-8 gap-2">
        <Hammer className="!size-4" />{status}
      </Badge >
  }
}